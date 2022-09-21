package linutil

import (
	"fmt"
	"github.com/go-delve/delve/pkg/proc"
)

// PPC64LERegisters implements the proc.Registers interface for the native/linux
// backend and core/linux backends, on PPC64LE.
type PPC64LERegisters struct {
	Regs       *PPC64LEPtraceRegs
	Fpregs     []proc.Register //Formatted floating point registers
	Fpregset   []byte          //holding all floating point register values
	loadFpRegs func(*PPC64LERegisters) error
}

func NewPPC64LERegisters(regs *PPC64LEPtraceRegs, loadFpRegs func(*PPC64LERegisters) error) *PPC64LERegisters {
	return &PPC64LERegisters{Regs: regs, loadFpRegs: loadFpRegs}
}

// PPC64LEPtraceRegs is the struct used by the linux kernel to return the
// general purpose registers for PPC64LE CPUs.
// TODO(alexsaezm) Remove the next line:
// Copied from src/syscall/ztypes_linux_ppc64le.go#L518-L532
type PPC64LEPtraceRegs struct {
	Gpr       [32]uint64 // 32 general-purpose registers, each 64 bits wide
	Nip       uint64
	Msr       uint64
	Orig_gpr3 uint64
	Ctr       uint64
	Link      uint64 // Link register -- LLDB dwarf_lr_ppc64le = 65
	Xer       uint64 // Fixed point exception register -- LLDB dwarf_xer_ppc64le = 76
	Ccr       uint64
	Softe     uint64
	Trap      uint64
	Dar       uint64
	Dsisr     uint64
	Result    uint64
	//SP  uint64 // SP pointer r[1]
	//TOC uint64 // TOC pointer -- r[2]
	//TLS uint64 // Thread pointer -- r[13]
	//	CTR uint64 // Loop count register -- LLDB dwarf_ctr_ppc64le = 66
}

func (regs *PPC64LEPtraceRegs) String() string {
	gprs := "\n"
	for i, r := range regs.Gpr {
		gprs += fmt.Sprintf("\t\t - gpr[%d]: %d\n", i, r)
	}
	return fmt.Sprintf("Registers:\n"+
		"\t - Gpr: %s"+
		"\t - Nip: %d\n"+
		"\t - Msr: %d\n"+
		"\t - Orig_gpr3: %d\n"+
		"\t - Ctr: %d\n"+
		"\t - Link: %d\n"+
		"\t - Xer: %d\n"+
		"\t - Ccr: %d\n"+
		"\t - Softe: %d\n"+
		"\t - Trap: %d\n"+
		"\t - Dar: %d\n"+
		"\t - Dsisr: %d\n"+
		"\t - Result: %d\n", gprs, regs.Nip, regs.Msr, regs.Orig_gpr3, regs.Ctr, regs.Link, regs.Xer, regs.Ccr, regs.Softe, regs.Trap, regs.Dar, regs.Dsisr, regs.Result)
}

// PC returns the value of the NIP register
// Also called the IAR/Instruction Address Register or NIP/Next Instruction Pointer
func (r *PPC64LERegisters) PC() uint64 {
	return r.Regs.Nip
}

// SP returns the value of Stack frame pointer stored in Gpr[1].
func (r *PPC64LERegisters) SP() uint64 {
	return r.Regs.Gpr[1]
}

// LR The Link Register is a 64-bit register. It can be
// used to provide the branch target address for the
// Branch Conditional to Link Register instruction, and it
// holds the return address after Branch instructions for
// which LK=1 and after System Call Vectored instructions.
// Extracted from the 2.3.2 section of the PowerISA Book 3.1
func (r *PPC64LERegisters) LR() uint64 {
	return r.Regs.Link
}

func (r *PPC64LERegisters) BP() uint64 {
	return r.Regs.Gpr[30]
}

// TLS returns the value of the thread pointer stored in Gpr[13]
func (r *PPC64LERegisters) TLS() uint64 {
	return r.Regs.Gpr[13]
}

// GAddr returns the address of the G variable
func (r *PPC64LERegisters) GAddr() (uint64, bool) {
	return r.Regs.Gpr[30], true
}

// Slice returns the registers as a list of (name, value) pairs.
func (r *PPC64LERegisters) Slice(floatingPoint bool) ([]proc.Register, error) {
	var regs = []struct {
		k string
		v uint64
	}{
		{"R0", r.Regs.Gpr[0]},
		{"R1", r.Regs.Gpr[1]},
		{"R2", r.Regs.Gpr[2]},
		{"R3", r.Regs.Gpr[3]},
		{"R4", r.Regs.Gpr[4]},
		{"R5", r.Regs.Gpr[5]},
		{"R6", r.Regs.Gpr[6]},
		{"R7", r.Regs.Gpr[7]},
		{"R8", r.Regs.Gpr[8]},
		{"R9", r.Regs.Gpr[9]},
		{"R10", r.Regs.Gpr[10]},
		{"R11", r.Regs.Gpr[11]},
		{"R12", r.Regs.Gpr[12]},
		{"R13", r.Regs.Gpr[13]},
		{"R14", r.Regs.Gpr[14]},
		{"R15", r.Regs.Gpr[15]},
		{"R16", r.Regs.Gpr[16]},
		{"R17", r.Regs.Gpr[17]},
		{"R18", r.Regs.Gpr[18]},
		{"R19", r.Regs.Gpr[19]},
		{"R20", r.Regs.Gpr[20]},
		{"R21", r.Regs.Gpr[21]},
		{"R22", r.Regs.Gpr[22]},
		{"R23", r.Regs.Gpr[23]},
		{"R24", r.Regs.Gpr[24]},
		{"R25", r.Regs.Gpr[25]},
		{"R26", r.Regs.Gpr[26]},
		{"R27", r.Regs.Gpr[27]},
		{"R28", r.Regs.Gpr[28]},
		{"R29", r.Regs.Gpr[29]},
		{"R30", r.Regs.Gpr[30]},
		{"R31", r.Regs.Gpr[31]},
		{"NIP", r.Regs.Nip},
	}
	out := make([]proc.Register, 0, len(regs)+len(r.Fpregs))
	for _, reg := range regs {
		out = proc.AppendUint64Register(out, reg.k, reg.v)
	}
	var floatLoadError error
	if floatingPoint {
		if r.loadFpRegs != nil {
			floatLoadError = r.loadFpRegs(r)
			r.loadFpRegs = nil
		}
		out = append(out, r.Fpregs...)
	}
	return out, floatLoadError
}

// Copy returns a copy of these registers that is guaranteed not to change.
func (r *PPC64LERegisters) Copy() (proc.Registers, error) {
	//TODO(alexsaezm) implement me
	panic("implement me: Copy")
}

type PPC64LEPtraceFpRegs struct {
	//TODO(alexsaezm) Not quite sure about this, review in the future
	Fp []byte
}

func (fpregs *PPC64LEPtraceFpRegs) Decode() (regs []proc.Register) {
	for i := 0; i < len(fpregs.Fp); i += 16 {
		regs = proc.AppendBytesRegister(regs, fmt.Sprintf("V%d", i/16), fpregs.Fp[i:i+16])
	}
	return
}
