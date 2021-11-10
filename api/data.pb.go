// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: data.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Image_OS int32

const (
	Image_OS_UNSPECIFIEED Image_OS = 0
	Image_DEBIAN_BUSTER   Image_OS = 1
	Image_DEBIAN_BULLSEYE Image_OS = 2
	Image_DEBIAN_BOOKWORM Image_OS = 3
)

// Enum value maps for Image_OS.
var (
	Image_OS_name = map[int32]string{
		0: "OS_UNSPECIFIEED",
		1: "DEBIAN_BUSTER",
		2: "DEBIAN_BULLSEYE",
		3: "DEBIAN_BOOKWORM",
	}
	Image_OS_value = map[string]int32{
		"OS_UNSPECIFIEED": 0,
		"DEBIAN_BUSTER":   1,
		"DEBIAN_BULLSEYE": 2,
		"DEBIAN_BOOKWORM": 3,
	}
)

func (x Image_OS) Enum() *Image_OS {
	p := new(Image_OS)
	*p = x
	return p
}

func (x Image_OS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Image_OS) Descriptor() protoreflect.EnumDescriptor {
	return file_data_proto_enumTypes[0].Descriptor()
}

func (Image_OS) Type() protoreflect.EnumType {
	return &file_data_proto_enumTypes[0]
}

func (x Image_OS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Image_OS.Descriptor instead.
func (Image_OS) EnumDescriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{1, 0}
}

type Machine_Status int32

const (
	Machine_STATUS_UNSPECIFIED Machine_Status = 0
	Machine_CREATED            Machine_Status = 1
	Machine_STOPPED            Machine_Status = 2
	Machine_RUNNING            Machine_Status = 3
	Machine_DESTROYED          Machine_Status = 4
)

// Enum value maps for Machine_Status.
var (
	Machine_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATED",
		2: "STOPPED",
		3: "RUNNING",
		4: "DESTROYED",
	}
	Machine_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATED":            1,
		"STOPPED":            2,
		"RUNNING":            3,
		"DESTROYED":          4,
	}
)

func (x Machine_Status) Enum() *Machine_Status {
	p := new(Machine_Status)
	*p = x
	return p
}

func (x Machine_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Machine_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_data_proto_enumTypes[1].Descriptor()
}

func (Machine_Status) Type() protoreflect.EnumType {
	return &file_data_proto_enumTypes[1]
}

func (x Machine_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Machine_Status.Descriptor instead.
func (Machine_Status) EnumDescriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{2, 0}
}

type SSHKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Pubkey string `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
}

func (x *SSHKey) Reset() {
	*x = SSHKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSHKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSHKey) ProtoMessage() {}

func (x *SSHKey) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSHKey.ProtoReflect.Descriptor instead.
func (*SSHKey) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{0}
}

func (x *SSHKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SSHKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SSHKey) GetPubkey() string {
	if x != nil {
		return x.Pubkey
	}
	return ""
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	System Image_OS `protobuf:"varint,3,opt,name=system,proto3,enum=virtm.Image_OS" json:"system,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{1}
}

func (x *Image) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Image) GetSystem() Image_OS {
	if x != nil {
		return x.System
	}
	return Image_OS_UNSPECIFIEED
}

type Machine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status    Machine_Status      `protobuf:"varint,3,opt,name=status,proto3,enum=virtm.Machine_Status" json:"status,omitempty"`
	Specs     *Machine_Specs      `protobuf:"bytes,4,opt,name=specs,proto3" json:"specs,omitempty"`
	Networks  []*NetworkInterface `protobuf:"bytes,5,rep,name=networks,proto3" json:"networks,omitempty"`
	ImageId   string              `protobuf:"bytes,6,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	SshKeyIds []string            `protobuf:"bytes,7,rep,name=ssh_key_ids,json=sshKeyIds,proto3" json:"ssh_key_ids,omitempty"`
}

func (x *Machine) Reset() {
	*x = Machine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Machine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Machine) ProtoMessage() {}

func (x *Machine) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Machine.ProtoReflect.Descriptor instead.
func (*Machine) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{2}
}

func (x *Machine) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Machine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Machine) GetStatus() Machine_Status {
	if x != nil {
		return x.Status
	}
	return Machine_STATUS_UNSPECIFIED
}

func (x *Machine) GetSpecs() *Machine_Specs {
	if x != nil {
		return x.Specs
	}
	return nil
}

func (x *Machine) GetNetworks() []*NetworkInterface {
	if x != nil {
		return x.Networks
	}
	return nil
}

func (x *Machine) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *Machine) GetSshKeyIds() []string {
	if x != nil {
		return x.SshKeyIds
	}
	return nil
}

type NetworkInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	IpV4      string `protobuf:"bytes,2,opt,name=ip_v4,json=ipV4,proto3" json:"ip_v4,omitempty"`
	IpV6      string `protobuf:"bytes,3,opt,name=ip_v6,json=ipV6,proto3" json:"ip_v6,omitempty"`
}

func (x *NetworkInterface) Reset() {
	*x = NetworkInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInterface) ProtoMessage() {}

func (x *NetworkInterface) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInterface.ProtoReflect.Descriptor instead.
func (*NetworkInterface) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{3}
}

func (x *NetworkInterface) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *NetworkInterface) GetIpV4() string {
	if x != nil {
		return x.IpV4
	}
	return ""
}

func (x *NetworkInterface) GetIpV6() string {
	if x != nil {
		return x.IpV6
	}
	return ""
}

type Network struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IpV4          *Network_IpNetwork `protobuf:"bytes,3,opt,name=ip_v4,json=ipV4,proto3" json:"ip_v4,omitempty"`
	IpV6          *Network_IpNetwork `protobuf:"bytes,4,opt,name=ip_v6,json=ipV6,proto3" json:"ip_v6,omitempty"`
	Nameservers   []string           `protobuf:"bytes,5,rep,name=nameservers,proto3" json:"nameservers,omitempty"`
	SearchDomains []string           `protobuf:"bytes,6,rep,name=searchDomains,proto3" json:"searchDomains,omitempty"`
}

func (x *Network) Reset() {
	*x = Network{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Network) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Network) ProtoMessage() {}

func (x *Network) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Network.ProtoReflect.Descriptor instead.
func (*Network) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{4}
}

func (x *Network) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Network) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Network) GetIpV4() *Network_IpNetwork {
	if x != nil {
		return x.IpV4
	}
	return nil
}

func (x *Network) GetIpV6() *Network_IpNetwork {
	if x != nil {
		return x.IpV6
	}
	return nil
}

func (x *Network) GetNameservers() []string {
	if x != nil {
		return x.Nameservers
	}
	return nil
}

func (x *Network) GetSearchDomains() []string {
	if x != nil {
		return x.SearchDomains
	}
	return nil
}

type Machine_Specs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpus   int64 `protobuf:"varint,1,opt,name=cpus,proto3" json:"cpus,omitempty"`
	Memory int64 `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
	Disk   int64 `protobuf:"varint,3,opt,name=disk,proto3" json:"disk,omitempty"`
}

func (x *Machine_Specs) Reset() {
	*x = Machine_Specs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Machine_Specs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Machine_Specs) ProtoMessage() {}

func (x *Machine_Specs) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Machine_Specs.ProtoReflect.Descriptor instead.
func (*Machine_Specs) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Machine_Specs) GetCpus() int64 {
	if x != nil {
		return x.Cpus
	}
	return 0
}

func (x *Machine_Specs) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *Machine_Specs) GetDisk() int64 {
	if x != nil {
		return x.Disk
	}
	return 0
}

type Network_IpNetwork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subnet  string `protobuf:"bytes,1,opt,name=subnet,proto3" json:"subnet,omitempty"`
	Gateway string `protobuf:"bytes,2,opt,name=gateway,proto3" json:"gateway,omitempty"`
}

func (x *Network_IpNetwork) Reset() {
	*x = Network_IpNetwork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Network_IpNetwork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Network_IpNetwork) ProtoMessage() {}

func (x *Network_IpNetwork) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Network_IpNetwork.ProtoReflect.Descriptor instead.
func (*Network_IpNetwork) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Network_IpNetwork) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

func (x *Network_IpNetwork) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

var File_data_proto protoreflect.FileDescriptor

var file_data_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x76, 0x69,
	0x72, 0x74, 0x6d, 0x22, 0x44, 0x0a, 0x06, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x22, 0xac, 0x01, 0x0a, 0x05, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x6d, 0x2e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x53, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x22, 0x56, 0x0a, 0x02, 0x4f, 0x53, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x44,
	0x45, 0x42, 0x49, 0x41, 0x4e, 0x5f, 0x42, 0x55, 0x53, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x44, 0x45, 0x42, 0x49, 0x41, 0x4e, 0x5f, 0x42, 0x55, 0x4c, 0x4c, 0x53, 0x45, 0x59,
	0x45, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x45, 0x42, 0x49, 0x41, 0x4e, 0x5f, 0x42, 0x4f,
	0x4f, 0x4b, 0x57, 0x4f, 0x52, 0x4d, 0x10, 0x03, 0x22, 0x99, 0x03, 0x0a, 0x07, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x6d,
	0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x63, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x6d, 0x2e, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x73, 0x52, 0x05, 0x73, 0x70,
	0x65, 0x63, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x6d, 0x2e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x08,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x73, 0x73, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x73, 0x68, 0x4b, 0x65, 0x79,
	0x49, 0x64, 0x73, 0x1a, 0x47, 0x0a, 0x05, 0x53, 0x70, 0x65, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x70, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x70, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x22, 0x56, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e,
	0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x53, 0x54, 0x52, 0x4f, 0x59,
	0x45, 0x44, 0x10, 0x04, 0x22, 0x5b, 0x0a, 0x10, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x70, 0x5f, 0x76, 0x34,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x56, 0x34, 0x12, 0x13, 0x0a, 0x05,
	0x69, 0x70, 0x5f, 0x76, 0x36, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x56,
	0x36, 0x22, 0x92, 0x02, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x70, 0x5f, 0x76, 0x34, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x6d, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x49, 0x70, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x04, 0x69, 0x70, 0x56, 0x34,
	0x12, 0x2d, 0x0a, 0x05, 0x69, 0x70, 0x5f, 0x76, 0x36, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x6d, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x49, 0x70, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x04, 0x69, 0x70, 0x56, 0x36, 0x12,
	0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x1a, 0x3d, 0x0a, 0x09, 0x49, 0x70, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x61, 0x6c, 0x61, 0x72, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x6d,
	0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_proto_rawDescOnce sync.Once
	file_data_proto_rawDescData = file_data_proto_rawDesc
)

func file_data_proto_rawDescGZIP() []byte {
	file_data_proto_rawDescOnce.Do(func() {
		file_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_proto_rawDescData)
	})
	return file_data_proto_rawDescData
}

var file_data_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_data_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_data_proto_goTypes = []interface{}{
	(Image_OS)(0),             // 0: virtm.Image.OS
	(Machine_Status)(0),       // 1: virtm.Machine.Status
	(*SSHKey)(nil),            // 2: virtm.SSHKey
	(*Image)(nil),             // 3: virtm.Image
	(*Machine)(nil),           // 4: virtm.Machine
	(*NetworkInterface)(nil),  // 5: virtm.NetworkInterface
	(*Network)(nil),           // 6: virtm.Network
	(*Machine_Specs)(nil),     // 7: virtm.Machine.Specs
	(*Network_IpNetwork)(nil), // 8: virtm.Network.IpNetwork
}
var file_data_proto_depIdxs = []int32{
	0, // 0: virtm.Image.system:type_name -> virtm.Image.OS
	1, // 1: virtm.Machine.status:type_name -> virtm.Machine.Status
	7, // 2: virtm.Machine.specs:type_name -> virtm.Machine.Specs
	5, // 3: virtm.Machine.networks:type_name -> virtm.NetworkInterface
	8, // 4: virtm.Network.ip_v4:type_name -> virtm.Network.IpNetwork
	8, // 5: virtm.Network.ip_v6:type_name -> virtm.Network.IpNetwork
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_data_proto_init() }
func file_data_proto_init() {
	if File_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSHKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Machine); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInterface); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Network); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Machine_Specs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Network_IpNetwork); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_data_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_proto_goTypes,
		DependencyIndexes: file_data_proto_depIdxs,
		EnumInfos:         file_data_proto_enumTypes,
		MessageInfos:      file_data_proto_msgTypes,
	}.Build()
	File_data_proto = out.File
	file_data_proto_rawDesc = nil
	file_data_proto_goTypes = nil
	file_data_proto_depIdxs = nil
}
