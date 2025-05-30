/*
   Copyright 2020 Docker Compose CLI authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package compose

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"sync"

	"github.com/distribution/reference"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/errdefs"
	"golang.org/x/sync/errgroup"

	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Images(ctx context.Context, projectName string, options api.ImagesOptions) ([]api.ImageSummary, error) {
	projectName = strings.ToLower(projectName)
	allContainers, err := s.apiClient().ContainerList(ctx, container.ListOptions{
		All:     true,
		Filters: filters.NewArgs(projectFilter(projectName)),
	})
	if err != nil {
		return nil, err
	}
	var containers []container.Summary
	if len(options.Services) > 0 {
		// filter service containers
		for _, c := range allContainers {
			if slices.Contains(options.Services, c.Labels[api.ServiceLabel]) {
				containers = append(containers, c)
			}
		}
	} else {
		containers = allContainers
	}

	images := []string{}
	for _, c := range containers {
		if !slices.Contains(images, c.Image) {
			images = append(images, c.Image)
		}
	}
	imageSummaries, err := s.getImageSummaries(ctx, images)
	if err != nil {
		return nil, err
	}
	summary := make([]api.ImageSummary, len(containers))
	for i, c := range containers {
		img, ok := imageSummaries[c.Image]
		if !ok {
			return nil, fmt.Errorf("failed to retrieve image for container %s", getCanonicalContainerName(c))
		}

		summary[i] = img
		summary[i].ContainerName = getCanonicalContainerName(c)
	}
	return summary, nil
}

func (s *composeService) getImageSummaries(ctx context.Context, repoTags []string) (map[string]api.ImageSummary, error) {
	summary := map[string]api.ImageSummary{}
	l := sync.Mutex{}
	eg, ctx := errgroup.WithContext(ctx)
	for _, repoTag := range repoTags {
		eg.Go(func() error {
			inspect, err := s.apiClient().ImageInspect(ctx, repoTag)
			if err != nil {
				if errdefs.IsNotFound(err) {
					return nil
				}
				return fmt.Errorf("unable to get image '%s': %w", repoTag, err)
			}
			tag := ""
			repository := ""
			ref, err := reference.ParseDockerRef(repoTag)
			if err == nil {
				// ParseDockerRef will reject a local image ID
				repository = reference.FamiliarName(ref)
				if tagged, ok := ref.(reference.Tagged); ok {
					tag = tagged.Tag()
				}
			}
			l.Lock()
			summary[repoTag] = api.ImageSummary{
				ID:          inspect.ID,
				Repository:  repository,
				Tag:         tag,
				Size:        inspect.Size,
				LastTagTime: inspect.Metadata.LastTagTime,
			}
			l.Unlock()
			return nil
		})
	}
	return summary, eg.Wait()
}
