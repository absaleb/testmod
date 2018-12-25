package testmod

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

func Sum(a int, b int, c int, name string) string {
	return fmt.Sprintf("for %s sum of %d and %d and %d is %d !!!", name, a, b, c, a+b+c)

}

func GetExifDate(fname string) (*time.Time, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}

	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		return nil, err
	}
	tm, err := x.DateTime()

	return &tm, err
}

func ListDirectory(dir string, outputDir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			dt, err := GetExifDate(path)
			if err != nil {
				fmt.Println(path)
			} else {
				fmt.Println(fmt.Sprintf("%4d%2s%2d", dt.Year(), dt.Month().String(), dt.Day()))
				fmt.Println(path, dt)
			}

			return nil
		})

	return err
}
