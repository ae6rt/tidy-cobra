package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

type database struct {
}

func main() {
	var rootCmd = &cobra.Command{Use: "app"}

	newSubcommand := func(collaborators ...interface{}) *cobra.Command {
		cmd := &cobra.Command{
			Use: "subcommand",
			RunE: func(cmd *cobra.Command, args []string) error {
				timesOption, err := cmd.Flags().GetInt("times")
				if err != nil {
					panic("flag times must exist")
				}
				fmt.Printf("got times: %d\n", timesOption)
				return nil
			},
		}

		var times int
		cmd.Flags().IntVarP(&times, "times", "t", 1, "times to do a thing")
		return cmd
	}

	rootCmd.AddCommand(newSubcommand(&database{}))

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
