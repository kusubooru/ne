// +build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	release   = flag.Bool("release", false, "Build binaries for all target platforms.")
	clean     = flag.Bool("clean", false, "Remove all created binaries from current directory.")
	buildARCH = flag.String("arch", runtime.GOARCH, "Architecture to build for.")
	buildOS   = flag.String("os", runtime.GOOS, "Operating system to build for.")
	stub      = flag.String("pb", "", "Operating system to build for.")
	gateway   = flag.String("gw", "", "Operating system to build for.")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: go run make.go [OPTIONS]\n\n")
	fmt.Fprintln(os.Stderr, "OPTIONS:")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		fmt.Fprintln(os.Stderr, "GOPATH is not set")
		return
	}

	generateStub(gopath)
	generateGateway(gopath)

}

func gatherProtoFiles() []string {
	files, err := filepath.Glob("*.proto")
	if err != nil {
		log.Fatalln("error gathering *.proto files:", err)
	}
	return files
}

func generateStub(gopath string, protoFiles ...string) {
	if len(protoFiles) == 0 {
		protoFiles = gatherProtoFiles()
	}

	var stderr bytes.Buffer
	cmdArgs := []string{
		//"-I/usr/local/include",
		"-I.",
		//fmt.Sprintf("-I%s/src", gopath),
		fmt.Sprintf("-I%s/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis", gopath),
		"--go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:.",
	}
	cmdArgs = append(cmdArgs, protoFiles...)
	cmd := exec.Command("protoc", cmdArgs...)
	cmd.Stderr = &stderr

	cmd.Env = copyGoEnv()
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error generating gRPC stub:", stderr.String())
		os.Exit(1)
	}
}

func generateGateway(gopath string, protoFiles ...string) {
	if len(protoFiles) == 0 {
		protoFiles = gatherProtoFiles()
	}

	var stderr bytes.Buffer
	cmdArgs := []string{
		//"-I/usr/local/include",
		"-I.",
		//fmt.Sprintf("-I%s/src", gopath),
		fmt.Sprintf("-I%s/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis", gopath),
		"--grpc-gateway_out=logtostderr=true:.",
	}
	cmdArgs = append(cmdArgs, protoFiles...)
	cmd := exec.Command("protoc", cmdArgs...)
	cmd.Stderr = &stderr

	cmd.Env = copyGoEnv()
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error generating gRPC gateway:", stderr.String())
		os.Exit(1)
	}
}

func copyGoEnv() (environ []string) {
	for _, env := range os.Environ() {
		environ = append(environ, env)
	}
	return
}

func setEnv(env []string, key, value string) []string {
	for i, e := range env {
		if strings.HasPrefix(e, fmt.Sprintf("%s=", key)) {
			env[i] = fmt.Sprintf("%s=%s", key, value)
			return env
		}
	}
	env = append(env, fmt.Sprintf("%s=%s", key, value))
	return env
}
