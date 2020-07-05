package main

import (
	"log"

	"github.com/spf13/cobra"
)

const (
	timesFlag = "times"
)

type database struct {
}

func main() {
	var rootCmd = &cobra.Command{Use: "app"}

	newSubcommand := func(collaborators ...interface{}) *cobra.Command {
		cmd := &cobra.Command{
			Use: "subcommand",
			RunE: func(cmd *cobra.Command, args []string) error {
				timesOption, err := cmd.Flags().GetInt(timesFlag)
				if err != nil {
					panic("flag " + timesFlag + " must exist")
				}
				log.Printf("got times: %d\n", timesOption)
				return nil
			},
		}

		var times int
		cmd.Flags().IntVarP(&times, timesFlag, "t", 1, "times to do a thing")
		return cmd
	}

	rootCmd.AddCommand(newSubcommand(&database{}))

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
