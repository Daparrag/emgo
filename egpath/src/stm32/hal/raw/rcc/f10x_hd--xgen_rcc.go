// +build f10x_hd

package rcc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_hd/mmap"
)

type RCC_Periph struct {
	CR       CR
	CFGR     CFGR
	CIR      CIR
	APB2RSTR APB2RSTR
	APB1RSTR APB1RSTR
	AHBENR   AHBENR
	APB2ENR  APB2ENR
	APB1ENR  APB1ENR
	BDCR     BDCR
	CSR      CSR
}

func (p *RCC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var RCC = (*RCC_Periph)(unsafe.Pointer(uintptr(mmap.RCC_BASE)))

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

func (p *RCC_Periph) HSION() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSION)}}
}

func (p *RCC_Periph) HSIRDY() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSIRDY)}}
}

func (p *RCC_Periph) HSITRIM() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSITRIM)}}
}

func (p *RCC_Periph) HSICAL() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSICAL)}}
}

func (p *RCC_Periph) HSEON() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSEON)}}
}

func (p *RCC_Periph) HSERDY() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSERDY)}}
}

func (p *RCC_Periph) HSEBYP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSEBYP)}}
}

func (p *RCC_Periph) CSSON() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(CSSON)}}
}

func (p *RCC_Periph) PLLON() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PLLON)}}
}

func (p *RCC_Periph) PLLRDY() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PLLRDY)}}
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

func (p *RCC_Periph) SW() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(SW)}}
}

func (p *RCC_Periph) SWS() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(SWS)}}
}

func (p *RCC_Periph) HPRE() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(HPRE)}}
}

func (p *RCC_Periph) PPRE1() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PPRE1)}}
}

func (p *RCC_Periph) PPRE2() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PPRE2)}}
}

func (p *RCC_Periph) ADCPRE() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(ADCPRE)}}
}

func (p *RCC_Periph) PLLSRC() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PLLSRC)}}
}

func (p *RCC_Periph) PLLXTPRE() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PLLXTPRE)}}
}

func (p *RCC_Periph) PLLMULL() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PLLMULL)}}
}

func (p *RCC_Periph) USBPRE() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(USBPRE)}}
}

func (p *RCC_Periph) MCO() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(MCO)}}
}

type CIR_Bits uint32

func (b CIR_Bits) Field(mask CIR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CIR_Bits) J(v int) CIR_Bits {
	return CIR_Bits(bits.Make32(v, uint32(mask)))
}

type CIR struct{ mmio.U32 }

func (r *CIR) Bits(mask CIR_Bits) CIR_Bits { return CIR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CIR) StoreBits(mask, b CIR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CIR) SetBits(mask CIR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CIR) ClearBits(mask CIR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CIR) Load() CIR_Bits              { return CIR_Bits(r.U32.Load()) }
func (r *CIR) Store(b CIR_Bits)            { r.U32.Store(uint32(b)) }

func (r *CIR) AtomicSetBits(mask CIR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CIR) AtomicClearBits(mask CIR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CIR_Mask struct{ mmio.UM32 }

func (rm CIR_Mask) Load() CIR_Bits   { return CIR_Bits(rm.UM32.Load()) }
func (rm CIR_Mask) Store(b CIR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) LSIRDYF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSIRDYF)}}
}

func (p *RCC_Periph) LSERDYF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSERDYF)}}
}

func (p *RCC_Periph) HSIRDYF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSIRDYF)}}
}

func (p *RCC_Periph) HSERDYF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSERDYF)}}
}

func (p *RCC_Periph) PLLRDYF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(PLLRDYF)}}
}

func (p *RCC_Periph) CSSF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(CSSF)}}
}

func (p *RCC_Periph) LSIRDYIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSIRDYIE)}}
}

func (p *RCC_Periph) LSERDYIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSERDYIE)}}
}

func (p *RCC_Periph) HSIRDYIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSIRDYIE)}}
}

func (p *RCC_Periph) HSERDYIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSERDYIE)}}
}

func (p *RCC_Periph) PLLRDYIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(PLLRDYIE)}}
}

func (p *RCC_Periph) LSIRDYC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSIRDYC)}}
}

func (p *RCC_Periph) LSERDYC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSERDYC)}}
}

func (p *RCC_Periph) HSIRDYC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSIRDYC)}}
}

func (p *RCC_Periph) HSERDYC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSERDYC)}}
}

func (p *RCC_Periph) PLLRDYC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(PLLRDYC)}}
}

func (p *RCC_Periph) CSSC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(CSSC)}}
}

type APB2RSTR_Bits uint32

func (b APB2RSTR_Bits) Field(mask APB2RSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2RSTR_Bits) J(v int) APB2RSTR_Bits {
	return APB2RSTR_Bits(bits.Make32(v, uint32(mask)))
}

type APB2RSTR struct{ mmio.U32 }

func (r *APB2RSTR) Bits(mask APB2RSTR_Bits) APB2RSTR_Bits {
	return APB2RSTR_Bits(r.U32.Bits(uint32(mask)))
}
func (r *APB2RSTR) StoreBits(mask, b APB2RSTR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *APB2RSTR) SetBits(mask APB2RSTR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *APB2RSTR) ClearBits(mask APB2RSTR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *APB2RSTR) Load() APB2RSTR_Bits             { return APB2RSTR_Bits(r.U32.Load()) }
func (r *APB2RSTR) Store(b APB2RSTR_Bits)           { r.U32.Store(uint32(b)) }

func (r *APB2RSTR) AtomicSetBits(mask APB2RSTR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *APB2RSTR) AtomicClearBits(mask APB2RSTR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type APB2RSTR_Mask struct{ mmio.UM32 }

func (rm APB2RSTR_Mask) Load() APB2RSTR_Bits   { return APB2RSTR_Bits(rm.UM32.Load()) }
func (rm APB2RSTR_Mask) Store(b APB2RSTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) AFIORST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(AFIORST)}}
}

func (p *RCC_Periph) IOPARST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(IOPARST)}}
}

func (p *RCC_Periph) IOPBRST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(IOPBRST)}}
}

func (p *RCC_Periph) IOPCRST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(IOPCRST)}}
}

func (p *RCC_Periph) IOPDRST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(IOPDRST)}}
}

func (p *RCC_Periph) ADC1RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(ADC1RST)}}
}

func (p *RCC_Periph) ADC2RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(ADC2RST)}}
}

func (p *RCC_Periph) TIM1RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM1RST)}}
}

func (p *RCC_Periph) SPI1RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(SPI1RST)}}
}

func (p *RCC_Periph) USART1RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(USART1RST)}}
}

func (p *RCC_Periph) IOPERST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(IOPERST)}}
}

func (p *RCC_Periph) IOPFRST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(IOPFRST)}}
}

func (p *RCC_Periph) IOPGRST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(IOPGRST)}}
}

func (p *RCC_Periph) TIM8RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM8RST)}}
}

func (p *RCC_Periph) ADC3RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(ADC3RST)}}
}

type APB1RSTR_Bits uint32

func (b APB1RSTR_Bits) Field(mask APB1RSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1RSTR_Bits) J(v int) APB1RSTR_Bits {
	return APB1RSTR_Bits(bits.Make32(v, uint32(mask)))
}

type APB1RSTR struct{ mmio.U32 }

func (r *APB1RSTR) Bits(mask APB1RSTR_Bits) APB1RSTR_Bits {
	return APB1RSTR_Bits(r.U32.Bits(uint32(mask)))
}
func (r *APB1RSTR) StoreBits(mask, b APB1RSTR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *APB1RSTR) SetBits(mask APB1RSTR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *APB1RSTR) ClearBits(mask APB1RSTR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *APB1RSTR) Load() APB1RSTR_Bits             { return APB1RSTR_Bits(r.U32.Load()) }
func (r *APB1RSTR) Store(b APB1RSTR_Bits)           { r.U32.Store(uint32(b)) }

func (r *APB1RSTR) AtomicSetBits(mask APB1RSTR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *APB1RSTR) AtomicClearBits(mask APB1RSTR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type APB1RSTR_Mask struct{ mmio.UM32 }

func (rm APB1RSTR_Mask) Load() APB1RSTR_Bits   { return APB1RSTR_Bits(rm.UM32.Load()) }
func (rm APB1RSTR_Mask) Store(b APB1RSTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) TIM2RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM2RST)}}
}

func (p *RCC_Periph) TIM3RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM3RST)}}
}

func (p *RCC_Periph) WWDGRST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(WWDGRST)}}
}

func (p *RCC_Periph) USART2RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(USART2RST)}}
}

func (p *RCC_Periph) I2C1RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(I2C1RST)}}
}

func (p *RCC_Periph) CAN1RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(CAN1RST)}}
}

func (p *RCC_Periph) BKPRST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(BKPRST)}}
}

func (p *RCC_Periph) PWRRST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(PWRRST)}}
}

func (p *RCC_Periph) TIM4RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM4RST)}}
}

func (p *RCC_Periph) SPI2RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(SPI2RST)}}
}

func (p *RCC_Periph) USART3RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(USART3RST)}}
}

func (p *RCC_Periph) I2C2RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(I2C2RST)}}
}

func (p *RCC_Periph) USBRST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(USBRST)}}
}

func (p *RCC_Periph) TIM5RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM5RST)}}
}

func (p *RCC_Periph) TIM6RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM6RST)}}
}

func (p *RCC_Periph) TIM7RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM7RST)}}
}

func (p *RCC_Periph) SPI3RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(SPI3RST)}}
}

func (p *RCC_Periph) UART4RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(UART4RST)}}
}

func (p *RCC_Periph) UART5RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(UART5RST)}}
}

func (p *RCC_Periph) DACRST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(DACRST)}}
}

type AHBENR_Bits uint32

func (b AHBENR_Bits) Field(mask AHBENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AHBENR_Bits) J(v int) AHBENR_Bits {
	return AHBENR_Bits(bits.Make32(v, uint32(mask)))
}

type AHBENR struct{ mmio.U32 }

func (r *AHBENR) Bits(mask AHBENR_Bits) AHBENR_Bits { return AHBENR_Bits(r.U32.Bits(uint32(mask))) }
func (r *AHBENR) StoreBits(mask, b AHBENR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *AHBENR) SetBits(mask AHBENR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *AHBENR) ClearBits(mask AHBENR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *AHBENR) Load() AHBENR_Bits                 { return AHBENR_Bits(r.U32.Load()) }
func (r *AHBENR) Store(b AHBENR_Bits)               { r.U32.Store(uint32(b)) }

func (r *AHBENR) AtomicSetBits(mask AHBENR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *AHBENR) AtomicClearBits(mask AHBENR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type AHBENR_Mask struct{ mmio.UM32 }

func (rm AHBENR_Mask) Load() AHBENR_Bits   { return AHBENR_Bits(rm.UM32.Load()) }
func (rm AHBENR_Mask) Store(b AHBENR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) DMA1EN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(DMA1EN)}}
}

func (p *RCC_Periph) SRAMEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(SRAMEN)}}
}

func (p *RCC_Periph) FLITFEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(FLITFEN)}}
}

func (p *RCC_Periph) CRCEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(CRCEN)}}
}

func (p *RCC_Periph) DMA2EN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(DMA2EN)}}
}

func (p *RCC_Periph) FSMCEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(FSMCEN)}}
}

func (p *RCC_Periph) SDIOEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(SDIOEN)}}
}

type APB2ENR_Bits uint32

func (b APB2ENR_Bits) Field(mask APB2ENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2ENR_Bits) J(v int) APB2ENR_Bits {
	return APB2ENR_Bits(bits.Make32(v, uint32(mask)))
}

type APB2ENR struct{ mmio.U32 }

func (r *APB2ENR) Bits(mask APB2ENR_Bits) APB2ENR_Bits { return APB2ENR_Bits(r.U32.Bits(uint32(mask))) }
func (r *APB2ENR) StoreBits(mask, b APB2ENR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *APB2ENR) SetBits(mask APB2ENR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *APB2ENR) ClearBits(mask APB2ENR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *APB2ENR) Load() APB2ENR_Bits                  { return APB2ENR_Bits(r.U32.Load()) }
func (r *APB2ENR) Store(b APB2ENR_Bits)                { r.U32.Store(uint32(b)) }

func (r *APB2ENR) AtomicSetBits(mask APB2ENR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *APB2ENR) AtomicClearBits(mask APB2ENR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type APB2ENR_Mask struct{ mmio.UM32 }

func (rm APB2ENR_Mask) Load() APB2ENR_Bits   { return APB2ENR_Bits(rm.UM32.Load()) }
func (rm APB2ENR_Mask) Store(b APB2ENR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) AFIOEN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(AFIOEN)}}
}

func (p *RCC_Periph) IOPAEN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(IOPAEN)}}
}

func (p *RCC_Periph) IOPBEN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(IOPBEN)}}
}

func (p *RCC_Periph) IOPCEN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(IOPCEN)}}
}

func (p *RCC_Periph) IOPDEN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(IOPDEN)}}
}

func (p *RCC_Periph) ADC1EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(ADC1EN)}}
}

func (p *RCC_Periph) ADC2EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(ADC2EN)}}
}

func (p *RCC_Periph) TIM1EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(TIM1EN)}}
}

func (p *RCC_Periph) SPI1EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(SPI1EN)}}
}

func (p *RCC_Periph) USART1EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(USART1EN)}}
}

func (p *RCC_Periph) IOPEEN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(IOPEEN)}}
}

func (p *RCC_Periph) IOPFEN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(IOPFEN)}}
}

func (p *RCC_Periph) IOPGEN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(IOPGEN)}}
}

func (p *RCC_Periph) TIM8EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(TIM8EN)}}
}

func (p *RCC_Periph) ADC3EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(ADC3EN)}}
}

type APB1ENR_Bits uint32

func (b APB1ENR_Bits) Field(mask APB1ENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1ENR_Bits) J(v int) APB1ENR_Bits {
	return APB1ENR_Bits(bits.Make32(v, uint32(mask)))
}

type APB1ENR struct{ mmio.U32 }

func (r *APB1ENR) Bits(mask APB1ENR_Bits) APB1ENR_Bits { return APB1ENR_Bits(r.U32.Bits(uint32(mask))) }
func (r *APB1ENR) StoreBits(mask, b APB1ENR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *APB1ENR) SetBits(mask APB1ENR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *APB1ENR) ClearBits(mask APB1ENR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *APB1ENR) Load() APB1ENR_Bits                  { return APB1ENR_Bits(r.U32.Load()) }
func (r *APB1ENR) Store(b APB1ENR_Bits)                { r.U32.Store(uint32(b)) }

func (r *APB1ENR) AtomicSetBits(mask APB1ENR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *APB1ENR) AtomicClearBits(mask APB1ENR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type APB1ENR_Mask struct{ mmio.UM32 }

func (rm APB1ENR_Mask) Load() APB1ENR_Bits   { return APB1ENR_Bits(rm.UM32.Load()) }
func (rm APB1ENR_Mask) Store(b APB1ENR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) TIM2EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM2EN)}}
}

func (p *RCC_Periph) TIM3EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM3EN)}}
}

func (p *RCC_Periph) WWDGEN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(WWDGEN)}}
}

func (p *RCC_Periph) USART2EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(USART2EN)}}
}

func (p *RCC_Periph) I2C1EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(I2C1EN)}}
}

func (p *RCC_Periph) CAN1EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(CAN1EN)}}
}

func (p *RCC_Periph) BKPEN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(BKPEN)}}
}

func (p *RCC_Periph) PWREN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(PWREN)}}
}

func (p *RCC_Periph) TIM4EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM4EN)}}
}

func (p *RCC_Periph) SPI2EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(SPI2EN)}}
}

func (p *RCC_Periph) USART3EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(USART3EN)}}
}

func (p *RCC_Periph) I2C2EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(I2C2EN)}}
}

func (p *RCC_Periph) USBEN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(USBEN)}}
}

func (p *RCC_Periph) TIM5EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM5EN)}}
}

func (p *RCC_Periph) TIM6EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM6EN)}}
}

func (p *RCC_Periph) TIM7EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM7EN)}}
}

func (p *RCC_Periph) SPI3EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(SPI3EN)}}
}

func (p *RCC_Periph) UART4EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(UART4EN)}}
}

func (p *RCC_Periph) UART5EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(UART5EN)}}
}

func (p *RCC_Periph) DACEN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(DACEN)}}
}

type BDCR_Bits uint32

func (b BDCR_Bits) Field(mask BDCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BDCR_Bits) J(v int) BDCR_Bits {
	return BDCR_Bits(bits.Make32(v, uint32(mask)))
}

type BDCR struct{ mmio.U32 }

func (r *BDCR) Bits(mask BDCR_Bits) BDCR_Bits { return BDCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BDCR) StoreBits(mask, b BDCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BDCR) SetBits(mask BDCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *BDCR) ClearBits(mask BDCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *BDCR) Load() BDCR_Bits               { return BDCR_Bits(r.U32.Load()) }
func (r *BDCR) Store(b BDCR_Bits)             { r.U32.Store(uint32(b)) }

func (r *BDCR) AtomicSetBits(mask BDCR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *BDCR) AtomicClearBits(mask BDCR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type BDCR_Mask struct{ mmio.UM32 }

func (rm BDCR_Mask) Load() BDCR_Bits   { return BDCR_Bits(rm.UM32.Load()) }
func (rm BDCR_Mask) Store(b BDCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) LSEON() BDCR_Mask {
	return BDCR_Mask{mmio.UM32{&p.BDCR.U32, uint32(LSEON)}}
}

func (p *RCC_Periph) LSERDY() BDCR_Mask {
	return BDCR_Mask{mmio.UM32{&p.BDCR.U32, uint32(LSERDY)}}
}

func (p *RCC_Periph) LSEBYP() BDCR_Mask {
	return BDCR_Mask{mmio.UM32{&p.BDCR.U32, uint32(LSEBYP)}}
}

func (p *RCC_Periph) RTCSEL() BDCR_Mask {
	return BDCR_Mask{mmio.UM32{&p.BDCR.U32, uint32(RTCSEL)}}
}

func (p *RCC_Periph) RTCEN() BDCR_Mask {
	return BDCR_Mask{mmio.UM32{&p.BDCR.U32, uint32(RTCEN)}}
}

func (p *RCC_Periph) BDRST() BDCR_Mask {
	return BDCR_Mask{mmio.UM32{&p.BDCR.U32, uint32(BDRST)}}
}

type CSR_Bits uint32

func (b CSR_Bits) Field(mask CSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSR_Bits) J(v int) CSR_Bits {
	return CSR_Bits(bits.Make32(v, uint32(mask)))
}

type CSR struct{ mmio.U32 }

func (r *CSR) Bits(mask CSR_Bits) CSR_Bits { return CSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSR) StoreBits(mask, b CSR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSR) SetBits(mask CSR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CSR) ClearBits(mask CSR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CSR) Load() CSR_Bits              { return CSR_Bits(r.U32.Load()) }
func (r *CSR) Store(b CSR_Bits)            { r.U32.Store(uint32(b)) }

func (r *CSR) AtomicSetBits(mask CSR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CSR) AtomicClearBits(mask CSR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CSR_Mask struct{ mmio.UM32 }

func (rm CSR_Mask) Load() CSR_Bits   { return CSR_Bits(rm.UM32.Load()) }
func (rm CSR_Mask) Store(b CSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) LSION() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(LSION)}}
}

func (p *RCC_Periph) LSIRDY() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(LSIRDY)}}
}

func (p *RCC_Periph) RMVF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(RMVF)}}
}

func (p *RCC_Periph) PINRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(PINRSTF)}}
}

func (p *RCC_Periph) PORRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(PORRSTF)}}
}

func (p *RCC_Periph) SFTRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(SFTRSTF)}}
}

func (p *RCC_Periph) IWDGRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(IWDGRSTF)}}
}

func (p *RCC_Periph) WWDGRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(WWDGRSTF)}}
}

func (p *RCC_Periph) LPWRRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(LPWRRSTF)}}
}
