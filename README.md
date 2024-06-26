# DevPod Provider for Nomad

Author: Brian Cain

[![Go](https://github.com/briancain/devpod-provider-nomad/actions/workflows/go.yml/badge.svg)](https://github.com/briancain/devpod-provider-nomad/actions/workflows/go.yml) [![Release](https://github.com/briancain/devpod-provider-nomad/actions/workflows/release.yml/badge.svg)](https://github.com/briancain/devpod-provider-nomad/actions/workflows/release.yml)

This is a provider for [DevPod](https://devpod.sh/) that allows you to create a
DevPod using [HashiCorp Nomad](https://www.nomadproject.io/).

Please report any issues or feature requests to the
[Github Issues](https://github.com/briancain/devpod-provider-nomad/issues) page.

This project is still a work in progress, excuse our mess! <3

[devpod.sh](https://devpod.sh/)

## Getting Started

1. Install the provider to your local machine

From Github:

```shell
devpod provider add briancain/devpod-provider-nomad
```

2. Use the provider

```shell
devpod up <repository-url> --provider nomad
```

### Provider Configurations

Set this options through DevPod to configure them when DevPod launches the
Nomad job during a workspace creation.

- NOMAD_NAMESPACE:
  + description: The namespace for the Nomad job
  + default:
- NOMAD_REGION:
  + description: The region for the Nomad job
  + default:
- NOMAD_CPU:
  + description: The cpu in mhz to use for the Nomad Job
  + default: "200"
- NOMAD_MEMORYMB:
  + description: The memory in mb to use for the Nomad Job
  + default: "512"

## Testing Locally

1. Build the provider locally

```shell
RELEASE_VERSION=0.0.1-dev ./hack/build.sh --dev
```

2. Delete the old provider from DevPod

```shell
devpod provider delete nomad
```

3. Install the new provider from a local build

```shell
devpod provider add --name nomad --use ./release/provider.yaml 
```

4. Test the provider

```shell
devpod up <repository-url> --provider nomad --debug 
```
