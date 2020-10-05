package heic2jpg

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const heicExt = ".heic"

// Converter contains the runtime settings for the convertion process.
type Converter struct {
	recursive bool
	rm        bool
	quality   int
}

// New creates the new converter.
func New(recursive bool, rm bool, quality int) (*Converter, error) {
	if quality < 0 || 100 < quality {
		return nil, errors.New("invalid quality value")
	}
	return &Converter{recursive: recursive, rm: rm, quality: quality}, nil
}

// WalkFn is the function that should be passed to os.Walk to do the conversion
// of files in some directory.
func (c *Converter) WalkFn(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !c.recursive && info.IsDir() && path != "." {
		return filepath.SkipDir
	}
	if strings.ToLower(filepath.Ext(path)) != heicExt {
		return nil
	}

	if err := c.convert(path, ReplaceExt(path, ".jpg")); err != nil {
		return err
	}
	return nil
}

func ReplaceExt(path string, newExt string) string {
	return strings.Replace(path, filepath.Ext(path), newExt, 1)
}

func (c *Converter) Convert(src, dst string) error {
	if !strings.EqualFold(strings.ToLower(filepath.Ext(src)), heicExt) {
		return fmt.Errorf("only %s files are supported", strings.ToUpper(heicExt))
	}
	return c.convert(src, dst)
}
