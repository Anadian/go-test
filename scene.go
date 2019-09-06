/**
* @file scene.go
* @brief Defines interface ans structs for the Scene type.
* @author Anadian
* @copyright MITlicensetm(2019,Canosw)
*/
package scene;

//Dependencies
import(
	//Internal
	//Standard
	"time"
	//External
);

//Constants
const(
	//Exported Constants
	//Private Constants
);

//Types, structs, and methods
type Appilication_type struct{
	IsRunning bool
	IsAwake bool
	IsFocused bool
	StartTime Time
	Engine *Engine_interface
}
type Scene_interface{
	GetClass()
	GetID()
	Input()
	Logic()
	Output()
	GetEntities() []*Entity_interface
	GetEntity( entity_index uint8 ) Entity_interface
	AddEntity( entity Entity_interface )
	ClearEntities()
	GetRules() []*Rule_interface
	GetRule( rule_index uint8 ) Rule_interface
	AddRule( rule Rule_interface )
	ClearRules()
}



//Global Variables
var(
	//Exported Variables
	//Private Variables
);

//Exported Functions

//Private Functions
