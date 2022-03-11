// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs /home/ngc-auth-ldap-kklues/go-nvml/pkg/nvml/types.go

package nvml

import "unsafe"

type Device struct {
	Handle *_Ctype_struct_nvmlDevice_st
}

type PciInfo struct {
	BusIdLegacy    [16]int8
	Domain         uint32
	Bus            uint32
	Device         uint32
	PciDeviceId    uint32
	PciSubSystemId uint32
	BusId          [32]int8
}

type EccErrorCounts struct {
	L1Cache      uint64
	L2Cache      uint64
	DeviceMemory uint64
	RegisterFile uint64
}

type Utilization struct {
	Gpu    uint32
	Memory uint32
}

type Memory struct {
	Total uint64
	Free  uint64
	Used  uint64
}

type BAR1Memory struct {
	Bar1Total uint64
	Bar1Free  uint64
	Bar1Used  uint64
}

type ProcessInfo_v1 struct {
	Pid           uint32
	UsedGpuMemory uint64
}

type ProcessInfo struct {
	Pid               uint32
	UsedGpuMemory     uint64
	GpuInstanceId     uint32
	ComputeInstanceId uint32
}

type DeviceAttributes struct {
	MultiprocessorCount       uint32
	SharedCopyEngineCount     uint32
	SharedDecoderCount        uint32
	SharedEncoderCount        uint32
	SharedJpegCount           uint32
	SharedOfaCount            uint32
	GpuInstanceSliceCount     uint32
	ComputeInstanceSliceCount uint32
	MemorySizeMB              uint64
}

type RowRemapperHistogramValues struct {
	Max     uint32
	High    uint32
	Partial uint32
	Low     uint32
	None    uint32
}

type NvLinkUtilizationControl struct {
	Units     uint32
	Pktfilter uint32
}

type BridgeChipInfo struct {
	Type      uint32
	FwVersion uint32
}

type BridgeChipHierarchy struct {
	BridgeCount    uint8
	BridgeChipInfo [128]BridgeChipInfo
}

const sizeofValue = unsafe.Sizeof([8]byte{})

type Value [sizeofValue]byte

type Sample struct {
	TimeStamp   uint64
	SampleValue [8]byte
}

type ViolationTime struct {
	ReferenceTime uint64
	ViolationTime uint64
}

type ClkMonFaultInfo struct {
	ClkApiDomain       uint32
	ClkDomainFaultMask uint32
}

type ClkMonStatus struct {
	BGlobalStatus  uint32
	ClkMonListSize uint32
	ClkMonList     [32]ClkMonFaultInfo
}

type VgpuTypeId uint32

type VgpuInstance uint32

type VgpuInstanceUtilizationSample struct {
	VgpuInstance uint32
	TimeStamp    uint64
	SmUtil       [8]byte
	MemUtil      [8]byte
	EncUtil      [8]byte
	DecUtil      [8]byte
}

type VgpuProcessUtilizationSample struct {
	VgpuInstance uint32
	Pid          uint32
	ProcessName  [64]int8
	TimeStamp    uint64
	SmUtil       uint32
	MemUtil      uint32
	EncUtil      uint32
	DecUtil      uint32
}

type ProcessUtilizationSample struct {
	Pid       uint32
	TimeStamp uint64
	SmUtil    uint32
	MemUtil   uint32
	EncUtil   uint32
	DecUtil   uint32
}

type GridLicenseExpiry struct {
	Year      uint32
	Month     uint16
	Day       uint16
	Hour      uint16
	Min       uint16
	Sec       uint16
	Status    uint8
	Pad_cgo_0 [1]byte
}

type GridLicensableFeature struct {
	FeatureCode    uint32
	FeatureState   uint32
	LicenseInfo    [128]int8
	ProductName    [128]int8
	FeatureEnabled uint32
	LicenseExpiry  GridLicenseExpiry
}

type GridLicensableFeatures struct {
	IsGridLicenseSupported  int32
	LicensableFeaturesCount uint32
	GridLicensableFeatures  [3]GridLicensableFeature
}

type DeviceArchitecture uint32

type FieldValue struct {
	FieldId     uint32
	ScopeId     uint32
	Timestamp   int64
	LatencyUsec int64
	ValueType   uint32
	NvmlReturn  uint32
	Value       [8]byte
}

type Unit struct {
	Handle *_Ctype_struct_nvmlUnit_st
}

type HwbcEntry struct {
	HwbcId          uint32
	FirmwareVersion [32]int8
}

type LedState struct {
	Cause [256]int8
	Color uint32
}

type UnitInfo struct {
	Name            [96]int8
	Id              [96]int8
	Serial          [96]int8
	FirmwareVersion [96]int8
}

type PSUInfo struct {
	State   [256]int8
	Current uint32
	Voltage uint32
	Power   uint32
}

type UnitFanInfo struct {
	Speed uint32
	State uint32
}

type UnitFanSpeeds struct {
	Fans  [24]UnitFanInfo
	Count uint32
}

type EventSet struct {
	Handle *_Ctype_struct_nvmlEventSet_st
}

type EventData struct {
	Device            Device
	EventType         uint64
	EventData         uint64
	GpuInstanceId     uint32
	ComputeInstanceId uint32
}

type AccountingStats struct {
	GpuUtilization    uint32
	MemoryUtilization uint32
	MaxMemoryUsage    uint64
	Time              uint64
	StartTime         uint64
	IsRunning         uint32
	Reserved          [5]uint32
}

type EncoderSessionInfo struct {
	SessionId      uint32
	Pid            uint32
	VgpuInstance   uint32
	CodecType      uint32
	HResolution    uint32
	VResolution    uint32
	AverageFps     uint32
	AverageLatency uint32
}

type FBCStats struct {
	SessionsCount  uint32
	AverageFPS     uint32
	AverageLatency uint32
}

type FBCSessionInfo struct {
	SessionId      uint32
	Pid            uint32
	VgpuInstance   uint32
	DisplayOrdinal uint32
	SessionType    uint32
	SessionFlags   uint32
	HMaxResolution uint32
	VMaxResolution uint32
	HResolution    uint32
	VResolution    uint32
	AverageFPS     uint32
	AverageLatency uint32
}

type AffinityScope uint32

type VgpuVersion struct {
	MinVersion uint32
	MaxVersion uint32
}

type nvmlVgpuMetadata struct {
	Version                uint32
	Revision               uint32
	GuestInfoState         uint32
	GuestDriverVersion     [80]int8
	HostDriverVersion      [80]int8
	Reserved               [6]uint32
	VgpuVirtualizationCaps uint32
	GuestVgpuVersion       uint32
	OpaqueDataSize         uint32
	OpaqueData             [4]int8
}

type nvmlVgpuPgpuMetadata struct {
	Version                uint32
	Revision               uint32
	HostDriverVersion      [80]int8
	PgpuVirtualizationCaps uint32
	Reserved               [5]uint32
	HostSupportedVgpuRange VgpuVersion
	OpaqueDataSize         uint32
	OpaqueData             [4]int8
}

type VgpuPgpuCompatibility struct {
	VgpuVmCompatibility    uint32
	CompatibilityLimitCode uint32
}

type ExcludedDeviceInfo struct {
	PciInfo PciInfo
	Uuid    [80]int8
}

type GpuInstancePlacement struct {
	Start uint32
	Size  uint32
}

type GpuInstanceProfileInfo struct {
	Id                  uint32
	IsP2pSupported      uint32
	SliceCount          uint32
	InstanceCount       uint32
	MultiprocessorCount uint32
	CopyEngineCount     uint32
	DecoderCount        uint32
	EncoderCount        uint32
	JpegCount           uint32
	OfaCount            uint32
	MemorySizeMB        uint64
}

type GpuInstanceInfo struct {
	Device    Device
	Id        uint32
	ProfileId uint32
	Placement GpuInstancePlacement
}

type GpuInstance struct {
	Handle *_Ctype_struct_nvmlGpuInstance_st
}

type ComputeInstancePlacement struct {
	Start uint32
	Size  uint32
}

type ComputeInstanceProfileInfo struct {
	Id                    uint32
	SliceCount            uint32
	InstanceCount         uint32
	MultiprocessorCount   uint32
	SharedCopyEngineCount uint32
	SharedDecoderCount    uint32
	SharedEncoderCount    uint32
	SharedJpegCount       uint32
	SharedOfaCount        uint32
}

type ComputeInstanceInfo struct {
	Device      Device
	GpuInstance GpuInstance
	Id          uint32
	ProfileId   uint32
	Placement   ComputeInstancePlacement
}

type ComputeInstance struct {
	Handle *_Ctype_struct_nvmlComputeInstance_st
}
