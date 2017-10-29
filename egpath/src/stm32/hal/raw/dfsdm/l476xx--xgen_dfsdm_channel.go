// +build l476xx

package dfsdm

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type DFSDM_Channel_Periph struct {
	CHCFGR1  CHCFGR1
	CHCFGR2  CHCFGR2
	CHAWSCDR CHAWSCDR
	CHWDATAR CHWDATAR
	CHDATINR CHDATINR
}

func (p *DFSDM_Channel_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var DFSDM1_Channel0 = (*DFSDM_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DFSDM1_Channel0_BASE)))

//emgo:const
var DFSDM1_Channel1 = (*DFSDM_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DFSDM1_Channel1_BASE)))

//emgo:const
var DFSDM1_Channel2 = (*DFSDM_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DFSDM1_Channel2_BASE)))

//emgo:const
var DFSDM1_Channel3 = (*DFSDM_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DFSDM1_Channel3_BASE)))

//emgo:const
var DFSDM1_Channel4 = (*DFSDM_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DFSDM1_Channel4_BASE)))

//emgo:const
var DFSDM1_Channel5 = (*DFSDM_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DFSDM1_Channel5_BASE)))

//emgo:const
var DFSDM1_Channel6 = (*DFSDM_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DFSDM1_Channel6_BASE)))

//emgo:const
var DFSDM1_Channel7 = (*DFSDM_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DFSDM1_Channel7_BASE)))

type CHCFGR1_Bits uint32

func (b CHCFGR1_Bits) Field(mask CHCFGR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CHCFGR1_Bits) J(v int) CHCFGR1_Bits {
	return CHCFGR1_Bits(bits.Make32(v, uint32(mask)))
}

type CHCFGR1 struct{ mmio.U32 }

func (r *CHCFGR1) Bits(mask CHCFGR1_Bits) CHCFGR1_Bits { return CHCFGR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *CHCFGR1) StoreBits(mask, b CHCFGR1_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CHCFGR1) SetBits(mask CHCFGR1_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *CHCFGR1) ClearBits(mask CHCFGR1_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *CHCFGR1) Load() CHCFGR1_Bits                  { return CHCFGR1_Bits(r.U32.Load()) }
func (r *CHCFGR1) Store(b CHCFGR1_Bits)                { r.U32.Store(uint32(b)) }

func (r *CHCFGR1) AtomicSetBits(mask CHCFGR1_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CHCFGR1) AtomicClearBits(mask CHCFGR1_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CHCFGR1_Mask struct{ mmio.UM32 }

func (rm CHCFGR1_Mask) Load() CHCFGR1_Bits   { return CHCFGR1_Bits(rm.UM32.Load()) }
func (rm CHCFGR1_Mask) Store(b CHCFGR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DFSDM_Channel_Periph) DFSDMEN() CHCFGR1_Mask {
	return CHCFGR1_Mask{mmio.UM32{&p.CHCFGR1.U32, uint32(DFSDMEN)}}
}

func (p *DFSDM_Channel_Periph) CKOUTSRC() CHCFGR1_Mask {
	return CHCFGR1_Mask{mmio.UM32{&p.CHCFGR1.U32, uint32(CKOUTSRC)}}
}

func (p *DFSDM_Channel_Periph) CKOUTDIV() CHCFGR1_Mask {
	return CHCFGR1_Mask{mmio.UM32{&p.CHCFGR1.U32, uint32(CKOUTDIV)}}
}

func (p *DFSDM_Channel_Periph) DATPACK() CHCFGR1_Mask {
	return CHCFGR1_Mask{mmio.UM32{&p.CHCFGR1.U32, uint32(DATPACK)}}
}

func (p *DFSDM_Channel_Periph) DATMPX() CHCFGR1_Mask {
	return CHCFGR1_Mask{mmio.UM32{&p.CHCFGR1.U32, uint32(DATMPX)}}
}

func (p *DFSDM_Channel_Periph) CHINSEL() CHCFGR1_Mask {
	return CHCFGR1_Mask{mmio.UM32{&p.CHCFGR1.U32, uint32(CHINSEL)}}
}

func (p *DFSDM_Channel_Periph) CHEN() CHCFGR1_Mask {
	return CHCFGR1_Mask{mmio.UM32{&p.CHCFGR1.U32, uint32(CHEN)}}
}

func (p *DFSDM_Channel_Periph) CKABEN() CHCFGR1_Mask {
	return CHCFGR1_Mask{mmio.UM32{&p.CHCFGR1.U32, uint32(CKABEN)}}
}

func (p *DFSDM_Channel_Periph) SCDEN() CHCFGR1_Mask {
	return CHCFGR1_Mask{mmio.UM32{&p.CHCFGR1.U32, uint32(SCDEN)}}
}

func (p *DFSDM_Channel_Periph) SPICKSEL() CHCFGR1_Mask {
	return CHCFGR1_Mask{mmio.UM32{&p.CHCFGR1.U32, uint32(SPICKSEL)}}
}

func (p *DFSDM_Channel_Periph) SITP() CHCFGR1_Mask {
	return CHCFGR1_Mask{mmio.UM32{&p.CHCFGR1.U32, uint32(SITP)}}
}

type CHCFGR2_Bits uint32

func (b CHCFGR2_Bits) Field(mask CHCFGR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CHCFGR2_Bits) J(v int) CHCFGR2_Bits {
	return CHCFGR2_Bits(bits.Make32(v, uint32(mask)))
}

type CHCFGR2 struct{ mmio.U32 }

func (r *CHCFGR2) Bits(mask CHCFGR2_Bits) CHCFGR2_Bits { return CHCFGR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *CHCFGR2) StoreBits(mask, b CHCFGR2_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CHCFGR2) SetBits(mask CHCFGR2_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *CHCFGR2) ClearBits(mask CHCFGR2_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *CHCFGR2) Load() CHCFGR2_Bits                  { return CHCFGR2_Bits(r.U32.Load()) }
func (r *CHCFGR2) Store(b CHCFGR2_Bits)                { r.U32.Store(uint32(b)) }

func (r *CHCFGR2) AtomicSetBits(mask CHCFGR2_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CHCFGR2) AtomicClearBits(mask CHCFGR2_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CHCFGR2_Mask struct{ mmio.UM32 }

func (rm CHCFGR2_Mask) Load() CHCFGR2_Bits   { return CHCFGR2_Bits(rm.UM32.Load()) }
func (rm CHCFGR2_Mask) Store(b CHCFGR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DFSDM_Channel_Periph) OFFSET() CHCFGR2_Mask {
	return CHCFGR2_Mask{mmio.UM32{&p.CHCFGR2.U32, uint32(OFFSET)}}
}

func (p *DFSDM_Channel_Periph) DTRBS() CHCFGR2_Mask {
	return CHCFGR2_Mask{mmio.UM32{&p.CHCFGR2.U32, uint32(DTRBS)}}
}

type CHAWSCDR_Bits uint32

func (b CHAWSCDR_Bits) Field(mask CHAWSCDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CHAWSCDR_Bits) J(v int) CHAWSCDR_Bits {
	return CHAWSCDR_Bits(bits.Make32(v, uint32(mask)))
}

type CHAWSCDR struct{ mmio.U32 }

func (r *CHAWSCDR) Bits(mask CHAWSCDR_Bits) CHAWSCDR_Bits {
	return CHAWSCDR_Bits(r.U32.Bits(uint32(mask)))
}
func (r *CHAWSCDR) StoreBits(mask, b CHAWSCDR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CHAWSCDR) SetBits(mask CHAWSCDR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CHAWSCDR) ClearBits(mask CHAWSCDR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CHAWSCDR) Load() CHAWSCDR_Bits             { return CHAWSCDR_Bits(r.U32.Load()) }
func (r *CHAWSCDR) Store(b CHAWSCDR_Bits)           { r.U32.Store(uint32(b)) }

func (r *CHAWSCDR) AtomicSetBits(mask CHAWSCDR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CHAWSCDR) AtomicClearBits(mask CHAWSCDR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CHAWSCDR_Mask struct{ mmio.UM32 }

func (rm CHAWSCDR_Mask) Load() CHAWSCDR_Bits   { return CHAWSCDR_Bits(rm.UM32.Load()) }
func (rm CHAWSCDR_Mask) Store(b CHAWSCDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DFSDM_Channel_Periph) AWFORD() CHAWSCDR_Mask {
	return CHAWSCDR_Mask{mmio.UM32{&p.CHAWSCDR.U32, uint32(AWFORD)}}
}

func (p *DFSDM_Channel_Periph) AWFOSR() CHAWSCDR_Mask {
	return CHAWSCDR_Mask{mmio.UM32{&p.CHAWSCDR.U32, uint32(AWFOSR)}}
}

func (p *DFSDM_Channel_Periph) BKSCD() CHAWSCDR_Mask {
	return CHAWSCDR_Mask{mmio.UM32{&p.CHAWSCDR.U32, uint32(BKSCD)}}
}

func (p *DFSDM_Channel_Periph) SCDT() CHAWSCDR_Mask {
	return CHAWSCDR_Mask{mmio.UM32{&p.CHAWSCDR.U32, uint32(SCDT)}}
}

type CHWDATAR_Bits uint32

func (b CHWDATAR_Bits) Field(mask CHWDATAR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CHWDATAR_Bits) J(v int) CHWDATAR_Bits {
	return CHWDATAR_Bits(bits.Make32(v, uint32(mask)))
}

type CHWDATAR struct{ mmio.U32 }

func (r *CHWDATAR) Bits(mask CHWDATAR_Bits) CHWDATAR_Bits {
	return CHWDATAR_Bits(r.U32.Bits(uint32(mask)))
}
func (r *CHWDATAR) StoreBits(mask, b CHWDATAR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CHWDATAR) SetBits(mask CHWDATAR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CHWDATAR) ClearBits(mask CHWDATAR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CHWDATAR) Load() CHWDATAR_Bits             { return CHWDATAR_Bits(r.U32.Load()) }
func (r *CHWDATAR) Store(b CHWDATAR_Bits)           { r.U32.Store(uint32(b)) }

func (r *CHWDATAR) AtomicSetBits(mask CHWDATAR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CHWDATAR) AtomicClearBits(mask CHWDATAR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CHWDATAR_Mask struct{ mmio.UM32 }

func (rm CHWDATAR_Mask) Load() CHWDATAR_Bits   { return CHWDATAR_Bits(rm.UM32.Load()) }
func (rm CHWDATAR_Mask) Store(b CHWDATAR_Bits) { rm.UM32.Store(uint32(b)) }

type CHDATINR_Bits uint32

func (b CHDATINR_Bits) Field(mask CHDATINR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CHDATINR_Bits) J(v int) CHDATINR_Bits {
	return CHDATINR_Bits(bits.Make32(v, uint32(mask)))
}

type CHDATINR struct{ mmio.U32 }

func (r *CHDATINR) Bits(mask CHDATINR_Bits) CHDATINR_Bits {
	return CHDATINR_Bits(r.U32.Bits(uint32(mask)))
}
func (r *CHDATINR) StoreBits(mask, b CHDATINR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CHDATINR) SetBits(mask CHDATINR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CHDATINR) ClearBits(mask CHDATINR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CHDATINR) Load() CHDATINR_Bits             { return CHDATINR_Bits(r.U32.Load()) }
func (r *CHDATINR) Store(b CHDATINR_Bits)           { r.U32.Store(uint32(b)) }

func (r *CHDATINR) AtomicSetBits(mask CHDATINR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CHDATINR) AtomicClearBits(mask CHDATINR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CHDATINR_Mask struct{ mmio.UM32 }

func (rm CHDATINR_Mask) Load() CHDATINR_Bits   { return CHDATINR_Bits(rm.UM32.Load()) }
func (rm CHDATINR_Mask) Store(b CHDATINR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DFSDM_Channel_Periph) INDAT0() CHDATINR_Mask {
	return CHDATINR_Mask{mmio.UM32{&p.CHDATINR.U32, uint32(INDAT0)}}
}

func (p *DFSDM_Channel_Periph) INDAT1() CHDATINR_Mask {
	return CHDATINR_Mask{mmio.UM32{&p.CHDATINR.U32, uint32(INDAT1)}}
}
