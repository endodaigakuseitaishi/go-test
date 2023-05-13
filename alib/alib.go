package alib

func Average(nums []int) float64 {
	total := 0
	for _, num := range nums {
    total += num
  }
	return float64(total / len(nums))
}