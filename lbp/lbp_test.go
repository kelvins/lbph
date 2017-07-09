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

	expectedlbp := []int64{93, 183, 91, 173, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	lbp, err := ApplyLBP(img)
	if err != nil {
		t.Error(err)
	}
	if len(lbp) == len(expectedlbp) {
		for index := 0; index < len(lbp); index++ {
			/*if lbp[index] != expectedlbp[index] {
				t.Error(
					"Expected value: ", expectedlbp[index],
					"Received value: ", lbp[index],
				)
			}*/
		}
	} else {
		t.Error("Different size")
	}
}
