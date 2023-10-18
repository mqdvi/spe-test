package main

import (
	"fmt"
	"math"
	"strconv"
)

func NarcissisticNumber(n int) bool {
	numberText := strconv.Itoa(n)

	var total int
	for i := 0; i < len(numberText); i++ {
		digit, _ := strconv.Atoi(string(numberText[i]))
		total += int(math.Pow(float64(digit), float64(len(numberText))))
	}

	return total == n
}

func ParityOutlier(numbers []int) interface{} {
	var (
		evenCount = 0
		oddCount  = 0

		result, previoustEven, previoustOdd int
	)
	for _, num := range numbers {
		if num%2 == 0 {
			evenCount++
			previoustEven = num
		} else {
			oddCount++
			previoustOdd = num
		}

		if evenCount > 1 && oddCount == 1 {
			result = previoustOdd
			break
		} else if oddCount > 1 && evenCount == 1 {
			result = previoustEven
			break
		}
	}

	if result == 0 {
		return false
	}

	return result
}

func FindNeedleInHaystack(data []string, needle string) int{
	for i, value := range data {
		if value == needle {
			return i
		}
	}
	return -1
}

func BlueOceanReverse(list1 []int, list2 []int) []int {
	result := []int{}

	list2Map := make(map[int]bool)
	for _, num := range list2 {
		list2Map[num] = true
	}

	for _, num := range list1 {
		if !list2Map[num] {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	fmt.Println(NarcissisticNumber(153))

	fmt.Println(ParityOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))

	fmt.Println(FindNeedleInHaystack([]string{"red", "blue", "yellow", "black", "grey"}, "blue"))

	fmt.Println(BlueOceanReverse([]int{1, 2, 3, 4, 6, 10}, []int{1}))
}
