package devices

import "github.com/vapor-ware/synse-snmp-base/pkg/mibs"

// Plugin device definitions for UPS-MIB "upsTest" objects.
//
// See UPS-MIB 1.3.6.1.2.1.33.1.7
// http://www.oidview.com/mibs/0/UPS-MIB.html
// https://tools.ietf.org/html/rfc1628
var (
	UpsTestID = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.7.1",
		Info:    "upsTestId",
		Type:    "identity",
		Output:  "identity",
		Handler: "read-write",
	}

	UpsTestSpinLock = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.7.2",
		Info:    "upsTestSpinLock",
		Type:    "number",
		Output:  "number",
		Handler: "read-write",
	}

	UpsTestResultsSummary = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.7.3",
		Info:    "upsTestResultsSummary",
		Type:    "status",
		Output:  "status",
		Handler: "read-only",
		Data: map[string]interface{}{
			"enum": map[interface{}]interface{}{
				1: "donePass",
				2: "doneWarning",
				3: "doneError",
				4: "aborted",
				5: "inProgress",
				6: "noTestsInitiated",
			},
		},
	}

	UpsTestResultsDetail = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.7.4",
		Info:    "upsTestResultsDetail",
		Type:    "string",
		Output:  "string",
		Handler: "read-only",
	}

	UpsTestStartTime = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.7.5",
		Info:    "upsTestStartTime",
		Type:    "timestamp",
		Output:  "timestamp",
		Handler: "read-only",
	}

	UpsTestElapsedTime = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.7.6",
		Info:    "upsTestElapsedTime",
		Type:    "duration",
		Output:  "number",
		Handler: "read-only",
	}
)
