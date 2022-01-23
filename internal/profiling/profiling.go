package profiling

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

const (
	BASE_PATH = "/usr/local/var/log/"
)

func ProfileCPU() (stop func()) {
	f, err := os.Create(BASE_PATH + "profilecpu.out")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	return pprof.StopCPUProfile
}

func ProfileMem() {
	f, err := os.Create(BASE_PATH + "profilemem.out")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}
