/**
* @file bitflag.go
* @brief Interfaces and structs for compact bit-efficient flag types.
* @author Anadian
* @copyright MITlicensetm(2019,Canosw)
*/
//package bitflag;
package main;

//Dependencies
import(
	//Internal
	//Standard
	"math"
	//External
);

//Constants
const(
	//Exported Constants
	/*POWER_OF_TWO []uint64 = []uint64{
		1,2,4,8,16,32,64,128,256,
		512,1024,2048,4096,8192,16384,32768,65536,
	}*/
	//Private Constants
);

//Types, structs, and methods
type BitFlag interface{
	GetSize() uint8
	HasFlagBit( flag_bit uint64 ) bool
	GetFlagBit( flag_bit uint64 ) uint8
	SetFlagBit( flag_bit uint64 )
	UnsetFlagBit( flag_bit uint64 )
	GetBit( bit_index uint8 ) uint8
	SetBit( bit_index uint8, bit_value uint8 )
}

///BitFlag (8-bit)
type BitFlag8 uint8;

func (bit_flag BitFlag8) GetSize() uint8{
	return uint8(8);
}

func (bit_flag BitFlag8) HasFlagBit( flag_bit uint64 ) bool{
	if( (uint8(bit_flag) & uint8(flag_bit)) == uint8(flag_bit) ){
		return true;
	} else{
		return false;
	}
}

func (bit_flag BitFlag8) GetFlagBit( flag_bit uint64 ) uint8{
	if( (uint8(bit_flag) & uint8(flag_bit)) == uint8(flag_bit) ){
		return uint8(1);
	} else{
		return uint8(0);
	}
}

func (bit_flag *BitFlag8) SetFlagBit( flag_bit uint64 ){
	*bit_flag |= BitFlag8(flag_bit);
}

func (bit_flag *BitFlag8) UnsetFlagBit( flag_bit uint64 ){
	*bit_flag -= BitFlag8(flag_bit);
}

func (bit_flag BitFlag8) GetBit( bit_index uint8 ) uint8{
	var flag_bit uint8 = uint8( math.Pow(2, float64(bit_index)) );
	if( (uint8(bit_flag) & flag_bit) == flag_bit ){
		return uint8(1);
	} else{
		return uint8(0);
	}
}

func (bit_flag *BitFlag8) SetBit( bit_index uint8, bit_value uint8 ){
	var flag_bit BitFlag8 = BitFlag8( math.Pow(2, float64(bit_index)) );
	if( bit_value == uint8(0) ){
		*bit_flag -= flag_bit;
	} else{
		*bit_flag |= flag_bit;
	}
}

///BitFlag (16-bit)
type BitFlag16 uint16;

func (bit_flag BitFlag16) GetSize() uint8{
	return uint8(16);
}

func (bit_flag BitFlag16) HasFlagBit( flag_bit uint64 ) bool{
	if( (uint16(bit_flag) & uint16(flag_bit)) == uint16(flag_bit) ){
		return true;
	} else{
		return false;
	}
}

func (bit_flag BitFlag16) GetFlagBit( flag_bit uint64 ) uint8{
	if( (uint16(bit_flag) & uint16(flag_bit)) == uint16(flag_bit) ){
		return uint8(1);
	} else{
		return uint8(0);
	}
}

func (bit_flag *BitFlag16) SetFlagBit( flag_bit uint64 ){
	*bit_flag |= BitFlag16(flag_bit);
}

func (bit_flag *BitFlag16) UnsetFlagBit( flag_bit uint64 ){
	*bit_flag -= BitFlag16(flag_bit);
}

func (bit_flag BitFlag16) GetBit( bit_index uint8 ) uint8{
	var flag_bit uint16 = uint16( math.Pow(2, float64(bit_index)) );
	if( (uint16(bit_flag) & flag_bit) == flag_bit ){
		return uint8(1);
	} else{
		return uint8(0);
	}
}

func (bit_flag *BitFlag16) SetBit( bit_index uint8, bit_value uint8 ){
	var flag_bit BitFlag16 = BitFlag16( math.Pow(2, float64(bit_index)) );
	if( bit_value == uint8(0) ){
		*bit_flag -= flag_bit;
	} else{
		*bit_flag |= flag_bit;
	}
}

///BitFlag (32-bit)
type BitFlag32 uint32;

func (bit_flag BitFlag32) GetSize() uint8{
	return uint8(32);
}

func (bit_flag BitFlag32) HasFlagBit( flag_bit uint64 ) bool{
	if( (uint32(bit_flag) & uint32(flag_bit)) == uint32(flag_bit) ){
		return true;
	} else{
		return false;
	}
}

func (bit_flag BitFlag32) GetFlagBit( flag_bit uint64 ) uint8{
	if( (uint32(bit_flag) & uint32(flag_bit)) == uint32(flag_bit) ){
		return uint8(1);
	} else{
		return uint8(0);
	}
}

func (bit_flag *BitFlag32) SetFlagBit( flag_bit uint64 ){
	*bit_flag |= BitFlag32(flag_bit);
}

func (bit_flag *BitFlag32) UnsetFlagBit( flag_bit uint64 ){
	*bit_flag -= BitFlag32(flag_bit);
}

func (bit_flag BitFlag32) GetBit( bit_index uint8 ) uint8{
	var flag_bit uint32 = uint32( math.Pow(2, float64(bit_index)) );
	if( (uint32(bit_flag) & flag_bit) == flag_bit ){
		return uint8(1);
	} else{
		return uint8(0);
	}
}

func (bit_flag *BitFlag32) SetBit( bit_index uint8, bit_value uint8 ){
	var flag_bit BitFlag32 = BitFlag32( math.Pow(2, float64(bit_index)) );
	if( bit_value == uint8(0) ){
		*bit_flag -= flag_bit;
	} else{
		*bit_flag |= flag_bit;
	}
}

///BitFlag (64-bit)
type BitFlag64 uint64;

func (bit_flag BitFlag64) GetSize() uint8{
	return uint8(64);
}

func (bit_flag BitFlag64) HasFlagBit( flag_bit uint64 ) bool{
	if( (uint64(bit_flag) & uint64(flag_bit)) == uint64(flag_bit) ){
		return true;
	} else{
		return false;
	}
}

func (bit_flag BitFlag64) GetFlagBit( flag_bit uint64 ) uint8{
	if( (uint64(bit_flag) & uint64(flag_bit)) == uint64(flag_bit) ){
		return uint8(1);
	} else{
		return uint8(0);
	}
}

func (bit_flag *BitFlag64) SetFlagBit( flag_bit uint64 ){
	*bit_flag |= BitFlag64(flag_bit);
}

func (bit_flag *BitFlag64) UnsetFlagBit( flag_bit uint64 ){
	*bit_flag -= BitFlag64(flag_bit);
}

func (bit_flag BitFlag64) GetBit( bit_index uint8 ) uint8{
	var flag_bit uint64 = uint64( math.Pow(2, float64(bit_index)) );
	if( (uint64(bit_flag) & flag_bit) == flag_bit ){
		return uint8(1);
	} else{
		return uint8(0);
	}
}

func (bit_flag *BitFlag64) SetBit( bit_index uint8, bit_value uint8 ){
	var flag_bit BitFlag64 = BitFlag64( math.Pow(2, float64(bit_index)) );
	if( bit_value == uint8(0) ){
		*bit_flag -= flag_bit;
	} else{
		*bit_flag |= flag_bit;
	}
}

//Global Variables
var(
	//Exported Variables
	//Private Variables
);

//Exported Functions

//Private Functions

