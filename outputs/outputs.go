package outputs

import "github.com/vapor-ware/synse-sdk/sdk/output"

var (
	// Identity describes readings with identity outputs.
	Identity = output.Output{
		Name: "identity",
		Type: "identity",
	}

	// VoltAmpere describes readings with values for apparent power.
	VoltAmpere = output.Output{
		Name: "volt-ampere",
		Type: "power",
		Precision: 4,
		Unit: &output.Unit{
			Name: "volt-ampere",
			Symbol: "VA",
		},
	}
)
