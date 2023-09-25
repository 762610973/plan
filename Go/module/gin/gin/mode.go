package gin

import (
	"flag"
	"io"
	"os"

	"github.com/gin-gonic/gin/binding"
)

// EnvGinMode gin mode的环境变量名
const EnvGinMode = "GIN_MODE"

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)

const (
	debugCode = iota
	releaseCode
	testCode
)

//		DefaultWriter
//	 gin默认将信息输出到标准输出中
//		import "github.com/mattn/go-colorable"
var DefaultWriter io.Writer = os.Stdout

// DefaultErrorWriter debug error默认写入标准错误
var DefaultErrorWriter io.Writer = os.Stderr

var (
	ginMode  = debugCode
	modeName = DebugMode
)

func init() {
	// 可以从环境变量中读取mode
	mode := os.Getenv(EnvGinMode)
	SetMode(mode)
}

// SetMode 设置gin的模式
func SetMode(value string) {
	if value == "" {
		if flag.Lookup("test.v") != nil {
			value = TestMode
		} else {
			value = DebugMode
		}
	}

	switch value {
	case DebugMode:
		ginMode = debugCode
	case ReleaseMode:
		ginMode = releaseCode
	case TestMode:
		ginMode = testCode
	default:
		panic("gin mode unknown: " + value + " (available mode: debug release test)")
	}

	modeName = value
}

func DisableBindValidation() {
	binding.Validator = nil
}

func EnableJsonDecoderUseNumber() {
	binding.EnableDecoderUseNumber = true
}

func EnableJsonDecoderDisallowUnknownFields() {
	binding.EnableDecoderDisallowUnknownFields = true
}

func Mode() string {
	return modeName
}
