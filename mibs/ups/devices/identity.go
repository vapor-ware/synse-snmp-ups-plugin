package devices

import "github.com/vapor-ware/synse-snmp-base/pkg/mibs"

// Plugin device definitions for UPS-MIB "upsIdent" objects.
//
// See UPS-MIB 1.3.6.1.2.1.33.1.1
// http://www.oidview.com/mibs/0/UPS-MIB.html
// https://tools.ietf.org/html/rfc1628
var (
	UpsIdentManufacturer = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.1.1.0",
		Info:    "upsIdentManufacturer",
		Type:    "string",
		Output:  "string",
		Handler: "read-only",
	}

	UpsIdentModel = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.1.2.0",
		Info:    "upsIdentModel",
		Type:    "string",
		Output:  "string",
		Handler: "read-only",
	}

	UpsIdentUPSSoftwareVersion = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.1.3.0",
		Info:    "upsIdentUPSSoftwareVersion",
		Type:    "string",
		Output:  "string",
		Handler: "read-only",
	}

	UpsIdentAgentSoftwareVersion = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.1.4.0",
		Info:    "upsIdentAgentSoftwareVersion",
		Type:    "string",
		Output:  "string",
		Handler: "read-only",
	}

	UpsIdentName = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.1.5.0",
		Info:    "upsIdentName",
		Type:    "string",
		Output:  "string",
		Handler: "read-write",
	}

	UpsIdentAttachedDevices = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.1.6.0",
		Info:    "upsIdentAttachedDevices",
		Type:    "string",
		Output:  "string",
		Handler: "read-write",
	}
)
