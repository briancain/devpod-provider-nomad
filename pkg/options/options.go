package options

import "os"

type Options struct {
	MachineID     string
	MachineFolder string

	DiskImage   string
	DiskSize    string
	MachineType string
	Token       string

	JobId       string
	NomadBinary string
}

// Read ENV Vars for option overrides
func FromEnv() (*Options, error) {
	newopts := DefaultOptions()
	return newopts, nil
}

func DefaultOptions() *Options {
	return &Options{
		MachineID:     "devpod",
		MachineFolder: "/tmp/devpod",
		DiskImage:     "ubuntu",
		DiskSize:      "10G",
		MachineType:   "qemu",
		Token:         "",
		JobId:         "",
		NomadBinary:   getEnv("NOMAD_BINARY", "nomad"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
