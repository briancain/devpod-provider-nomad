package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/briancain/devpod-provider-nomad/pkg/nomad"
	"github.com/briancain/devpod-provider-nomad/pkg/options"
	"github.com/spf13/cobra"
)

// CommandCmd holds the cmd flags
type CommandCmd struct{}

// NewCommandCmd defines a command
func NewCommandCmd() *cobra.Command {
	cmd := &CommandCmd{}
	commandCmd := &cobra.Command{
		Use:   "command",
		Short: "Run a command on the Nomad instance",
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

func (cmd *CommandCmd) Run(
	ctx context.Context,
	options *options.Options,
) error {
	nomad, err := nomad.NewNomad(options)
	if err != nil {
		return err
	}

	code, err := nomad.CommandDevContainer(ctx,
		options.JobId,
		os.Getenv("DEVCONTAINER_USER"),
		os.Getenv("DEVCONTAINER_COMMAND"),
		os.Stdin,
		os.Stdout,
		os.Stderr,
	)
	if err != nil {
		return err
	}
	if code != 0 {
		return errors.New(fmt.Sprintf("command failed with exit code %d", code))
	}

	return nil
}
