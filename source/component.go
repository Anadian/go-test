/**
* @file component.go
* @brief Components for possible ECS.
* @author Anadian
* @copyright MITlicensetm(2019,Canosw)
*/
package main;

//Dependencies
import( 
	//Internal
	//Standard
	"fmt"
	//"log"
	//External
);

//Constants
const(
	//Exported Constants
	//Private Constants
);

//Types, structs, and methods
type Scene_type interface{
	GetEntities() []Entity_type
	GetEntity( entity_index uint64 ) (int,Entity_type)
	AddEntity( entity Entity_type ) int
	RemoveEntity( entity_index uint64 ) int
	ToJSON() []byte
	FromJSON( json []byte ) int
}

type Entity_type interface{
	Enable() int
	Disable() int
	IsEnabled() bool
	GetClass() uint64
	GetID() uint64
	GetComponents() []Component_type
	GetComponent( component_index uint16 ) (int,Component_type)
	AddComponent( component Component_type ) int
	RemoveComponent( component_index uint16 ) int
	ToJSON() []byte
	FromJSON( json []byte ) int
}

type Component_type interface{
	Enable() int
	Disable() int
	IsEnabled() bool
	GetClass() uint64
	GetID() uint64
	ToJSON() []byte
	FromJSON( json []byte ) int
}

type MainScene_Scene_type struct{
	Entities []Entity_type
}

func (main_scene MainScene_Scene_type) GetEntities() []Entity_type{
	var _return []Entity_type;
	for i := 0; i < len(main_scene.Entities); i++{
		_return = append(_return, main_scene.Entities[i]);
	}
	return _return;
}

func (main_scene MainScene_Scene_type) GetEntity( entity_index uint64 ) (int, Entity_type){
	var _return int = 1;
	var entity Entity_type;
	if( int(entity_index) < len(main_scene.Entities) ){
		entity = main_scene.Entities[entity_index];
		_return = 0;
	}
	return _return, entity;
}

func (main_scene *MainScene_Scene_type) AddEntity( entity Entity_type ) int{
	var _return int = 1;
	main_scene.Entities = append( main_scene.Entities, entity );
	_return = 0;
	return _return;
}

func (main_scene *MainScene_Scene_type) RemoveEntity( entity_index uint64 ) int{
	return 1;
}

func (main_scene MainScene_Scene_type) ToJSON() []byte{
	return []byte{1,1};
}

func (main_scene *MainScene_Scene_type) FromJSON( json []byte ) int{
	return 1;
}

type PlayerShip_Entity_type struct{
	Enabled bool
	Class uint64
	ID uint64
	Components []Component_type
}

func (player_ship *PlayerShip_Entity_type) Enable() int{
	var _return = 1;
	player_ship.Enabled = true;
	_return = 0;
	return _return;
}

func (player_ship *PlayerShip_Entity_type) Disable() int{
	var _return = 1;
	player_ship.Enabled = false;
	_return = 0;
	return _return;
}

func (player_ship PlayerShip_Entity_type) IsEnabled() bool{
	return player_ship.Enabled;
}

func (player_ship PlayerShip_Entity_type) GetClass() uint64{
	return player_ship.Class;
}

func (player_ship PlayerShip_Entity_type) GetID() uint64{
	return player_ship.ID;
}

func (player_ship PlayerShip_Entity_type) GetComponents() []Component_type{
	return nil;
}

func (player_ship PlayerShip_Entity_type) GetComponent( component_index uint16 ) (int,Component_type){
	return 1, nil;
}

func (player_ship PlayerShip_Entity_type) AddComponent( component Component_type ) int{
	return 1;
}

func (player_ship PlayerShip_Entity_type) RemoveComponent( component_index uint16 ) int{
	return 1;
}

func (player_ship PlayerShip_Entity_type) ToJSON() []byte{
	return nil;
}

func (player_ship PlayerShip_Entity_type) FromJSON( json []byte ) int{
	return 1;
}


//Global Variables
var(
	//Exported Variables
	//Private Variables
);

//Exported Functions

//Private Functions
func main(){
	var main_scene MainScene_Scene_type = MainScene_Scene_type{
		nil,
	};
	var player_ship PlayerShip_Entity_type = PlayerShip_Entity_type{
		true,
		0,
		0,
		nil,
	};
	main_scene.AddEntity( &player_ship );
	var entities []Entity_type = main_scene.GetEntities();
	fmt.Printf("main_scene: %v\nplayer_ship: %v\nentities: %v\n", main_scene, player_ship, entities);
}
