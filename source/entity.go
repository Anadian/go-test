/**
* @file entity.go
* @brief Defines the Entity structure and interface for the Entity-Component System.
* @author Anadian
* @copyright MITlicensetm(2019,Canosw)
*/
package entity;

//Dependencies
import( 
	//Internal
	//Standard
	//External
	"github.com/emirpasic/gods/lists/arraylist"
);

//Constants
const(
	//Exported Constants
	///Class
	ENTITY_CLASS_UNSPECIFIED uint64 = 0,
	ENTITY_CLASS_SHIP_PLAYER uint64 = 1,
	ENTITY_CLASS_SHIP_ENEMY uint64 = 2,
	ENTITY_CLASS_PROJECTILE uint64 = 8,
	///ID
	//Private Constants
);

//Types, structs, and methods
/*type Entity_interface interface{
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
}*/

type Entity_type struct{
	Enabled bool
	Class uint64
	ID uint64
	Components *arraylist.List
}

//Global Variables
var(
	//Exported Variables
	//Private Variables
);

//Exported Functions

//Private Functions

