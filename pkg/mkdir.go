package pkg

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/NexizOne/go-box/internal"
	"github.com/urfave/cli/v3"
)

const (
	// arguments
	mkFlagMode  = "mode"
	mkAliasMode = "m"
)

var CommandMkdir *cli.Command = &cli.Command{
	Name:                  internal.CmdMkdir,
	Version:               internal.Version,
	Usage:                 "Creates directory",
	ArgsUsage:             "[directory]",
	EnableShellCompletion: true,
	Action:                mkdirAction,
	Flags: []cli.Flag{
		&cli.UintFlag{
			Name:    mkFlagMode,
			Aliases: []string{mkAliasMode},
			Value:   777,
			Usage:   "mode",
		},
	},
}

func mkdirAction(ctx context.Context, cmd *cli.Command) error {

	if cmd.Args().Len() < 1 {
		return cli.Exit("path not provided", 1)
	}

	path := cmd.Args().First()
	mode := os.ModePerm
	modeString := fmt.Sprint(cmd.Uint(mkFlagMode))
	if modeOcta, err := strconv.ParseUint(modeString, 8, 32); err == nil {
		mode = os.FileMode(modeOcta)
	} else {
		fmt.Println(err)
	}

	if absPath, err := MakeDir(path, mode); err != nil {
		return cli.Exit(err, 1)
	} else if absPath != nil {
		fmt.Printf("%s\n", *absPath)
	}
	return nil
}

func MakeDir(path string, perm os.FileMode) (*string, error) {
	var err error
	var absPath string

	if absPath, err = filepath.Abs(path); err != nil {
		return nil, err
	}

	if _, err = os.Stat(absPath); !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	if err = os.MkdirAll(absPath, perm); err != nil {
		return nil, err
	}

	return &absPath, nil
}
