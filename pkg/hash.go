package pkg

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"os"
	"strings"

	"github.com/NexizOne/go-box/internal"
	"github.com/urfave/cli/v3"
)

const (
	// arguments
	FlagAlgoritm  = "algorithm"
	FlagFile      = "file"
	FlagString    = "string"
	AliasAlgoritm = "a"
	AliasFile     = "f"
	AliasString   = "s"

	// algorithms
	AlgSha256 = "sha256"
	AlgSha1   = "sha1"
	AlgMd5    = "md5"
)

// all algorithms
var AlgAll = []string{AlgSha256, AlgSha1, AlgMd5}

var CommandHash *cli.Command = &cli.Command{
	Name:                  internal.CommandHash,
	Version:               internal.Version,
	Usage:                 "Hash string or file",
	EnableShellCompletion: true,
	ArgsUsage:             "[options]",
	Action:                hashAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    FlagAlgoritm,
			Aliases: []string{AliasAlgoritm},
			Value:   AlgSha256,
			Usage:   fmt.Sprintf("hashing algorithm (%s)", strings.Join(AlgAll, ", ")),
		},
		&cli.StringFlag{
			Name:    FlagFile,
			Aliases: []string{AliasFile},
			Usage:   "file to hash",
		},
		&cli.StringFlag{
			Name:    FlagString,
			Aliases: []string{AliasString},
			Usage:   "string to hash",
		},
	},
}

func hashAction(ctx context.Context, cmd *cli.Command) error {
	var hash *string
	var err error
	alg := cmd.String(AliasAlgoritm)

	if file := cmd.String(FlagFile); len(file) > 0 {
		hash, err = HashFile(alg, file)
		if hash == nil {
			return cli.Exit("no hash", 1)
		}

		if err != nil {
			return cli.Exit(err, 1)
		}
		fmt.Printf("%s", *hash)
		return nil
	}

	if text := cmd.String(FlagString); len(text) > 0 {
		hash, err = HashString(alg, text)
		if hash == nil {
			return cli.Exit("no hash", 1)
		}
		if err != nil {
			return cli.Exit(err, 1)
		}
		fmt.Printf("%s", *hash)
		return nil
	}

	return cli.Exit(fmt.Sprintf("Required flag \"%s\" or \"%s\" not set", FlagFile, FlagString), 1)
}

func HashString(alg string, text string) (*string, error) {
	hash, err := algorithm(alg)
	if err != nil {
		return nil, err
	}
	data := strings.TrimSpace(text)
	if _, err := hash.Write([]byte(data)); err != nil {
		return nil, err
	}
	result := fmt.Sprintf("%x", hash.Sum(nil))
	return &result, nil
}

func HashFile(alg string, path string) (*string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash, err := algorithm(alg)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}
	result := fmt.Sprintf("%x", hash.Sum(nil))
	return &result, nil
}

func algorithm(alg string) (hash.Hash, error) {
	switch name := strings.ToLower(alg); name {
	case AlgSha256:
		return sha256.New(), nil
	case AlgSha1:
		return sha1.New(), nil
	case AlgMd5:
		return md5.New(), nil
	default:
		return nil, fmt.Errorf("Unknown algoritm %s", alg)
	}
}
