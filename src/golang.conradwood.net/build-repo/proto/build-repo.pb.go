// Code generated by protoc-gen-go.
// source: proto/build-repo.proto
// DO NOT EDIT!

/*
Package buildrepo is a generated protocol buffer package.

It is generated from these files:
	proto/build-repo.proto

It has these top-level messages:
	CreateBuildRequest
	CreateBuildResponse
	UploadSlotRequest
	UploadSlotResponse
	UploadDoneRequest
	UploadDoneResponse
	RepoEntry
	ListReposRequest
	ListReposResponse
	ListBranchesRequest
	ListBranchesResponse
	ListVersionsRequest
	ListVersionsResponse
	ListFilesRequest
	ListFilesResponse
	GetLatestVersionRequest
	GetLatestVersionResponse
	GetBlockRequest
	GetBlockResponse
	File
	GetMetaRequest
	GetMetaResponse
*/
package buildrepo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//
// import "google/protobuf/empty.proto";
// import "google/protobuf/duration.proto";
// import "examples/sub/message.proto";
// import "examples/sub2/message.proto";
// import "google/protobuf/timestamp.proto";
type CreateBuildRequest struct {
	Repository string `protobuf:"bytes,1,opt,name=Repository,json=repository" json:"Repository,omitempty"`
	CommitID   string `protobuf:"bytes,2,opt,name=CommitID,json=commitID" json:"CommitID,omitempty"`
	Branch     string `protobuf:"bytes,3,opt,name=Branch,json=branch" json:"Branch,omitempty"`
	BuildID    uint64 `protobuf:"varint,4,opt,name=BuildID,json=buildID" json:"BuildID,omitempty"`
	CommitMSG  string `protobuf:"bytes,5,opt,name=CommitMSG,json=commitMSG" json:"CommitMSG,omitempty"`
}

func (m *CreateBuildRequest) Reset()                    { *m = CreateBuildRequest{} }
func (m *CreateBuildRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateBuildRequest) ProtoMessage()               {}
func (*CreateBuildRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateBuildRequest) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *CreateBuildRequest) GetCommitID() string {
	if m != nil {
		return m.CommitID
	}
	return ""
}

func (m *CreateBuildRequest) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *CreateBuildRequest) GetBuildID() uint64 {
	if m != nil {
		return m.BuildID
	}
	return 0
}

func (m *CreateBuildRequest) GetCommitMSG() string {
	if m != nil {
		return m.CommitMSG
	}
	return ""
}

type CreateBuildResponse struct {
	BuildStoreid string `protobuf:"bytes,1,opt,name=BuildStoreid,json=buildStoreid" json:"BuildStoreid,omitempty"`
}

func (m *CreateBuildResponse) Reset()                    { *m = CreateBuildResponse{} }
func (m *CreateBuildResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateBuildResponse) ProtoMessage()               {}
func (*CreateBuildResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateBuildResponse) GetBuildStoreid() string {
	if m != nil {
		return m.BuildStoreid
	}
	return ""
}

type UploadSlotRequest struct {
	BuildStoreid string `protobuf:"bytes,1,opt,name=BuildStoreid,json=buildStoreid" json:"BuildStoreid,omitempty"`
	Filename     string `protobuf:"bytes,2,opt,name=Filename,json=filename" json:"Filename,omitempty"`
}

func (m *UploadSlotRequest) Reset()                    { *m = UploadSlotRequest{} }
func (m *UploadSlotRequest) String() string            { return proto.CompactTextString(m) }
func (*UploadSlotRequest) ProtoMessage()               {}
func (*UploadSlotRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UploadSlotRequest) GetBuildStoreid() string {
	if m != nil {
		return m.BuildStoreid
	}
	return ""
}

func (m *UploadSlotRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type UploadSlotResponse struct {
	Token string `protobuf:"bytes,1,opt,name=Token,json=token" json:"Token,omitempty"`
	Port  int32  `protobuf:"varint,2,opt,name=Port,json=port" json:"Port,omitempty"`
}

func (m *UploadSlotResponse) Reset()                    { *m = UploadSlotResponse{} }
func (m *UploadSlotResponse) String() string            { return proto.CompactTextString(m) }
func (*UploadSlotResponse) ProtoMessage()               {}
func (*UploadSlotResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UploadSlotResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UploadSlotResponse) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type UploadDoneRequest struct {
	BuildStoreid string `protobuf:"bytes,1,opt,name=BuildStoreid,json=buildStoreid" json:"BuildStoreid,omitempty"`
}

func (m *UploadDoneRequest) Reset()                    { *m = UploadDoneRequest{} }
func (m *UploadDoneRequest) String() string            { return proto.CompactTextString(m) }
func (*UploadDoneRequest) ProtoMessage()               {}
func (*UploadDoneRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UploadDoneRequest) GetBuildStoreid() string {
	if m != nil {
		return m.BuildStoreid
	}
	return ""
}

type UploadDoneResponse struct {
}

func (m *UploadDoneResponse) Reset()                    { *m = UploadDoneResponse{} }
func (m *UploadDoneResponse) String() string            { return proto.CompactTextString(m) }
func (*UploadDoneResponse) ProtoMessage()               {}
func (*UploadDoneResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type RepoEntry struct {
	Name string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
	Type int32  `protobuf:"varint,2,opt,name=Type,json=type" json:"Type,omitempty"`
}

func (m *RepoEntry) Reset()                    { *m = RepoEntry{} }
func (m *RepoEntry) String() string            { return proto.CompactTextString(m) }
func (*RepoEntry) ProtoMessage()               {}
func (*RepoEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RepoEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RepoEntry) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ListReposRequest struct {
}

func (m *ListReposRequest) Reset()                    { *m = ListReposRequest{} }
func (m *ListReposRequest) String() string            { return proto.CompactTextString(m) }
func (*ListReposRequest) ProtoMessage()               {}
func (*ListReposRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type ListReposResponse struct {
	Entries []*RepoEntry `protobuf:"bytes,1,rep,name=Entries,json=entries" json:"Entries,omitempty"`
}

func (m *ListReposResponse) Reset()                    { *m = ListReposResponse{} }
func (m *ListReposResponse) String() string            { return proto.CompactTextString(m) }
func (*ListReposResponse) ProtoMessage()               {}
func (*ListReposResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListReposResponse) GetEntries() []*RepoEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type ListBranchesRequest struct {
	Repository string `protobuf:"bytes,1,opt,name=Repository,json=repository" json:"Repository,omitempty"`
}

func (m *ListBranchesRequest) Reset()                    { *m = ListBranchesRequest{} }
func (m *ListBranchesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListBranchesRequest) ProtoMessage()               {}
func (*ListBranchesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListBranchesRequest) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

type ListBranchesResponse struct {
	Entries []*RepoEntry `protobuf:"bytes,1,rep,name=Entries,json=entries" json:"Entries,omitempty"`
}

func (m *ListBranchesResponse) Reset()                    { *m = ListBranchesResponse{} }
func (m *ListBranchesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListBranchesResponse) ProtoMessage()               {}
func (*ListBranchesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ListBranchesResponse) GetEntries() []*RepoEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type ListVersionsRequest struct {
	Repository string `protobuf:"bytes,1,opt,name=Repository,json=repository" json:"Repository,omitempty"`
	Branch     string `protobuf:"bytes,2,opt,name=Branch,json=branch" json:"Branch,omitempty"`
}

func (m *ListVersionsRequest) Reset()                    { *m = ListVersionsRequest{} }
func (m *ListVersionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListVersionsRequest) ProtoMessage()               {}
func (*ListVersionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ListVersionsRequest) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *ListVersionsRequest) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

type ListVersionsResponse struct {
	Entries []*RepoEntry `protobuf:"bytes,1,rep,name=Entries,json=entries" json:"Entries,omitempty"`
}

func (m *ListVersionsResponse) Reset()                    { *m = ListVersionsResponse{} }
func (m *ListVersionsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListVersionsResponse) ProtoMessage()               {}
func (*ListVersionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ListVersionsResponse) GetEntries() []*RepoEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type ListFilesRequest struct {
	Repository string `protobuf:"bytes,1,opt,name=Repository,json=repository" json:"Repository,omitempty"`
	Branch     string `protobuf:"bytes,2,opt,name=Branch,json=branch" json:"Branch,omitempty"`
	Version    string `protobuf:"bytes,3,opt,name=Version,json=version" json:"Version,omitempty"`
}

func (m *ListFilesRequest) Reset()                    { *m = ListFilesRequest{} }
func (m *ListFilesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListFilesRequest) ProtoMessage()               {}
func (*ListFilesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ListFilesRequest) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *ListFilesRequest) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *ListFilesRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type ListFilesResponse struct {
	Entries []*RepoEntry `protobuf:"bytes,1,rep,name=Entries,json=entries" json:"Entries,omitempty"`
}

func (m *ListFilesResponse) Reset()                    { *m = ListFilesResponse{} }
func (m *ListFilesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListFilesResponse) ProtoMessage()               {}
func (*ListFilesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ListFilesResponse) GetEntries() []*RepoEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type GetLatestVersionRequest struct {
	Repository string `protobuf:"bytes,1,opt,name=Repository,json=repository" json:"Repository,omitempty"`
	Branch     string `protobuf:"bytes,2,opt,name=Branch,json=branch" json:"Branch,omitempty"`
}

func (m *GetLatestVersionRequest) Reset()                    { *m = GetLatestVersionRequest{} }
func (m *GetLatestVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLatestVersionRequest) ProtoMessage()               {}
func (*GetLatestVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *GetLatestVersionRequest) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *GetLatestVersionRequest) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

type GetLatestVersionResponse struct {
	BuildID uint64 `protobuf:"varint,1,opt,name=BuildID,json=buildID" json:"BuildID,omitempty"`
}

func (m *GetLatestVersionResponse) Reset()                    { *m = GetLatestVersionResponse{} }
func (m *GetLatestVersionResponse) String() string            { return proto.CompactTextString(m) }
func (*GetLatestVersionResponse) ProtoMessage()               {}
func (*GetLatestVersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *GetLatestVersionResponse) GetBuildID() uint64 {
	if m != nil {
		return m.BuildID
	}
	return 0
}

type GetBlockRequest struct {
	File   *File  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=Offset,json=offset" json:"Offset,omitempty"`
	Size   uint32 `protobuf:"varint,3,opt,name=Size,json=size" json:"Size,omitempty"`
}

func (m *GetBlockRequest) Reset()                    { *m = GetBlockRequest{} }
func (m *GetBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBlockRequest) ProtoMessage()               {}
func (*GetBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *GetBlockRequest) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *GetBlockRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetBlockRequest) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type GetBlockResponse struct {
	File   *File  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=Offset,json=offset" json:"Offset,omitempty"`
	Size   uint32 `protobuf:"varint,3,opt,name=Size,json=size" json:"Size,omitempty"`
	Data   []byte `protobuf:"bytes,4,opt,name=Data,json=data,proto3" json:"Data,omitempty"`
}

func (m *GetBlockResponse) Reset()                    { *m = GetBlockResponse{} }
func (m *GetBlockResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBlockResponse) ProtoMessage()               {}
func (*GetBlockResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *GetBlockResponse) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *GetBlockResponse) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetBlockResponse) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GetBlockResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// identify a unique file
type File struct {
	Repository string `protobuf:"bytes,1,opt,name=Repository,json=repository" json:"Repository,omitempty"`
	Branch     string `protobuf:"bytes,2,opt,name=Branch,json=branch" json:"Branch,omitempty"`
	BuildID    uint64 `protobuf:"varint,3,opt,name=BuildID,json=buildID" json:"BuildID,omitempty"`
	Filename   string `protobuf:"bytes,4,opt,name=Filename,json=filename" json:"Filename,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *File) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *File) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *File) GetBuildID() uint64 {
	if m != nil {
		return m.BuildID
	}
	return 0
}

func (m *File) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type GetMetaRequest struct {
	File *File `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
}

func (m *GetMetaRequest) Reset()                    { *m = GetMetaRequest{} }
func (m *GetMetaRequest) String() string            { return proto.CompactTextString(m) }
func (*GetMetaRequest) ProtoMessage()               {}
func (*GetMetaRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *GetMetaRequest) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

type GetMetaResponse struct {
	File *File  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Size uint64 `protobuf:"varint,2,opt,name=Size,json=size" json:"Size,omitempty"`
}

func (m *GetMetaResponse) Reset()                    { *m = GetMetaResponse{} }
func (m *GetMetaResponse) String() string            { return proto.CompactTextString(m) }
func (*GetMetaResponse) ProtoMessage()               {}
func (*GetMetaResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *GetMetaResponse) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *GetMetaResponse) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateBuildRequest)(nil), "buildrepo.CreateBuildRequest")
	proto.RegisterType((*CreateBuildResponse)(nil), "buildrepo.CreateBuildResponse")
	proto.RegisterType((*UploadSlotRequest)(nil), "buildrepo.UploadSlotRequest")
	proto.RegisterType((*UploadSlotResponse)(nil), "buildrepo.UploadSlotResponse")
	proto.RegisterType((*UploadDoneRequest)(nil), "buildrepo.UploadDoneRequest")
	proto.RegisterType((*UploadDoneResponse)(nil), "buildrepo.UploadDoneResponse")
	proto.RegisterType((*RepoEntry)(nil), "buildrepo.RepoEntry")
	proto.RegisterType((*ListReposRequest)(nil), "buildrepo.ListReposRequest")
	proto.RegisterType((*ListReposResponse)(nil), "buildrepo.ListReposResponse")
	proto.RegisterType((*ListBranchesRequest)(nil), "buildrepo.ListBranchesRequest")
	proto.RegisterType((*ListBranchesResponse)(nil), "buildrepo.ListBranchesResponse")
	proto.RegisterType((*ListVersionsRequest)(nil), "buildrepo.ListVersionsRequest")
	proto.RegisterType((*ListVersionsResponse)(nil), "buildrepo.ListVersionsResponse")
	proto.RegisterType((*ListFilesRequest)(nil), "buildrepo.ListFilesRequest")
	proto.RegisterType((*ListFilesResponse)(nil), "buildrepo.ListFilesResponse")
	proto.RegisterType((*GetLatestVersionRequest)(nil), "buildrepo.GetLatestVersionRequest")
	proto.RegisterType((*GetLatestVersionResponse)(nil), "buildrepo.GetLatestVersionResponse")
	proto.RegisterType((*GetBlockRequest)(nil), "buildrepo.GetBlockRequest")
	proto.RegisterType((*GetBlockResponse)(nil), "buildrepo.GetBlockResponse")
	proto.RegisterType((*File)(nil), "buildrepo.File")
	proto.RegisterType((*GetMetaRequest)(nil), "buildrepo.GetMetaRequest")
	proto.RegisterType((*GetMetaResponse)(nil), "buildrepo.GetMetaResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BuildRepoManager service

type BuildRepoManagerClient interface {
	CreateBuild(ctx context.Context, in *CreateBuildRequest, opts ...grpc.CallOption) (*CreateBuildResponse, error)
	GetUploadSlot(ctx context.Context, in *UploadSlotRequest, opts ...grpc.CallOption) (*UploadSlotResponse, error)
	UploadsComplete(ctx context.Context, in *UploadDoneRequest, opts ...grpc.CallOption) (*UploadDoneResponse, error)
	ListRepos(ctx context.Context, in *ListReposRequest, opts ...grpc.CallOption) (*ListReposResponse, error)
	ListBranches(ctx context.Context, in *ListBranchesRequest, opts ...grpc.CallOption) (*ListBranchesResponse, error)
	ListVersions(ctx context.Context, in *ListVersionsRequest, opts ...grpc.CallOption) (*ListVersionsResponse, error)
	ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error)
	GetLatestVersion(ctx context.Context, in *GetLatestVersionRequest, opts ...grpc.CallOption) (*GetLatestVersionResponse, error)
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error)
	GetFileMetaData(ctx context.Context, in *GetMetaRequest, opts ...grpc.CallOption) (*GetMetaResponse, error)
}

type buildRepoManagerClient struct {
	cc *grpc.ClientConn
}

func NewBuildRepoManagerClient(cc *grpc.ClientConn) BuildRepoManagerClient {
	return &buildRepoManagerClient{cc}
}

func (c *buildRepoManagerClient) CreateBuild(ctx context.Context, in *CreateBuildRequest, opts ...grpc.CallOption) (*CreateBuildResponse, error) {
	out := new(CreateBuildResponse)
	err := grpc.Invoke(ctx, "/buildrepo.BuildRepoManager/CreateBuild", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildRepoManagerClient) GetUploadSlot(ctx context.Context, in *UploadSlotRequest, opts ...grpc.CallOption) (*UploadSlotResponse, error) {
	out := new(UploadSlotResponse)
	err := grpc.Invoke(ctx, "/buildrepo.BuildRepoManager/GetUploadSlot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildRepoManagerClient) UploadsComplete(ctx context.Context, in *UploadDoneRequest, opts ...grpc.CallOption) (*UploadDoneResponse, error) {
	out := new(UploadDoneResponse)
	err := grpc.Invoke(ctx, "/buildrepo.BuildRepoManager/UploadsComplete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildRepoManagerClient) ListRepos(ctx context.Context, in *ListReposRequest, opts ...grpc.CallOption) (*ListReposResponse, error) {
	out := new(ListReposResponse)
	err := grpc.Invoke(ctx, "/buildrepo.BuildRepoManager/ListRepos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildRepoManagerClient) ListBranches(ctx context.Context, in *ListBranchesRequest, opts ...grpc.CallOption) (*ListBranchesResponse, error) {
	out := new(ListBranchesResponse)
	err := grpc.Invoke(ctx, "/buildrepo.BuildRepoManager/ListBranches", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildRepoManagerClient) ListVersions(ctx context.Context, in *ListVersionsRequest, opts ...grpc.CallOption) (*ListVersionsResponse, error) {
	out := new(ListVersionsResponse)
	err := grpc.Invoke(ctx, "/buildrepo.BuildRepoManager/ListVersions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildRepoManagerClient) ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error) {
	out := new(ListFilesResponse)
	err := grpc.Invoke(ctx, "/buildrepo.BuildRepoManager/ListFiles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildRepoManagerClient) GetLatestVersion(ctx context.Context, in *GetLatestVersionRequest, opts ...grpc.CallOption) (*GetLatestVersionResponse, error) {
	out := new(GetLatestVersionResponse)
	err := grpc.Invoke(ctx, "/buildrepo.BuildRepoManager/GetLatestVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildRepoManagerClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error) {
	out := new(GetBlockResponse)
	err := grpc.Invoke(ctx, "/buildrepo.BuildRepoManager/GetBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildRepoManagerClient) GetFileMetaData(ctx context.Context, in *GetMetaRequest, opts ...grpc.CallOption) (*GetMetaResponse, error) {
	out := new(GetMetaResponse)
	err := grpc.Invoke(ctx, "/buildrepo.BuildRepoManager/GetFileMetaData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BuildRepoManager service

type BuildRepoManagerServer interface {
	CreateBuild(context.Context, *CreateBuildRequest) (*CreateBuildResponse, error)
	GetUploadSlot(context.Context, *UploadSlotRequest) (*UploadSlotResponse, error)
	UploadsComplete(context.Context, *UploadDoneRequest) (*UploadDoneResponse, error)
	ListRepos(context.Context, *ListReposRequest) (*ListReposResponse, error)
	ListBranches(context.Context, *ListBranchesRequest) (*ListBranchesResponse, error)
	ListVersions(context.Context, *ListVersionsRequest) (*ListVersionsResponse, error)
	ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error)
	GetLatestVersion(context.Context, *GetLatestVersionRequest) (*GetLatestVersionResponse, error)
	GetBlock(context.Context, *GetBlockRequest) (*GetBlockResponse, error)
	GetFileMetaData(context.Context, *GetMetaRequest) (*GetMetaResponse, error)
}

func RegisterBuildRepoManagerServer(s *grpc.Server, srv BuildRepoManagerServer) {
	s.RegisterService(&_BuildRepoManager_serviceDesc, srv)
}

func _BuildRepoManager_CreateBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildRepoManagerServer).CreateBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildrepo.BuildRepoManager/CreateBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildRepoManagerServer).CreateBuild(ctx, req.(*CreateBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildRepoManager_GetUploadSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildRepoManagerServer).GetUploadSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildrepo.BuildRepoManager/GetUploadSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildRepoManagerServer).GetUploadSlot(ctx, req.(*UploadSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildRepoManager_UploadsComplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadDoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildRepoManagerServer).UploadsComplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildrepo.BuildRepoManager/UploadsComplete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildRepoManagerServer).UploadsComplete(ctx, req.(*UploadDoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildRepoManager_ListRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReposRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildRepoManagerServer).ListRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildrepo.BuildRepoManager/ListRepos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildRepoManagerServer).ListRepos(ctx, req.(*ListReposRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildRepoManager_ListBranches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBranchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildRepoManagerServer).ListBranches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildrepo.BuildRepoManager/ListBranches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildRepoManagerServer).ListBranches(ctx, req.(*ListBranchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildRepoManager_ListVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildRepoManagerServer).ListVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildrepo.BuildRepoManager/ListVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildRepoManagerServer).ListVersions(ctx, req.(*ListVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildRepoManager_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildRepoManagerServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildrepo.BuildRepoManager/ListFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildRepoManagerServer).ListFiles(ctx, req.(*ListFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildRepoManager_GetLatestVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildRepoManagerServer).GetLatestVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildrepo.BuildRepoManager/GetLatestVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildRepoManagerServer).GetLatestVersion(ctx, req.(*GetLatestVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildRepoManager_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildRepoManagerServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildrepo.BuildRepoManager/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildRepoManagerServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildRepoManager_GetFileMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildRepoManagerServer).GetFileMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildrepo.BuildRepoManager/GetFileMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildRepoManagerServer).GetFileMetaData(ctx, req.(*GetMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BuildRepoManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buildrepo.BuildRepoManager",
	HandlerType: (*BuildRepoManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBuild",
			Handler:    _BuildRepoManager_CreateBuild_Handler,
		},
		{
			MethodName: "GetUploadSlot",
			Handler:    _BuildRepoManager_GetUploadSlot_Handler,
		},
		{
			MethodName: "UploadsComplete",
			Handler:    _BuildRepoManager_UploadsComplete_Handler,
		},
		{
			MethodName: "ListRepos",
			Handler:    _BuildRepoManager_ListRepos_Handler,
		},
		{
			MethodName: "ListBranches",
			Handler:    _BuildRepoManager_ListBranches_Handler,
		},
		{
			MethodName: "ListVersions",
			Handler:    _BuildRepoManager_ListVersions_Handler,
		},
		{
			MethodName: "ListFiles",
			Handler:    _BuildRepoManager_ListFiles_Handler,
		},
		{
			MethodName: "GetLatestVersion",
			Handler:    _BuildRepoManager_GetLatestVersion_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _BuildRepoManager_GetBlock_Handler,
		},
		{
			MethodName: "GetFileMetaData",
			Handler:    _BuildRepoManager_GetFileMetaData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/build-repo.proto",
}

func init() { proto.RegisterFile("proto/build-repo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 777 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0x4b, 0x53, 0x1a, 0x41,
	0x10, 0x2e, 0x74, 0x01, 0x69, 0x31, 0xea, 0x48, 0x99, 0xcd, 0x8a, 0xc6, 0x1a, 0x2f, 0x5e, 0x82,
	0x55, 0x1a, 0x2b, 0x95, 0x4b, 0x0e, 0x42, 0x24, 0xa6, 0x40, 0x93, 0xc5, 0xe4, 0x92, 0xd3, 0x02,
	0xad, 0xd9, 0x12, 0x76, 0x36, 0xbb, 0x63, 0x52, 0xf8, 0x5b, 0x72, 0xc8, 0x4f, 0x4d, 0xcd, 0x63,
	0x61, 0x96, 0x85, 0x94, 0x8f, 0x1c, 0xa7, 0x1f, 0x5f, 0x7f, 0xdd, 0xdb, 0xfd, 0x01, 0x6c, 0x86,
	0x11, 0xe3, 0xec, 0xa0, 0x7b, 0xeb, 0x0f, 0xfa, 0xaf, 0x22, 0x0c, 0x59, 0x4d, 0x1a, 0x48, 0x49,
	0x5a, 0x84, 0xc1, 0xa9, 0x5e, 0x33, 0x76, 0x3d, 0xc0, 0x03, 0x2f, 0xf4, 0x0f, 0xbc, 0x20, 0x60,
	0xdc, 0xe3, 0x3e, 0x0b, 0x62, 0x15, 0x48, 0xff, 0xe4, 0x80, 0xd4, 0x23, 0xf4, 0x38, 0x9e, 0x88,
	0x0c, 0x17, 0x7f, 0xdc, 0x62, 0xcc, 0xc9, 0x0e, 0x80, 0x8b, 0x21, 0x8b, 0x7d, 0xce, 0xa2, 0x91,
	0x9d, 0xdb, 0xcd, 0xed, 0x97, 0x5c, 0x88, 0xc6, 0x16, 0xe2, 0xc0, 0x52, 0x9d, 0x0d, 0x87, 0x3e,
	0x3f, 0x6b, 0xd8, 0x0b, 0xd2, 0xbb, 0xd4, 0xd3, 0x6f, 0xb2, 0x09, 0x85, 0x93, 0xc8, 0x0b, 0x7a,
	0xdf, 0xed, 0x45, 0xe9, 0x29, 0x74, 0xe5, 0x8b, 0xd8, 0x50, 0x94, 0x35, 0xce, 0x1a, 0xb6, 0xb5,
	0x9b, 0xdb, 0xb7, 0xdc, 0x62, 0x57, 0x3d, 0x49, 0x15, 0x4a, 0x0a, 0xad, 0xdd, 0x69, 0xda, 0x79,
	0x99, 0x54, 0xea, 0x25, 0x06, 0xfa, 0x16, 0x36, 0x52, 0x0c, 0xe3, 0x90, 0x05, 0x31, 0x12, 0x0a,
	0x65, 0x69, 0xe8, 0x70, 0x16, 0xa1, 0xdf, 0xd7, 0x24, 0xcb, 0x5d, 0xc3, 0x46, 0x3b, 0xb0, 0xfe,
	0x25, 0x1c, 0x30, 0xaf, 0xdf, 0x19, 0x30, 0x9e, 0xf4, 0x76, 0x8f, 0x44, 0xd1, 0xdf, 0xa9, 0x3f,
	0xc0, 0xc0, 0x1b, 0x62, 0xd2, 0xdf, 0x95, 0x7e, 0xd3, 0x77, 0x40, 0x4c, 0x50, 0x4d, 0xa7, 0x02,
	0xf9, 0x4b, 0x76, 0x83, 0x81, 0x86, 0xcb, 0x73, 0xf1, 0x20, 0x04, 0xac, 0x4f, 0x2c, 0xe2, 0x12,
	0x23, 0xef, 0x5a, 0x21, 0x8b, 0x38, 0x7d, 0x93, 0x90, 0x6a, 0xb0, 0x00, 0x1f, 0x40, 0x8a, 0x56,
	0x92, 0xc2, 0x2a, 0x51, 0x15, 0xa6, 0x47, 0x50, 0x12, 0x9f, 0xea, 0x7d, 0xc0, 0xa3, 0x91, 0xa8,
	0x77, 0x2e, 0x38, 0xab, 0x74, 0x4b, 0xf0, 0x15, 0xb6, 0xcb, 0x51, 0x88, 0x09, 0x07, 0x3e, 0x0a,
	0x91, 0x12, 0x58, 0x6b, 0xf9, 0x31, 0x97, 0xdf, 0x58, 0x53, 0xa0, 0x75, 0x58, 0x37, 0x6c, 0xba,
	0xad, 0x1a, 0x14, 0x05, 0xb2, 0x8f, 0xb1, 0x9d, 0xdb, 0x5d, 0xdc, 0x5f, 0x3e, 0xac, 0xd4, 0xc6,
	0xab, 0x55, 0x1b, 0xd7, 0x75, 0x8b, 0xa8, 0x82, 0xe8, 0x31, 0x6c, 0x08, 0x10, 0xb5, 0x00, 0x18,
	0xdf, 0x73, 0x9f, 0xe8, 0x29, 0x54, 0xd2, 0x69, 0x8f, 0x2c, 0xdf, 0x56, 0xe5, 0xbf, 0x62, 0x14,
	0x8b, 0x25, 0xbf, 0xef, 0x3a, 0x4f, 0x56, 0x76, 0xc1, 0x5c, 0xd9, 0x84, 0xd6, 0x04, 0xee, 0x91,
	0xb4, 0xfa, 0x6a, 0xdc, 0x62, 0xa5, 0x9e, 0xca, 0x49, 0x9c, 0x91, 0xe6, 0xa3, 0xef, 0xab, 0xf8,
	0x53, 0x3d, 0x93, 0x0f, 0xa8, 0xab, 0x3c, 0x92, 0xea, 0x67, 0x78, 0xde, 0x44, 0xde, 0xf2, 0x38,
	0x8e, 0xfb, 0x7e, 0xea, 0x14, 0x5f, 0x83, 0x9d, 0x85, 0xd4, 0xf4, 0x0c, 0x51, 0xc8, 0xa5, 0x44,
	0x81, 0x76, 0x61, 0xb5, 0x89, 0xfc, 0x64, 0xc0, 0x7a, 0x37, 0x09, 0x81, 0x3d, 0xb0, 0xc4, 0x15,
	0xca, 0xc8, 0xe5, 0xc3, 0x55, 0xa3, 0x11, 0xd1, 0xb3, 0x2b, 0x9d, 0x82, 0xc5, 0xc5, 0xd5, 0x55,
	0x8c, 0xea, 0xe8, 0x2c, 0xb7, 0xc0, 0xe4, 0x4b, 0x9c, 0x41, 0xc7, 0xbf, 0x43, 0x39, 0xb4, 0x15,
	0xd7, 0x8a, 0xfd, 0x3b, 0xa4, 0xbf, 0x60, 0x6d, 0x52, 0x43, 0x33, 0xfa, 0xdf, 0x45, 0x84, 0xad,
	0xe1, 0x71, 0x4f, 0x8a, 0x5e, 0xd9, 0xb5, 0xfa, 0x1e, 0xf7, 0x28, 0x07, 0x4b, 0xa0, 0x3d, 0x65,
	0x09, 0x92, 0xb1, 0x2d, 0xa6, 0xb5, 0xd4, 0x54, 0x2e, 0x6b, 0x4a, 0xb9, 0x8e, 0xe1, 0x59, 0x13,
	0x79, 0x1b, 0xb9, 0xf7, 0x90, 0x89, 0xd2, 0x8f, 0xf2, 0x4b, 0xa8, 0xb4, 0x87, 0x0c, 0x29, 0x19,
	0x86, 0x1a, 0x91, 0x1c, 0xc6, 0xe1, 0xef, 0x02, 0xac, 0x69, 0x1d, 0x0f, 0x59, 0xdb, 0x0b, 0xbc,
	0x6b, 0x8c, 0x48, 0x0b, 0x96, 0x0d, 0x85, 0x27, 0xdb, 0x06, 0x5c, 0xf6, 0xb7, 0xc9, 0xd9, 0x99,
	0xe7, 0xd6, 0xdc, 0x5a, 0xb0, 0xd2, 0x44, 0x3e, 0x91, 0x68, 0x52, 0x35, 0x12, 0x32, 0x3f, 0x07,
	0xce, 0xf6, 0x1c, 0xaf, 0x46, 0x3b, 0x87, 0x55, 0x65, 0x8d, 0xeb, 0x6c, 0x18, 0x0e, 0x90, 0xe3,
	0x0c, 0x3c, 0x43, 0xc9, 0x67, 0xe0, 0x99, 0x72, 0x4d, 0x4e, 0xa1, 0x34, 0x56, 0x59, 0xb2, 0x65,
	0xc4, 0x4e, 0xeb, 0xb1, 0x53, 0x9d, 0xed, 0xd4, 0x38, 0x17, 0x50, 0x36, 0x15, 0x93, 0xec, 0x4c,
	0x45, 0x4f, 0x29, 0xb0, 0xf3, 0x72, 0xae, 0x3f, 0x0d, 0x98, 0x68, 0x5d, 0x06, 0x70, 0x4a, 0x53,
	0x33, 0x80, 0x19, 0x91, 0xd4, 0x9d, 0x4a, 0x39, 0xca, 0x74, 0x6a, 0x4a, 0x61, 0xa6, 0xd3, 0xb4,
	0x82, 0x7d, 0x93, 0x47, 0x9a, 0x92, 0x0f, 0x42, 0x8d, 0x8c, 0x39, 0x72, 0xe5, 0xec, 0xfd, 0x33,
	0x46, 0x83, 0xd7, 0x61, 0x29, 0x51, 0x00, 0xe2, 0xa4, 0x13, 0x4c, 0xe9, 0x71, 0xb6, 0x66, 0xfa,
	0x34, 0xc8, 0x07, 0x79, 0x20, 0x82, 0xb5, 0x38, 0x12, 0x71, 0xec, 0xe4, 0x45, 0x3a, 0xde, 0xb8,
	0x39, 0xc7, 0x99, 0xe5, 0x52, 0x48, 0xdd, 0x82, 0xfc, 0x57, 0x76, 0xf4, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x70, 0x37, 0x7d, 0x6e, 0xd8, 0x09, 0x00, 0x00,
}
