/**
* @file bitflag_arraylist_test.go
* @brief Test storing and managing BitFlags in a `gods` arraylist.
* @author Anadian
* @copyright MITlicensetm(2019,Canosw)
*/
package main;

//Dependencies
import(
	//Internal
	//Standard
	"testing"
	//External
	"github.com/emirpasic/gods/lists/arraylist"
);

//Constants
const(
	//Exported Constants
	//Private Constants
);

//Types, structs, and methods

//Global Variables
var(
	//Exported Variables
	//Private Variables
);

//Exported Functions

//Private Functions
func TestBitFlags(t *testing.T){
	var bitflag8 BitFlag8 = 16;
	var bitflag16 BitFlag16 = 256;
	var bitflag32 BitFlag32 = 65536;
	var bitflag64 BitFlag64 = 4294967296;
	var bitflag8_p *BitFlag8 = &(bitflag8);
	var bitflag16_p *BitFlag16 = &(bitflag16);
	var bitflag32_p *BitFlag32 = &(bitflag32);
	var bitflag64_p *BitFlag64 = &(bitflag64);
	//BitFlag8 test
	_ = bitflag8;
	t.Logf("%T: %d @ %p (%d)", bitflag8, bitflag8, bitflag8_p, *bitflag8_p);
	t.Logf("GetSize: %d (%d)", bitflag8.GetSize(), bitflag8_p.GetSize());
	t.Logf("HasFlagBit(8): %t (%t) HasFlagBit(16): %t (%t)", bitflag8.HasFlagBit(uint64(8)), bitflag8_p.HasFlagBit(uint64(8)), bitflag8.HasFlagBit(uint64(16)), bitflag8_p.HasFlagBit(uint64(16)));
	t.Logf("GetFlagBit(8): %d (%d) GetFlagBit(16): %d (%d)", bitflag8.GetFlagBit(uint64(8)), bitflag8_p.GetFlagBit(uint64(8)), bitflag8.GetFlagBit(uint64(16)), bitflag8_p.GetFlagBit(uint64(16)));
	t.Logf("SetFlagBit(8)");
	bitflag8.SetFlagBit(uint64(8));
	t.Logf("GetFlagBit(8): %d (%d) GetFlagBit(16): %d (%d)", bitflag8.GetFlagBit(uint64(8)), bitflag8_p.GetFlagBit(uint64(8)), bitflag8.GetFlagBit(uint64(16)), bitflag8_p.GetFlagBit(uint64(16)));
	t.Logf("UnsetFlagBit(16)");
	bitflag8.UnsetFlagBit(uint64(16));
	t.Logf("GetFlagBit(8): %d (%d) GetFlagBit(16): %d (%d)", bitflag8.GetFlagBit(uint64(8)), bitflag8_p.GetFlagBit(uint64(8)), bitflag8.GetFlagBit(uint64(16)), bitflag8_p.GetFlagBit(uint64(16)));
	t.Logf("GetBit(2): %d (%d) GetBit(5): %d (%d)", bitflag8.GetBit(uint8(2)), bitflag8_p.GetBit(uint8(2)), bitflag8.GetBit(uint8(5)), bitflag8_p.GetBit(uint8(5)));
	t.Logf("SetBit(5, 1)");
	bitflag8.SetBit(uint8(5), uint8(1));
	t.Logf("GetBit(2): %d (%d) GetBit(5): %d (%d)", bitflag8.GetBit(uint8(2)), bitflag8_p.GetBit(uint8(2)), bitflag8.GetBit(uint8(5)), bitflag8_p.GetBit(uint8(5)));
	//BitFlag16 test
	_ = bitflag16;
	t.Logf("%T: %d @ %p (%d)", bitflag16, bitflag16, bitflag16_p, *bitflag16_p);
	t.Logf("GetSize: %d (%d)", bitflag16.GetSize(), bitflag16_p.GetSize());
	t.Logf("HasFlagBit(8): %t (%t) HasFlagBit(16): %t (%t)", bitflag16.HasFlagBit(uint64(8)), bitflag16_p.HasFlagBit(uint64(8)), bitflag16.HasFlagBit(uint64(256)), bitflag16_p.HasFlagBit(uint64(256)));
	t.Logf("GetFlagBit(8): %d (%d) GetFlagBit(16): %d (%d)", bitflag16.GetFlagBit(uint64(8)), bitflag16_p.GetFlagBit(uint64(8)), bitflag16.GetFlagBit(uint64(256)), bitflag16_p.GetFlagBit(uint64(256)));
	t.Logf("SetFlagBit(8)");
	bitflag16.SetFlagBit(uint64(8));
	t.Logf("GetFlagBit(8): %d (%d) GetFlagBit(16): %d (%d)", bitflag16.GetFlagBit(uint64(8)), bitflag16_p.GetFlagBit(uint64(8)), bitflag16.GetFlagBit(uint64(256)), bitflag16_p.GetFlagBit(uint64(256)));
	t.Logf("UnsetFlagBit(16)");
	bitflag16.UnsetFlagBit(uint64(256));
	t.Logf("GetFlagBit(8): %d (%d) GetFlagBit(16): %d (%d)", bitflag16.GetFlagBit(uint64(8)), bitflag16_p.GetFlagBit(uint64(8)), bitflag16.GetFlagBit(uint64(256)), bitflag16_p.GetFlagBit(uint64(256)));
	t.Logf("GetBit(2): %d (%d) GetBit(5): %d (%d)", bitflag16.GetBit(uint8(2)), bitflag16_p.GetBit(uint8(2)), bitflag16.GetBit(uint8(5)), bitflag16_p.GetBit(uint8(5)));
	t.Logf("SetBit(5, 1)");
	bitflag16.SetBit(uint8(5), uint8(1));
	t.Logf("GetBit(2): %d (%d) GetBit(5): %d (%d)", bitflag16.GetBit(uint8(2)), bitflag16_p.GetBit(uint8(2)), bitflag16.GetBit(uint8(5)), bitflag16_p.GetBit(uint8(5)));
	//BitFlag32 test
	_ = bitflag32;
	t.Logf("%T: %d @ %p (%d)", bitflag32, bitflag32, bitflag32_p, *bitflag32_p);
	t.Logf("GetSize: %d (%d)", bitflag32.GetSize(), bitflag32_p.GetSize());
	t.Logf("HasFlagBit(8): %t (%t) HasFlagBit(16): %t (%t)", bitflag32.HasFlagBit(uint64(8)), bitflag32_p.HasFlagBit(uint64(8)), bitflag32.HasFlagBit(uint64(65536)), bitflag32_p.HasFlagBit(uint64(65536)));
	t.Logf("GetFlagBit(8): %d (%d) GetFlagBit(16): %d (%d)", bitflag32.GetFlagBit(uint64(8)), bitflag32_p.GetFlagBit(uint64(8)), bitflag32.GetFlagBit(uint64(65536)), bitflag32_p.GetFlagBit(uint64(65536)));
	t.Logf("SetFlagBit(8)");
	bitflag32.SetFlagBit(uint64(8));
	t.Logf("GetFlagBit(8): %d (%d) GetFlagBit(16): %d (%d)", bitflag32.GetFlagBit(uint64(8)), bitflag32_p.GetFlagBit(uint64(8)), bitflag32.GetFlagBit(uint64(65536)), bitflag32_p.GetFlagBit(uint64(65536)));
	t.Logf("UnsetFlagBit(16)");
	bitflag32.UnsetFlagBit(uint64(65536));
	t.Logf("GetFlagBit(8): %d (%d) GetFlagBit(16): %d (%d)", bitflag32.GetFlagBit(uint64(8)), bitflag32_p.GetFlagBit(uint64(8)), bitflag32.GetFlagBit(uint64(65536)), bitflag32_p.GetFlagBit(uint64(65536)));
	t.Logf("GetBit(2): %d (%d) GetBit(5): %d (%d)", bitflag32.GetBit(uint8(2)), bitflag32_p.GetBit(uint8(2)), bitflag32.GetBit(uint8(5)), bitflag32_p.GetBit(uint8(5)));
	t.Logf("SetBit(5, 1)");
	bitflag32.SetBit(uint8(5), uint8(1));
	t.Logf("GetBit(2): %d (%d) GetBit(5): %d (%d)", bitflag32.GetBit(uint8(2)), bitflag32_p.GetBit(uint8(2)), bitflag32.GetBit(uint8(5)), bitflag32_p.GetBit(uint8(5)));
	//BitFlag64 test
	_ = bitflag64;
	t.Logf("%T: %d @ %p (%d)", bitflag64, bitflag64, bitflag64_p, *bitflag64_p);
	t.Logf("GetSize: %d (%d)", bitflag64.GetSize(), bitflag64_p.GetSize());
	t.Logf("HasFlagBit(8): %t (%t) HasFlagBit(16): %t (%t)", bitflag64.HasFlagBit(uint64(8)), bitflag64_p.HasFlagBit(uint64(8)), bitflag64.HasFlagBit(uint64(4294967296)), bitflag64_p.HasFlagBit(uint64(4294967296)));
	t.Logf("GetFlagBit(8): %d (%d) GetFlagBit(16): %d (%d)", bitflag64.GetFlagBit(uint64(8)), bitflag64_p.GetFlagBit(uint64(8)), bitflag64.GetFlagBit(uint64(4294967296)), bitflag64_p.GetFlagBit(uint64(4294967296)));
	t.Logf("SetFlagBit(8)");
	bitflag64.SetFlagBit(uint64(8));
	t.Logf("GetFlagBit(8): %d (%d) GetFlagBit(16): %d (%d)", bitflag64.GetFlagBit(uint64(8)), bitflag64_p.GetFlagBit(uint64(8)), bitflag64.GetFlagBit(uint64(4294967296)), bitflag64_p.GetFlagBit(uint64(4294967296)));
	t.Logf("UnsetFlagBit(16)");
	bitflag64.UnsetFlagBit(uint64(4294967296));
	t.Logf("GetFlagBit(8): %d (%d) GetFlagBit(16): %d (%d)", bitflag64.GetFlagBit(uint64(8)), bitflag64_p.GetFlagBit(uint64(8)), bitflag64.GetFlagBit(uint64(4294967296)), bitflag64_p.GetFlagBit(uint64(4294967296)));
	t.Logf("GetBit(2): %d (%d) GetBit(5): %d (%d)", bitflag64.GetBit(uint8(2)), bitflag64_p.GetBit(uint8(2)), bitflag64.GetBit(uint8(5)), bitflag64_p.GetBit(uint8(5)));
	t.Logf("SetBit(5, 1)");
	bitflag64.SetBit(uint8(5), uint8(1));
	t.Logf("GetBit(2): %d (%d) GetBit(5): %d (%d)", bitflag64.GetBit(uint8(2)), bitflag64_p.GetBit(uint8(2)), bitflag64.GetBit(uint8(5)), bitflag64_p.GetBit(uint8(5)));
	//arraylist test
	array_list := arraylist.New();
	array_list.Add( bitflag8 );
	t.Logf("%T: %v Size: %d", array_list, array_list, array_list.Size());
	array_list.Add( bitflag16, bitflag32, bitflag64 );
	t.Logf("%v (size: %d)", array_list, array_list.Size());
	array_list.Insert(2, bitflag32_p);
	t.Logf("%v (size: %d)", array_list, array_list.Size());
	bitflag32.SetBit(uint8(7), uint8(1));
	t.Logf("%v (size: %d)", array_list, array_list.Size());
	array_list.Swap(2, 3);
	t.Logf("%v (size: %d)", array_list, array_list.Size());
	array_list.Remove(2);
	t.Logf("%v (size: %d)", array_list, array_list.Size());
	array_list.Remove(2);
	t.Logf("%v (size: %d)", array_list, array_list.Size());
	array_list.Insert(2, *bitflag32_p);
	t.Logf("%v (size: %d)", array_list, array_list.Size());
	array_list.Clear();
	t.Logf("%v (size: %d)", array_list, array_list.Size());

	return;
}
