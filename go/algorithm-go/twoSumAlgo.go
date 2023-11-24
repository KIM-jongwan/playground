package main

import "fmt"

/*
	정수가 저장된 배열 nums가 주어졌을 때, nums의 원소 중 두 숫자를 더해
	target이 될 수 있으면 true를, 불가능하면 false 를 반환 (단, 같은 원소를 두 번 사용할 수 없음)

	input: nums []int / target int
	output: boolean

*/

	/*
	1.단순 이중 for문을 활용하여 해결
		for(i range nums)
			for(j range nums)
					i == j continue;
					nums[i] + nums[j] == target return true;
					nums[i] + nums[j] != target continue;

			시간복잡도: O(n^2)
	*/
func twoSum(nums []int, target int) (bool) {	
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++{
				if (i == j) {
					continue;
				} else if (nums[i] + nums[j] == target) {
					return true;
				}	else if (nums[i] + nums[j] != target) {
					continue;
				}
		}
	}
	return false;
}

func main(){

	fmt.Println(twoSum([]int{1,5,2,4,7,8,10}, 14)) // true
	fmt.Println(twoSum([]int{3,5,3,6,77,2}, 21)) // false


}