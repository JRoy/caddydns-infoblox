package infoblox

import (
	infoblox "github.com/JRoy/libdns-infoblox"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

type Provider struct{ *infoblox.Provider }

func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.infoblox",
		New: func() caddy.Module { return &Provider{new(infoblox.Provider)} },
	}
}

func init() {
	caddy.RegisterModule(Provider{})
}

func (p *Provider) Provision(ctx caddy.Context) error {
	p.Host = caddy.NewReplacer().ReplaceAll(p.Host, "")
	p.Version = caddy.NewReplacer().ReplaceAll(p.Version, "")
	p.Username = caddy.NewReplacer().ReplaceAll(p.Username, "")
	p.Password = caddy.NewReplacer().ReplaceAll(p.Password, "")
	p.View = caddy.NewReplacer().ReplaceAll(p.View, "")
	return nil
}

func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}

		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "host":
				if d.NextArg() {
					p.Host = d.Val()
				} else {
					return d.ArgErr()
				}
			case "version":
				if d.NextArg() {
					p.Version = d.Val()
				} else {
					return d.ArgErr()
				}
			case "username":
				if d.NextArg() {
					p.Username = d.Val()
				} else {
					return d.ArgErr()
				}
			case "password":
				if d.NextArg() {
					p.Password = d.Val()
				} else {
					return d.ArgErr()
				}
			case "view":
				if d.NextArg() {
					p.View = d.Val()
				} else {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}

	if p.Host == "" || p.Version == "" || p.Username == "" || p.Password == "" {
		return d.Err("missing config!")
	}

	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
