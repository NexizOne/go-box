package internal

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v3"
)

// TODO https://cli.urfave.org/v3/examples/full-api-example/
func FlagStringer(fl cli.Flag) string {
	flagsKeys := fmt.Sprintf("--%s", strings.Join(fl.Names(), ",\t-"))
	df, ok := fl.(cli.DocGenerationFlag)
	if !ok {
		return flagsKeys
	}
	defText := ""
	if def := df.GetDefaultText(); len(def) > 0 {
		if fl == cli.HelpFlag || fl == cli.VersionFlag {
			defText = ""
		} else {
			defText = fmt.Sprintf(" (default: %s)", def)
		}
	}

	return fmt.Sprintf("%s\t%s%s", flagsKeys, df.GetUsage(), defText)
}

var RootCommandHelpTemplate = `NAME:
   {{template "helpNameTemplate" .}}

USAGE:
   {{if .UsageText}}{{wrap .UsageText 3}}{{else}}{{.FullName}}{{if .ArgsUsage}} {{.ArgsUsage}}{{else}}{{if .Arguments}} [arguments...]{{end}}{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}

VERSION:
   {{.Version}}{{end}}{{end}}{{if .Description}}

DESCRIPTION:
   {{template "descriptionTemplate" .}}{{end}}
{{- if len .Authors}}

AUTHOR{{template "authorsTemplate" .}}{{end}}{{if .VisibleCommands}}

COMMANDS:{{template "visibleCommandCategoryTemplate" .}}{{end}}{{if .VisibleFlagCategories}}

GLOBAL OPTIONS:{{template "visibleFlagCategoryTemplate" .}}{{else if .VisibleFlags}}

GLOBAL OPTIONS:{{template "visibleFlagTemplate" .}}{{end}}{{if .Copyright}}

COPYRIGHT:
   {{template "copyrightTemplate" .}}{{end}}
`
