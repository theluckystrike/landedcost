package landedcost

import "testing"

func almost(a, b float64) bool { d := a - b; if d < 0 { d = -d }; return d < 1e-9 }

func TestDutyBases(t *testing.T) {
	s := Shipment{CustomsValue: 10000, DutyRate: 0.25, Freight: 1500, Insurance: 500, Fees: 200}
	if !almost(s.Duty(), 2500) {
		t.Fatalf("FOB duty: got %v want 2500", s.Duty())
	}
	if !almost(s.DutyCIF(), 3000) {
		t.Fatalf("CIF duty: got %v want 3000", s.DutyCIF())
	}
	if s.DutyCIF() <= s.Duty() {
		t.Fatal("CIF duty must exceed FOB duty when carriage is non-zero")
	}
}

func TestTotals(t *testing.T) {
	s := Shipment{CustomsValue: 10000, DutyRate: 0.25, Freight: 1500, Insurance: 500, Fees: 200}
	if !almost(s.Total(), 14700) {
		t.Fatalf("total: got %v want 14700", s.Total())
	}
	if !almost(s.TotalCIF(), 15200) {
		t.Fatalf("totalCIF: got %v want 15200", s.TotalCIF())
	}
}

func TestEffectiveRateIsBelowHeadline(t *testing.T) {
	s := Shipment{CustomsValue: 10000, DutyRate: 0.25, Freight: 1500, Insurance: 500, Fees: 200}
	if s.EffectiveRate() >= s.DutyRate {
		t.Fatalf("effective %v should be below headline %v", s.EffectiveRate(), s.DutyRate)
	}
}

func TestUnitCostGuards(t *testing.T) {
	s := Shipment{CustomsValue: 1000, DutyRate: 0.1}
	if s.UnitLandedCost(0) != 0 || s.UnitLandedCost(-5) != 0 {
		t.Fatal("non-positive unit counts must return 0")
	}
	if !almost(s.UnitLandedCost(100), 11) {
		t.Fatalf("unit cost: got %v want 11", s.UnitLandedCost(100))
	}
}
