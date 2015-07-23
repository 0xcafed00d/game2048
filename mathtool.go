package main

func MapValue(val, srcMin, srcMax, dstMin, dstMax int64) int64 {
	return (val-srcMin)*(dstMax-dstMin)/(srcMax-srcMin) + dstMin
}

func MapValuef(val, srcMin, srcMax, dstMin, dstMax float64) float64 {
	return (val-srcMin)*(dstMax-dstMin)/(srcMax-srcMin) + dstMin
}
