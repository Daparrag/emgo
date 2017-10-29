// +build f10x_md

package gpio

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_md/mmap"
)

type GPIO_Periph struct {
	CRL  CRL
	CRH  CRH
	IDR  IDR
	ODR  ODR
	BSRR BSRR
	BRR  BRR
	LCKR LCKR
}

func (p *GPIO_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var GPIOA = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOA_BASE)))

//emgo:const
var GPIOB = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOB_BASE)))

//emgo:const
var GPIOC = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOC_BASE)))

//emgo:const
var GPIOD = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOD_BASE)))

//emgo:const
var GPIOE = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOE_BASE)))

//emgo:const
var GPIOF = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOF_BASE)))

//emgo:const
var GPIOG = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOG_BASE)))

type CRL_Bits uint32

func (b CRL_Bits) Field(mask CRL_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CRL_Bits) J(v int) CRL_Bits {
	return CRL_Bits(bits.Make32(v, uint32(mask)))
}

type CRL struct{ mmio.U32 }

func (r *CRL) Bits(mask CRL_Bits) CRL_Bits { return CRL_Bits(r.U32.Bits(uint32(mask))) }
func (r *CRL) StoreBits(mask, b CRL_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CRL) SetBits(mask CRL_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CRL) ClearBits(mask CRL_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CRL) Load() CRL_Bits              { return CRL_Bits(r.U32.Load()) }
func (r *CRL) Store(b CRL_Bits)            { r.U32.Store(uint32(b)) }

func (r *CRL) AtomicSetBits(mask CRL_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CRL) AtomicClearBits(mask CRL_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CRL_Mask struct{ mmio.UM32 }

func (rm CRL_Mask) Load() CRL_Bits   { return CRL_Bits(rm.UM32.Load()) }
func (rm CRL_Mask) Store(b CRL_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) MODE() CRL_Mask {
	return CRL_Mask{mmio.UM32{&p.CRL.U32, uint32(MODE)}}
}

func (p *GPIO_Periph) CNF() CRL_Mask {
	return CRL_Mask{mmio.UM32{&p.CRL.U32, uint32(CNF)}}
}

type CRH_Bits uint32

func (b CRH_Bits) Field(mask CRH_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CRH_Bits) J(v int) CRH_Bits {
	return CRH_Bits(bits.Make32(v, uint32(mask)))
}

type CRH struct{ mmio.U32 }

func (r *CRH) Bits(mask CRH_Bits) CRH_Bits { return CRH_Bits(r.U32.Bits(uint32(mask))) }
func (r *CRH) StoreBits(mask, b CRH_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CRH) SetBits(mask CRH_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CRH) ClearBits(mask CRH_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CRH) Load() CRH_Bits              { return CRH_Bits(r.U32.Load()) }
func (r *CRH) Store(b CRH_Bits)            { r.U32.Store(uint32(b)) }

func (r *CRH) AtomicSetBits(mask CRH_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CRH) AtomicClearBits(mask CRH_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CRH_Mask struct{ mmio.UM32 }

func (rm CRH_Mask) Load() CRH_Bits   { return CRH_Bits(rm.UM32.Load()) }
func (rm CRH_Mask) Store(b CRH_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) MODE() CRH_Mask {
	return CRH_Mask{mmio.UM32{&p.CRH.U32, uint32(MODE)}}
}

func (p *GPIO_Periph) CNF() CRH_Mask {
	return CRH_Mask{mmio.UM32{&p.CRH.U32, uint32(CNF)}}
}

type IDR_Bits uint32

func (b IDR_Bits) Field(mask IDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IDR_Bits) J(v int) IDR_Bits {
	return IDR_Bits(bits.Make32(v, uint32(mask)))
}

type IDR struct{ mmio.U32 }

func (r *IDR) Bits(mask IDR_Bits) IDR_Bits { return IDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IDR) StoreBits(mask, b IDR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IDR) SetBits(mask IDR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *IDR) ClearBits(mask IDR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *IDR) Load() IDR_Bits              { return IDR_Bits(r.U32.Load()) }
func (r *IDR) Store(b IDR_Bits)            { r.U32.Store(uint32(b)) }

func (r *IDR) AtomicSetBits(mask IDR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *IDR) AtomicClearBits(mask IDR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type IDR_Mask struct{ mmio.UM32 }

func (rm IDR_Mask) Load() IDR_Bits   { return IDR_Bits(rm.UM32.Load()) }
func (rm IDR_Mask) Store(b IDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) IDR0() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR0)}}
}

func (p *GPIO_Periph) IDR1() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR1)}}
}

func (p *GPIO_Periph) IDR2() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR2)}}
}

func (p *GPIO_Periph) IDR3() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR3)}}
}

func (p *GPIO_Periph) IDR4() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR4)}}
}

func (p *GPIO_Periph) IDR5() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR5)}}
}

func (p *GPIO_Periph) IDR6() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR6)}}
}

func (p *GPIO_Periph) IDR7() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR7)}}
}

func (p *GPIO_Periph) IDR8() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR8)}}
}

func (p *GPIO_Periph) IDR9() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR9)}}
}

func (p *GPIO_Periph) IDR10() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR10)}}
}

func (p *GPIO_Periph) IDR11() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR11)}}
}

func (p *GPIO_Periph) IDR12() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR12)}}
}

func (p *GPIO_Periph) IDR13() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR13)}}
}

func (p *GPIO_Periph) IDR14() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR14)}}
}

func (p *GPIO_Periph) IDR15() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR15)}}
}

type ODR_Bits uint32

func (b ODR_Bits) Field(mask ODR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ODR_Bits) J(v int) ODR_Bits {
	return ODR_Bits(bits.Make32(v, uint32(mask)))
}

type ODR struct{ mmio.U32 }

func (r *ODR) Bits(mask ODR_Bits) ODR_Bits { return ODR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ODR) StoreBits(mask, b ODR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ODR) SetBits(mask ODR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ODR) ClearBits(mask ODR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ODR) Load() ODR_Bits              { return ODR_Bits(r.U32.Load()) }
func (r *ODR) Store(b ODR_Bits)            { r.U32.Store(uint32(b)) }

func (r *ODR) AtomicSetBits(mask ODR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ODR) AtomicClearBits(mask ODR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type ODR_Mask struct{ mmio.UM32 }

func (rm ODR_Mask) Load() ODR_Bits   { return ODR_Bits(rm.UM32.Load()) }
func (rm ODR_Mask) Store(b ODR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) ODR0() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR0)}}
}

func (p *GPIO_Periph) ODR1() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR1)}}
}

func (p *GPIO_Periph) ODR2() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR2)}}
}

func (p *GPIO_Periph) ODR3() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR3)}}
}

func (p *GPIO_Periph) ODR4() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR4)}}
}

func (p *GPIO_Periph) ODR5() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR5)}}
}

func (p *GPIO_Periph) ODR6() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR6)}}
}

func (p *GPIO_Periph) ODR7() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR7)}}
}

func (p *GPIO_Periph) ODR8() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR8)}}
}

func (p *GPIO_Periph) ODR9() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR9)}}
}

func (p *GPIO_Periph) ODR10() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR10)}}
}

func (p *GPIO_Periph) ODR11() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR11)}}
}

func (p *GPIO_Periph) ODR12() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR12)}}
}

func (p *GPIO_Periph) ODR13() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR13)}}
}

func (p *GPIO_Periph) ODR14() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR14)}}
}

func (p *GPIO_Periph) ODR15() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR15)}}
}

type BSRR_Bits uint32

func (b BSRR_Bits) Field(mask BSRR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BSRR_Bits) J(v int) BSRR_Bits {
	return BSRR_Bits(bits.Make32(v, uint32(mask)))
}

type BSRR struct{ mmio.U32 }

func (r *BSRR) Bits(mask BSRR_Bits) BSRR_Bits { return BSRR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BSRR) StoreBits(mask, b BSRR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BSRR) SetBits(mask BSRR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *BSRR) ClearBits(mask BSRR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *BSRR) Load() BSRR_Bits               { return BSRR_Bits(r.U32.Load()) }
func (r *BSRR) Store(b BSRR_Bits)             { r.U32.Store(uint32(b)) }

func (r *BSRR) AtomicSetBits(mask BSRR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *BSRR) AtomicClearBits(mask BSRR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type BSRR_Mask struct{ mmio.UM32 }

func (rm BSRR_Mask) Load() BSRR_Bits   { return BSRR_Bits(rm.UM32.Load()) }
func (rm BSRR_Mask) Store(b BSRR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) BS0() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS0)}}
}

func (p *GPIO_Periph) BS1() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS1)}}
}

func (p *GPIO_Periph) BS2() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS2)}}
}

func (p *GPIO_Periph) BS3() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS3)}}
}

func (p *GPIO_Periph) BS4() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS4)}}
}

func (p *GPIO_Periph) BS5() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS5)}}
}

func (p *GPIO_Periph) BS6() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS6)}}
}

func (p *GPIO_Periph) BS7() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS7)}}
}

func (p *GPIO_Periph) BS8() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS8)}}
}

func (p *GPIO_Periph) BS9() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS9)}}
}

func (p *GPIO_Periph) BS10() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS10)}}
}

func (p *GPIO_Periph) BS11() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS11)}}
}

func (p *GPIO_Periph) BS12() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS12)}}
}

func (p *GPIO_Periph) BS13() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS13)}}
}

func (p *GPIO_Periph) BS14() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS14)}}
}

func (p *GPIO_Periph) BS15() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS15)}}
}

func (p *GPIO_Periph) BR0() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR0)}}
}

func (p *GPIO_Periph) BR1() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR1)}}
}

func (p *GPIO_Periph) BR2() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR2)}}
}

func (p *GPIO_Periph) BR3() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR3)}}
}

func (p *GPIO_Periph) BR4() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR4)}}
}

func (p *GPIO_Periph) BR5() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR5)}}
}

func (p *GPIO_Periph) BR6() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR6)}}
}

func (p *GPIO_Periph) BR7() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR7)}}
}

func (p *GPIO_Periph) BR8() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR8)}}
}

func (p *GPIO_Periph) BR9() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR9)}}
}

func (p *GPIO_Periph) BR10() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR10)}}
}

func (p *GPIO_Periph) BR11() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR11)}}
}

func (p *GPIO_Periph) BR12() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR12)}}
}

func (p *GPIO_Periph) BR13() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR13)}}
}

func (p *GPIO_Periph) BR14() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR14)}}
}

func (p *GPIO_Periph) BR15() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR15)}}
}

type BRR_Bits uint32

func (b BRR_Bits) Field(mask BRR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BRR_Bits) J(v int) BRR_Bits {
	return BRR_Bits(bits.Make32(v, uint32(mask)))
}

type BRR struct{ mmio.U32 }

func (r *BRR) Bits(mask BRR_Bits) BRR_Bits { return BRR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BRR) StoreBits(mask, b BRR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BRR) SetBits(mask BRR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *BRR) ClearBits(mask BRR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *BRR) Load() BRR_Bits              { return BRR_Bits(r.U32.Load()) }
func (r *BRR) Store(b BRR_Bits)            { r.U32.Store(uint32(b)) }

func (r *BRR) AtomicSetBits(mask BRR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *BRR) AtomicClearBits(mask BRR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type BRR_Mask struct{ mmio.UM32 }

func (rm BRR_Mask) Load() BRR_Bits   { return BRR_Bits(rm.UM32.Load()) }
func (rm BRR_Mask) Store(b BRR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) BR0() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR0)}}
}

func (p *GPIO_Periph) BR1() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR1)}}
}

func (p *GPIO_Periph) BR2() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR2)}}
}

func (p *GPIO_Periph) BR3() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR3)}}
}

func (p *GPIO_Periph) BR4() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR4)}}
}

func (p *GPIO_Periph) BR5() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR5)}}
}

func (p *GPIO_Periph) BR6() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR6)}}
}

func (p *GPIO_Periph) BR7() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR7)}}
}

func (p *GPIO_Periph) BR8() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR8)}}
}

func (p *GPIO_Periph) BR9() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR9)}}
}

func (p *GPIO_Periph) BR10() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR10)}}
}

func (p *GPIO_Periph) BR11() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR11)}}
}

func (p *GPIO_Periph) BR12() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR12)}}
}

func (p *GPIO_Periph) BR13() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR13)}}
}

func (p *GPIO_Periph) BR14() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR14)}}
}

func (p *GPIO_Periph) BR15() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR15)}}
}

type LCKR_Bits uint32

func (b LCKR_Bits) Field(mask LCKR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LCKR_Bits) J(v int) LCKR_Bits {
	return LCKR_Bits(bits.Make32(v, uint32(mask)))
}

type LCKR struct{ mmio.U32 }

func (r *LCKR) Bits(mask LCKR_Bits) LCKR_Bits { return LCKR_Bits(r.U32.Bits(uint32(mask))) }
func (r *LCKR) StoreBits(mask, b LCKR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *LCKR) SetBits(mask LCKR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *LCKR) ClearBits(mask LCKR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *LCKR) Load() LCKR_Bits               { return LCKR_Bits(r.U32.Load()) }
func (r *LCKR) Store(b LCKR_Bits)             { r.U32.Store(uint32(b)) }

func (r *LCKR) AtomicSetBits(mask LCKR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *LCKR) AtomicClearBits(mask LCKR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type LCKR_Mask struct{ mmio.UM32 }

func (rm LCKR_Mask) Load() LCKR_Bits   { return LCKR_Bits(rm.UM32.Load()) }
func (rm LCKR_Mask) Store(b LCKR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) LCK0() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK0)}}
}

func (p *GPIO_Periph) LCK1() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK1)}}
}

func (p *GPIO_Periph) LCK2() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK2)}}
}

func (p *GPIO_Periph) LCK3() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK3)}}
}

func (p *GPIO_Periph) LCK4() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK4)}}
}

func (p *GPIO_Periph) LCK5() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK5)}}
}

func (p *GPIO_Periph) LCK6() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK6)}}
}

func (p *GPIO_Periph) LCK7() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK7)}}
}

func (p *GPIO_Periph) LCK8() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK8)}}
}

func (p *GPIO_Periph) LCK9() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK9)}}
}

func (p *GPIO_Periph) LCK10() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK10)}}
}

func (p *GPIO_Periph) LCK11() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK11)}}
}

func (p *GPIO_Periph) LCK12() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK12)}}
}

func (p *GPIO_Periph) LCK13() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK13)}}
}

func (p *GPIO_Periph) LCK14() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK14)}}
}

func (p *GPIO_Periph) LCK15() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK15)}}
}

func (p *GPIO_Periph) LCKK() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCKK)}}
}
