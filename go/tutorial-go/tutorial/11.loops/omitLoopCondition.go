package main

func maxMessages(thresh float64) int {
	totalCost := 0.0

	for i := 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > thresh {
			return i
		}
	}

}

func main() {

}
