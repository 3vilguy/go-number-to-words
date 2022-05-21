package cmd

import (
	"errors"
	"fmt"
	"numbers-to-words/convert"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:                   "numbers-to-words [number]",
	Short:                 "Convert number to English words",
	Long:                  "Takes a number, from 0 to 100000, and converts the number into grammatically correct English words.",
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires exactly one argument\n")
		}

		val, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New(fmt.Sprintf("%s is not a number\n", args[0]))
		}
		if val < 0 || val > 100000 {
			return errors.New(fmt.Sprintf("%s has to be a number from 0 to 100000\n", args[0]))
		}

		return nil
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		val, _ := strconv.Atoi(args[0])
		fmt.Println(convert.Number2Words(val))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.numbers-to-words.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
