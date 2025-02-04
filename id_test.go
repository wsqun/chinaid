package chinaid

import (
	"testing"
	"time"
)

// TestIDNum 正确性测试
func TestIDNum(t *testing.T) {
	testID := "341302199006041233"
	detail, err := IDCard(testID).Decode()
	if err != nil {
		t.Error("Valid id parse failed.")
	}
	if detail.Province != "安徽省" {
		t.Error("Province check failed")
	}
	if detail.City != "宿州市" {
		t.Error("City check failed")
	}
	if detail.District != "埇桥区" {
		t.Error("District check failed")
	}
	birth, err := time.ParseInLocation("20060102", "19900604", location)
	if detail.Birthday != birth {
		t.Error("Birthday check failed")
	}
	if detail.Sex != Male {
		t.Error("Sex check failed")
	}
}

func TestTransform15To18(t *testing.T) {
	testID := "513425330222071"
	newId,err := Transform15To18(testID)
	if err != nil {
		t.Error("Valid id parse failed")
	}
	if newId != "513425193302220718" {
		t.Error("transform fail")
	}
}