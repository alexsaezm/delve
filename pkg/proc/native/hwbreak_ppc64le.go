package native

import (
	"github.com/go-delve/delve/pkg/proc"
)

func (thread *nativeThread) findHardwareBreakpoint() (*proc.Breakpoint, error) {
	// TODO(alexsaezm) Implement this method
	panic("Unimplemented findHardwareBreakpoint method in threads_linux_ppc64le.go")
	return nil, nil
}

func (thread *nativeThread) writeHardwareBreakpoint(addr uint64, wtype proc.WatchType, idx uint8) error {
	// TODO(alexsaezm) Implement this method
	panic("Unimplemented writeHardwareBreakpoint method in threads_linux_ppc64le.go")
	return nil
}

func (thread *nativeThread) clearHardwareBreakpoint(addr uint64, wtype proc.WatchType, idx uint8) error {
	// TODO(alexsaezm) Implement this method
	panic("Unimplemented clearHardwareBreakpoint method in threads_linux_ppc64le.go")
	return nil
}
