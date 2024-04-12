package nomad

import "github.com/hashicorp/nomad/api"

type Nomad struct {
	// Nomad client
	client *api.Client
}

func NewNomad() (*Nomad, error) {
	return &Nomad{}, nil
}
