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
	// TODO(alexsaezm) Implement me
	panic("Unimplemented restoreRegisters method in threads_linux_ppc64le.go")
	return nil
}
