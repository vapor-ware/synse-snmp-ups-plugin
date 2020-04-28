package ups

import (
	"github.com/vapor-ware/synse-snmp-base/pkg/mibs"
	"github.com/vapor-ware/synse-snmp-ups-plugin/mibs/ups/devices"
)

// Mib is the MIB definition for the UPS MIB.
//
// See also: http://www.oidview.com/mibs/0/UPS-MIB.html
var Mib = mibs.NewMIB(
	"UPS-MIB",
	// Alarm Devices

	// Battery Devices
	&devices.UpsBatteryStatus,
	&devices.UpsSecondsOnBattery,
	&devices.UpsEstimatedMinutesRemaining,
	&devices.UpsEstimatedChargeRemaining,
	&devices.UpsBatteryVoltage,
	&devices.UpsBatteryCurrent,
	&devices.UpsBatteryTemperature,

	// Bypass Devices

	// Config Devices

	// Control Devices

	// Identity Devices
	&devices.UpsIdentManufacturer,
	&devices.UpsIdentModel,
	&devices.UpsIdentUPSSoftwareVersion,
	&devices.UpsIdentAgentSoftwareVersion,
	&devices.UpsIdentName,
	&devices.UpsIdentAttachedDevices,

	// Input Devices

	// Output Devices

	// Test Devices
)
