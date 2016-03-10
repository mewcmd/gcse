// Code generated by protoc-gen-go.
// source: github.com/daviddengcn/gcse/proto/spider/spider.proto
// DO NOT EDIT!

/*
Package gcse_spider is a generated protocol buffer package.

It is generated from these files:
	github.com/daviddengcn/gcse/proto/spider/spider.proto

It has these top-level messages:
	GoFileInfo
	RepoInfo
	FolderInfo
	CrawlingInfo
	HistoryEvent
	HistoryInfo
*/
package gcse_spider

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type GoFileInfo_Status int32

const (
	GoFileInfo_Unknown      GoFileInfo_Status = 0
	GoFileInfo_ParseSuccess GoFileInfo_Status = 1
	GoFileInfo_ParseFailed  GoFileInfo_Status = 2
	GoFileInfo_ShouldIgnore GoFileInfo_Status = 3
)

var GoFileInfo_Status_name = map[int32]string{
	0: "Unknown",
	1: "ParseSuccess",
	2: "ParseFailed",
	3: "ShouldIgnore",
}
var GoFileInfo_Status_value = map[string]int32{
	"Unknown":      0,
	"ParseSuccess": 1,
	"ParseFailed":  2,
	"ShouldIgnore": 3,
}

func (x GoFileInfo_Status) String() string {
	return proto.EnumName(GoFileInfo_Status_name, int32(x))
}
func (GoFileInfo_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type HistoryEvent_Action_Enum int32

const (
	HistoryEvent_Action_None    HistoryEvent_Action_Enum = 0
	HistoryEvent_Action_Success HistoryEvent_Action_Enum = 1
	HistoryEvent_Action_Failed  HistoryEvent_Action_Enum = 2
	HistoryEvent_Action_Invalid HistoryEvent_Action_Enum = 3
)

var HistoryEvent_Action_Enum_name = map[int32]string{
	0: "None",
	1: "Success",
	2: "Failed",
	3: "Invalid",
}
var HistoryEvent_Action_Enum_value = map[string]int32{
	"None":    0,
	"Success": 1,
	"Failed":  2,
	"Invalid": 3,
}

func (x HistoryEvent_Action_Enum) String() string {
	return proto.EnumName(HistoryEvent_Action_Enum_name, int32(x))
}
func (HistoryEvent_Action_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0, 0}
}

type GoFileInfo struct {
	Status      GoFileInfo_Status `protobuf:"varint,1,opt,name=status,enum=gcse.spider.GoFileInfo_Status" json:"status,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description string            `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	IsTest      bool              `protobuf:"varint,4,opt,name=is_test,json=isTest" json:"is_test,omitempty"`
	Imports     []string          `protobuf:"bytes,5,rep,name=imports" json:"imports,omitempty"`
}

func (m *GoFileInfo) Reset()                    { *m = GoFileInfo{} }
func (m *GoFileInfo) String() string            { return proto.CompactTextString(m) }
func (*GoFileInfo) ProtoMessage()               {}
func (*GoFileInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RepoInfo struct {
	// The timestamp this repo-info is crawled
	CrawlingTime *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=crawling_time,json=crawlingTime" json:"crawling_time,omitempty"`
	Stars        int32                      `protobuf:"varint,2,opt,name=stars" json:"stars,omitempty"`
	Description  string                     `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// Where this project was forked from, full path
	Source string `protobuf:"bytes,5,opt,name=source" json:"source,omitempty"`
	// As far as we know, when this repo was updated
	LastUpdated *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=last_updated,json=lastUpdated" json:"last_updated,omitempty"`
}

func (m *RepoInfo) Reset()                    { *m = RepoInfo{} }
func (m *RepoInfo) String() string            { return proto.CompactTextString(m) }
func (*RepoInfo) ProtoMessage()               {}
func (*RepoInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RepoInfo) GetCrawlingTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CrawlingTime
	}
	return nil
}

func (m *RepoInfo) GetLastUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// Information for a non-repository folder.
type FolderInfo struct {
	// E.g. "sub"
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// E.g. "spider/sub"
	Path    string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Sha     string `protobuf:"bytes,3,opt,name=sha" json:"sha,omitempty"`
	HtmlUrl string `protobuf:"bytes,4,opt,name=html_url,json=htmlUrl" json:"html_url,omitempty"`
	// The timestamp this folder-info is crawled
	CrawlingTime *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=crawling_time,json=crawlingTime" json:"crawling_time,omitempty"`
}

func (m *FolderInfo) Reset()                    { *m = FolderInfo{} }
func (m *FolderInfo) String() string            { return proto.CompactTextString(m) }
func (*FolderInfo) ProtoMessage()               {}
func (*FolderInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FolderInfo) GetCrawlingTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CrawlingTime
	}
	return nil
}

type CrawlingInfo struct {
	// The timestamp the related entry was crawled
	CrawlingTime *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=crawling_time,json=crawlingTime" json:"crawling_time,omitempty"`
	Etag         string                     `protobuf:"bytes,2,opt,name=etag" json:"etag,omitempty"`
}

func (m *CrawlingInfo) Reset()                    { *m = CrawlingInfo{} }
func (m *CrawlingInfo) String() string            { return proto.CompactTextString(m) }
func (*CrawlingInfo) ProtoMessage()               {}
func (*CrawlingInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CrawlingInfo) GetCrawlingTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CrawlingTime
	}
	return nil
}

type HistoryEvent struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Action    HistoryEvent_Action_Enum   `protobuf:"varint,2,opt,name=action,enum=gcse.spider.HistoryEvent_Action_Enum" json:"action,omitempty"`
}

func (m *HistoryEvent) Reset()                    { *m = HistoryEvent{} }
func (m *HistoryEvent) String() string            { return proto.CompactTextString(m) }
func (*HistoryEvent) ProtoMessage()               {}
func (*HistoryEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *HistoryEvent) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type HistoryEvent_Action struct {
}

func (m *HistoryEvent_Action) Reset()                    { *m = HistoryEvent_Action{} }
func (m *HistoryEvent_Action) String() string            { return proto.CompactTextString(m) }
func (*HistoryEvent_Action) ProtoMessage()               {}
func (*HistoryEvent_Action) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type HistoryInfo struct {
	Events    []*HistoryEvent            `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
	FoundTime *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=found_time,json=foundTime" json:"found_time,omitempty"`
	// Possible value:
	//   web                  added from web
	//   user:<user>          found from user crawling
	//   parent               found by crawling his parent
	//   imported:<pkg>       imported by a <pkg>
	//   testimported:<pkg>   test imported by a <pkg>
	//   package:<pkg>
	//   reference:<pkg>      referenced in the readme file of <pkg>
	//   godoc                found by godoc.org/api
	FoundWay      string                     `protobuf:"bytes,3,opt,name=found_way,json=foundWay" json:"found_way,omitempty"`
	LatestSuccess *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=latest_success,json=latestSuccess" json:"latest_success,omitempty"`
	LatestFailed  *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=latest_failed,json=latestFailed" json:"latest_failed,omitempty"`
}

func (m *HistoryInfo) Reset()                    { *m = HistoryInfo{} }
func (m *HistoryInfo) String() string            { return proto.CompactTextString(m) }
func (*HistoryInfo) ProtoMessage()               {}
func (*HistoryInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *HistoryInfo) GetEvents() []*HistoryEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *HistoryInfo) GetFoundTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.FoundTime
	}
	return nil
}

func (m *HistoryInfo) GetLatestSuccess() *google_protobuf.Timestamp {
	if m != nil {
		return m.LatestSuccess
	}
	return nil
}

func (m *HistoryInfo) GetLatestFailed() *google_protobuf.Timestamp {
	if m != nil {
		return m.LatestFailed
	}
	return nil
}

func init() {
	proto.RegisterType((*GoFileInfo)(nil), "gcse.spider.GoFileInfo")
	proto.RegisterType((*RepoInfo)(nil), "gcse.spider.RepoInfo")
	proto.RegisterType((*FolderInfo)(nil), "gcse.spider.FolderInfo")
	proto.RegisterType((*CrawlingInfo)(nil), "gcse.spider.CrawlingInfo")
	proto.RegisterType((*HistoryEvent)(nil), "gcse.spider.HistoryEvent")
	proto.RegisterType((*HistoryEvent_Action)(nil), "gcse.spider.HistoryEvent.Action")
	proto.RegisterType((*HistoryInfo)(nil), "gcse.spider.HistoryInfo")
	proto.RegisterEnum("gcse.spider.GoFileInfo_Status", GoFileInfo_Status_name, GoFileInfo_Status_value)
	proto.RegisterEnum("gcse.spider.HistoryEvent_Action_Enum", HistoryEvent_Action_Enum_name, HistoryEvent_Action_Enum_value)
}

var fileDescriptor0 = []byte{
	// 636 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x53, 0xdd, 0x6a, 0x14, 0x4d,
	0x10, 0xfd, 0x66, 0x7f, 0x26, 0x9b, 0x9a, 0x4d, 0xbe, 0xa5, 0x11, 0xdd, 0x44, 0x90, 0x30, 0x20,
	0x78, 0x35, 0x8b, 0x11, 0x83, 0x22, 0xa2, 0x41, 0x12, 0x8d, 0x17, 0x22, 0x93, 0x04, 0x2f, 0x97,
	0xce, 0x4c, 0xef, 0x6c, 0x63, 0x4f, 0xf7, 0xd0, 0xdd, 0x93, 0x65, 0xdf, 0xc5, 0x2b, 0xdf, 0xc6,
	0x57, 0xf0, 0x4d, 0xbc, 0xb3, 0x7f, 0x66, 0xb2, 0x2b, 0xa2, 0xae, 0xe0, 0xd5, 0x74, 0x55, 0x9d,
	0x9a, 0xaa, 0x3a, 0xa7, 0x0a, 0x1e, 0x17, 0x54, 0xcf, 0xeb, 0xab, 0x24, 0x13, 0xe5, 0x24, 0xc7,
	0xd7, 0x34, 0xcf, 0x09, 0x2f, 0x32, 0x3e, 0x29, 0x32, 0x45, 0x26, 0x95, 0x14, 0x5a, 0x4c, 0x54,
	0x45, 0x73, 0x22, 0x9b, 0x4f, 0xe2, 0x7c, 0x28, 0xb2, 0xf1, 0xc4, 0xbb, 0xf6, 0x9f, 0xad, 0xfd,
	0xa3, 0x10, 0x0c, 0xf3, 0xc2, 0x67, 0x5e, 0xd5, 0xb3, 0x49, 0xa5, 0x97, 0x15, 0x51, 0x13, 0x4d,
	0x4b, 0xa2, 0x34, 0x2e, 0xab, 0xd5, 0xcb, 0xff, 0x29, 0xfe, 0x16, 0x00, 0xbc, 0x16, 0xa7, 0x94,
	0x91, 0x33, 0x3e, 0x13, 0xe8, 0x08, 0x42, 0x13, 0xd5, 0xb5, 0x1a, 0x07, 0x07, 0xc1, 0x83, 0xdd,
	0xc3, 0x7b, 0xc9, 0x5a, 0xa5, 0x64, 0x05, 0x4c, 0xce, 0x1d, 0x2a, 0x6d, 0xd0, 0x08, 0x41, 0x8f,
	0xe3, 0x92, 0x8c, 0x3b, 0x26, 0x6b, 0x3b, 0x75, 0x6f, 0x74, 0x00, 0x51, 0x4e, 0x54, 0x26, 0x69,
	0xa5, 0xa9, 0xe0, 0xe3, 0xae, 0x0b, 0xad, 0xbb, 0xd0, 0x1d, 0xd8, 0xa2, 0x6a, 0xaa, 0x4d, 0x43,
	0xe3, 0x9e, 0x89, 0x0e, 0xd2, 0x90, 0xaa, 0x0b, 0x63, 0xa1, 0xb1, 0x09, 0x94, 0x95, 0x90, 0x5a,
	0x8d, 0xfb, 0x07, 0x5d, 0x93, 0xd6, 0x9a, 0xf1, 0x5b, 0x08, 0x7d, 0x69, 0x14, 0xc1, 0xd6, 0x25,
	0xff, 0xc8, 0xc5, 0x82, 0x8f, 0xfe, 0x43, 0x23, 0x18, 0xbe, 0xc7, 0x52, 0x91, 0xf3, 0x3a, 0xcb,
	0x88, 0x52, 0xa3, 0x00, 0xfd, 0x0f, 0x91, 0xf3, 0x9c, 0x62, 0xd3, 0x72, 0x3e, 0xea, 0x58, 0xc8,
	0xf9, 0x5c, 0xd4, 0x2c, 0x3f, 0x2b, 0xb8, 0x90, 0x64, 0xd4, 0x8d, 0xbf, 0x06, 0x30, 0x48, 0x49,
	0x25, 0xdc, 0xe4, 0x2f, 0x60, 0x27, 0x93, 0x78, 0xc1, 0x28, 0x2f, 0xa6, 0x96, 0x24, 0x47, 0x40,
	0x74, 0xb8, 0x9f, 0x14, 0x42, 0x14, 0x8c, 0x24, 0x2d, 0xa5, 0xc9, 0x45, 0xcb, 0x60, 0x3a, 0x6c,
	0x13, 0xac, 0x0b, 0xdd, 0x82, 0xbe, 0x71, 0x4b, 0xe5, 0x38, 0xe8, 0xa7, 0xde, 0xd8, 0x80, 0x84,
	0xdb, 0x86, 0x72, 0x51, 0xcb, 0x8c, 0x98, 0x51, 0x6d, 0xb0, 0xb1, 0xd0, 0x73, 0x18, 0x32, 0xac,
	0xf4, 0xb4, 0xae, 0x72, 0xac, 0x49, 0xee, 0x18, 0xfa, 0x7d, 0x3f, 0x91, 0xc5, 0x5f, 0x7a, 0x78,
	0xfc, 0xd9, 0x08, 0x7b, 0x2a, 0x98, 0x91, 0xcd, 0x8d, 0xd7, 0x0a, 0x14, 0xac, 0x09, 0x64, 0x7c,
	0x15, 0xd6, 0xf3, 0x56, 0x34, 0xfb, 0x36, 0x2c, 0x75, 0xd5, 0x1c, 0x37, 0x7d, 0xda, 0x27, 0xda,
	0x83, 0xc1, 0x5c, 0x97, 0x6c, 0x5a, 0x4b, 0xe6, 0x7a, 0x30, 0x62, 0x58, 0xfb, 0x52, 0xb2, 0x9f,
	0x39, 0xeb, 0xff, 0x1d, 0x67, 0x71, 0x06, 0xc3, 0x57, 0x8d, 0xfd, 0x6f, 0x44, 0x30, 0x23, 0x11,
	0x8d, 0x8b, 0x76, 0x24, 0xfb, 0x8e, 0xbf, 0x04, 0x30, 0x7c, 0x43, 0x95, 0x16, 0x72, 0x79, 0x72,
	0x4d, 0xb8, 0x46, 0x4f, 0x60, 0xfb, 0xe6, 0x0c, 0x36, 0xa8, 0xb0, 0x02, 0x1b, 0x4d, 0x42, 0x9c,
	0x39, 0x21, 0x3b, 0xee, 0x3c, 0xee, 0xff, 0x70, 0x1e, 0xeb, 0x45, 0x92, 0x63, 0x87, 0x4b, 0x4e,
	0x78, 0x5d, 0xa6, 0x4d, 0xd2, 0xfe, 0x4b, 0x08, 0xbd, 0x3b, 0x3e, 0x82, 0x9e, 0x8d, 0xa0, 0x01,
	0xf4, 0xde, 0x09, 0x4e, 0xcc, 0x06, 0x9b, 0x75, 0x5e, 0x2d, 0x2f, 0x40, 0x78, 0xb3, 0xb7, 0x26,
	0x70, 0xc6, 0xaf, 0x31, 0xa3, 0xb9, 0x59, 0xd9, 0x4f, 0x1d, 0x88, 0x9a, 0x32, 0x8e, 0xb0, 0x87,
	0x10, 0x12, 0x5b, 0xce, 0xde, 0x6b, 0xd7, 0xcc, 0xb1, 0xf7, 0xcb, 0x86, 0xd2, 0x06, 0x88, 0x9e,
	0x02, 0xcc, 0x44, 0xcd, 0x73, 0x4f, 0x70, 0xe7, 0xcf, 0xe3, 0x3b, 0xb4, 0x63, 0xf7, 0x2e, 0x78,
	0x63, 0xba, 0xc0, 0xcb, 0x66, 0x45, 0x06, 0xce, 0xf1, 0x01, 0x2f, 0xd1, 0x31, 0xec, 0x32, 0x6c,
	0x6f, 0x79, 0xaa, 0xfc, 0x1c, 0x1b, 0x6c, 0xec, 0x8e, 0xcf, 0x68, 0x06, 0xb7, 0xf2, 0x37, 0xbf,
	0x98, 0xb9, 0xe9, 0x37, 0xd9, 0x27, 0x9f, 0xe0, 0xd9, 0xba, 0x0a, 0x1d, 0xe2, 0xd1, 0xf7, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x97, 0x0e, 0x34, 0x31, 0x57, 0x05, 0x00, 0x00,
}
