package lazors

import(
	"fmt"
	"github.com/daviddengcn/go-colortext"
)

type Placement struct{
	Location byte
	Piece byte
	Facing byte
}

func ClassicSetup() Board{
	var b Board
	var placements [13]Placement
	placements[0] = Placement{0,Lazor,South}
	placements[1] = Placement{4,Shield,South}
	placements[2] = Placement{5,Target,South}
	placements[3] = Placement{6,Shield,South}
	placements[4] = Placement{7,Mirror,East}
	placements[5] = Placement{12,Mirror,South}
	placements[6] = Placement{30,Mirror,North}
	placements[7] = Placement{34,DoubleMirror,North}
	placements[8] = Placement{35,DoubleMirror,East}
	placements[9] = Placement{37,Mirror,East}
	placements[10] = Placement{40,Mirror,East}
	placements[11] = Placement{47,Mirror,North}
	placements[12] = Placement{56,Mirror,East}
	addPlacementsToBoard(placements,Red, false,&b)
	addPlacementsToBoard(placements,Silver, true,&b)
	return b
}

func addPlacementsToBoard(p [13]Placement, color byte, shouldRotate bool, b *Board){
	for i := 0; i < len(p); i++ {
		cell := p[i];
		//Horrible hack to simulate ternary operator
		location := map[bool]byte{false: cell.Location, true: 79 - cell.Location}[shouldRotate]
		facing := map[bool]byte{false: cell.Facing, true: rotate(cell.Facing,2)}[shouldRotate]
		b[location] = color | cell.Piece | facing
	}
}

func PrettyPrint(b *Board){
	for i := 0; i < len(b); i++ {
		if(i%10 == 0){fmt.Println()}
		print(b[i])
	}
}

func print(c byte){
	colorVal := color(c)
	
	if colorVal == Red{
		ct.ChangeColor(ct.Red,false,ct.Black,false)
	}
	if colorVal == Silver && c != 0{
		ct.ChangeColor(ct.Yellow,false,ct.Black,false)
	}
	out := ""
	switch pieceType(c){
		case Empty:
			out += "0"
		case Mirror:
			out += "M"
		case Target:
			out += "T"
		case Lazor:
			out += "L"
		case Shield:
			out += "S"
		case DoubleMirror:
			out += "D"
		default:
			out += "?"
	}
	if(c != 0){
		switch facing(c){
			case North:
				out += "^"
			case East:
				out += ">"
			case South: 
				out += "v"
			case West:
				out += "<"
		}
	}else{out+=" "}
	fmt.Print(out)
	ct.ChangeColor(ct.White,false,ct.Black,false)
}
	

