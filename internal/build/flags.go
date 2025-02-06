package build

import (
	"flag"
	"fmt"
	"os"
)

var verbose = flag.Bool("verbose", false, "Runs the program in verbose mode (all logs)")
var parsed = false

func init() {
	if parsed {
		return
	}

	flag.BoolFunc("version", "Prints build version", func(s string) error {
		fmt.Println(Version())

		os.Exit(0)
		return nil
	})

	parsed = true
	flag.Parse()
}
