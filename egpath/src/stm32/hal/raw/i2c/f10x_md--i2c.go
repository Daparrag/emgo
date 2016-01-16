// +build f10x_md

// Peripheral: I2C_Periph  Inter Integrated Circuit Interface.
// Instances:
//  I2C1  mmap.I2C1_BASE
//  I2C2  mmap.I2C2_BASE
// Registers:
//  0x00 16  CR1
//  0x04 16  CR2
//  0x08 16  OAR1
//  0x0C 16  OAR2
//  0x10 16  DR
//  0x14 16  SR1
//  0x18 16  SR2
//  0x1C 16  CCR
//  0x20 16  TRISE
// Import:
//  stm32/o/f10x_md/mmap
package i2c

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	PE        CR1_Bits = 0x01 << 0  //+ Peripheral Enable.
	SMBUS     CR1_Bits = 0x01 << 1  //+ SMBus Mode.
	SMBTYPE   CR1_Bits = 0x01 << 3  //+ SMBus Type.
	ENARP     CR1_Bits = 0x01 << 4  //+ ARP Enable.
	ENPEC     CR1_Bits = 0x01 << 5  //+ PEC Enable.
	ENGC      CR1_Bits = 0x01 << 6  //+ General Call Enable.
	NOSTRETCH CR1_Bits = 0x01 << 7  //+ Clock Stretching Disable (Slave mode).
	START     CR1_Bits = 0x01 << 8  //+ Start Generation.
	STOP      CR1_Bits = 0x01 << 9  //+ Stop Generation.
	ACK       CR1_Bits = 0x01 << 10 //+ Acknowledge Enable.
	POS       CR1_Bits = 0x01 << 11 //+ Acknowledge/PEC Position (for data reception).
	PEC       CR1_Bits = 0x01 << 12 //+ Packet Error Checking.
	ALERT     CR1_Bits = 0x01 << 13 //+ SMBus Alert.
	SWRST     CR1_Bits = 0x01 << 15 //+ Software Reset.
)

const (
	FREQ    CR2_Bits = 0x3F << 0  //+ FREQ[5:0] bits (Peripheral Clock Frequency).
	FREQ_0  CR2_Bits = 0x01 << 0  //  Bit 0.
	FREQ_1  CR2_Bits = 0x02 << 0  //  Bit 1.
	FREQ_2  CR2_Bits = 0x04 << 0  //  Bit 2.
	FREQ_3  CR2_Bits = 0x08 << 0  //  Bit 3.
	FREQ_4  CR2_Bits = 0x10 << 0  //  Bit 4.
	FREQ_5  CR2_Bits = 0x20 << 0  //  Bit 5.
	ITERREN CR2_Bits = 0x01 << 8  //+ Error Interrupt Enable.
	ITEVTEN CR2_Bits = 0x01 << 9  //+ Event Interrupt Enable.
	ITBUFEN CR2_Bits = 0x01 << 10 //+ Buffer Interrupt Enable.
	DMAEN   CR2_Bits = 0x01 << 11 //+ DMA Requests Enable.
	LAST    CR2_Bits = 0x01 << 12 //+ DMA Last Transfer.
)

const (
	ADD1_7  OAR1_Bits = 0x7F << 1  //+ Interface Address.
	ADD8_9  OAR1_Bits = 0x03 << 8  //+ Interface Address.
	ADD0    OAR1_Bits = 0x01 << 0  //+ Bit 0.
	ADD1    OAR1_Bits = 0x01 << 1  //  Bit 1.
	ADD2    OAR1_Bits = 0x02 << 1  //  Bit 2.
	ADD3    OAR1_Bits = 0x04 << 1  //  Bit 3.
	ADD4    OAR1_Bits = 0x08 << 1  //  Bit 4.
	ADD5    OAR1_Bits = 0x10 << 1  //  Bit 5.
	ADD6    OAR1_Bits = 0x20 << 1  //  Bit 6.
	ADD7    OAR1_Bits = 0x40 << 1  //  Bit 7.
	ADD8    OAR1_Bits = 0x01 << 8  //  Bit 8.
	ADD9    OAR1_Bits = 0x02 << 8  //  Bit 9.
	ADDMODE OAR1_Bits = 0x01 << 15 //+ Addressing Mode (Slave mode).
)

const (
	ENDUAL OAR2_Bits = 0x01 << 0 //+ Dual addressing mode enable.
	ADD2   OAR2_Bits = 0x7F << 1 //+ Interface address.
)

const ()

const (
	SB       SR1_Bits = 0x01 << 0  //+ Start Bit (Master mode).
	ADDR     SR1_Bits = 0x01 << 1  //+ Address sent (master mode)/matched (slave mode).
	BTF      SR1_Bits = 0x01 << 2  //+ Byte Transfer Finished.
	ADD10    SR1_Bits = 0x01 << 3  //+ 10-bit header sent (Master mode).
	STOPF    SR1_Bits = 0x01 << 4  //+ Stop detection (Slave mode).
	RXNE     SR1_Bits = 0x01 << 6  //+ Data Register not Empty (receivers).
	TXE      SR1_Bits = 0x01 << 7  //+ Data Register Empty (transmitters).
	BERR     SR1_Bits = 0x01 << 8  //+ Bus Error.
	ARLO     SR1_Bits = 0x01 << 9  //+ Arbitration Lost (master mode).
	AF       SR1_Bits = 0x01 << 10 //+ Acknowledge Failure.
	OVR      SR1_Bits = 0x01 << 11 //+ Overrun/Underrun.
	PECERR   SR1_Bits = 0x01 << 12 //+ PEC Error in reception.
	TIMEOUT  SR1_Bits = 0x01 << 14 //+ Timeout or Tlow Error.
	SMBALERT SR1_Bits = 0x01 << 15 //+ SMBus Alert.
)

const (
	MSL        SR2_Bits = 0x01 << 0 //+ Master/Slave.
	BUSY       SR2_Bits = 0x01 << 1 //+ Bus Busy.
	TRA        SR2_Bits = 0x01 << 2 //+ Transmitter/Receiver.
	GENCALL    SR2_Bits = 0x01 << 4 //+ General Call Address (Slave mode).
	SMBDEFAULT SR2_Bits = 0x01 << 5 //+ SMBus Device Default Address (Slave mode).
	SMBHOST    SR2_Bits = 0x01 << 6 //+ SMBus Host Header (Slave mode).
	DUALF      SR2_Bits = 0x01 << 7 //+ Dual Flag (Slave mode).
	PEC        SR2_Bits = 0xFF << 8 //+ Packet Error Checking Register.
)

const (
	CCR  CCR_Bits = 0xFFF << 0 //+ Clock Control Register in Fast/Standard mode (Master mode).
	DUTY CCR_Bits = 0x01 << 14 //+ Fast Mode Duty Cycle.
	FS   CCR_Bits = 0x01 << 15 //+ I2C Master Mode Selection.
)

const ()
