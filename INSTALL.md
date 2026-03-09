# INSTALL

In addition to OS packages, octane also supports alternative installation methods.

# INSTALL (GO REMOTE)

octane is packaged as a Go module.

```sh
go install github.com/mcandre/octane/cmd/octane@latest
```

## Prerequisites

* a [C++](https://isocpp.org/) compiler
* [Go](https://go.dev/)
* FreeBSD users require enabling the `snd_uaudio` driver
* Linux distros require an [ALSA](https://www.alsa-project.org/wiki/Main_Page) driver

# INSTALL (GO LOCAL)

kirill may be compiled from source.

```sh
git clone https://github.com/mcandre/octane.git
cd octane
go install ./...
```

## Prerequisites

* a [C++](https://isocpp.org/) compiler
* [git](https://git-scm.com/)
* [Go](https://go.dev/)
* FreeBSD users require enabling the `snd_uaudio` driver
* Linux distros require an [ALSA](https://www.alsa-project.org/wiki/Main_Page) driver

For more details on developing octane, see our [development guide](DEVELOPMENT.md).
