package main

import (
	"github.com/byzk-project-deploy/base-interface"
	"github.com/hashicorp/go-hclog"
	"time"
)

var createTime, _ = time.Parse("2006-01-02 15:04:05", "2022-07-28 14:58:00")

var pluginInfo = &rpcinterfaces.PluginInfo{
	Author:     "无痕",
	Name:       "docker swarm extends",
	ShortDesc:  "基于DockerSwarm的扩展",
	Desc:       "基于DockerSwarm的扩展",
	CreateTime: createTime,
	Type:       rpcinterfaces.PluginTypeCmd,
}

type PluginInfoImpl struct {
	logger hclog.Logger
}

func (p *PluginInfoImpl) Info() (*rpcinterfaces.PluginInfo, error) {
	return pluginInfo, nil
}
