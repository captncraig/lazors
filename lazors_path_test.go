package lazors

import (
	"testing"
)

func TestDefaultPath(t *testing.T){
	b := ClassicSetup()
	x := GetFullPath(&b, 0, South)
	if(x.Len() != 12){
		t.Errorf("Was expecting path of length 12, but got %v",x.Len()); 
		return;
	}
	last := x.Back().Value.(*PathSegment)
	if(last.Cell != 70){
		t.Errorf("Exits Wrong cell. Should be 70, but was %v",last.Cell); 
		return;
	}
	if(!last.ExitsBoard){
		t.Errorf("Should exit board but doesn't"); 
		return;
	}
}

func TestKillPharoh(t *testing.T){
	b := ClassicSetup()
	b[32] = 0
	b[44] = 0
	x := GetFullPath(&b, 0, South)
	if(x.Len() != 12){
		t.Errorf("Was expecting path of length 12, but got %v",x.Len()); 
		return;
	}
	last := x.Back().Value.(*PathSegment)
	if(last.Cell != 74){
		t.Errorf("Exits Wrong cell. Should be 74, but was %v",last.Cell); 
		return;
	}
	if(last.ExitsBoard){
		t.Errorf("Should not exit board but does"); 
		return;
	}
	if(!last.IsDestroyed){
		t.Errorf("Should kill but doesn't"); 
		return;
	}
}

func TestKillPyramid(t *testing.T){
	b := ClassicSetup()
	b[32] = Silver | Mirror | North
	x := GetFullPath(&b, 0, South)
	if(x.Len() != 6){
		t.Errorf("Was expecting path of length 6, but got %v",x.Len()); 
		return;
	}
	last := x.Back().Value.(*PathSegment)
	if(last.Cell != 32){
		t.Errorf("Exits Wrong cell. Should be 32, but was %v",last.Cell); 
		return;
	}
	if(last.ExitsBoard){
		t.Errorf("Should not exit board but does"); 
		return;
	}
	if(!last.IsDestroyed){
		t.Errorf("Should kill but doesn't"); 
		return;
	}
}