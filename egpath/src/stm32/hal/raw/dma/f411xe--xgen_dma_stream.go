// +build f411xe

package dma

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f411xe/mmap"
)

type DMA_Stream_Periph struct {
	CR   CR
	NDTR NDTR
	PAR  PAR
	M0AR M0AR
	M1AR M1AR
	FCR  FCR
}

func (p *DMA_Stream_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var DMA1_Stream0 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Stream0_BASE)))

//emgo:const
var DMA1_Stream1 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Stream1_BASE)))

//emgo:const
var DMA1_Stream2 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Stream2_BASE)))

//emgo:const
var DMA1_Stream3 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Stream3_BASE)))

//emgo:const
var DMA1_Stream4 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Stream4_BASE)))

//emgo:const
var DMA1_Stream5 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Stream5_BASE)))

//emgo:const
var DMA1_Stream6 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Stream6_BASE)))

//emgo:const
var DMA1_Stream7 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Stream7_BASE)))

//emgo:const
var DMA2_Stream0 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Stream0_BASE)))

//emgo:const
var DMA2_Stream1 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Stream1_BASE)))

//emgo:const
var DMA2_Stream2 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Stream2_BASE)))

//emgo:const
var DMA2_Stream3 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Stream3_BASE)))

//emgo:const
var DMA2_Stream4 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Stream4_BASE)))

//emgo:const
var DMA2_Stream5 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Stream5_BASE)))

//emgo:const
var DMA2_Stream6 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Stream6_BASE)))

//emgo:const
var DMA2_Stream7 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Stream7_BASE)))

type CR_Bits uint32

func (b CR_Bits) Field(mask CR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR_Bits) J(v int) CR_Bits {
	return CR_Bits(bits.Make32(v, uint32(mask)))
}

type CR struct{ mmio.U32 }

func (r *CR) Bits(mask CR_Bits) CR_Bits { return CR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR) StoreBits(mask, b CR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR) SetBits(mask CR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CR) ClearBits(mask CR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CR) Load() CR_Bits             { return CR_Bits(r.U32.Load()) }
func (r *CR) Store(b CR_Bits)           { r.U32.Store(uint32(b)) }

func (r *CR) AtomicSetBits(mask CR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CR) AtomicClearBits(mask CR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CR_Mask struct{ mmio.UM32 }

func (rm CR_Mask) Load() CR_Bits   { return CR_Bits(rm.UM32.Load()) }
func (rm CR_Mask) Store(b CR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Stream_Periph) CHSEL() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(CHSEL)}}
}

func (p *DMA_Stream_Periph) MBURST() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(MBURST)}}
}

func (p *DMA_Stream_Periph) PBURST() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PBURST)}}
}

func (p *DMA_Stream_Periph) ACK() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ACK)}}
}

func (p *DMA_Stream_Periph) CT() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(CT)}}
}

func (p *DMA_Stream_Periph) DBM() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBM)}}
}

func (p *DMA_Stream_Periph) PL() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PL)}}
}

func (p *DMA_Stream_Periph) PINCOS() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PINCOS)}}
}

func (p *DMA_Stream_Periph) MSIZE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(MSIZE)}}
}

func (p *DMA_Stream_Periph) PSIZE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PSIZE)}}
}

func (p *DMA_Stream_Periph) MINC() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(MINC)}}
}

func (p *DMA_Stream_Periph) PINC() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PINC)}}
}

func (p *DMA_Stream_Periph) CIRC() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(CIRC)}}
}

func (p *DMA_Stream_Periph) DIR() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DIR)}}
}

func (p *DMA_Stream_Periph) PFCTRL() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PFCTRL)}}
}

func (p *DMA_Stream_Periph) TCIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TCIE)}}
}

func (p *DMA_Stream_Periph) HTIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HTIE)}}
}

func (p *DMA_Stream_Periph) TEIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TEIE)}}
}

func (p *DMA_Stream_Periph) DMEIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DMEIE)}}
}

func (p *DMA_Stream_Periph) EN() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(EN)}}
}

type NDTR_Bits uint32

func (b NDTR_Bits) Field(mask NDTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask NDTR_Bits) J(v int) NDTR_Bits {
	return NDTR_Bits(bits.Make32(v, uint32(mask)))
}

type NDTR struct{ mmio.U32 }

func (r *NDTR) Bits(mask NDTR_Bits) NDTR_Bits { return NDTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *NDTR) StoreBits(mask, b NDTR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *NDTR) SetBits(mask NDTR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *NDTR) ClearBits(mask NDTR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *NDTR) Load() NDTR_Bits               { return NDTR_Bits(r.U32.Load()) }
func (r *NDTR) Store(b NDTR_Bits)             { r.U32.Store(uint32(b)) }

func (r *NDTR) AtomicSetBits(mask NDTR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *NDTR) AtomicClearBits(mask NDTR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type NDTR_Mask struct{ mmio.UM32 }

func (rm NDTR_Mask) Load() NDTR_Bits   { return NDTR_Bits(rm.UM32.Load()) }
func (rm NDTR_Mask) Store(b NDTR_Bits) { rm.UM32.Store(uint32(b)) }

type PAR_Bits uint32

func (b PAR_Bits) Field(mask PAR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PAR_Bits) J(v int) PAR_Bits {
	return PAR_Bits(bits.Make32(v, uint32(mask)))
}

type PAR struct{ mmio.U32 }

func (r *PAR) Bits(mask PAR_Bits) PAR_Bits { return PAR_Bits(r.U32.Bits(uint32(mask))) }
func (r *PAR) StoreBits(mask, b PAR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PAR) SetBits(mask PAR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *PAR) ClearBits(mask PAR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *PAR) Load() PAR_Bits              { return PAR_Bits(r.U32.Load()) }
func (r *PAR) Store(b PAR_Bits)            { r.U32.Store(uint32(b)) }

func (r *PAR) AtomicSetBits(mask PAR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PAR) AtomicClearBits(mask PAR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type PAR_Mask struct{ mmio.UM32 }

func (rm PAR_Mask) Load() PAR_Bits   { return PAR_Bits(rm.UM32.Load()) }
func (rm PAR_Mask) Store(b PAR_Bits) { rm.UM32.Store(uint32(b)) }

type M0AR_Bits uint32

func (b M0AR_Bits) Field(mask M0AR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask M0AR_Bits) J(v int) M0AR_Bits {
	return M0AR_Bits(bits.Make32(v, uint32(mask)))
}

type M0AR struct{ mmio.U32 }

func (r *M0AR) Bits(mask M0AR_Bits) M0AR_Bits { return M0AR_Bits(r.U32.Bits(uint32(mask))) }
func (r *M0AR) StoreBits(mask, b M0AR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *M0AR) SetBits(mask M0AR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *M0AR) ClearBits(mask M0AR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *M0AR) Load() M0AR_Bits               { return M0AR_Bits(r.U32.Load()) }
func (r *M0AR) Store(b M0AR_Bits)             { r.U32.Store(uint32(b)) }

func (r *M0AR) AtomicSetBits(mask M0AR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *M0AR) AtomicClearBits(mask M0AR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type M0AR_Mask struct{ mmio.UM32 }

func (rm M0AR_Mask) Load() M0AR_Bits   { return M0AR_Bits(rm.UM32.Load()) }
func (rm M0AR_Mask) Store(b M0AR_Bits) { rm.UM32.Store(uint32(b)) }

type M1AR_Bits uint32

func (b M1AR_Bits) Field(mask M1AR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask M1AR_Bits) J(v int) M1AR_Bits {
	return M1AR_Bits(bits.Make32(v, uint32(mask)))
}

type M1AR struct{ mmio.U32 }

func (r *M1AR) Bits(mask M1AR_Bits) M1AR_Bits { return M1AR_Bits(r.U32.Bits(uint32(mask))) }
func (r *M1AR) StoreBits(mask, b M1AR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *M1AR) SetBits(mask M1AR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *M1AR) ClearBits(mask M1AR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *M1AR) Load() M1AR_Bits               { return M1AR_Bits(r.U32.Load()) }
func (r *M1AR) Store(b M1AR_Bits)             { r.U32.Store(uint32(b)) }

func (r *M1AR) AtomicSetBits(mask M1AR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *M1AR) AtomicClearBits(mask M1AR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type M1AR_Mask struct{ mmio.UM32 }

func (rm M1AR_Mask) Load() M1AR_Bits   { return M1AR_Bits(rm.UM32.Load()) }
func (rm M1AR_Mask) Store(b M1AR_Bits) { rm.UM32.Store(uint32(b)) }

type FCR_Bits uint32

func (b FCR_Bits) Field(mask FCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FCR_Bits) J(v int) FCR_Bits {
	return FCR_Bits(bits.Make32(v, uint32(mask)))
}

type FCR struct{ mmio.U32 }

func (r *FCR) Bits(mask FCR_Bits) FCR_Bits { return FCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *FCR) StoreBits(mask, b FCR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FCR) SetBits(mask FCR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *FCR) ClearBits(mask FCR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *FCR) Load() FCR_Bits              { return FCR_Bits(r.U32.Load()) }
func (r *FCR) Store(b FCR_Bits)            { r.U32.Store(uint32(b)) }

func (r *FCR) AtomicSetBits(mask FCR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *FCR) AtomicClearBits(mask FCR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type FCR_Mask struct{ mmio.UM32 }

func (rm FCR_Mask) Load() FCR_Bits   { return FCR_Bits(rm.UM32.Load()) }
func (rm FCR_Mask) Store(b FCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Stream_Periph) FEIE() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(FEIE)}}
}

func (p *DMA_Stream_Periph) FS() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(FS)}}
}

func (p *DMA_Stream_Periph) DMDIS() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(DMDIS)}}
}

func (p *DMA_Stream_Periph) FTH() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(FTH)}}
}
