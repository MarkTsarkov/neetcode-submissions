func twoSum(numbers []int, target int) []int {
	res := make([]int, 0, 2)	
	l:=0
	r:=len(numbers)-1
	for{
		n1 := numbers[l]
		n2 := numbers[r]
		if n1 + n2 == target{
			return []int{l+1, r+1}
		}
		if n1 + n2 > target {
			r--
			continue
		}
		if n1 + n2 < target {
			l++
		}
	} 	
	

	return res
}
