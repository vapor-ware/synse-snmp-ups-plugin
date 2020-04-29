package devices

import (
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-snmp-base/pkg/mibs"
)

// Plugin device definitions for UPS-MIB "upsOutput" objects.
//
// See UPS-MIB 1.3.6.1.2.1.33.1.4
// http://www.oidview.com/mibs/0/UPS-MIB.html
// https://tools.ietf.org/html/rfc1628
var (
	UpsOutputSource = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.4.1",
		Info:    "upsOutputSource",
		Type:    "status",
		Output:  "status",
		Handler: "read-only",
		Data: map[string]interface{}{
			"enum": map[interface{}]interface{}{
				1: "other",
				2: "none",
				3: "normal",
				4: "bypass",
				5: "battery",
				6: "booster",
				7: "reducer",
			},
		},
	}

	UpsOutputFrequency = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.4.2",
		Info:    "upsOutputFrequency",
		Type:    "frequency",
		Output:  "frequency",
		Handler: "read-only",
		Transforms: []sdk.Transformer{
			// Units are "0.1 Hertz", so to encode the value as Hertz,
			// multiply by 10 (0.1 * 10 = 1).
			&sdk.ScaleTransformer{
				Factor: 10,
			},
		},
	}

	UpsOutputNumLines = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.4.3",
		Info:    "upsOutputNumLines",
		Type:    "number",
		Output:  "number",
		Handler: "read-only",
	}

	UpsOutputVoltage = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.4.4.1.2",
		Info:    "upsOutputVoltage",
		Type:    "voltage",
		Output:  "voltage",
		Handler: "read-only",
	}

	UpsOutputCurrent = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.4.4.1.3",
		Info:    "upsOutputCurrent",
		Type:    "current",
		Output:  "electric-current",
		Handler: "read-only",
		Transforms: []sdk.Transformer{
			// Units are "0.1 RMS Amp", so to encode the value as Amps,
			// multiply by 10 (0.1 * 10 = 1).
			&sdk.ScaleTransformer{
				Factor: 10,
			},
		},
	}

	UpsOutputPower = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.4.4.1.4",
		Info:    "upsOutputPower",
		Type:    "power",
		Output:  "watts",
		Handler: "read-only",
	}

	UpsOutputPercentLoad = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.4.4.1.5",
		Info:    "upsOutputPercentLoad",
		Type:    "percent",
		Output:  "percentage",
		Handler: "read-only",
	}
)
