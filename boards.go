package lazors

import(
	"fmt"
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
	

