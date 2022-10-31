package native

import (
	"fmt"
	"github.com/go-delve/delve/pkg/proc"
	"github.com/go-delve/delve/pkg/proc/linutil"
)

func (t *nativeThread) fpRegisters() ([]proc.Register, []byte, error) {
	var err error
	var ppc64leFpregs linutil.PPC64LEPtraceFpRegs
	t.dbp.execPtraceFunc(func() { ppc64leFpregs.Fp, err = ptraceGetFpRegset(t.ID) })
	fpregs := ppc64leFpregs.Decode()
	if err != nil {
		err = fmt.Errorf("could not get floating point registers: %v", err.Error())
	}
	return fpregs, ppc64leFpregs.Fp, err
}

func (t *nativeThread) restoreRegisters(savedRegs proc.Registers) error {
	// TODO(alexsaezm) Implement restoreRegisters
	panic("Unimplemented restoreRegisters method in threads_linux_ppc64le.go")
	return nil
}

// TODO(alexsaezm) Verify if this makes sense in PPC64LE, I copied from threads_linux_arm64.go
type watchpointState struct {
	num      uint8
	debugVer uint8
	words    []uint64
}

// TODO(alexsaezm) Verify if this makes sense in PPC64LE, I copied from threads_linux_arm64.go
func (t *nativeThread) getWatchpoints() (*watchpointState, error) {
	panic("Unimplemented getWatchpoints method in threads_linux_ppc64le.go")
}

// TODO(alexsaezm) Verify if this makes sense in PPC64LE, I copied from threads_linux_arm64.go
func (t *nativeThread) setWatchpoints(wpstate *watchpointState) error {
	panic("Unimplemented setWatchpoints method in threads_linux_ppc64le.go")
}
