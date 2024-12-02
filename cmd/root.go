package cmd

import "github.com/spf13/cobra"

var RootCommand = cobra.Command{
	Use: "semis",
}

func init() {
	RootCommand.AddCommand(newCommand(equal), newCommand(notequal), newCommand(greaterthan), newCommand(lessthan), newCommand(constraint))
}
