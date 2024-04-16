package nomad

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/briancain/devpod-provider-nomad/pkg/options"
	"github.com/hashicorp/nomad/api"
	"github.com/loft-sh/devpod/pkg/client"
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
) (client.Status, *api.Job, error) {
	job, _, err := n.client.Jobs().Info(jobID, nil)
	if err != nil {
		return client.StatusNotFound, job, err
	}

	status := *job.Status
	switch status {
	case "pending":
		return client.StatusBusy, job, nil
	case "running":
		return client.StatusRunning, job, nil
	case "complete":
		return client.StatusStopped, job, nil
	case "dead":
		return client.StatusStopped, job, nil
	case "":
		return client.StatusNotFound, job, nil
	default:
		return client.StatusNotFound, job, nil
	}

	return client.StatusNotFound, job, nil
}

// Run a command on the instance
func (n *Nomad) CommandDevContainer(
	ctx context.Context,
	jobID string,
	user string,
	command string,
	stdin io.Reader,
	stdout io.Writer,
	stderr io.Writer,
) (int, error) {
	ctx, cancelFn := context.WithCancel(ctx)
	defer cancelFn()

	// Check if the job is running
	status, _, err := n.Status(ctx, jobID)
	if err != nil {
		return -1, err
	}
	if status != client.StatusRunning {
		return -1, errors.New("job is not running")
	}

	// Get our allocation ID to exec into
	allocs, _, err := n.client.Jobs().Allocations(jobID, false, nil)
	if err != nil {
		return -1, err
	}
	if len(allocs) == 0 {
		return -1, fmt.Errorf("job %q has no allocations found", jobID)
	}
	// Check for running allocations
	var allocID string
	for _, alloc := range allocs {
		if alloc.ClientStatus == "running" {
			// Pick the first one
			allocID = alloc.ID
			break
		}
	}
	if allocID == "" {
		return -1, fmt.Errorf("job %q has no running allocations found", jobID)
	}

	alloc, _, err := n.client.Allocations().Info(allocID, nil)
	if err != nil {
		return -1, err
	}
	// TODO: make this an options
	task := "devpod"
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		for range signalCh {
			cancelFn()
		}
	}()

	sizeCh := make(chan api.TerminalSize, 1)

	return n.client.Allocations().Exec(ctx, alloc, task, true, []string{command},
		stdin, stdout, stderr, sizeCh, nil)
}
