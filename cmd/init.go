package cmd

import (
	"context"
	"fmt"

	"github.com/briancain/devpod-provider-nomad/pkg/nomad"
	"github.com/briancain/devpod-provider-nomad/pkg/options"
	"github.com/spf13/cobra"
)

// InitCmd holds the cmd flags
type InitCmd struct{}

// NewCommandCmd defines a command
func NewInitCmd() *cobra.Command {
	cmd := &InitCmd{}
	commandCmd := &cobra.Command{
		Use:   "init",
		Short: "Check that we can connect to Nomad",
		RunE: func(_ *cobra.Command, args []string) error {
			options, err := options.FromEnv()
			if err != nil {
				return err
			}

			return cmd.Run(context.Background(), options)
		},
	}

	return commandCmd
}

func (cmd *InitCmd) Run(
	ctx context.Context,
	options *options.Options,
) error {
	nomad, err := nomad.NewNomad(options)
	if err != nil {
		return err
	}

	if err := nomad.Init(ctx); err != nil {
		return err
	}

	fmt.Println("Nomad is ready")
	return nil
}
