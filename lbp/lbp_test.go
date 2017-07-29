package lbp

import (
	"testing"
	"github.com/kelvins/lbph/common"
)

func TestApplyLBP(t *testing.T) {
	img, err := common.LoadImage("../dataset/test/4.png")
	if err != nil {
		t.Error(err)
	}

	var expectedlbp [][]uint8
	expectedlbp = append(expectedlbp, []uint8{ 91, 190,  93, 179})
	expectedlbp = append(expectedlbp, []uint8{238, 245, 255, 206})
	expectedlbp = append(expectedlbp, []uint8{115, 255, 175, 119})
	expectedlbp = append(expectedlbp, []uint8{205, 186, 125, 218})

	lbp, err := ApplyLBP(img)
	if err != nil {
		t.Error(err)
	}

	for row := 0; row < len(lbp); row++ {
		for col := 0; col < len(lbp[row]); col++ {
			if lbp[row][col] != expectedlbp[row][col] {
				t.Error(
					"Expected value: ", expectedlbp[row][col],
					"Received value: ", lbp[row][col],
				)
			}
		}
	}
}