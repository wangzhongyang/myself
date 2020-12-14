package main

import (
	"encoding/json"
	"testing"

	json2 "github.com/json-iterator/go"
	"github.com/pquerna/ffjson/ffjson"
)

//easyjson:json
type VoiceImportRec struct {
	FileId            string   `protobuf:"bytes,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
	OriginPath        string   `protobuf:"bytes,2,opt,name=originPath,proto3" json:"originPath,omitempty"`
	TotalNum          int32    `protobuf:"varint,3,opt,name=totalNum,proto3" json:"totalNum,omitempty"`
	FailedNum         int32    `protobuf:"varint,4,opt,name=failedNum,proto3" json:"failedNum,omitempty"`
	SucceedNum        int32    `protobuf:"varint,5,opt,name=succeedNum,proto3" json:"succeedNum,omitempty"`
	Status            string   `protobuf:"varint,6,opt,name=status,proto3,enum=data_def.VOICE_IMPORT_STATUS" json:"status,omitempty"`
	PreImpRN          int32    `protobuf:"varint,7,opt,name=preImpRN,proto3" json:"preImpRN,omitempty"`
	PersistRN         int32    `protobuf:"varint,8,opt,name=persistRN,proto3" json:"persistRN,omitempty"`
	CopiedFiles       []string `protobuf:"bytes,9,rep,name=copiedFiles,proto3" json:"copiedFiles,omitempty"`
	CreateTime        int64    `protobuf:"varint,10,opt,name=createTime,proto3" json:"createTime,omitempty"`
	LogPath           string   `protobuf:"bytes,11,opt,name=logPath,proto3" json:"logPath,omitempty"`
	StorageSucceedNum int32    `protobuf:"varint,12,opt,name=storageSucceedNum,proto3" json:"storageSucceedNum,omitempty"`
	StorageFailedNum  int32    `protobuf:"varint,13,opt,name=storageFailedNum,proto3" json:"storageFailedNum,omitempty"`
	StorageNum        int32    `protobuf:"varint,14,opt,name=storageNum,proto3" json:"storageNum,omitempty"`
}

var voice = VoiceImportRec{
	FileId:            "11111",
	OriginPath:        "22222",
	TotalNum:          1,
	FailedNum:         2,
	SucceedNum:        3,
	Status:            "33333",
	PreImpRN:          4,
	PersistRN:         5,
	CopiedFiles:       []string{"11", "22", "33"},
	CreateTime:        2312312,
	LogPath:           "44444",
	StorageSucceedNum: 6,
	StorageFailedNum:  7,
	StorageNum:        8,
}

func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(voice)
	}
}

func BenchmarkIterator(b *testing.B) {
	//var json22 = jsoniter.ConfigCompatibleWithStandardLibrary
	//b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = json2.Marshal(voice)
	}
}

func BenchmarkFFjson(b *testing.B) {
	//b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = ffjson.Marshal(voice)
	}
}
