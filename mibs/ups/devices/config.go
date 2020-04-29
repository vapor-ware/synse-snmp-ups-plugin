package devices

import (
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-snmp-base/pkg/mibs"
)

// Plugin device definitions for UPS-MIB "upsConfig" objects.
//
// See UPS-MIB 1.3.6.1.2.1.33.1.9
// http://www.oidview.com/mibs/0/UPS-MIB.html
// https://tools.ietf.org/html/rfc1628
var (
	UpsConfigInputVoltage = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.9.1",
		Info:    "upsConfigInputVoltage",
		Type:    "voltage",
		Output:  "voltage",
		Handler: "read-write",
	}

	UpsConfigInputFreq = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.9.2",
		Info:    "upsConfigInputFreq",
		Type:    "frequency",
		Output:  "frequency",
		Handler: "read-write",
		Transforms: []sdk.Transformer{
			// Units are "0.1 Hertz", so to encode the value as Hertz,
			// multiply by 10 (0.1 * 10 = 1).
			&sdk.ScaleTransformer{
				Factor: 10,
			},
		},
	}

	UpsConfigOutputVoltage = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.9.3",
		Info:    "upsConfigOutputVoltage",
		Type:    "voltage",
		Output:  "voltage",
		Handler: "read-write",
	}

	UpsConfigOutputFreq = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.9.4",
		Info:    "upsConfigOutputFreq",
		Type:    "frequency",
		Output:  "frequency",
		Handler: "read-write",
		Transforms: []sdk.Transformer{
			// Units are "0.1 Hertz", so to encode the value as Hertz,
			// multiply by 10 (0.1 * 10 = 1).
			&sdk.ScaleTransformer{
				Factor: 10,
			},
		},
	}

	UpsConfigOutputVA = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.9.5",
		Info:    "upsConfigOutputVA",
		Type:    "power",
		Output:  "volt-ampere",
		Handler: "read-only",
	}

	UpsConfigOutputPower = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.9.6",
		Info:    "upsConfigOutputPower",
		Type:    "power",
		Output:  "watts",
		Handler: "read-only",
	}

	UpsConfigLowBattTime = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.9.7",
		Info:    "upsConfigLowBattTime",
		Type:    "duration",
		Output:  "minutes",
		Handler: "read-write",
	}

	UpsConfigAudibleStatus = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.9.8",
		Info:    "upsConfigAudibleStatus",
		Type:    "status",
		Output:  "status",
		Handler: "read-write",
		Data: map[string]interface{}{
			"enum": map[interface{}]interface{}{
				1: "disabled",
				2: "enabled",
				3: "muted",
			},
		},
	}

	UpsConfigLowVoltageTransferPoint = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.9.9",
		Info:    "upsConfigLowVoltageTransferPoint",
		Type:    "voltage",
		Output:  "voltage",
		Handler: "read-write",
	}

	UpsConfigHighVoltageTransferPoint = mibs.SnmpDevice{
		OID:     "1.3.6.1.2.1.33.1.9.10",
		Info:    "upsConfigHighVoltageTransferPoint",
		Type:    "voltage",
		Output:  "voltage",
		Handler: "read-write",
	}
)
