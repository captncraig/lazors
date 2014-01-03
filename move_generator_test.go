package lazors

import (
	"testing"
)

func TestEmptyBoard(t *testing.T){
	b := Board{}
	b[12] = Red | Mirror | North
	moves := b.GenerateMoves(Red)
	if(len(moves) != 10){
		t.Errorf("Expected length of 10, but got %v",len(moves))
	}
}

func TestMirrorWithMirror(t *testing.T){
	b := Board{}
	b[12] = Red | Mirror | North
	b[1] = Silver | Mirror | North
	moves := b.GenerateMoves(Red)
	if(len(moves) != 9){
		t.Errorf("Expected length of 9, but got %v",len(moves))
	}
}

func TestDoubleMirrorWithMirror(t *testing.T){
	b := Board{}
	b[12] = Red | DoubleMirror | North
	b[1] = Silver | Mirror | North
	b[2] = Silver | Shield | North
	moves := b.GenerateMoves(Red)
	if(len(moves) != 9){
		t.Errorf("Expected length of 9, but got %v",len(moves))
	}
}

func TestDoubleMirrorWithUnmovables(t *testing.T){
	b := Board{}
	b[12] = Red | DoubleMirror | North
	b[1] = Silver | DoubleMirror | North
	b[2] = Silver | Target | North
	b[3] = Silver | Lazor | North
	moves := b.GenerateMoves(Red)
	if(len(moves) != 6){
		t.Errorf("Expected length of 6, but got %v",len(moves))
	}
}

func TestClassicSetup(t *testing.T){
	b:=ClassicSetup()
	moves := b.GenerateMoves(Red)
	if(len(moves) != 79){
		t.Errorf("Expected length of 79, but got %v",len(moves))
	}

}