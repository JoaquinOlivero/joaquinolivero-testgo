package main

func ZeroSum(arr []int) []int {
	if arr == nil {
		return []int{}
	}

	var sum int
	mapSum := make(map[int]int)
	mapSum[0] = -1
	nodes := []int{}

	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		idx, found := mapSum[sum]
		if found {
			nodes = nodes[:idx+1]
			continue
		}

		mapSum[sum] = len(nodes)
		nodes = append(nodes, arr[i])
	}

	return nodes

}
