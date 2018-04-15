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
var runes string
var silent bool

func init() {
	command.PersistentFlags().BoolVarP(&check, "check", "c", false, "check correctness of result (big performance hit)")
	command.PersistentFlags().IntVarP(&length, "length", "l", 5, "set input string length (max 13)")
	command.PersistentFlags().StringVarP(&out, "out", "o", "", "write result to a file")
	command.PersistentFlags().BoolVarP(&print, "print", "p", false, "print the result (may be very large)")
	command.PersistentFlags().StringVarP(&runes, "runes", "r", "", "custom list of chars (looped if < length)")
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

func fatal(f string, a ...interface{}) {
	if !silent {
		color.New(color.FgRed).Fprintf(os.Stderr, "Error: "+f, a...)
	}
	os.Exit(1)
}

var command = &cobra.Command{
	Use:     "superpermutations",
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		var t time.Time

		min := 0
		max := 13
		if length <= min {
			fatal("length must be bigger than %d\n", min)
		} else if length > max {
			fatal("lengths above %d are not supported (maximum slice size)\n", max)
		}

		if runes != "" {
			for len(runes) < length {
				runes += runes
			}
		} else {
			runes = "0123456789abc"
		}
		runes = runes[:length]

		log(color.White, "Computing for length %d ...", length)
		t = time.Now()

		sp := superpermutations.Find(runes)

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
			if superpermutations.Check(runes, sp) {
				log(color.Cyan, "Check has passed! (%s)", time.Since(t))
			} else {
				fatal("cannot not confirm result is a superpermutation")
			}
		}

		if out != "" {
			log(color.White, "Writing ...")
			err := ioutil.WriteFile(out, []byte(sp), 0644)
			if err != nil {
				fatal("could not write to file \"%v\"\n", out)
			} else {
				log(color.Cyan, "Written successfully to \"%v\"", out)
			}
		}
	},
}
