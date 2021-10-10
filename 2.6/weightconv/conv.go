package weightconv

// PToK converts pounds to kilograms
func PToK(p Pound) Kilogram { return Kilogram(p * 0.453592) }

// KToP converts kilograms to pounds
func KToP(k Kilogram) Pound { return Pound(k * 2.20462) }
