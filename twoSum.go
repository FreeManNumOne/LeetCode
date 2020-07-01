package main

import "fmt"

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

//把数组内的值取出和其它数遍历，进行相加，判断是否和目标数相等，然后将索引位置记录下来
func twoSum(nums []int, target int) []int {
	result := []int{}
	for i, v := range nums {
		//fmt.Println(v)
		for i1, v2 := range nums {
			if v+v2 == target {
				fmt.Println("result", v, v2, target)
				result = append(result, i)               //记录第一个索引
				result = append(result, i1)              //记录第二个索引位置
				nums = append(nums[:i], nums[i+1:]...)   //移除第一个数
				nums = append(nums[:i1], nums[i1+1:]...) //移除第二个数

			}
		}
	}
	return result
}

func twoSum1(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		//fmt.Println(v)
		v := nums[i]
		for j := i + 1; j < len(nums); j++ {
			v2 := nums[j]
			if v+v2 == target {
				//fmt.Println("result", v, v2, target)
				result = append(result, i) //记录第一个索引
				result = append(result, j) //记录第二个索引位置
				fmt.Println(result)
				return result
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 1, 7, 15, 3, 6}
	//nums := []int{2, 7, 11, 15}
	target := 9
	resetult := twoSum1(nums, target)
	fmt.Println(resetult)

}
