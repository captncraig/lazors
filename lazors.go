package lazors

import (
)

//00FFCTTT
//Where FF is Facing (00-north, 01-East, 10-South, 11-West)
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

//Constants to define Cells. eg. Mirror | Silver | South
const(
	Empty = 		byte(0)
	Lazor = 		byte(1)
	Target = 		byte(2)
	Shield = 		byte(3)
	Mirror = 		byte(4)
	DoubleMirror = 	byte(5)
	
	Silver = 		byte(0 << 3)
	Red = 			byte(1 << 3)
	
	North = 		byte(0 << 4)
	East = 			byte(1 << 4)
	South = 		byte(2 << 4)
	West = 			byte(3 << 4)
	NoExit = 		byte(4 << 4) //Only valid in PathSegments. Don't put this in a cell
)

type Board [80]byte

type PathSegment struct{
	ExitDirection byte
	IsDestroyed bool
}

func (b Board) GetPath(colorToFire byte) []PathSegment{
	return nil
}

func pieceType(cell byte) byte{
	return cell & 7
}

func facing(cell byte) byte{
	return cell & West
}

func color(cell byte) byte{
	return cell & Red
}

func rotate(facingDir byte, steps byte) byte {
	return facing(facingDir + (East * steps)) 
}

func getPathSegment(cell byte, enterDirection byte) PathSegment{
	pieceType := pieceType(cell)
	facingDir := facing(cell)
	if(pieceType == Empty){
		return PathSegment{rotate(enterDirection,2), false}
	}
	
	if(pieceType == Mirror){
		if(enterDirection == facingDir){
			return PathSegment{rotate(enterDirection,1),false}
		}
		if(enterDirection == rotate(facingDir,1)){
			return PathSegment{facing(cell), false}
		}
		return PathSegment {NoExit, true}
	}
	
	if(pieceType == DoubleMirror){
		if(enterDirection == facingDir || enterDirection == rotate(facingDir,2)){
			return PathSegment{rotate(enterDirection,1),false}
		}
		return PathSegment{rotate(enterDirection,3),false}
	}
	
	if pieceType == Target{
		return PathSegment{NoExit, true}
	}
	
	if pieceType == Shield{
		return PathSegment{NoExit, enterDirection != facingDir}
	}
	
	if pieceType == Lazor{
		return PathSegment{NoExit, false}
	}
	return PathSegment{NoExit, false}
}



