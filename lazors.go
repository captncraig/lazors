package lazors

import (
)



//OOCTTT
//Where OO is orientaion (00-north, 01-East, 10-South, 11-West)
//C is color (0-Silver, 1-Red)
//TTT is Type:
//000 - Empty
//001 - Lazor - Shoots in direction of orientation
//010 - Target - Try to hit this
//011 - Shield - Impervious from one direction
//100 - Mirror - Bounces according to direction:	Destroyed from back. 
//					00: X| 		01:	XXX		10:	XXX		11:	  X
//						X└-			X			  X			  X
//						XXX			X			  X			XXX		(Bounce directions are OO and OO+1)
//              Left hand Rule
//101 - DoubleMirror - Bounces according to direction:
//					00/10:	X			01/11:	  X
//							 X					 X
//							  X					X
type Cell byte

//Constants to define Cells. eg. Mirror | Silver | South
const(
	Empty = Cell(0)
	Lazor = Cell(1)
	Target = Cell(2)
	Shield = Cell(3)
	Mirror = Cell(4)
	DoubleMirror = Cell(5)
	
	Silver = Cell(0 << 3)
	Red = Cell(1 << 3)
	
	North = Cell(0 << 4)
	East = Cell(1 << 4)
	South = Cell(2 <<4)
	West = Cell(3 << 4)
)

//Constants for standalone facings (once extracted from cell)
const(
	NorthFacing = byte(0)
	EastFacing = byte(1)
	SouthFacing = byte(2)
	WestFacing = byte(3)
	NoExit = byte(4)
)

type Board [100]Cell

type PathSegment struct{
	EnterDirection byte
	ExitDirection byte
	IsDestroyed bool
}

func (b Board) GetPath(colorToFire byte) []PathSegment{
	return nil
}

func (c Cell) getPathSegment(enterDirection byte) PathSegment{
	pieceType := Cell(c & 7)
	//Empty passthrough
	if(pieceType == Empty){
		return PathSegment{enterDirection,(enterDirection + 2) % 4, false}
	}
	//Target case:
	if pieceType == Target{
		return PathSegment{enterDirection, NoExit, true}
	}
	return PathSegment{}
}



