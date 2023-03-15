package mincosttickets

func min(xs ...int) int {
	answer := xs[0]
	for _, n := range xs[1:] {
		if n < answer {
			answer = n
		}
	}
	return answer
}

type Key struct {
	day        int
	passEndDay int
}

func mincostTicketsRecur(days []int, costs []int, k int, passEndDay int, memo map[Key]int) int {
	if k == len(days) {
		return 0
	}
	currDay := days[k]
	key := Key{currDay, passEndDay}
	ans, ok := memo[key]
	if ok {
		return ans
	}
	var oneDayPass, sevenDayPass, thirtyDayPass int
	if currDay <= passEndDay {
		noPass := mincostTicketsRecur(days, costs, k+1, passEndDay, memo)
		memo[key] = noPass
		return noPass
	} else {
		oneDayPass = costs[0] + mincostTicketsRecur(days, costs, k+1, currDay, memo)
		sevenDayPass = costs[1] + mincostTicketsRecur(days, costs, k+1, currDay+7-1, memo)
		thirtyDayPass = costs[2] + mincostTicketsRecur(days, costs, k+1, currDay+30-1, memo)
		minCost := min(oneDayPass, sevenDayPass, thirtyDayPass)
		memo[key] = minCost
		return minCost
	}
}

func mincostTickets(days []int, costs []int) int {
	memo := make(map[Key]int)
	return mincostTicketsRecur(days, costs, 0, -1, memo)
}
