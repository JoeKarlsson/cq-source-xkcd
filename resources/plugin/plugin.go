package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Name    = "xkcd"
	Kind    = "source"
	Team    = "JoeKarlsson"
	Version = "development"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(Name, Version, Configure, plugin.WithKind(Kind), plugin.WithTeam(Team))
}
