package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-snmp-base/pkg/mibs"
	plugin "github.com/vapor-ware/synse-snmp-base/pkg/plugin"
	"github.com/vapor-ware/synse-snmp-ups-plugin/mibs/ups"
	"github.com/vapor-ware/synse-snmp-ups-plugin/outputs"
)

// MakePlugin creates a new instance of the plugin.
//
// It ensures all plugin-specific configuration is done and all necessary items
// are registered with the plugin.
func MakePlugin() *sdk.Plugin {

	// Build a base SNMP plugin instance.
	p, err := plugin.NewSnmpBasePlugin(&plugin.PluginMetadata{
		Name:        "RFC1628 UPS SNMP",
		Maintainer:  "vaporio",
		Description: "SNMP plugin for the RFC1628 UPS MIB",
		VCS:         "https://github.com/vapor-ware/synse-snmp-ups-plugin",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Register custom output types.
	err = p.RegisterOutputs(
		&outputs.Identity,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Register the UPS MIB.
	err = mibs.Register(
		ups.Mib,
	)
	if err != nil {
		log.Fatal(err)
	}

	return p
}

func main() {
	p := MakePlugin()

	if err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
