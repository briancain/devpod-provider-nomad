package nomad

import (
	"context"
	"errors"
	"io"

	"github.com/briancain/devpod-provider-nomad/pkg/options"
	"github.com/hashicorp/nomad/api"
)

type Nomad struct {
	// Nomad client
	client *api.Client
}

func NewNomad(opts *options.Options) (*Nomad, error) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}

	return &Nomad{
		client: client,
	}, nil
}

func (n *Nomad) Init(
	ctx context.Context,
) error {
	// List nomad jobs to confirm we can connect
	_, _, err := n.client.Jobs().List(nil)
	if err != nil {
		return err
	}

	return nil
}

func (n *Nomad) Create(
	ctx context.Context,
	job *api.Job,
) (*api.JobRegisterResponse, error) {
	resp, _, err := n.client.Jobs().Register(job, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *Nomad) Delete(
	ctx context.Context,
	jobID string,
) error {
	_, _, err := n.client.Jobs().Deregister(jobID, true, nil)
	if err != nil {
		return err
	}

	return nil
}

func (n *Nomad) Status(
	ctx context.Context,
	jobID string,
) (*api.Job, error) {
	job, _, err := n.client.Jobs().Info(jobID, nil)
	if err != nil {
		return nil, err
	}

	return job, nil
}

// Untested
// Run a command on the instance
func (n *Nomad) CommandDevContainer(
	ctx context.Context,
	jobID string,
	user string,
	command string,
	stdin io.Reader,
	stdout io.Writer,
	stderr io.Writer,
) error {
	return errors.New("not implemented")
}
