package main

import "fmt"
/**
 * 题目
 *     给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
 *     在杨辉三角中，每个数是它左上方和右上方的数的和。
 *     杨辉三角是个动图，具体的看下链接
 *         https://leetcode-cn.com/problems/pascals-triangle/
 * 示例:
 *     输入: 5
 *     输出:
 *        [
 *             [1],
 *            [1,1],
 *           [1,2,1],
 *          [1,3,3,1],
 *         [1,4,6,4,1]
 *        ]
 */
func main () {
	x := generate(0)
	fmt.Println("-", x)
}
/**
 * 思路
 *    找规律，循环
 *    GO写了好几天了，终于写的有点样子
 * 耗时
 *    执行用时 :0 ms, 在所有Go提交中击败了100.00%的用户
 *    内存消耗 :2.4 MB, 在所有Go提交中击败了6.56%的用户
 */
func generate(numRows int) [][]int {

	s := [][]int{}

	// 第一层
	s = append(s, []int{1})
	// 第二层
	s = append(s, []int{1, 1})

	// 第三层
	for i := 3; i <= numRows; i++  {

		x := []int{1} // 前补一

		for j := 0; j < len(s[i-2]) - 1; j++  {
			x = append(x, s[i-2][j] + s[i-2][j+1]) // 补中间
		}

		x = append(x, 1) // 后补一

		s = append(s, x)
	}

	return s[:numRows]
}

