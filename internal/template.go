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
