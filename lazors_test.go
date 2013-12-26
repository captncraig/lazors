package lazors

import (
	"testing"
	"strconv"
)

func TestEmptySquare(t *testing.T){
	var c Cell = 0
	x := c.getPathSegment(0)
	checkSegment(0,2,false,x,t,"Passthrough N->S")
	x = c.getPathSegment(1)
	checkSegment(1,3,false,x,t,"Passthrough E->W")
	x = c.getPathSegment(2)
	checkSegment(2,0,false,x,t,"Passthrough S->N")
	x = c.getPathSegment(3)
	checkSegment(3,1,false,x,t,"Passthrough W->E")
}

func TestTarget(t *testing.T){
	for orientation := Cell(0); orientation <= 3; orientation++{
		var c Cell = 1 | (orientation << 4)
		for direction := byte(0); direction <=3; direction++{
			x:=c.getPathSegment(direction)
			checkSegment(direction,4,true,x,t,"Target facing " + strconv.Itoa(int(orientation)))
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