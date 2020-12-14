package main

import (
	"fmt"
	"os"
	"testing"
)

const replaceStr = "${project_config_path}/go/src/webhook/test.sh${project_config_path}/go/src/webhook/test.sh${project_config_path}/go/src/webhook/test.sh${project_config_path}/go/src/webhook/test.sh${project_config_path}/go/src/webhook/test.sh${project_config_path}/go/src/webhook/test.sh${project_config_path}/go/src/webhook/test.sh${project_config_path}/go/src/webhook/test.sh${project_config_path}/go/src/webhook/test.sh"

var replaceStrByte []byte

func TestMain(m *testing.M) {
	fmt.Println("this is main")
	_ = os.Setenv("project_config_path", "hello")
	replaceStrByte = []byte(replaceStr)
	m.Run()
}

func BenchmarkReplaceEnv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ReplaceEnv(replaceStr)
	}
}

func BenchmarkReplaceEnvB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ReplaceEnvB(replaceStrByte)
	}
}
