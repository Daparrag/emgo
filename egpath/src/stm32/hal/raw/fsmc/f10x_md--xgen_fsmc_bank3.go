// +build f10x_md

package fsmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_md/mmap"
)

type FSMC_Bank3_Periph struct {
	PCR3  PCR3
	SR3   SR3
	PMEM3 PMEM3
	PATT3 PATT3
	_     uint32
	ECCR3 ECCR3
}

func (p *FSMC_Bank3_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FSMC_Bank3 = (*FSMC_Bank3_Periph)(unsafe.Pointer(uintptr(mmap.FSMC_Bank3_R_BASE)))

type PCR3_Bits uint32

func (b PCR3_Bits) Field(mask PCR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PCR3_Bits) J(v int) PCR3_Bits {
	return PCR3_Bits(bits.Make32(v, uint32(mask)))
}

type PCR3 struct{ mmio.U32 }

func (r *PCR3) Bits(mask PCR3_Bits) PCR3_Bits { return PCR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *PCR3) StoreBits(mask, b PCR3_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PCR3) SetBits(mask PCR3_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *PCR3) ClearBits(mask PCR3_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *PCR3) Load() PCR3_Bits               { return PCR3_Bits(r.U32.Load()) }
func (r *PCR3) Store(b PCR3_Bits)             { r.U32.Store(uint32(b)) }

func (r *PCR3) AtomicSetBits(mask PCR3_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PCR3) AtomicClearBits(mask PCR3_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type PCR3_Mask struct{ mmio.UM32 }

func (rm PCR3_Mask) Load() PCR3_Bits   { return PCR3_Bits(rm.UM32.Load()) }
func (rm PCR3_Mask) Store(b PCR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank3_Periph) PWAITEN() PCR3_Mask {
	return PCR3_Mask{mmio.UM32{&p.PCR3.U32, uint32(PWAITEN)}}
}

func (p *FSMC_Bank3_Periph) PBKEN() PCR3_Mask {
	return PCR3_Mask{mmio.UM32{&p.PCR3.U32, uint32(PBKEN)}}
}

func (p *FSMC_Bank3_Periph) PTYP() PCR3_Mask {
	return PCR3_Mask{mmio.UM32{&p.PCR3.U32, uint32(PTYP)}}
}

func (p *FSMC_Bank3_Periph) PWID() PCR3_Mask {
	return PCR3_Mask{mmio.UM32{&p.PCR3.U32, uint32(PWID)}}
}

func (p *FSMC_Bank3_Periph) ECCEN() PCR3_Mask {
	return PCR3_Mask{mmio.UM32{&p.PCR3.U32, uint32(ECCEN)}}
}

func (p *FSMC_Bank3_Periph) TCLR() PCR3_Mask {
	return PCR3_Mask{mmio.UM32{&p.PCR3.U32, uint32(TCLR)}}
}

func (p *FSMC_Bank3_Periph) TAR() PCR3_Mask {
	return PCR3_Mask{mmio.UM32{&p.PCR3.U32, uint32(TAR)}}
}

func (p *FSMC_Bank3_Periph) ECCPS() PCR3_Mask {
	return PCR3_Mask{mmio.UM32{&p.PCR3.U32, uint32(ECCPS)}}
}

type SR3_Bits uint32

func (b SR3_Bits) Field(mask SR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR3_Bits) J(v int) SR3_Bits {
	return SR3_Bits(bits.Make32(v, uint32(mask)))
}

type SR3 struct{ mmio.U32 }

func (r *SR3) Bits(mask SR3_Bits) SR3_Bits { return SR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR3) StoreBits(mask, b SR3_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR3) SetBits(mask SR3_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *SR3) ClearBits(mask SR3_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *SR3) Load() SR3_Bits              { return SR3_Bits(r.U32.Load()) }
func (r *SR3) Store(b SR3_Bits)            { r.U32.Store(uint32(b)) }

func (r *SR3) AtomicSetBits(mask SR3_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *SR3) AtomicClearBits(mask SR3_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type SR3_Mask struct{ mmio.UM32 }

func (rm SR3_Mask) Load() SR3_Bits   { return SR3_Bits(rm.UM32.Load()) }
func (rm SR3_Mask) Store(b SR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank3_Periph) IRS() SR3_Mask {
	return SR3_Mask{mmio.UM32{&p.SR3.U32, uint32(IRS)}}
}

func (p *FSMC_Bank3_Periph) ILS() SR3_Mask {
	return SR3_Mask{mmio.UM32{&p.SR3.U32, uint32(ILS)}}
}

func (p *FSMC_Bank3_Periph) IFS() SR3_Mask {
	return SR3_Mask{mmio.UM32{&p.SR3.U32, uint32(IFS)}}
}

func (p *FSMC_Bank3_Periph) IREN() SR3_Mask {
	return SR3_Mask{mmio.UM32{&p.SR3.U32, uint32(IREN)}}
}

func (p *FSMC_Bank3_Periph) ILEN() SR3_Mask {
	return SR3_Mask{mmio.UM32{&p.SR3.U32, uint32(ILEN)}}
}

func (p *FSMC_Bank3_Periph) IFEN() SR3_Mask {
	return SR3_Mask{mmio.UM32{&p.SR3.U32, uint32(IFEN)}}
}

func (p *FSMC_Bank3_Periph) FEMPT() SR3_Mask {
	return SR3_Mask{mmio.UM32{&p.SR3.U32, uint32(FEMPT)}}
}

type PMEM3_Bits uint32

func (b PMEM3_Bits) Field(mask PMEM3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PMEM3_Bits) J(v int) PMEM3_Bits {
	return PMEM3_Bits(bits.Make32(v, uint32(mask)))
}

type PMEM3 struct{ mmio.U32 }

func (r *PMEM3) Bits(mask PMEM3_Bits) PMEM3_Bits { return PMEM3_Bits(r.U32.Bits(uint32(mask))) }
func (r *PMEM3) StoreBits(mask, b PMEM3_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PMEM3) SetBits(mask PMEM3_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PMEM3) ClearBits(mask PMEM3_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PMEM3) Load() PMEM3_Bits                { return PMEM3_Bits(r.U32.Load()) }
func (r *PMEM3) Store(b PMEM3_Bits)              { r.U32.Store(uint32(b)) }

func (r *PMEM3) AtomicSetBits(mask PMEM3_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PMEM3) AtomicClearBits(mask PMEM3_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type PMEM3_Mask struct{ mmio.UM32 }

func (rm PMEM3_Mask) Load() PMEM3_Bits   { return PMEM3_Bits(rm.UM32.Load()) }
func (rm PMEM3_Mask) Store(b PMEM3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank3_Periph) MEMSET3() PMEM3_Mask {
	return PMEM3_Mask{mmio.UM32{&p.PMEM3.U32, uint32(MEMSET3)}}
}

func (p *FSMC_Bank3_Periph) MEMWAIT3() PMEM3_Mask {
	return PMEM3_Mask{mmio.UM32{&p.PMEM3.U32, uint32(MEMWAIT3)}}
}

func (p *FSMC_Bank3_Periph) MEMHOLD3() PMEM3_Mask {
	return PMEM3_Mask{mmio.UM32{&p.PMEM3.U32, uint32(MEMHOLD3)}}
}

func (p *FSMC_Bank3_Periph) MEMHIZ3() PMEM3_Mask {
	return PMEM3_Mask{mmio.UM32{&p.PMEM3.U32, uint32(MEMHIZ3)}}
}

type PATT3_Bits uint32

func (b PATT3_Bits) Field(mask PATT3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PATT3_Bits) J(v int) PATT3_Bits {
	return PATT3_Bits(bits.Make32(v, uint32(mask)))
}

type PATT3 struct{ mmio.U32 }

func (r *PATT3) Bits(mask PATT3_Bits) PATT3_Bits { return PATT3_Bits(r.U32.Bits(uint32(mask))) }
func (r *PATT3) StoreBits(mask, b PATT3_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PATT3) SetBits(mask PATT3_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PATT3) ClearBits(mask PATT3_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PATT3) Load() PATT3_Bits                { return PATT3_Bits(r.U32.Load()) }
func (r *PATT3) Store(b PATT3_Bits)              { r.U32.Store(uint32(b)) }

func (r *PATT3) AtomicSetBits(mask PATT3_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PATT3) AtomicClearBits(mask PATT3_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type PATT3_Mask struct{ mmio.UM32 }

func (rm PATT3_Mask) Load() PATT3_Bits   { return PATT3_Bits(rm.UM32.Load()) }
func (rm PATT3_Mask) Store(b PATT3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank3_Periph) ATTSET3() PATT3_Mask {
	return PATT3_Mask{mmio.UM32{&p.PATT3.U32, uint32(ATTSET3)}}
}

func (p *FSMC_Bank3_Periph) ATTWAIT3() PATT3_Mask {
	return PATT3_Mask{mmio.UM32{&p.PATT3.U32, uint32(ATTWAIT3)}}
}

func (p *FSMC_Bank3_Periph) ATTHOLD3() PATT3_Mask {
	return PATT3_Mask{mmio.UM32{&p.PATT3.U32, uint32(ATTHOLD3)}}
}

func (p *FSMC_Bank3_Periph) ATTHIZ3() PATT3_Mask {
	return PATT3_Mask{mmio.UM32{&p.PATT3.U32, uint32(ATTHIZ3)}}
}

type ECCR3_Bits uint32

func (b ECCR3_Bits) Field(mask ECCR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ECCR3_Bits) J(v int) ECCR3_Bits {
	return ECCR3_Bits(bits.Make32(v, uint32(mask)))
}

type ECCR3 struct{ mmio.U32 }

func (r *ECCR3) Bits(mask ECCR3_Bits) ECCR3_Bits { return ECCR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *ECCR3) StoreBits(mask, b ECCR3_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ECCR3) SetBits(mask ECCR3_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *ECCR3) ClearBits(mask ECCR3_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *ECCR3) Load() ECCR3_Bits                { return ECCR3_Bits(r.U32.Load()) }
func (r *ECCR3) Store(b ECCR3_Bits)              { r.U32.Store(uint32(b)) }

func (r *ECCR3) AtomicSetBits(mask ECCR3_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ECCR3) AtomicClearBits(mask ECCR3_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type ECCR3_Mask struct{ mmio.UM32 }

func (rm ECCR3_Mask) Load() ECCR3_Bits   { return ECCR3_Bits(rm.UM32.Load()) }
func (rm ECCR3_Mask) Store(b ECCR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank3_Periph) ECC3() ECCR3_Mask {
	return ECCR3_Mask{mmio.UM32{&p.ECCR3.U32, uint32(ECC3)}}
}
