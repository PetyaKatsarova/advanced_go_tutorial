package main

import (
	"crypto/md5"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	// for each regular file, start a gorouting to sum the file and sends
	c 		:= make(chan result)
	errc 	:= make(chan error, 1)

	go func() {
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil { return err  }
            if !info.Mode().IsRegular() { return nil }
            wg.Add(1)
			go func() {
				data, err := os.ReadFile(path)
				select {
				case c <- result{path, md5.Sum(data), err}:
				case <-done:
				}
				wg.Done()
			}()
			//abort the walk if done is closed
			select {
			case <- done:
				return errors.New("walk canceled")
			default: return nil
			}
		})

		go func(){
			wg.Wait()
			close(c)
		}()
		errc <- err
	}()
	return c, errc
}

func MD5ALL2(root string)(map[string][md5.Size]byte, error) {
	// MD5All closes the done channel when it returns; it may do so before
    // receiving all the values from c and errc.
	done := make(chan struct{})
	defer close(done)
	 c, errc := sumFiles(done, root)

	 m := make(map[string][md5.Size]byte)
	 for val := range c {
		if val.err != nil { return nil, val.err}
		m[val.path] = val.sum
	 }
	 if err := <-errc; err != nil { return nil, err }
	 return m, nil
}