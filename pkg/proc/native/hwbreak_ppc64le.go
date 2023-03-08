package native

import (
	"github.com/go-delve/delve/pkg/proc"
)

func (t *nativeThread) findHardwareBreakpoint() (*proc.Breakpoint, error) {
	panic("Unimplemented findHardwareBreakpoint method in threads_linux_ppc64le.go")
}

func (t *nativeThread) writeHardwareBreakpoint(addr uint64, wtype proc.WatchType, idx uint8) error {
	panic("hardware breakpoints not implemented on ppc64le")
}

func (t *nativeThread) clearHardwareBreakpoint(addr uint64, wtype proc.WatchType, idx uint8) error {
	panic("Unimplemented clearHardwareBreakpoint method in threads_linux_ppc64le.go")
}
