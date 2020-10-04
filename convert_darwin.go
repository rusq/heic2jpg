package heic2jpg

import (
	"log"
	"os"
	"os/exec"
	"strconv"
)

func (c *Converter) convert(src, dst string) error {
	log.Printf("%s -> %s", src, dst)

	cmd := exec.Command("sips", "-s", "format", "jpeg", "-s", "formatOptions", strconv.Itoa(c.quality), src, "-o", dst)
	if out, err := cmd.CombinedOutput(); err != nil {
		return err
	} else {
		log.Println(string(out))
	}

	if c.rm {
		return os.Remove(src)
	}
	return nil
}
