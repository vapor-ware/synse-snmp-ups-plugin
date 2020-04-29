package devices

import "github.com/vapor-ware/synse-snmp-base/pkg/mibs"

// Plugin device definitions for UPS-MIB "upsControl" objects.
//
// See UPS-MIB 1.3.6.1.2.1.33.1.8
// http://www.oidview.com/mibs/0/UPS-MIB.html
// https://tools.ietf.org/html/rfc1628
var (
	UpsShutdownType = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.8.1",
		Info:    "upsShutdownType",
		Type:    "state",
		Output:  "state",
		Handler: "read-write",
		Data: map[string]interface{}{
			"enum": map[interface{}]interface{}{
				1: "output",
				2: "system",
			},
		},
	}

	UpsShutdownAfterDelay = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.8.2",
		Info:    "upsShutdownAfterDelay",
		Type:    "duration",
		Output:  "seconds",
		Handler: "read-write",
	}

	UpsStartupAfterDelay = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.8.3",
		Info:    "upsStartupAfterDelay",
		Type:    "duration",
		Output:  "seconds",
		Handler: "read-write",
	}

	UpsRebootWithDuration = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.8.4",
		Info:    "upsRebootWithDuration",
		Type:    "duration",
		Output:  "seconds",
		Handler: "read-write",
	}

	UpsAutoRestart = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.8.5",
		Info:    "upsAutoRestart",
		Type:    "state",
		Output:  "state",
		Handler: "read-write",
		Data: map[string]interface{}{
			"enum": map[interface{}]interface{}{
				1: "on",
				2: "off",
			},
		},
	}
)
