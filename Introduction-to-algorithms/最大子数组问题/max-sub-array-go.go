package main

import (
	"fmt"
	"math"
)

// 找出穿过中点的最大子数组
func findCrossSubarray(A []float64, low, mid, high int32) (int32, int32, float64) {
	leftSum := math.Inf(-1)
	var sum float64 = 0
	var leftMax int32 = -1
	for i := mid; i > low-1; i-- {
		sum += A[i]
		if sum > leftSum {
			leftSum = sum
			leftMax = i
		}
	}

	rightSum := math.Inf(-1)
	sum = 0
	var rightMax int32 = -1
	for i := mid + 1; i < high+1; i++ {
		sum += A[i]
		if sum > rightSum {
			rightSum = sum
			rightMax = i
		}
	}

	return leftMax, rightMax, leftSum + rightSum
}

// 找出穿过中点的最大子数组
func findMaximumSubarray(A []float64, low, high int32) (int32, int32, float64) {
	if low == high {
		return low, high, A[low]
	}

	var mid int32 = int32((low + high) / 2)
	leftLow, leftHigh, leftSum := findMaximumSubarray(A, low, mid)
	rightLow, rightHigh, rightSum := findMaximumSubarray(A, mid+1, high)
	crossLow, crossHigh, crossSum := findCrossSubarray(A, low, mid, high)

	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	}

	if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	}

	if crossSum >= leftSum && crossSum >= rightSum {
		return crossLow, crossHigh, crossSum
	}
	return crossLow, crossHigh, crossSum
}

func main() {
	a := []float64{
		13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7,
	}
	low := int32(0)
	high := int32(len(a) - 1)
	//mid := int32((low + high) / 2)

	left_max, right_max, sum := findMaximumSubarray(a, low, high)

	fmt.Println(left_max, "\t", right_max, "\t", sum)
}
