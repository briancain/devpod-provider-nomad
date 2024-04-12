package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/briancain/devpod-provider-nomad/pkg/nomad"
	"github.com/briancain/devpod-provider-nomad/pkg/options"
	"github.com/spf13/cobra"
)

// StatusCmd holds the cmd flags
type StatusCmd struct{}

// NewCommandCmd defines a command
func NewStatusCmd() *cobra.Command {
	cmd := &StatusCmd{}
	commandCmd := &cobra.Command{
		Use:   "status",
		Short: "Get the status of a devpod instance on Nomad",
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

func (cmd *StatusCmd) Run(
	ctx context.Context,
	options *options.Options,
) error {
	nomad, err := nomad.NewNomad(options)
	if err != nil {
		return err
	}

	respJob, err := nomad.Status(ctx, options.JobId)
	if err != nil {
		return err
	}

	status := respJob.Status
	_, err = fmt.Fprint(os.Stdout, status)

	return nil
}
