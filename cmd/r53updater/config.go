package main

import (
	"errors"
	"flag"
	"log"
	"os"
)

// Config cmdline config struct
type Config struct {
	fqdn string
	zone string
	tick int
}

// NewConfig parse cmdline args and env vars for config
func NewConfig() Config {
	cmdFQDN := flag.String("fqdn", "", "fqdn")
	cmdZone := flag.String("zone", "", "aws route53 zone")
	cmdTick := flag.Int("tick", 600, "interval in seconds between each ip check")

	flag.Parse()

	fqdn := ""
	if *cmdFQDN != "" {
		fqdn = *cmdFQDN
	} else {
		fqdn = os.Getenv("AWSFQDN")
		if fqdn == "" {
			log.Fatal(errors.New("fqdn not provided"))
		}
	}

	zone := ""
	if *cmdZone != "" {
		zone = *cmdZone
	} else {
		zone = os.Getenv("AWSZONE")
		if zone == "" {
			log.Fatal(errors.New("zone not provided"))
		}
	}

	return Config{
		fqdn: fqdn,
		zone: zone,
		tick: *cmdTick,
	}
}
