package devices

import (
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-snmp-base/pkg/mibs"
)

// Plugin device definitions for UPS-MIB "upsBattery" objects.
//
// See UPS-MIB 1.3.6.1.2.1.33.1.2
// http://www.oidview.com/mibs/0/UPS-MIB.html
// https://tools.ietf.org/html/rfc1628
var (
	UpsBatteryStatus = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.2.1.0",
		Info:    "upsBatteryStatus",
		Type:    "status",
		Output:  "status",
		Handler: "read-only",
		Data: map[string]interface{}{
			"enum": map[interface{}]interface{}{
				1: "unknown",
				2: "batteryNormal",
				3: "batteryLow",
				4: "batteryDepleted",
			},
		},
	}

	UpsSecondsOnBattery = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.2.2.0",
		Info:    "upsSecondsOnBattery",
		Type:    "duration",
		Output:  "seconds",
		Handler: "read-only",
	}

	UpsEstimatedMinutesRemaining = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.2.3.0",
		Info:    "upsEstimatedMinutesRemaining",
		Type:    "duration",
		Output:  "minutes",
		Handler: "read-only",
	}

	UpsEstimatedChargeRemaining = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.2.4.0",
		Info:    "upsEstimatedChargeRemaining",
		Type:    "percent",
		Output:  "percentage",
		Handler: "read-only",
	}

	UpsBatteryVoltage = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.2.5.0",
		Info:    "upsBatteryVoltage",
		Type:    "voltage",
		Output:  "voltage",
		Handler: "read-only",
		Transforms: []sdk.Transformer{
			// Units are "0.1 Volt DC", so to encode the value as Volts,
			// multiply by 10 (0.1 * 10 = 1).
			&sdk.ScaleTransformer{
				Factor: 10,
			},
		},
	}

	UpsBatteryCurrent = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.2.6.0",
		Info:    "upsBatteryCurrent",
		Type:    "current",
		Output:  "electric-current",
		Handler: "read-only",
		Transforms: []sdk.Transformer{
			// Units are "0.1 Amp DC", so to encode the value as Amps,
			// multiply by 10 (0.1 * 10 = 1).
			&sdk.ScaleTransformer{
				Factor: 10,
			},
		},
	}

	UpsBatteryTemperature = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.2.7.0",
		Info:    "upsBatteryTemperature",
		Type:    "temperature",
		Output:  "temperature",
		Handler: "read-only",
	}
)
