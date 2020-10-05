package main

import (
	"errors"
	"flag"
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/rusq/heic2jpg"
)

var (
	root      = flag.String("d", "", "specify root `directory` to start processing from")
	recursive = flag.Bool("r", false, "process directories recursively")
	rm        = flag.Bool("rm", false, "remove source file")
	quality   = flag.Int("q", jpeg.DefaultQuality, "jpeg `quality`")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags] < <-d dir> | <filename.heic> >\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		flag.Usage()
		log.Fatal(err)
	}

}

func run() error {
	c, err := heic2jpg.New(*recursive, *rm, *quality)
	if err != nil {
		return err
	}

	if flag.Arg(0) != "" {
		// process single file
		if err := c.Convert(flag.Arg(0), heic2jpg.ReplaceExt(flag.Arg(0), ".jpg")); err != nil {
			return err
		}
	} else {
		// walk the directory
		if *root == "" {
			return errors.New("no dir specified")
		}
		if err := filepath.Walk(*root, c.WalkFn); err != nil {
			return err
		}
	}
	return nil
}
