package main

import (
	"context"
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
	cmd := pkg.CommandCp

	cmd.Name = internal.Basename(os.Args[0])

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
