package nomad

import (
	"context"

	"github.com/hashicorp/nomad/api"
)

type Nomad struct {
	// Nomad client
	client *api.Client
}

func NewNomad() (*Nomad, error) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}

	return &Nomad{
		client: client,
	}, nil
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
