package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	flagName string
)

func getProcs(d string) []string {

	pid := strconv.Itoa(os.Getpid())
	ppid := strconv.Itoa(os.Getppid())

	files, err := ioutil.ReadDir(d)
	if err != nil {
		log.Fatal(err)
	}

	var v []string

	for _, f := range files {
		if f.IsDir() {
			if _, err := strconv.Atoi(f.Name()); err == nil {
				if f.Name() == ppid || f.Name() == pid {
					// Skipping own PID
				} else {
					v = append(v, f.Name())
				}
			}
		}
	}

	return v
}

func seachProc(n string) []byte {

	// fmt.Printf("Own PID %s %s\n", pid, ppid)
	d := "/proc"
	procs := getProcs(d)
	r, _ := regexp.Compile(n)
	for _, p := range procs {

		f := fmt.Sprintf("%s/%s/cmdline", d, p)
		o, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}

		if r.Match([]byte(o)) {
			fmt.Printf("%s\n", p)
		}
	}

	return nil
}

func init() {
	flag.StringVar(&flagName, "n", "", "Search for process by name")
}

func main() {
	flag.Parse()
	if flagName == "" {
		fmt.Println("n value is required.")
		os.Exit(1)
	}

	seachProc(flagName)
}
