package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	root := "/home/alaa"
	fileSizes := make(chan int64)
	var wg sync.WaitGroup

	wg.Add(1)
	go walk(root, fileSizes, &wg)

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	tick = time.Tick(100 * time.Millisecond)
	var nbfiles, nbytes int64

loop:
	for {
		select {
		case <-tick:
			stats(nbfiles, nbytes)
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nbfiles++
			nbytes += size
		}
	}
	stats(nbfiles, nbytes)
}

func stats(nbfiles, nbytes int64) {
	fmt.Println("Number of files found: ", nbfiles)
	fmt.Println("Total Directory Size: ", nbytes)
	fmt.Println("\n")
}

func walk(root string, sizesCh chan int64, wg *sync.WaitGroup) {
	defer wg.Done()

	dirs := dirEntries(root)
	for _, f := range dirs {
		if f.IsDir() {
			wg.Add(1)
			go walk(filepath.Join(root, f.Name()), sizesCh, wg)
		}
		sizesCh <- f.Size()
	}
}

var sema = make(chan struct{}, 10)

func dirEntries(path string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return entries
}
