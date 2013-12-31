package lazors

import (
	"container/list"

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
//					    X└-			X			  X			  X
//					    XXX			X			  X			XXX		(Bounce directions are OO and OO+1)
//              Left hand Rule
//101 - DoubleMirror - Bounces according to direction:
//					00/10:	X			01/11:	  X
//						 X				 X
//						  X			        X

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
	EnterDirection byte
	ExitDirection byte
	IsDestroyed bool
	Cell byte
	ExitsBoard bool
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

func nextCell(loc byte, direction byte) byte{
	if(direction == North){
		if(loc >= 10){return loc - 10}
		return 255
	}
	if(direction == South){
		if(loc <= 69){return loc + 10}
		return 255
	}
	if(direction == East){
		if(loc % 10 < 9){return loc + 1}
		return 255
	}
	if(direction == West){
		if(loc % 10 > 0){return loc - 1}
		return 255
	}
	return 255
}

func (b *Board) GetFullPath( startLoc byte, startFacing byte) *list.List{
	l := list.New()
	l.PushBack(&PathSegment{NoExit,startFacing,false,startLoc,false})
	for{
		last:= l.Back().Value.(*PathSegment)
		enterDirection := rotate(last.ExitDirection,2)
		cell := nextCell(last.Cell,last.ExitDirection)
		if cell == 255{
			last.ExitsBoard = true
			break
		}
		seg := getPathSegment(b[cell],enterDirection)
		seg.Cell = cell
		seg.EnterDirection = enterDirection
		l.PushBack(&seg)
		if(seg.IsDestroyed){
			break
		}
	}
	return l
}

func getPathSegment(cell byte, enterDirection byte) PathSegment{
	pieceType := pieceType(cell)
	facingDir := facing(cell)
	if(pieceType == Empty){
		return PathSegment{ExitDirection:rotate(enterDirection,2)}
	}
	
	if(pieceType == Mirror){
		if(enterDirection == facingDir){
			return PathSegment{ExitDirection:rotate(enterDirection,1)}
		}
		if(enterDirection == rotate(facingDir,1)){
			return PathSegment{ExitDirection:facing(cell)}
		}
		return PathSegment {ExitDirection:NoExit, IsDestroyed:true}
	}
	
	if(pieceType == DoubleMirror){
		if(enterDirection == facingDir || enterDirection == rotate(facingDir,2)){
			return PathSegment{ExitDirection:rotate(enterDirection,1)}
		}
		return PathSegment{ExitDirection:rotate(enterDirection,3)}
	}
	
	if pieceType == Target{
		return PathSegment{ExitDirection:NoExit,IsDestroyed: true}
	}
	
	if pieceType == Shield{
		return PathSegment{ExitDirection:NoExit, IsDestroyed: enterDirection != facingDir}
	}
	
	if pieceType == Lazor{
		return PathSegment{ExitDirection:NoExit}
	}
	return PathSegment{ExitDirection:NoExit}
}



