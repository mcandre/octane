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
var portBasename = "octane"

// artifactsPathDist is the parent directory of xgo artifacts.
var artifactsPathDist = path.Join(artifactsPath, portBasename)

// repoNamespace identifies the Go namespace for this project.
var repoNamespace = "github.com/mcandre/octane"

// imageXgo denotes a Docker image for building this project.
var imageXgo = "n4jm4/octane-xgo"

// Clean deletes build artifacts.
func Clean() error { mg.Deps(CleanBin); return CleanPackages() }

// CleanBin deletes Go artifacts.
func CleanBin() error { return os.RemoveAll(artifactsPath) }

// CleanPackages deletes OS package artifacts.
func CleanPackages() error { return os.RemoveAll(".rockhopper") }

// Audit runs security audits.
func Audit() error {
	mg.Deps(Govulncheck)
	return DockerScout()
}

// Deadcode runs deadcode.
func Deadcode() error { return mageextras.Deadcode("./...") }

// DockerBuild generates Docker images.
func DockerBuild() error {
	cmd := exec.Command("docker", "buildx", "bake", "all")
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// DockerPush pushes Docker images.
func DockerPush() error {
	cmd := exec.Command("docker", "buildx", "bake", "production", "--push")
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// DockerScout runs docker scout scans.
func DockerScout() error {
	if err := mageextras.DockerScout("-e", imageXgo); err != nil {
		return err
	}

	cmd := exec.Command("docker", "scout", "cves", "-e", "fs://.")
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// DockerTest tests pushing Docker images.
func DockerTest() error {
	cmd := exec.Command("docker", "buildx", "bake", "test", "--push")
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// GoFix runs go fix.
func GoFix() error { return mageextras.GoFix("./...") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// GoVet runs default go vet analyzers.
func GoVet() error { return mageextras.GoVet() }

// Govulncheck runs govulncheck.
func Govulncheck() error { return mageextras.Govulncheck("-scan", "package", "./...") }

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(Deadcode)
	mg.Deps(GoFix)
	mg.Deps(GoImports)
	mg.Deps(GoVet)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	mg.Deps(Shadow)
	mg.Deps(Staticcheck)
	return nil
}

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Package generates OS packages.
func Package() error {
	cmd := exec.Command("rockhopper", "-r", fmt.Sprintf("version=%s", octane.Version))
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// Shadow runs go vet with shadow checks enabled.
func Shadow() error { return mageextras.GoVetShadow() }

// Staticcheck runs staticcheck.
func Staticcheck() error { return mageextras.Staticcheck("./...") }

// Test runs a test suite.
func Test() error { return mageextras.UnitTest() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("octane") }

// Upload copies packages to CloudFlare R2.
func Upload() error {
	cmd := exec.Command("./upload")
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// Xgo cross-compiles (c)Go binaries with additional targets enabled.
func Xgo() error {
	return mageextras.Xgo(
		artifactsPathDist,
		"-image",
		imageXgo,
		"-targets",
		"darwin/amd64,darwin/arm64,freebsd/amd64,linux/amd64,linux/arm64,windows/amd64,windows/arm64",
		".",
	)
}
