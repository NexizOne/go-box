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
	FlagMode  = "mode"
	AliasMode = "m"

	// algorithms
	// AlgSha256 = "sha256"
	// AlgSha1   = "sha1"
	// AlgMd5    = "md5"
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
			Name:    FlagMode,
			Aliases: []string{AliasMode},
			Value:   777,
			Usage:   "mode",
		},
	},
}

func mkdirAction(ctx context.Context, cmd *cli.Command) error {

	path := cmd.Args().First()
	if cmd.Args().Len() < 1 {
		return cli.Exit("path not provided", 1)
	}

	mode := os.ModePerm
	modeString := fmt.Sprint(cmd.Uint(FlagMode))
	if modeOcta, err := strconv.ParseUint(modeString, 8, 32); err == nil {
		mode = os.FileMode(modeOcta)
	} else {
		fmt.Println(err)
	}

	if absPath, err := MakeDir(path, mode); err != nil {
		return cli.Exit(err, 1)
	} else if absPath != nil {
		fmt.Printf("%s\n", *absPath)
		// fmt.Printf("%04o\n%s\n", mode, mode.String())
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
