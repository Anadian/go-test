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
	//"time"
	//External
);

//Constants
const(
	//Exported Constants
	//Private Constants
);

//Types, structs, and methods
/*type Appilication_type struct{
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
}*/
//Rules are functions which are run every tick. Have a providor and an id. Rules trigger event-handling functions as needed.
func Rule_ApplyPhysics();
func Rule_CollisionEvents();
func Rule_BoundCheck_Entities();
func Rule_BoundCheck_Entity_Player();

type Scene_type struct{



//Global Variables
var(
	//Exported Variables
	//Private Variables
);

//Exported Functions

//Private Functions
