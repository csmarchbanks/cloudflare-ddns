package main

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type config struct {
	APIToken   string   `yaml:"api_token"`
	ZoneID     string   `yaml:"zone_id"`
	DNSRecords []string `yaml:"dns_records"`
}

func loadConfig(filename string) (*config, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &config{}
	err = yaml.Unmarshal(bytes, c)
	if err != nil {
		return nil, err
	}

	if c.APIToken == "" {
		return nil, errors.New("must provide an API Token")
	}
	if c.ZoneID == "" {
		return nil, errors.New("must provide a Zone ID")
	}
	if len(c.DNSRecords) == 0 {
		return nil, errors.New("must provide at least one DNS Record")
	}
	return c, err
}
