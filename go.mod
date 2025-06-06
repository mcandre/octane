module github.com/mcandre/octane

go 1.24.2

require (
	github.com/magefile/mage v1.15.0
	github.com/mcandre/mage-extras v0.0.21
	gitlab.com/gomidi/midi/v2 v2.0.28
)

require (
	github.com/BurntSushi/toml v1.4.1-0.20240526193622-a339e1f7089c // indirect
	github.com/alexkohler/nakedret/v2 v2.0.5 // indirect
	github.com/chavacava/garif v0.1.0 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/kisielk/errcheck v1.9.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mcandre/factorio v0.0.10 // indirect
	github.com/mcandre/tug v0.0.12 // indirect
	github.com/mgechev/dots v0.0.0-20210922191527-e955255bf517 // indirect
	github.com/mgechev/revive v1.7.0 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/spf13/afero v1.12.0 // indirect
	golang.org/x/exp/typeparams v0.0.0-20231108232855-2478ac86f678 // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	golang.org/x/tools v0.30.0 // indirect
	honnef.co/go/tools v0.6.0 // indirect
	src.techknowlogick.com/xgo v1.8.1-0.20250211184208-9779ef77bfbb // indirect
)

tool (
	github.com/alexkohler/nakedret/v2/cmd/nakedret
	github.com/kisielk/errcheck
	github.com/magefile/mage
	github.com/mcandre/factorio/cmd/factorio
	github.com/mcandre/tug/cmd/tug
	github.com/mgechev/revive
	honnef.co/go/tools/cmd/staticcheck
	src.techknowlogick.com/xgo
)
