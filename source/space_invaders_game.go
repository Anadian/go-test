/**
* @file space_invaders_game.go
* @brief The games-specific logic and constructs for Space Invaders.
* @author Anadian
* @copyright MITlicensetm(2019,Canosw)
*/
package space_invaders_game;

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
	SPACE_INVADERS_ENGINE_STATE_UNSPECIFIED uint8 = 0
	SPACE_INVADERS_ENGINE_STATE_INITIALISING uint8 = 1
	SPACE_INVADERS_ENGINE_STATE_QUITING uint8 = 2
	SPACE_INVADERS_ENGINE_STATE_RUNNING uint8 = 3
	SPACE_INVADERS_ENGINE_STATE_PAUSED uint8 = 4
	SPACE_INVADERS_ENGINE_STATE_STOPPED uint8 = 5
	//Private Constants
);

//Types, structs, and methods
type SpaceInvaders_Game_type struct{
	state uint8
	start_time time.Time
	iterations uint64
	frames uint64
	

//Global Variables
var(
	//Exported Variables
	//Private Variables
);

//Exported Functions

//Private Functions

