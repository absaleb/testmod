package testmod

import (
	"flag"
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

func getNumberOfFiles(dir string) int {
	count := 0

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			count = count + 1
		}
		return nil
	})

	if err != nil {
		return 0
	}

	return count
}

func ListDirectory(dir string, outputRootDir string, maxGoroutines int) error {
	maxNumGoroutines := flag.Int("maxNumGoroutines", maxGoroutines, "max number of goroutines")
	numOfJobs := flag.Int("numOfJobs", getNumberOfFiles(dir)*maxGoroutines, "number of jobs")

	flag.Parse()

	goroutines := make(chan struct{}, *maxNumGoroutines)
	for i := 0; i < *maxNumGoroutines; i++ {
		goroutines <- struct{}{}
	}

	done := make(chan bool)
	waitAll := make(chan bool)

	go func() {
		for i := 0; i < *numOfJobs; i++ {
			<-done
			goroutines <- struct{}{}
		}
		waitAll <- true
	}()

	var files map[string]string
	files = make(map[string]string)

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			var outputDir string
			dt, err := GetExifDate(path)
			if err != nil {
				fmt.Println(path)
				outputDir = filepath.Join(outputRootDir, "19700101")
			} else {
				outputDir = filepath.Join(outputRootDir, fmt.Sprintf("%d%02d%02d", dt.Year(), dt.Month(), dt.Day()))
			}

			if _, err = os.Stat(outputDir); os.IsNotExist(err) {
				err = os.MkdirAll(outputDir, os.ModePerm)
				if err != nil {
					fmt.Printf("###err os.Open : %s\n", err)
					return err
				}
			}
			outputPath := filepath.Join(outputDir, info.Name())

			files[path] = outputPath

			return nil
		})

	for path, outputPath := range files {
		<-goroutines
		go func() {
			r, err := os.Open(path)
			if err != nil {
				fmt.Printf("###err os.Open : %s\n", err)
				return
			}
			defer r.Close()

			f, err := os.Create(outputPath)
			if err != nil {
				fmt.Printf("###err os.Create : %s\n", err)
				return
			}
			defer f.Close()

			_, err = io.Copy(f, r)
			if err != nil {
				fmt.Printf("###err os.Copy : %s\n", err)
				return
			}

			fmt.Println(outputPath)

			done <- true
		}()
	}
	<-waitAll
	return err
}
