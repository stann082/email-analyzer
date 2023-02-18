package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
)

func main() {
	emailPath := path.Join(os.Getenv("USERPROFILE"), "Documents\\email")
	files, err := os.ReadDir(emailPath)
	if err != nil {
		log.Fatal(err)
	}

	addresses := createAddressMap(files)
	keys := getSortedKeys(addresses)
	for _, k := range keys {
		fmt.Println(k, addresses[k])
	}
}

func createAddressMap(files []fs.DirEntry) map[string]int {
	var addresses = make(map[string]int)
	var rgx = regexp.MustCompile(`\((.*?)\)`)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		rs := rgx.FindStringSubmatch(file.Name())
		var length = len(rs)
		if length == 0 {
			continue
		}

		var address string
		if length > 1 {
			address = rs[1]
		} else if length == 1 {
			address = rs[0]
		}

		addresses[address] = addresses[address] + 1
	}
	return addresses
}

func getSortedKeys(addresses map[string]int) []string {
	keys := make([]string, 0, len(addresses))
	for k := range addresses {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return addresses[keys[i]] < addresses[keys[j]]
	})

	return keys
}
