package main

import (
	"fmt"
	"sort"
)

/*
	정수가 저장된 배열 nums가 주어졌을 때, nums의 원소 중 두 숫자를 더해
	target이 될 수 있으면 true를, 불가능하면 false 를 반환 (단, 같은 원소를 두 번 사용할 수 없음)

	input: nums []int / target int
	output: boolean

*/

/*
	1.golang sort library를 사용하여 nums를 우선 정렬
	  -> golang의 sort는 기본적으로 퀵정렬을 사용하며 시간복잡도는 O(nlogn)

	2. 정렬된 배열의 앞에서 시작하는 포인터i, 뒤에서 시작하는 포인터 j를 선언

	3. i + j 값을 계산
	 -> target과 동일할 경우 return true;
		-> target보다 작을 경우 i+1
		-> target보다 클 경우 j-1
		-> i == j일 경우 return false;
*/
func twoSum(nums []int, target int) (bool){
	sort.Ints(nums);

 var i int = 0
	var j int = len(nums) -1

	for {
 	var sum int = nums[i] + nums[j]

		if(i == j) {
			return false
		}

		if (sum == target) {
			return true
		}	else if (sum > target) {
			j = j	-1
		} else if (sum < target) {
			i = i +1
		}
	}

}


func main(){
	fmt.Println(twoSum([]int{1,5,2,4,7,8,10}, 14)) // true
	fmt.Println(twoSum([]int{3,5,3,6,77,2}, 21)) // false
}