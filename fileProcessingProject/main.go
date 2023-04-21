package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type result struct {
	path string
	hash string
}

func main() {
	files, _ := ioutil.ReadDir("./files")
	filesCount := len(files)

	workers := runtime.GOMAXPROCS(0) * 2
	pathChannel := make(chan string)
	resultChannel := make(chan result, filesCount)
	doneChannel := make(chan struct{})
	doneChannel2 := make(chan struct{})
	resultMap := make(map[string][]string)

	for i := 0; i < workers; i++ {
		go processor(pathChannel, resultChannel, doneChannel)
	}

	go collector(resultChannel, doneChannel2, resultMap)

	filepath.Walk("./files/", func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			pathChannel <- path
		}
		return nil
	})
	close(pathChannel)

	for i := 0; i < workers; i++ {
		<-doneChannel
	}
	close(doneChannel)
	close(resultChannel)

	<-doneChannel2
	close(doneChannel2)

	for x, y := range resultMap {
		fmt.Println(x)
		fmt.Println(len(y))
	}

}

func processor(pathChan <-chan string, resultChan chan<- result, done chan<- struct{}) {
	for path := range pathChan {
		hash := md5.New()
		file, err := os.Open(path)
		if err != nil {
			log.Fatal("error reading file")
			return
		}
		defer file.Close()

		if _, err := io.Copy(hash, file); err != nil {
			log.Fatal("error copying file", err)
			return
		}
		resultChan <- result{path, fmt.Sprintf("%x", hash.Sum(nil))}
	}
	done <- struct{}{}
}

func collector(resultChan <-chan result, done chan<- struct{}, resultmap map[string][]string) {
	for res := range resultChan {
		resultmap[res.hash] = append(resultmap[res.hash], res.path)
	}
	done <- struct{}{}
}
