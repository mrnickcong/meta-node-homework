package main

import "fmt"

/*
*

  - @param nums:  An array of Integer

  - @param target: target = nums[index1] + nums[index2]

  - @return: [index1 + 1, index2 + 1] (index1 < index2)

    复杂度分析

时间复杂度：O(N 2)，其中 N 是数组中的元素数量。最坏情况下数组中任意两个数都要被匹配一次。
空间复杂度：O(1)。
*/

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

/**
 * 优化
 *
 * @param nums:  An array of Integer
 * @param target: target = nums[index1] + nums[index2]
 * @return: [index1 + 1, index2 + 1] (index1 < index2)
 *
 * 思路：
 *
 * 创建一个字典，将数组中的元素作为键，索引作为值。
 * 遍历数组，对于每一个元素，检查字典中是否存在 target - nums[i]。
 * 如果存在，则返回 [m[target - nums[i]], i]。
 * 如果不存在，则将当前元素及其索引添加到字典中。
 * 如果没有找到满足条件的两个元素，则返回 []。
 *
 * 复杂度分析
 *
 * 时间复杂度：O(N)，其中 N 是数组中的元素数量。
 * 空间复杂度：O(N)。
 */
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	//时间复杂度为O(N 2)
	result := twoSum(nums, target)
	fmt.Println(result)

	//时间复杂度为O(N)
	result2 := twoSum2(nums, target)
	fmt.Println(result2)
}
