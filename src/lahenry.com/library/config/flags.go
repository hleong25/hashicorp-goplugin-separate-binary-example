package config

import "flag"

type Flags struct {
	RunAsPlugin bool
}

func GetFlags() Flags {
	var flags Flags

	flag.BoolVar(&flags.RunAsPlugin, "plugin", true, "Run as a plugin")

	flag.Parse()

	return flags
}
