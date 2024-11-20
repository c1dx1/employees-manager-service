// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: proto/employee.proto

package proto

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

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname    string               `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	Phone      string               `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	CompanyId  int32                `protobuf:"varint,5,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Passport   *Employee_Passport   `protobuf:"bytes,6,opt,name=passport,proto3" json:"passport,omitempty"`
	Department *Employee_Department `protobuf:"bytes,7,opt,name=department,proto3" json:"department,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	mi := &file_proto_employee_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{0}
}

func (x *Employee) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Employee) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Employee) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *Employee) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Employee) GetCompanyId() int32 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *Employee) GetPassport() *Employee_Passport {
	if x != nil {
		return x.Passport
	}
	return nil
}

func (x *Employee) GetDepartment() *Employee_Department {
	if x != nil {
		return x.Department
	}
	return nil
}

type AddEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname    string               `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	Phone      string               `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	CompanyId  int32                `protobuf:"varint,4,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Passport   *Employee_Passport   `protobuf:"bytes,5,opt,name=passport,proto3" json:"passport,omitempty"`
	Department *Employee_Department `protobuf:"bytes,6,opt,name=department,proto3" json:"department,omitempty"`
}

func (x *AddEmployeeRequest) Reset() {
	*x = AddEmployeeRequest{}
	mi := &file_proto_employee_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEmployeeRequest) ProtoMessage() {}

func (x *AddEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEmployeeRequest.ProtoReflect.Descriptor instead.
func (*AddEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{1}
}

func (x *AddEmployeeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddEmployeeRequest) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *AddEmployeeRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AddEmployeeRequest) GetCompanyId() int32 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *AddEmployeeRequest) GetPassport() *Employee_Passport {
	if x != nil {
		return x.Passport
	}
	return nil
}

func (x *AddEmployeeRequest) GetDepartment() *Employee_Department {
	if x != nil {
		return x.Department
	}
	return nil
}

type AddEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddEmployeeResponse) Reset() {
	*x = AddEmployeeResponse{}
	mi := &file_proto_employee_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEmployeeResponse) ProtoMessage() {}

func (x *AddEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEmployeeResponse.ProtoReflect.Descriptor instead.
func (*AddEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{2}
}

func (x *AddEmployeeResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteEmployeeRequest) Reset() {
	*x = DeleteEmployeeRequest{}
	mi := &file_proto_employee_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmployeeRequest) ProtoMessage() {}

func (x *DeleteEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmployeeRequest.ProtoReflect.Descriptor instead.
func (*DeleteEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteEmployeeRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success string `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteEmployeeResponse) Reset() {
	*x = DeleteEmployeeResponse{}
	mi := &file_proto_employee_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmployeeResponse) ProtoMessage() {}

func (x *DeleteEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmployeeResponse.ProtoReflect.Descriptor instead.
func (*DeleteEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteEmployeeResponse) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

type CompanyEmployeesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId  int32                `protobuf:"varint,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Department *Employee_Department `protobuf:"bytes,2,opt,name=department,proto3" json:"department,omitempty"`
}

func (x *CompanyEmployeesRequest) Reset() {
	*x = CompanyEmployeesRequest{}
	mi := &file_proto_employee_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompanyEmployeesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyEmployeesRequest) ProtoMessage() {}

func (x *CompanyEmployeesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyEmployeesRequest.ProtoReflect.Descriptor instead.
func (*CompanyEmployeesRequest) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{5}
}

func (x *CompanyEmployeesRequest) GetCompanyId() int32 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *CompanyEmployeesRequest) GetDepartment() *Employee_Department {
	if x != nil {
		return x.Department
	}
	return nil
}

type EmployeesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employees []*Employee `protobuf:"bytes,1,rep,name=employees,proto3" json:"employees,omitempty"`
}

func (x *EmployeesResponse) Reset() {
	*x = EmployeesResponse{}
	mi := &file_proto_employee_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmployeesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeesResponse) ProtoMessage() {}

func (x *EmployeesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeesResponse.ProtoReflect.Descriptor instead.
func (*EmployeesResponse) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{6}
}

func (x *EmployeesResponse) GetEmployees() []*Employee {
	if x != nil {
		return x.Employees
	}
	return nil
}

type UpdateEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname    string               `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	Phone      string               `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	CompanyId  int32                `protobuf:"varint,5,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Passport   *Employee_Passport   `protobuf:"bytes,6,opt,name=passport,proto3" json:"passport,omitempty"`
	Department *Employee_Department `protobuf:"bytes,7,opt,name=department,proto3" json:"department,omitempty"`
}

func (x *UpdateEmployeeRequest) Reset() {
	*x = UpdateEmployeeRequest{}
	mi := &file_proto_employee_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmployeeRequest) ProtoMessage() {}

func (x *UpdateEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmployeeRequest.ProtoReflect.Descriptor instead.
func (*UpdateEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateEmployeeRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateEmployeeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetCompanyId() int32 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *UpdateEmployeeRequest) GetPassport() *Employee_Passport {
	if x != nil {
		return x.Passport
	}
	return nil
}

func (x *UpdateEmployeeRequest) GetDepartment() *Employee_Department {
	if x != nil {
		return x.Department
	}
	return nil
}

type UpdateEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success string `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UpdateEmployeeResponse) Reset() {
	*x = UpdateEmployeeResponse{}
	mi := &file_proto_employee_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmployeeResponse) ProtoMessage() {}

func (x *UpdateEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmployeeResponse.ProtoReflect.Descriptor instead.
func (*UpdateEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateEmployeeResponse) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

type Employee_Passport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Number string `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Employee_Passport) Reset() {
	*x = Employee_Passport{}
	mi := &file_proto_employee_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Employee_Passport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee_Passport) ProtoMessage() {}

func (x *Employee_Passport) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee_Passport.ProtoReflect.Descriptor instead.
func (*Employee_Passport) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Employee_Passport) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Employee_Passport) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type Employee_Department struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *Employee_Department) Reset() {
	*x = Employee_Department{}
	mi := &file_proto_employee_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Employee_Department) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee_Department) ProtoMessage() {}

func (x *Employee_Department) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee_Department.ProtoReflect.Descriptor instead.
func (*Employee_Department) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Employee_Department) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Employee_Department) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_proto_employee_proto protoreflect.FileDescriptor

var file_proto_employee_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x02,
	0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x34, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x2e, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x1a,
	0x36, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x36, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22,
	0xe9, 0x01, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x50, 0x61, 0x73,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x3a, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x13, 0x41,
	0x64, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x16, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x74, 0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x44,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x42, 0x0a, 0x11, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x09,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x22, 0xfc, 0x01, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3a, 0x0a, 0x0a,
	0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xcf, 0x02, 0x0a,
	0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x46, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x14, 0x53, 0x68, 0x6f,
	0x77, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e,
	0x5a, 0x0c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_employee_proto_rawDescOnce sync.Once
	file_proto_employee_proto_rawDescData = file_proto_employee_proto_rawDesc
)

func file_proto_employee_proto_rawDescGZIP() []byte {
	file_proto_employee_proto_rawDescOnce.Do(func() {
		file_proto_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_employee_proto_rawDescData)
	})
	return file_proto_employee_proto_rawDescData
}

var file_proto_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_employee_proto_goTypes = []any{
	(*Employee)(nil),                // 0: proto.Employee
	(*AddEmployeeRequest)(nil),      // 1: proto.AddEmployeeRequest
	(*AddEmployeeResponse)(nil),     // 2: proto.AddEmployeeResponse
	(*DeleteEmployeeRequest)(nil),   // 3: proto.DeleteEmployeeRequest
	(*DeleteEmployeeResponse)(nil),  // 4: proto.DeleteEmployeeResponse
	(*CompanyEmployeesRequest)(nil), // 5: proto.CompanyEmployeesRequest
	(*EmployeesResponse)(nil),       // 6: proto.EmployeesResponse
	(*UpdateEmployeeRequest)(nil),   // 7: proto.UpdateEmployeeRequest
	(*UpdateEmployeeResponse)(nil),  // 8: proto.UpdateEmployeeResponse
	(*Employee_Passport)(nil),       // 9: proto.Employee.Passport
	(*Employee_Department)(nil),     // 10: proto.Employee.Department
}
var file_proto_employee_proto_depIdxs = []int32{
	9,  // 0: proto.Employee.passport:type_name -> proto.Employee.Passport
	10, // 1: proto.Employee.department:type_name -> proto.Employee.Department
	9,  // 2: proto.AddEmployeeRequest.passport:type_name -> proto.Employee.Passport
	10, // 3: proto.AddEmployeeRequest.department:type_name -> proto.Employee.Department
	10, // 4: proto.CompanyEmployeesRequest.department:type_name -> proto.Employee.Department
	0,  // 5: proto.EmployeesResponse.employees:type_name -> proto.Employee
	9,  // 6: proto.UpdateEmployeeRequest.passport:type_name -> proto.Employee.Passport
	10, // 7: proto.UpdateEmployeeRequest.department:type_name -> proto.Employee.Department
	1,  // 8: proto.EmployeeService.AddEmployee:input_type -> proto.AddEmployeeRequest
	3,  // 9: proto.EmployeeService.DeleteEmployee:input_type -> proto.DeleteEmployeeRequest
	5,  // 10: proto.EmployeeService.ShowCompanyEmployees:input_type -> proto.CompanyEmployeesRequest
	7,  // 11: proto.EmployeeService.UpdateEmployee:input_type -> proto.UpdateEmployeeRequest
	2,  // 12: proto.EmployeeService.AddEmployee:output_type -> proto.AddEmployeeResponse
	4,  // 13: proto.EmployeeService.DeleteEmployee:output_type -> proto.DeleteEmployeeResponse
	6,  // 14: proto.EmployeeService.ShowCompanyEmployees:output_type -> proto.EmployeesResponse
	8,  // 15: proto.EmployeeService.UpdateEmployee:output_type -> proto.UpdateEmployeeResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_employee_proto_init() }
func file_proto_employee_proto_init() {
	if File_proto_employee_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_employee_proto_goTypes,
		DependencyIndexes: file_proto_employee_proto_depIdxs,
		MessageInfos:      file_proto_employee_proto_msgTypes,
	}.Build()
	File_proto_employee_proto = out.File
	file_proto_employee_proto_rawDesc = nil
	file_proto_employee_proto_goTypes = nil
	file_proto_employee_proto_depIdxs = nil
}