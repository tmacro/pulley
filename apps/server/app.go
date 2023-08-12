package server

import "github.com/caddyserver/caddy/v2"

func init() {
	caddy.RegisterModule(PulleyServer{})
}

type PulleyServer struct {
}

func (PulleyServer) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "pulley.server",
		New: func() caddy.Module { return new(PulleyServer) },
	}
}
