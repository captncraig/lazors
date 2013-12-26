package lazors

import (
)


//OOCTTT
//Where OO is orientaion (00-north, 01-East, 10-South, 11-West)
//C is color (0-Silver, 1-Red)
//TTT is Type:
//000 - Empty
//001 - Lazor - Shoots in direction of orientation
//010 - Targer - Try to hit this
//011 - Shield - Impervious from one direction
//100 - Mirror - Bounces according to direction:	Destroyed from back. 
//					00: X| 		01:	XXX		10:	XXX		11:	  X
//						X└-			X			  X			  X
//						XXX			X			  X			XXX		(Bounce directions are OO and OO+1)
//101 - DoubleMirror - Bounces according to direction:
//					00/10:	X			01/11:	  X
//							 X					 X
//							  X					X
type Cell byte

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
	pieceType := byte(c & 7)
	//Target case:
	if pieceType == 1{
		return PathSegment{enterDirection, 4, true}
	}
	
	//Empty passthrough
	return PathSegment{enterDirection,(enterDirection + 2) % 4, false}

}



