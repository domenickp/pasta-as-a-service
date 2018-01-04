package main

import (
	"fmt"
	"math"

	"github.com/divan/num2words"
)

func main() {

	var people float64 = 1
	var flourmix float64 = people * 0.75
	var flour float64 = flourmix / 2
	var semolina float64 = flourmix / 2
	var liquidmix = flourmix * (1 / 2.5)
	var eggs string = num2words.Convert(getEggs(liquidmix, "large"))

	fmt.Println("Use " + getMeasurement(flour) + " cups of flour.")
	fmt.Println("Use " + getMeasurement(semolina) + " cups of semolina.")
	fmt.Println("Make " + getMeasurement(liquidmix) + " cups of the liquid mix.")
	fmt.Println("You will need " + eggs + " large eggs.")
}

func getEggs(totaleggmix float64, eggsize string) int {
	// egg equivalents in cups according to
	// https://www.incredibleegg.org/cooking-school/tips-tricks/egg-sizes-equivalents-and-substitutions/
	const eggMedLrg = 0.2
	const eggXLrgJmb = 0.25

	// all decimal multiples of the values above
	eggMedLrgDivisions := []float64{0, .2, .4, .6, .8, 1}
	eggXLrgJmbDivisions := []float64{0, .25, .5, .75, 1}

	var eggmix float64
	switch eggsize {
	case "medium":
		eggmix = binsearch(eggMedLrgDivisions, totaleggmix)
		return int(eggmix / eggMedLrg)
	case "large":
		eggmix = binsearch(eggMedLrgDivisions, totaleggmix)
		return int(eggmix / eggMedLrg)
	case "xlarge":
		eggmix = binsearch(eggXLrgJmbDivisions, totaleggmix)
		return int(eggmix / eggXLrgJmb)
	case "jumbo":
		eggmix = binsearch(eggXLrgJmbDivisions, totaleggmix)
		return int(eggmix / eggXLrgJmb)
	default:
		return 0
	}

	return 0

}

func getMeasurement(value float64) string {

	// a slice of the decimal forms of fractional cup measurements
	measurements := []float64{0, .25, .33, .5, .66, .75, 1}

	value = binsearch(measurements, value)

	// split the integer out from the fractional value, and save it.
	// the integer gets added to the final value when it is returned.
	var integer float64 // boy that looks stupid
	var fraction float64
	if value > 1 {
		integer = math.Floor(value)
	}
	fraction = math.Mod(value, 1)

	var integerwords string
	if integer != 0 {
		integerwords = num2words.Convert(int(integer))
	}

	var fracwords string
	if integer != fraction {
		fracwords = frac2words(fraction)
	}

	var conjunction string
	if integerwords != "" {
		if fracwords != "" {
			conjunction = " and "
		}
	}

	return integerwords + conjunction + fracwords

}

func frac2words(value float64) string {
	switch value {
	case .25:
		return "one fourth"
	case .33:
		return "one third"
	case .5:
		return "one half"
	case .66:
		return "two thirds"
	case .75:
		return "three fourths"
	default:
		return ""
	}
	return ""
}

// use a binary search algorithm to find the closest decimal value,
// and when we're close, round by hand.
func binsearch(values []float64, amount float64) float64 {

	// get the middle value.
	mid := (len(values) / 2)

	// split the integer out from the fractional value, and save it.
	// the integer gets added to the final value when it is returned.
	var wholeval float64
	if amount > 1 {
		wholeval = math.Floor(amount)
	}
	amount = math.Mod(amount, 1)

	// check to see if we have an exact value of a standard measurement.
	// if so, return.
	if amount == values[mid] {
		return values[mid] + wholeval
	}

	// once we are down to two possible measurements, find the closest
	// and return that value, because math.round isn't available yet.
	// https://github.com/golang/go/issues/20100
	if len(values) == 2 {
		var zero float64 = amount - values[0]
		var one float64 = values[1] - amount

		// return [1] first to round up in the event of an amount precicely
		// in the middle of two values.
		if zero >= one {
			return values[1] + wholeval
		} else {
			return values[0] + wholeval
		}

	}

	// loop on a binary search to find the two closest values for
	// the given amount.
	var result float64
	if amount < values[mid] {
		result = binsearch(values[0:mid+1], amount)
	} else {
		result = binsearch(values[mid:len(values)], amount)
	}
	return result + wholeval
}
