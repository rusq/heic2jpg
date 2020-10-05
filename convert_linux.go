package heic2jpg

import (
	"image/jpeg"
	"os"

	"github.com/jdeng/goheif"
)

func (c *Converter) convert(src, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	img, err := goheif.Decode(f)
	if err != nil {
		return err
	}
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if err := jpeg.Encode(out, img, &jpeg.Options{Quality: c.quality}); err != nil {
		return err
	}
	return nil
}
