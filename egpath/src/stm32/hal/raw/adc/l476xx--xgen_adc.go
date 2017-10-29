// +build l476xx

package adc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type ADC_Periph struct {
	ISR     ISR
	IER     IER
	CR      CR
	CFGR    CFGR
	CFGR2   CFGR2
	SMPR1   SMPR1
	SMPR2   SMPR2
	_       uint32
	TR1     TR1
	TR2     TR2
	TR3     TR3
	_       uint32
	SQR1    SQR1
	SQR2    SQR2
	SQR3    SQR3
	SQR4    SQR4
	DR      DR
	_       [2]uint32
	JSQR    JSQR
	_       [4]uint32
	OFR     [4]OFR
	_       [4]uint32
	JDR     [4]JDR
	_       [4]uint32
	AWD2CR  AWD2CR
	AWD3CR  AWD3CR
	_       [2]uint32
	DIFSEL  DIFSEL
	CALFACT CALFACT
}

func (p *ADC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var ADC1 = (*ADC_Periph)(unsafe.Pointer(uintptr(mmap.ADC1_BASE)))

//emgo:const
var ADC2 = (*ADC_Periph)(unsafe.Pointer(uintptr(mmap.ADC2_BASE)))

//emgo:const
var ADC3 = (*ADC_Periph)(unsafe.Pointer(uintptr(mmap.ADC3_BASE)))

type ISR_Bits uint32

func (b ISR_Bits) Field(mask ISR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ISR_Bits) J(v int) ISR_Bits {
	return ISR_Bits(bits.Make32(v, uint32(mask)))
}

type ISR struct{ mmio.U32 }

func (r *ISR) Bits(mask ISR_Bits) ISR_Bits { return ISR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ISR) StoreBits(mask, b ISR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ISR) SetBits(mask ISR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ISR) ClearBits(mask ISR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ISR) Load() ISR_Bits              { return ISR_Bits(r.U32.Load()) }
func (r *ISR) Store(b ISR_Bits)            { r.U32.Store(uint32(b)) }

func (r *ISR) AtomicSetBits(mask ISR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ISR) AtomicClearBits(mask ISR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type ISR_Mask struct{ mmio.UM32 }

func (rm ISR_Mask) Load() ISR_Bits   { return ISR_Bits(rm.UM32.Load()) }
func (rm ISR_Mask) Store(b ISR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) ADRDY() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(ADRDY)}}
}

func (p *ADC_Periph) EOSMP() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(EOSMP)}}
}

func (p *ADC_Periph) EOC() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(EOC)}}
}

func (p *ADC_Periph) EOS() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(EOS)}}
}

func (p *ADC_Periph) OVR() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(OVR)}}
}

func (p *ADC_Periph) JEOC() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(JEOC)}}
}

func (p *ADC_Periph) JEOS() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(JEOS)}}
}

func (p *ADC_Periph) AWD1() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(AWD1)}}
}

func (p *ADC_Periph) AWD2() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(AWD2)}}
}

func (p *ADC_Periph) AWD3() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(AWD3)}}
}

func (p *ADC_Periph) JQOVF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(JQOVF)}}
}

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

func (p *ADC_Periph) ADRDYIEIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(ADRDYIEIE)}}
}

func (p *ADC_Periph) EOSMPIEIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(EOSMPIEIE)}}
}

func (p *ADC_Periph) EOCIEIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(EOCIEIE)}}
}

func (p *ADC_Periph) EOSIEIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(EOSIEIE)}}
}

func (p *ADC_Periph) OVRIEIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(OVRIEIE)}}
}

func (p *ADC_Periph) JEOCIEIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(JEOCIEIE)}}
}

func (p *ADC_Periph) JEOSIEIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(JEOSIEIE)}}
}

func (p *ADC_Periph) AWD1IEIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(AWD1IEIE)}}
}

func (p *ADC_Periph) AWD2IEIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(AWD2IEIE)}}
}

func (p *ADC_Periph) AWD3IEIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(AWD3IEIE)}}
}

func (p *ADC_Periph) JQOVFIEIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(JQOVFIEIE)}}
}

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

func (p *ADC_Periph) ADEN() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ADEN)}}
}

func (p *ADC_Periph) ADDIS() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ADDIS)}}
}

func (p *ADC_Periph) ADSTART() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ADSTART)}}
}

func (p *ADC_Periph) JADSTART() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(JADSTART)}}
}

func (p *ADC_Periph) ADSTP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ADSTP)}}
}

func (p *ADC_Periph) JADSTP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(JADSTP)}}
}

func (p *ADC_Periph) ADVREGEN() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ADVREGEN)}}
}

func (p *ADC_Periph) DEEPPWD() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DEEPPWD)}}
}

func (p *ADC_Periph) ADCALDIF() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ADCALDIF)}}
}

func (p *ADC_Periph) ADCAL() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ADCAL)}}
}

type CFGR_Bits uint32

func (b CFGR_Bits) Field(mask CFGR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR_Bits) J(v int) CFGR_Bits {
	return CFGR_Bits(bits.Make32(v, uint32(mask)))
}

type CFGR struct{ mmio.U32 }

func (r *CFGR) Bits(mask CFGR_Bits) CFGR_Bits { return CFGR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CFGR) StoreBits(mask, b CFGR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CFGR) SetBits(mask CFGR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CFGR) ClearBits(mask CFGR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CFGR) Load() CFGR_Bits               { return CFGR_Bits(r.U32.Load()) }
func (r *CFGR) Store(b CFGR_Bits)             { r.U32.Store(uint32(b)) }

func (r *CFGR) AtomicSetBits(mask CFGR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CFGR) AtomicClearBits(mask CFGR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CFGR_Mask struct{ mmio.UM32 }

func (rm CFGR_Mask) Load() CFGR_Bits   { return CFGR_Bits(rm.UM32.Load()) }
func (rm CFGR_Mask) Store(b CFGR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) DMAEN() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(DMAEN)}}
}

func (p *ADC_Periph) DMACFG() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(DMACFG)}}
}

func (p *ADC_Periph) RES() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(RES)}}
}

func (p *ADC_Periph) ALIGN() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(ALIGN)}}
}

func (p *ADC_Periph) EXTSEL() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(EXTSEL)}}
}

func (p *ADC_Periph) EXTEN() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(EXTEN)}}
}

func (p *ADC_Periph) OVRMOD() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(OVRMOD)}}
}

func (p *ADC_Periph) CONT() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(CONT)}}
}

func (p *ADC_Periph) AUTDLY() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(AUTDLY)}}
}

func (p *ADC_Periph) DISCEN() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(DISCEN)}}
}

func (p *ADC_Periph) DISCNUM() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(DISCNUM)}}
}

func (p *ADC_Periph) JDISCEN() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(JDISCEN)}}
}

func (p *ADC_Periph) JQM() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(JQM)}}
}

func (p *ADC_Periph) AWD1SGL() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(AWD1SGL)}}
}

func (p *ADC_Periph) AWD1EN() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(AWD1EN)}}
}

func (p *ADC_Periph) JAWD1EN() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(JAWD1EN)}}
}

func (p *ADC_Periph) JAUTO() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(JAUTO)}}
}

func (p *ADC_Periph) AWD1CH() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(AWD1CH)}}
}

func (p *ADC_Periph) JQDIS() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(JQDIS)}}
}

type CFGR2_Bits uint32

func (b CFGR2_Bits) Field(mask CFGR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR2_Bits) J(v int) CFGR2_Bits {
	return CFGR2_Bits(bits.Make32(v, uint32(mask)))
}

type CFGR2 struct{ mmio.U32 }

func (r *CFGR2) Bits(mask CFGR2_Bits) CFGR2_Bits { return CFGR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *CFGR2) StoreBits(mask, b CFGR2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CFGR2) SetBits(mask CFGR2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CFGR2) ClearBits(mask CFGR2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CFGR2) Load() CFGR2_Bits                { return CFGR2_Bits(r.U32.Load()) }
func (r *CFGR2) Store(b CFGR2_Bits)              { r.U32.Store(uint32(b)) }

func (r *CFGR2) AtomicSetBits(mask CFGR2_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CFGR2) AtomicClearBits(mask CFGR2_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CFGR2_Mask struct{ mmio.UM32 }

func (rm CFGR2_Mask) Load() CFGR2_Bits   { return CFGR2_Bits(rm.UM32.Load()) }
func (rm CFGR2_Mask) Store(b CFGR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) ROVSE() CFGR2_Mask {
	return CFGR2_Mask{mmio.UM32{&p.CFGR2.U32, uint32(ROVSE)}}
}

func (p *ADC_Periph) JOVSE() CFGR2_Mask {
	return CFGR2_Mask{mmio.UM32{&p.CFGR2.U32, uint32(JOVSE)}}
}

func (p *ADC_Periph) OVSR() CFGR2_Mask {
	return CFGR2_Mask{mmio.UM32{&p.CFGR2.U32, uint32(OVSR)}}
}

func (p *ADC_Periph) OVSS() CFGR2_Mask {
	return CFGR2_Mask{mmio.UM32{&p.CFGR2.U32, uint32(OVSS)}}
}

func (p *ADC_Periph) TROVS() CFGR2_Mask {
	return CFGR2_Mask{mmio.UM32{&p.CFGR2.U32, uint32(TROVS)}}
}

func (p *ADC_Periph) ROVSM() CFGR2_Mask {
	return CFGR2_Mask{mmio.UM32{&p.CFGR2.U32, uint32(ROVSM)}}
}

type SMPR1_Bits uint32

func (b SMPR1_Bits) Field(mask SMPR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SMPR1_Bits) J(v int) SMPR1_Bits {
	return SMPR1_Bits(bits.Make32(v, uint32(mask)))
}

type SMPR1 struct{ mmio.U32 }

func (r *SMPR1) Bits(mask SMPR1_Bits) SMPR1_Bits { return SMPR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *SMPR1) StoreBits(mask, b SMPR1_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SMPR1) SetBits(mask SMPR1_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *SMPR1) ClearBits(mask SMPR1_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *SMPR1) Load() SMPR1_Bits                { return SMPR1_Bits(r.U32.Load()) }
func (r *SMPR1) Store(b SMPR1_Bits)              { r.U32.Store(uint32(b)) }

func (r *SMPR1) AtomicSetBits(mask SMPR1_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *SMPR1) AtomicClearBits(mask SMPR1_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type SMPR1_Mask struct{ mmio.UM32 }

func (rm SMPR1_Mask) Load() SMPR1_Bits   { return SMPR1_Bits(rm.UM32.Load()) }
func (rm SMPR1_Mask) Store(b SMPR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SMP0() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP0)}}
}

func (p *ADC_Periph) SMP1() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP1)}}
}

func (p *ADC_Periph) SMP2() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP2)}}
}

func (p *ADC_Periph) SMP3() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP3)}}
}

func (p *ADC_Periph) SMP4() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP4)}}
}

func (p *ADC_Periph) SMP5() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP5)}}
}

func (p *ADC_Periph) SMP6() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP6)}}
}

func (p *ADC_Periph) SMP7() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP7)}}
}

func (p *ADC_Periph) SMP8() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP8)}}
}

func (p *ADC_Periph) SMP9() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP9)}}
}

type SMPR2_Bits uint32

func (b SMPR2_Bits) Field(mask SMPR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SMPR2_Bits) J(v int) SMPR2_Bits {
	return SMPR2_Bits(bits.Make32(v, uint32(mask)))
}

type SMPR2 struct{ mmio.U32 }

func (r *SMPR2) Bits(mask SMPR2_Bits) SMPR2_Bits { return SMPR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *SMPR2) StoreBits(mask, b SMPR2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SMPR2) SetBits(mask SMPR2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *SMPR2) ClearBits(mask SMPR2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *SMPR2) Load() SMPR2_Bits                { return SMPR2_Bits(r.U32.Load()) }
func (r *SMPR2) Store(b SMPR2_Bits)              { r.U32.Store(uint32(b)) }

func (r *SMPR2) AtomicSetBits(mask SMPR2_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *SMPR2) AtomicClearBits(mask SMPR2_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type SMPR2_Mask struct{ mmio.UM32 }

func (rm SMPR2_Mask) Load() SMPR2_Bits   { return SMPR2_Bits(rm.UM32.Load()) }
func (rm SMPR2_Mask) Store(b SMPR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SMP10() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP10)}}
}

func (p *ADC_Periph) SMP11() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP11)}}
}

func (p *ADC_Periph) SMP12() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP12)}}
}

func (p *ADC_Periph) SMP13() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP13)}}
}

func (p *ADC_Periph) SMP14() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP14)}}
}

func (p *ADC_Periph) SMP15() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP15)}}
}

func (p *ADC_Periph) SMP16() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP16)}}
}

func (p *ADC_Periph) SMP17() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP17)}}
}

func (p *ADC_Periph) SMP18() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP18)}}
}

type TR1_Bits uint32

func (b TR1_Bits) Field(mask TR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TR1_Bits) J(v int) TR1_Bits {
	return TR1_Bits(bits.Make32(v, uint32(mask)))
}

type TR1 struct{ mmio.U32 }

func (r *TR1) Bits(mask TR1_Bits) TR1_Bits { return TR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *TR1) StoreBits(mask, b TR1_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TR1) SetBits(mask TR1_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *TR1) ClearBits(mask TR1_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *TR1) Load() TR1_Bits              { return TR1_Bits(r.U32.Load()) }
func (r *TR1) Store(b TR1_Bits)            { r.U32.Store(uint32(b)) }

func (r *TR1) AtomicSetBits(mask TR1_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *TR1) AtomicClearBits(mask TR1_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type TR1_Mask struct{ mmio.UM32 }

func (rm TR1_Mask) Load() TR1_Bits   { return TR1_Bits(rm.UM32.Load()) }
func (rm TR1_Mask) Store(b TR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) LT1() TR1_Mask {
	return TR1_Mask{mmio.UM32{&p.TR1.U32, uint32(LT1)}}
}

func (p *ADC_Periph) HT1() TR1_Mask {
	return TR1_Mask{mmio.UM32{&p.TR1.U32, uint32(HT1)}}
}

type TR2_Bits uint32

func (b TR2_Bits) Field(mask TR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TR2_Bits) J(v int) TR2_Bits {
	return TR2_Bits(bits.Make32(v, uint32(mask)))
}

type TR2 struct{ mmio.U32 }

func (r *TR2) Bits(mask TR2_Bits) TR2_Bits { return TR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *TR2) StoreBits(mask, b TR2_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TR2) SetBits(mask TR2_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *TR2) ClearBits(mask TR2_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *TR2) Load() TR2_Bits              { return TR2_Bits(r.U32.Load()) }
func (r *TR2) Store(b TR2_Bits)            { r.U32.Store(uint32(b)) }

func (r *TR2) AtomicSetBits(mask TR2_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *TR2) AtomicClearBits(mask TR2_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type TR2_Mask struct{ mmio.UM32 }

func (rm TR2_Mask) Load() TR2_Bits   { return TR2_Bits(rm.UM32.Load()) }
func (rm TR2_Mask) Store(b TR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) LT2() TR2_Mask {
	return TR2_Mask{mmio.UM32{&p.TR2.U32, uint32(LT2)}}
}

func (p *ADC_Periph) HT2() TR2_Mask {
	return TR2_Mask{mmio.UM32{&p.TR2.U32, uint32(HT2)}}
}

type TR3_Bits uint32

func (b TR3_Bits) Field(mask TR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TR3_Bits) J(v int) TR3_Bits {
	return TR3_Bits(bits.Make32(v, uint32(mask)))
}

type TR3 struct{ mmio.U32 }

func (r *TR3) Bits(mask TR3_Bits) TR3_Bits { return TR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *TR3) StoreBits(mask, b TR3_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TR3) SetBits(mask TR3_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *TR3) ClearBits(mask TR3_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *TR3) Load() TR3_Bits              { return TR3_Bits(r.U32.Load()) }
func (r *TR3) Store(b TR3_Bits)            { r.U32.Store(uint32(b)) }

func (r *TR3) AtomicSetBits(mask TR3_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *TR3) AtomicClearBits(mask TR3_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type TR3_Mask struct{ mmio.UM32 }

func (rm TR3_Mask) Load() TR3_Bits   { return TR3_Bits(rm.UM32.Load()) }
func (rm TR3_Mask) Store(b TR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) LT3() TR3_Mask {
	return TR3_Mask{mmio.UM32{&p.TR3.U32, uint32(LT3)}}
}

func (p *ADC_Periph) HT3() TR3_Mask {
	return TR3_Mask{mmio.UM32{&p.TR3.U32, uint32(HT3)}}
}

type SQR1_Bits uint32

func (b SQR1_Bits) Field(mask SQR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR1_Bits) J(v int) SQR1_Bits {
	return SQR1_Bits(bits.Make32(v, uint32(mask)))
}

type SQR1 struct{ mmio.U32 }

func (r *SQR1) Bits(mask SQR1_Bits) SQR1_Bits { return SQR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *SQR1) StoreBits(mask, b SQR1_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SQR1) SetBits(mask SQR1_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *SQR1) ClearBits(mask SQR1_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *SQR1) Load() SQR1_Bits               { return SQR1_Bits(r.U32.Load()) }
func (r *SQR1) Store(b SQR1_Bits)             { r.U32.Store(uint32(b)) }

func (r *SQR1) AtomicSetBits(mask SQR1_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *SQR1) AtomicClearBits(mask SQR1_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type SQR1_Mask struct{ mmio.UM32 }

func (rm SQR1_Mask) Load() SQR1_Bits   { return SQR1_Bits(rm.UM32.Load()) }
func (rm SQR1_Mask) Store(b SQR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) L() SQR1_Mask {
	return SQR1_Mask{mmio.UM32{&p.SQR1.U32, uint32(L)}}
}

func (p *ADC_Periph) SQ1() SQR1_Mask {
	return SQR1_Mask{mmio.UM32{&p.SQR1.U32, uint32(SQ1)}}
}

func (p *ADC_Periph) SQ2() SQR1_Mask {
	return SQR1_Mask{mmio.UM32{&p.SQR1.U32, uint32(SQ2)}}
}

func (p *ADC_Periph) SQ3() SQR1_Mask {
	return SQR1_Mask{mmio.UM32{&p.SQR1.U32, uint32(SQ3)}}
}

func (p *ADC_Periph) SQ4() SQR1_Mask {
	return SQR1_Mask{mmio.UM32{&p.SQR1.U32, uint32(SQ4)}}
}

type SQR2_Bits uint32

func (b SQR2_Bits) Field(mask SQR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR2_Bits) J(v int) SQR2_Bits {
	return SQR2_Bits(bits.Make32(v, uint32(mask)))
}

type SQR2 struct{ mmio.U32 }

func (r *SQR2) Bits(mask SQR2_Bits) SQR2_Bits { return SQR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *SQR2) StoreBits(mask, b SQR2_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SQR2) SetBits(mask SQR2_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *SQR2) ClearBits(mask SQR2_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *SQR2) Load() SQR2_Bits               { return SQR2_Bits(r.U32.Load()) }
func (r *SQR2) Store(b SQR2_Bits)             { r.U32.Store(uint32(b)) }

func (r *SQR2) AtomicSetBits(mask SQR2_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *SQR2) AtomicClearBits(mask SQR2_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type SQR2_Mask struct{ mmio.UM32 }

func (rm SQR2_Mask) Load() SQR2_Bits   { return SQR2_Bits(rm.UM32.Load()) }
func (rm SQR2_Mask) Store(b SQR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SQ5() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{&p.SQR2.U32, uint32(SQ5)}}
}

func (p *ADC_Periph) SQ6() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{&p.SQR2.U32, uint32(SQ6)}}
}

func (p *ADC_Periph) SQ7() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{&p.SQR2.U32, uint32(SQ7)}}
}

func (p *ADC_Periph) SQ8() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{&p.SQR2.U32, uint32(SQ8)}}
}

func (p *ADC_Periph) SQ9() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{&p.SQR2.U32, uint32(SQ9)}}
}

type SQR3_Bits uint32

func (b SQR3_Bits) Field(mask SQR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR3_Bits) J(v int) SQR3_Bits {
	return SQR3_Bits(bits.Make32(v, uint32(mask)))
}

type SQR3 struct{ mmio.U32 }

func (r *SQR3) Bits(mask SQR3_Bits) SQR3_Bits { return SQR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *SQR3) StoreBits(mask, b SQR3_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SQR3) SetBits(mask SQR3_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *SQR3) ClearBits(mask SQR3_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *SQR3) Load() SQR3_Bits               { return SQR3_Bits(r.U32.Load()) }
func (r *SQR3) Store(b SQR3_Bits)             { r.U32.Store(uint32(b)) }

func (r *SQR3) AtomicSetBits(mask SQR3_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *SQR3) AtomicClearBits(mask SQR3_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type SQR3_Mask struct{ mmio.UM32 }

func (rm SQR3_Mask) Load() SQR3_Bits   { return SQR3_Bits(rm.UM32.Load()) }
func (rm SQR3_Mask) Store(b SQR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SQ10() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{&p.SQR3.U32, uint32(SQ10)}}
}

func (p *ADC_Periph) SQ11() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{&p.SQR3.U32, uint32(SQ11)}}
}

func (p *ADC_Periph) SQ12() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{&p.SQR3.U32, uint32(SQ12)}}
}

func (p *ADC_Periph) SQ13() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{&p.SQR3.U32, uint32(SQ13)}}
}

func (p *ADC_Periph) SQ14() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{&p.SQR3.U32, uint32(SQ14)}}
}

type SQR4_Bits uint32

func (b SQR4_Bits) Field(mask SQR4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR4_Bits) J(v int) SQR4_Bits {
	return SQR4_Bits(bits.Make32(v, uint32(mask)))
}

type SQR4 struct{ mmio.U32 }

func (r *SQR4) Bits(mask SQR4_Bits) SQR4_Bits { return SQR4_Bits(r.U32.Bits(uint32(mask))) }
func (r *SQR4) StoreBits(mask, b SQR4_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SQR4) SetBits(mask SQR4_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *SQR4) ClearBits(mask SQR4_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *SQR4) Load() SQR4_Bits               { return SQR4_Bits(r.U32.Load()) }
func (r *SQR4) Store(b SQR4_Bits)             { r.U32.Store(uint32(b)) }

func (r *SQR4) AtomicSetBits(mask SQR4_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *SQR4) AtomicClearBits(mask SQR4_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type SQR4_Mask struct{ mmio.UM32 }

func (rm SQR4_Mask) Load() SQR4_Bits   { return SQR4_Bits(rm.UM32.Load()) }
func (rm SQR4_Mask) Store(b SQR4_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SQ15() SQR4_Mask {
	return SQR4_Mask{mmio.UM32{&p.SQR4.U32, uint32(SQ15)}}
}

func (p *ADC_Periph) SQ16() SQR4_Mask {
	return SQR4_Mask{mmio.UM32{&p.SQR4.U32, uint32(SQ16)}}
}

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

func (p *ADC_Periph) RDATA() DR_Mask {
	return DR_Mask{mmio.UM32{&p.DR.U32, uint32(RDATA)}}
}

type JSQR_Bits uint32

func (b JSQR_Bits) Field(mask JSQR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JSQR_Bits) J(v int) JSQR_Bits {
	return JSQR_Bits(bits.Make32(v, uint32(mask)))
}

type JSQR struct{ mmio.U32 }

func (r *JSQR) Bits(mask JSQR_Bits) JSQR_Bits { return JSQR_Bits(r.U32.Bits(uint32(mask))) }
func (r *JSQR) StoreBits(mask, b JSQR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JSQR) SetBits(mask JSQR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *JSQR) ClearBits(mask JSQR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *JSQR) Load() JSQR_Bits               { return JSQR_Bits(r.U32.Load()) }
func (r *JSQR) Store(b JSQR_Bits)             { r.U32.Store(uint32(b)) }

func (r *JSQR) AtomicSetBits(mask JSQR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *JSQR) AtomicClearBits(mask JSQR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type JSQR_Mask struct{ mmio.UM32 }

func (rm JSQR_Mask) Load() JSQR_Bits   { return JSQR_Bits(rm.UM32.Load()) }
func (rm JSQR_Mask) Store(b JSQR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JL() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{&p.JSQR.U32, uint32(JL)}}
}

func (p *ADC_Periph) JEXTSEL() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{&p.JSQR.U32, uint32(JEXTSEL)}}
}

func (p *ADC_Periph) JEXTEN() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{&p.JSQR.U32, uint32(JEXTEN)}}
}

func (p *ADC_Periph) JSQ1() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{&p.JSQR.U32, uint32(JSQ1)}}
}

func (p *ADC_Periph) JSQ2() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{&p.JSQR.U32, uint32(JSQ2)}}
}

func (p *ADC_Periph) JSQ3() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{&p.JSQR.U32, uint32(JSQ3)}}
}

func (p *ADC_Periph) JSQ4() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{&p.JSQR.U32, uint32(JSQ4)}}
}

type OFR_Bits uint32

func (b OFR_Bits) Field(mask OFR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OFR_Bits) J(v int) OFR_Bits {
	return OFR_Bits(bits.Make32(v, uint32(mask)))
}

type OFR struct{ mmio.U32 }

func (r *OFR) Bits(mask OFR_Bits) OFR_Bits { return OFR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OFR) StoreBits(mask, b OFR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OFR) SetBits(mask OFR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *OFR) ClearBits(mask OFR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *OFR) Load() OFR_Bits              { return OFR_Bits(r.U32.Load()) }
func (r *OFR) Store(b OFR_Bits)            { r.U32.Store(uint32(b)) }

func (r *OFR) AtomicSetBits(mask OFR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *OFR) AtomicClearBits(mask OFR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type OFR_Mask struct{ mmio.UM32 }

func (rm OFR_Mask) Load() OFR_Bits   { return OFR_Bits(rm.UM32.Load()) }
func (rm OFR_Mask) Store(b OFR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) OFFSET1(n int) OFR_Mask {
	return OFR_Mask{mmio.UM32{&p.OFR[n].U32, uint32(OFFSET1)}}
}

func (p *ADC_Periph) OFFSET1_CH(n int) OFR_Mask {
	return OFR_Mask{mmio.UM32{&p.OFR[n].U32, uint32(OFFSET1_CH)}}
}

func (p *ADC_Periph) OFFSET1_EN(n int) OFR_Mask {
	return OFR_Mask{mmio.UM32{&p.OFR[n].U32, uint32(OFFSET1_EN)}}
}

type JDR_Bits uint32

func (b JDR_Bits) Field(mask JDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JDR_Bits) J(v int) JDR_Bits {
	return JDR_Bits(bits.Make32(v, uint32(mask)))
}

type JDR struct{ mmio.U32 }

func (r *JDR) Bits(mask JDR_Bits) JDR_Bits { return JDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *JDR) StoreBits(mask, b JDR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JDR) SetBits(mask JDR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *JDR) ClearBits(mask JDR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *JDR) Load() JDR_Bits              { return JDR_Bits(r.U32.Load()) }
func (r *JDR) Store(b JDR_Bits)            { r.U32.Store(uint32(b)) }

func (r *JDR) AtomicSetBits(mask JDR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *JDR) AtomicClearBits(mask JDR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type JDR_Mask struct{ mmio.UM32 }

func (rm JDR_Mask) Load() JDR_Bits   { return JDR_Bits(rm.UM32.Load()) }
func (rm JDR_Mask) Store(b JDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JDATA(n int) JDR_Mask {
	return JDR_Mask{mmio.UM32{&p.JDR[n].U32, uint32(JDATA)}}
}

type AWD2CR_Bits uint32

func (b AWD2CR_Bits) Field(mask AWD2CR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AWD2CR_Bits) J(v int) AWD2CR_Bits {
	return AWD2CR_Bits(bits.Make32(v, uint32(mask)))
}

type AWD2CR struct{ mmio.U32 }

func (r *AWD2CR) Bits(mask AWD2CR_Bits) AWD2CR_Bits { return AWD2CR_Bits(r.U32.Bits(uint32(mask))) }
func (r *AWD2CR) StoreBits(mask, b AWD2CR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *AWD2CR) SetBits(mask AWD2CR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *AWD2CR) ClearBits(mask AWD2CR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *AWD2CR) Load() AWD2CR_Bits                 { return AWD2CR_Bits(r.U32.Load()) }
func (r *AWD2CR) Store(b AWD2CR_Bits)               { r.U32.Store(uint32(b)) }

func (r *AWD2CR) AtomicSetBits(mask AWD2CR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *AWD2CR) AtomicClearBits(mask AWD2CR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type AWD2CR_Mask struct{ mmio.UM32 }

func (rm AWD2CR_Mask) Load() AWD2CR_Bits   { return AWD2CR_Bits(rm.UM32.Load()) }
func (rm AWD2CR_Mask) Store(b AWD2CR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) AWD2CH() AWD2CR_Mask {
	return AWD2CR_Mask{mmio.UM32{&p.AWD2CR.U32, uint32(AWD2CH)}}
}

type AWD3CR_Bits uint32

func (b AWD3CR_Bits) Field(mask AWD3CR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AWD3CR_Bits) J(v int) AWD3CR_Bits {
	return AWD3CR_Bits(bits.Make32(v, uint32(mask)))
}

type AWD3CR struct{ mmio.U32 }

func (r *AWD3CR) Bits(mask AWD3CR_Bits) AWD3CR_Bits { return AWD3CR_Bits(r.U32.Bits(uint32(mask))) }
func (r *AWD3CR) StoreBits(mask, b AWD3CR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *AWD3CR) SetBits(mask AWD3CR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *AWD3CR) ClearBits(mask AWD3CR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *AWD3CR) Load() AWD3CR_Bits                 { return AWD3CR_Bits(r.U32.Load()) }
func (r *AWD3CR) Store(b AWD3CR_Bits)               { r.U32.Store(uint32(b)) }

func (r *AWD3CR) AtomicSetBits(mask AWD3CR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *AWD3CR) AtomicClearBits(mask AWD3CR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type AWD3CR_Mask struct{ mmio.UM32 }

func (rm AWD3CR_Mask) Load() AWD3CR_Bits   { return AWD3CR_Bits(rm.UM32.Load()) }
func (rm AWD3CR_Mask) Store(b AWD3CR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) AWD3CH() AWD3CR_Mask {
	return AWD3CR_Mask{mmio.UM32{&p.AWD3CR.U32, uint32(AWD3CH)}}
}

type DIFSEL_Bits uint32

func (b DIFSEL_Bits) Field(mask DIFSEL_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIFSEL_Bits) J(v int) DIFSEL_Bits {
	return DIFSEL_Bits(bits.Make32(v, uint32(mask)))
}

type DIFSEL struct{ mmio.U32 }

func (r *DIFSEL) Bits(mask DIFSEL_Bits) DIFSEL_Bits { return DIFSEL_Bits(r.U32.Bits(uint32(mask))) }
func (r *DIFSEL) StoreBits(mask, b DIFSEL_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DIFSEL) SetBits(mask DIFSEL_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *DIFSEL) ClearBits(mask DIFSEL_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *DIFSEL) Load() DIFSEL_Bits                 { return DIFSEL_Bits(r.U32.Load()) }
func (r *DIFSEL) Store(b DIFSEL_Bits)               { r.U32.Store(uint32(b)) }

func (r *DIFSEL) AtomicSetBits(mask DIFSEL_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DIFSEL) AtomicClearBits(mask DIFSEL_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DIFSEL_Mask struct{ mmio.UM32 }

func (rm DIFSEL_Mask) Load() DIFSEL_Bits   { return DIFSEL_Bits(rm.UM32.Load()) }
func (rm DIFSEL_Mask) Store(b DIFSEL_Bits) { rm.UM32.Store(uint32(b)) }

type CALFACT_Bits uint32

func (b CALFACT_Bits) Field(mask CALFACT_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CALFACT_Bits) J(v int) CALFACT_Bits {
	return CALFACT_Bits(bits.Make32(v, uint32(mask)))
}

type CALFACT struct{ mmio.U32 }

func (r *CALFACT) Bits(mask CALFACT_Bits) CALFACT_Bits { return CALFACT_Bits(r.U32.Bits(uint32(mask))) }
func (r *CALFACT) StoreBits(mask, b CALFACT_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CALFACT) SetBits(mask CALFACT_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *CALFACT) ClearBits(mask CALFACT_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *CALFACT) Load() CALFACT_Bits                  { return CALFACT_Bits(r.U32.Load()) }
func (r *CALFACT) Store(b CALFACT_Bits)                { r.U32.Store(uint32(b)) }

func (r *CALFACT) AtomicSetBits(mask CALFACT_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CALFACT) AtomicClearBits(mask CALFACT_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CALFACT_Mask struct{ mmio.UM32 }

func (rm CALFACT_Mask) Load() CALFACT_Bits   { return CALFACT_Bits(rm.UM32.Load()) }
func (rm CALFACT_Mask) Store(b CALFACT_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) CALFACT_S() CALFACT_Mask {
	return CALFACT_Mask{mmio.UM32{&p.CALFACT.U32, uint32(CALFACT_S)}}
}

func (p *ADC_Periph) CALFACT_D() CALFACT_Mask {
	return CALFACT_Mask{mmio.UM32{&p.CALFACT.U32, uint32(CALFACT_D)}}
}
