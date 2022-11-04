package mdy

import "testing"

func TestDataType(t *testing.T) {
	if DateTime != 16 {
		t.Errorf("Expected DataType DateTime  16 actual, got %v", DateTime)
	}

	if Amount != 8 {
		t.Errorf("Expected DataType Amount 8 actual, got %v", Amount)
	}

	if FreeConnection != 21 {
		t.Errorf("Expected DataType FreeConnection 21 actual, got %v", FreeConnection)
	}

	if OtherFields != 30 {
		t.Errorf("Expected DataType OtherFields 30 actual, got %v", OtherFields)
	}
	if Embed != 45 {
		t.Errorf("Expected DataType Embed 45 actual, got %v", Embed)
	}
	if Remark != 10010 {
		t.Errorf("Expected DataType Embed 10010 actual, got %v", Remark)
	}
}

func TestFilterType(t *testing.T) {
	if HasValue != 8 {
		t.Errorf("Expected FilterType DateTime 8 actual, got %v", HasValue)
	}

	if Between != 11 {
		t.Errorf("Expected FilterType Between 11 actual, got %v", Between)
	}

	if Lte != 16 {
		t.Errorf("Expected FilterType Lte 16 actual, got %v", Lte)
	}

	if NDateEnum != 18 {
		t.Errorf("Expected FilterType NDateEnum 18 actual, got %v", NDateEnum)
	}
	if MySelf != 21 {
		t.Errorf("Expected FilterType MySelf 21 actual, got %v", MySelf)
	}
	if ArrNe != 27 {
		t.Errorf("Expected FilterType ArrNe 27 actual, got %v", ArrNe)
	}
	if DateBetween != 31 {
		t.Errorf("Expected FilterType ArrNe 31 actual, got %v", DateBetween)
	}
	if DateBetween != 31 {
		t.Errorf("Expected FilterType ArrNe 31 actual, got %v", DateBetween)
	}
	if DateLte != 36 {
		t.Errorf("Expected FilterType DateLte 36 actual, got %v", DateLte)
	}
	if NormalUser != 41 {
		t.Errorf("Expected FilterType NormalUser 41 actual, got %v", NormalUser)
	}
}
