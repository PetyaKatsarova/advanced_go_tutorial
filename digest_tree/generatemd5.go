package main

import (
	"crypto/md5"
	"fmt"
)
// command: go env
// go build
// .\digest_tree.exe
/*
Package md5 implements the MD5 hash algorithm as defined in RFC 1321.
MD5 is cryptographically broken and should not be used for secure applications.
*/

func GenerateMD5() {
	exampleMap 	:= make(map[string][md5.Size]byte) //size of an MD5 checksum in bytes.
	keys 		:= []string{"example1", "example2", "example3"}

	for _, key := range keys {
		hash := md5.Sum([]byte(key)) // returns the MD5 checksum of the data.
		exampleMap[key] = hash
	}
	for key, val := range exampleMap {
		fmt.Printf("key: %s, MD5 hash: %x\n", key, val)
	}
}