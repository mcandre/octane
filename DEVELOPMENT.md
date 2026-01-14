# BUILDTIME REQUIREMENTS

* a [C++](https://isocpp.org/) compiler
* [Docker](https://www.docker.com/) 27+
* [Go](https://go.dev/)
* POSIX compliant [make](https://pubs.opengroup.org/onlinepubs/9799919799/utilities/make.html)
* [Rust](https://rust-lang.org/)
* Provision additional dev tools with `make`

## Recommended

* a host capable of running musl/Linux containers (e.g. a GNU/Linux, musl/Linux, macOS, or Windows host)
* a UNIX-like environment (e.g. [WSL](https://learn.microsoft.com/en-us/windows/wsl/))
* [ASDF](https://asdf-vm.com/) 0.18 (run `asdf reshim` after provisioning)
* [direnv](https://direnv.net/) 2

## Linux

Linux build environments have additional requirements.

* [ALSA](https://alsa-project.org/wiki/Main_Page) development headers (Debian: `libasound2-dev`, RHEL: `alsa-lib-devel`, Alpine: `alsa-lib-dev`, etc.)

## Windows

Apply a user environment variable `GODEBUG=modcacheunzipinplace=1` per [access denied resolution](https://github.com/golang/go/wiki/Modules/e93463d3e853031af84204dc5d3e2a9a710a7607#go-115), for native Windows development environments (Command Prompt / PowerShell, not WLS, not Cygwin, not MSYS2, not MinGW, not msysGit, not Git Bash, not etc).

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

# PORT

```sh
mage port
```

# TEST DOCKER IMAGES

```sh
mage dockerTest
```

# PUSH DOCKER IMAGES

```sh
mage dockerPush
```
