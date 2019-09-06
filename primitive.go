/**
* @file primitive.go
* @brief Mathematical primitives like 2D vector and rectangles.
* @author Anadian
* @copyright MITlicensetm(2019,Canosw)
*/
package main;

//Dependencies
import( 
	//Internal
	//Standard
	//External
);

//Constants
const(
	//Exported Constants
		//Intersection Constants
		INTERSECTION_NONE uint8 = 0 //No collision
		INTERSECTION_INTERSECT uint8 = 1 //One or more corners overlap but the rect_a centre is not in rect_b
		INTERSECTION_ON uint8 = 2 //rect_a's centre is within the bounds of rect_b
		INTERSECTION_IN uint8 = 4 //rect_a is contained entirely within rect_b
		INTERSECTION_HOLDS uint8 = 8 //rect_b's centre is contained within rect_a
		INTERSECTION_CONTAINS uint8 = 16 //rect_b is contained entirely within rect_a
		//"Edge" Constants
		EDGE_NONE uint16 = 0
		EDGE_TOPLEFT uint16 = 1
		EDGE_TOPCENTRE uint16 = 2
		EDGE_TOPRIGHT uint16 = 4
		EDGE_CENTRELEFT uint16 = 8
		EDGE_CENTRE uint16 = 16
		EDGE_CENTRERIGHT uint16 = 32
		EDGE_BOTTOMLEFT uint16 = 64
		EDGE_BOTTOMCENTRE uint16 = 128
		EDGE_BOTTOMRIGHT uint16 = 256
		EDGE_SUMOFALLEDGES uint16 = 511
	//Private Constants
);

//Types, structs, and methods
type Vector_type struct{
	x float64
	y float64
}

type Rectangle_type struct{
	origin Vector_type
	magnitude Vector_type
}

func (rect *Rectangle_type) SetOriginFromCentre( centre Vector_type ){
	rect.origin.x = (centre.x - (rect.magnitude.x / 2));
	rect.origin.y = (centre.y - (rect.magnitude.y / 2));
}

func (vector Vector_type) InRectangle( rect Rectangle_type ) bool{
	if( (vector.x >= rect.origin.x) && (vector.x <= (rect.origin.x + rect.magnitude.x)) && (vector.y >= rect.origin.y) && (vector.y <= (rect.origin.y + rect.magnitude.y)) ){
		return true;
	} else{
		return false;
	}
}

func (rect Rectangle_type) ContainsVector( vector Vector_type ) bool{
	if( (vector.x >= rect.origin.x) && (vector.x <= (rect.origin.x + rect.magnitude.x)) && (vector.y >= rect.origin.y) && (vector.y <= (rect.origin.y + rect.magnitude.y)) ){
		return true;
	} else{
		return false;
	}
}

type Intersection_type struct{
	intersection uint8
	edges uint16 //Edges rect_a has within rect_b
	centre_from_centre Vector_type //The difference/distance rect_a's centre is from rect_b's centre
}

func (rect_a Rectangle_type) IntersectRectangle( rect_b Rectangle_type ) Intersection_type{
//	if( (rect_a.origin.x > rect_b.origin.x) && ((rect_a.origin.x + rect_a.magnitude.x) < (rect_b.origin.x + rect_b.magnitude.x)) && (rect_a.origin.y > rect_b.origin.y) && ((rect_a.origin.y + rect_a.magnitude.y) < (rect_b.origin.y + rect_b.magnitude.y)) ){
//		return true;
//	} else{
//		return false;
//	}
	var _return Intersection_type;
	//TL
	if( rect_b.ContainsVector( Vector_type{rect_a.origin.x,rect_a.origin.y,} ) == true ){
		_return.edges = (_return.edges | EDGE_TOPLEFT);
	}
	//TC
	if( rect_b.ContainsVector( Vector_type{(rect_a.origin.x + (rect_a.magnitude.x / 2)),rect_a.origin.y,} ) == true ){
		_return.edges = (_return.edges | EDGE_TOPCENTRE);
	}
	//TR
	if( rect_b.ContainsVector( Vector_type{(rect_a.origin.x + rect_a.magnitude.x),rect_a.origin.y,} ) == true ){
		_return.edges = (_return.edges | EDGE_TOPRIGHT);
	}
	//CL
	if( rect_b.ContainsVector( Vector_type{rect_a.origin.x,(rect_a.origin.y + (rect_a.magnitude.y / 2)),} ) == true ){
		_return.edges = (_return.edges | EDGE_CENTRELEFT);
	}
	//C
	if( rect_b.ContainsVector( Vector_type{(rect_a.origin.x + (rect_a.magnitude.x / 2)),(rect_a.origin.y + (rect_a.magnitude.y / 2)),} ) == true ){
		_return.edges = (_return.edges | EDGE_CENTRE);
	}
	//CR
	if( rect_b.ContainsVector( Vector_type{(rect_a.origin.x + rect_a.magnitude.x),(rect_a.origin.y + (rect_a.magnitude.y / 2)),} ) == true ){
		_return.edges = (_return.edges | EDGE_CENTRERIGHT);
	}
	//BL
	if( rect_b.ContainsVector( Vector_type{rect_a.origin.x,(rect_a.origin.y + rect_a.magnitude.y),} ) == true ){
		_return.edges = (_return.edges | EDGE_BOTTOMLEFT);
	}
	//BC
	if( rect_b.ContainsVector( Vector_type{(rect_a.origin.x + (rect_a.magnitude.x / 2)),(rect_a.origin.y + rect_a.magnitude.y),} ) == true ){
		_return.edges = (_return.edges | EDGE_BOTTOMCENTRE);
	}
	//BR
	if( rect_b.ContainsVector( Vector_type{(rect_a.origin.x + rect_a.magnitude.x),(rect_a.origin.y + rect_a.magnitude.y),} ) == true ){
		_return.edges = (_return.edges | EDGE_BOTTOMRIGHT);
	}
	//Intersection
	if( _return.edges == EDGE_SUMOFALLEDGES /*Sum of all edges*/ ){
		_return.intersection = (_return.intersection | INTERSECTION_IN);
	}
	if( (_return.edges & EDGE_CENTRE) == EDGE_CENTRE ){
		_return.intersection = (_return.intersection | INTERSECTION_ON);
	}
	if( _return.edges > 0 ){
		_return.intersection = (_return.intersection | INTERSECTION_INTERSECT);
	}
	//Centre from centre
	_return.centre_from_centre.x = (rect_a.origin.x + (rect_a.magnitude.x / 2)) - (rect_b.origin.x + (rect_b.magnitude.x / 2));
	_return.centre_from_centre.y = (rect_a.origin.y + (rect_a.magnitude.y / 2)) - (rect_b.origin.y + (rect_b.magnitude.y / 2));
	return _return;
}

//Global Variables
var(
	//Exported Variables
	//Private Variables
);

//Exported Functions
func MutualIntersectRectangles( rect_a Rectangle_type, rect_b Rectangle_type ) [2]Intersection_type{
	var _return [2]Intersection_type;
	_return[0] = rect_a.IntersectRectangle( rect_b );
	_return[1] = rect_b.IntersectRectangle( rect_a );
	if( (_return[0].intersection & INTERSECTION_IN) == INTERSECTION_IN ){
		_return[1].intersection = (_return[1].intersection | INTERSECTION_CONTAINS);
	}
	if( (_return[1].intersection & INTERSECTION_IN) == INTERSECTION_IN ){
		_return[0].intersection = (_return[0].intersection | INTERSECTION_CONTAINS);
	}
	if( (_return[0].intersection & INTERSECTION_ON) == INTERSECTION_ON ){
		_return[1].intersection = (_return[1].intersection | INTERSECTION_HOLDS);
	}
	if( (_return[1].intersection & INTERSECTION_ON) == INTERSECTION_ON ){
		_return[0].intersection = (_return[0].intersection | INTERSECTION_HOLDS);
	}
	return _return;
}

//Private Functions

