package dbgmcu

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type DBGMCU_Periph struct {
	IDCODE IDCODE
	CR     CR
	APB1FZ APB1FZ
	APB2FZ APB2FZ
}

func (p *DBGMCU_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var DBGMCU = (*DBGMCU_Periph)(unsafe.Pointer(uintptr(mmap.DBGMCU_BASE)))

type IDCODE_Bits uint32

func (b IDCODE_Bits) Field(mask IDCODE_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IDCODE_Bits) J(v int) IDCODE_Bits {
	return IDCODE_Bits(bits.Make32(v, uint32(mask)))
}

type IDCODE struct{ mmio.U32 }

func (r *IDCODE) Bits(mask IDCODE_Bits) IDCODE_Bits { return IDCODE_Bits(r.U32.Bits(uint32(mask))) }
func (r *IDCODE) StoreBits(mask, b IDCODE_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IDCODE) SetBits(mask IDCODE_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *IDCODE) ClearBits(mask IDCODE_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *IDCODE) Load() IDCODE_Bits                 { return IDCODE_Bits(r.U32.Load()) }
func (r *IDCODE) Store(b IDCODE_Bits)               { r.U32.Store(uint32(b)) }

func (r *IDCODE) AtomicSetBits(mask IDCODE_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *IDCODE) AtomicClearBits(mask IDCODE_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type IDCODE_Mask struct{ mmio.UM32 }

func (rm IDCODE_Mask) Load() IDCODE_Bits   { return IDCODE_Bits(rm.UM32.Load()) }
func (rm IDCODE_Mask) Store(b IDCODE_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DBGMCU_Periph) DEV_ID() IDCODE_Mask {
	return IDCODE_Mask{mmio.UM32{&p.IDCODE.U32, uint32(DEV_ID)}}
}

func (p *DBGMCU_Periph) REV_ID() IDCODE_Mask {
	return IDCODE_Mask{mmio.UM32{&p.IDCODE.U32, uint32(REV_ID)}}
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

func (p *DBGMCU_Periph) DBG_SLEEP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_SLEEP)}}
}

func (p *DBGMCU_Periph) DBG_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_STOP)}}
}

func (p *DBGMCU_Periph) DBG_STANDBY() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_STANDBY)}}
}

func (p *DBGMCU_Periph) TRACE_IOEN() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TRACE_IOEN)}}
}

func (p *DBGMCU_Periph) TRACE_MODE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TRACE_MODE)}}
}

type APB1FZ_Bits uint32

func (b APB1FZ_Bits) Field(mask APB1FZ_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1FZ_Bits) J(v int) APB1FZ_Bits {
	return APB1FZ_Bits(bits.Make32(v, uint32(mask)))
}

type APB1FZ struct{ mmio.U32 }

func (r *APB1FZ) Bits(mask APB1FZ_Bits) APB1FZ_Bits { return APB1FZ_Bits(r.U32.Bits(uint32(mask))) }
func (r *APB1FZ) StoreBits(mask, b APB1FZ_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *APB1FZ) SetBits(mask APB1FZ_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *APB1FZ) ClearBits(mask APB1FZ_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *APB1FZ) Load() APB1FZ_Bits                 { return APB1FZ_Bits(r.U32.Load()) }
func (r *APB1FZ) Store(b APB1FZ_Bits)               { r.U32.Store(uint32(b)) }

func (r *APB1FZ) AtomicSetBits(mask APB1FZ_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *APB1FZ) AtomicClearBits(mask APB1FZ_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type APB1FZ_Mask struct{ mmio.UM32 }

func (rm APB1FZ_Mask) Load() APB1FZ_Bits   { return APB1FZ_Bits(rm.UM32.Load()) }
func (rm APB1FZ_Mask) Store(b APB1FZ_Bits) { rm.UM32.Store(uint32(b)) }

type APB2FZ_Bits uint32

func (b APB2FZ_Bits) Field(mask APB2FZ_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2FZ_Bits) J(v int) APB2FZ_Bits {
	return APB2FZ_Bits(bits.Make32(v, uint32(mask)))
}

type APB2FZ struct{ mmio.U32 }

func (r *APB2FZ) Bits(mask APB2FZ_Bits) APB2FZ_Bits { return APB2FZ_Bits(r.U32.Bits(uint32(mask))) }
func (r *APB2FZ) StoreBits(mask, b APB2FZ_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *APB2FZ) SetBits(mask APB2FZ_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *APB2FZ) ClearBits(mask APB2FZ_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *APB2FZ) Load() APB2FZ_Bits                 { return APB2FZ_Bits(r.U32.Load()) }
func (r *APB2FZ) Store(b APB2FZ_Bits)               { r.U32.Store(uint32(b)) }

func (r *APB2FZ) AtomicSetBits(mask APB2FZ_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *APB2FZ) AtomicClearBits(mask APB2FZ_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type APB2FZ_Mask struct{ mmio.UM32 }

func (rm APB2FZ_Mask) Load() APB2FZ_Bits   { return APB2FZ_Bits(rm.UM32.Load()) }
func (rm APB2FZ_Mask) Store(b APB2FZ_Bits) { rm.UM32.Store(uint32(b)) }
