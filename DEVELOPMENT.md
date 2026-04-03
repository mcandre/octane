# DEVELOPMENT

We follow standard, `go` based operations for compiling and unit testing Go code.

For advanced operations, such as linting, we further supplement with some software industry tools.

# DEV ENVIRONMENT

## Prerequisites

* a UNIX-like environment (e.g. [WSL](https://learn.microsoft.com/en-us/windows/wsl/))
* [awscli](https://aws.amazon.com/cli/)
* a [C++](https://isocpp.org/) compiler
* [Docker](https://www.docker.com/)
* [Go](https://go.dev/)
* [make](https://pubs.opengroup.org/onlinepubs/9799919799/utilities/make.html)
* [Rust](https://rust-lang.org/)
* FreeBSD users require enabling the `snd_uaudio` driver
* Linux distros require an [ALSA](https://www.alsa-project.org/wiki/Main_Page) driver
* Provision additional dev tools with `make`

## Recommended

* [asdf](https://asdf-vm.com/)

## Postinstall

Register output of `go env GOBIN` to `PATH` environment variable.

Register `~/.cargo/bin` to `PATH` environment variable.

# TASKS

We automate engineering tasks.

## Build

```sh
mage
```

## Install

```sh
mage install
```

## Uninstall

```sh
mage uninstall
```

## Security Audit

```sh
mage audit
```

## Lint

```sh
mage lint
```

## Test

```sh
mage test
```

## Build Images

```sh
mage dockerBuild
```

## Test Push Images

```sh
mage dockerTest
```

## Push Images

```sh
mage dockerPush
```

## Crosscompile Binaries

```sh
mage xgo
```

## Package Binaries

```sh
mage package
```

## Upload Packages

```sh
mage upload
```

## Clean Workspace

```sh
mage clean
```
