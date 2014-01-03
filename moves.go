package lazors

type Move struct{
	CellFrom byte
	CellTo byte
	Rotations byte
}

func (b *Board) GenerateMoves(c byte) []Move{
	moves := make([]Move, 0, 118) //118 is max number of moves
	for i, cell := range b {
		if cell != 0 && color(cell) == c{
			moves = addAdjacentMoves(b,byte(i),moves)
			switch pieceType(cell){
				case Target, DoubleMirror:
					moves = append(moves,Move{byte(i),byte(i),1})
				case Mirror, Shield :
					moves = append(moves,Move{byte(i),byte(i),1})
					moves = append(moves,Move{byte(i),byte(i),3})
				case Lazor:
					if(facing(cell) == North || facing(cell) == South){
						moves = append(moves,Move{byte(i),byte(i),3})
					}else{
						moves = append(moves,Move{byte(i),byte(i),1})
					}
			}
		}
	}
	return moves
}

func addAdjacentMoves(b *Board, i byte,m []Move) []Move{
	row := int(i / 10)
	col := int(i % 10)
	
	for r := row - 1; r <= row + 1; r++{
		for c := col - 1; c<=col + 1; c++{
			if(r >= 0 && r < 8 && c >= 0 && c<10){
				idx := byte(r*10+c)
				if idx == i{continue}
				src := pieceType(b[i])
				target := pieceType(b[idx])
				if target != Empty{
					if(src != DoubleMirror || (target != Mirror && target != Shield)){continue}
				}
				if src == Lazor{continue}
				m = append(m,Move{i,idx,0})
			}
		}
	}
	return m
}