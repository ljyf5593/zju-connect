package log

import (
	"bytes"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"os"
)

var debug bool

func Init() {
	log.SetOutput(os.Stdout)
}

func EnableDebug() {
	debug = true
}

func DisableDebug() {
	debug = false
}

func Print(v ...any) {
	log.Print(v...)
}

func DebugPrint(v ...any) {
	if debug {
		log.Print(v...)
	}
}

func Println(v ...any) {
	log.Println(v...)
}

func DebugPrintln(v ...any) {
	if debug {
		log.Println(v...)
	}
}

func Printf(format string, v ...any) {
	log.Printf(format, v...)
}

func DebugPrintf(format string, v ...any) {
	if debug {
		log.Printf(format, v...)
	}
}

func Fatal(v ...any) {
	log.Fatal(v...)
}

func Fatalf(format string, v ...any) {
	log.Fatalf(format, v...)
}

func PrintRequest(req *http.Request) {
	log.Println("Request Details:")
	log.Printf("Method: %s\t,URL: %s\n", req.Method, req.URL.String())
	log.Printf("Header:\n")

	// 打印请求头
	for k, v := range req.Header {
		log.Printf("\t%s: %v\n", k, v)
	}

	// 打印请求体
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 重置请求体

	log.Printf("Body: %s\n", string(bodyBytes))

	// 打印查询参数
	queryParams := req.URL.Query()
	log.Printf("Query Parameters:\n")
	for k, v := range queryParams {
		log.Printf("\t%s: %v\n", k, v)
	}

	// 打印 POST 表单数据
	err = req.ParseForm()
	if err == nil {
		formData := req.Form
		log.Printf("Form Data:\n")
		for k, v := range formData {
			log.Printf("\t%s: %v\n", k, v)
		}
	} else {
		log.Println("No form data found.")
	}
}

func DumpHex(buf []byte) {
	stdoutDumper := hex.Dumper(os.Stdout)
	defer func(stdoutDumper io.WriteCloser) {
		_ = stdoutDumper.Close()
	}(stdoutDumper)
	_, _ = stdoutDumper.Write(buf)
}

func DebugDumpHex(buf []byte) {
	if debug {
		stdoutDumper := hex.Dumper(os.Stdout)
		defer func(stdoutDumper io.WriteCloser) {
			_ = stdoutDumper.Close()
		}(stdoutDumper)
		_, _ = stdoutDumper.Write(buf)
	}
}

func NewLogger(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix, log.LstdFlags)
}
