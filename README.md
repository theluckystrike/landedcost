# landedcost

Landed-cost arithmetic for an imported shipment: duty, carriage, fees, and the
effective rate that actually lands in a margin calculation.

## What it does not do

It ships **no rate table**. Which duty rate applies to a good is a question of
tariff classification and of trade measures in force on the day of entry, and both
change by proclamation. A library that hardcoded rates would go wrong silently and
on a schedule nobody controls. You supply the rate; this computes what follows.

## The part people get wrong

The valuation basis, not the rate. On an FOB basis the duty applies to the customs
value alone. On a CIF basis it applies to the customs value plus freight and
insurance. On a shipment whose carriage is a large fraction of its value those are
materially different numbers at the same headline rate, which is why the basis is
worth establishing first.

The other one is the effective rate. Duty as a fraction of the landed total is
always lower than the headline rate, because the rate applies to the customs value
while the total includes carriage and fees.

## Usage

    s := landedcost.Shipment{
        CustomsValue: 10000, DutyRate: 0.25,
        Freight: 1500, Insurance: 500, Fees: 200,
    }
    s.Duty()          // FOB basis
    s.DutyCIF()       // CIF basis
    s.Total()
    s.EffectiveRate()
    s.UnitLandedCost(100)

An interactive version is at [ustariffcalc.com](https://ustariffcalc.com/).

## Licence

MIT.
