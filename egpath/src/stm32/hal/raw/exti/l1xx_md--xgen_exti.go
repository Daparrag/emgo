// +build l1xx_md

package exti

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type EXTI_Periph struct {
	IMR   IMR
	EMR   EMR
	RTSR  RTSR
	FTSR  FTSR
	SWIER SWIER
	PR    PR
}

func (p *EXTI_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var EXTI = (*EXTI_Periph)(unsafe.Pointer(uintptr(mmap.EXTI_BASE)))

type IMR_Bits uint32

func (b IMR_Bits) Field(mask IMR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IMR_Bits) J(v int) IMR_Bits {
	return IMR_Bits(bits.Make32(v, uint32(mask)))
}

type IMR struct{ mmio.U32 }

func (r *IMR) Bits(mask IMR_Bits) IMR_Bits { return IMR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IMR) StoreBits(mask, b IMR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IMR) SetBits(mask IMR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *IMR) ClearBits(mask IMR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *IMR) Load() IMR_Bits              { return IMR_Bits(r.U32.Load()) }
func (r *IMR) Store(b IMR_Bits)            { r.U32.Store(uint32(b)) }

func (r *IMR) AtomicSetBits(mask IMR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *IMR) AtomicClearBits(mask IMR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type IMR_Mask struct{ mmio.UM32 }

func (rm IMR_Mask) Load() IMR_Bits   { return IMR_Bits(rm.UM32.Load()) }
func (rm IMR_Mask) Store(b IMR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) IL0() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL0)}}
}

func (p *EXTI_Periph) IL1() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL1)}}
}

func (p *EXTI_Periph) IL2() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL2)}}
}

func (p *EXTI_Periph) IL3() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL3)}}
}

func (p *EXTI_Periph) IL4() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL4)}}
}

func (p *EXTI_Periph) IL5() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL5)}}
}

func (p *EXTI_Periph) IL6() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL6)}}
}

func (p *EXTI_Periph) IL7() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL7)}}
}

func (p *EXTI_Periph) IL8() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL8)}}
}

func (p *EXTI_Periph) IL9() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL9)}}
}

func (p *EXTI_Periph) IL10() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL10)}}
}

func (p *EXTI_Periph) IL11() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL11)}}
}

func (p *EXTI_Periph) IL12() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL12)}}
}

func (p *EXTI_Periph) IL13() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL13)}}
}

func (p *EXTI_Periph) IL14() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL14)}}
}

func (p *EXTI_Periph) IL15() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL15)}}
}

func (p *EXTI_Periph) IL16() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL16)}}
}

func (p *EXTI_Periph) IL17() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL17)}}
}

func (p *EXTI_Periph) IL18() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL18)}}
}

func (p *EXTI_Periph) IL19() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL19)}}
}

func (p *EXTI_Periph) IL20() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL20)}}
}

func (p *EXTI_Periph) IL21() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL21)}}
}

func (p *EXTI_Periph) IL22() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL22)}}
}

func (p *EXTI_Periph) IL23() IMR_Mask {
	return IMR_Mask{mmio.UM32{&p.IMR.U32, uint32(IL23)}}
}

type EMR_Bits uint32

func (b EMR_Bits) Field(mask EMR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask EMR_Bits) J(v int) EMR_Bits {
	return EMR_Bits(bits.Make32(v, uint32(mask)))
}

type EMR struct{ mmio.U32 }

func (r *EMR) Bits(mask EMR_Bits) EMR_Bits { return EMR_Bits(r.U32.Bits(uint32(mask))) }
func (r *EMR) StoreBits(mask, b EMR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *EMR) SetBits(mask EMR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *EMR) ClearBits(mask EMR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *EMR) Load() EMR_Bits              { return EMR_Bits(r.U32.Load()) }
func (r *EMR) Store(b EMR_Bits)            { r.U32.Store(uint32(b)) }

func (r *EMR) AtomicSetBits(mask EMR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *EMR) AtomicClearBits(mask EMR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type EMR_Mask struct{ mmio.UM32 }

func (rm EMR_Mask) Load() EMR_Bits   { return EMR_Bits(rm.UM32.Load()) }
func (rm EMR_Mask) Store(b EMR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) EL0() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL0)}}
}

func (p *EXTI_Periph) EL1() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL1)}}
}

func (p *EXTI_Periph) EL2() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL2)}}
}

func (p *EXTI_Periph) EL3() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL3)}}
}

func (p *EXTI_Periph) EL4() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL4)}}
}

func (p *EXTI_Periph) EL5() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL5)}}
}

func (p *EXTI_Periph) EL6() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL6)}}
}

func (p *EXTI_Periph) EL7() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL7)}}
}

func (p *EXTI_Periph) EL8() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL8)}}
}

func (p *EXTI_Periph) EL9() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL9)}}
}

func (p *EXTI_Periph) EL10() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL10)}}
}

func (p *EXTI_Periph) EL11() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL11)}}
}

func (p *EXTI_Periph) EL12() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL12)}}
}

func (p *EXTI_Periph) EL13() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL13)}}
}

func (p *EXTI_Periph) EL14() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL14)}}
}

func (p *EXTI_Periph) EL15() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL15)}}
}

func (p *EXTI_Periph) EL16() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL16)}}
}

func (p *EXTI_Periph) EL17() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL17)}}
}

func (p *EXTI_Periph) EL18() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL18)}}
}

func (p *EXTI_Periph) EL19() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL19)}}
}

func (p *EXTI_Periph) EL20() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL20)}}
}

func (p *EXTI_Periph) EL21() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL21)}}
}

func (p *EXTI_Periph) EL22() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL22)}}
}

func (p *EXTI_Periph) EL23() EMR_Mask {
	return EMR_Mask{mmio.UM32{&p.EMR.U32, uint32(EL23)}}
}

type RTSR_Bits uint32

func (b RTSR_Bits) Field(mask RTSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RTSR_Bits) J(v int) RTSR_Bits {
	return RTSR_Bits(bits.Make32(v, uint32(mask)))
}

type RTSR struct{ mmio.U32 }

func (r *RTSR) Bits(mask RTSR_Bits) RTSR_Bits { return RTSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *RTSR) StoreBits(mask, b RTSR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTSR) SetBits(mask RTSR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *RTSR) ClearBits(mask RTSR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *RTSR) Load() RTSR_Bits               { return RTSR_Bits(r.U32.Load()) }
func (r *RTSR) Store(b RTSR_Bits)             { r.U32.Store(uint32(b)) }

func (r *RTSR) AtomicSetBits(mask RTSR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTSR) AtomicClearBits(mask RTSR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type RTSR_Mask struct{ mmio.UM32 }

func (rm RTSR_Mask) Load() RTSR_Bits   { return RTSR_Bits(rm.UM32.Load()) }
func (rm RTSR_Mask) Store(b RTSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) TR0() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR0)}}
}

func (p *EXTI_Periph) TR1() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR1)}}
}

func (p *EXTI_Periph) TR2() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR2)}}
}

func (p *EXTI_Periph) TR3() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR3)}}
}

func (p *EXTI_Periph) TR4() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR4)}}
}

func (p *EXTI_Periph) TR5() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR5)}}
}

func (p *EXTI_Periph) TR6() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR6)}}
}

func (p *EXTI_Periph) TR7() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR7)}}
}

func (p *EXTI_Periph) TR8() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR8)}}
}

func (p *EXTI_Periph) TR9() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR9)}}
}

func (p *EXTI_Periph) TR10() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR10)}}
}

func (p *EXTI_Periph) TR11() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR11)}}
}

func (p *EXTI_Periph) TR12() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR12)}}
}

func (p *EXTI_Periph) TR13() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR13)}}
}

func (p *EXTI_Periph) TR14() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR14)}}
}

func (p *EXTI_Periph) TR15() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR15)}}
}

func (p *EXTI_Periph) TR16() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR16)}}
}

func (p *EXTI_Periph) TR17() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR17)}}
}

func (p *EXTI_Periph) TR18() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR18)}}
}

func (p *EXTI_Periph) TR19() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR19)}}
}

func (p *EXTI_Periph) TR20() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR20)}}
}

func (p *EXTI_Periph) TR21() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR21)}}
}

func (p *EXTI_Periph) TR22() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR22)}}
}

func (p *EXTI_Periph) TR23() RTSR_Mask {
	return RTSR_Mask{mmio.UM32{&p.RTSR.U32, uint32(TR23)}}
}

type FTSR_Bits uint32

func (b FTSR_Bits) Field(mask FTSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FTSR_Bits) J(v int) FTSR_Bits {
	return FTSR_Bits(bits.Make32(v, uint32(mask)))
}

type FTSR struct{ mmio.U32 }

func (r *FTSR) Bits(mask FTSR_Bits) FTSR_Bits { return FTSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *FTSR) StoreBits(mask, b FTSR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FTSR) SetBits(mask FTSR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *FTSR) ClearBits(mask FTSR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *FTSR) Load() FTSR_Bits               { return FTSR_Bits(r.U32.Load()) }
func (r *FTSR) Store(b FTSR_Bits)             { r.U32.Store(uint32(b)) }

func (r *FTSR) AtomicSetBits(mask FTSR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *FTSR) AtomicClearBits(mask FTSR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type FTSR_Mask struct{ mmio.UM32 }

func (rm FTSR_Mask) Load() FTSR_Bits   { return FTSR_Bits(rm.UM32.Load()) }
func (rm FTSR_Mask) Store(b FTSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) TF0() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF0)}}
}

func (p *EXTI_Periph) TF1() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF1)}}
}

func (p *EXTI_Periph) TF2() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF2)}}
}

func (p *EXTI_Periph) TF3() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF3)}}
}

func (p *EXTI_Periph) TF4() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF4)}}
}

func (p *EXTI_Periph) TF5() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF5)}}
}

func (p *EXTI_Periph) TF6() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF6)}}
}

func (p *EXTI_Periph) TF7() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF7)}}
}

func (p *EXTI_Periph) TF8() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF8)}}
}

func (p *EXTI_Periph) TF9() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF9)}}
}

func (p *EXTI_Periph) TF10() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF10)}}
}

func (p *EXTI_Periph) TF11() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF11)}}
}

func (p *EXTI_Periph) TF12() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF12)}}
}

func (p *EXTI_Periph) TF13() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF13)}}
}

func (p *EXTI_Periph) TF14() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF14)}}
}

func (p *EXTI_Periph) TF15() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF15)}}
}

func (p *EXTI_Periph) TF16() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF16)}}
}

func (p *EXTI_Periph) TF17() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF17)}}
}

func (p *EXTI_Periph) TF18() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF18)}}
}

func (p *EXTI_Periph) TF19() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF19)}}
}

func (p *EXTI_Periph) TF20() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF20)}}
}

func (p *EXTI_Periph) TF21() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF21)}}
}

func (p *EXTI_Periph) TF22() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF22)}}
}

func (p *EXTI_Periph) TF23() FTSR_Mask {
	return FTSR_Mask{mmio.UM32{&p.FTSR.U32, uint32(TF23)}}
}

type SWIER_Bits uint32

func (b SWIER_Bits) Field(mask SWIER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SWIER_Bits) J(v int) SWIER_Bits {
	return SWIER_Bits(bits.Make32(v, uint32(mask)))
}

type SWIER struct{ mmio.U32 }

func (r *SWIER) Bits(mask SWIER_Bits) SWIER_Bits { return SWIER_Bits(r.U32.Bits(uint32(mask))) }
func (r *SWIER) StoreBits(mask, b SWIER_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SWIER) SetBits(mask SWIER_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *SWIER) ClearBits(mask SWIER_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *SWIER) Load() SWIER_Bits                { return SWIER_Bits(r.U32.Load()) }
func (r *SWIER) Store(b SWIER_Bits)              { r.U32.Store(uint32(b)) }

func (r *SWIER) AtomicSetBits(mask SWIER_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *SWIER) AtomicClearBits(mask SWIER_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type SWIER_Mask struct{ mmio.UM32 }

func (rm SWIER_Mask) Load() SWIER_Bits   { return SWIER_Bits(rm.UM32.Load()) }
func (rm SWIER_Mask) Store(b SWIER_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) SWIER0() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER0)}}
}

func (p *EXTI_Periph) SWIER1() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER1)}}
}

func (p *EXTI_Periph) SWIER2() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER2)}}
}

func (p *EXTI_Periph) SWIER3() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER3)}}
}

func (p *EXTI_Periph) SWIER4() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER4)}}
}

func (p *EXTI_Periph) SWIER5() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER5)}}
}

func (p *EXTI_Periph) SWIER6() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER6)}}
}

func (p *EXTI_Periph) SWIER7() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER7)}}
}

func (p *EXTI_Periph) SWIER8() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER8)}}
}

func (p *EXTI_Periph) SWIER9() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER9)}}
}

func (p *EXTI_Periph) SWIER10() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER10)}}
}

func (p *EXTI_Periph) SWIER11() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER11)}}
}

func (p *EXTI_Periph) SWIER12() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER12)}}
}

func (p *EXTI_Periph) SWIER13() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER13)}}
}

func (p *EXTI_Periph) SWIER14() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER14)}}
}

func (p *EXTI_Periph) SWIER15() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER15)}}
}

func (p *EXTI_Periph) SWIER16() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER16)}}
}

func (p *EXTI_Periph) SWIER17() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER17)}}
}

func (p *EXTI_Periph) SWIER18() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER18)}}
}

func (p *EXTI_Periph) SWIER19() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER19)}}
}

func (p *EXTI_Periph) SWIER20() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER20)}}
}

func (p *EXTI_Periph) SWIER21() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER21)}}
}

func (p *EXTI_Periph) SWIER22() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER22)}}
}

func (p *EXTI_Periph) SWIER23() SWIER_Mask {
	return SWIER_Mask{mmio.UM32{&p.SWIER.U32, uint32(SWIER23)}}
}

type PR_Bits uint32

func (b PR_Bits) Field(mask PR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PR_Bits) J(v int) PR_Bits {
	return PR_Bits(bits.Make32(v, uint32(mask)))
}

type PR struct{ mmio.U32 }

func (r *PR) Bits(mask PR_Bits) PR_Bits { return PR_Bits(r.U32.Bits(uint32(mask))) }
func (r *PR) StoreBits(mask, b PR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PR) SetBits(mask PR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *PR) ClearBits(mask PR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *PR) Load() PR_Bits             { return PR_Bits(r.U32.Load()) }
func (r *PR) Store(b PR_Bits)           { r.U32.Store(uint32(b)) }

func (r *PR) AtomicSetBits(mask PR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PR) AtomicClearBits(mask PR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type PR_Mask struct{ mmio.UM32 }

func (rm PR_Mask) Load() PR_Bits   { return PR_Bits(rm.UM32.Load()) }
func (rm PR_Mask) Store(b PR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) PR0() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR0)}}
}

func (p *EXTI_Periph) PR1() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR1)}}
}

func (p *EXTI_Periph) PR2() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR2)}}
}

func (p *EXTI_Periph) PR3() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR3)}}
}

func (p *EXTI_Periph) PR4() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR4)}}
}

func (p *EXTI_Periph) PR5() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR5)}}
}

func (p *EXTI_Periph) PR6() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR6)}}
}

func (p *EXTI_Periph) PR7() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR7)}}
}

func (p *EXTI_Periph) PR8() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR8)}}
}

func (p *EXTI_Periph) PR9() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR9)}}
}

func (p *EXTI_Periph) PR10() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR10)}}
}

func (p *EXTI_Periph) PR11() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR11)}}
}

func (p *EXTI_Periph) PR12() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR12)}}
}

func (p *EXTI_Periph) PR13() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR13)}}
}

func (p *EXTI_Periph) PR14() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR14)}}
}

func (p *EXTI_Periph) PR15() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR15)}}
}

func (p *EXTI_Periph) PR16() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR16)}}
}

func (p *EXTI_Periph) PR17() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR17)}}
}

func (p *EXTI_Periph) PR18() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR18)}}
}

func (p *EXTI_Periph) PR19() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR19)}}
}

func (p *EXTI_Periph) PR20() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR20)}}
}

func (p *EXTI_Periph) PR21() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR21)}}
}

func (p *EXTI_Periph) PR22() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR22)}}
}

func (p *EXTI_Periph) PR23() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR23)}}
}
