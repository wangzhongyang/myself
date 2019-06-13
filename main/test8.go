package main

import (
	"fmt"
)

func main() {
	gas := []int{5, 8, 2, 8}
	cost := []int{6, 5, 6, 6}
	//gas := []int{1, 1, 3}
	//cost := []int{2, 2, 1}
	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	if !IsResidue(gas, cost) {
		return -1
	}
	maxCostKey := MaxCost(gas, cost)
	beginKey, temp, surplus := maxCostKey+1, maxCostKey+1, 0
	if temp > len(gas)-1 {
		temp = 0
		beginKey = 0
	}
	for {
		if temp > len(gas)-1 {
			temp = 0
		}
		if gas[temp]+surplus >= cost[temp] {
			surplus = gas[temp] + surplus - cost[temp]
			temp++
		} else {
			return -1
		}
		if temp == beginKey {
			return beginKey
		}
	}
}

// 油站的油是否 >= 消耗
func IsResidue(gas []int, cost []int) bool {
	gasSum, costSum := 0, 0
	for k, gasV := range gas {
		gasSum += gasV
		costSum += cost[k]
	}
	if gasSum-costSum >= 0 {
		return true
	} else {
		return false
	}
}

func MaxCost(gas, cost []int) int {
	maxValKey, maxVal, isEquality := 0, 0, false
	for k, v := range cost {
		if maxVal == v {
			kCost, maxCost := gas[k]-cost[k], gas[maxValKey]-cost[maxValKey]
			if kCost > maxCost {
				isEquality = true
				maxVal, maxValKey = v, k
			} else if kCost == maxCost {
				isEquality = false
				maxVal, maxValKey = v, k
			}
		} else if v > maxVal {
			maxVal, maxValKey, isEquality = v, k, false
		}
	}
	if isEquality {
		return maxValKey - 1
	}
	return maxValKey
}
