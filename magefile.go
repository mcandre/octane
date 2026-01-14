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

// imageXgo denotes a Docker image for building this project.
var imageXgo = "n4jm4/octane-builder"

// DockerBuildXgo creates local Docker buildx xgo images.
func DockerBuildXgo() error {
	return mageextras.Tuggy(
		"-c", "tuggy.xgo.toml",
		"-t", imageXgo,
		"-f", "xgo.Dockerfile",
		"--load",
	)
}

// DockerPushXgo creates and tag aliases remote Docker buildx xgo images.
func DockerPushXgo() error {
	return mageextras.Tuggy(
		"-c", "tuggy.xgo.toml",
		"-t", imageXgo,
		"-f", "xgo.Dockerfile",
		"--push",
	)
}

// DockerTestXgo creates and tag aliases remote test Docker buildx xgo images.
func DockerTestXgo() error {
	if err := mageextras.Tuggy("-c", "tuggy.xgo.toml", "-t", fmt.Sprintf("%s:test", imageXgo), "-f", "xgo.Dockerfile", "--load"); err != nil {
		return err
	}

	return mageextras.Tuggy(
		"-c", "tuggy.xgo.toml",
		"-t", fmt.Sprintf("%s:test", imageXgo),
		"-f", "xgo.Dockerfile",
		"--push",
	)
}

// imageApp denotes a Docker image for running this project.
var imageApp = "n4jm4/octane"

// DockerBuildApp creates Docker buildx images.
func DockerBuildApp() error {
	return mageextras.Tuggy(
		"-c", "tuggy.app.toml",
		"-t", imageApp,
		"-f", "app.Dockerfile",
		"--load",
	)
}

// DockerPushApp creates and tag aliases remote Docker buildx app images.
func DockerPushApp() error {
	return mageextras.Tuggy(
		"-c", "tuggy.app.toml",
		"-t", imageApp,
		"-f", "app.Dockerfile",
		"-a", fmt.Sprintf("%s:%s", imageApp, octane.Version),
		"--push",
	)
}

// DockerTestApp creates and tag aliases remote test Docker buildx app images.
func DockerTestApp() error {
	if err := mageextras.Tuggy("-c", "tuggy.app.toml", "-t", fmt.Sprintf("%s:test", imageApp), "-f", "app.Dockerfile", "--load"); err != nil {
		return err
	}

	return mageextras.Tuggy(
		"-c", "tuggy.app.toml",
		"-t", fmt.Sprintf("%s:test", imageApp),
		"-f", "app.Dockerfile",
		"--push",
	)
}

// DockerBuild creates buildx images.
func DockerBuild() error {
	mg.Deps(DockerBuildXgo)
	return DockerBuildApp()
}

// DockerPush creates and tags remote Docker buildx images.
func DockerPush() error {
	mg.Deps(DockerPushXgo)
	return DockerPushApp()
}

// DockerTest creates and tag aliases remote test Docker buildx images.
func DockerTest() error {
	mg.Deps(DockerTestXgo)
	return DockerTestApp()
}

// Govulncheck runs govulncheck.
func Govulncheck() error { return mageextras.Govulncheck("-scan", "package", "./...") }

// DockerScout runs a Docker security audit.
func DockerScout() error {
	mg.Deps(DockerBuildXgo)
	mg.Deps(DockerBuildApp)

	if err := mageextras.DockerScout("-e", imageXgo); err != nil {
		return err
	}

	return mageextras.DockerScout("-e", imageApp)
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
	mg.Deps(DockerBuildXgo)

	return mageextras.Xgo(
		artifactsPathDist,
		"-image",
		imageXgo,
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
