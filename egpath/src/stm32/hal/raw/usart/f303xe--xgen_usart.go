// +build f303xe

package usart

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f303xe/mmap"
)

type USART_Periph struct {
	CR1  CR1
	CR2  CR2
	CR3  CR3
	BRR  BRR
	_    uint16
	GTPR GTPR
	_    uint16
	RTOR RTOR
	RQR  RQR
	_    uint16
	ISR  ISR
	ICR  ICR
	RDR  RDR
	_    uint16
	TDR  TDR
}

func (p *USART_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var USART2 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.USART2_BASE)))

//emgo:const
var USART3 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.USART3_BASE)))

//emgo:const
var UART4 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.UART4_BASE)))

//emgo:const
var UART5 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.UART5_BASE)))

//emgo:const
var USART1 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.USART1_BASE)))

type CR1_Bits uint32

func (b CR1_Bits) Field(mask CR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR1_Bits) J(v int) CR1_Bits {
	return CR1_Bits(bits.Make32(v, uint32(mask)))
}

type CR1 struct{ mmio.U32 }

func (r *CR1) Bits(mask CR1_Bits) CR1_Bits { return CR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR1) StoreBits(mask, b CR1_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR1) SetBits(mask CR1_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CR1) ClearBits(mask CR1_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CR1) Load() CR1_Bits              { return CR1_Bits(r.U32.Load()) }
func (r *CR1) Store(b CR1_Bits)            { r.U32.Store(uint32(b)) }

func (r *CR1) AtomicSetBits(mask CR1_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CR1) AtomicClearBits(mask CR1_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CR1_Mask struct{ mmio.UM32 }

func (rm CR1_Mask) Load() CR1_Bits   { return CR1_Bits(rm.UM32.Load()) }
func (rm CR1_Mask) Store(b CR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USART_Periph) UE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(UE)}}
}

func (p *USART_Periph) UESM() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(UESM)}}
}

func (p *USART_Periph) RE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(RE)}}
}

func (p *USART_Periph) TE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(TE)}}
}

func (p *USART_Periph) IDLEIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(IDLEIE)}}
}

func (p *USART_Periph) RXNEIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(RXNEIE)}}
}

func (p *USART_Periph) TCIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(TCIE)}}
}

func (p *USART_Periph) TXEIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(TXEIE)}}
}

func (p *USART_Periph) PEIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(PEIE)}}
}

func (p *USART_Periph) PS() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(PS)}}
}

func (p *USART_Periph) PCE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(PCE)}}
}

func (p *USART_Periph) WAKE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(WAKE)}}
}

func (p *USART_Periph) M() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(M)}}
}

func (p *USART_Periph) MME() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(MME)}}
}

func (p *USART_Periph) CMIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(CMIE)}}
}

func (p *USART_Periph) OVER8() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(OVER8)}}
}

func (p *USART_Periph) DEDT() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(DEDT)}}
}

func (p *USART_Periph) DEAT() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(DEAT)}}
}

func (p *USART_Periph) RTOIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(RTOIE)}}
}

func (p *USART_Periph) EOBIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(EOBIE)}}
}

type CR2_Bits uint32

func (b CR2_Bits) Field(mask CR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR2_Bits) J(v int) CR2_Bits {
	return CR2_Bits(bits.Make32(v, uint32(mask)))
}

type CR2 struct{ mmio.U32 }

func (r *CR2) Bits(mask CR2_Bits) CR2_Bits { return CR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR2) StoreBits(mask, b CR2_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR2) SetBits(mask CR2_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CR2) ClearBits(mask CR2_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CR2) Load() CR2_Bits              { return CR2_Bits(r.U32.Load()) }
func (r *CR2) Store(b CR2_Bits)            { r.U32.Store(uint32(b)) }

func (r *CR2) AtomicSetBits(mask CR2_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CR2) AtomicClearBits(mask CR2_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CR2_Mask struct{ mmio.UM32 }

func (rm CR2_Mask) Load() CR2_Bits   { return CR2_Bits(rm.UM32.Load()) }
func (rm CR2_Mask) Store(b CR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USART_Periph) ADDM7() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(ADDM7)}}
}

func (p *USART_Periph) LBDL() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(LBDL)}}
}

func (p *USART_Periph) LBDIE() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(LBDIE)}}
}

func (p *USART_Periph) LBCL() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(LBCL)}}
}

func (p *USART_Periph) CPHA() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(CPHA)}}
}

func (p *USART_Periph) CPOL() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(CPOL)}}
}

func (p *USART_Periph) CLKEN() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(CLKEN)}}
}

func (p *USART_Periph) STOP() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(STOP)}}
}

func (p *USART_Periph) LINEN() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(LINEN)}}
}

func (p *USART_Periph) SWAP() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(SWAP)}}
}

func (p *USART_Periph) RXINV() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(RXINV)}}
}

func (p *USART_Periph) TXINV() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(TXINV)}}
}

func (p *USART_Periph) DATAINV() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(DATAINV)}}
}

func (p *USART_Periph) MSBFIRST() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(MSBFIRST)}}
}

func (p *USART_Periph) ABREN() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(ABREN)}}
}

func (p *USART_Periph) ABRMODE() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(ABRMODE)}}
}

func (p *USART_Periph) RTOEN() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(RTOEN)}}
}

func (p *USART_Periph) ADD() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(ADD)}}
}

type CR3_Bits uint32

func (b CR3_Bits) Field(mask CR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR3_Bits) J(v int) CR3_Bits {
	return CR3_Bits(bits.Make32(v, uint32(mask)))
}

type CR3 struct{ mmio.U32 }

func (r *CR3) Bits(mask CR3_Bits) CR3_Bits { return CR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR3) StoreBits(mask, b CR3_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR3) SetBits(mask CR3_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CR3) ClearBits(mask CR3_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CR3) Load() CR3_Bits              { return CR3_Bits(r.U32.Load()) }
func (r *CR3) Store(b CR3_Bits)            { r.U32.Store(uint32(b)) }

func (r *CR3) AtomicSetBits(mask CR3_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CR3) AtomicClearBits(mask CR3_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CR3_Mask struct{ mmio.UM32 }

func (rm CR3_Mask) Load() CR3_Bits   { return CR3_Bits(rm.UM32.Load()) }
func (rm CR3_Mask) Store(b CR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USART_Periph) EIE() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(EIE)}}
}

func (p *USART_Periph) IREN() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(IREN)}}
}

func (p *USART_Periph) IRLP() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(IRLP)}}
}

func (p *USART_Periph) HDSEL() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(HDSEL)}}
}

func (p *USART_Periph) NACK() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(NACK)}}
}

func (p *USART_Periph) SCEN() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(SCEN)}}
}

func (p *USART_Periph) DMAR() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(DMAR)}}
}

func (p *USART_Periph) DMAT() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(DMAT)}}
}

func (p *USART_Periph) RTSE() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(RTSE)}}
}

func (p *USART_Periph) CTSE() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(CTSE)}}
}

func (p *USART_Periph) CTSIE() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(CTSIE)}}
}

func (p *USART_Periph) ONEBIT() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(ONEBIT)}}
}

func (p *USART_Periph) OVRDIS() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(OVRDIS)}}
}

func (p *USART_Periph) DDRE() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(DDRE)}}
}

func (p *USART_Periph) DEM() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(DEM)}}
}

func (p *USART_Periph) DEP() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(DEP)}}
}

func (p *USART_Periph) SCARCNT() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(SCARCNT)}}
}

func (p *USART_Periph) WUS() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(WUS)}}
}

func (p *USART_Periph) WUFIE() CR3_Mask {
	return CR3_Mask{mmio.UM32{&p.CR3.U32, uint32(WUFIE)}}
}

type BRR_Bits uint16

func (b BRR_Bits) Field(mask BRR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BRR_Bits) J(v int) BRR_Bits {
	return BRR_Bits(bits.Make32(v, uint32(mask)))
}

type BRR struct{ mmio.U16 }

func (r *BRR) Bits(mask BRR_Bits) BRR_Bits { return BRR_Bits(r.U16.Bits(uint16(mask))) }
func (r *BRR) StoreBits(mask, b BRR_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *BRR) SetBits(mask BRR_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *BRR) ClearBits(mask BRR_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *BRR) Load() BRR_Bits              { return BRR_Bits(r.U16.Load()) }
func (r *BRR) Store(b BRR_Bits)            { r.U16.Store(uint16(b)) }

type BRR_Mask struct{ mmio.UM16 }

func (rm BRR_Mask) Load() BRR_Bits   { return BRR_Bits(rm.UM16.Load()) }
func (rm BRR_Mask) Store(b BRR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *USART_Periph) DIV_FRACTION() BRR_Mask {
	return BRR_Mask{mmio.UM16{&p.BRR.U16, uint16(DIV_FRACTION)}}
}

func (p *USART_Periph) DIV_MANTISSA() BRR_Mask {
	return BRR_Mask{mmio.UM16{&p.BRR.U16, uint16(DIV_MANTISSA)}}
}

type GTPR_Bits uint16

func (b GTPR_Bits) Field(mask GTPR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask GTPR_Bits) J(v int) GTPR_Bits {
	return GTPR_Bits(bits.Make32(v, uint32(mask)))
}

type GTPR struct{ mmio.U16 }

func (r *GTPR) Bits(mask GTPR_Bits) GTPR_Bits { return GTPR_Bits(r.U16.Bits(uint16(mask))) }
func (r *GTPR) StoreBits(mask, b GTPR_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *GTPR) SetBits(mask GTPR_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *GTPR) ClearBits(mask GTPR_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *GTPR) Load() GTPR_Bits               { return GTPR_Bits(r.U16.Load()) }
func (r *GTPR) Store(b GTPR_Bits)             { r.U16.Store(uint16(b)) }

type GTPR_Mask struct{ mmio.UM16 }

func (rm GTPR_Mask) Load() GTPR_Bits   { return GTPR_Bits(rm.UM16.Load()) }
func (rm GTPR_Mask) Store(b GTPR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *USART_Periph) PSC() GTPR_Mask {
	return GTPR_Mask{mmio.UM16{&p.GTPR.U16, uint16(PSC)}}
}

func (p *USART_Periph) GT() GTPR_Mask {
	return GTPR_Mask{mmio.UM16{&p.GTPR.U16, uint16(GT)}}
}

type RTOR_Bits uint32

func (b RTOR_Bits) Field(mask RTOR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RTOR_Bits) J(v int) RTOR_Bits {
	return RTOR_Bits(bits.Make32(v, uint32(mask)))
}

type RTOR struct{ mmio.U32 }

func (r *RTOR) Bits(mask RTOR_Bits) RTOR_Bits { return RTOR_Bits(r.U32.Bits(uint32(mask))) }
func (r *RTOR) StoreBits(mask, b RTOR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTOR) SetBits(mask RTOR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *RTOR) ClearBits(mask RTOR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *RTOR) Load() RTOR_Bits               { return RTOR_Bits(r.U32.Load()) }
func (r *RTOR) Store(b RTOR_Bits)             { r.U32.Store(uint32(b)) }

func (r *RTOR) AtomicSetBits(mask RTOR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTOR) AtomicClearBits(mask RTOR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type RTOR_Mask struct{ mmio.UM32 }

func (rm RTOR_Mask) Load() RTOR_Bits   { return RTOR_Bits(rm.UM32.Load()) }
func (rm RTOR_Mask) Store(b RTOR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USART_Periph) RTO() RTOR_Mask {
	return RTOR_Mask{mmio.UM32{&p.RTOR.U32, uint32(RTO)}}
}

func (p *USART_Periph) BLEN() RTOR_Mask {
	return RTOR_Mask{mmio.UM32{&p.RTOR.U32, uint32(BLEN)}}
}

type RQR_Bits uint16

func (b RQR_Bits) Field(mask RQR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RQR_Bits) J(v int) RQR_Bits {
	return RQR_Bits(bits.Make32(v, uint32(mask)))
}

type RQR struct{ mmio.U16 }

func (r *RQR) Bits(mask RQR_Bits) RQR_Bits { return RQR_Bits(r.U16.Bits(uint16(mask))) }
func (r *RQR) StoreBits(mask, b RQR_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RQR) SetBits(mask RQR_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *RQR) ClearBits(mask RQR_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *RQR) Load() RQR_Bits              { return RQR_Bits(r.U16.Load()) }
func (r *RQR) Store(b RQR_Bits)            { r.U16.Store(uint16(b)) }

type RQR_Mask struct{ mmio.UM16 }

func (rm RQR_Mask) Load() RQR_Bits   { return RQR_Bits(rm.UM16.Load()) }
func (rm RQR_Mask) Store(b RQR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *USART_Periph) ABRRQ() RQR_Mask {
	return RQR_Mask{mmio.UM16{&p.RQR.U16, uint16(ABRRQ)}}
}

func (p *USART_Periph) SBKRQ() RQR_Mask {
	return RQR_Mask{mmio.UM16{&p.RQR.U16, uint16(SBKRQ)}}
}

func (p *USART_Periph) MMRQ() RQR_Mask {
	return RQR_Mask{mmio.UM16{&p.RQR.U16, uint16(MMRQ)}}
}

func (p *USART_Periph) RXFRQ() RQR_Mask {
	return RQR_Mask{mmio.UM16{&p.RQR.U16, uint16(RXFRQ)}}
}

func (p *USART_Periph) TXFRQ() RQR_Mask {
	return RQR_Mask{mmio.UM16{&p.RQR.U16, uint16(TXFRQ)}}
}

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

func (p *USART_Periph) PE() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(PE)}}
}

func (p *USART_Periph) FE() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(FE)}}
}

func (p *USART_Periph) NE() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(NE)}}
}

func (p *USART_Periph) ORE() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(ORE)}}
}

func (p *USART_Periph) IDLE() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(IDLE)}}
}

func (p *USART_Periph) RXNE() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(RXNE)}}
}

func (p *USART_Periph) TC() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TC)}}
}

func (p *USART_Periph) TXE() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TXE)}}
}

func (p *USART_Periph) LBD() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(LBD)}}
}

func (p *USART_Periph) CTSIF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(CTSIF)}}
}

func (p *USART_Periph) CTS() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(CTS)}}
}

func (p *USART_Periph) RTOF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(RTOF)}}
}

func (p *USART_Periph) EOBF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(EOBF)}}
}

func (p *USART_Periph) ABRE() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(ABRE)}}
}

func (p *USART_Periph) ABRF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(ABRF)}}
}

func (p *USART_Periph) BUSY() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(BUSY)}}
}

func (p *USART_Periph) CMF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(CMF)}}
}

func (p *USART_Periph) SBKF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(SBKF)}}
}

func (p *USART_Periph) RWU() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(RWU)}}
}

func (p *USART_Periph) WUF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(WUF)}}
}

func (p *USART_Periph) TEACK() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TEACK)}}
}

func (p *USART_Periph) REACK() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(REACK)}}
}

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

func (p *USART_Periph) PECF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(PECF)}}
}

func (p *USART_Periph) FECF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(FECF)}}
}

func (p *USART_Periph) NCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(NCF)}}
}

func (p *USART_Periph) ORECF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(ORECF)}}
}

func (p *USART_Periph) IDLECF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(IDLECF)}}
}

func (p *USART_Periph) TCCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(TCCF)}}
}

func (p *USART_Periph) LBDCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(LBDCF)}}
}

func (p *USART_Periph) CTSCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(CTSCF)}}
}

func (p *USART_Periph) RTOCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(RTOCF)}}
}

func (p *USART_Periph) EOBCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(EOBCF)}}
}

func (p *USART_Periph) CMCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(CMCF)}}
}

func (p *USART_Periph) WUCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(WUCF)}}
}

type RDR_Bits uint16

func (b RDR_Bits) Field(mask RDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RDR_Bits) J(v int) RDR_Bits {
	return RDR_Bits(bits.Make32(v, uint32(mask)))
}

type RDR struct{ mmio.U16 }

func (r *RDR) Bits(mask RDR_Bits) RDR_Bits { return RDR_Bits(r.U16.Bits(uint16(mask))) }
func (r *RDR) StoreBits(mask, b RDR_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RDR) SetBits(mask RDR_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *RDR) ClearBits(mask RDR_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *RDR) Load() RDR_Bits              { return RDR_Bits(r.U16.Load()) }
func (r *RDR) Store(b RDR_Bits)            { r.U16.Store(uint16(b)) }

type RDR_Mask struct{ mmio.UM16 }

func (rm RDR_Mask) Load() RDR_Bits   { return RDR_Bits(rm.UM16.Load()) }
func (rm RDR_Mask) Store(b RDR_Bits) { rm.UM16.Store(uint16(b)) }

type TDR_Bits uint16

func (b TDR_Bits) Field(mask TDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TDR_Bits) J(v int) TDR_Bits {
	return TDR_Bits(bits.Make32(v, uint32(mask)))
}

type TDR struct{ mmio.U16 }

func (r *TDR) Bits(mask TDR_Bits) TDR_Bits { return TDR_Bits(r.U16.Bits(uint16(mask))) }
func (r *TDR) StoreBits(mask, b TDR_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *TDR) SetBits(mask TDR_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *TDR) ClearBits(mask TDR_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *TDR) Load() TDR_Bits              { return TDR_Bits(r.U16.Load()) }
func (r *TDR) Store(b TDR_Bits)            { r.U16.Store(uint16(b)) }

type TDR_Mask struct{ mmio.UM16 }

func (rm TDR_Mask) Load() TDR_Bits   { return TDR_Bits(rm.UM16.Load()) }
func (rm TDR_Mask) Store(b TDR_Bits) { rm.UM16.Store(uint16(b)) }
