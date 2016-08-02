package engine

import (
	"go-chess/constants"
	"go-chess/models"
	"go-chess/utils"
	"log"
	"strconv"
)

// ParseFEN parses FEN notation
func ParseFEN(fen string, pos *models.SBoard) int {
	if len(fen) <= 0 || pos == nil {
		log.Fatalln("FEN or board structure is nil")
	}
	rank := constants.Rank8
	file := constants.FileA
	piece := 0
	ResetBoard(pos)
	for (rank >= constants.Rank1) && (len(fen) > 0) {
		count := 1
		switch string(fen[0]) {
		case "p":
			piece = constants.Bp
			break
		case "r":
			piece = constants.Br
			break
		case "n":
			piece = constants.Bn
			break
		case "b":
			piece = constants.Bb
			break
		case "k":
			piece = constants.Bk
			break
		case "q":
			piece = constants.Bq
			break
		case "P":
			piece = constants.Wp
			break
		case "R":
			piece = constants.Wr
			break
		case "N":
			piece = constants.Wn
			break
		case "B":
			piece = constants.Wb
			break
		case "K":
			piece = constants.Wk
			break
		case "Q":
			piece = constants.Wq
			break
		case "1":
		case "2":
		case "3":
		case "4":
		case "5":
		case "6":
		case "7":
		case "8":
			piece = constants.Empty
			count, _ = strconv.Atoi(string(fen[0]))
			break
		case "/":
		case " ":
			rank--
			file = constants.FileA
			fen = fen[1:len(fen)]
			continue

		default:
			log.Println("Error in parsing FEN")
			return -1
		}
		for i := 0; i < count; i++ {
			sq64 := (rank * 8) + file
			sq120 := utils.SQ120(sq64)
			if piece != constants.Empty {
				pos.Pieces[sq120] = piece
			}
			file++
		}
		fen = fen[1:len(fen)]
	}
	return 0
}
