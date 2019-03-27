package main;
import (
	"fmt"
	"log"
	"image/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
);

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
	graphic Graphic_type
	position Vector_type
	velocity Vector_type
	acceleration Vector_type
	rest uint16
}

type enemy_ship_type struct{
	graphic Graphic_type
	position Vector_type
	hp uint8
}

//Global Variables
var player_ship player_ship_type;
var enemy_ship_graphic Graphic_type;
var enemy_ships_slice []enemy_ship_type;

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
func logic(){
	//Respond to input.
	var left_pressed, right_pressed bool = (Actions_map["left"].state > 0), (Actions_map["right"].state > 0);
	if( left_pressed && right_pressed ){
		if( Actions_map["left"].state <  Actions_map["right"].state ){
			player_ship.acceleration.x = -1;
		} else{
			player_ship.acceleration.x = 1;
		}
	} else if( left_pressed ){
		player_ship.acceleration.x = -1;
	} else if( right_pressed ){
		player_ship.acceleration.x = 1;
	} else{
		player_ship.acceleration.x = 0;
	}
	//Apply physics.
	player_ship.velocity.x += player_ship.acceleration.x;
	//Terminal velocity and Friction
	if( player_ship.velocity.x > 4 ){
		player_ship.velocity.x = 4;
	} else if( player_ship.velocity.x < -4){
		player_ship.velocity.x = -4;
	} else if( player_ship.velocity.x > 0 ){
		player_ship.velocity.x -= 0.5;
	} else if( player_ship.velocity.x < 0 ){
		player_ship.velocity.x += 0.5;
	}
	player_ship.position.x += player_ship.velocity.x;
	//Check bounds.
	if( (player_ship.position.x + 16) < 0 ){
		player_ship.position.x = -16;
	} else if( (player_ship.position.x + 16) > 640 ){
		player_ship.position.x  = 640 - 16;
	}
	player_ship.graphic.op = &ebiten.DrawImageOptions{};
	player_ship.graphic.op.GeoM.Translate(player_ship.position.x, player_ship.position.y);
	for i := 0; i < len(enemy_ships_slice); i++{
		enemy_ships_slice[i].graphic.op = &ebiten.DrawImageOptions{};
		enemy_ships_slice[i].graphic.op.GeoM.Translate(enemy_ships_slice[i].position.x, enemy_ships_slice[i].position.y);
	}
}
func main(){
	//Init
	player_ship = player_ship_type{
		NewGraphic_Rect_Colour(32,16,0,255,0,255),
		Vector_type{float64((640-32)/2),float64(480-32),},
		Vector_type{0,0,},
		Vector_type{0,0,},
		5,
	};
	enemy_ship_graphic = NewGraphic_Rect_Colour(20,20,255,126,0,255);
	var enemy_ship enemy_ship_type;
	for i := 0; i < 100; i++{
		enemy_ship = enemy_ship_type{
			enemy_ship_graphic,
			Vector_type{float64((i % 16)*32),float64(((i/20)*32)+16),},
			1,
		};
		enemy_ships_slice = append(enemy_ships_slice,enemy_ship);
	}
	var ebiten_error error = ebiten.Run( MainLoop, 640, 480, 1, "Hello, world?" );
	fmt.Printf("ebiten_error: %v\n", ebiten_error);
}
func MainLoop( screen *ebiten.Image ) error{
	//Input
	input(); //Process input
	logic(); //Process logic

	log.Printf("screen: %v", screen);
	if( !ebiten.IsDrawingSkipped() ){
		ebitenutil.DebugPrint( screen, "Hello, World!" );
		//Draw player ship
		screen.DrawImage(player_ship.graphic.image, player_ship.graphic.op);
		//Enemy ships
		for i := 0; i < len(enemy_ships_slice); i++{
			screen.DrawImage(enemy_ships_slice[i].graphic.image, enemy_ships_slice[i].graphic.op);
		}
	}
	return nil;
}
