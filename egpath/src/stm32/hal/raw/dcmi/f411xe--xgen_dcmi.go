// +build f411xe

package dcmi

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f411xe/mmap"
)

type DCMI_Periph struct {
	CR      CR
	SR      SR
	RISR    RISR
	IER     IER
	MISR    MISR
	ICR     ICR
	ESCR    ESCR
	ESUR    ESUR
	CWSTRTR CWSTRTR
	CWSIZER CWSIZER
	DR      DR
}

func (p *DCMI_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var DCMI = (*DCMI_Periph)(unsafe.Pointer(uintptr(mmap.DCMI_BASE)))

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

func (p *DCMI_Periph) CAPTURE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(CAPTURE)}}
}

func (p *DCMI_Periph) CM() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(CM)}}
}

func (p *DCMI_Periph) CROP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(CROP)}}
}

func (p *DCMI_Periph) JPEG() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(JPEG)}}
}

func (p *DCMI_Periph) ESS() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ESS)}}
}

func (p *DCMI_Periph) PCKPOL() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PCKPOL)}}
}

func (p *DCMI_Periph) HSPOL() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSPOL)}}
}

func (p *DCMI_Periph) VSPOL() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(VSPOL)}}
}

func (p *DCMI_Periph) FCRC_0() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(FCRC_0)}}
}

func (p *DCMI_Periph) FCRC_1() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(FCRC_1)}}
}

func (p *DCMI_Periph) EDM_0() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(EDM_0)}}
}

func (p *DCMI_Periph) EDM_1() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(EDM_1)}}
}

func (p *DCMI_Periph) CRE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(CRE)}}
}

func (p *DCMI_Periph) ENABLE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ENABLE)}}
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

func (p *DCMI_Periph) HSYNC() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(HSYNC)}}
}

func (p *DCMI_Periph) VSYNC() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(VSYNC)}}
}

func (p *DCMI_Periph) FNE() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(FNE)}}
}

type RISR_Bits uint32

func (b RISR_Bits) Field(mask RISR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RISR_Bits) J(v int) RISR_Bits {
	return RISR_Bits(bits.Make32(v, uint32(mask)))
}

type RISR struct{ mmio.U32 }

func (r *RISR) Bits(mask RISR_Bits) RISR_Bits { return RISR_Bits(r.U32.Bits(uint32(mask))) }
func (r *RISR) StoreBits(mask, b RISR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RISR) SetBits(mask RISR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *RISR) ClearBits(mask RISR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *RISR) Load() RISR_Bits               { return RISR_Bits(r.U32.Load()) }
func (r *RISR) Store(b RISR_Bits)             { r.U32.Store(uint32(b)) }

func (r *RISR) AtomicSetBits(mask RISR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RISR) AtomicClearBits(mask RISR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type RISR_Mask struct{ mmio.UM32 }

func (rm RISR_Mask) Load() RISR_Bits   { return RISR_Bits(rm.UM32.Load()) }
func (rm RISR_Mask) Store(b RISR_Bits) { rm.UM32.Store(uint32(b)) }

type IER_Bits uint32

func (b IER_Bits) Field(mask IER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IER_Bits) J(v int) IER_Bits {
	return IER_Bits(bits.Make32(v, uint32(mask)))
}

type IER struct{ mmio.U32 }

func (r *IER) Bits(mask IER_Bits) IER_Bits { return IER_Bits(r.U32.Bits(uint32(mask))) }
func (r *IER) StoreBits(mask, b IER_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IER) SetBits(mask IER_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *IER) ClearBits(mask IER_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *IER) Load() IER_Bits              { return IER_Bits(r.U32.Load()) }
func (r *IER) Store(b IER_Bits)            { r.U32.Store(uint32(b)) }

func (r *IER) AtomicSetBits(mask IER_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *IER) AtomicClearBits(mask IER_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type IER_Mask struct{ mmio.UM32 }

func (rm IER_Mask) Load() IER_Bits   { return IER_Bits(rm.UM32.Load()) }
func (rm IER_Mask) Store(b IER_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DCMI_Periph) FRAME_IE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(FRAME_IE)}}
}

func (p *DCMI_Periph) OVR_IE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(OVR_IE)}}
}

func (p *DCMI_Periph) ERR_IE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(ERR_IE)}}
}

func (p *DCMI_Periph) VSYNC_IE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(VSYNC_IE)}}
}

func (p *DCMI_Periph) LINE_IE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(LINE_IE)}}
}

type MISR_Bits uint32

func (b MISR_Bits) Field(mask MISR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MISR_Bits) J(v int) MISR_Bits {
	return MISR_Bits(bits.Make32(v, uint32(mask)))
}

type MISR struct{ mmio.U32 }

func (r *MISR) Bits(mask MISR_Bits) MISR_Bits { return MISR_Bits(r.U32.Bits(uint32(mask))) }
func (r *MISR) StoreBits(mask, b MISR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *MISR) SetBits(mask MISR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *MISR) ClearBits(mask MISR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *MISR) Load() MISR_Bits               { return MISR_Bits(r.U32.Load()) }
func (r *MISR) Store(b MISR_Bits)             { r.U32.Store(uint32(b)) }

func (r *MISR) AtomicSetBits(mask MISR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *MISR) AtomicClearBits(mask MISR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type MISR_Mask struct{ mmio.UM32 }

func (rm MISR_Mask) Load() MISR_Bits   { return MISR_Bits(rm.UM32.Load()) }
func (rm MISR_Mask) Store(b MISR_Bits) { rm.UM32.Store(uint32(b)) }

type ICR_Bits uint32

func (b ICR_Bits) Field(mask ICR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ICR_Bits) J(v int) ICR_Bits {
	return ICR_Bits(bits.Make32(v, uint32(mask)))
}

type ICR struct{ mmio.U32 }

func (r *ICR) Bits(mask ICR_Bits) ICR_Bits { return ICR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ICR) StoreBits(mask, b ICR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ICR) SetBits(mask ICR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ICR) ClearBits(mask ICR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ICR) Load() ICR_Bits              { return ICR_Bits(r.U32.Load()) }
func (r *ICR) Store(b ICR_Bits)            { r.U32.Store(uint32(b)) }

func (r *ICR) AtomicSetBits(mask ICR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ICR) AtomicClearBits(mask ICR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type ICR_Mask struct{ mmio.UM32 }

func (rm ICR_Mask) Load() ICR_Bits   { return ICR_Bits(rm.UM32.Load()) }
func (rm ICR_Mask) Store(b ICR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DCMI_Periph) FRAME_ISC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(FRAME_ISC)}}
}

func (p *DCMI_Periph) OVR_ISC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(OVR_ISC)}}
}

func (p *DCMI_Periph) ERR_ISC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(ERR_ISC)}}
}

func (p *DCMI_Periph) VSYNC_ISC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(VSYNC_ISC)}}
}

func (p *DCMI_Periph) LINE_ISC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(LINE_ISC)}}
}

type ESCR_Bits uint32

func (b ESCR_Bits) Field(mask ESCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ESCR_Bits) J(v int) ESCR_Bits {
	return ESCR_Bits(bits.Make32(v, uint32(mask)))
}

type ESCR struct{ mmio.U32 }

func (r *ESCR) Bits(mask ESCR_Bits) ESCR_Bits { return ESCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ESCR) StoreBits(mask, b ESCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ESCR) SetBits(mask ESCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *ESCR) ClearBits(mask ESCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *ESCR) Load() ESCR_Bits               { return ESCR_Bits(r.U32.Load()) }
func (r *ESCR) Store(b ESCR_Bits)             { r.U32.Store(uint32(b)) }

func (r *ESCR) AtomicSetBits(mask ESCR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ESCR) AtomicClearBits(mask ESCR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type ESCR_Mask struct{ mmio.UM32 }

func (rm ESCR_Mask) Load() ESCR_Bits   { return ESCR_Bits(rm.UM32.Load()) }
func (rm ESCR_Mask) Store(b ESCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DCMI_Periph) FSC() ESCR_Mask {
	return ESCR_Mask{mmio.UM32{&p.ESCR.U32, uint32(FSC)}}
}

func (p *DCMI_Periph) LSC() ESCR_Mask {
	return ESCR_Mask{mmio.UM32{&p.ESCR.U32, uint32(LSC)}}
}

func (p *DCMI_Periph) LEC() ESCR_Mask {
	return ESCR_Mask{mmio.UM32{&p.ESCR.U32, uint32(LEC)}}
}

func (p *DCMI_Periph) FEC() ESCR_Mask {
	return ESCR_Mask{mmio.UM32{&p.ESCR.U32, uint32(FEC)}}
}

type ESUR_Bits uint32

func (b ESUR_Bits) Field(mask ESUR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ESUR_Bits) J(v int) ESUR_Bits {
	return ESUR_Bits(bits.Make32(v, uint32(mask)))
}

type ESUR struct{ mmio.U32 }

func (r *ESUR) Bits(mask ESUR_Bits) ESUR_Bits { return ESUR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ESUR) StoreBits(mask, b ESUR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ESUR) SetBits(mask ESUR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *ESUR) ClearBits(mask ESUR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *ESUR) Load() ESUR_Bits               { return ESUR_Bits(r.U32.Load()) }
func (r *ESUR) Store(b ESUR_Bits)             { r.U32.Store(uint32(b)) }

func (r *ESUR) AtomicSetBits(mask ESUR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ESUR) AtomicClearBits(mask ESUR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type ESUR_Mask struct{ mmio.UM32 }

func (rm ESUR_Mask) Load() ESUR_Bits   { return ESUR_Bits(rm.UM32.Load()) }
func (rm ESUR_Mask) Store(b ESUR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DCMI_Periph) FSU() ESUR_Mask {
	return ESUR_Mask{mmio.UM32{&p.ESUR.U32, uint32(FSU)}}
}

func (p *DCMI_Periph) LSU() ESUR_Mask {
	return ESUR_Mask{mmio.UM32{&p.ESUR.U32, uint32(LSU)}}
}

func (p *DCMI_Periph) LEU() ESUR_Mask {
	return ESUR_Mask{mmio.UM32{&p.ESUR.U32, uint32(LEU)}}
}

func (p *DCMI_Periph) FEU() ESUR_Mask {
	return ESUR_Mask{mmio.UM32{&p.ESUR.U32, uint32(FEU)}}
}

type CWSTRTR_Bits uint32

func (b CWSTRTR_Bits) Field(mask CWSTRTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CWSTRTR_Bits) J(v int) CWSTRTR_Bits {
	return CWSTRTR_Bits(bits.Make32(v, uint32(mask)))
}

type CWSTRTR struct{ mmio.U32 }

func (r *CWSTRTR) Bits(mask CWSTRTR_Bits) CWSTRTR_Bits { return CWSTRTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CWSTRTR) StoreBits(mask, b CWSTRTR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CWSTRTR) SetBits(mask CWSTRTR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *CWSTRTR) ClearBits(mask CWSTRTR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *CWSTRTR) Load() CWSTRTR_Bits                  { return CWSTRTR_Bits(r.U32.Load()) }
func (r *CWSTRTR) Store(b CWSTRTR_Bits)                { r.U32.Store(uint32(b)) }

func (r *CWSTRTR) AtomicSetBits(mask CWSTRTR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CWSTRTR) AtomicClearBits(mask CWSTRTR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CWSTRTR_Mask struct{ mmio.UM32 }

func (rm CWSTRTR_Mask) Load() CWSTRTR_Bits   { return CWSTRTR_Bits(rm.UM32.Load()) }
func (rm CWSTRTR_Mask) Store(b CWSTRTR_Bits) { rm.UM32.Store(uint32(b)) }

type CWSIZER_Bits uint32

func (b CWSIZER_Bits) Field(mask CWSIZER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CWSIZER_Bits) J(v int) CWSIZER_Bits {
	return CWSIZER_Bits(bits.Make32(v, uint32(mask)))
}

type CWSIZER struct{ mmio.U32 }

func (r *CWSIZER) Bits(mask CWSIZER_Bits) CWSIZER_Bits { return CWSIZER_Bits(r.U32.Bits(uint32(mask))) }
func (r *CWSIZER) StoreBits(mask, b CWSIZER_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CWSIZER) SetBits(mask CWSIZER_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *CWSIZER) ClearBits(mask CWSIZER_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *CWSIZER) Load() CWSIZER_Bits                  { return CWSIZER_Bits(r.U32.Load()) }
func (r *CWSIZER) Store(b CWSIZER_Bits)                { r.U32.Store(uint32(b)) }

func (r *CWSIZER) AtomicSetBits(mask CWSIZER_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CWSIZER) AtomicClearBits(mask CWSIZER_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CWSIZER_Mask struct{ mmio.UM32 }

func (rm CWSIZER_Mask) Load() CWSIZER_Bits   { return CWSIZER_Bits(rm.UM32.Load()) }
func (rm CWSIZER_Mask) Store(b CWSIZER_Bits) { rm.UM32.Store(uint32(b)) }

type DR_Bits uint32

func (b DR_Bits) Field(mask DR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR_Bits) J(v int) DR_Bits {
	return DR_Bits(bits.Make32(v, uint32(mask)))
}

type DR struct{ mmio.U32 }

func (r *DR) Bits(mask DR_Bits) DR_Bits { return DR_Bits(r.U32.Bits(uint32(mask))) }
func (r *DR) StoreBits(mask, b DR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DR) SetBits(mask DR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DR) ClearBits(mask DR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DR) Load() DR_Bits             { return DR_Bits(r.U32.Load()) }
func (r *DR) Store(b DR_Bits)           { r.U32.Store(uint32(b)) }

func (r *DR) AtomicSetBits(mask DR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DR) AtomicClearBits(mask DR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DR_Mask struct{ mmio.UM32 }

func (rm DR_Mask) Load() DR_Bits   { return DR_Bits(rm.UM32.Load()) }
func (rm DR_Mask) Store(b DR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DCMI_Periph) BYTE0() DR_Mask {
	return DR_Mask{mmio.UM32{&p.DR.U32, uint32(BYTE0)}}
}

func (p *DCMI_Periph) BYTE1() DR_Mask {
	return DR_Mask{mmio.UM32{&p.DR.U32, uint32(BYTE1)}}
}

func (p *DCMI_Periph) BYTE2() DR_Mask {
	return DR_Mask{mmio.UM32{&p.DR.U32, uint32(BYTE2)}}
}

func (p *DCMI_Periph) BYTE3() DR_Mask {
	return DR_Mask{mmio.UM32{&p.DR.U32, uint32(BYTE3)}}
}
