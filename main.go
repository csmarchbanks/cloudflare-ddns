package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

func main() {
	configFile := flag.String("config", "config.yaml", "Config file for dynamic dns.")
	flag.Parse()

	cfg, err := loadConfig(*configFile)
	if err != nil {
		fmt.Printf("Error loading config: %+v\n", err)
		os.Exit(1)
	}

	api, err := cloudflare.NewWithAPIToken(cfg.APIToken)
	if err != nil {
		fmt.Printf("Error connecting to cloudflare API: %+v\n", err)
		os.Exit(1)
	}

	ip, err := externalIP()
	if err != nil {
		fmt.Printf("Error getting external IP address: %+v\n", err)
		os.Exit(1)
	}

	records, err := api.DNSRecords(cfg.ZoneID, cloudflare.DNSRecord{
		Name: cfg.DNSRecord,
	})
	if err != nil {
		fmt.Printf("Error listing DNS records: %+v\n", err)
		os.Exit(1)
	}

	if len(records) > 1 {
		fmt.Printf("More than one DNS record found, cannot proceed")
		os.Exit(1)
	}

	dnsRecord := cloudflare.DNSRecord{
		Type:    "A",
		Name:    cfg.DNSRecord,
		Content: ip.String(),
		TTL:     120,
	}

	if len(records) > 0 {
		err := api.UpdateDNSRecord(cfg.ZoneID, records[0].ID, dnsRecord)
		if err != nil {
			fmt.Printf("Error updating record: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("DNS Record successfully updated")
	} else {
		_, err := api.CreateDNSRecord(cfg.ZoneID, dnsRecord)
		if err != nil {
			fmt.Printf("Error creating record: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("DNS Record successfully created")
	}
}

func externalIP() (net.IP, error) {
	r, err := http.Get("https://ifconfig.me/ip")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	ip := net.ParseIP(string(bytes))
	if ip == nil {
		return nil, fmt.Errorf("could not parse %s as an IP address", string(bytes))
	}
	return ip, nil
}
