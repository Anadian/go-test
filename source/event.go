/**
* @file event.go
* @brief Defines Event type and interface.
* @author Anadian
* @copyright MITlicensetm(2019,Canosw)
*/
package event;

//Dependencies
import( 
	//Internal
	//Standard
	"time"
	//External
	"github.com/emirpasic/gods/lists/arraylist"
);

//Constants
const(
	//Exported Constants
	//Private Constants
);

//Types, structs, and methods
type Event_interface interface{
}

type Event_type struct{
	Class uint64
	ID uint64
	Time time.Time
	DVars *arraylist.List
}

//Global Variables
var(
	//Exported Variables
	//Private Variables
);

//Exported Functions

//Private Functions

