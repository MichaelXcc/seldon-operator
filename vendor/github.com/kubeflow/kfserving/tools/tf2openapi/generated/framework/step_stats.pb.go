// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/step_stats.proto

package framework

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// An allocation/de-allocation operation performed by the allocator.
type AllocationRecord struct {
	// The timestamp of the operation.
	AllocMicros int64 `protobuf:"varint,1,opt,name=alloc_micros,json=allocMicros,proto3" json:"alloc_micros,omitempty"`
	// Number of bytes allocated, or de-allocated if negative.
	AllocBytes           int64    `protobuf:"varint,2,opt,name=alloc_bytes,json=allocBytes,proto3" json:"alloc_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllocationRecord) Reset()         { *m = AllocationRecord{} }
func (m *AllocationRecord) String() string { return proto.CompactTextString(m) }
func (*AllocationRecord) ProtoMessage()    {}
func (*AllocationRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e915309f7ed52e5, []int{0}
}

func (m *AllocationRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocationRecord.Unmarshal(m, b)
}
func (m *AllocationRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocationRecord.Marshal(b, m, deterministic)
}
func (m *AllocationRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocationRecord.Merge(m, src)
}
func (m *AllocationRecord) XXX_Size() int {
	return xxx_messageInfo_AllocationRecord.Size(m)
}
func (m *AllocationRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocationRecord.DiscardUnknown(m)
}

var xxx_messageInfo_AllocationRecord proto.InternalMessageInfo

func (m *AllocationRecord) GetAllocMicros() int64 {
	if m != nil {
		return m.AllocMicros
	}
	return 0
}

func (m *AllocationRecord) GetAllocBytes() int64 {
	if m != nil {
		return m.AllocBytes
	}
	return 0
}

type AllocatorMemoryUsed struct {
	AllocatorName string `protobuf:"bytes,1,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
	// These are per-node allocator memory stats.
	TotalBytes int64 `protobuf:"varint,2,opt,name=total_bytes,json=totalBytes,proto3" json:"total_bytes,omitempty"`
	PeakBytes  int64 `protobuf:"varint,3,opt,name=peak_bytes,json=peakBytes,proto3" json:"peak_bytes,omitempty"`
	// The bytes that are not deallocated.
	LiveBytes int64 `protobuf:"varint,4,opt,name=live_bytes,json=liveBytes,proto3" json:"live_bytes,omitempty"`
	// The allocation and deallocation timeline.
	AllocationRecords []*AllocationRecord `protobuf:"bytes,6,rep,name=allocation_records,json=allocationRecords,proto3" json:"allocation_records,omitempty"`
	// These are snapshots of the overall allocator memory stats.
	// The number of live bytes currently allocated by the allocator.
	AllocatorBytesInUse  int64    `protobuf:"varint,5,opt,name=allocator_bytes_in_use,json=allocatorBytesInUse,proto3" json:"allocator_bytes_in_use,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllocatorMemoryUsed) Reset()         { *m = AllocatorMemoryUsed{} }
func (m *AllocatorMemoryUsed) String() string { return proto.CompactTextString(m) }
func (*AllocatorMemoryUsed) ProtoMessage()    {}
func (*AllocatorMemoryUsed) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e915309f7ed52e5, []int{1}
}

func (m *AllocatorMemoryUsed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocatorMemoryUsed.Unmarshal(m, b)
}
func (m *AllocatorMemoryUsed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocatorMemoryUsed.Marshal(b, m, deterministic)
}
func (m *AllocatorMemoryUsed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocatorMemoryUsed.Merge(m, src)
}
func (m *AllocatorMemoryUsed) XXX_Size() int {
	return xxx_messageInfo_AllocatorMemoryUsed.Size(m)
}
func (m *AllocatorMemoryUsed) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocatorMemoryUsed.DiscardUnknown(m)
}

var xxx_messageInfo_AllocatorMemoryUsed proto.InternalMessageInfo

func (m *AllocatorMemoryUsed) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

func (m *AllocatorMemoryUsed) GetTotalBytes() int64 {
	if m != nil {
		return m.TotalBytes
	}
	return 0
}

func (m *AllocatorMemoryUsed) GetPeakBytes() int64 {
	if m != nil {
		return m.PeakBytes
	}
	return 0
}

func (m *AllocatorMemoryUsed) GetLiveBytes() int64 {
	if m != nil {
		return m.LiveBytes
	}
	return 0
}

func (m *AllocatorMemoryUsed) GetAllocationRecords() []*AllocationRecord {
	if m != nil {
		return m.AllocationRecords
	}
	return nil
}

func (m *AllocatorMemoryUsed) GetAllocatorBytesInUse() int64 {
	if m != nil {
		return m.AllocatorBytesInUse
	}
	return 0
}

// Output sizes recorded for a single execution of a graph node.
type NodeOutput struct {
	Slot                 int32              `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	TensorDescription    *TensorDescription `protobuf:"bytes,3,opt,name=tensor_description,json=tensorDescription,proto3" json:"tensor_description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NodeOutput) Reset()         { *m = NodeOutput{} }
func (m *NodeOutput) String() string { return proto.CompactTextString(m) }
func (*NodeOutput) ProtoMessage()    {}
func (*NodeOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e915309f7ed52e5, []int{2}
}

func (m *NodeOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeOutput.Unmarshal(m, b)
}
func (m *NodeOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeOutput.Marshal(b, m, deterministic)
}
func (m *NodeOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeOutput.Merge(m, src)
}
func (m *NodeOutput) XXX_Size() int {
	return xxx_messageInfo_NodeOutput.Size(m)
}
func (m *NodeOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeOutput.DiscardUnknown(m)
}

var xxx_messageInfo_NodeOutput proto.InternalMessageInfo

func (m *NodeOutput) GetSlot() int32 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *NodeOutput) GetTensorDescription() *TensorDescription {
	if m != nil {
		return m.TensorDescription
	}
	return nil
}

// For memory tracking.
type MemoryStats struct {
	TempMemorySize                 int64    `protobuf:"varint,1,opt,name=temp_memory_size,json=tempMemorySize,proto3" json:"temp_memory_size,omitempty"`
	PersistentMemorySize           int64    `protobuf:"varint,3,opt,name=persistent_memory_size,json=persistentMemorySize,proto3" json:"persistent_memory_size,omitempty"`
	PersistentTensorAllocIds       []int64  `protobuf:"varint,5,rep,packed,name=persistent_tensor_alloc_ids,json=persistentTensorAllocIds,proto3" json:"persistent_tensor_alloc_ids,omitempty"`
	DeviceTempMemorySize           int64    `protobuf:"varint,2,opt,name=device_temp_memory_size,json=deviceTempMemorySize,proto3" json:"device_temp_memory_size,omitempty"`                                        // Deprecated: Do not use.
	DevicePersistentMemorySize     int64    `protobuf:"varint,4,opt,name=device_persistent_memory_size,json=devicePersistentMemorySize,proto3" json:"device_persistent_memory_size,omitempty"`                      // Deprecated: Do not use.
	DevicePersistentTensorAllocIds []int64  `protobuf:"varint,6,rep,packed,name=device_persistent_tensor_alloc_ids,json=devicePersistentTensorAllocIds,proto3" json:"device_persistent_tensor_alloc_ids,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *MemoryStats) Reset()         { *m = MemoryStats{} }
func (m *MemoryStats) String() string { return proto.CompactTextString(m) }
func (*MemoryStats) ProtoMessage()    {}
func (*MemoryStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e915309f7ed52e5, []int{3}
}

func (m *MemoryStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemoryStats.Unmarshal(m, b)
}
func (m *MemoryStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemoryStats.Marshal(b, m, deterministic)
}
func (m *MemoryStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemoryStats.Merge(m, src)
}
func (m *MemoryStats) XXX_Size() int {
	return xxx_messageInfo_MemoryStats.Size(m)
}
func (m *MemoryStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MemoryStats.DiscardUnknown(m)
}

var xxx_messageInfo_MemoryStats proto.InternalMessageInfo

func (m *MemoryStats) GetTempMemorySize() int64 {
	if m != nil {
		return m.TempMemorySize
	}
	return 0
}

func (m *MemoryStats) GetPersistentMemorySize() int64 {
	if m != nil {
		return m.PersistentMemorySize
	}
	return 0
}

func (m *MemoryStats) GetPersistentTensorAllocIds() []int64 {
	if m != nil {
		return m.PersistentTensorAllocIds
	}
	return nil
}

// Deprecated: Do not use.
func (m *MemoryStats) GetDeviceTempMemorySize() int64 {
	if m != nil {
		return m.DeviceTempMemorySize
	}
	return 0
}

// Deprecated: Do not use.
func (m *MemoryStats) GetDevicePersistentMemorySize() int64 {
	if m != nil {
		return m.DevicePersistentMemorySize
	}
	return 0
}

// Deprecated: Do not use.
func (m *MemoryStats) GetDevicePersistentTensorAllocIds() []int64 {
	if m != nil {
		return m.DevicePersistentTensorAllocIds
	}
	return nil
}

// Time/size stats recorded for a single execution of a graph node.
type NodeExecStats struct {
	// TODO(tucker): Use some more compact form of node identity than
	// the full string name.  Either all processes should agree on a
	// global id (cost_id?) for each node, or we should use a hash of
	// the name.
	NodeName             string                   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	AllStartMicros       int64                    `protobuf:"varint,2,opt,name=all_start_micros,json=allStartMicros,proto3" json:"all_start_micros,omitempty"`
	OpStartRelMicros     int64                    `protobuf:"varint,3,opt,name=op_start_rel_micros,json=opStartRelMicros,proto3" json:"op_start_rel_micros,omitempty"`
	OpEndRelMicros       int64                    `protobuf:"varint,4,opt,name=op_end_rel_micros,json=opEndRelMicros,proto3" json:"op_end_rel_micros,omitempty"`
	AllEndRelMicros      int64                    `protobuf:"varint,5,opt,name=all_end_rel_micros,json=allEndRelMicros,proto3" json:"all_end_rel_micros,omitempty"`
	Memory               []*AllocatorMemoryUsed   `protobuf:"bytes,6,rep,name=memory,proto3" json:"memory,omitempty"`
	Output               []*NodeOutput            `protobuf:"bytes,7,rep,name=output,proto3" json:"output,omitempty"`
	TimelineLabel        string                   `protobuf:"bytes,8,opt,name=timeline_label,json=timelineLabel,proto3" json:"timeline_label,omitempty"`
	ScheduledMicros      int64                    `protobuf:"varint,9,opt,name=scheduled_micros,json=scheduledMicros,proto3" json:"scheduled_micros,omitempty"`
	ThreadId             uint32                   `protobuf:"varint,10,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
	ReferencedTensor     []*AllocationDescription `protobuf:"bytes,11,rep,name=referenced_tensor,json=referencedTensor,proto3" json:"referenced_tensor,omitempty"`
	MemoryStats          *MemoryStats             `protobuf:"bytes,12,opt,name=memory_stats,json=memoryStats,proto3" json:"memory_stats,omitempty"`
	AllStartNanos        int64                    `protobuf:"varint,13,opt,name=all_start_nanos,json=allStartNanos,proto3" json:"all_start_nanos,omitempty"`
	OpStartRelNanos      int64                    `protobuf:"varint,14,opt,name=op_start_rel_nanos,json=opStartRelNanos,proto3" json:"op_start_rel_nanos,omitempty"`
	OpEndRelNanos        int64                    `protobuf:"varint,15,opt,name=op_end_rel_nanos,json=opEndRelNanos,proto3" json:"op_end_rel_nanos,omitempty"`
	AllEndRelNanos       int64                    `protobuf:"varint,16,opt,name=all_end_rel_nanos,json=allEndRelNanos,proto3" json:"all_end_rel_nanos,omitempty"`
	ScheduledNanos       int64                    `protobuf:"varint,17,opt,name=scheduled_nanos,json=scheduledNanos,proto3" json:"scheduled_nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *NodeExecStats) Reset()         { *m = NodeExecStats{} }
func (m *NodeExecStats) String() string { return proto.CompactTextString(m) }
func (*NodeExecStats) ProtoMessage()    {}
func (*NodeExecStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e915309f7ed52e5, []int{4}
}

func (m *NodeExecStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecStats.Unmarshal(m, b)
}
func (m *NodeExecStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecStats.Marshal(b, m, deterministic)
}
func (m *NodeExecStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecStats.Merge(m, src)
}
func (m *NodeExecStats) XXX_Size() int {
	return xxx_messageInfo_NodeExecStats.Size(m)
}
func (m *NodeExecStats) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecStats.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecStats proto.InternalMessageInfo

func (m *NodeExecStats) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *NodeExecStats) GetAllStartMicros() int64 {
	if m != nil {
		return m.AllStartMicros
	}
	return 0
}

func (m *NodeExecStats) GetOpStartRelMicros() int64 {
	if m != nil {
		return m.OpStartRelMicros
	}
	return 0
}

func (m *NodeExecStats) GetOpEndRelMicros() int64 {
	if m != nil {
		return m.OpEndRelMicros
	}
	return 0
}

func (m *NodeExecStats) GetAllEndRelMicros() int64 {
	if m != nil {
		return m.AllEndRelMicros
	}
	return 0
}

func (m *NodeExecStats) GetMemory() []*AllocatorMemoryUsed {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *NodeExecStats) GetOutput() []*NodeOutput {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *NodeExecStats) GetTimelineLabel() string {
	if m != nil {
		return m.TimelineLabel
	}
	return ""
}

func (m *NodeExecStats) GetScheduledMicros() int64 {
	if m != nil {
		return m.ScheduledMicros
	}
	return 0
}

func (m *NodeExecStats) GetThreadId() uint32 {
	if m != nil {
		return m.ThreadId
	}
	return 0
}

func (m *NodeExecStats) GetReferencedTensor() []*AllocationDescription {
	if m != nil {
		return m.ReferencedTensor
	}
	return nil
}

func (m *NodeExecStats) GetMemoryStats() *MemoryStats {
	if m != nil {
		return m.MemoryStats
	}
	return nil
}

func (m *NodeExecStats) GetAllStartNanos() int64 {
	if m != nil {
		return m.AllStartNanos
	}
	return 0
}

func (m *NodeExecStats) GetOpStartRelNanos() int64 {
	if m != nil {
		return m.OpStartRelNanos
	}
	return 0
}

func (m *NodeExecStats) GetOpEndRelNanos() int64 {
	if m != nil {
		return m.OpEndRelNanos
	}
	return 0
}

func (m *NodeExecStats) GetAllEndRelNanos() int64 {
	if m != nil {
		return m.AllEndRelNanos
	}
	return 0
}

func (m *NodeExecStats) GetScheduledNanos() int64 {
	if m != nil {
		return m.ScheduledNanos
	}
	return 0
}

type DeviceStepStats struct {
	Device               string           `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	NodeStats            []*NodeExecStats `protobuf:"bytes,2,rep,name=node_stats,json=nodeStats,proto3" json:"node_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DeviceStepStats) Reset()         { *m = DeviceStepStats{} }
func (m *DeviceStepStats) String() string { return proto.CompactTextString(m) }
func (*DeviceStepStats) ProtoMessage()    {}
func (*DeviceStepStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e915309f7ed52e5, []int{5}
}

func (m *DeviceStepStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceStepStats.Unmarshal(m, b)
}
func (m *DeviceStepStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceStepStats.Marshal(b, m, deterministic)
}
func (m *DeviceStepStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceStepStats.Merge(m, src)
}
func (m *DeviceStepStats) XXX_Size() int {
	return xxx_messageInfo_DeviceStepStats.Size(m)
}
func (m *DeviceStepStats) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceStepStats.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceStepStats proto.InternalMessageInfo

func (m *DeviceStepStats) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *DeviceStepStats) GetNodeStats() []*NodeExecStats {
	if m != nil {
		return m.NodeStats
	}
	return nil
}

type StepStats struct {
	DevStats             []*DeviceStepStats `protobuf:"bytes,1,rep,name=dev_stats,json=devStats,proto3" json:"dev_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StepStats) Reset()         { *m = StepStats{} }
func (m *StepStats) String() string { return proto.CompactTextString(m) }
func (*StepStats) ProtoMessage()    {}
func (*StepStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e915309f7ed52e5, []int{6}
}

func (m *StepStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepStats.Unmarshal(m, b)
}
func (m *StepStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepStats.Marshal(b, m, deterministic)
}
func (m *StepStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepStats.Merge(m, src)
}
func (m *StepStats) XXX_Size() int {
	return xxx_messageInfo_StepStats.Size(m)
}
func (m *StepStats) XXX_DiscardUnknown() {
	xxx_messageInfo_StepStats.DiscardUnknown(m)
}

var xxx_messageInfo_StepStats proto.InternalMessageInfo

func (m *StepStats) GetDevStats() []*DeviceStepStats {
	if m != nil {
		return m.DevStats
	}
	return nil
}

func init() {
	proto.RegisterType((*AllocationRecord)(nil), "tensorflow.AllocationRecord")
	proto.RegisterType((*AllocatorMemoryUsed)(nil), "tensorflow.AllocatorMemoryUsed")
	proto.RegisterType((*NodeOutput)(nil), "tensorflow.NodeOutput")
	proto.RegisterType((*MemoryStats)(nil), "tensorflow.MemoryStats")
	proto.RegisterType((*NodeExecStats)(nil), "tensorflow.NodeExecStats")
	proto.RegisterType((*DeviceStepStats)(nil), "tensorflow.DeviceStepStats")
	proto.RegisterType((*StepStats)(nil), "tensorflow.StepStats")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/step_stats.proto", fileDescriptor_1e915309f7ed52e5)
}

var fileDescriptor_1e915309f7ed52e5 = []byte{
	// 891 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xdf, 0x6f, 0xdb, 0x36,
	0x10, 0x86, 0xed, 0xc4, 0x8b, 0xcf, 0x71, 0x6c, 0x33, 0x45, 0xaa, 0x25, 0xcb, 0x9a, 0x1a, 0xd8,
	0xea, 0x6c, 0x98, 0x03, 0xa4, 0xc3, 0xd6, 0x0d, 0xe8, 0xc3, 0x82, 0xe6, 0x21, 0x58, 0xeb, 0x05,
	0x4a, 0xbb, 0x87, 0xbd, 0x08, 0x8a, 0x78, 0x49, 0x84, 0x4a, 0xa2, 0x40, 0xd2, 0xe9, 0xda, 0x7f,
	0x63, 0xcf, 0xfb, 0x3f, 0xf7, 0x34, 0x14, 0x3c, 0xd2, 0x12, 0xfd, 0x23, 0x6f, 0xd4, 0xc7, 0xef,
	0x4e, 0x77, 0xc7, 0xef, 0x23, 0xe1, 0x3b, 0x8d, 0x85, 0x12, 0xf2, 0x26, 0x13, 0x1f, 0x4e, 0x12,
	0x21, 0xf1, 0xe4, 0x46, 0xc6, 0x39, 0x7e, 0x10, 0xf2, 0xfd, 0x89, 0xd2, 0x58, 0x46, 0x4a, 0xc7,
	0x5a, 0x4d, 0x4a, 0x29, 0xb4, 0x60, 0x50, 0x73, 0xf7, 0x7f, 0x7a, 0x38, 0x2e, 0xce, 0x32, 0x91,
	0xc4, 0x3a, 0x15, 0x45, 0xc4, 0x51, 0x25, 0x32, 0x2d, 0xcd, 0xda, 0xe6, 0xd8, 0x3f, 0x7d, 0x38,
	0xce, 0xee, 0xac, 0xc6, 0x8c, 0xfe, 0x84, 0xc1, 0x6f, 0x55, 0xce, 0x10, 0x13, 0x21, 0x39, 0x7b,
	0x0a, 0xdb, 0xf4, 0x9f, 0x28, 0x4f, 0x13, 0x29, 0x54, 0xd0, 0x38, 0x6a, 0x8c, 0x5b, 0x61, 0x97,
	0xb0, 0x37, 0x04, 0xb1, 0x27, 0x60, 0x3f, 0xa3, 0xeb, 0x8f, 0x1a, 0x55, 0xd0, 0x24, 0x06, 0x10,
	0x74, 0x66, 0x90, 0xd1, 0xbf, 0x4d, 0xd8, 0x75, 0x89, 0x85, 0x7c, 0x83, 0xb9, 0x90, 0x1f, 0xdf,
	0x29, 0xe4, 0xec, 0x1b, 0xd8, 0x89, 0xe7, 0x70, 0x54, 0xc4, 0x39, 0x52, 0xf6, 0x4e, 0xd8, 0xab,
	0xd0, 0x69, 0x9c, 0xa3, 0xc9, 0xaf, 0x85, 0x8e, 0xb3, 0xc5, 0xfc, 0x04, 0x51, 0x7e, 0x76, 0x08,
	0x50, 0x62, 0xfc, 0xde, 0xed, 0xb7, 0x68, 0xbf, 0x63, 0x90, 0x6a, 0x3b, 0x4b, 0xef, 0xd1, 0x6d,
	0x6f, 0xd8, 0x6d, 0x83, 0xd8, 0xed, 0xdf, 0x81, 0x79, 0x93, 0x94, 0xd4, 0xb6, 0x0a, 0xda, 0x47,
	0xad, 0x71, 0xf7, 0xf4, 0xab, 0x49, 0x3d, 0xc6, 0xc9, 0xf2, 0x6c, 0xc2, 0x61, 0xbc, 0x84, 0x28,
	0xf6, 0x1c, 0xf6, 0xea, 0x96, 0xe8, 0x87, 0x51, 0x5a, 0x44, 0x33, 0x85, 0xc1, 0x26, 0xfd, 0x77,
	0xb7, 0xda, 0xa5, 0x9f, 0x5f, 0x14, 0xef, 0x14, 0x8e, 0x0a, 0x80, 0xa9, 0xe0, 0xf8, 0xc7, 0x4c,
	0x97, 0x33, 0xcd, 0x18, 0x6c, 0xa8, 0x4c, 0x68, 0x9a, 0xc5, 0x66, 0x48, 0x6b, 0xf6, 0x1a, 0xd8,
	0xea, 0xa9, 0x51, 0xa7, 0xdd, 0xd3, 0x43, 0xbf, 0xc6, 0xb7, 0xb4, 0x7c, 0x55, 0x93, 0xc2, 0xa1,
	0x5e, 0x86, 0x46, 0xff, 0x37, 0xa1, 0x6b, 0x8f, 0xe1, 0xca, 0xa8, 0x8e, 0x8d, 0x61, 0xa0, 0x31,
	0x2f, 0xa3, 0x9c, 0xb0, 0x48, 0xa5, 0x9f, 0xd0, 0x9d, 0xf3, 0x8e, 0xc1, 0x1d, 0x35, 0xfd, 0x84,
	0xec, 0x47, 0xd8, 0x2b, 0x51, 0xaa, 0x54, 0x69, 0x2c, 0xf4, 0x02, 0xdf, 0x4e, 0xfd, 0x51, 0xbd,
	0xeb, 0x45, 0xbd, 0x84, 0x03, 0x2f, 0xca, 0x35, 0x62, 0x25, 0x93, 0x72, 0x15, 0x6c, 0x1e, 0xb5,
	0xc6, 0xad, 0x30, 0xa8, 0x29, 0xb6, 0x09, 0x1a, 0xf7, 0x05, 0x57, 0xec, 0x17, 0x78, 0xcc, 0xf1,
	0x3e, 0x4d, 0x30, 0x5a, 0xa9, 0x92, 0xb4, 0x70, 0xd6, 0x0c, 0x1a, 0xe1, 0x23, 0x4b, 0x79, 0xbb,
	0x58, 0xef, 0x39, 0x1c, 0xba, 0xd0, 0x07, 0xca, 0xde, 0xa8, 0x12, 0xec, 0x5b, 0xe2, 0xe5, 0xba,
	0x06, 0xa6, 0x30, 0x5a, 0x4d, 0xb3, 0xd2, 0x87, 0x91, 0x8c, 0xcd, 0xf5, 0xf5, 0x72, 0xae, 0xc5,
	0x8e, 0x46, 0xff, 0xb4, 0xa1, 0x67, 0x4e, 0xfc, 0xfc, 0x6f, 0x4c, 0xec, 0x11, 0x1c, 0x40, 0xa7,
	0x10, 0x1c, 0x7d, 0x17, 0x6c, 0x19, 0x80, 0x0c, 0x30, 0x86, 0x41, 0x9c, 0x65, 0xe6, 0x8a, 0x90,
	0x7a, 0xee, 0x43, 0xeb, 0x02, 0xe3, 0x9f, 0x2b, 0x03, 0x3b, 0x2b, 0xfe, 0x00, 0xbb, 0xa2, 0x74,
	0x44, 0x89, 0xd9, 0x9c, 0x6c, 0x0f, 0x67, 0x20, 0x4a, 0xe2, 0x86, 0x98, 0x39, 0xfa, 0x31, 0x0c,
	0x45, 0x19, 0x61, 0xc1, 0x7d, 0xb2, 0x35, 0xc8, 0x8e, 0x28, 0xcf, 0x0b, 0x5e, 0x53, 0xbf, 0x27,
	0x97, 0x2c, 0x73, 0xad, 0xa8, 0xfb, 0x71, 0x96, 0x2d, 0x90, 0x7f, 0x86, 0xb6, 0x1d, 0xb2, 0xb3,
	0xd1, 0x93, 0x35, 0x36, 0xf2, 0x6f, 0x82, 0xd0, 0xd1, 0xd9, 0x04, 0xda, 0x82, 0x5c, 0x10, 0x7c,
	0x41, 0x81, 0x7b, 0x7e, 0x60, 0xed, 0x91, 0xd0, 0xb1, 0xcc, 0x0d, 0xa2, 0xd3, 0x1c, 0xb3, 0xb4,
	0xc0, 0x28, 0x8b, 0xaf, 0x31, 0x0b, 0xb6, 0xec, 0x0d, 0x32, 0x47, 0x5f, 0x1b, 0x90, 0x1d, 0xc3,
	0x40, 0x25, 0x77, 0xc8, 0x67, 0x19, 0xf2, 0x79, 0xe9, 0x1d, 0x5b, 0x7a, 0x85, 0xbb, 0xd2, 0x0f,
	0xa0, 0xa3, 0xef, 0x24, 0xc6, 0x3c, 0x4a, 0x79, 0x00, 0x47, 0x8d, 0x71, 0x2f, 0xdc, 0xb2, 0xc0,
	0x05, 0x67, 0x53, 0x18, 0x4a, 0xbc, 0x41, 0x89, 0x45, 0x82, 0xdc, 0x09, 0x20, 0xe8, 0x52, 0xa5,
	0x4f, 0xd7, 0xdf, 0x14, 0xbe, 0x13, 0x07, 0x75, 0xac, 0xd5, 0x03, 0xfb, 0x15, 0xb6, 0xe7, 0x62,
	0x34, 0x2a, 0x08, 0xb6, 0xc9, 0xd0, 0x8f, 0xfd, 0x54, 0x9e, 0x4f, 0xc3, 0x6e, 0xee, 0x99, 0xf6,
	0x5b, 0xe8, 0xd7, 0xa2, 0x28, 0xe2, 0x42, 0xa8, 0xa0, 0x47, 0x2d, 0xf5, 0xe6, 0x9a, 0x98, 0x1a,
	0xd0, 0x1c, 0xdc, 0x82, 0x24, 0x2c, 0x75, 0xc7, 0x76, 0x5f, 0x2b, 0xc2, 0x92, 0x9f, 0xc1, 0xc0,
	0x13, 0x84, 0xa5, 0xf6, 0x6d, 0xd6, 0xb9, 0x1e, 0x2c, 0xf1, 0x18, 0x86, 0xbe, 0x1c, 0x2c, 0x73,
	0x50, 0x69, 0xd2, 0xa7, 0x3e, 0x83, 0x7a, 0xc8, 0x8e, 0x38, 0xb4, 0xc4, 0x0a, 0x26, 0xe2, 0x28,
	0x81, 0xfe, 0x2b, 0xf2, 0xcd, 0x95, 0xc6, 0xd2, 0x36, 0xb9, 0x07, 0x6d, 0x6b, 0x25, 0xe7, 0x09,
	0xf7, 0xc5, 0x5e, 0x00, 0x90, 0x5d, 0xec, 0xd8, 0x9a, 0x74, 0x02, 0x5f, 0x2e, 0x6b, 0xa5, 0x72,
	0x57, 0x48, 0xde, 0xa2, 0xe5, 0xe8, 0x1c, 0x3a, 0x75, 0xfa, 0x17, 0xd0, 0xe1, 0x78, 0xef, 0xb2,
	0x34, 0x28, 0xcb, 0x81, 0x9f, 0x65, 0xa9, 0x9c, 0x70, 0x8b, 0xe3, 0x3d, 0xad, 0xce, 0x04, 0x04,
	0x42, 0xde, 0xfa, 0xdc, 0xea, 0x7d, 0x3d, 0xeb, 0x57, 0x01, 0x97, 0xe6, 0x59, 0x55, 0x97, 0x8d,
	0xbf, 0x5e, 0xde, 0xa6, 0xfa, 0x6e, 0x76, 0x3d, 0x49, 0x44, 0x7e, 0xe2, 0x3d, 0xcc, 0xeb, 0x97,
	0xb7, 0x62, 0xe9, 0xc5, 0xfe, 0xaf, 0xd1, 0xb8, 0x6e, 0xd3, 0x13, 0xfd, 0xfc, 0x73, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb7, 0x8d, 0xf8, 0x8a, 0x48, 0x08, 0x00, 0x00,
}
