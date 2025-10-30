package pkg

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/NexizOne/go-box/internal"
	"github.com/urfave/cli/v3"
)

const (
	// arguments
	mvFlagMode  = "force"
	mvAliasMode = "f"
)

// TODO https://man7.org/linux/man-pages/man1/mv.1.html
var CommandMv *cli.Command = &cli.Command{
	Name:                  internal.CmdMv,
	Version:               internal.Version,
	Usage:                 "Mode file or directory",
	ArgsUsage:             "[from] [to] [options]",
	Description:           "from\t\tsource (masks supported, example: *.txt)\nto\t\tdestination",
	CustomHelpTemplate:    cli.SubcommandHelpTemplate,
	EnableShellCompletion: true,
	Action:                mvAction,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    mvFlagMode,
			Aliases: []string{mvAliasMode},
			Value:   false,
			Usage:   "force",
		},
	},
}

func mvAction(ctx context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 2 {
		return cli.Exit("argument count is incorrect required 2", 1)
	}

	var from = cmd.Args().Get(0)
	var to = cmd.Args().Get(1)
	var overwrite = cmd.Bool(mvFlagMode)

	files, err := filepath.Glob(from)
	if err != nil {
		return cli.Exit(err, 1)
	}

	if len(files) > 0 {
		info, err := os.Stat(to)
		if os.IsNotExist(err) {
			if err = os.MkdirAll(to, os.ModePerm); err != nil {
				return cli.Exit(err, 1)
			}
		} else {
			if !info.IsDir() {
				return cli.Exit(fmt.Sprintf("destination \"%s\" is not a directory", to), 1)
			}
		}
	} else {
		absFrom, err := filepath.Abs(from)
		if err != nil {
			return cli.Exit(fmt.Sprintf("no files in %s\n", from), 1)
		}
		return cli.Exit(fmt.Sprintf("no files in %s\n", absFrom), 1)
	}

	for _, fromFile := range files {
		filename := filepath.Base(fromFile)
		toFile := filepath.Join(to, filename)

		toInfo, err := os.Stat(to)
		if os.IsNotExist(err) {
			if err := os.Rename(fromFile, toFile); err != nil {
				fmt.Println(err)
				continue
			}
		} else if toInfo.IsDir() {
			if overwrite == false && ifFileExists(toFile) {
				fmt.Printf(fmt.Sprintf("file \"%s\" exists\n", toFile))
				continue
			}

			if err := os.Rename(fromFile, toFile); err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			if overwrite == false {
				fmt.Printf(fmt.Sprintf("file \"%s\" exists\n", toFile))
				continue
			}

			if err := os.Rename(fromFile, to); err != nil {
				fmt.Println(err)
				continue
			}
			toFile = to
		}
		fmt.Printf("move from: %s to: %s\n", fromFile, toFile)
	}

	return nil
}

func ifFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
