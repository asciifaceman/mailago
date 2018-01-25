package cmd

import (
  "github.com/spf13/cobra"
)

const (
  defaultPort = 3031
)

var (
  port int
)

// RunCmd command assembly
var RunCmd = &cobra.Command{
  Use:   "run",
  Short: "Run MailaGo",
  Run:   nil,
}

func init() {
  RootCmd.AddCommand(RunCmd)
  RunCmd.Flags().IntVarP(&port, "port", "p", defaultPort, "Start MailaGo server. Default port is 3031")
}
