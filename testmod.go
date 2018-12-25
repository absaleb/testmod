package testmod

import (
	"fmt"
	"log"
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
		log.Fatal(err)
		return nil, err
	}

	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	tm, err := x.DateTime()

	return &tm, err
}

func ListDirectory(dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error{
			if err != nil{
				return err
			}

			dt, err := GetExifDate(path)
			fmt.Println(path, dt)
			return nil
		})

	return err
}