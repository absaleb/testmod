package testmod

import (
	"fmt"
	"io"
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

func ListDirectory(dir string, outputRootDir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			dt, err := GetExifDate(path)
			if err != nil {
				fmt.Println(path)
			} else {
				outputDir := filepath.Join(outputRootDir, fmt.Sprintf("%d%02d%02d", dt.Year(), dt.Month(), dt.Day()))
				err = os.MkdirAll(outputDir, os.ModePerm)
				if err != nil {
					fmt.Printf("###err os.Open : %s\n", err)
				}else {
					outputPath := filepath.Join(outputDir, info.Name())
					r, err := os.Open(path)
					if err != nil {
						fmt.Printf("###err os.Open : %s\n", err)
					} else {
						f, err := os.Create(outputPath)
						if err != nil {
							fmt.Printf("###err os.Create : %s\n", err)
						} else {
							_, err = io.Copy(f, r)
							if err != nil {
								fmt.Printf("###err os.Copy : %s\n", err)
							}
						}
					}
					fmt.Println(outputPath)
				}

				fmt.Println(path, dt)
			}

			return nil
		})

	return err
}
