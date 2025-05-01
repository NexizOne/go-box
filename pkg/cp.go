package pkg

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/NexizOne/go-box/internal"
	"github.com/urfave/cli/v3"
)

const (
	// arguments
	cpFlagMode  = "force"
	cpAliasMode = "f"
)

// TODO https://man7.org/linux/man-pages/man1/cp.1.html
var CommandCp *cli.Command = &cli.Command{
	Name:                  internal.CmdCp,
	Version:               internal.Version,
	Usage:                 "Copy file or directory",
	ArgsUsage:             "[from] [to]",
	EnableShellCompletion: true,
	Action:                cpAction,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    cpFlagMode,
			Aliases: []string{cpAliasMode},
			Value:   false,
			Usage:   "force",
		},
	},
}

func cpAction(ctx context.Context, cmd *cli.Command) error {
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

	if len(files) > 1 {
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
	}

	for _, fromFile := range files {
		filename := filepath.Base(fromFile)
		toFile := filepath.Join(to, filename)

		toInfo, err := os.Stat(to)
		if os.IsNotExist(err) {
			if err := copyFile(fromFile, toFile); err != nil {
				fmt.Println(err)
				continue
			}
		} else if toInfo.IsDir() {
			if overwrite == false && ifFileExists(toFile) {
				fmt.Printf(fmt.Sprintf("file \"%s\" exists\n", toFile))
				continue
			}

			if err := copyFile(fromFile, toFile); err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			if overwrite == false {
				fmt.Printf(fmt.Sprintf("file \"%s\" exists\n", toFile))
				continue
			}

			if err := copyFile(fromFile, to); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}

	return nil
}

func copyFile(src string, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	if err = os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
		return err
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if _, err := io.Copy(destination, source); err != nil {
		return err
	}

	return nil
}
