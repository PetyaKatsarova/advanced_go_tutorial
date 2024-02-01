/*
Get-ChildItem -Filter *.go | Get-FileHash -Algorithm MD5for linux:
md5sum *.go
*/

package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	fmt.Println("func GenerateMD5: ")
	GenerateMD5()
	fmt.Println("hello world")
	
	m, err := MD5All(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	fmt.Println("sorted MD5ALL: ")
	for _, path := range paths {
		fmt.Printf("%x %s\n", m[path], path) // %x is hexadecimal string
	}
}

func MD5All(root string) (map[string][md5.Size]byte, error) {
	m := make(map[string][md5.Size]byte)
	//filepath.Walk is a function in Go's path/filepath package that is used to traverse a directory tree.
	//It visits all the files and directories in the tree, including the root directory. This function is
	//especially useful when you need to process files in a directory and all its subdirectories.
	//func Walk(root string, walkFn WalkFunc) error
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {	return err }
			if !info.Mode().IsRegular() { return nil }
			data, err := os.ReadFile(path)
			if err != nil {	return err }
			m[path] = md5.Sum(data) // checksum of the data in the file?
			return nil
		})
		if err != nil { return nil, err }
		return m, nil
}

/*
In Go, info.Mode().IsRegular() is a method call used to check whether a file system entry (like a file or directory)
 is a regular file. This is typically used in conjunction with file traversal functions like filepath.Walk.
  Let's break down what this method call means:
  info: This is an os.FileInfo object. os.FileInfo is an interface in the Go standard library that
  provides basic information about a file system entry (which could be a file, directory, symbolic link,
	 etc.). This interface is returned by several functions in the os and path/filepath packages when they need to provide details about file system entries.
Mode(): The Mode method is a function of the os.FileInfo interface. It returns a os.FileMode value,
which is a bitmask representing the file mode and permission bits.
IsRegular(): IsRegular is a method on os.FileMode. It returns a boolean value indicating whether
the file mode describes a regular file. A regular file in this context means it's not a directory,
 symbolic link, device file, etc. It's a standard file that can hold content like text, binary data, etc.
If IsRegular() returns true, it means the file system entry is a regular file.
If it returns false, the file system entry might be a directory, symbolic link, or some other
non-regular file type.
*/
