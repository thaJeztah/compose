# docker compose build

<!---MARKER_GEN_START-->
Services are built once and then tagged, by default as `project-service`.

If the Compose file specifies an
[image](https://github.com/compose-spec/compose-spec/blob/main/spec.md#image) name,
the image is tagged with that name, substituting any variables beforehand. See
[variable interpolation](https://github.com/compose-spec/compose-spec/blob/main/spec.md#interpolation).

If you change a service's `Dockerfile` or the contents of its build directory,
run `docker compose build` to rebuild it.

### Options

| Name                  | Type          | Default | Description                                                                                                 |
|:----------------------|:--------------|:--------|:------------------------------------------------------------------------------------------------------------|
| `--build-arg`         | `stringArray` |         | Set build-time variables for services                                                                       |
| `--builder`           | `string`      |         | Set builder to use                                                                                          |
| `--check`             | `bool`        |         | Check build configuration                                                                                   |
| `--dry-run`           | `bool`        |         | Execute command in dry run mode                                                                             |
| `-m`, `--memory`      | `bytes`       | `0`     | Set memory limit for the build container. Not supported by BuildKit.                                        |
| `--no-cache`          | `bool`        |         | Do not use cache when building the image                                                                    |
| `--print`             | `bool`        |         | Print equivalent bake file                                                                                  |
| `--provenance`        | `string`      |         | Add a provenance attestation                                                                                |
| `--pull`              | `bool`        |         | Always attempt to pull a newer version of the image                                                         |
| `--push`              | `bool`        |         | Push service images                                                                                         |
| `-q`, `--quiet`       | `bool`        |         | Suppress the build output                                                                                   |
| `--sbom`              | `string`      |         | Add a SBOM attestation                                                                                      |
| `--ssh`               | `string`      |         | Set SSH authentications used when building service images. (use 'default' for using your default SSH Agent) |
| `--with-dependencies` | `bool`        |         | Also build dependencies (transitively)                                                                      |


<!---MARKER_GEN_END-->

## Description

Services are built once and then tagged, by default as `project-service`.

If the Compose file specifies an
[image](https://github.com/compose-spec/compose-spec/blob/main/spec.md#image) name,
the image is tagged with that name, substituting any variables beforehand. See
[variable interpolation](https://github.com/compose-spec/compose-spec/blob/main/spec.md#interpolation).

If you change a service's `Dockerfile` or the contents of its build directory,
run `docker compose build` to rebuild it.
