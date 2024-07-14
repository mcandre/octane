# BUILDTIME REQUIREMENTS

* a [C++](https://en.wikipedia.org/wiki/List_of_compilers#C++_compilers) compiler
* [Docker](https://www.docker.com/) 19+
* [Go](https://go.dev/) 1.22.5+
* POSIX compatible [make](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/make.html)
* [Node.js](https://nodejs.org/en) 20.10.0+
* [Rust](https://www.rust-lang.org/) 1.75.0+
* POSIX compatible [tar](https://pubs.opengroup.org/onlinepubs/7908799/xcu/tar.html)
* Provision additional dev tools with `make`

## Recommended

* [ASDF](https://asdf-vm.com/) 0.10 (run `asdf reshim` after provisioning)
* [direnv](https://direnv.net/) 2

## Linux

Linux build environments have additional requirements.

* [ALSA](https://alsa-project.org/wiki/Main_Page) development headers (Debian: `libasound2-dev`, RHEL: `alsa-lib-devel`, Alpine: `alsa-lib-dev`, etc.)

# AUDIT

```console
$ mage audit
```

# INSTALL

```console
$ mage install
```

# UNINSTALL

```console
$ mage uninstall
```

# LINT

Keep the code tidy:

```console
$ mage lint
```

# BUILD + PUSH DOCKER IMAGE

```console
$ mage dockerBuild
$ mage dockerPush
```

# PORT

```console
$ mage port
```
