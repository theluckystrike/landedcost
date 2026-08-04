// Package landedcost computes the landed cost of an imported shipment from a
// declared customs value, an ad valorem duty rate, and any per-shipment fees.
//
// It deliberately ships NO rate table. Which duty rate applies to a given good is
// a question of tariff classification and of trade measures in force on the day of
// entry, and both change by proclamation. A library that hardcoded rates would be
// wrong silently and on a schedule nobody controls. Supply the rate; this computes
// what follows from it.
//
// Worked through at https://ustariffcalc.com/ if you want the arithmetic in a form
// you can poke at.
package landedcost

// Shipment is one customs entry.
type Shipment struct {
	// CustomsValue is the declared value the ad valorem duty applies to,
	// in the same currency as every other field here.
	CustomsValue float64
	// DutyRate is the ad valorem rate as a fraction, so 0.25 means 25 percent.
	DutyRate float64
	// Freight and Insurance are carriage costs. Whether they sit inside the
	// dutiable value depends on the valuation basis: they are excluded under an
	// FOB basis and included under a CIF basis. See Duty and DutyCIF.
	Freight   float64
	Insurance float64
	// Fees are flat per-entry charges that are not ad valorem.
	Fees float64
}

// Duty returns the ad valorem duty on an FOB basis: the rate applies to the
// customs value alone, and carriage is excluded from the dutiable base.
func (s Shipment) Duty() float64 {
	return s.CustomsValue * s.DutyRate
}

// DutyCIF returns the ad valorem duty on a CIF basis: the rate applies to the
// customs value plus freight and insurance.
//
// The two bases are not a rounding difference. On a shipment whose carriage is a
// large fraction of its value, CIF duty is materially higher than FOB duty at the
// same rate, which is why the valuation basis is worth establishing before the rate.
func (s Shipment) DutyCIF() float64 {
	return (s.CustomsValue + s.Freight + s.Insurance) * s.DutyRate
}

// Total returns the landed cost on an FOB duty basis: value, carriage, duty, fees.
func (s Shipment) Total() float64 {
	return s.CustomsValue + s.Freight + s.Insurance + s.Duty() + s.Fees
}

// TotalCIF returns the landed cost with duty computed on a CIF basis.
func (s Shipment) TotalCIF() float64 {
	return s.CustomsValue + s.Freight + s.Insurance + s.DutyCIF() + s.Fees
}

// EffectiveRate returns duty as a fraction of the landed total, which is the
// number that actually shows up in a margin calculation. It is always lower than
// the headline duty rate, because the rate applies to the customs value while the
// denominator includes carriage and fees.
func (s Shipment) EffectiveRate() float64 {
	t := s.Total()
	if t == 0 {
		return 0
	}
	return s.Duty() / t
}

// UnitLandedCost divides the landed total across a unit count. Returns 0 for a
// non-positive count rather than panicking or returning an infinity.
func (s Shipment) UnitLandedCost(units int) float64 {
	if units <= 0 {
		return 0
	}
	return s.Total() / float64(units)
}
