// Peripheral: HRTIM_Master_Periph  High resolution Timer (HRTIM).
// Instances:
// Registers:
//  0x00 32  MCR    HRTIM Master Timer control register.
//  0x04 32  MISR   HRTIM Master Timer interrupt status register.
//  0x08 32  MICR   HRTIM Master Timer interrupt clear register.
//  0x0C 32  MDIER  HRTIM Master Timer DMA/interrupt enable register.
//  0x10 32  MCNTR  HRTIM Master Timer counter register.
//  0x14 32  MPER   HRTIM Master Timer period register.
//  0x18 32  MREP   HRTIM Master Timer repetition register.
//  0x1C 32  MCMP1R HRTIM Master Timer compare 1 register.
//  0x24 32  MCMP2R HRTIM Master Timer compare 2 register.
//  0x28 32  MCMP3R HRTIM Master Timer compare 3 register.
//  0x2C 32  MCMP4R HRTIM Master Timer compare 4 register.
// Import:
//  stm32/o/f303xe/mmap
package hrtim

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	CK_PSC     MCR_Bits = 0x07 << 0  //+ Prescaler mask.
	CK_PSC_0   MCR_Bits = 0x01 << 0  //  Prescaler bit 0.
	CK_PSC_1   MCR_Bits = 0x02 << 0  //  Prescaler bit 1.
	CK_PSC_2   MCR_Bits = 0x04 << 0  //  Prescaler bit 2.
	CONT       MCR_Bits = 0x01 << 3  //+ Continuous mode.
	RETRIG     MCR_Bits = 0x01 << 4  //+ Rettrigreable mode.
	HALF       MCR_Bits = 0x01 << 5  //+ Half mode.
	SYNC_IN    MCR_Bits = 0x03 << 8  //+ Synchronization input master.
	SYNC_IN_0  MCR_Bits = 0x01 << 8  //  Synchronization input bit 0.
	SYNC_IN_1  MCR_Bits = 0x02 << 8  //  Synchronization input bit 1.
	SYNCRSTM   MCR_Bits = 0x01 << 10 //+ Synchronization reset master.
	SYNCSTRTM  MCR_Bits = 0x01 << 11 //+ Synchronization start master.
	SYNC_OUT   MCR_Bits = 0x03 << 12 //+ Synchronization output master.
	SYNC_OUT_0 MCR_Bits = 0x01 << 12 //  Synchronization output bit 0.
	SYNC_OUT_1 MCR_Bits = 0x02 << 12 //  Synchronization output bit 1.
	SYNC_SRC   MCR_Bits = 0x03 << 14 //+ Synchronization source.
	SYNC_SRC_0 MCR_Bits = 0x01 << 14 //  Synchronization source bit 0.
	SYNC_SRC_1 MCR_Bits = 0x02 << 14 //  Synchronization source bit 1.
	MCEN       MCR_Bits = 0x01 << 16 //+ Master counter enable.
	TACEN      MCR_Bits = 0x01 << 17 //+ Timer A counter enable.
	TBCEN      MCR_Bits = 0x01 << 18 //+ Timer B counter enable.
	TCCEN      MCR_Bits = 0x01 << 19 //+ Timer C counter enable.
	TDCEN      MCR_Bits = 0x01 << 20 //+ Timer D counter enable.
	TECEN      MCR_Bits = 0x01 << 21 //+ Timer E counter enable.
	DACSYNC    MCR_Bits = 0x03 << 25 //+ DAC synchronization mask.
	DACSYNC_0  MCR_Bits = 0x01 << 25 //  DAC synchronization bit 0.
	DACSYNC_1  MCR_Bits = 0x02 << 25 //  DAC synchronization bit 1.
	PREEN      MCR_Bits = 0x01 << 27 //+ Master preload enable.
	MREPU      MCR_Bits = 0x01 << 29 //+ Master repetition update.
	BRSTDMA    MCR_Bits = 0x03 << 30 //+ Burst DMA update.
	BRSTDMA_0  MCR_Bits = 0x01 << 30 //  Burst DMA update bit 0.
	BRSTDMA_1  MCR_Bits = 0x02 << 30 //  Burst DMA update bit 1.
)

const (
	CK_PSCn    = 0
	CONTn      = 3
	RETRIGn    = 4
	HALFn      = 5
	SYNC_INn   = 8
	SYNCRSTMn  = 10
	SYNCSTRTMn = 11
	SYNC_OUTn  = 12
	SYNC_SRCn  = 14
	MCENn      = 16
	TACENn     = 17
	TBCENn     = 18
	TCCENn     = 19
	TDCENn     = 20
	TECENn     = 21
	DACSYNCn   = 25
	PREENn     = 27
	MREPUn     = 29
	BRSTDMAn   = 30
)

const (
	MCMP1 MISR_Bits = 0x01 << 0 //+ Master compare 1 interrupt flag.
	MCMP2 MISR_Bits = 0x01 << 1 //+ Master compare 2 interrupt flag.
	MCMP3 MISR_Bits = 0x01 << 2 //+ Master compare 3 interrupt flag.
	MCMP4 MISR_Bits = 0x01 << 3 //+ Master compare 4 interrupt flag.
	MREP  MISR_Bits = 0x01 << 4 //+ Master Repetition interrupt flag.
	SYNC  MISR_Bits = 0x01 << 5 //+ Synchronization input interrupt flag.
	MUPD  MISR_Bits = 0x01 << 6 //+ Master update interrupt flag.
)

const (
	MCMP1n = 0
	MCMP2n = 1
	MCMP3n = 2
	MCMP4n = 3
	MREPn  = 4
	SYNCn  = 5
	MUPDn  = 6
)

const (
	MCMP1 MICR_Bits = 0x01 << 0 //+ Master compare 1 interrupt flag clear.
	MCMP2 MICR_Bits = 0x01 << 1 //+ Master compare 2 interrupt flag clear.
	MCMP3 MICR_Bits = 0x01 << 2 //+ Master compare 3 interrupt flag clear.
	MCMP4 MICR_Bits = 0x01 << 3 //+ Master compare 4 interrupt flag clear.
	MREP  MICR_Bits = 0x01 << 4 //+ Master Repetition interrupt flag clear.
	SYNC  MICR_Bits = 0x01 << 5 //+ Synchronization input interrupt flag clear.
	MUPD  MICR_Bits = 0x01 << 6 //+ Master update interrupt flag clear.
)

const (
	MCMP1n = 0
	MCMP2n = 1
	MCMP3n = 2
	MCMP4n = 3
	MREPn  = 4
	SYNCn  = 5
	MUPDn  = 6
)

const (
	MCMP1IE MDIER_Bits = 0x01 << 0  //+ Master compare 1 interrupt enable.
	MCMP2IE MDIER_Bits = 0x01 << 1  //+ Master compare 2 interrupt enable.
	MCMP3IE MDIER_Bits = 0x01 << 2  //+ Master compare 3 interrupt enable.
	MCMP4IE MDIER_Bits = 0x01 << 3  //+ Master compare 4 interrupt enable.
	MREPIE  MDIER_Bits = 0x01 << 4  //+ Master Repetition interrupt enable.
	SYNCIE  MDIER_Bits = 0x01 << 5  //+ Synchronization input interrupt enable.
	MUPDIE  MDIER_Bits = 0x01 << 6  //+ Master update interrupt enable.
	MCMP1DE MDIER_Bits = 0x01 << 16 //+ Master compare 1 DMA enable.
	MCMP2DE MDIER_Bits = 0x01 << 17 //+ Master compare 2 DMA enable.
	MCMP3DE MDIER_Bits = 0x01 << 18 //+ Master compare 3 DMA enable.
	MCMP4DE MDIER_Bits = 0x01 << 19 //+ Master compare 4 DMA enable.
	MREPDE  MDIER_Bits = 0x01 << 20 //+ Master Repetition DMA enable.
	SYNCDE  MDIER_Bits = 0x01 << 21 //+ Synchronization input DMA enable.
	MUPDDE  MDIER_Bits = 0x01 << 22 //+ Master update DMA enable.
)

const (
	MCMP1IEn = 0
	MCMP2IEn = 1
	MCMP3IEn = 2
	MCMP4IEn = 3
	MREPIEn  = 4
	SYNCIEn  = 5
	MUPDIEn  = 6
	MCMP1DEn = 16
	MCMP2DEn = 17
	MCMP3DEn = 18
	MCMP4DEn = 19
	MREPDEn  = 20
	SYNCDEn  = 21
	MUPDDEn  = 22
)

const (
	MCMP1R MCMP1R_Bits = 0xFFFFFFFF << 0 //+ Compare Value.
	MCMP2R MCMP1R_Bits = 0xFFFFFFFF << 0 //  Compare Value.
	MCMP3R MCMP1R_Bits = 0xFFFFFFFF << 0 //  Compare Value.
	MCMP4R MCMP1R_Bits = 0xFFFFFFFF << 0 //  Compare Value.
)

const (
	MCMP1Rn = 0
)