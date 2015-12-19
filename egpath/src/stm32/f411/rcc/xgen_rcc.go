package rcc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"
)


func rcc(n uint) *mmio.U32 {
	return &(*[36]mmio.U32)(unsafe.Pointer(uintptr(0x40023800)))[n]
}


type CR_Bits uint32

func (m CR_Bits) Set()           { rcc(0).SetBits(uint32(m)) }
func (m CR_Bits) Clear()         { rcc(0).ClearBits(uint32(m)) }
func (m CR_Bits) Load() uint32   { return rcc(0).Bits(uint32(m)) }
func (m CR_Bits) Store(b uint32) { rcc(0).StoreBits(uint32(m), b) }
func (m CR_Bits) LoadVal() int   { return rcc(0).Field(uint32(m)) }
func (m CR_Bits) StoreVal(v int) { rcc(0).SetField(uint32(m), v) }

func CR_Load() CR_Bits   { return CR_Bits(rcc(0).Load()) }
func CR_Store(b CR_Bits) { rcc(0).Store(uint32(b)) }

func (b CR_Bits) Field(mask CR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_CR(v int, mask CR_Bits) CR_Bits {
	return CR_Bits(bits.Make32(v, uint32(mask)))
}


type PLLCFGR_Bits uint32

func (m PLLCFGR_Bits) Set()           { rcc(1).SetBits(uint32(m)) }
func (m PLLCFGR_Bits) Clear()         { rcc(1).ClearBits(uint32(m)) }
func (m PLLCFGR_Bits) Load() uint32   { return rcc(1).Bits(uint32(m)) }
func (m PLLCFGR_Bits) Store(b uint32) { rcc(1).StoreBits(uint32(m), b) }
func (m PLLCFGR_Bits) LoadVal() int   { return rcc(1).Field(uint32(m)) }
func (m PLLCFGR_Bits) StoreVal(v int) { rcc(1).SetField(uint32(m), v) }

func PLLCFGR_Load() PLLCFGR_Bits   { return PLLCFGR_Bits(rcc(1).Load()) }
func PLLCFGR_Store(b PLLCFGR_Bits) { rcc(1).Store(uint32(b)) }

func (b PLLCFGR_Bits) Field(mask PLLCFGR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_PLLCFGR(v int, mask PLLCFGR_Bits) PLLCFGR_Bits {
	return PLLCFGR_Bits(bits.Make32(v, uint32(mask)))
}


type CFGR_Bits uint32

func (m CFGR_Bits) Set()           { rcc(2).SetBits(uint32(m)) }
func (m CFGR_Bits) Clear()         { rcc(2).ClearBits(uint32(m)) }
func (m CFGR_Bits) Load() uint32   { return rcc(2).Bits(uint32(m)) }
func (m CFGR_Bits) Store(b uint32) { rcc(2).StoreBits(uint32(m), b) }
func (m CFGR_Bits) LoadVal() int   { return rcc(2).Field(uint32(m)) }
func (m CFGR_Bits) StoreVal(v int) { rcc(2).SetField(uint32(m), v) }

func CFGR_Load() CFGR_Bits   { return CFGR_Bits(rcc(2).Load()) }
func CFGR_Store(b CFGR_Bits) { rcc(2).Store(uint32(b)) }

func (b CFGR_Bits) Field(mask CFGR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_CFGR(v int, mask CFGR_Bits) CFGR_Bits {
	return CFGR_Bits(bits.Make32(v, uint32(mask)))
}


type CIR_Bits uint32

func (m CIR_Bits) Set()           { rcc(3).SetBits(uint32(m)) }
func (m CIR_Bits) Clear()         { rcc(3).ClearBits(uint32(m)) }
func (m CIR_Bits) Load() uint32   { return rcc(3).Bits(uint32(m)) }
func (m CIR_Bits) Store(b uint32) { rcc(3).StoreBits(uint32(m), b) }
func (m CIR_Bits) LoadVal() int   { return rcc(3).Field(uint32(m)) }
func (m CIR_Bits) StoreVal(v int) { rcc(3).SetField(uint32(m), v) }

func CIR_Load() CIR_Bits   { return CIR_Bits(rcc(3).Load()) }
func CIR_Store(b CIR_Bits) { rcc(3).Store(uint32(b)) }

func (b CIR_Bits) Field(mask CIR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_CIR(v int, mask CIR_Bits) CIR_Bits {
	return CIR_Bits(bits.Make32(v, uint32(mask)))
}


type AHB1RSTR_Bits uint32

func (m AHB1RSTR_Bits) Set()           { rcc(4).SetBits(uint32(m)) }
func (m AHB1RSTR_Bits) Clear()         { rcc(4).ClearBits(uint32(m)) }
func (m AHB1RSTR_Bits) Load() uint32   { return rcc(4).Bits(uint32(m)) }
func (m AHB1RSTR_Bits) Store(b uint32) { rcc(4).StoreBits(uint32(m), b) }
func (m AHB1RSTR_Bits) LoadVal() int   { return rcc(4).Field(uint32(m)) }
func (m AHB1RSTR_Bits) StoreVal(v int) { rcc(4).SetField(uint32(m), v) }

func AHB1RSTR_Load() AHB1RSTR_Bits   { return AHB1RSTR_Bits(rcc(4).Load()) }
func AHB1RSTR_Store(b AHB1RSTR_Bits) { rcc(4).Store(uint32(b)) }

func (b AHB1RSTR_Bits) Field(mask AHB1RSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_AHB1RSTR(v int, mask AHB1RSTR_Bits) AHB1RSTR_Bits {
	return AHB1RSTR_Bits(bits.Make32(v, uint32(mask)))
}


type AHB2RSTR_Bits uint32

func (m AHB2RSTR_Bits) Set()           { rcc(5).SetBits(uint32(m)) }
func (m AHB2RSTR_Bits) Clear()         { rcc(5).ClearBits(uint32(m)) }
func (m AHB2RSTR_Bits) Load() uint32   { return rcc(5).Bits(uint32(m)) }
func (m AHB2RSTR_Bits) Store(b uint32) { rcc(5).StoreBits(uint32(m), b) }
func (m AHB2RSTR_Bits) LoadVal() int   { return rcc(5).Field(uint32(m)) }
func (m AHB2RSTR_Bits) StoreVal(v int) { rcc(5).SetField(uint32(m), v) }

func AHB2RSTR_Load() AHB2RSTR_Bits   { return AHB2RSTR_Bits(rcc(5).Load()) }
func AHB2RSTR_Store(b AHB2RSTR_Bits) { rcc(5).Store(uint32(b)) }

func (b AHB2RSTR_Bits) Field(mask AHB2RSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_AHB2RSTR(v int, mask AHB2RSTR_Bits) AHB2RSTR_Bits {
	return AHB2RSTR_Bits(bits.Make32(v, uint32(mask)))
}


type APB1RSTR_Bits uint32

func (m APB1RSTR_Bits) Set()           { rcc(8).SetBits(uint32(m)) }
func (m APB1RSTR_Bits) Clear()         { rcc(8).ClearBits(uint32(m)) }
func (m APB1RSTR_Bits) Load() uint32   { return rcc(8).Bits(uint32(m)) }
func (m APB1RSTR_Bits) Store(b uint32) { rcc(8).StoreBits(uint32(m), b) }
func (m APB1RSTR_Bits) LoadVal() int   { return rcc(8).Field(uint32(m)) }
func (m APB1RSTR_Bits) StoreVal(v int) { rcc(8).SetField(uint32(m), v) }

func APB1RSTR_Load() APB1RSTR_Bits   { return APB1RSTR_Bits(rcc(8).Load()) }
func APB1RSTR_Store(b APB1RSTR_Bits) { rcc(8).Store(uint32(b)) }

func (b APB1RSTR_Bits) Field(mask APB1RSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_APB1RSTR(v int, mask APB1RSTR_Bits) APB1RSTR_Bits {
	return APB1RSTR_Bits(bits.Make32(v, uint32(mask)))
}


type APB2RSTR_Bits uint32

func (m APB2RSTR_Bits) Set()           { rcc(9).SetBits(uint32(m)) }
func (m APB2RSTR_Bits) Clear()         { rcc(9).ClearBits(uint32(m)) }
func (m APB2RSTR_Bits) Load() uint32   { return rcc(9).Bits(uint32(m)) }
func (m APB2RSTR_Bits) Store(b uint32) { rcc(9).StoreBits(uint32(m), b) }
func (m APB2RSTR_Bits) LoadVal() int   { return rcc(9).Field(uint32(m)) }
func (m APB2RSTR_Bits) StoreVal(v int) { rcc(9).SetField(uint32(m), v) }

func APB2RSTR_Load() APB2RSTR_Bits   { return APB2RSTR_Bits(rcc(9).Load()) }
func APB2RSTR_Store(b APB2RSTR_Bits) { rcc(9).Store(uint32(b)) }

func (b APB2RSTR_Bits) Field(mask APB2RSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_APB2RSTR(v int, mask APB2RSTR_Bits) APB2RSTR_Bits {
	return APB2RSTR_Bits(bits.Make32(v, uint32(mask)))
}


type AHB1ENR_Bits uint32

func (m AHB1ENR_Bits) Set()           { rcc(12).SetBits(uint32(m)) }
func (m AHB1ENR_Bits) Clear()         { rcc(12).ClearBits(uint32(m)) }
func (m AHB1ENR_Bits) Load() uint32   { return rcc(12).Bits(uint32(m)) }
func (m AHB1ENR_Bits) Store(b uint32) { rcc(12).StoreBits(uint32(m), b) }
func (m AHB1ENR_Bits) LoadVal() int   { return rcc(12).Field(uint32(m)) }
func (m AHB1ENR_Bits) StoreVal(v int) { rcc(12).SetField(uint32(m), v) }

func AHB1ENR_Load() AHB1ENR_Bits   { return AHB1ENR_Bits(rcc(12).Load()) }
func AHB1ENR_Store(b AHB1ENR_Bits) { rcc(12).Store(uint32(b)) }

func (b AHB1ENR_Bits) Field(mask AHB1ENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_AHB1ENR(v int, mask AHB1ENR_Bits) AHB1ENR_Bits {
	return AHB1ENR_Bits(bits.Make32(v, uint32(mask)))
}


type AHB2ENR_Bits uint32

func (m AHB2ENR_Bits) Set()           { rcc(13).SetBits(uint32(m)) }
func (m AHB2ENR_Bits) Clear()         { rcc(13).ClearBits(uint32(m)) }
func (m AHB2ENR_Bits) Load() uint32   { return rcc(13).Bits(uint32(m)) }
func (m AHB2ENR_Bits) Store(b uint32) { rcc(13).StoreBits(uint32(m), b) }
func (m AHB2ENR_Bits) LoadVal() int   { return rcc(13).Field(uint32(m)) }
func (m AHB2ENR_Bits) StoreVal(v int) { rcc(13).SetField(uint32(m), v) }

func AHB2ENR_Load() AHB2ENR_Bits   { return AHB2ENR_Bits(rcc(13).Load()) }
func AHB2ENR_Store(b AHB2ENR_Bits) { rcc(13).Store(uint32(b)) }

func (b AHB2ENR_Bits) Field(mask AHB2ENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_AHB2ENR(v int, mask AHB2ENR_Bits) AHB2ENR_Bits {
	return AHB2ENR_Bits(bits.Make32(v, uint32(mask)))
}


type APB1ENR_Bits uint32

func (m APB1ENR_Bits) Set()           { rcc(16).SetBits(uint32(m)) }
func (m APB1ENR_Bits) Clear()         { rcc(16).ClearBits(uint32(m)) }
func (m APB1ENR_Bits) Load() uint32   { return rcc(16).Bits(uint32(m)) }
func (m APB1ENR_Bits) Store(b uint32) { rcc(16).StoreBits(uint32(m), b) }
func (m APB1ENR_Bits) LoadVal() int   { return rcc(16).Field(uint32(m)) }
func (m APB1ENR_Bits) StoreVal(v int) { rcc(16).SetField(uint32(m), v) }

func APB1ENR_Load() APB1ENR_Bits   { return APB1ENR_Bits(rcc(16).Load()) }
func APB1ENR_Store(b APB1ENR_Bits) { rcc(16).Store(uint32(b)) }

func (b APB1ENR_Bits) Field(mask APB1ENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_APB1ENR(v int, mask APB1ENR_Bits) APB1ENR_Bits {
	return APB1ENR_Bits(bits.Make32(v, uint32(mask)))
}


type APB2ENR_Bits uint32

func (m APB2ENR_Bits) Set()           { rcc(17).SetBits(uint32(m)) }
func (m APB2ENR_Bits) Clear()         { rcc(17).ClearBits(uint32(m)) }
func (m APB2ENR_Bits) Load() uint32   { return rcc(17).Bits(uint32(m)) }
func (m APB2ENR_Bits) Store(b uint32) { rcc(17).StoreBits(uint32(m), b) }
func (m APB2ENR_Bits) LoadVal() int   { return rcc(17).Field(uint32(m)) }
func (m APB2ENR_Bits) StoreVal(v int) { rcc(17).SetField(uint32(m), v) }

func APB2ENR_Load() APB2ENR_Bits   { return APB2ENR_Bits(rcc(17).Load()) }
func APB2ENR_Store(b APB2ENR_Bits) { rcc(17).Store(uint32(b)) }

func (b APB2ENR_Bits) Field(mask APB2ENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_APB2ENR(v int, mask APB2ENR_Bits) APB2ENR_Bits {
	return APB2ENR_Bits(bits.Make32(v, uint32(mask)))
}


type AHB1LPENR_Bits uint32

func (m AHB1LPENR_Bits) Set()           { rcc(20).SetBits(uint32(m)) }
func (m AHB1LPENR_Bits) Clear()         { rcc(20).ClearBits(uint32(m)) }
func (m AHB1LPENR_Bits) Load() uint32   { return rcc(20).Bits(uint32(m)) }
func (m AHB1LPENR_Bits) Store(b uint32) { rcc(20).StoreBits(uint32(m), b) }
func (m AHB1LPENR_Bits) LoadVal() int   { return rcc(20).Field(uint32(m)) }
func (m AHB1LPENR_Bits) StoreVal(v int) { rcc(20).SetField(uint32(m), v) }

func AHB1LPENR_Load() AHB1LPENR_Bits   { return AHB1LPENR_Bits(rcc(20).Load()) }
func AHB1LPENR_Store(b AHB1LPENR_Bits) { rcc(20).Store(uint32(b)) }

func (b AHB1LPENR_Bits) Field(mask AHB1LPENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_AHB1LPENR(v int, mask AHB1LPENR_Bits) AHB1LPENR_Bits {
	return AHB1LPENR_Bits(bits.Make32(v, uint32(mask)))
}


type AHB2LPENR_Bits uint32

func (m AHB2LPENR_Bits) Set()           { rcc(21).SetBits(uint32(m)) }
func (m AHB2LPENR_Bits) Clear()         { rcc(21).ClearBits(uint32(m)) }
func (m AHB2LPENR_Bits) Load() uint32   { return rcc(21).Bits(uint32(m)) }
func (m AHB2LPENR_Bits) Store(b uint32) { rcc(21).StoreBits(uint32(m), b) }
func (m AHB2LPENR_Bits) LoadVal() int   { return rcc(21).Field(uint32(m)) }
func (m AHB2LPENR_Bits) StoreVal(v int) { rcc(21).SetField(uint32(m), v) }

func AHB2LPENR_Load() AHB2LPENR_Bits   { return AHB2LPENR_Bits(rcc(21).Load()) }
func AHB2LPENR_Store(b AHB2LPENR_Bits) { rcc(21).Store(uint32(b)) }

func (b AHB2LPENR_Bits) Field(mask AHB2LPENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_AHB2LPENR(v int, mask AHB2LPENR_Bits) AHB2LPENR_Bits {
	return AHB2LPENR_Bits(bits.Make32(v, uint32(mask)))
}


type APB1LPENR_Bits uint32

func (m APB1LPENR_Bits) Set()           { rcc(24).SetBits(uint32(m)) }
func (m APB1LPENR_Bits) Clear()         { rcc(24).ClearBits(uint32(m)) }
func (m APB1LPENR_Bits) Load() uint32   { return rcc(24).Bits(uint32(m)) }
func (m APB1LPENR_Bits) Store(b uint32) { rcc(24).StoreBits(uint32(m), b) }
func (m APB1LPENR_Bits) LoadVal() int   { return rcc(24).Field(uint32(m)) }
func (m APB1LPENR_Bits) StoreVal(v int) { rcc(24).SetField(uint32(m), v) }

func APB1LPENR_Load() APB1LPENR_Bits   { return APB1LPENR_Bits(rcc(24).Load()) }
func APB1LPENR_Store(b APB1LPENR_Bits) { rcc(24).Store(uint32(b)) }

func (b APB1LPENR_Bits) Field(mask APB1LPENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_APB1LPENR(v int, mask APB1LPENR_Bits) APB1LPENR_Bits {
	return APB1LPENR_Bits(bits.Make32(v, uint32(mask)))
}


type APB2LPENR_Bits uint32

func (m APB2LPENR_Bits) Set()           { rcc(25).SetBits(uint32(m)) }
func (m APB2LPENR_Bits) Clear()         { rcc(25).ClearBits(uint32(m)) }
func (m APB2LPENR_Bits) Load() uint32   { return rcc(25).Bits(uint32(m)) }
func (m APB2LPENR_Bits) Store(b uint32) { rcc(25).StoreBits(uint32(m), b) }
func (m APB2LPENR_Bits) LoadVal() int   { return rcc(25).Field(uint32(m)) }
func (m APB2LPENR_Bits) StoreVal(v int) { rcc(25).SetField(uint32(m), v) }

func APB2LPENR_Load() APB2LPENR_Bits   { return APB2LPENR_Bits(rcc(25).Load()) }
func APB2LPENR_Store(b APB2LPENR_Bits) { rcc(25).Store(uint32(b)) }

func (b APB2LPENR_Bits) Field(mask APB2LPENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_APB2LPENR(v int, mask APB2LPENR_Bits) APB2LPENR_Bits {
	return APB2LPENR_Bits(bits.Make32(v, uint32(mask)))
}


type BDCR_Bits uint32

func (m BDCR_Bits) Set()           { rcc(28).SetBits(uint32(m)) }
func (m BDCR_Bits) Clear()         { rcc(28).ClearBits(uint32(m)) }
func (m BDCR_Bits) Load() uint32   { return rcc(28).Bits(uint32(m)) }
func (m BDCR_Bits) Store(b uint32) { rcc(28).StoreBits(uint32(m), b) }
func (m BDCR_Bits) LoadVal() int   { return rcc(28).Field(uint32(m)) }
func (m BDCR_Bits) StoreVal(v int) { rcc(28).SetField(uint32(m), v) }

func BDCR_Load() BDCR_Bits   { return BDCR_Bits(rcc(28).Load()) }
func BDCR_Store(b BDCR_Bits) { rcc(28).Store(uint32(b)) }

func (b BDCR_Bits) Field(mask BDCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_BDCR(v int, mask BDCR_Bits) BDCR_Bits {
	return BDCR_Bits(bits.Make32(v, uint32(mask)))
}


type CSR_Bits uint32

func (m CSR_Bits) Set()           { rcc(29).SetBits(uint32(m)) }
func (m CSR_Bits) Clear()         { rcc(29).ClearBits(uint32(m)) }
func (m CSR_Bits) Load() uint32   { return rcc(29).Bits(uint32(m)) }
func (m CSR_Bits) Store(b uint32) { rcc(29).StoreBits(uint32(m), b) }
func (m CSR_Bits) LoadVal() int   { return rcc(29).Field(uint32(m)) }
func (m CSR_Bits) StoreVal(v int) { rcc(29).SetField(uint32(m), v) }

func CSR_Load() CSR_Bits   { return CSR_Bits(rcc(29).Load()) }
func CSR_Store(b CSR_Bits) { rcc(29).Store(uint32(b)) }

func (b CSR_Bits) Field(mask CSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_CSR(v int, mask CSR_Bits) CSR_Bits {
	return CSR_Bits(bits.Make32(v, uint32(mask)))
}


type SSCGR_Bits uint32

func (m SSCGR_Bits) Set()           { rcc(32).SetBits(uint32(m)) }
func (m SSCGR_Bits) Clear()         { rcc(32).ClearBits(uint32(m)) }
func (m SSCGR_Bits) Load() uint32   { return rcc(32).Bits(uint32(m)) }
func (m SSCGR_Bits) Store(b uint32) { rcc(32).StoreBits(uint32(m), b) }
func (m SSCGR_Bits) LoadVal() int   { return rcc(32).Field(uint32(m)) }
func (m SSCGR_Bits) StoreVal(v int) { rcc(32).SetField(uint32(m), v) }

func SSCGR_Load() SSCGR_Bits   { return SSCGR_Bits(rcc(32).Load()) }
func SSCGR_Store(b SSCGR_Bits) { rcc(32).Store(uint32(b)) }

func (b SSCGR_Bits) Field(mask SSCGR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_SSCGR(v int, mask SSCGR_Bits) SSCGR_Bits {
	return SSCGR_Bits(bits.Make32(v, uint32(mask)))
}


type PLLI2SCFGR_Bits uint32

func (m PLLI2SCFGR_Bits) Set()           { rcc(33).SetBits(uint32(m)) }
func (m PLLI2SCFGR_Bits) Clear()         { rcc(33).ClearBits(uint32(m)) }
func (m PLLI2SCFGR_Bits) Load() uint32   { return rcc(33).Bits(uint32(m)) }
func (m PLLI2SCFGR_Bits) Store(b uint32) { rcc(33).StoreBits(uint32(m), b) }
func (m PLLI2SCFGR_Bits) LoadVal() int   { return rcc(33).Field(uint32(m)) }
func (m PLLI2SCFGR_Bits) StoreVal(v int) { rcc(33).SetField(uint32(m), v) }

func PLLI2SCFGR_Load() PLLI2SCFGR_Bits   { return PLLI2SCFGR_Bits(rcc(33).Load()) }
func PLLI2SCFGR_Store(b PLLI2SCFGR_Bits) { rcc(33).Store(uint32(b)) }

func (b PLLI2SCFGR_Bits) Field(mask PLLI2SCFGR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_PLLI2SCFGR(v int, mask PLLI2SCFGR_Bits) PLLI2SCFGR_Bits {
	return PLLI2SCFGR_Bits(bits.Make32(v, uint32(mask)))
}


type RCC_DCKCFGR_Bits uint32

func (m RCC_DCKCFGR_Bits) Set()           { rcc(35).SetBits(uint32(m)) }
func (m RCC_DCKCFGR_Bits) Clear()         { rcc(35).ClearBits(uint32(m)) }
func (m RCC_DCKCFGR_Bits) Load() uint32   { return rcc(35).Bits(uint32(m)) }
func (m RCC_DCKCFGR_Bits) Store(b uint32) { rcc(35).StoreBits(uint32(m), b) }
func (m RCC_DCKCFGR_Bits) LoadVal() int   { return rcc(35).Field(uint32(m)) }
func (m RCC_DCKCFGR_Bits) StoreVal(v int) { rcc(35).SetField(uint32(m), v) }

func RCC_DCKCFGR_Load() RCC_DCKCFGR_Bits   { return RCC_DCKCFGR_Bits(rcc(35).Load()) }
func RCC_DCKCFGR_Store(b RCC_DCKCFGR_Bits) { rcc(35).Store(uint32(b)) }

func (b RCC_DCKCFGR_Bits) Field(mask RCC_DCKCFGR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_RCC_DCKCFGR(v int, mask RCC_DCKCFGR_Bits) RCC_DCKCFGR_Bits {
	return RCC_DCKCFGR_Bits(bits.Make32(v, uint32(mask)))
}
