package main

import (
	"bytes"
	"compress/gzip"
	"example/test_pb/helloworld"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/descriptor"
	"github.com/jhump/protoreflect/desc"

	"github.com/golang/protobuf/proto"
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func main() {
	//dir, _ := os.Getwd()
	//fdTest, fdTestByte := loadFileDesc("helloworld.proto")
	//
	//fmt.Println(fdTest.Dependency)
	//fmt.Println(string(fdTestByte))

	fd, err := loadProtoset("./helloworld/desc_test1.protoset")
	if err != nil {
		panic(err)
	}

	fmt.Println(fd.GetFile())

	req := helloworld.HelloRequest{}
	msg, md := descriptor.ForMessage(&req)
	fmt.Println(msg)
	fmt.Println(md)
}

func loadProtoset(path string) (*desc.FileDescriptor, error) {
	var fds dpb.FileDescriptorSet
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bb, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if err = proto.Unmarshal(bb, &fds); err != nil {
		return nil, err
	}
	return desc.CreateFileDescriptorFromSet(&fds)
}

func loadFileDesc(filename string) (*dpb.FileDescriptorProto, []byte) {
	enc := proto.FileDescriptor(filename)
	if enc == nil {
		panic(fmt.Sprintf("failed to find fd for file: %v", filename))
	}
	fd, err := decodeFileDesc(enc)
	if err != nil {
		panic(fmt.Sprintf("failed to decode enc: %v", err))
	}
	b, err := proto.Marshal(fd)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal fd: %v", err))
	}
	return fd, b
}

// decodeFileDesc does decompression and unmarshalling on the given
// file descriptor byte slice.
func decodeFileDesc(enc []byte) (*dpb.FileDescriptorProto, error) {
	raw, err := decompress(enc)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress enc: %v", err)
	}

	fd := new(dpb.FileDescriptorProto)
	if err := proto.Unmarshal(raw, fd); err != nil {
		return nil, fmt.Errorf("bad descriptor: %v", err)
	}
	return fd, nil
}

// decompress does gzip decompression.
func decompress(b []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("bad gzipped descriptor: %v", err)
	}
	out, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("bad gzipped descriptor: %v", err)
	}
	return out, nil
}
