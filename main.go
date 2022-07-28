package main

import (
	rpcinterfaces "github.com/byzk-project-deploy/base-interface"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func main() {
	rpcinterfaces.PluginServe(func(logger hclog.Logger) *rpcinterfaces.PluginServeCallbackResult {
		return &rpcinterfaces.PluginServeCallbackResult{
			HandshakeConfig: &plugin.HandshakeConfig{
				ProtocolVersion:  1,
				MagicCookieKey:   "BASIC_PLUGIN",
				MagicCookieValue: "hello",
			},
			InfoPlugin: &PluginInfoImpl{
				logger: logger,
			},
		}
	})
}
