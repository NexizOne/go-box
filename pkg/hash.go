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
	hashFlagAlgoritm  = "algorithm"
	hashFlagFile      = "file"
	hashFlagString    = "string"
	hashAliasAlgoritm = "a"
	hashAliasFile     = "f"
	hashAliasString   = "s"

	// algorithms
	hashAlgSha256 = "sha256"
	hashAlgSha1   = "sha1"
	hashAlgMd5    = "md5"
)

// all algorithms
var AlgAll = []string{hashAlgSha256, hashAlgSha1, hashAlgMd5}

var CommandHash *cli.Command = &cli.Command{
	Name:                  internal.CmdHash,
	Version:               internal.Version,
	Usage:                 "Hash string or file",
	EnableShellCompletion: true,
	ArgsUsage:             "[options]",
	Action:                hashAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    hashFlagAlgoritm,
			Aliases: []string{hashAliasAlgoritm},
			Value:   hashAlgSha256,
			Usage:   fmt.Sprintf("hashing algorithm (%s)", strings.Join(AlgAll, ", ")),
		},
		&cli.StringFlag{
			Name:    hashFlagFile,
			Aliases: []string{hashAliasFile},
			Usage:   "file to hash",
		},
		&cli.StringFlag{
			Name:    hashFlagString,
			Aliases: []string{hashAliasString},
			Usage:   "string to hash",
		},
	},
}

func hashAction(ctx context.Context, cmd *cli.Command) error {
	var hash *string
	var err error
	alg := cmd.String(hashAliasAlgoritm)

	if file := cmd.String(hashFlagFile); len(file) > 0 {
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

	if text := cmd.String(hashFlagString); len(text) > 0 {
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

	return cli.Exit(fmt.Sprintf("Required flag \"%s\" or \"%s\" not set", hashFlagFile, hashFlagString), 1)
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
	case hashAlgSha256:
		return sha256.New(), nil
	case hashAlgSha1:
		return sha1.New(), nil
	case hashAlgMd5:
		return md5.New(), nil
	default:
		return nil, fmt.Errorf("Unknown algoritm %s", alg)
	}
}
