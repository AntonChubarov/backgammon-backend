package main

import (
	"github.com/jbrodriguez/mlog"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)

	go func() {
		mlog.Fatal(http.ListenAndServe("localhost:6060", nil))
	}()

	SetupAndRun()
}
