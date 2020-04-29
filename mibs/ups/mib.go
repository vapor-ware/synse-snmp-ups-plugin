package ups

import (
	"github.com/vapor-ware/synse-snmp-base/pkg/mibs"
	"github.com/vapor-ware/synse-snmp-ups-plugin/mibs/ups/devices"
)

// Mib is the MIB definition for the UPS MIB.
//
// See also:
//  - http://www.oidview.com/mibs/0/UPS-MIB.html
//  - https://tools.ietf.org/html/rfc1628
var Mib = mibs.NewMIB(
	"UPS-MIB",

	// Alarm Devices
	&devices.UpsAlarmsPresent,
	&devices.UpsAlarmDescr,
	&devices.UpsAlarmTime,

	// Battery Devices
	&devices.UpsBatteryStatus,
	&devices.UpsSecondsOnBattery,
	&devices.UpsEstimatedMinutesRemaining,
	&devices.UpsEstimatedChargeRemaining,
	&devices.UpsBatteryVoltage,
	&devices.UpsBatteryCurrent,
	&devices.UpsBatteryTemperature,

	// Bypass Devices
	&devices.UpsBypassFrequency,
	&devices.UpsBypassNumLines,
	&devices.UpsBypassVoltage,
	&devices.UpsBypassCurrent,
	&devices.UpsBypassPower,

	// Config Devices
	&devices.UpsConfigInputVoltage,
	&devices.UpsConfigInputFreq,
	&devices.UpsConfigOutputVoltage,
	&devices.UpsConfigOutputFreq,
	&devices.UpsConfigOutputVA,
	&devices.UpsConfigOutputPower,
	&devices.UpsConfigLowBattTime,
	&devices.UpsConfigAudibleStatus,
	&devices.UpsConfigLowVoltageTransferPoint,
	&devices.UpsConfigHighVoltageTransferPoint,

	// Control Devices
	&devices.UpsShutdownType,
	&devices.UpsShutdownAfterDelay,
	&devices.UpsStartupAfterDelay,
	&devices.UpsRebootWithDuration,
	&devices.UpsAutoRestart,

	// Identity Devices
	&devices.UpsIdentManufacturer,
	&devices.UpsIdentModel,
	&devices.UpsIdentUPSSoftwareVersion,
	&devices.UpsIdentAgentSoftwareVersion,
	&devices.UpsIdentName,
	&devices.UpsIdentAttachedDevices,

	// Input Devices
	&devices.UpsInputLineBads,
	&devices.UpsInputNumLines,
	&devices.UpsInputFrequency,
	&devices.UpsInputVoltage,
	&devices.UpsInputCurrent,
	&devices.UpsInputTruePower,

	// Output Devices
	&devices.UpsOutputSource,
	&devices.UpsOutputFrequency,
	&devices.UpsOutputNumLines,
	&devices.UpsOutputVoltage,
	&devices.UpsOutputCurrent,
	&devices.UpsOutputPower,
	&devices.UpsOutputPercentLoad,

	// Test Devices
	&devices.UpsTestID,
	&devices.UpsTestSpinLock,
	&devices.UpsTestResultsSummary,
	&devices.UpsTestResultsDetail,
	&devices.UpsTestStartTime,
	&devices.UpsTestElapsedTime,
)
