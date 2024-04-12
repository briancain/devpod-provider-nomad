package cmd

import (
	"context"

	"github.com/briancain/devpod-provider-nomad/pkg/nomad"
	"github.com/briancain/devpod-provider-nomad/pkg/options"
	"github.com/hashicorp/nomad/api"
	"github.com/spf13/cobra"
)

const (
	defaultImage = "alpine"
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

	image := defaultImage
	if options.DriverOpts != nil && options.DriverOpts.Image != "" {
		image = options.DriverOpts.Image
	}

	// TODO(briancain): add more fields
	jobName := "devpod"
	job := &api.Job{
		ID:        &options.JobId,
		Name:      &jobName,
		Namespace: &options.Namespace,
		TaskGroups: []*api.TaskGroup{
			{
				Name: &jobName,
				Tasks: []*api.Task{
					{
						Name: "devpod",
						Config: map[string]interface{}{
							"image": image,
						},
						Driver: "docker",
					},
				},
			},
		},
	}

	_, err = nomad.Create(ctx, job)
	if err != nil {
		return err
	}

	return nil
}
