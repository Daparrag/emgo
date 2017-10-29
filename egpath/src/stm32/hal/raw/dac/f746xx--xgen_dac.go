// +build f746xx

package dac

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type DAC_Periph struct {
	CR      CR
	SWTRIGR SWTRIGR
	DHR12R1 DHR12R1
	DHR12L1 DHR12L1
	DHR8R1  DHR8R1
	DHR12R2 DHR12R2
	DHR12L2 DHR12L2
	DHR8R2  DHR8R2
	DHR12RD DHR12RD
	DHR12LD DHR12LD
	DHR8RD  DHR8RD
	DOR1    DOR1
	DOR2    DOR2
	SR      SR
}

func (p *DAC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var DAC = (*DAC_Periph)(unsafe.Pointer(uintptr(mmap.DAC_BASE)))

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

func (p *DAC_Periph) EN1() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(EN1)}}
}

func (p *DAC_Periph) BOFF1() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(BOFF1)}}
}

func (p *DAC_Periph) TEN1() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TEN1)}}
}

func (p *DAC_Periph) TSEL1() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TSEL1)}}
}

func (p *DAC_Periph) WAVE1() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(WAVE1)}}
}

func (p *DAC_Periph) MAMP1() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(MAMP1)}}
}

func (p *DAC_Periph) DMAEN1() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DMAEN1)}}
}

func (p *DAC_Periph) DMAUDRIE1() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DMAUDRIE1)}}
}

func (p *DAC_Periph) EN2() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(EN2)}}
}

func (p *DAC_Periph) BOFF2() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(BOFF2)}}
}

func (p *DAC_Periph) TEN2() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TEN2)}}
}

func (p *DAC_Periph) TSEL2() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TSEL2)}}
}

func (p *DAC_Periph) WAVE2() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(WAVE2)}}
}

func (p *DAC_Periph) MAMP2() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(MAMP2)}}
}

func (p *DAC_Periph) DMAEN2() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DMAEN2)}}
}

func (p *DAC_Periph) DMAUDRIE2() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DMAUDRIE2)}}
}

type SWTRIGR_Bits uint32

func (b SWTRIGR_Bits) Field(mask SWTRIGR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SWTRIGR_Bits) J(v int) SWTRIGR_Bits {
	return SWTRIGR_Bits(bits.Make32(v, uint32(mask)))
}

type SWTRIGR struct{ mmio.U32 }

func (r *SWTRIGR) Bits(mask SWTRIGR_Bits) SWTRIGR_Bits { return SWTRIGR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SWTRIGR) StoreBits(mask, b SWTRIGR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SWTRIGR) SetBits(mask SWTRIGR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *SWTRIGR) ClearBits(mask SWTRIGR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *SWTRIGR) Load() SWTRIGR_Bits                  { return SWTRIGR_Bits(r.U32.Load()) }
func (r *SWTRIGR) Store(b SWTRIGR_Bits)                { r.U32.Store(uint32(b)) }

func (r *SWTRIGR) AtomicSetBits(mask SWTRIGR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *SWTRIGR) AtomicClearBits(mask SWTRIGR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type SWTRIGR_Mask struct{ mmio.UM32 }

func (rm SWTRIGR_Mask) Load() SWTRIGR_Bits   { return SWTRIGR_Bits(rm.UM32.Load()) }
func (rm SWTRIGR_Mask) Store(b SWTRIGR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) SWTRIG1() SWTRIGR_Mask {
	return SWTRIGR_Mask{mmio.UM32{&p.SWTRIGR.U32, uint32(SWTRIG1)}}
}

func (p *DAC_Periph) SWTRIG2() SWTRIGR_Mask {
	return SWTRIGR_Mask{mmio.UM32{&p.SWTRIGR.U32, uint32(SWTRIG2)}}
}

type DHR12R1_Bits uint32

func (b DHR12R1_Bits) Field(mask DHR12R1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR12R1_Bits) J(v int) DHR12R1_Bits {
	return DHR12R1_Bits(bits.Make32(v, uint32(mask)))
}

type DHR12R1 struct{ mmio.U32 }

func (r *DHR12R1) Bits(mask DHR12R1_Bits) DHR12R1_Bits { return DHR12R1_Bits(r.U32.Bits(uint32(mask))) }
func (r *DHR12R1) StoreBits(mask, b DHR12R1_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DHR12R1) SetBits(mask DHR12R1_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DHR12R1) ClearBits(mask DHR12R1_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DHR12R1) Load() DHR12R1_Bits                  { return DHR12R1_Bits(r.U32.Load()) }
func (r *DHR12R1) Store(b DHR12R1_Bits)                { r.U32.Store(uint32(b)) }

func (r *DHR12R1) AtomicSetBits(mask DHR12R1_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DHR12R1) AtomicClearBits(mask DHR12R1_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DHR12R1_Mask struct{ mmio.UM32 }

func (rm DHR12R1_Mask) Load() DHR12R1_Bits   { return DHR12R1_Bits(rm.UM32.Load()) }
func (rm DHR12R1_Mask) Store(b DHR12R1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC1DHR() DHR12R1_Mask {
	return DHR12R1_Mask{mmio.UM32{&p.DHR12R1.U32, uint32(DACC1DHR)}}
}

type DHR12L1_Bits uint32

func (b DHR12L1_Bits) Field(mask DHR12L1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR12L1_Bits) J(v int) DHR12L1_Bits {
	return DHR12L1_Bits(bits.Make32(v, uint32(mask)))
}

type DHR12L1 struct{ mmio.U32 }

func (r *DHR12L1) Bits(mask DHR12L1_Bits) DHR12L1_Bits { return DHR12L1_Bits(r.U32.Bits(uint32(mask))) }
func (r *DHR12L1) StoreBits(mask, b DHR12L1_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DHR12L1) SetBits(mask DHR12L1_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DHR12L1) ClearBits(mask DHR12L1_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DHR12L1) Load() DHR12L1_Bits                  { return DHR12L1_Bits(r.U32.Load()) }
func (r *DHR12L1) Store(b DHR12L1_Bits)                { r.U32.Store(uint32(b)) }

func (r *DHR12L1) AtomicSetBits(mask DHR12L1_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DHR12L1) AtomicClearBits(mask DHR12L1_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DHR12L1_Mask struct{ mmio.UM32 }

func (rm DHR12L1_Mask) Load() DHR12L1_Bits   { return DHR12L1_Bits(rm.UM32.Load()) }
func (rm DHR12L1_Mask) Store(b DHR12L1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC1DHR() DHR12L1_Mask {
	return DHR12L1_Mask{mmio.UM32{&p.DHR12L1.U32, uint32(DACC1DHR)}}
}

type DHR8R1_Bits uint32

func (b DHR8R1_Bits) Field(mask DHR8R1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR8R1_Bits) J(v int) DHR8R1_Bits {
	return DHR8R1_Bits(bits.Make32(v, uint32(mask)))
}

type DHR8R1 struct{ mmio.U32 }

func (r *DHR8R1) Bits(mask DHR8R1_Bits) DHR8R1_Bits { return DHR8R1_Bits(r.U32.Bits(uint32(mask))) }
func (r *DHR8R1) StoreBits(mask, b DHR8R1_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DHR8R1) SetBits(mask DHR8R1_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *DHR8R1) ClearBits(mask DHR8R1_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *DHR8R1) Load() DHR8R1_Bits                 { return DHR8R1_Bits(r.U32.Load()) }
func (r *DHR8R1) Store(b DHR8R1_Bits)               { r.U32.Store(uint32(b)) }

func (r *DHR8R1) AtomicSetBits(mask DHR8R1_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DHR8R1) AtomicClearBits(mask DHR8R1_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DHR8R1_Mask struct{ mmio.UM32 }

func (rm DHR8R1_Mask) Load() DHR8R1_Bits   { return DHR8R1_Bits(rm.UM32.Load()) }
func (rm DHR8R1_Mask) Store(b DHR8R1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC1DHR() DHR8R1_Mask {
	return DHR8R1_Mask{mmio.UM32{&p.DHR8R1.U32, uint32(DACC1DHR)}}
}

type DHR12R2_Bits uint32

func (b DHR12R2_Bits) Field(mask DHR12R2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR12R2_Bits) J(v int) DHR12R2_Bits {
	return DHR12R2_Bits(bits.Make32(v, uint32(mask)))
}

type DHR12R2 struct{ mmio.U32 }

func (r *DHR12R2) Bits(mask DHR12R2_Bits) DHR12R2_Bits { return DHR12R2_Bits(r.U32.Bits(uint32(mask))) }
func (r *DHR12R2) StoreBits(mask, b DHR12R2_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DHR12R2) SetBits(mask DHR12R2_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DHR12R2) ClearBits(mask DHR12R2_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DHR12R2) Load() DHR12R2_Bits                  { return DHR12R2_Bits(r.U32.Load()) }
func (r *DHR12R2) Store(b DHR12R2_Bits)                { r.U32.Store(uint32(b)) }

func (r *DHR12R2) AtomicSetBits(mask DHR12R2_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DHR12R2) AtomicClearBits(mask DHR12R2_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DHR12R2_Mask struct{ mmio.UM32 }

func (rm DHR12R2_Mask) Load() DHR12R2_Bits   { return DHR12R2_Bits(rm.UM32.Load()) }
func (rm DHR12R2_Mask) Store(b DHR12R2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC2DHR() DHR12R2_Mask {
	return DHR12R2_Mask{mmio.UM32{&p.DHR12R2.U32, uint32(DACC2DHR)}}
}

type DHR12L2_Bits uint32

func (b DHR12L2_Bits) Field(mask DHR12L2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR12L2_Bits) J(v int) DHR12L2_Bits {
	return DHR12L2_Bits(bits.Make32(v, uint32(mask)))
}

type DHR12L2 struct{ mmio.U32 }

func (r *DHR12L2) Bits(mask DHR12L2_Bits) DHR12L2_Bits { return DHR12L2_Bits(r.U32.Bits(uint32(mask))) }
func (r *DHR12L2) StoreBits(mask, b DHR12L2_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DHR12L2) SetBits(mask DHR12L2_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DHR12L2) ClearBits(mask DHR12L2_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DHR12L2) Load() DHR12L2_Bits                  { return DHR12L2_Bits(r.U32.Load()) }
func (r *DHR12L2) Store(b DHR12L2_Bits)                { r.U32.Store(uint32(b)) }

func (r *DHR12L2) AtomicSetBits(mask DHR12L2_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DHR12L2) AtomicClearBits(mask DHR12L2_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DHR12L2_Mask struct{ mmio.UM32 }

func (rm DHR12L2_Mask) Load() DHR12L2_Bits   { return DHR12L2_Bits(rm.UM32.Load()) }
func (rm DHR12L2_Mask) Store(b DHR12L2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC2DHR() DHR12L2_Mask {
	return DHR12L2_Mask{mmio.UM32{&p.DHR12L2.U32, uint32(DACC2DHR)}}
}

type DHR8R2_Bits uint32

func (b DHR8R2_Bits) Field(mask DHR8R2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR8R2_Bits) J(v int) DHR8R2_Bits {
	return DHR8R2_Bits(bits.Make32(v, uint32(mask)))
}

type DHR8R2 struct{ mmio.U32 }

func (r *DHR8R2) Bits(mask DHR8R2_Bits) DHR8R2_Bits { return DHR8R2_Bits(r.U32.Bits(uint32(mask))) }
func (r *DHR8R2) StoreBits(mask, b DHR8R2_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DHR8R2) SetBits(mask DHR8R2_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *DHR8R2) ClearBits(mask DHR8R2_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *DHR8R2) Load() DHR8R2_Bits                 { return DHR8R2_Bits(r.U32.Load()) }
func (r *DHR8R2) Store(b DHR8R2_Bits)               { r.U32.Store(uint32(b)) }

func (r *DHR8R2) AtomicSetBits(mask DHR8R2_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DHR8R2) AtomicClearBits(mask DHR8R2_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DHR8R2_Mask struct{ mmio.UM32 }

func (rm DHR8R2_Mask) Load() DHR8R2_Bits   { return DHR8R2_Bits(rm.UM32.Load()) }
func (rm DHR8R2_Mask) Store(b DHR8R2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC2DHR() DHR8R2_Mask {
	return DHR8R2_Mask{mmio.UM32{&p.DHR8R2.U32, uint32(DACC2DHR)}}
}

type DHR12RD_Bits uint32

func (b DHR12RD_Bits) Field(mask DHR12RD_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR12RD_Bits) J(v int) DHR12RD_Bits {
	return DHR12RD_Bits(bits.Make32(v, uint32(mask)))
}

type DHR12RD struct{ mmio.U32 }

func (r *DHR12RD) Bits(mask DHR12RD_Bits) DHR12RD_Bits { return DHR12RD_Bits(r.U32.Bits(uint32(mask))) }
func (r *DHR12RD) StoreBits(mask, b DHR12RD_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DHR12RD) SetBits(mask DHR12RD_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DHR12RD) ClearBits(mask DHR12RD_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DHR12RD) Load() DHR12RD_Bits                  { return DHR12RD_Bits(r.U32.Load()) }
func (r *DHR12RD) Store(b DHR12RD_Bits)                { r.U32.Store(uint32(b)) }

func (r *DHR12RD) AtomicSetBits(mask DHR12RD_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DHR12RD) AtomicClearBits(mask DHR12RD_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DHR12RD_Mask struct{ mmio.UM32 }

func (rm DHR12RD_Mask) Load() DHR12RD_Bits   { return DHR12RD_Bits(rm.UM32.Load()) }
func (rm DHR12RD_Mask) Store(b DHR12RD_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC1DHR() DHR12RD_Mask {
	return DHR12RD_Mask{mmio.UM32{&p.DHR12RD.U32, uint32(DACC1DHR)}}
}

func (p *DAC_Periph) DACC2DHR() DHR12RD_Mask {
	return DHR12RD_Mask{mmio.UM32{&p.DHR12RD.U32, uint32(DACC2DHR)}}
}

type DHR12LD_Bits uint32

func (b DHR12LD_Bits) Field(mask DHR12LD_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR12LD_Bits) J(v int) DHR12LD_Bits {
	return DHR12LD_Bits(bits.Make32(v, uint32(mask)))
}

type DHR12LD struct{ mmio.U32 }

func (r *DHR12LD) Bits(mask DHR12LD_Bits) DHR12LD_Bits { return DHR12LD_Bits(r.U32.Bits(uint32(mask))) }
func (r *DHR12LD) StoreBits(mask, b DHR12LD_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DHR12LD) SetBits(mask DHR12LD_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DHR12LD) ClearBits(mask DHR12LD_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DHR12LD) Load() DHR12LD_Bits                  { return DHR12LD_Bits(r.U32.Load()) }
func (r *DHR12LD) Store(b DHR12LD_Bits)                { r.U32.Store(uint32(b)) }

func (r *DHR12LD) AtomicSetBits(mask DHR12LD_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DHR12LD) AtomicClearBits(mask DHR12LD_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DHR12LD_Mask struct{ mmio.UM32 }

func (rm DHR12LD_Mask) Load() DHR12LD_Bits   { return DHR12LD_Bits(rm.UM32.Load()) }
func (rm DHR12LD_Mask) Store(b DHR12LD_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC1DHR() DHR12LD_Mask {
	return DHR12LD_Mask{mmio.UM32{&p.DHR12LD.U32, uint32(DACC1DHR)}}
}

func (p *DAC_Periph) DACC2DHR() DHR12LD_Mask {
	return DHR12LD_Mask{mmio.UM32{&p.DHR12LD.U32, uint32(DACC2DHR)}}
}

type DHR8RD_Bits uint32

func (b DHR8RD_Bits) Field(mask DHR8RD_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR8RD_Bits) J(v int) DHR8RD_Bits {
	return DHR8RD_Bits(bits.Make32(v, uint32(mask)))
}

type DHR8RD struct{ mmio.U32 }

func (r *DHR8RD) Bits(mask DHR8RD_Bits) DHR8RD_Bits { return DHR8RD_Bits(r.U32.Bits(uint32(mask))) }
func (r *DHR8RD) StoreBits(mask, b DHR8RD_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DHR8RD) SetBits(mask DHR8RD_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *DHR8RD) ClearBits(mask DHR8RD_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *DHR8RD) Load() DHR8RD_Bits                 { return DHR8RD_Bits(r.U32.Load()) }
func (r *DHR8RD) Store(b DHR8RD_Bits)               { r.U32.Store(uint32(b)) }

func (r *DHR8RD) AtomicSetBits(mask DHR8RD_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DHR8RD) AtomicClearBits(mask DHR8RD_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DHR8RD_Mask struct{ mmio.UM32 }

func (rm DHR8RD_Mask) Load() DHR8RD_Bits   { return DHR8RD_Bits(rm.UM32.Load()) }
func (rm DHR8RD_Mask) Store(b DHR8RD_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC1DHR() DHR8RD_Mask {
	return DHR8RD_Mask{mmio.UM32{&p.DHR8RD.U32, uint32(DACC1DHR)}}
}

func (p *DAC_Periph) DACC2DHR() DHR8RD_Mask {
	return DHR8RD_Mask{mmio.UM32{&p.DHR8RD.U32, uint32(DACC2DHR)}}
}

type DOR1_Bits uint32

func (b DOR1_Bits) Field(mask DOR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOR1_Bits) J(v int) DOR1_Bits {
	return DOR1_Bits(bits.Make32(v, uint32(mask)))
}

type DOR1 struct{ mmio.U32 }

func (r *DOR1) Bits(mask DOR1_Bits) DOR1_Bits { return DOR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *DOR1) StoreBits(mask, b DOR1_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DOR1) SetBits(mask DOR1_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *DOR1) ClearBits(mask DOR1_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *DOR1) Load() DOR1_Bits               { return DOR1_Bits(r.U32.Load()) }
func (r *DOR1) Store(b DOR1_Bits)             { r.U32.Store(uint32(b)) }

func (r *DOR1) AtomicSetBits(mask DOR1_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DOR1) AtomicClearBits(mask DOR1_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DOR1_Mask struct{ mmio.UM32 }

func (rm DOR1_Mask) Load() DOR1_Bits   { return DOR1_Bits(rm.UM32.Load()) }
func (rm DOR1_Mask) Store(b DOR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC1DOR() DOR1_Mask {
	return DOR1_Mask{mmio.UM32{&p.DOR1.U32, uint32(DACC1DOR)}}
}

type DOR2_Bits uint32

func (b DOR2_Bits) Field(mask DOR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOR2_Bits) J(v int) DOR2_Bits {
	return DOR2_Bits(bits.Make32(v, uint32(mask)))
}

type DOR2 struct{ mmio.U32 }

func (r *DOR2) Bits(mask DOR2_Bits) DOR2_Bits { return DOR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *DOR2) StoreBits(mask, b DOR2_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DOR2) SetBits(mask DOR2_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *DOR2) ClearBits(mask DOR2_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *DOR2) Load() DOR2_Bits               { return DOR2_Bits(r.U32.Load()) }
func (r *DOR2) Store(b DOR2_Bits)             { r.U32.Store(uint32(b)) }

func (r *DOR2) AtomicSetBits(mask DOR2_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DOR2) AtomicClearBits(mask DOR2_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DOR2_Mask struct{ mmio.UM32 }

func (rm DOR2_Mask) Load() DOR2_Bits   { return DOR2_Bits(rm.UM32.Load()) }
func (rm DOR2_Mask) Store(b DOR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC2DOR() DOR2_Mask {
	return DOR2_Mask{mmio.UM32{&p.DOR2.U32, uint32(DACC2DOR)}}
}

type SR_Bits uint32

func (b SR_Bits) Field(mask SR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR_Bits) J(v int) SR_Bits {
	return SR_Bits(bits.Make32(v, uint32(mask)))
}

type SR struct{ mmio.U32 }

func (r *SR) Bits(mask SR_Bits) SR_Bits { return SR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR) StoreBits(mask, b SR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR) SetBits(mask SR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *SR) ClearBits(mask SR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *SR) Load() SR_Bits             { return SR_Bits(r.U32.Load()) }
func (r *SR) Store(b SR_Bits)           { r.U32.Store(uint32(b)) }

func (r *SR) AtomicSetBits(mask SR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *SR) AtomicClearBits(mask SR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type SR_Mask struct{ mmio.UM32 }

func (rm SR_Mask) Load() SR_Bits   { return SR_Bits(rm.UM32.Load()) }
func (rm SR_Mask) Store(b SR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DMAUDR1() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(DMAUDR1)}}
}

func (p *DAC_Periph) DMAUDR2() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(DMAUDR2)}}
}
