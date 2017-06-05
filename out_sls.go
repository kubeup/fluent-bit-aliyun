package main

import "github.com/fluent/fluent-bit-go/output"
import (
	"C"
	"fmt"
	sls "github.com/galaxydi/go-loghub"
	"github.com/gogo/protobuf/proto"
	"github.com/ugorji/go/codec"
	"os"
	"reflect"
	"unsafe"
)

var project *sls.LogProject
var logstore *sls.LogStore

//export FLBPluginInit
func FLBPluginInit(ctx unsafe.Pointer) int {
	return FLBPluginRegister(ctx)
}

//export FLBPluginRegister
func FLBPluginRegister(ctx unsafe.Pointer) int {
	accessKey := os.Getenv("ALIYUN_ACCESS_KEY")
	accessKeySecret := os.Getenv("ALIYUN_ACCESS_KEY_SECRET")
	projectName := os.Getenv("ALIYUN_SLS_PROJECT")
	endpoint := os.Getenv("ALIYUN_SLS_ENDPOINT")
	logstoreName := os.Getenv("ALIYUN_SLS_LOGSTORE")

	project = &sls.LogProject{
		Name:            projectName,
		Endpoint:        endpoint,
		AccessKeyID:     accessKey,
		AccessKeySecret: accessKeySecret,
	}

	var err error
	logstore, err = project.GetLogStore(logstoreName)
	if err != nil {
		fmt.Printf("Unable to get logstore %v: %v", logstoreName, err)
		return output.FLB_ERROR
	}
	return output.FLBPluginRegister(ctx, "sls", "Aliyun SLS output")
}

//export FLBPluginFlush
func FLBPluginFlush(data unsafe.Pointer, length C.int, tag *C.char) int {
	var h codec.Handle = new(codec.MsgpackHandle)
	var b []byte
	var m interface{}
	var err error

	if logstore == nil {
		fmt.Printf("logstore is nil")
		return output.FLB_ERROR
	}

	b = C.GoBytes(data, length)
	dec := codec.NewDecoderBytes(b, h)

	// Iterate the original MessagePack array
	logs := []*sls.Log{}
	for {
		// Decode the entry
		err = dec.Decode(&m)
		if err != nil {
			break
		}

		// Get a slice and their two entries: timestamp and map
		slice := reflect.ValueOf(m)
		timestampData := slice.Index(0)
		data := slice.Index(1)
		timestamp, ok := timestampData.Interface().(uint64)
		if !ok {
			fmt.Printf("Unable to convert timestamp: %+v", timestampData)
			return output.FLB_ERROR
		}

		// Convert slice data to a real map and iterate
		mapData := data.Interface().(map[interface{}]interface{})
		content := []*sls.LogContent{}
		for k, v := range mapData {
			content = append(content, &sls.LogContent{
				Key:   proto.String(fmt.Sprintf("%s", k)),
				Value: proto.String(fmt.Sprintf("%s", v)),
			})
		}
		log := &sls.Log{
			Time:     proto.Uint32(uint32(timestamp)),
			Contents: content,
		}
		logs = append(logs, log)
	}
	loggroup := &sls.LogGroup{
		Topic:  proto.String(""),
		Source: proto.String("10.230.201.117"),
		Logs:   logs,
	}
	err = logstore.PutLogs(loggroup)
	if err != nil {
		return output.FLB_ERROR
	}

	// Return options:
	//
	// output.FLB_OK    = data have been processed.
	// output.FLB_ERROR = unrecoverable error, do not try this again.
	// output.FLB_RETRY = retry to flush later.
	return output.FLB_OK
}

//export FLBPluginExit
func FLBPluginExit() int {
	return 0
}

func main() {
}
