package internal

import (
	"path/filepath"
	"strings"
)

func WithoutExtension(basename string) string {
	return strings.TrimSuffix(basename, filepath.Ext(basename))
}

func Basename(name string) string {
	return WithoutExtension(filepath.Base(name))
}
