# Cloudflare Dynamic DNS

This is a simple program to update a [Cloudflare](https://www.cloudflare.com/)
A record with the external IP address of the server it is being run on. This is
targeted at allowing access to home computers behind a router.

## Installation

```
go get github.com/csmarchbanks/cloudflare-ddns
```

## Usage

```
cloudflare-ddns -config config.yaml
```

Commonly this will be run in cron, for example to update the DNS entry every minute:
```
* * * * * /path/to/binary/cloudflare-ddns --config /path/to/config.yaml
```

## Configuration

A configuration file with the following is required to run this:
```
api_token: <your_cloudflare_api_token>
zone_id: <your_cloudflare_zone_id>
dns_record: target.dns.example.com
```

You can find the Zone ID and create an API Token from your Cloudflare dashboard.
