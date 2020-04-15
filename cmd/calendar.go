package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tempor1s/msconsole/modules"
)

func init() {
	rootCmd.AddCommand(calendarCommand)
}

var calendarCommand = &cobra.Command{
	Use:   "calendar",
	Short: "keep track of calendar events, open links, and more!",
	Long:  "This command is used for opening links from your calendar and keeping track of events.",
	Run: func(cmd *cobra.Command, args []string) {
		modules.CalendarModule(cmd)
	},
}
