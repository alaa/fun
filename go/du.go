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

	var n sync.WaitGroup

	n.Add(1)
	go walk(root, fileSizes, &n)

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	tick = time.Tick(50 * time.Millisecond)
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

func walk(root string, sizesCh chan int64, n *sync.WaitGroup) {
	defer n.Done()

	dirs := dirEntries(root)
	for _, f := range dirs {
		if f.IsDir() {
			n.Add(1)
			go walk(filepath.Join(root, f.Name()), sizesCh, n)
		}
		sizesCh <- f.Size()
	}
}

func dirEntries(path string) []os.FileInfo {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return nil
	}
	return entries
}
