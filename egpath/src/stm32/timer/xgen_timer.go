package timer

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"
)

type Timer struct{}

func (p *Timer) r(n uint) *mmio.U32 {
	return &(*[20]mmio.U32)(unsafe.Pointer(p))[n]
}


var TIM1 = (*Timer)(unsafe.Pointer(uintptr(0x40010000)))
var TIM2 = (*Timer)(unsafe.Pointer(uintptr(0x40000000)))
var TIM3 = (*Timer)(unsafe.Pointer(uintptr(0x40000400)))
var TIM4 = (*Timer)(unsafe.Pointer(uintptr(0x40000800)))
var TIM5 = (*Timer)(unsafe.Pointer(uintptr(0x40000C00)))
var TIM6 = (*Timer)(unsafe.Pointer(uintptr(0x40001000)))
var TIM7 = (*Timer)(unsafe.Pointer(uintptr(0x40001400)))
var TIM8 = (*Timer)(unsafe.Pointer(uintptr(0x40010400)))
var TIM9 = (*Timer)(unsafe.Pointer(uintptr(0x40014000)))
var TIM10 = (*Timer)(unsafe.Pointer(uintptr(0x40014400)))
var TIM11 = (*Timer)(unsafe.Pointer(uintptr(0x40014800)))
var TIM12 = (*Timer)(unsafe.Pointer(uintptr(0x40001800)))
var TIM13 = (*Timer)(unsafe.Pointer(uintptr(0x40001C00)))
var TIM14 = (*Timer)(unsafe.Pointer(uintptr(0x40002000)))


type CR1_Bits uint32

func (p *Timer) CEN() mmio.Bits32 {return mmio.Bits32{p.r(0), uint32(CEN)} }
func (p *Timer) UDIS() mmio.Bits32 {return mmio.Bits32{p.r(0), uint32(UDIS)} }
func (p *Timer) URS() mmio.Bits32 {return mmio.Bits32{p.r(0), uint32(URS)} }
func (p *Timer) OPM() mmio.Bits32 {return mmio.Bits32{p.r(0), uint32(OPM)} }
func (p *Timer) DIR() mmio.Bits32 {return mmio.Bits32{p.r(0), uint32(DIR)} }
func (p *Timer) CMS() mmio.Bits32 {return mmio.Bits32{p.r(0), uint32(CMS)} }
func (p *Timer) ARPE() mmio.Bits32 {return mmio.Bits32{p.r(0), uint32(ARPE)} }
func (p *Timer) CKD() mmio.Bits32 {return mmio.Bits32{p.r(0), uint32(CKD)} }

func (p *Timer) CR1_Load() CR1_Bits   { return CR1_Bits(p.r(0).Load()) }
func (p *Timer) CR1_Store(b CR1_Bits) { p.r(0).Store(uint32(b)) }

func (b CR1_Bits) Field(mask CR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_CR1(v int, mask CR1_Bits) CR1_Bits {
	return CR1_Bits(bits.Make32(v, uint32(mask)))
}


type CR2_Bits uint32

func (p *Timer) CCPC() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(CCPC)} }
func (p *Timer) CCUS() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(CCUS)} }
func (p *Timer) CCDS() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(CCDS)} }
func (p *Timer) MMS() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(MMS)} }
func (p *Timer) TI1S() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(TI1S)} }
func (p *Timer) OIS1() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(OIS1)} }
func (p *Timer) OIS1N() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(OIS1N)} }
func (p *Timer) OIS2() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(OIS2)} }
func (p *Timer) OIS2N() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(OIS2N)} }
func (p *Timer) OIS3() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(OIS3)} }
func (p *Timer) OIS3N() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(OIS3N)} }
func (p *Timer) OIS4() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(OIS4)} }

func (p *Timer) CR2_Load() CR2_Bits   { return CR2_Bits(p.r(1).Load()) }
func (p *Timer) CR2_Store(b CR2_Bits) { p.r(1).Store(uint32(b)) }

func (b CR2_Bits) Field(mask CR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_CR2(v int, mask CR2_Bits) CR2_Bits {
	return CR2_Bits(bits.Make32(v, uint32(mask)))
}


type SMCR_Bits uint32

func (p *Timer) SMS() mmio.Bits32 {return mmio.Bits32{p.r(2), uint32(SMS)} }
func (p *Timer) TS() mmio.Bits32 {return mmio.Bits32{p.r(2), uint32(TS)} }
func (p *Timer) MSM() mmio.Bits32 {return mmio.Bits32{p.r(2), uint32(MSM)} }
func (p *Timer) ETF() mmio.Bits32 {return mmio.Bits32{p.r(2), uint32(ETF)} }
func (p *Timer) ETPS() mmio.Bits32 {return mmio.Bits32{p.r(2), uint32(ETPS)} }
func (p *Timer) ECE() mmio.Bits32 {return mmio.Bits32{p.r(2), uint32(ECE)} }
func (p *Timer) ETP() mmio.Bits32 {return mmio.Bits32{p.r(2), uint32(ETP)} }

func (p *Timer) SMCR_Load() SMCR_Bits   { return SMCR_Bits(p.r(2).Load()) }
func (p *Timer) SMCR_Store(b SMCR_Bits) { p.r(2).Store(uint32(b)) }

func (b SMCR_Bits) Field(mask SMCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_SMCR(v int, mask SMCR_Bits) SMCR_Bits {
	return SMCR_Bits(bits.Make32(v, uint32(mask)))
}


type DIER_Bits uint32

func (p *Timer) UIE() mmio.Bits32 {return mmio.Bits32{p.r(3), uint32(UIE)} }
func (p *Timer) CC1IE() mmio.Bits32 {return mmio.Bits32{p.r(3), uint32(CC1IE)} }
func (p *Timer) CC2IE() mmio.Bits32 {return mmio.Bits32{p.r(3), uint32(CC2IE)} }
func (p *Timer) CC3IE() mmio.Bits32 {return mmio.Bits32{p.r(3), uint32(CC3IE)} }
func (p *Timer) CC4IE() mmio.Bits32 {return mmio.Bits32{p.r(3), uint32(CC4IE)} }
func (p *Timer) COMIE() mmio.Bits32 {return mmio.Bits32{p.r(3), uint32(COMIE)} }
func (p *Timer) TIE() mmio.Bits32 {return mmio.Bits32{p.r(3), uint32(TIE)} }
func (p *Timer) BIE() mmio.Bits32 {return mmio.Bits32{p.r(3), uint32(BIE)} }
func (p *Timer) UDE() mmio.Bits32 {return mmio.Bits32{p.r(3), uint32(UDE)} }
func (p *Timer) CC1DE() mmio.Bits32 {return mmio.Bits32{p.r(3), uint32(CC1DE)} }
func (p *Timer) CC2DE() mmio.Bits32 {return mmio.Bits32{p.r(3), uint32(CC2DE)} }
func (p *Timer) CC3DE() mmio.Bits32 {return mmio.Bits32{p.r(3), uint32(CC3DE)} }
func (p *Timer) CC4DE() mmio.Bits32 {return mmio.Bits32{p.r(3), uint32(CC4DE)} }
func (p *Timer) COMDE() mmio.Bits32 {return mmio.Bits32{p.r(3), uint32(COMDE)} }
func (p *Timer) TDE() mmio.Bits32 {return mmio.Bits32{p.r(3), uint32(TDE)} }

func (p *Timer) DIER_Load() DIER_Bits   { return DIER_Bits(p.r(3).Load()) }
func (p *Timer) DIER_Store(b DIER_Bits) { p.r(3).Store(uint32(b)) }

func (b DIER_Bits) Field(mask DIER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_DIER(v int, mask DIER_Bits) DIER_Bits {
	return DIER_Bits(bits.Make32(v, uint32(mask)))
}


type SR_Bits uint32

func (p *Timer) UIF() mmio.Bits32 {return mmio.Bits32{p.r(4), uint32(UIF)} }
func (p *Timer) CC1IF() mmio.Bits32 {return mmio.Bits32{p.r(4), uint32(CC1IF)} }
func (p *Timer) CC2IF() mmio.Bits32 {return mmio.Bits32{p.r(4), uint32(CC2IF)} }
func (p *Timer) CC3IF() mmio.Bits32 {return mmio.Bits32{p.r(4), uint32(CC3IF)} }
func (p *Timer) CC4IF() mmio.Bits32 {return mmio.Bits32{p.r(4), uint32(CC4IF)} }
func (p *Timer) COMIF() mmio.Bits32 {return mmio.Bits32{p.r(4), uint32(COMIF)} }
func (p *Timer) TIF() mmio.Bits32 {return mmio.Bits32{p.r(4), uint32(TIF)} }
func (p *Timer) BIF() mmio.Bits32 {return mmio.Bits32{p.r(4), uint32(BIF)} }
func (p *Timer) CC1OF() mmio.Bits32 {return mmio.Bits32{p.r(4), uint32(CC1OF)} }
func (p *Timer) CC2OF() mmio.Bits32 {return mmio.Bits32{p.r(4), uint32(CC2OF)} }
func (p *Timer) CC3OF() mmio.Bits32 {return mmio.Bits32{p.r(4), uint32(CC3OF)} }
func (p *Timer) CC4OF() mmio.Bits32 {return mmio.Bits32{p.r(4), uint32(CC4OF)} }

func (p *Timer) SR_Load() SR_Bits   { return SR_Bits(p.r(4).Load()) }
func (p *Timer) SR_Store(b SR_Bits) { p.r(4).Store(uint32(b)) }

func (b SR_Bits) Field(mask SR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_SR(v int, mask SR_Bits) SR_Bits {
	return SR_Bits(bits.Make32(v, uint32(mask)))
}


type EGR_Bits uint32

func (p *Timer) UG() mmio.Bits32 {return mmio.Bits32{p.r(5), uint32(UG)} }
func (p *Timer) CC1G() mmio.Bits32 {return mmio.Bits32{p.r(5), uint32(CC1G)} }
func (p *Timer) CC2G() mmio.Bits32 {return mmio.Bits32{p.r(5), uint32(CC2G)} }
func (p *Timer) CC3G() mmio.Bits32 {return mmio.Bits32{p.r(5), uint32(CC3G)} }
func (p *Timer) CC4G() mmio.Bits32 {return mmio.Bits32{p.r(5), uint32(CC4G)} }
func (p *Timer) COMG() mmio.Bits32 {return mmio.Bits32{p.r(5), uint32(COMG)} }
func (p *Timer) TG() mmio.Bits32 {return mmio.Bits32{p.r(5), uint32(TG)} }
func (p *Timer) BG() mmio.Bits32 {return mmio.Bits32{p.r(5), uint32(BG)} }

func (p *Timer) EGR_Load() EGR_Bits   { return EGR_Bits(p.r(5).Load()) }
func (p *Timer) EGR_Store(b EGR_Bits) { p.r(5).Store(uint32(b)) }

func (b EGR_Bits) Field(mask EGR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_EGR(v int, mask EGR_Bits) EGR_Bits {
	return EGR_Bits(bits.Make32(v, uint32(mask)))
}


type CCMR1_Bits uint32

func (p *Timer) CC1S() mmio.Bits32 {return mmio.Bits32{p.r(6), uint32(CC1S)} }
func (p *Timer) OC1FE() mmio.Bits32 {return mmio.Bits32{p.r(6), uint32(OC1FE)} }
func (p *Timer) OC1PE() mmio.Bits32 {return mmio.Bits32{p.r(6), uint32(OC1PE)} }
func (p *Timer) OC1M() mmio.Bits32 {return mmio.Bits32{p.r(6), uint32(OC1M)} }
func (p *Timer) OC1CE() mmio.Bits32 {return mmio.Bits32{p.r(6), uint32(OC1CE)} }
func (p *Timer) CC2S() mmio.Bits32 {return mmio.Bits32{p.r(6), uint32(CC2S)} }
func (p *Timer) OC2FE() mmio.Bits32 {return mmio.Bits32{p.r(6), uint32(OC2FE)} }
func (p *Timer) OC2PE() mmio.Bits32 {return mmio.Bits32{p.r(6), uint32(OC2PE)} }
func (p *Timer) OC2M() mmio.Bits32 {return mmio.Bits32{p.r(6), uint32(OC2M)} }
func (p *Timer) OC2CE() mmio.Bits32 {return mmio.Bits32{p.r(6), uint32(OC2CE)} }
func (p *Timer) IC1PSC() mmio.Bits32 {return mmio.Bits32{p.r(6), uint32(IC1PSC)} }
func (p *Timer) IC1F() mmio.Bits32 {return mmio.Bits32{p.r(6), uint32(IC1F)} }
func (p *Timer) IC2PSC() mmio.Bits32 {return mmio.Bits32{p.r(6), uint32(IC2PSC)} }
func (p *Timer) IC2F() mmio.Bits32 {return mmio.Bits32{p.r(6), uint32(IC2F)} }

func (p *Timer) CCMR1_Load() CCMR1_Bits   { return CCMR1_Bits(p.r(6).Load()) }
func (p *Timer) CCMR1_Store(b CCMR1_Bits) { p.r(6).Store(uint32(b)) }

func (b CCMR1_Bits) Field(mask CCMR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_CCMR1(v int, mask CCMR1_Bits) CCMR1_Bits {
	return CCMR1_Bits(bits.Make32(v, uint32(mask)))
}


type CCMR2_Bits uint32

func (p *Timer) CC3S() mmio.Bits32 {return mmio.Bits32{p.r(7), uint32(CC3S)} }
func (p *Timer) OC3FE() mmio.Bits32 {return mmio.Bits32{p.r(7), uint32(OC3FE)} }
func (p *Timer) OC3PE() mmio.Bits32 {return mmio.Bits32{p.r(7), uint32(OC3PE)} }
func (p *Timer) OC3M() mmio.Bits32 {return mmio.Bits32{p.r(7), uint32(OC3M)} }
func (p *Timer) OC3CE() mmio.Bits32 {return mmio.Bits32{p.r(7), uint32(OC3CE)} }
func (p *Timer) CC4S() mmio.Bits32 {return mmio.Bits32{p.r(7), uint32(CC4S)} }
func (p *Timer) OC4FE() mmio.Bits32 {return mmio.Bits32{p.r(7), uint32(OC4FE)} }
func (p *Timer) OC4PE() mmio.Bits32 {return mmio.Bits32{p.r(7), uint32(OC4PE)} }
func (p *Timer) OC4M() mmio.Bits32 {return mmio.Bits32{p.r(7), uint32(OC4M)} }
func (p *Timer) OC4CE() mmio.Bits32 {return mmio.Bits32{p.r(7), uint32(OC4CE)} }
func (p *Timer) IC3PSC() mmio.Bits32 {return mmio.Bits32{p.r(7), uint32(IC3PSC)} }
func (p *Timer) IC3F() mmio.Bits32 {return mmio.Bits32{p.r(7), uint32(IC3F)} }
func (p *Timer) IC4PSC() mmio.Bits32 {return mmio.Bits32{p.r(7), uint32(IC4PSC)} }
func (p *Timer) IC4F() mmio.Bits32 {return mmio.Bits32{p.r(7), uint32(IC4F)} }

func (p *Timer) CCMR2_Load() CCMR2_Bits   { return CCMR2_Bits(p.r(7).Load()) }
func (p *Timer) CCMR2_Store(b CCMR2_Bits) { p.r(7).Store(uint32(b)) }

func (b CCMR2_Bits) Field(mask CCMR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_CCMR2(v int, mask CCMR2_Bits) CCMR2_Bits {
	return CCMR2_Bits(bits.Make32(v, uint32(mask)))
}


type CCER_Bits uint32

func (p *Timer) CC1E() mmio.Bits32 {return mmio.Bits32{p.r(8), uint32(CC1E)} }
func (p *Timer) CC1P() mmio.Bits32 {return mmio.Bits32{p.r(8), uint32(CC1P)} }
func (p *Timer) CC1NE() mmio.Bits32 {return mmio.Bits32{p.r(8), uint32(CC1NE)} }
func (p *Timer) CC1NP() mmio.Bits32 {return mmio.Bits32{p.r(8), uint32(CC1NP)} }
func (p *Timer) CC2E() mmio.Bits32 {return mmio.Bits32{p.r(8), uint32(CC2E)} }
func (p *Timer) CC2P() mmio.Bits32 {return mmio.Bits32{p.r(8), uint32(CC2P)} }
func (p *Timer) CC2NE() mmio.Bits32 {return mmio.Bits32{p.r(8), uint32(CC2NE)} }
func (p *Timer) CC2NP() mmio.Bits32 {return mmio.Bits32{p.r(8), uint32(CC2NP)} }
func (p *Timer) CC3E() mmio.Bits32 {return mmio.Bits32{p.r(8), uint32(CC3E)} }
func (p *Timer) CC3P() mmio.Bits32 {return mmio.Bits32{p.r(8), uint32(CC3P)} }
func (p *Timer) CC3NE() mmio.Bits32 {return mmio.Bits32{p.r(8), uint32(CC3NE)} }
func (p *Timer) CC3NP() mmio.Bits32 {return mmio.Bits32{p.r(8), uint32(CC3NP)} }
func (p *Timer) CC4E() mmio.Bits32 {return mmio.Bits32{p.r(8), uint32(CC4E)} }
func (p *Timer) CC4P() mmio.Bits32 {return mmio.Bits32{p.r(8), uint32(CC4P)} }

func (p *Timer) CCER_Load() CCER_Bits   { return CCER_Bits(p.r(8).Load()) }
func (p *Timer) CCER_Store(b CCER_Bits) { p.r(8).Store(uint32(b)) }

func (b CCER_Bits) Field(mask CCER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_CCER(v int, mask CCER_Bits) CCER_Bits {
	return CCER_Bits(bits.Make32(v, uint32(mask)))
}


type CNT_Bits uint32


func (p *Timer) CNT_Load() CNT_Bits   { return CNT_Bits(p.r(9).Load()) }
func (p *Timer) CNT_Store(b CNT_Bits) { p.r(9).Store(uint32(b)) }

func (b CNT_Bits) Field(mask CNT_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_CNT(v int, mask CNT_Bits) CNT_Bits {
	return CNT_Bits(bits.Make32(v, uint32(mask)))
}


type PSC_Bits uint32


func (p *Timer) PSC_Load() PSC_Bits   { return PSC_Bits(p.r(10).Load()) }
func (p *Timer) PSC_Store(b PSC_Bits) { p.r(10).Store(uint32(b)) }

func (b PSC_Bits) Field(mask PSC_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_PSC(v int, mask PSC_Bits) PSC_Bits {
	return PSC_Bits(bits.Make32(v, uint32(mask)))
}


type ARR_Bits uint32


func (p *Timer) ARR_Load() ARR_Bits   { return ARR_Bits(p.r(11).Load()) }
func (p *Timer) ARR_Store(b ARR_Bits) { p.r(11).Store(uint32(b)) }

func (b ARR_Bits) Field(mask ARR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_ARR(v int, mask ARR_Bits) ARR_Bits {
	return ARR_Bits(bits.Make32(v, uint32(mask)))
}


type RCR_Bits uint32

func (p *Timer) REP() mmio.Bits32 {return mmio.Bits32{p.r(12), uint32(REP)} }

func (p *Timer) RCR_Load() RCR_Bits   { return RCR_Bits(p.r(12).Load()) }
func (p *Timer) RCR_Store(b RCR_Bits) { p.r(12).Store(uint32(b)) }

func (b RCR_Bits) Field(mask RCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_RCR(v int, mask RCR_Bits) RCR_Bits {
	return RCR_Bits(bits.Make32(v, uint32(mask)))
}


type CCR1_Bits uint32


func (p *Timer) CCR1_Load() CCR1_Bits   { return CCR1_Bits(p.r(13).Load()) }
func (p *Timer) CCR1_Store(b CCR1_Bits) { p.r(13).Store(uint32(b)) }

func (b CCR1_Bits) Field(mask CCR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_CCR1(v int, mask CCR1_Bits) CCR1_Bits {
	return CCR1_Bits(bits.Make32(v, uint32(mask)))
}


type CCR2_Bits uint32


func (p *Timer) CCR2_Load() CCR2_Bits   { return CCR2_Bits(p.r(14).Load()) }
func (p *Timer) CCR2_Store(b CCR2_Bits) { p.r(14).Store(uint32(b)) }

func (b CCR2_Bits) Field(mask CCR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_CCR2(v int, mask CCR2_Bits) CCR2_Bits {
	return CCR2_Bits(bits.Make32(v, uint32(mask)))
}


type CCR3_Bits uint32


func (p *Timer) CCR3_Load() CCR3_Bits   { return CCR3_Bits(p.r(15).Load()) }
func (p *Timer) CCR3_Store(b CCR3_Bits) { p.r(15).Store(uint32(b)) }

func (b CCR3_Bits) Field(mask CCR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_CCR3(v int, mask CCR3_Bits) CCR3_Bits {
	return CCR3_Bits(bits.Make32(v, uint32(mask)))
}


type CCR4_Bits uint32


func (p *Timer) CCR4_Load() CCR4_Bits   { return CCR4_Bits(p.r(16).Load()) }
func (p *Timer) CCR4_Store(b CCR4_Bits) { p.r(16).Store(uint32(b)) }

func (b CCR4_Bits) Field(mask CCR4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_CCR4(v int, mask CCR4_Bits) CCR4_Bits {
	return CCR4_Bits(bits.Make32(v, uint32(mask)))
}


type BDTR_Bits uint32

func (p *Timer) DTG() mmio.Bits32 {return mmio.Bits32{p.r(17), uint32(DTG)} }
func (p *Timer) LOCK() mmio.Bits32 {return mmio.Bits32{p.r(17), uint32(LOCK)} }
func (p *Timer) OSSI() mmio.Bits32 {return mmio.Bits32{p.r(17), uint32(OSSI)} }
func (p *Timer) OSSR() mmio.Bits32 {return mmio.Bits32{p.r(17), uint32(OSSR)} }
func (p *Timer) BKE() mmio.Bits32 {return mmio.Bits32{p.r(17), uint32(BKE)} }
func (p *Timer) BKP() mmio.Bits32 {return mmio.Bits32{p.r(17), uint32(BKP)} }
func (p *Timer) AOE() mmio.Bits32 {return mmio.Bits32{p.r(17), uint32(AOE)} }
func (p *Timer) MOE() mmio.Bits32 {return mmio.Bits32{p.r(17), uint32(MOE)} }

func (p *Timer) BDTR_Load() BDTR_Bits   { return BDTR_Bits(p.r(17).Load()) }
func (p *Timer) BDTR_Store(b BDTR_Bits) { p.r(17).Store(uint32(b)) }

func (b BDTR_Bits) Field(mask BDTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_BDTR(v int, mask BDTR_Bits) BDTR_Bits {
	return BDTR_Bits(bits.Make32(v, uint32(mask)))
}


type DCR_Bits uint32

func (p *Timer) DBA() mmio.Bits32 {return mmio.Bits32{p.r(18), uint32(DBA)} }
func (p *Timer) DBL() mmio.Bits32 {return mmio.Bits32{p.r(18), uint32(DBL)} }

func (p *Timer) DCR_Load() DCR_Bits   { return DCR_Bits(p.r(18).Load()) }
func (p *Timer) DCR_Store(b DCR_Bits) { p.r(18).Store(uint32(b)) }

func (b DCR_Bits) Field(mask DCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_DCR(v int, mask DCR_Bits) DCR_Bits {
	return DCR_Bits(bits.Make32(v, uint32(mask)))
}


type DMAR_Bits uint32


func (p *Timer) DMAR_Load() DMAR_Bits   { return DMAR_Bits(p.r(19).Load()) }
func (p *Timer) DMAR_Store(b DMAR_Bits) { p.r(19).Store(uint32(b)) }

func (b DMAR_Bits) Field(mask DMAR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_DMAR(v int, mask DMAR_Bits) DMAR_Bits {
	return DMAR_Bits(bits.Make32(v, uint32(mask)))
}
