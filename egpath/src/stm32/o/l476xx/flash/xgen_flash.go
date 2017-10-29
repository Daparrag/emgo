package flash

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type FLASH_Periph struct {
	ACR       ACR
	PDKEYR    PDKEYR
	KEYR      KEYR
	OPTKEYR   OPTKEYR
	SR        SR
	CR        CR
	ECCR      ECCR
	RESERVED1 RESERVED1
	OPTR      OPTR
	PCROP1SR  PCROP1SR
	PCROP1ER  PCROP1ER
	WRP1AR    WRP1AR
	WRP1BR    WRP1BR
	_         [4]uint32
	PCROP2SR  PCROP2SR
	PCROP2ER  PCROP2ER
	WRP2AR    WRP2AR
	WRP2BR    WRP2BR
}

func (p *FLASH_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FLASH = (*FLASH_Periph)(unsafe.Pointer(uintptr(mmap.FLASH_R_BASE)))

type ACR_Bits uint32

func (b ACR_Bits) Field(mask ACR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ACR_Bits) J(v int) ACR_Bits {
	return ACR_Bits(bits.Make32(v, uint32(mask)))
}

type ACR struct{ mmio.U32 }

func (r *ACR) Bits(mask ACR_Bits) ACR_Bits { return ACR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ACR) StoreBits(mask, b ACR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ACR) SetBits(mask ACR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ACR) ClearBits(mask ACR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ACR) Load() ACR_Bits              { return ACR_Bits(r.U32.Load()) }
func (r *ACR) Store(b ACR_Bits)            { r.U32.Store(uint32(b)) }

func (r *ACR) AtomicSetBits(mask ACR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ACR) AtomicClearBits(mask ACR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type ACR_Mask struct{ mmio.UM32 }

func (rm ACR_Mask) Load() ACR_Bits   { return ACR_Bits(rm.UM32.Load()) }
func (rm ACR_Mask) Store(b ACR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) LATENCY() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(LATENCY)}}
}

func (p *FLASH_Periph) PRFTEN() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(PRFTEN)}}
}

func (p *FLASH_Periph) ICEN() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(ICEN)}}
}

func (p *FLASH_Periph) DCEN() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(DCEN)}}
}

func (p *FLASH_Periph) ICRST() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(ICRST)}}
}

func (p *FLASH_Periph) DCRST() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(DCRST)}}
}

func (p *FLASH_Periph) RUN_PD() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(RUN_PD)}}
}

func (p *FLASH_Periph) SLEEP_PD() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(SLEEP_PD)}}
}

type PDKEYR_Bits uint32

func (b PDKEYR_Bits) Field(mask PDKEYR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PDKEYR_Bits) J(v int) PDKEYR_Bits {
	return PDKEYR_Bits(bits.Make32(v, uint32(mask)))
}

type PDKEYR struct{ mmio.U32 }

func (r *PDKEYR) Bits(mask PDKEYR_Bits) PDKEYR_Bits { return PDKEYR_Bits(r.U32.Bits(uint32(mask))) }
func (r *PDKEYR) StoreBits(mask, b PDKEYR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PDKEYR) SetBits(mask PDKEYR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *PDKEYR) ClearBits(mask PDKEYR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *PDKEYR) Load() PDKEYR_Bits                 { return PDKEYR_Bits(r.U32.Load()) }
func (r *PDKEYR) Store(b PDKEYR_Bits)               { r.U32.Store(uint32(b)) }

func (r *PDKEYR) AtomicSetBits(mask PDKEYR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PDKEYR) AtomicClearBits(mask PDKEYR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type PDKEYR_Mask struct{ mmio.UM32 }

func (rm PDKEYR_Mask) Load() PDKEYR_Bits   { return PDKEYR_Bits(rm.UM32.Load()) }
func (rm PDKEYR_Mask) Store(b PDKEYR_Bits) { rm.UM32.Store(uint32(b)) }

type KEYR_Bits uint32

func (b KEYR_Bits) Field(mask KEYR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask KEYR_Bits) J(v int) KEYR_Bits {
	return KEYR_Bits(bits.Make32(v, uint32(mask)))
}

type KEYR struct{ mmio.U32 }

func (r *KEYR) Bits(mask KEYR_Bits) KEYR_Bits { return KEYR_Bits(r.U32.Bits(uint32(mask))) }
func (r *KEYR) StoreBits(mask, b KEYR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *KEYR) SetBits(mask KEYR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *KEYR) ClearBits(mask KEYR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *KEYR) Load() KEYR_Bits               { return KEYR_Bits(r.U32.Load()) }
func (r *KEYR) Store(b KEYR_Bits)             { r.U32.Store(uint32(b)) }

func (r *KEYR) AtomicSetBits(mask KEYR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *KEYR) AtomicClearBits(mask KEYR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type KEYR_Mask struct{ mmio.UM32 }

func (rm KEYR_Mask) Load() KEYR_Bits   { return KEYR_Bits(rm.UM32.Load()) }
func (rm KEYR_Mask) Store(b KEYR_Bits) { rm.UM32.Store(uint32(b)) }

type OPTKEYR_Bits uint32

func (b OPTKEYR_Bits) Field(mask OPTKEYR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OPTKEYR_Bits) J(v int) OPTKEYR_Bits {
	return OPTKEYR_Bits(bits.Make32(v, uint32(mask)))
}

type OPTKEYR struct{ mmio.U32 }

func (r *OPTKEYR) Bits(mask OPTKEYR_Bits) OPTKEYR_Bits { return OPTKEYR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OPTKEYR) StoreBits(mask, b OPTKEYR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OPTKEYR) SetBits(mask OPTKEYR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *OPTKEYR) ClearBits(mask OPTKEYR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *OPTKEYR) Load() OPTKEYR_Bits                  { return OPTKEYR_Bits(r.U32.Load()) }
func (r *OPTKEYR) Store(b OPTKEYR_Bits)                { r.U32.Store(uint32(b)) }

func (r *OPTKEYR) AtomicSetBits(mask OPTKEYR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *OPTKEYR) AtomicClearBits(mask OPTKEYR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type OPTKEYR_Mask struct{ mmio.UM32 }

func (rm OPTKEYR_Mask) Load() OPTKEYR_Bits   { return OPTKEYR_Bits(rm.UM32.Load()) }
func (rm OPTKEYR_Mask) Store(b OPTKEYR_Bits) { rm.UM32.Store(uint32(b)) }

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

func (p *FLASH_Periph) EOP() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(EOP)}}
}

func (p *FLASH_Periph) OPERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(OPERR)}}
}

func (p *FLASH_Periph) PROGERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(PROGERR)}}
}

func (p *FLASH_Periph) WRPERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(WRPERR)}}
}

func (p *FLASH_Periph) PGAERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(PGAERR)}}
}

func (p *FLASH_Periph) SIZERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(SIZERR)}}
}

func (p *FLASH_Periph) PGSERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(PGSERR)}}
}

func (p *FLASH_Periph) MISERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(MISERR)}}
}

func (p *FLASH_Periph) FASTERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(FASTERR)}}
}

func (p *FLASH_Periph) RDERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(RDERR)}}
}

func (p *FLASH_Periph) OPTVERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(OPTVERR)}}
}

func (p *FLASH_Periph) BSY() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(BSY)}}
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

func (p *FLASH_Periph) PG() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PG)}}
}

func (p *FLASH_Periph) PER() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PER)}}
}

func (p *FLASH_Periph) MER1() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(MER1)}}
}

func (p *FLASH_Periph) PNB() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PNB)}}
}

func (p *FLASH_Periph) BKER() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(BKER)}}
}

func (p *FLASH_Periph) MER2() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(MER2)}}
}

func (p *FLASH_Periph) STRT() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(STRT)}}
}

func (p *FLASH_Periph) OPTSTRT() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(OPTSTRT)}}
}

func (p *FLASH_Periph) FSTPG() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(FSTPG)}}
}

func (p *FLASH_Periph) EOPIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(EOPIE)}}
}

func (p *FLASH_Periph) ERRIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ERRIE)}}
}

func (p *FLASH_Periph) RDERRIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(RDERRIE)}}
}

func (p *FLASH_Periph) OBL_LAUNCH() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(OBL_LAUNCH)}}
}

func (p *FLASH_Periph) OPTLOCK() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(OPTLOCK)}}
}

func (p *FLASH_Periph) LOCK() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(LOCK)}}
}

type ECCR_Bits uint32

func (b ECCR_Bits) Field(mask ECCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ECCR_Bits) J(v int) ECCR_Bits {
	return ECCR_Bits(bits.Make32(v, uint32(mask)))
}

type ECCR struct{ mmio.U32 }

func (r *ECCR) Bits(mask ECCR_Bits) ECCR_Bits { return ECCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ECCR) StoreBits(mask, b ECCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ECCR) SetBits(mask ECCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *ECCR) ClearBits(mask ECCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *ECCR) Load() ECCR_Bits               { return ECCR_Bits(r.U32.Load()) }
func (r *ECCR) Store(b ECCR_Bits)             { r.U32.Store(uint32(b)) }

func (r *ECCR) AtomicSetBits(mask ECCR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ECCR) AtomicClearBits(mask ECCR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type ECCR_Mask struct{ mmio.UM32 }

func (rm ECCR_Mask) Load() ECCR_Bits   { return ECCR_Bits(rm.UM32.Load()) }
func (rm ECCR_Mask) Store(b ECCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) ADDR_ECC() ECCR_Mask {
	return ECCR_Mask{mmio.UM32{&p.ECCR.U32, uint32(ADDR_ECC)}}
}

func (p *FLASH_Periph) BK_ECC() ECCR_Mask {
	return ECCR_Mask{mmio.UM32{&p.ECCR.U32, uint32(BK_ECC)}}
}

func (p *FLASH_Periph) SYSF_ECC() ECCR_Mask {
	return ECCR_Mask{mmio.UM32{&p.ECCR.U32, uint32(SYSF_ECC)}}
}

func (p *FLASH_Periph) ECCIE() ECCR_Mask {
	return ECCR_Mask{mmio.UM32{&p.ECCR.U32, uint32(ECCIE)}}
}

func (p *FLASH_Periph) ECCC() ECCR_Mask {
	return ECCR_Mask{mmio.UM32{&p.ECCR.U32, uint32(ECCC)}}
}

func (p *FLASH_Periph) ECCD() ECCR_Mask {
	return ECCR_Mask{mmio.UM32{&p.ECCR.U32, uint32(ECCD)}}
}

type RESERVED1_Bits uint32

func (b RESERVED1_Bits) Field(mask RESERVED1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESERVED1_Bits) J(v int) RESERVED1_Bits {
	return RESERVED1_Bits(bits.Make32(v, uint32(mask)))
}

type RESERVED1 struct{ mmio.U32 }

func (r *RESERVED1) Bits(mask RESERVED1_Bits) RESERVED1_Bits {
	return RESERVED1_Bits(r.U32.Bits(uint32(mask)))
}
func (r *RESERVED1) StoreBits(mask, b RESERVED1_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RESERVED1) SetBits(mask RESERVED1_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *RESERVED1) ClearBits(mask RESERVED1_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *RESERVED1) Load() RESERVED1_Bits             { return RESERVED1_Bits(r.U32.Load()) }
func (r *RESERVED1) Store(b RESERVED1_Bits)           { r.U32.Store(uint32(b)) }

func (r *RESERVED1) AtomicSetBits(mask RESERVED1_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RESERVED1) AtomicClearBits(mask RESERVED1_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type RESERVED1_Mask struct{ mmio.UM32 }

func (rm RESERVED1_Mask) Load() RESERVED1_Bits   { return RESERVED1_Bits(rm.UM32.Load()) }
func (rm RESERVED1_Mask) Store(b RESERVED1_Bits) { rm.UM32.Store(uint32(b)) }

type OPTR_Bits uint32

func (b OPTR_Bits) Field(mask OPTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OPTR_Bits) J(v int) OPTR_Bits {
	return OPTR_Bits(bits.Make32(v, uint32(mask)))
}

type OPTR struct{ mmio.U32 }

func (r *OPTR) Bits(mask OPTR_Bits) OPTR_Bits { return OPTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OPTR) StoreBits(mask, b OPTR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OPTR) SetBits(mask OPTR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *OPTR) ClearBits(mask OPTR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *OPTR) Load() OPTR_Bits               { return OPTR_Bits(r.U32.Load()) }
func (r *OPTR) Store(b OPTR_Bits)             { r.U32.Store(uint32(b)) }

func (r *OPTR) AtomicSetBits(mask OPTR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *OPTR) AtomicClearBits(mask OPTR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type OPTR_Mask struct{ mmio.UM32 }

func (rm OPTR_Mask) Load() OPTR_Bits   { return OPTR_Bits(rm.UM32.Load()) }
func (rm OPTR_Mask) Store(b OPTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) RDP() OPTR_Mask {
	return OPTR_Mask{mmio.UM32{&p.OPTR.U32, uint32(RDP)}}
}

func (p *FLASH_Periph) BOR_LEV() OPTR_Mask {
	return OPTR_Mask{mmio.UM32{&p.OPTR.U32, uint32(BOR_LEV)}}
}

func (p *FLASH_Periph) nRST_STOP() OPTR_Mask {
	return OPTR_Mask{mmio.UM32{&p.OPTR.U32, uint32(nRST_STOP)}}
}

func (p *FLASH_Periph) nRST_STDBY() OPTR_Mask {
	return OPTR_Mask{mmio.UM32{&p.OPTR.U32, uint32(nRST_STDBY)}}
}

func (p *FLASH_Periph) nRST_SHDW() OPTR_Mask {
	return OPTR_Mask{mmio.UM32{&p.OPTR.U32, uint32(nRST_SHDW)}}
}

func (p *FLASH_Periph) IWDG_SW() OPTR_Mask {
	return OPTR_Mask{mmio.UM32{&p.OPTR.U32, uint32(IWDG_SW)}}
}

func (p *FLASH_Periph) IWDG_STOP() OPTR_Mask {
	return OPTR_Mask{mmio.UM32{&p.OPTR.U32, uint32(IWDG_STOP)}}
}

func (p *FLASH_Periph) IWDG_STDBY() OPTR_Mask {
	return OPTR_Mask{mmio.UM32{&p.OPTR.U32, uint32(IWDG_STDBY)}}
}

func (p *FLASH_Periph) WWDG_SW() OPTR_Mask {
	return OPTR_Mask{mmio.UM32{&p.OPTR.U32, uint32(WWDG_SW)}}
}

func (p *FLASH_Periph) BFB2() OPTR_Mask {
	return OPTR_Mask{mmio.UM32{&p.OPTR.U32, uint32(BFB2)}}
}

func (p *FLASH_Periph) DUALBANK() OPTR_Mask {
	return OPTR_Mask{mmio.UM32{&p.OPTR.U32, uint32(DUALBANK)}}
}

func (p *FLASH_Periph) nBOOT1() OPTR_Mask {
	return OPTR_Mask{mmio.UM32{&p.OPTR.U32, uint32(nBOOT1)}}
}

func (p *FLASH_Periph) SRAM2_PE() OPTR_Mask {
	return OPTR_Mask{mmio.UM32{&p.OPTR.U32, uint32(SRAM2_PE)}}
}

func (p *FLASH_Periph) SRAM2_RST() OPTR_Mask {
	return OPTR_Mask{mmio.UM32{&p.OPTR.U32, uint32(SRAM2_RST)}}
}

type PCROP1SR_Bits uint32

func (b PCROP1SR_Bits) Field(mask PCROP1SR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PCROP1SR_Bits) J(v int) PCROP1SR_Bits {
	return PCROP1SR_Bits(bits.Make32(v, uint32(mask)))
}

type PCROP1SR struct{ mmio.U32 }

func (r *PCROP1SR) Bits(mask PCROP1SR_Bits) PCROP1SR_Bits {
	return PCROP1SR_Bits(r.U32.Bits(uint32(mask)))
}
func (r *PCROP1SR) StoreBits(mask, b PCROP1SR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PCROP1SR) SetBits(mask PCROP1SR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *PCROP1SR) ClearBits(mask PCROP1SR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *PCROP1SR) Load() PCROP1SR_Bits             { return PCROP1SR_Bits(r.U32.Load()) }
func (r *PCROP1SR) Store(b PCROP1SR_Bits)           { r.U32.Store(uint32(b)) }

func (r *PCROP1SR) AtomicSetBits(mask PCROP1SR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PCROP1SR) AtomicClearBits(mask PCROP1SR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type PCROP1SR_Mask struct{ mmio.UM32 }

func (rm PCROP1SR_Mask) Load() PCROP1SR_Bits   { return PCROP1SR_Bits(rm.UM32.Load()) }
func (rm PCROP1SR_Mask) Store(b PCROP1SR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) PCROP1_STRT() PCROP1SR_Mask {
	return PCROP1SR_Mask{mmio.UM32{&p.PCROP1SR.U32, uint32(PCROP1_STRT)}}
}

type PCROP1ER_Bits uint32

func (b PCROP1ER_Bits) Field(mask PCROP1ER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PCROP1ER_Bits) J(v int) PCROP1ER_Bits {
	return PCROP1ER_Bits(bits.Make32(v, uint32(mask)))
}

type PCROP1ER struct{ mmio.U32 }

func (r *PCROP1ER) Bits(mask PCROP1ER_Bits) PCROP1ER_Bits {
	return PCROP1ER_Bits(r.U32.Bits(uint32(mask)))
}
func (r *PCROP1ER) StoreBits(mask, b PCROP1ER_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PCROP1ER) SetBits(mask PCROP1ER_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *PCROP1ER) ClearBits(mask PCROP1ER_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *PCROP1ER) Load() PCROP1ER_Bits             { return PCROP1ER_Bits(r.U32.Load()) }
func (r *PCROP1ER) Store(b PCROP1ER_Bits)           { r.U32.Store(uint32(b)) }

func (r *PCROP1ER) AtomicSetBits(mask PCROP1ER_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PCROP1ER) AtomicClearBits(mask PCROP1ER_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type PCROP1ER_Mask struct{ mmio.UM32 }

func (rm PCROP1ER_Mask) Load() PCROP1ER_Bits   { return PCROP1ER_Bits(rm.UM32.Load()) }
func (rm PCROP1ER_Mask) Store(b PCROP1ER_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) PCROP1_END() PCROP1ER_Mask {
	return PCROP1ER_Mask{mmio.UM32{&p.PCROP1ER.U32, uint32(PCROP1_END)}}
}

func (p *FLASH_Periph) PCROP_RDP() PCROP1ER_Mask {
	return PCROP1ER_Mask{mmio.UM32{&p.PCROP1ER.U32, uint32(PCROP_RDP)}}
}

type WRP1AR_Bits uint32

func (b WRP1AR_Bits) Field(mask WRP1AR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP1AR_Bits) J(v int) WRP1AR_Bits {
	return WRP1AR_Bits(bits.Make32(v, uint32(mask)))
}

type WRP1AR struct{ mmio.U32 }

func (r *WRP1AR) Bits(mask WRP1AR_Bits) WRP1AR_Bits { return WRP1AR_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRP1AR) StoreBits(mask, b WRP1AR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRP1AR) SetBits(mask WRP1AR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *WRP1AR) ClearBits(mask WRP1AR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *WRP1AR) Load() WRP1AR_Bits                 { return WRP1AR_Bits(r.U32.Load()) }
func (r *WRP1AR) Store(b WRP1AR_Bits)               { r.U32.Store(uint32(b)) }

func (r *WRP1AR) AtomicSetBits(mask WRP1AR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WRP1AR) AtomicClearBits(mask WRP1AR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type WRP1AR_Mask struct{ mmio.UM32 }

func (rm WRP1AR_Mask) Load() WRP1AR_Bits   { return WRP1AR_Bits(rm.UM32.Load()) }
func (rm WRP1AR_Mask) Store(b WRP1AR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) WRP1A_STRT() WRP1AR_Mask {
	return WRP1AR_Mask{mmio.UM32{&p.WRP1AR.U32, uint32(WRP1A_STRT)}}
}

func (p *FLASH_Periph) WRP1A_END() WRP1AR_Mask {
	return WRP1AR_Mask{mmio.UM32{&p.WRP1AR.U32, uint32(WRP1A_END)}}
}

type WRP1BR_Bits uint32

func (b WRP1BR_Bits) Field(mask WRP1BR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP1BR_Bits) J(v int) WRP1BR_Bits {
	return WRP1BR_Bits(bits.Make32(v, uint32(mask)))
}

type WRP1BR struct{ mmio.U32 }

func (r *WRP1BR) Bits(mask WRP1BR_Bits) WRP1BR_Bits { return WRP1BR_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRP1BR) StoreBits(mask, b WRP1BR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRP1BR) SetBits(mask WRP1BR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *WRP1BR) ClearBits(mask WRP1BR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *WRP1BR) Load() WRP1BR_Bits                 { return WRP1BR_Bits(r.U32.Load()) }
func (r *WRP1BR) Store(b WRP1BR_Bits)               { r.U32.Store(uint32(b)) }

func (r *WRP1BR) AtomicSetBits(mask WRP1BR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WRP1BR) AtomicClearBits(mask WRP1BR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type WRP1BR_Mask struct{ mmio.UM32 }

func (rm WRP1BR_Mask) Load() WRP1BR_Bits   { return WRP1BR_Bits(rm.UM32.Load()) }
func (rm WRP1BR_Mask) Store(b WRP1BR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) WRP1B_STRT() WRP1BR_Mask {
	return WRP1BR_Mask{mmio.UM32{&p.WRP1BR.U32, uint32(WRP1B_STRT)}}
}

func (p *FLASH_Periph) WRP1B_END() WRP1BR_Mask {
	return WRP1BR_Mask{mmio.UM32{&p.WRP1BR.U32, uint32(WRP1B_END)}}
}

type PCROP2SR_Bits uint32

func (b PCROP2SR_Bits) Field(mask PCROP2SR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PCROP2SR_Bits) J(v int) PCROP2SR_Bits {
	return PCROP2SR_Bits(bits.Make32(v, uint32(mask)))
}

type PCROP2SR struct{ mmio.U32 }

func (r *PCROP2SR) Bits(mask PCROP2SR_Bits) PCROP2SR_Bits {
	return PCROP2SR_Bits(r.U32.Bits(uint32(mask)))
}
func (r *PCROP2SR) StoreBits(mask, b PCROP2SR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PCROP2SR) SetBits(mask PCROP2SR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *PCROP2SR) ClearBits(mask PCROP2SR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *PCROP2SR) Load() PCROP2SR_Bits             { return PCROP2SR_Bits(r.U32.Load()) }
func (r *PCROP2SR) Store(b PCROP2SR_Bits)           { r.U32.Store(uint32(b)) }

func (r *PCROP2SR) AtomicSetBits(mask PCROP2SR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PCROP2SR) AtomicClearBits(mask PCROP2SR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type PCROP2SR_Mask struct{ mmio.UM32 }

func (rm PCROP2SR_Mask) Load() PCROP2SR_Bits   { return PCROP2SR_Bits(rm.UM32.Load()) }
func (rm PCROP2SR_Mask) Store(b PCROP2SR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) PCROP2_STRT() PCROP2SR_Mask {
	return PCROP2SR_Mask{mmio.UM32{&p.PCROP2SR.U32, uint32(PCROP2_STRT)}}
}

type PCROP2ER_Bits uint32

func (b PCROP2ER_Bits) Field(mask PCROP2ER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PCROP2ER_Bits) J(v int) PCROP2ER_Bits {
	return PCROP2ER_Bits(bits.Make32(v, uint32(mask)))
}

type PCROP2ER struct{ mmio.U32 }

func (r *PCROP2ER) Bits(mask PCROP2ER_Bits) PCROP2ER_Bits {
	return PCROP2ER_Bits(r.U32.Bits(uint32(mask)))
}
func (r *PCROP2ER) StoreBits(mask, b PCROP2ER_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PCROP2ER) SetBits(mask PCROP2ER_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *PCROP2ER) ClearBits(mask PCROP2ER_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *PCROP2ER) Load() PCROP2ER_Bits             { return PCROP2ER_Bits(r.U32.Load()) }
func (r *PCROP2ER) Store(b PCROP2ER_Bits)           { r.U32.Store(uint32(b)) }

func (r *PCROP2ER) AtomicSetBits(mask PCROP2ER_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PCROP2ER) AtomicClearBits(mask PCROP2ER_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type PCROP2ER_Mask struct{ mmio.UM32 }

func (rm PCROP2ER_Mask) Load() PCROP2ER_Bits   { return PCROP2ER_Bits(rm.UM32.Load()) }
func (rm PCROP2ER_Mask) Store(b PCROP2ER_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) PCROP2_END() PCROP2ER_Mask {
	return PCROP2ER_Mask{mmio.UM32{&p.PCROP2ER.U32, uint32(PCROP2_END)}}
}

type WRP2AR_Bits uint32

func (b WRP2AR_Bits) Field(mask WRP2AR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP2AR_Bits) J(v int) WRP2AR_Bits {
	return WRP2AR_Bits(bits.Make32(v, uint32(mask)))
}

type WRP2AR struct{ mmio.U32 }

func (r *WRP2AR) Bits(mask WRP2AR_Bits) WRP2AR_Bits { return WRP2AR_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRP2AR) StoreBits(mask, b WRP2AR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRP2AR) SetBits(mask WRP2AR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *WRP2AR) ClearBits(mask WRP2AR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *WRP2AR) Load() WRP2AR_Bits                 { return WRP2AR_Bits(r.U32.Load()) }
func (r *WRP2AR) Store(b WRP2AR_Bits)               { r.U32.Store(uint32(b)) }

func (r *WRP2AR) AtomicSetBits(mask WRP2AR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WRP2AR) AtomicClearBits(mask WRP2AR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type WRP2AR_Mask struct{ mmio.UM32 }

func (rm WRP2AR_Mask) Load() WRP2AR_Bits   { return WRP2AR_Bits(rm.UM32.Load()) }
func (rm WRP2AR_Mask) Store(b WRP2AR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) WRP2A_STRT() WRP2AR_Mask {
	return WRP2AR_Mask{mmio.UM32{&p.WRP2AR.U32, uint32(WRP2A_STRT)}}
}

func (p *FLASH_Periph) WRP2A_END() WRP2AR_Mask {
	return WRP2AR_Mask{mmio.UM32{&p.WRP2AR.U32, uint32(WRP2A_END)}}
}

type WRP2BR_Bits uint32

func (b WRP2BR_Bits) Field(mask WRP2BR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP2BR_Bits) J(v int) WRP2BR_Bits {
	return WRP2BR_Bits(bits.Make32(v, uint32(mask)))
}

type WRP2BR struct{ mmio.U32 }

func (r *WRP2BR) Bits(mask WRP2BR_Bits) WRP2BR_Bits { return WRP2BR_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRP2BR) StoreBits(mask, b WRP2BR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRP2BR) SetBits(mask WRP2BR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *WRP2BR) ClearBits(mask WRP2BR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *WRP2BR) Load() WRP2BR_Bits                 { return WRP2BR_Bits(r.U32.Load()) }
func (r *WRP2BR) Store(b WRP2BR_Bits)               { r.U32.Store(uint32(b)) }

func (r *WRP2BR) AtomicSetBits(mask WRP2BR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WRP2BR) AtomicClearBits(mask WRP2BR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type WRP2BR_Mask struct{ mmio.UM32 }

func (rm WRP2BR_Mask) Load() WRP2BR_Bits   { return WRP2BR_Bits(rm.UM32.Load()) }
func (rm WRP2BR_Mask) Store(b WRP2BR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) WRP2B_STRT() WRP2BR_Mask {
	return WRP2BR_Mask{mmio.UM32{&p.WRP2BR.U32, uint32(WRP2B_STRT)}}
}

func (p *FLASH_Periph) WRP2B_END() WRP2BR_Mask {
	return WRP2BR_Mask{mmio.UM32{&p.WRP2BR.U32, uint32(WRP2B_END)}}
}
