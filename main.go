package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/NexizOne/go-box/internal"
	"github.com/NexizOne/go-box/pkg"
	"github.com/urfave/cli/v3"
)

func init() {
	cli.FlagStringer = internal.FlagStringer
	cli.RootCommandHelpTemplate = internal.RootCommandHelpTemplate
}

func main() {
	cmd := &cli.Command{
		Usage:     "Box of simple utils",
		Version:   fmt.Sprintf("%s %s", internal.Version, internal.Revision),
		ArgsUsage: "command [command options]",
		Commands: []*cli.Command{
			pkg.CommandCp,
			pkg.CommandMkdir,
			pkg.CommandMv,
			pkg.CommandHash,
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return nil
		},
	}

	cmd.Name = internal.Basename(os.Args[0])

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
