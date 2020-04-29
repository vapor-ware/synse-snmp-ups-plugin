package devices

import (
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-snmp-base/pkg/mibs"
)

// Plugin device definitions for UPS-MIB "upsInput" objects.
//
// See UPS-MIB 1.3.6.1.2.1.33.1.3
// http://www.oidview.com/mibs/0/UPS-MIB.html
// https://tools.ietf.org/html/rfc1628
var (
	UpsInputLineBads = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.3.1",
		Info:    "upsInputLineBads",
		Type:    "count",
		Output:  "count",
		Handler: "read-only",
	}

	UpsInputNumLines = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.3.2",
		Info:    "upsInputNumLines",
		Type:    "number",
		Output:  "number",
		Handler: "read-only",
	}

	UpsInputFrequency = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.3.3.1.2",
		Info:    "upsInputFrequency",
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

	UpsInputVoltage = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.3.3.1.3",
		Info:    "upsInputVoltage",
		Type:    "voltage",
		Output:  "voltage",
		Handler: "read-only",
	}

	UpsInputCurrent = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.3.3.1.4",
		Info:    "upsInputCurrent",
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

	UpsInputTruePower = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.3.3.1.5",
		Info:    "upsInputTruePower",
		Type:    "power",
		Output:  "watt",
		Handler: "read-only",
	}
)
