package regnum

import "fmt"

// The mapping between hardware registers and DWARF registers is specified
// in the 64-Bit ELF V2 ABI Specification of the Power Architecture in section
// 2.4 DWARF Definition
// https://openpowerfoundation.org/specifications/64bitelfabi/

const (
	PPC64LE_R0  = 0  // General Purpose Registers: from R0 to R31
	PPC64LE_F0  = 32 // Floating point registers: from F0 to F31
	PPC64LE_V0  = 64 // Vector (Altivec/VMX) registers: from V0 to V31
	PPC64LE_VS0 = 96 // Vector Scalar (VSX) registers: from VS0 to VS63
	PPC64LE_CR0 = 0  // Condition Registers: from CR0 to CR7
	PPC64LE_SP  = 1  // Stack frame pointer: Gpr[1]
	PPC64LE_PC  = 12 // The documentation refers to this as the CIA (Current Instruction Address)
	PPC64LE_LR  = 65 // Link register
)

func PPC64LEToName(num uint64) string {
	switch {
	case PPC64LE_R0 <= num && num <= PPC64LE_F0-1:
		return fmt.Sprintf("r%d", int(num-PPC64LE_R0))
	case PPC64LE_F0 <= num && num <= PPC64LE_V0-1:
		return fmt.Sprintf("f%d", int(num-PPC64LE_F0))
	case PPC64LE_V0 <= num && num <= PPC64LE_VS0-1:
		return fmt.Sprintf("v%d", int(num-PPC64LE_V0))
	case PPC64LE_VS0 <= num && num <= PPC64LE_VS0+63:
		return fmt.Sprintf("vs%d", int(num-PPC64LE_VS0))
	case num == PPC64LE_SP:
		return "SP"
	case num == PPC64LE_PC:
		return "PC"
	case num == PPC64LE_LR:
		return "LR"
	default:
		return fmt.Sprintf("unknown%d", num)
	}
}

func PPC64LEMaxRegNum() uint64 {
	return 160
}

var PPC64LENameToDwarf = func() map[string]int {
	r := make(map[string]int)

	r["nip"] = PPC64LE_PC
	r["sp"] = PPC64LE_SP
	r["bp"] = PPC64LE_SP
	r["link"] = PPC64LE_LR

	// General Purpose Registers: from R0 to R31
	for i := 0; i <= 31; i++ {
		r[fmt.Sprintf("r%d", i)] = PPC64LE_R0 + i
	}

	// Floating point registers: from F0 to F31
	for i := 0; i <= 31; i++ {
		r[fmt.Sprintf("f%d", i)] = PPC64LE_F0 + i
	}

	// Vector (Altivec/VMX) registers: from V0 to V31
	for i := 0; i <= 31; i++ {
		r[fmt.Sprintf("v%d", i)] = PPC64LE_V0 + i
	}

	// Vector Scalar (VSX) registers: from VS0 to VS63
	for i := 0; i <= 63; i++ {
		r[fmt.Sprintf("vs%d", i)] = PPC64LE_VS0 + i
	}

	// Condition Registers: from CR0 to CR7
	for i := 0; i <= 7; i++ {
		r[fmt.Sprintf("cr%d", i)] = PPC64LE_CR0 + i
	}
	return r
}()
