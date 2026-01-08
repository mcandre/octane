//go:build mage

package main

import (
	"fmt"
	"os"
	"path"

	"github.com/magefile/mage/mg"
	mageextras "github.com/mcandre/mage-extras"
	"github.com/mcandre/octane"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "bin"

// Default references the default build task.
var Default = Test

// portBasename labels the artifact basename.
var portBasename = fmt.Sprintf("octane-%s", octane.Version)

// artifactsPathDist is the parent directory of xgo artifacts.
var artifactsPathDist = path.Join(artifactsPath, portBasename)

// repoNamespace identifies the Go namespace for this project.
var repoNamespace = "github.com/mcandre/octane"

// image denotes a Docker image for building this project.
var image = "n4jm4/octane-builder"

// DockerBuild creates local Docker buildx images.
func DockerBuild() error {
	return mageextras.Tuggy(
		"-t", image,
		"--load",
	)
}

// DockerPush creates and tag aliases remote Docker buildx images.
func DockerPush() error {
	return mageextras.Tuggy(
		"-t", image,
		"--push",
	)
}

// DockerTest creates and tag aliases remote test Docker buildx images.
func DockerTest() error {
	if err := mageextras.Tuggy("-t", fmt.Sprintf("%s:test", image), "--load"); err != nil {
		return err
	}

	return mageextras.Tuggy(
		"-t", fmt.Sprintf("%s:test", image),
		"--push",
	)
}

// Govulncheck runs govulncheck.
func Govulncheck() error { return mageextras.Govulncheck("-scan", "package", "./...") }

// DockerScout runs a Docker security audit.
func DockerScout() error {
	mg.Deps(DockerBuild)
	return mageextras.DockerScout("-e", image)
}

// Audit runs security audits.
func Audit() error {
	mg.Deps(Govulncheck)
	return DockerScout()
}

// Deadcode runs deadcode.
func Deadcode() error { return mageextras.Deadcode("./...") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// GoVet runs default go vet analyzers.
func GoVet() error { return mageextras.GoVet() }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Shadow runs go vet with shadow checks enabled.
func Shadow() error { return mageextras.GoVetShadow() }

// Staticcheck runs staticcheck.
func Staticcheck() error { return mageextras.Staticcheck("./...") }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(Deadcode)
	mg.Deps(GoImports)
	mg.Deps(GoVet)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	mg.Deps(Shadow)
	mg.Deps(Staticcheck)
	return nil
}

// Xgo cross-compiles (c)Go binaries with additional targets enabled.
func Xgo() error {
	mg.Deps(DockerBuild)

	return mageextras.Xgo(
		artifactsPathDist,
		"-image",
		image,
		"-targets",
		"darwin/amd64,darwin/arm64,freebsd/amd64,linux/386,linux/amd64,linux/arm,linux/arm64,linux/mips,linux/mips64,linux/mips64le,linux/mipsle,linux/ppc64le,linux/riscv64,linux/s390x,windows/386,windows/amd64",
		".",
	)
}

// Port builds and compresses artifacts.
func Port() error {
	mg.Deps(Xgo)

	return mageextras.Chandler(
		"-C",
		artifactsPath,
		"-czf",
		fmt.Sprintf("%s.tgz", portBasename),
		portBasename,
	)
}

// Test runs a test suite.
func Test() error { return mageextras.UnitTest() }

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("octane") }

// Clean deletes artifacts.
func Clean() error { return os.RemoveAll(artifactsPath) }
