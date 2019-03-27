package main;
import (
	"fmt"
	"log"
	"image/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
);
/*var player_ship_image *ebiten.Image;
var player_ship_x, player_ship_y float64;*/
type Vector_type struct{
	x float64
	y float64
}
type Action_type struct{
	mapping ebiten.Key
	state int64
}
var Actions_map = map[string]Action_type{
	"pause": {ebiten.KeyEscape,0,},
	"left": {ebiten.KeyLeft,0,},
	"right": {ebiten.KeyRight,0,},
	"fire": {ebiten.KeySpace,0,},
};
type Graphic_type struct{
	image *ebiten.Image
	op *ebiten.DrawImageOptions
}
func (graphic *Graphic_type) Set_Rect_Colour( width uint32, height uint32, red uint8, green uint8, blue uint8, alpha uint8 ) *Graphic_type{
	graphic.image, _ = ebiten.NewImage( int(width), int(height), ebiten.FilterDefault );
	graphic.image.Fill(color.RGBA{red,green,blue,alpha,});
	graphic.op = &ebiten.DrawImageOptions{};
	return graphic;
}
func NewGraphic_Rect_Colour( width uint32, height uint32, red uint8, green uint8, blue uint8, alpha uint8 ) Graphic_type{
	var new_graphic Graphic_type;
	new_graphic.image, _ = ebiten.NewImage( int(width), int(height), ebiten.FilterDefault );
	new_graphic.image.Fill(color.RGBA{red,green,blue,alpha,});
	new_graphic.op = &ebiten.DrawImageOptions{};
	return new_graphic;
}
type player_ship_type struct{
	image *ebiten.Image
	x float64
	y float64
}
var player_ship player_ship_type;
var generic_sprite Graphic_type;
func input(){
	for name := range Actions_map{
		if( ebiten.IsKeyPressed(Actions_map[name].mapping) ){
			if( Actions_map[name].state >= 1 ){
				Actions_map[name] = Action_type{Actions_map[name].mapping,(Actions_map[name].state + 1),};
			} else{
				Actions_map[name] = Action_type{Actions_map[name].mapping,1,};
			}
		} else{
			if( Actions_map[name].state <= -1 ){
				Actions_map[name] = Action_type{Actions_map[name].mapping,(Actions_map[name].state - 1),};
			} else{
				Actions_map[name] = Action_type{Actions_map[name].mapping,-1,};
			}
		}
	}
}
func main(){
	fmt.Println("Will this work?");
	//##Variable declarations
	//###Multi-varibale declaration
	//var single_byte Uint8 = 26, full_integer Int64 = 6555000, some_float Float32 = 3.14159265358979323; Doesn't work.
	//var single_byte = 26, full_integer = 6555000, some_float = 3.14159265358979323; Also doesn't work.
	//var single_byte uint8, full_integer int64, some_float float32 = 65, 6555000, 3.1415926535899323; //Doesn't works
	//var single_byte = 26, full_integer = 6555000, some_float = 3.14159265358979323 Uint8, Int64, Float32; Doesn't work.
	//var single_byte uint8 = 26, full_integer int64 = 6555000, some_float float32 = 3.14159265358979323; //Doesn't work.
	var single_byte, full_integer, some_float = 65, 6555000, 3.1415926535899323; //Works
	//###Single explicit-type declaration
	//var another_byte = 26 Uint8; Doesn't work.
	//var another_byte = 26 uint8; //Doesn't work.
	var another_byte uint8 = 26; //Works
	//###Constants are declared as above.
	const this_is_a_constant uint8 = 47;
	fmt.Printf("Apparently, this should print a byte (%d), interger (%d), and float (%f).\n", single_byte, full_integer, some_float);
	fmt.Printf("Will this notation work? %d now with a constant %d\n", another_byte, this_is_a_constant);
	//##Operators
	//###Arithmetic
	// additon (x + y), subtraction (x - y), multiplication (x * y), division (x / y), modulo (x % y), postfix increment (x++), postfix decrement (x--)
	//###Relational
	// ==, !=, >, <, >=, <=
	//###Logical
	// and (&&), or (||), not (!)
	//###Bitwise
	// and (&), or (|), xor (^), shift left (<<), shift right (>>)
	//###Assignment
	// +=, -=, *=, /=, %=, <<=, >>=, &=, ^=, |=
	//Miscellaneous
	// address (&), pointer (*)
	/*if( condition ){

	} else if( other_condition ){

	} else{

	}*/
	/*player_ship_x = 0;
	player_ship_y = 0;
	player_ship_image, _ = ebiten.NewImage( 32, 16, ebiten.FilterDefault );
	log.Printf("player_ship_image: %v", player_ship_image);
	player_ship_image.Fill(color.RGBA{0,255,0,255,});
	log.Printf("player_ship_image: %v", player_ship_image);*/
	temp_image, _ := ebiten.NewImage( 32, 16, ebiten.FilterDefault );
	player_ship = player_ship_type{
		temp_image,
		0,
		0,
	};
	player_ship.image.Fill(color.RGBA{0,255,0,255});
	generic_sprite.NewGraphic_Rect_Colour(32, 48, 255, 0, 255, 200);
	var another_graphic = NewGraphic_Rect_Colour(26, 12, 0, 0, 255, 250);
	fmt.Printf("another_graphic: %v\n", another_graphic);
	var ebiten_error error = ebiten.Run( MainLoop, 640, 480, 1, "Hello, world?" );
	fmt.Printf("ebiten_error: %v\n", ebiten_error);
}
func MainLoop( screen *ebiten.Image ) error{
	//Input
	if( ebiten.IsKeyPressed( ebiten.KeyRight ) ){
		player_ship.x += 1;
	} else if( ebiten.IsKeyPressed( ebiten.KeyLeft ) ){
		player_ship.x -= 1;
	}
	if( ebiten.IsKeyPressed( ebiten.KeyUp ) ){
		player_ship.y -= 1;
	} else if( ebiten.IsKeyPressed( ebiten.KeyDown ) ){
		player_ship.y += 1;
	}
	input(); //Process input
	logic(); //Process logic

	log.Printf("screen: %v", screen);
	op := &ebiten.DrawImageOptions{};
	log.Printf("op: %v", op);
	if( !ebiten.IsDrawingSkipped() ){
		ebitenutil.DebugPrint( screen, "Hello, World!" );
		op.GeoM.Translate( player_ship.x, player_ship.y );
		log.Printf("op: %v", op);
		screen.DrawImage(player_ship.image, op);
		screen.DrawImage(generic_sprite.image, generic_sprite.op);
	}
	return nil;
}
