package cmd

import (
	"context"

	"github.com/briancain/devpod-provider-nomad/pkg/nomad"
	"github.com/briancain/devpod-provider-nomad/pkg/options"
	"github.com/hashicorp/nomad/api"
	"github.com/spf13/cobra"
)

// CreateCmd holds the cmd flags
type CreateCmd struct{}

// NewCommandCmd defines a command
func NewCreateCmd() *cobra.Command {
	cmd := &CreateCmd{}
	commandCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new devpod instance on Nomad",
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

func (cmd *CreateCmd) Run(
	ctx context.Context,
	options *options.Options,
) error {
	nomad, err := nomad.NewNomad(options)
	if err != nil {
		return err
	}

	// TODO(briancain): add more fields
	jobName := "devpod"
	job := &api.Job{
		ID:        &options.JobId,
		Name:      &jobName,
		Namespace: &options.Namespace,
	}

	// Do something with the evalID?
	_, err = nomad.Create(ctx, job)
	if err != nil {
		return err
	}
	return nil
}
