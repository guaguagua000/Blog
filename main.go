package main

import (
	"Blog/cmd/job"
	"Blog/cmd/server"
	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{Use: "sniper"}

	root.AddCommand(
		server.Cmd,
		job.Cmd,
	)

	root.Execute()
}
