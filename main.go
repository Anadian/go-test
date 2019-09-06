package main;
import (
	//Internal
	//Standard
	"fmt"
	"log"
	"image/color"
	//External
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
);

func NewImage_Rect_RGBA( width uint32, height uint32, red uint8, green uint8, blue uint8, alpha uint8) *ebiten.Image{
	var new_image *ebiten.Image;
	new_image, _ = ebiten.NewImage( int(width), int(height), ebiten.FilterDefault );
	new_image.Fill(color.RGBA{red,green,blue,alpha,});
	return new_image;
}

type Action_type struct{
	mapping ebiten.Key
	state int64
}

//graphic.go
type Graphic_type struct{
	image_index uint8
	op *ebiten.DrawImageOptions
}
func NewGraphic( image_index uint8 ) Graphic_type{
	var new_graphic Graphic_type;
	new_graphic.image_index = image_index;
	new_graphic.op = &ebiten.DrawImageOptions{};
	return new_graphic;
}

type Meter_type struct{
	current uint64
	maximum uint64
}

type ActorMeters_type struct{
	hp Meter_type
	sp Meter_type
//	ep Meter_type
}

type ActorStats_type struct{
	strength uint8 //Damage per shot
	constitution uint8 //Health, hits per life
	intelligence uint8 //Number of shots onscreen at once, faster SP recharge
	agility uint8 //Top Speed
	dexterity uint8 //Handling, turning ratio
}

const(
	//Box_type.alignment
	ALIGNMENT_UNIVERSAL uint8 = 0
	ALIGNMENT_PLAYER uint8 = 1
	ALIGNMENT_ENEMY uint8 = 2
	//Box_type.boxtype
	BOXTYPE_COLLISION uint8 = 1
	BOXTYPE_HIT uint8 = 2
	BOXTYPE_HURT uint8 = 4
	BOXTYPE_WIND uint8 = 8
	BOXTYPE_EVENT uint8 = 16
);
type Box_type struct{
	alignment uint8
	boxtype uint8
	magnitude Vector_type
	rect Rectangle_type
}

type PlayerShip_type struct{
	enabled bool
	graphic Graphic_type
	box Box_type
	centre Vector_type
	velocity Vector_type
	acceleration Vector_type
	rest uint16
	meters ActorMeters_type
	stats ActorStats_type
}

type EnemyShip_type struct{
	enabled bool
	graphic Graphic_type
	box Box_type
	centre Vector_type
	hp uint8
}

type Projectile_type struct{
	enabled bool
	graphic Graphic_type
	box Box_type
	centre Vector_type
	velocity Vector_type
}

/**
* @fn Projectile_New
* @brief Create a new projectile.
* @param alignment uint8 [in] The alignment of the projectile.
* @param damage uint32 [in] The damage dealt by the projectile.
* @param penetration uint32 [in] The number of enemies the projectile can pass through before disappearing.
* @param centre Vector_type [in] The centre vector of the projectile.
* @param velocity Vector_type [in] The velocity vector of the projectile.
* @return Projectile_type
* @retval nil Failure
* @retval <Projectile> Success
*/
func  Projectile_New( alignment uint8, damage uint32, penetration uint32, centre Vector_type, velocity Vector_type ) Projectile_type{
	/* Variables */
	var _return Projectile_type = Projectile_type{
		true,
		NewGraphic(GRAPHIC_INDEX_PROJECTILE),
		Box_type{
			alignment,
			(BOXTYPE_HIT),
			Vector_type{ float64(damage), float64(penetration), },
			Rectangle_type{
				Vector_type{ 0, 0, },
				Vector_type{ 16, 20, },
			},
		},
		centre,
		velocity,
	};
	/* Parametres */
	/* Function */
	/* Return */
	return _return;
}


type SpaceInvaders_Scene_type struct{
	rect Rectangle_type
	images []*ebiten.Image
	player_ship PlayerShip_type
	enemy_ships []EnemyShip_type
	projectiles []Projectile_type
}

func (scene SpaceInvaders_Scene_type) BoundsCheck( x float64, y float64, width float64, height float64 ) uint8{
	var _return uint8 = 0;
	if( uint32(x + width) < 0 ){
		_return += 1;
	} else if( uint32(x) > uint32(scene.rect.magnitude.x) ){
		_return += 2;
	}
	if( uint32(y + height) < 0 ){
		_return += 4;
	} else if( uint32(y) > uint32(scene.rect.magnitude.y) ){
		_return += 8;
	}
	return _return;
}

func (scene SpaceInvaders_Scene_type) BoundsCheck_Rectangle( rect Rectangle_type ) uint8{
	var _return uint8 = 0;
	if( uint32(rect.origin.x + rect.magnitude.x) < 0 ){
		_return += 1;
	} else if( uint32(rect.origin.x) > uint32(scene.rect.magnitude.x) ){
		_return += 2;
	}
	if( uint32(rect.origin.y + rect.magnitude.y) < 0 ){
		_return += 4;
	} else if( uint32(rect.origin.y) > uint32(scene.rect.magnitude.y) ){
		_return += 8;
	}
	return _return;
}

func (scene *SpaceInvaders_Scene_type) AddEnemy( x float64, y float64 ){
	var new_enemy_ship EnemyShip_type;
	new_enemy_ship = EnemyShip_type{
		true,
		NewGraphic( 1 ),
		Box_type{
			1,
			( BOXTYPE_HURT ),
			Vector_type{0,0,},
			Rectangle_type{
				Vector_type{
					x,
					y,
				},
				Vector_type{
					20,
					20,
				},
			},
		},
		Vector_type{
			x,
			y,
		},
		1,
	};
	scene.enemy_ships = append(scene.enemy_ships, new_enemy_ship);
}

func (scene *SpaceInvaders_Scene_type) AddProjectile( x float64, y float64, alignment uint8 ){
	var new_projectile Projectile_type;
	if( alignment == 0 ){
		new_projectile = Projectile_type{
			true,
			NewGraphic(2),
			Box_type{
				alignment,
				6,
				Vector_type{
					1,
					0,
				},
				Rectangle_type{
					Vector_type{
						0,
						0,
					},
					Vector_type{
						16,
						20,
					},
				},
			},
			Vector_type{
				x,
				y,
			},
			Vector_type{
				0,
				-3,
			},
		};
		new_projectile.box.rect.SetOriginFromCentre( new_projectile.centre );
		var projectile_added bool = false;
		for i := 0; i < len(scene.projectiles); i++{
			if( scene.projectiles[i].enabled == false ){
				scene.projectiles[i] = new_projectile;
				projectile_added = true;
				break;
			}
		}
		if( projectile_added == false ){
			scene.projectiles = append(scene.projectiles, new_projectile);
		}
	}
}
/**
* @fn AddProjectile
* @brief A method to add a new projectile to the scene.
* @struct scene *Scene_type
* @param alignment uint8 [in] 
* @param centre Vector_type [in] 
* @param velocity Vector_type [in] 
* @return uint8
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/
func (scene *Scene_type) AddProjectile( alignment uint8, centre Vector_type, velocity Vector_type ) uint8{
	var _return uint8 = nil;
	/* Variables */
	
	/* Parametres */

	/* Function */
	/* Return */
	return _return;
}


//Global Variables
var Actions_map = map[string]Action_type{
	"pause": {ebiten.KeyEscape,0,},
	"left": {ebiten.KeyLeft,0,},
	"right": {ebiten.KeyRight,0,},
	"fire": {ebiten.KeySpace,0,},
};
var SpaceInvaders_Images = []*ebiten.Image{
	NewImage_Rect_RGBA( 32, 16, 0, 255, 0, 255 ), //Player ship
	NewImage_Rect_RGBA( 20, 20, 255, 128, 0, 255 ), //Enemy ship
	NewImage_Rect_RGBA( 16, 20, 0, 0, 255, 255), //Projectile
};
var Scene SpaceInvaders_Scene_type;

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
	///Respond to input.
	var left_pressed, right_pressed bool = (Actions_map["left"].state > 0), (Actions_map["right"].state > 0);
	var intersection Intersection_type;
	if( left_pressed && right_pressed ){
		if( Actions_map["left"].state <  Actions_map["right"].state ){
			Scene.player_ship.acceleration.x = -1;
		} else{
			Scene.player_ship.acceleration.x = 1;
		}
	} else if( left_pressed ){
		Scene.player_ship.acceleration.x = -1;
	} else if( right_pressed ){
		Scene.player_ship.acceleration.x = 1;
	} else{
		Scene.player_ship.acceleration.x = 0;
	}
	///Apply physics.
	////Quick turn consideration
	if( (Scene.player_ship.velocity.x < -2 && Scene.player_ship.acceleration.x == 1) || (Scene.player_ship.velocity.x > 2 && Scene.player_ship.acceleration.x == -1) ){
		Scene.player_ship.velocity.x = -(Scene.player_ship.velocity.x);
	} else{
		Scene.player_ship.velocity.x += Scene.player_ship.acceleration.x;
	}
	////Terminal velocity and Friction
	if( Scene.player_ship.velocity.x > 4 ){
		Scene.player_ship.velocity.x = 4;
	} else if( Scene.player_ship.velocity.x < -4){
		Scene.player_ship.velocity.x = -4;
	} else if( Scene.player_ship.velocity.x > 0 ){
		Scene.player_ship.velocity.x -= 0.5;
	} else if( Scene.player_ship.velocity.x < 0 ){
		Scene.player_ship.velocity.x += 0.5;
	}
	Scene.player_ship.centre.x += Scene.player_ship.velocity.x;
	////Check bounds.
//	if( (Scene.player_ship.centre.x + 16) < 0 ){
//		Scene.player_ship.centre.x = -16;
//	} else if( (Scene.player_ship.centre.x + 16) > 640 ){
//		Scene.player_ship.centre.x  = 640 - 16;
//	}
//with .IntersectRectangle
//	var intersection Intersection_type = Scene.player_ship.box.rect.IntersectRectangle( Scene.rect );
//	if( (intersection.intersection & INTERSECTION_ON) != INTERSECTION_ON ){
//		if( intersection.centre_from_centre.x < 0 ){
//			Scene.player_ship.centre.x = 0;
//		} else if( intersection.centre_from_centre.x > float64(Scene.width) ){
//			Scene.player_ship.centre.x = float64(Scene.width);
//		}
//	}
	if( Scene.player_ship.centre.InRectangle( Scene.rect ) == false ){
		if( Scene.player_ship.centre.x < Scene.rect.origin.x ){
			Scene.player_ship.centre.x = Scene.rect.origin.x;
		} else if( Scene.player_ship.centre.x > (Scene.rect.origin.x + Scene.rect.magnitude.x) ){
			Scene.player_ship.centre.x = (Scene.rect.origin.x + Scene.rect.magnitude.x);
		}
	}
	Scene.player_ship.box.rect.SetOriginFromCentre( Scene.player_ship.centre );
	///Move enemy ships
	////Move bullets
	for i := 0; i < len(Scene.projectiles); i++{
		if( Scene.projectiles[i].enabled == true ){
			Scene.projectiles[i].centre.x += Scene.projectiles[i].velocity.x;
			Scene.projectiles[i].centre.y += Scene.projectiles[i].velocity.y;
			/////Update box
			Scene.projectiles[i].box.rect.SetOriginFromCentre( Scene.projectiles[i].centre );
			/////Check bounds
//		intersection = Scene.projectiles[i].box.rect.IntersectRectangle( Scene.rect );
//		if( (intersection.intersection & INTERSECTION_ON) != INTERSECTION_ON ){
			if( Scene.projectiles[i].centre.InRectangle( Scene.rect ) == false ){
				Scene.projectiles[i].enabled = false;
			} else{
				for j := 0; j < len(Scene.enemy_ships); j++{
					if( Scene.enemy_ships[j].enabled == true ){
						intersection = Scene.projectiles[i].box.rect.IntersectRectangle( Scene.enemy_ships[j].box.rect );
						if( (intersection.intersection & INTERSECTION_INTERSECT) == INTERSECTION_INTERSECT ){
							Scene.enemy_ships[j].enabled = false;
						}
					}
				}
			}
//			Scene.projectiles[i].box.rect.origin.x = (Scene.projectiles[i].centre.x - (Scene.projectiles[i].box.rect.magnitude.x / 2));
//			Scene.projectiles[i].box.rect.origin.y = (Scene.projectiles[i].centre.y - (Scene.projectiles[i].box.rect.magnitude.y / 2));
//			////Check bounds.
//			var bound_check uint8 = Scene.BoundsCheck_Rectangle( Scene.projectiles[i].box.rect );
//			if( bound_check != 0 ){
//				if( (bound_check & 1) > 0 ){
//					Scene.projectiles[i].centre.x = 0;
//				} else if( (bound_check & 2) > 0 ){
//					Scene.projectiles[i].centre.x = float64(Scene.width);
//				}
//				if( (bound_check & 4) > 0 ){
//					Scene.projectiles[i].centre.y = 0;
//				} else if( (bound_check & 8) > 0 ){
//					Scene.projectiles[i].centre.y = float64(Scene.height);
//				}
//				Scene.projectiles[i].box.rect.SetOriginFromCentre(Scene.projectiles[i].centre);
//				Scene.projectiles[i].enabled = false;
//			}
		}
	}
	///Player Fire
	if( Scene.player_ship.meters.sp.current < Scene.player_ship.meters.sp.maximum ){
		Scene.player_ship.meters.sp.current++;
	}
	if( Actions_map["fire"].state == 1 ){
		if( Scene.player_ship.meters.sp.current > 30 ){
			var projectile Projectile_type = Projectile_New( alignment, damage, penetration, centre, velocity , centre, velocity )
			Scene.AddProjectile( Scene.player_ship.centre.x , (Scene.player_ship.centre.y - 32), 0 );
			Scene.player_ship.meters.sp.current -= 30;
		}
	}
	///Drawing options
	////Update player_ship drawing options
	Scene.player_ship.graphic.op = &ebiten.DrawImageOptions{};
	Scene.player_ship.graphic.op.GeoM.Translate(Scene.player_ship.box.rect.origin.x, Scene.player_ship.box.rect.origin.y);
	////Enemy ships drawing options
	for i := 0; i < len(Scene.enemy_ships); i++{
		Scene.enemy_ships[i].graphic.op = &ebiten.DrawImageOptions{};
		Scene.enemy_ships[i].graphic.op.GeoM.Translate(Scene.enemy_ships[i].box.rect.origin.x, Scene.enemy_ships[i].box.rect.origin.y);
	}
	////Projectile drawing options
	for i:= 0; i < len(Scene.projectiles); i++{
		Scene.projectiles[i].graphic.op = &ebiten.DrawImageOptions{};
		Scene.projectiles[i].graphic.op.GeoM.Translate( Scene.projectiles[i].centre.x - (Scene.projectiles[i].box.rect.magnitude.x / 2), Scene.projectiles[i].centre.y - (Scene.projectiles[i].box.rect.magnitude.y / 2) );
	}
}
func main(){
	//Init
	Scene.rect = Rectangle_type{
		Vector_type{0,0,},
		Vector_type{640,480,},
	};
	Scene.player_ship = PlayerShip_type{
		true,
		NewGraphic(0),
		Box_type{
			1,
			( BOXTYPE_COLLISION | BOXTYPE_HURT ),
			Vector_type{0,0,},
			Rectangle_type{
				Vector_type{
					0,
					0,
				},
				Vector_type{
					32,
					16,
				},
			},
		},
		Vector_type{float64((640-32)/2),float64(480-32),},
		Vector_type{0,0,},
		Vector_type{0,0,},
		5,
		ActorMeters_type{
			Meter_type{100,100},
			Meter_type{50,50},
		},
		ActorStats_type{1,1,1,1,1,},
	};
	Scene.player_ship.box.rect.SetOriginFromCentre( Scene.player_ship.centre );
	for i := 0; i < 100; i++{
			Scene.AddEnemy(float64((i % 16)*32), float64(((i/20)*32)+16));
	}
	var ebiten_error error = ebiten.Run( MainLoop, int(Scene.rect.magnitude.x), int(Scene.rect.magnitude.y), 1, "Hello, world?" );
	fmt.Printf("ebiten_error: %v\n", ebiten_error);
}
func MainLoop( screen *ebiten.Image ) error{
	//Input
	input(); //Process input
	logic(); //Process logic

	log.Printf("screen: %v", screen);
	if( !ebiten.IsDrawingSkipped() ){
		var hud_string string = fmt.Sprintf( "SP: %d\nProjectiles: %d", Scene.player_ship.meters.sp.current, len(Scene.projectiles) );
		//Draw player ship
		screen.DrawImage(SpaceInvaders_Images[Scene.player_ship.graphic.image_index], Scene.player_ship.graphic.op);
		//Enemy ships
		for i := 0; i < len(Scene.enemy_ships); i++{
			if( Scene.enemy_ships[i].enabled == true ){
				screen.DrawImage(SpaceInvaders_Images[Scene.enemy_ships[i].graphic.image_index], Scene.enemy_ships[i].graphic.op);
			}
		}
		for i := 0; i < len(Scene.projectiles); i++{
			if( Scene.projectiles[i].enabled == true ){
				screen.DrawImage(SpaceInvaders_Images[Scene.projectiles[i].graphic.image_index], Scene.projectiles[i].graphic.op);
			}
		}
		ebitenutil.DebugPrint( screen, hud_string );
	}
	return nil;
}
