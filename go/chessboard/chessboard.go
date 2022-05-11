package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	counts := make(map[string]int)
	for key, value := range cb {
		for _,b := range value {
			if (b) {
				counts[key] += 1
			}
		}
	}

	return counts[rank]
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	counts := make(map[int]int)

	for _, value := range cb {
		for i,b := range value {
			if (b) {
				counts[i+1] += 1
			}
		}
	}

	return counts[file]
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	var counts int

	for _, value := range cb {
		counts += len(value)
	}

	return counts
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var counts int

	for _, value := range cb {
		for _,b := range value {
			if (b) {
				counts++
			}
		}
	}

	return counts
}
