package testmod

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	files, err := ioutil.ReadDir(dir)
	if err != nil{
		return err
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}