package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func scan(dir string) []string {
	f, err := os.Open(dir)
	if err != nil {
		return []string{}
	}
	defer f.Close()

	fs, err := f.Readdir(-1)
	if err != nil {
		panic(err)
	}

	ans := []string{}
	for _, v := range fs {
		name := dir + "/" + v.Name()
		if strings.Contains(name, "/vendor/") {
			continue
		}
		if v.IsDir() {
			ans = append(ans, scan(name)...)
		} else if strings.HasSuffix(name, ".go") {
			ans = append(ans, name)
		}
	}
	return ans
}

func format(f string, wg *sync.WaitGroup) {
	defer wg.Done()
	cmd := exec.Command("go", "fmt", f)
	if err := cmd.Run(); err != nil {
		fmt.Println("format " + f + " failed")
		return
	}
	fmt.Println("format " + f + " success")
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	files := scan(pwd)

	wg := &sync.WaitGroup{}
	for _, f := range files {
		wg.Add(1)
		go format(f, wg)
	}
	wg.Wait()
	fmt.Println("finished")
}
