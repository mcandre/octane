.PHONY: all

all:
	@cargo install --force unmake@0.0.3

	@go install github.com/alexkohler/nakedret@v1.0.1
	@go install github.com/crazy-max/xgo@v0.26.0
	@go install github.com/goware/modvendor@v0.5.0
	@go install github.com/kisielk/errcheck@v1.6.3
	@go install github.com/magefile/mage@v1.14.0
	@go install github.com/mcandre/factorio/cmd/factorio@v0.0.3
	@go install golang.org/x/lint/golint@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	@go install honnef.co/go/tools/cmd/staticcheck@2023.1.3
	@go mod tidy
	@modvendor -copy="**/*.h **/*.c **/*.hpp **/*.cpp"

	@npm install -g snyk@1.1140.0
