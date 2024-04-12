package cmd

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "devpod-provider-nomad",
		Short: "DevPod provider for Nomad",
	}

	return cmd
}

func Execute() {
	cmd := NewRootCmd()
	// Add commands

	if err := cmd.Execute(); err != nil {
		// TODO: handle this more gracefully
		panic(err)
	}
}
