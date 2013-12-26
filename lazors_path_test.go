package lazors

import (
	"testing"
)

func TestDefaultPath(t *testing.T){
	b := ClassicSetup()
	x := getFullPath(&b, 0, South)
	if(x.Len() != 12){
		t.Errorf("Was expecting path of length 12, but got %v",x.Len()); 
		return;
	}
}