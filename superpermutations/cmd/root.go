package cmd

import (
	"io/ioutil"

	"github.com/fatih/color"
	"github.com/g-harel/superpermutations"
	"github.com/spf13/cobra"
)

var check bool
var length int
var print bool
var write string

var rootCmd = &cobra.Command{
	Use: "superpermutations",
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&check, "check", "c", false, "check correctness of result")
	rootCmd.PersistentFlags().BoolVarP(&print, "print", "p", false, "print the result (may be very large)")
	rootCmd.PersistentFlags().IntVarP(&length, "length", "l", 5, "set input string length (max 16)")
	rootCmd.PersistentFlags().StringVarP(&write, "write", "w", "", "write result to a file")
}

// Execute reads arguments and runs the desired action.
func Execute() {
	rootCmd.Execute()

	min := 0
	max := 13
	if length <= min {
		color.Red("Error: length must be bigger than %d\n", min)
		return
	} else if length > max {
		color.Red("Error: lengths above %d are not supported (maximum slice size)\n", max)
		return
	}

	chars := "0123456789abcdef"
	value := ""
	for i := 0; i < length; i++ {
		value += string(chars[i])
	}

	color.White("Computing for length %d ...", length)

	sp := superpermutations.Find(value)

	if print {
		color.Magenta(sp)
	}

	color.Cyan("Found, size: %d chars\n", len(sp))

	if check {
		color.White("Checking ...")
		if superpermutations.Check(value, sp) {
			color.Cyan("Check has passed!")
		} else {
			color.Red("Error: cannot not confirm result is a superpermutation")
		}
	}

	if write != "" {
		color.White("Writing ...")
		err := ioutil.WriteFile(write, []byte(sp), 0644)
		if err != nil {
			color.Red("Error: could not write to file \"%v\"\n", write)
		} else {
			color.Cyan("Written successfully to \"%v\"", write)
		}
	}
}
