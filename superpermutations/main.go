package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/g-harel/superpermutations"
	"github.com/spf13/cobra"
)

var check bool
var length int
var out string
var print bool
var silent bool

func init() {
	command.PersistentFlags().BoolVarP(&check, "check", "c", false, "check correctness of result (big performance hit)")
	command.PersistentFlags().IntVarP(&length, "length", "l", 5, "set input string length (max 16)")
	command.PersistentFlags().StringVarP(&out, "out", "o", "", "write result to a file")
	command.PersistentFlags().BoolVarP(&print, "print", "p", false, "print the result (may be very large)")
	command.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "silence all output (except --print)")
}

func main() {
	command.Execute()
}

func log(p func(string, ...interface{}), f string, a ...interface{}) {
	if !silent {
		p(f, a...)
	}
}

var command = &cobra.Command{
	Use:     "superpermutations",
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		var t time.Time

		min := 0
		max := 13
		if length <= min {
			log(color.Red, "Error: length must be bigger than %d\n", min)
			os.Exit(1)
		} else if length > max {
			log(color.Red, "Error: lengths above %d are not supported (maximum slice size)\n", max)
			os.Exit(1)
		}

		value := "0123456789abc"[:length]

		log(color.White, "Computing for length %d ...", length)
		t = time.Now()

		sp := superpermutations.Find(value)

		if print {
			if silent {
				fmt.Print(sp)
			} else {
				log(color.Magenta, sp)
			}
		}

		log(color.Cyan, "Found, size: %d chars (%s)\n", len(sp), time.Since(t))

		if check {
			log(color.White, "Checking ...")
			t = time.Now()
			if superpermutations.Check(value, sp) {
				log(color.Cyan, "Check has passed! (%s)", time.Since(t))
			} else {
				log(color.Red, "Error: cannot not confirm result is a superpermutation")
				os.Exit(1)
			}
		}

		if out != "" {
			log(color.White, "Writing ...")
			err := ioutil.WriteFile(out, []byte(sp), 0644)
			if err != nil {
				log(color.Red, "Error: could not write to file \"%v\"\n", out)
				os.Exit(1)
			} else {
				log(color.Cyan, "Written successfully to \"%v\"", out)
			}
		}
	},
}
