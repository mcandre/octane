# DEVELOPMENT GUIDE

octane follows standard, Go based operations for compiling and unit testing Go code.

For advanced operations, such as linting, managing multiplatform Docker images, and so on, we further supplement with some software industry tools.

# BUILDTIME REQUIREMENTS

* a UNIX-like environment (e.g. [WSL](https://learn.microsoft.com/en-us/windows/wsl/))
* [awscli](https://aws.amazon.com/cli/)
* a [C++](https://isocpp.org/) compiler
* [Docker](https://www.docker.com/)
* [Go](https://go.dev/)
* POSIX compliant [make](https://pubs.opengroup.org/onlinepubs/9799919799/utilities/make.html)
* [Rust](https://rust-lang.org/)
* FreeBSD users require enabling the `snd_uaudio` driver
* Linux distros require an [ALSA](https://www.alsa-project.org/wiki/Main_Page) driver
* Provision additional dev tools with `make`

## Recommended

* [ASDF](https://asdf-vm.com/) 0.18 (run `asdf reshim` after provisioning)

# AUDIT

```sh
mage audit
```

# INSTALL

```sh
mage install
```

# UNINSTALL

```sh
mage uninstall
```

# LINT

Keep the code tidy:

```sh
mage lint
```

# TEST

```sh
mage [test]
```

# BUILD DOCKER IMAGES

```sh
mage dockerBuild
```

# TEST PUSH DOCKER IMAGES

```sh
mage dockerTest
```

# PUSH DOCKER IMAGES

```sh
mage dockerPush
```

# CROSSCOMPILE BINARIES

```sh
mage xgo
```

# PACKAGE BINARIES

```sh
mage package
```

# UPLOAD BINARIES

```sh
mage upload
```
