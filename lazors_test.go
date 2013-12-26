package lazors

import (
	"testing"
	"strconv"
)

func TestEmptySquare(t *testing.T){
	c := Empty
	x := getPathSegment(c,North)
	checkSegment(North,South,false,x,t,"Passthrough N->S")
	x = getPathSegment(c,East)
	checkSegment(East,West,false,x,t,"Passthrough E->W")
	x = getPathSegment(c,South)
	checkSegment(South,North,false,x,t,"Passthrough S->N")
	x = getPathSegment(c,West)
	checkSegment(West,East,false,x,t,"Passthrough W->E")
}

func TestTarget(t *testing.T){
	for orientation := North; orientation <= West; orientation+=East{
		c := Target | orientation
		for direction := North; direction <=West; direction+= East{
			x:= getPathSegment(c,direction)
			checkSegment(direction,NoExit,true,x,t,"Target facing " + strconv.Itoa(int(orientation)))
		}
	}
}

func TestShield(t *testing.T){
	for orientation := North; orientation <= West; orientation+=East{
		c := Shield | orientation
		for direction := North; direction <=West; direction+= East{
			x:= getPathSegment(c,direction)
			checkSegment(direction,NoExit,orientation != direction,x,t,"Shield facing " + strconv.Itoa(int(orientation)))
		}
	}
}

func checkSegment(inDir byte, outDir byte,kill bool, segment PathSegment, t *testing.T, description string){
	if segment.EnterDirection != inDir{
		t.Errorf("%v: Expected EnterDirection to be %v, but got %v",description, inDir, segment.EnterDirection)
	}
	if segment.ExitDirection != outDir{
		t.Errorf("%v: Expected ExitDirection to be %v, but got %v",description, outDir, segment.ExitDirection)
	}
	if(segment.IsDestroyed != kill){
		t.Errorf("%v: Expected Destroyed Flag to be %v, but got %v",description, kill, segment.IsDestroyed)
	}
}