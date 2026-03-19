# INSTALL

We support alternative installation methods.

# INSTALL (GO)

octane is packaged as a Go module.

```sh
go install github.com/mcandre/octane/cmd/octane@latest
```

## Prerequisites

* a [C++](https://isocpp.org/) compiler
* [Go](https://go.dev/)
* FreeBSD users require enabling the `snd_uaudio` driver
* Linux distros require an [ALSA](https://www.alsa-project.org/wiki/Main_Page) driver

## Postinstall

Register output of `go env GOBIN` to `PATH` environment variable.
