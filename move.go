package main

import (
	"errors"
	"fmt"
	"math"
)

type move struct {
	isFlip         bool
	source, target string // in ban chi notation
}

func parseMove(s string) (m *move, e error) {
	if len(s) < 3 {
		return nil, errors.New("Command string not long enough")
	}
	if s[0:1] == "?" {
		return &move{isFlip: true, source: s[1:3]}, nil
	} else if len(s) < 5 {
		return nil, errors.New("Command string not long enough")
	}
	return &move{isFlip: false, source: s[0:2], target: s[3:5]}, nil
}

func (m *move) getCoords() (srcFile, srcRank, tgtFile, tgtRank int) {
	fmt.Printf("Trying to get coords for: %+v\n", m)
	srcFile = int(m.source[0] - 'A')
	srcRank = int(m.source[1] - '1' + 1)
	if m.isFlip {
		tgtFile, tgtRank = -1, -1
		fmt.Printf("I think they are %d,%d\n", srcRank, srcFile)
		return
	}
	tgtFile = int(m.target[0] - 'A')
	tgtRank = int(m.target[1] - '1' + 1)
	return
}

func (m *move) isValid() bool {
	srcFile, srcRank, tgtFile, tgtRank := m.getCoords()

	if srcFile < 0 || srcRank < 0 || tgtRank < 0 || tgtFile < 0 {
		return false
	}
	if srcFile > 7 || srcRank > 3 || tgtRank > 3 || tgtFile > 7 {
		return false
	}
	if math.Abs(float64(srcFile-tgtFile)) > 1 || math.Abs(float64(srcRank-tgtRank)) > 1 {
		return false
	}
	return true
}
