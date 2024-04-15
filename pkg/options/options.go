package options

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/loft-sh/devpod/pkg/driver"
)

type Options struct {
	DiskSize string
	Token    string

	JobId     string
	Namespace string
	Region    string

	DriverOpts *driver.RunOptions
}

// Read ENV Vars for option overrides
func FromEnv() (*Options, error) {
	newopts, err := DefaultOptions()
	if err != nil {
		return nil, err
	}

	return newopts, nil
}

func DefaultOptions() (*Options, error) {
	var runOptions *driver.RunOptions
	runOptsEnv := os.Getenv("DEVCONTAINER_RUN_OPTIONS")
	if runOptsEnv != "" && runOptsEnv != "null" {
		runOptions = &driver.RunOptions{}
		err := json.Unmarshal([]byte(runOptsEnv), runOptions)
		if err != nil {
			return nil, fmt.Errorf("unmarshal run options: %w", err)
		}
	}

	return &Options{
		DiskSize:   "10G",
		Token:      "",
		Namespace:  "",
		Region:     "",
		JobId:      getEnv("DEVCONTAINER_ID", "devpod-nomad"), // set by devpod
		DriverOpts: runOptions,
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
