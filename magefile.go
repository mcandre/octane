//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/magefile/mage/mg"
	"github.com/mcandre/mage-extras"
	"github.com/mcandre/octane"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "bin"

// Default references the default build task.
var Default = Install

// Audit runs a security audit.
func Audit() error { return mageextras.SnykTest() }

// GoVet runs go vet with shadow checks enabled.
func GoVet() error { return mageextras.GoVetShadow() }

// GoLint runs golint.
func GoLint() error { return mageextras.GoLint() }

// Gofmt runs gofmt.
func GoFmt() error { return mageextras.GoFmt("-s", "-w") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Staticcheck runs staticcheck.
func Staticcheck() error { return mageextras.Staticcheck() }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(GoVet)
	mg.Deps(GoLint)
	mg.Deps(GoFmt)
	mg.Deps(GoImports)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	mg.Deps(Staticcheck)
	return nil
}

// portBasename labels the artifact basename.
var portBasename = fmt.Sprintf("octane-%s", octane.Version)

// repoNamespace identifies the Go namespace for this project.
var repoNamespace = "github.com/mcandre/octane"

// image denotes a Docker image for building this project.
var image = "mcandre/octane-builder"

// DockerBuild generates a Docker image.
func DockerBuild() error {
	cmd := exec.Command("docker")
	cmd.Args = []string{
		"docker",
		"build",
		"-t",
		image,
		".",
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// DockerPush registers a Docker image.
func DockerPush() error {
	cmd := exec.Command("docker")
	cmd.Args = []string{
		"docker",
		"push",
		image,
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Xgo cross-compiles (c)Go binaries with additional targets enabled.
func Xgo() error {
	artifactsPathDist := path.Join(artifactsPath, portBasename)

	return mageextras.Xgo(
		artifactsPathDist,
		"-docker-image",
		image,
		"-targets",
		"darwin/amd64,darwin/arm64,linux/amd64,windows/amd64",
		"github.com/mcandre/octane/cmd/octane",
	)
}

// Port builds and compresses artifacts.
func Port() error { mg.Deps(Xgo); return mageextras.Archive(portBasename, artifactsPath) }

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("octane") }

// Clean deletes artifacts.
func Clean() error { return os.RemoveAll(artifactsPath) }
