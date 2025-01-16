Infoblox module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with Infoblox accounts.

## Caddy module name

```
dns.providers.infoblox
```

## Configuration

### JSON Example

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuers/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "infoblox",
				"host": "INFOBLOX_HOST",
                "version": "INFOBLOX_VERSION",
                "username": "INFOBLOX_USERNAME",
                "password": "INFOBLOX_PASSWORD"
			}
		}
	}
}
```

### Caddyfile Examples

```Caddyfile
tls {
	dns infoblox {
		host {env.INFOBLOX_HOST}
		version {env.INFOBLOX_VERSION}
		username {env.INFOBLOX_USERNAME}
		password {env.INFOBLOX_PASSWORD}
	}
}
```