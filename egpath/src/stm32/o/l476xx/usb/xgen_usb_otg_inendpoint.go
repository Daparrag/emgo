package usb

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type USB_OTG_INEndpoint_Periph struct {
	DIEPCTL  DIEPCTL
	_        uint32
	DIEPINT  DIEPINT
	_        uint32
	DIEPTSIZ DIEPTSIZ
	DIEPDMA  DIEPDMA
	DTXFSTS  DTXFSTS
}

func (p *USB_OTG_INEndpoint_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type DIEPCTL_Bits uint32

func (b DIEPCTL_Bits) Field(mask DIEPCTL_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIEPCTL_Bits) J(v int) DIEPCTL_Bits {
	return DIEPCTL_Bits(bits.Make32(v, uint32(mask)))
}

type DIEPCTL struct{ mmio.U32 }

func (r *DIEPCTL) Bits(mask DIEPCTL_Bits) DIEPCTL_Bits { return DIEPCTL_Bits(r.U32.Bits(uint32(mask))) }
func (r *DIEPCTL) StoreBits(mask, b DIEPCTL_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DIEPCTL) SetBits(mask DIEPCTL_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DIEPCTL) ClearBits(mask DIEPCTL_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DIEPCTL) Load() DIEPCTL_Bits                  { return DIEPCTL_Bits(r.U32.Load()) }
func (r *DIEPCTL) Store(b DIEPCTL_Bits)                { r.U32.Store(uint32(b)) }

func (r *DIEPCTL) AtomicSetBits(mask DIEPCTL_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DIEPCTL) AtomicClearBits(mask DIEPCTL_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DIEPCTL_Mask struct{ mmio.UM32 }

func (rm DIEPCTL_Mask) Load() DIEPCTL_Bits   { return DIEPCTL_Bits(rm.UM32.Load()) }
func (rm DIEPCTL_Mask) Store(b DIEPCTL_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_INEndpoint_Periph) MPSIZ() DIEPCTL_Mask {
	return DIEPCTL_Mask{mmio.UM32{&p.DIEPCTL.U32, uint32(MPSIZ)}}
}

func (p *USB_OTG_INEndpoint_Periph) USBAEP() DIEPCTL_Mask {
	return DIEPCTL_Mask{mmio.UM32{&p.DIEPCTL.U32, uint32(USBAEP)}}
}

func (p *USB_OTG_INEndpoint_Periph) EONUM_DPID() DIEPCTL_Mask {
	return DIEPCTL_Mask{mmio.UM32{&p.DIEPCTL.U32, uint32(EONUM_DPID)}}
}

func (p *USB_OTG_INEndpoint_Periph) NAKSTS() DIEPCTL_Mask {
	return DIEPCTL_Mask{mmio.UM32{&p.DIEPCTL.U32, uint32(NAKSTS)}}
}

func (p *USB_OTG_INEndpoint_Periph) EPTYP() DIEPCTL_Mask {
	return DIEPCTL_Mask{mmio.UM32{&p.DIEPCTL.U32, uint32(EPTYP)}}
}

func (p *USB_OTG_INEndpoint_Periph) STALL() DIEPCTL_Mask {
	return DIEPCTL_Mask{mmio.UM32{&p.DIEPCTL.U32, uint32(STALL)}}
}

func (p *USB_OTG_INEndpoint_Periph) TXFNUM() DIEPCTL_Mask {
	return DIEPCTL_Mask{mmio.UM32{&p.DIEPCTL.U32, uint32(TXFNUM)}}
}

func (p *USB_OTG_INEndpoint_Periph) CNAK() DIEPCTL_Mask {
	return DIEPCTL_Mask{mmio.UM32{&p.DIEPCTL.U32, uint32(CNAK)}}
}

func (p *USB_OTG_INEndpoint_Periph) SNAK() DIEPCTL_Mask {
	return DIEPCTL_Mask{mmio.UM32{&p.DIEPCTL.U32, uint32(SNAK)}}
}

func (p *USB_OTG_INEndpoint_Periph) SD0PID_SEVNFRM() DIEPCTL_Mask {
	return DIEPCTL_Mask{mmio.UM32{&p.DIEPCTL.U32, uint32(SD0PID_SEVNFRM)}}
}

func (p *USB_OTG_INEndpoint_Periph) SODDFRM() DIEPCTL_Mask {
	return DIEPCTL_Mask{mmio.UM32{&p.DIEPCTL.U32, uint32(SODDFRM)}}
}

func (p *USB_OTG_INEndpoint_Periph) EPDIS() DIEPCTL_Mask {
	return DIEPCTL_Mask{mmio.UM32{&p.DIEPCTL.U32, uint32(EPDIS)}}
}

func (p *USB_OTG_INEndpoint_Periph) EPENA() DIEPCTL_Mask {
	return DIEPCTL_Mask{mmio.UM32{&p.DIEPCTL.U32, uint32(EPENA)}}
}

type DIEPINT_Bits uint32

func (b DIEPINT_Bits) Field(mask DIEPINT_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIEPINT_Bits) J(v int) DIEPINT_Bits {
	return DIEPINT_Bits(bits.Make32(v, uint32(mask)))
}

type DIEPINT struct{ mmio.U32 }

func (r *DIEPINT) Bits(mask DIEPINT_Bits) DIEPINT_Bits { return DIEPINT_Bits(r.U32.Bits(uint32(mask))) }
func (r *DIEPINT) StoreBits(mask, b DIEPINT_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DIEPINT) SetBits(mask DIEPINT_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DIEPINT) ClearBits(mask DIEPINT_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DIEPINT) Load() DIEPINT_Bits                  { return DIEPINT_Bits(r.U32.Load()) }
func (r *DIEPINT) Store(b DIEPINT_Bits)                { r.U32.Store(uint32(b)) }

func (r *DIEPINT) AtomicSetBits(mask DIEPINT_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DIEPINT) AtomicClearBits(mask DIEPINT_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DIEPINT_Mask struct{ mmio.UM32 }

func (rm DIEPINT_Mask) Load() DIEPINT_Bits   { return DIEPINT_Bits(rm.UM32.Load()) }
func (rm DIEPINT_Mask) Store(b DIEPINT_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_INEndpoint_Periph) XFRC() DIEPINT_Mask {
	return DIEPINT_Mask{mmio.UM32{&p.DIEPINT.U32, uint32(XFRC)}}
}

func (p *USB_OTG_INEndpoint_Periph) EPDISD() DIEPINT_Mask {
	return DIEPINT_Mask{mmio.UM32{&p.DIEPINT.U32, uint32(EPDISD)}}
}

func (p *USB_OTG_INEndpoint_Periph) TOC() DIEPINT_Mask {
	return DIEPINT_Mask{mmio.UM32{&p.DIEPINT.U32, uint32(TOC)}}
}

func (p *USB_OTG_INEndpoint_Periph) ITTXFE() DIEPINT_Mask {
	return DIEPINT_Mask{mmio.UM32{&p.DIEPINT.U32, uint32(ITTXFE)}}
}

func (p *USB_OTG_INEndpoint_Periph) INEPNE() DIEPINT_Mask {
	return DIEPINT_Mask{mmio.UM32{&p.DIEPINT.U32, uint32(INEPNE)}}
}

func (p *USB_OTG_INEndpoint_Periph) TXFE() DIEPINT_Mask {
	return DIEPINT_Mask{mmio.UM32{&p.DIEPINT.U32, uint32(TXFE)}}
}

func (p *USB_OTG_INEndpoint_Periph) TXFIFOUDRN() DIEPINT_Mask {
	return DIEPINT_Mask{mmio.UM32{&p.DIEPINT.U32, uint32(TXFIFOUDRN)}}
}

func (p *USB_OTG_INEndpoint_Periph) BNA() DIEPINT_Mask {
	return DIEPINT_Mask{mmio.UM32{&p.DIEPINT.U32, uint32(BNA)}}
}

func (p *USB_OTG_INEndpoint_Periph) PKTDRPSTS() DIEPINT_Mask {
	return DIEPINT_Mask{mmio.UM32{&p.DIEPINT.U32, uint32(PKTDRPSTS)}}
}

func (p *USB_OTG_INEndpoint_Periph) BERR() DIEPINT_Mask {
	return DIEPINT_Mask{mmio.UM32{&p.DIEPINT.U32, uint32(BERR)}}
}

func (p *USB_OTG_INEndpoint_Periph) NAK() DIEPINT_Mask {
	return DIEPINT_Mask{mmio.UM32{&p.DIEPINT.U32, uint32(NAK)}}
}

type DIEPTSIZ_Bits uint32

func (b DIEPTSIZ_Bits) Field(mask DIEPTSIZ_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIEPTSIZ_Bits) J(v int) DIEPTSIZ_Bits {
	return DIEPTSIZ_Bits(bits.Make32(v, uint32(mask)))
}

type DIEPTSIZ struct{ mmio.U32 }

func (r *DIEPTSIZ) Bits(mask DIEPTSIZ_Bits) DIEPTSIZ_Bits {
	return DIEPTSIZ_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DIEPTSIZ) StoreBits(mask, b DIEPTSIZ_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DIEPTSIZ) SetBits(mask DIEPTSIZ_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DIEPTSIZ) ClearBits(mask DIEPTSIZ_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DIEPTSIZ) Load() DIEPTSIZ_Bits             { return DIEPTSIZ_Bits(r.U32.Load()) }
func (r *DIEPTSIZ) Store(b DIEPTSIZ_Bits)           { r.U32.Store(uint32(b)) }

func (r *DIEPTSIZ) AtomicSetBits(mask DIEPTSIZ_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DIEPTSIZ) AtomicClearBits(mask DIEPTSIZ_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DIEPTSIZ_Mask struct{ mmio.UM32 }

func (rm DIEPTSIZ_Mask) Load() DIEPTSIZ_Bits   { return DIEPTSIZ_Bits(rm.UM32.Load()) }
func (rm DIEPTSIZ_Mask) Store(b DIEPTSIZ_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_INEndpoint_Periph) XFRSIZ() DIEPTSIZ_Mask {
	return DIEPTSIZ_Mask{mmio.UM32{&p.DIEPTSIZ.U32, uint32(XFRSIZ)}}
}

func (p *USB_OTG_INEndpoint_Periph) PKTCNT() DIEPTSIZ_Mask {
	return DIEPTSIZ_Mask{mmio.UM32{&p.DIEPTSIZ.U32, uint32(PKTCNT)}}
}

func (p *USB_OTG_INEndpoint_Periph) MULCNT() DIEPTSIZ_Mask {
	return DIEPTSIZ_Mask{mmio.UM32{&p.DIEPTSIZ.U32, uint32(MULCNT)}}
}

type DIEPDMA_Bits uint32

func (b DIEPDMA_Bits) Field(mask DIEPDMA_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIEPDMA_Bits) J(v int) DIEPDMA_Bits {
	return DIEPDMA_Bits(bits.Make32(v, uint32(mask)))
}

type DIEPDMA struct{ mmio.U32 }

func (r *DIEPDMA) Bits(mask DIEPDMA_Bits) DIEPDMA_Bits { return DIEPDMA_Bits(r.U32.Bits(uint32(mask))) }
func (r *DIEPDMA) StoreBits(mask, b DIEPDMA_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DIEPDMA) SetBits(mask DIEPDMA_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DIEPDMA) ClearBits(mask DIEPDMA_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DIEPDMA) Load() DIEPDMA_Bits                  { return DIEPDMA_Bits(r.U32.Load()) }
func (r *DIEPDMA) Store(b DIEPDMA_Bits)                { r.U32.Store(uint32(b)) }

func (r *DIEPDMA) AtomicSetBits(mask DIEPDMA_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DIEPDMA) AtomicClearBits(mask DIEPDMA_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DIEPDMA_Mask struct{ mmio.UM32 }

func (rm DIEPDMA_Mask) Load() DIEPDMA_Bits   { return DIEPDMA_Bits(rm.UM32.Load()) }
func (rm DIEPDMA_Mask) Store(b DIEPDMA_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_INEndpoint_Periph) DMAADDR() DIEPDMA_Mask {
	return DIEPDMA_Mask{mmio.UM32{&p.DIEPDMA.U32, uint32(DMAADDR)}}
}

type DTXFSTS_Bits uint32

func (b DTXFSTS_Bits) Field(mask DTXFSTS_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DTXFSTS_Bits) J(v int) DTXFSTS_Bits {
	return DTXFSTS_Bits(bits.Make32(v, uint32(mask)))
}

type DTXFSTS struct{ mmio.U32 }

func (r *DTXFSTS) Bits(mask DTXFSTS_Bits) DTXFSTS_Bits { return DTXFSTS_Bits(r.U32.Bits(uint32(mask))) }
func (r *DTXFSTS) StoreBits(mask, b DTXFSTS_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DTXFSTS) SetBits(mask DTXFSTS_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DTXFSTS) ClearBits(mask DTXFSTS_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DTXFSTS) Load() DTXFSTS_Bits                  { return DTXFSTS_Bits(r.U32.Load()) }
func (r *DTXFSTS) Store(b DTXFSTS_Bits)                { r.U32.Store(uint32(b)) }

func (r *DTXFSTS) AtomicSetBits(mask DTXFSTS_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DTXFSTS) AtomicClearBits(mask DTXFSTS_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DTXFSTS_Mask struct{ mmio.UM32 }

func (rm DTXFSTS_Mask) Load() DTXFSTS_Bits   { return DTXFSTS_Bits(rm.UM32.Load()) }
func (rm DTXFSTS_Mask) Store(b DTXFSTS_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_INEndpoint_Periph) INEPTFSAV() DTXFSTS_Mask {
	return DTXFSTS_Mask{mmio.UM32{&p.DTXFSTS.U32, uint32(INEPTFSAV)}}
}
