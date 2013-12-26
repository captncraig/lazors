package lazors

import (
	"testing"
	"strconv"
)

func TestEmptySquare(t *testing.T){
	c := Empty
	x := getPathSegment(c,North)
	checkSegment(South,false,x,t,"Passthrough N->S")
	x = getPathSegment(c,East)
	checkSegment(West,false,x,t,"Passthrough E->W")
	x = getPathSegment(c,South)
	checkSegment(North,false,x,t,"Passthrough S->N")
	x = getPathSegment(c,West)
	checkSegment(East,false,x,t,"Passthrough W->E")
}

func TestTarget(t *testing.T){
	for orientation := North; orientation <= West; orientation+=East{
		c := Target | orientation
		for direction := North; direction <=West; direction+= East{
			x:= getPathSegment(c,direction)
			checkSegment(NoExit,true,x,t,"Target facing " + strconv.Itoa(int(orientation)))
		}
	}
}

func TestLazor(t *testing.T){
	for orientation := North; orientation <= West; orientation+=East{
		c := Lazor | orientation
		for direction := North; direction <=West; direction+= East{
			x:= getPathSegment(c,direction)
			checkSegment(NoExit,false,x,t,"Lazor facing " + strconv.Itoa(int(orientation)))
		}
	}
}

func TestShield(t *testing.T){
	for orientation := North; orientation <= West; orientation+=East{
		c := Shield | orientation
		for direction := North; direction <=West; direction+= East{
			x:= getPathSegment(c,direction)
			checkSegment(NoExit,orientation != direction,x,t,"Shield facing " + strconv.Itoa(int(orientation)))
		}
	}
}

func TestMirrorFromFacingDirection(t *testing.T){
	for orientation := North; orientation <= West; orientation+=East{
		c := Mirror | orientation
		x := getPathSegment(c,orientation)
		checkSegment(rotate(orientation,1),false,x,t,"Mirror Straight")
	}
}
func TestMirrorFromOtherMirrorSide(t *testing.T){
	for orientation := North; orientation <= West; orientation+=East{
		c := Mirror | orientation
		x := getPathSegment(c,rotate(orientation,1))
		checkSegment(orientation,false,x,t,"Mirror Bounce")
	}
}
func TestMirrorDestruction(t *testing.T){
	for orientation := North; orientation <= West; orientation+=East{
		c := Mirror | orientation
		x := getPathSegment(c,rotate(orientation,2))
		checkSegment(NoExit,true,x,t,"Mirror Destruction")
		x = getPathSegment(c,rotate(orientation,3))
		checkSegment(NoExit,true,x,t,"Mirror Destruction")
	}
}

func TestDoubleMirrorFromDirection(t *testing.T){
	for orientation := North; orientation <= West; orientation+=East{
		c := DoubleMirror | orientation
		x := getPathSegment(c,orientation)
		checkSegment(rotate(orientation,1),false,x,t,"Double Mirror Straight")
	}
}

func TestDoubleMirrorFromOppositeDirection(t *testing.T){
	for orientation := North; orientation <= West; orientation+=East{
		c := DoubleMirror | orientation
		x := getPathSegment(c,rotate(orientation,2))
		checkSegment(rotate(orientation,3),false,x,t,"Double Mirror Backwards")
	}
}

func TestDoubleMirrorFromRight(t *testing.T){
	for orientation := North; orientation <= West; orientation+=East{
		c := DoubleMirror | orientation
		x := getPathSegment(c,rotate(orientation,1))
		checkSegment(orientation,false,x,t,"Double Mirror Straight")
	}
}

func TestDoubleMirrorFromLeft(t *testing.T){
	for orientation := North; orientation <= West; orientation+=East{
		c := DoubleMirror | orientation
		x := getPathSegment(c,rotate(orientation,3))
		checkSegment(rotate(orientation,2),false,x,t,"Double Mirror Straight")
	}
}

func checkSegment(outDir byte,kill bool, segment PathSegment, t *testing.T, description string){
	if segment.ExitDirection != outDir{
		t.Errorf("%v: Expected ExitDirection to be %v, but got %v",description, outDir, segment.ExitDirection)
	}
	if(segment.IsDestroyed != kill){
		t.Errorf("%v: Expected Destroyed Flag to be %v, but got %v",description, kill, segment.IsDestroyed)
	}
}