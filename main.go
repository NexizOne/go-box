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
}

func main() {
	cmd := &cli.Command{
		Usage:   "Box of simple utils",
		Version: fmt.Sprintf("%s %s", internal.Version, internal.Revision),
		Commands: []*cli.Command{
			pkg.CommandMkdir,
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
