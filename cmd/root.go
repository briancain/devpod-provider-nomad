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
	rootCmd := NewRootCmd()
	// Add commands
	rootCmd.AddCommand(NewCommandCmd())
	rootCmd.AddCommand(NewInitCmd())

	if err := rootCmd.Execute(); err != nil {
		// TODO: handle this more gracefully
		panic(err)
	}
}
