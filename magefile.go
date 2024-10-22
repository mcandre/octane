//go:build mage
package main

import (
	"fmt"
	"os"
	"os/exec"
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

// repoNamespace identifies the Go namespace for this project.
var repoNamespace = "github.com/mcandre/octane"

// image denotes a Docker image for building this project.
var image = "mcandre/octane-builder"

// DockerBuild generates Docker images.
func DockerBuild() error {
	cmd := exec.Command("tug")
	cmd.Args = []string{
		"tug",
		"-t",
		image,
		"-exclude-arch",
		"386,arm/v6,arm/v7,arm64,mips64le,ppc64le,riscv64,s390x",
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// DockerPush registers Docker images.
func DockerPush() error {
	cmd := exec.Command("tug")
	cmd.Args = []string{
		"tug",
		"-t",
		image,
		"-exclude-arch",
		"386,arm/v6,arm/v7,arm64,mips64le,ppc64le,riscv64,s390x",
		"-push",
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// DockerLoad loads Docker images of a given platform.
func DockerLoad(platform string) error {
	cmd := exec.Command("tug")
	cmd.Args = []string{
		"tug",
		"-t",
		image,
		"-load",
		platform,
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Govulncheck runs govulncheck.
func Govulncheck() error { return mageextras.Govulncheck("-scan", "package", "./...") }

// Snyk runs Snyk SCA.
func Snyk() error { return mageextras.SnykTest() }

// DockerScout runs a Docker security audit.
func DockerScout() error {
	if err := DockerLoad("linux/amd64"); err != nil {
		return err
	}

	return mageextras.DockerScout("-e", "mcandre/octane-builder")
}

// Audit runs security audits.
func Audit() error {
	mg.Deps(Govulncheck)
	mg.Deps(Snyk)
	return DockerScout()
}

// Deadcode runs deadcode.
func Deadcode() error { return mageextras.Deadcode("./...") }

// Gofmt runs gofmt.
func GoFmt() error { return mageextras.GoFmt("-s", "-w") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// GoVet runs default go vet analyzers.
func GoVet() error { return mageextras.GoVet() }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Revive runs revive.
func Revive() error { return mageextras.Revive() }

// Shadow runs go vet with shadow checks enabled.
func Shadow() error { return mageextras.GoVetShadow() }

// Staticcheck runs staticcheck.
func Staticcheck() error { return mageextras.Staticcheck() }

// Unmake runs unmake.
func Unmake() error {
	err := mageextras.Unmake(".")

	if err != nil {
		return err
	}

	return mageextras.Unmake("-n", ".")
}

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(Deadcode)
	mg.Deps(GoFmt)
	mg.Deps(GoImports)
	mg.Deps(GoVet)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	mg.Deps(Revive)
	mg.Deps(Shadow)
	mg.Deps(Staticcheck)
	mg.Deps(Unmake)
	return nil
}

// Xgo cross-compiles (c)Go binaries with additional targets enabled.
func Xgo() error {
	err := DockerLoad("linux/amd64")

	if err != nil {
		return err
	}

	artifactsPathDist := path.Join(artifactsPath, portBasename)

	return mageextras.Xgo(
		artifactsPathDist,
		"-image",
		image,
		"-targets",
		"darwin/amd64,darwin/arm64,linux/amd64,windows/amd64",
		".",
	)
}

// Port builds and compresses artifacts.
func Port() error { mg.Deps(Xgo); return mageextras.Archive(portBasename, artifactsPath) }

// Test runs a test suite.
func Test() error { return mageextras.UnitTest() }

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("octane") }

// Clean deletes artifacts.
func Clean() error { return os.RemoveAll(artifactsPath) }
