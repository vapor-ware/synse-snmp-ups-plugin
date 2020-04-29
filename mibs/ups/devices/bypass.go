package devices

import (
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-snmp-base/pkg/mibs"
)

// Plugin device definitions for UPS-MIB "upsBypass" objects.
//
// See UPS-MIB 1.3.6.1.2.1.33.1.5
// http://www.oidview.com/mibs/0/UPS-MIB.html
// https://tools.ietf.org/html/rfc1628
var (
	UpsBypassFrequency = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.5.1",
		Info:    "upsBypassFrequency",
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

	UpsBypassNumLines = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.5.2",
		Info:    "upsBypassNumLines",
		Type:    "number",
		Output:  "number",
		Handler: "read-only",
	}

	UpsBypassVoltage = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.5.3.1.2",
		Info:    "upsBypassVoltage",
		Type:    "voltage",
		Output:  "voltage",
		Handler: "read-only",
	}

	UpsBypassCurrent = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.5.3.1.3",
		Info:    "upsBypassCurrent",
		Type:    "current",
		Output:  "electric-current",
		Handler: "read-only",
		Transforms: []sdk.Transformer{
			// Units are "0.1 RMS Amps", so to encode the value as Amps,
			// multiply by 10 (0.1 * 10 = 1).
			&sdk.ScaleTransformer{
				Factor: 10,
			},
		},
	}

	UpsBypassPower = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.5.3.1.4",
		Info:    "upsBypassPower",
		Type:    "power",
		Output:  "watts",
		Handler: "read-only",
	}
)
