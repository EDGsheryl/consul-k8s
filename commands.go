package main

import (
	"os"

	cmdACLInit "github.com/hashicorp/consul-k8s/subcommand/acl-init"
	cmdInjectConnect "github.com/hashicorp/consul-k8s/subcommand/inject-connect"
	cmdServerACLInit "github.com/hashicorp/consul-k8s/subcommand/server-acl-init"
	cmdSyncCatalog "github.com/hashicorp/consul-k8s/subcommand/sync-catalog"
	cmdVersion "github.com/hashicorp/consul-k8s/subcommand/version"
	"github.com/hashicorp/consul-k8s/version"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all available consul-k8s commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}

	Commands = map[string]cli.CommandFactory{
		"acl-init": func() (cli.Command, error) {
			return &cmdACLInit.Command{UI: ui}, nil
		},

		"inject-connect": func() (cli.Command, error) {
			return &cmdInjectConnect.Command{UI: ui}, nil
		},

		"server-acl-init": func() (cli.Command, error) {
			return &cmdServerACLInit.Command{UI: ui}, nil
		},

		"sync-catalog": func() (cli.Command, error) {
			return &cmdSyncCatalog.Command{UI: ui}, nil
		},

		"version": func() (cli.Command, error) {
			return &cmdVersion.Command{UI: ui, Version: version.GetHumanVersion()}, nil
		},
	}
}

func helpFunc() cli.HelpFunc {
	// This should be updated for any commands we want to hide for any reason.
	// Hidden commands can still be executed if you know the command, but
	// aren't shown in any help output. We use this for prerelease functionality
	// or advanced features.
	hidden := map[string]struct{}{
		"inject-connect": struct{}{},
	}

	var include []string
	for k := range Commands {
		if _, ok := hidden[k]; !ok {
			include = append(include, k)
		}
	}

	return cli.FilteredHelpFunc(include, cli.BasicHelpFunc("consul-k8s"))
}
