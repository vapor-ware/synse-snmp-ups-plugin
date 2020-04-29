package devices

import "github.com/vapor-ware/synse-snmp-base/pkg/mibs"

// Plugin device definitions for UPS-MIB "upsAlarm" objects.
//
// See UPS-MIB 1.3.6.1.2.1.33.1.6
// http://www.oidview.com/mibs/0/UPS-MIB.html
// https://tools.ietf.org/html/rfc1628
var (
	UpsAlarmsPresent = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.6.1",
		Info:    "upsAlarmsPresent",
		Type:    "count",
		Output:  "count",
		Handler: "read-only",
	}

	UpsAlarmDescr = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.6.2.1.2",
		Info:    "upsAlarmDescr",
		Type:    "identity",
		Output:  "identity",
		Handler: "read-only",
	}

	UpsAlarmTime = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.6.2.1.3",
		Info:    "upsAlarmTime",
		Type:    "timestamp",
		Output:  "timestamp",
		Handler: "read-only",
	}
)
