/**
* @file rule.go
* @brief Implementation of the rule cache.
* @author Anadian
* @copyright MITlicensetm(2019,Canosw)
*/
package rule;

//Dependencies
import(
	//Internal
	//Standard
	//External
	"github.com/gofrs/uuid"
);

//Constants
const(
	//Exported Constants
	//Private Constants
);

//Types, structs, and methods
type Provider_interface interface{
	GetUUID() UUID
	GetVersion
type Rule_interface interface{
	GetProvider() string
	GetName() string
	GetID() uint64
	Call(args ..interface{})
}
type Rule_type struct{


//Global Variables
var(
	//Exported Variables
	//Private Variables
);

//Exported Functions

//Private Functions

