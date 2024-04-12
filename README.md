# Devpod Provider for Nomad

Author: Brian Cain

[devpod.sh](https://devpod.sh/)

## Getting Started

TODO

## Testing Locally

1. Build the provider locally

```shell
RELEASE_VERSION=0.0.1-dev ./hack/build.sh --dev
```

2. Delete the old provider from devpod

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
