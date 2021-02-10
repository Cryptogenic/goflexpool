package utils

import "math"

const GweiToETHRatio = 0.000000001

// HashrateUnit type values.
const (
	HashesPerSecond HashrateUnit = iota
	KiloHashesPerSecond
	MegaHashesPerSecond
	GigaHashesPerSecond
	TeraHashesPerSecond
	PetaHashesPerSecond
)

// Power exponentials for unit conversions.
const (
	BasePow10Exponential = 0
	KiloPow10Exponential = 3
	MegaPow10Exponential = 6
	GigaPow10Exponential = 9
	TeraPow10Exponential = 12
	PetaPow10Exponential = 15
)

// Units for measuring hashrates.
type HashrateUnit int

// ConvertHashrate takes an input hashrate, as well as an input and output HashrateUnit, and converts the input to the
// output hashrate with the given unit conversion. Returns the converted hashrate as an integer.
func ConvertHashrate(inputHashrate int, inputHashrateUnit HashrateUnit, outputHashrateUnit HashrateUnit) int {
	// First we convert the input to hashes/second
	hashesPerSecond := 0

	switch inputHashrateUnit {
	case HashesPerSecond:
		hashesPerSecond = inputHashrate * int(math.Pow10(BasePow10Exponential))
	case KiloHashesPerSecond:
		hashesPerSecond = inputHashrate * int(math.Pow10(KiloPow10Exponential))
	case MegaHashesPerSecond:
		hashesPerSecond = inputHashrate * int(math.Pow10(MegaPow10Exponential))
	case GigaHashesPerSecond:
		hashesPerSecond = inputHashrate * int(math.Pow10(GigaPow10Exponential))
	case TeraHashesPerSecond:
		hashesPerSecond = inputHashrate * int(math.Pow10(TeraPow10Exponential))
	case PetaHashesPerSecond:
		hashesPerSecond = inputHashrate * int(math.Pow10(PetaPow10Exponential))
	}

	// Now convert to output units
	switch outputHashrateUnit {
	case HashesPerSecond:
		return hashesPerSecond / int(math.Pow10(BasePow10Exponential))
	case KiloHashesPerSecond:
		return hashesPerSecond / int(math.Pow10(KiloPow10Exponential))
	case MegaHashesPerSecond:
		return hashesPerSecond / int(math.Pow10(MegaPow10Exponential))
	case GigaHashesPerSecond:
		return hashesPerSecond / int(math.Pow10(GigaPow10Exponential))
	case TeraHashesPerSecond:
		return hashesPerSecond / int(math.Pow10(TeraPow10Exponential))
	case PetaHashesPerSecond:
		return hashesPerSecond / int(math.Pow10(PetaPow10Exponential))
	}

	return -1
}

// ConvertGweiToEth takes a gwei value as an int and returns it's value in ethereum as a float64.
func ConvertGweiToEth(gwei int) float64 {
	return float64(gwei) * GweiToETHRatio
}

// ConvertEthToGwei takes an eth value as a float64 and returns it's value in gwei as an int.
func ConvertEthToGwei(eth float64) int {
	return int(eth / GweiToETHRatio)
}
