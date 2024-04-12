package options

import "os"

type Options struct {
	DiskSize string
	Token    string

	JobId       string
	NomadBinary string
	Namespace   string
	Region      string
}

// Read ENV Vars for option overrides
func FromEnv() (*Options, error) {
	newopts := DefaultOptions()
	return newopts, nil
}

func DefaultOptions() *Options {
	return &Options{
		DiskSize:    "10G",
		Token:       "",
		Namespace:   "",
		Region:      "",
		JobId:       "devpod-job",
		NomadBinary: getEnv("NOMAD_BINARY", "nomad"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
