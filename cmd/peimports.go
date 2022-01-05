package main

import (
	"flag"
	"fmt"
	"github.com/mikewiacek/peimports"
	"os"
)

func main() {
	flag.Parse()

	for _, fname := range flag.Args() {
		func() {
			f, err := os.Open(fname)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			peObj, err := peimports.NewFile(f)
			if err != nil {
				panic(err)
			}

			s, err := peObj.ImportedSymbols()
			if err != nil {
				panic(err)
			}

			fmt.Println(fname, " = ", s)
		}()
	}
}
