// +build l1xx_md

package dma

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type DMA_Channel_Periph struct {
	CCR   CCR
	CNDTR CNDTR
	CPAR  CPAR
	CMAR  CMAR
}

func (p *DMA_Channel_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var DMA1_Channel1 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Channel1_BASE)))

//emgo:const
var DMA1_Channel2 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Channel2_BASE)))

//emgo:const
var DMA1_Channel3 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Channel3_BASE)))

//emgo:const
var DMA1_Channel4 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Channel4_BASE)))

//emgo:const
var DMA1_Channel5 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Channel5_BASE)))

//emgo:const
var DMA1_Channel6 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Channel6_BASE)))

//emgo:const
var DMA1_Channel7 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Channel7_BASE)))

//emgo:const
var DMA2_Channel1 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Channel1_BASE)))

//emgo:const
var DMA2_Channel2 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Channel2_BASE)))

//emgo:const
var DMA2_Channel3 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Channel3_BASE)))

//emgo:const
var DMA2_Channel4 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Channel4_BASE)))

//emgo:const
var DMA2_Channel5 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Channel5_BASE)))

type CCR_Bits uint32

func (b CCR_Bits) Field(mask CCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR_Bits) J(v int) CCR_Bits {
	return CCR_Bits(bits.Make32(v, uint32(mask)))
}

type CCR struct{ mmio.U32 }

func (r *CCR) Bits(mask CCR_Bits) CCR_Bits { return CCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CCR) StoreBits(mask, b CCR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CCR) SetBits(mask CCR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CCR) ClearBits(mask CCR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CCR) Load() CCR_Bits              { return CCR_Bits(r.U32.Load()) }
func (r *CCR) Store(b CCR_Bits)            { r.U32.Store(uint32(b)) }

func (r *CCR) AtomicSetBits(mask CCR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CCR) AtomicClearBits(mask CCR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CCR_Mask struct{ mmio.UM32 }

func (rm CCR_Mask) Load() CCR_Bits   { return CCR_Bits(rm.UM32.Load()) }
func (rm CCR_Mask) Store(b CCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Channel_Periph) EN() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(EN)}}
}

func (p *DMA_Channel_Periph) TCIE() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(TCIE)}}
}

func (p *DMA_Channel_Periph) HTIE() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(HTIE)}}
}

func (p *DMA_Channel_Periph) TEIE() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(TEIE)}}
}

func (p *DMA_Channel_Periph) DIR() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(DIR)}}
}

func (p *DMA_Channel_Periph) CIRC() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(CIRC)}}
}

func (p *DMA_Channel_Periph) PINC() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(PINC)}}
}

func (p *DMA_Channel_Periph) MINC() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(MINC)}}
}

func (p *DMA_Channel_Periph) PSIZE() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(PSIZE)}}
}

func (p *DMA_Channel_Periph) MSIZE() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(MSIZE)}}
}

func (p *DMA_Channel_Periph) PL() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(PL)}}
}

func (p *DMA_Channel_Periph) MEM2MEM() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(MEM2MEM)}}
}

type CNDTR_Bits uint32

func (b CNDTR_Bits) Field(mask CNDTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CNDTR_Bits) J(v int) CNDTR_Bits {
	return CNDTR_Bits(bits.Make32(v, uint32(mask)))
}

type CNDTR struct{ mmio.U32 }

func (r *CNDTR) Bits(mask CNDTR_Bits) CNDTR_Bits { return CNDTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CNDTR) StoreBits(mask, b CNDTR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CNDTR) SetBits(mask CNDTR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CNDTR) ClearBits(mask CNDTR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CNDTR) Load() CNDTR_Bits                { return CNDTR_Bits(r.U32.Load()) }
func (r *CNDTR) Store(b CNDTR_Bits)              { r.U32.Store(uint32(b)) }

func (r *CNDTR) AtomicSetBits(mask CNDTR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CNDTR) AtomicClearBits(mask CNDTR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CNDTR_Mask struct{ mmio.UM32 }

func (rm CNDTR_Mask) Load() CNDTR_Bits   { return CNDTR_Bits(rm.UM32.Load()) }
func (rm CNDTR_Mask) Store(b CNDTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Channel_Periph) NDT() CNDTR_Mask {
	return CNDTR_Mask{mmio.UM32{&p.CNDTR.U32, uint32(NDT)}}
}

type CPAR_Bits uint32

func (b CPAR_Bits) Field(mask CPAR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CPAR_Bits) J(v int) CPAR_Bits {
	return CPAR_Bits(bits.Make32(v, uint32(mask)))
}

type CPAR struct{ mmio.U32 }

func (r *CPAR) Bits(mask CPAR_Bits) CPAR_Bits { return CPAR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CPAR) StoreBits(mask, b CPAR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CPAR) SetBits(mask CPAR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CPAR) ClearBits(mask CPAR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CPAR) Load() CPAR_Bits               { return CPAR_Bits(r.U32.Load()) }
func (r *CPAR) Store(b CPAR_Bits)             { r.U32.Store(uint32(b)) }

func (r *CPAR) AtomicSetBits(mask CPAR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CPAR) AtomicClearBits(mask CPAR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CPAR_Mask struct{ mmio.UM32 }

func (rm CPAR_Mask) Load() CPAR_Bits   { return CPAR_Bits(rm.UM32.Load()) }
func (rm CPAR_Mask) Store(b CPAR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Channel_Periph) PA() CPAR_Mask {
	return CPAR_Mask{mmio.UM32{&p.CPAR.U32, uint32(PA)}}
}

type CMAR_Bits uint32

func (b CMAR_Bits) Field(mask CMAR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CMAR_Bits) J(v int) CMAR_Bits {
	return CMAR_Bits(bits.Make32(v, uint32(mask)))
}

type CMAR struct{ mmio.U32 }

func (r *CMAR) Bits(mask CMAR_Bits) CMAR_Bits { return CMAR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CMAR) StoreBits(mask, b CMAR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CMAR) SetBits(mask CMAR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CMAR) ClearBits(mask CMAR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CMAR) Load() CMAR_Bits               { return CMAR_Bits(r.U32.Load()) }
func (r *CMAR) Store(b CMAR_Bits)             { r.U32.Store(uint32(b)) }

func (r *CMAR) AtomicSetBits(mask CMAR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CMAR) AtomicClearBits(mask CMAR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CMAR_Mask struct{ mmio.UM32 }

func (rm CMAR_Mask) Load() CMAR_Bits   { return CMAR_Bits(rm.UM32.Load()) }
func (rm CMAR_Mask) Store(b CMAR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Channel_Periph) MA() CMAR_Mask {
	return CMAR_Mask{mmio.UM32{&p.CMAR.U32, uint32(MA)}}
}
