package array

func canCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		totalCost, totalGas, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			totalCost += cost[j]
			totalGas += gas[j]
			if totalCost > totalGas {
				break
			}
			cnt++
		}

		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}
	}
	return -1
}

func canCompleteCircuit2(gas []int, cost []int) int {
	for i := 0; i < len(gas); {
		totalCost, totalGas := 0, 0
		count := 0
		pos := 0
		for count < len(gas) {
			pos = (i + count) % len(gas)
			totalCost += cost[pos]
			totalGas += gas[pos]
			if totalCost > totalGas {
				break
			}
			pos++
			count++
		}
		if count == len(gas) {
			return i
		} else {
			i += count + 1
		}
	}
	return -1
}
