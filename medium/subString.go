package medium

import (
	"fmt"
	"sort"
	"strings"
)

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start

}

//找出这两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	s := append(nums1, nums2...)

	sort.Ints(s)

	mid := len(s) / 2
	if len(s)%2 == 0 {
		return float64(s[mid-1]+s[mid]) / 2
	} else {
		return float64(s[mid])
	}

}

//最长回文子串
func longestPalindrome(s string) string {
	max := 0
	var subS string
	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			if isPalindrome(s[i:j]) > max {
				max = isPalindrome(s[i:j])
				subS = s[i:j]
			}
		}

	}
	return subS

}

func isPalindrome(s string) int {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return 0
		}
	}
	return len(s)
}

/* type TreeNode struct {
 	Val int
 	Left *TreeNode
 	Right *TreeNode
 }
//先序、中序建树
func buildTree(preorder []int, inorder []int) *TreeNode {
	for i := range inorder{
		if inorder[i] == preorder[0]{
			T := &TreeNode{
				Val:   preorder[0],
				Left:  buildTree(preorder[1:i+1],inorder[0:i]),
				Right: buildTree(preorder[i+1:],inorder[i+1:]),
			}
			return T
		}
	}
	return nil
}*/

//台阶问题
func numWays(n int) int {
	val1 := 1
	val2 := 1
	for i := 2; i < n+1; i++ {
		val1, val2 = val1+val2, val1
		val1 = val1 % 1000000007
	}
	return val1
}

//查找任意重复数字
func findRepeatNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return 0
}

//二维数组查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if target == matrix[i][j] {
				return true
			}
		}
	}
	return false
}

//字符串替换
func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

/*func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j, k int, word string) bool {
	//已经完全匹配到该word
	if k == len(word) {
		return true
	}

	//i,j边界
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}

	if board[i][j] == word[k] {
		temp := board[i][j]
		//如果board[i][j]的上下左右都找不到符合条件的值，进行还原
		board[i][j] = '*'
		//对board[i][j]上下左右的元素进行匹配
		if dfs(board, i-1, j, k+1, word) ||
			dfs(board, i+1, j, k+1, word) ||
			dfs(board, i, j-1, k+1, word) ||
			dfs(board, i, j+1, k+1, word) {
			return true
		} else {
			//board[i][j]上下左右的元素都没有匹配，还原
			board[i][j] = temp
		}
	}
	return false
}*/

/*
 *机器人移动
func movingCount(m int, n int, k int) int {
	v:=make(map[int]bool)
	// 为了让m,n两个值产生一个唯一的key作为访问过的标记
	// 这里提前交换m,n的值，让m永远不小于n，这样i*m+j作为key时，不会撞到
	if m<n {
		m,n=n,m
	}
	count := 0
	count+=dfs(0,0,m,n,k,v)

	return count
}

// 用来算各位之和
func sum (num int) int {
	s := 0
	for num!=0 {
		s+=num%10
		num/=10
	}
	return s
}

func dfs (i int, j int, m int, n int, k int, v map[int]bool) int {
	count := 0
	// 超过边界条件就不计数
	if  sum(i)+sum(j)>k || i>=m || j>=n || v[i*m+j]{
		return 0
	}
	v[i*m+j] = true
	// 当前的格子算一格
	// 以直角坐标系来看，只需要向右和向下搜索
	count = 1 + dfs(i+1,j,m,n,k,v) + dfs(i,j+1,m,n,k,v)
	return count
}*/

/*
 *股票收益最大
func maxProfit(prices []int) int {
	if len(prices)==0{
		return 0
	}
	profit_0 := make([]int, len(prices)) //不持有股票
	profit_1 := make([]int, len(prices)) //持有股票

	profit_1[0] = -prices[0]
	//状态转移方程
	for x := 1; x < len(prices); x++ { //遍历每天股票价格
		profit_0[x] = max(profit_0[x-1], profit_1[x-1]+prices[x])
		if x-2 < 0{
			profit_1[x] = max(profit_1[x-1], -prices[x]) //对第二天的操作进行判断，是第一天买入第二天持有 还是第一天不动，第二天买入
		}else {
			profit_1[x] = max(profit_1[x-1], profit_0[x-2]-prices[x])
		}
	}
	return profit_0[len(prices)-1]
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
*/

func cuttingRope(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			if dp[j]*dp[i-j] > max {
				max = dp[j] * dp[i-j]
			}
		}
		dp[i] = max
	}
	return dp[n]
}

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	fmt.Print(triangle)
	return triangle[0][0]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMin() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	minimumTotal(triangle)
}

func numTrees(n int) int {
	var res = 1
	for i := 0; i < n; i++ {
		res = res * 2 * (1 + 2*i) / (i + 2)
	}
	return res

}

func isBipartite(graph [][]int) bool {

	flag := make([]int, len(graph))
	var cal func(int, int) bool
	cal = func(tag, i int) bool {
		if flag[i] != 0 {
			return flag[i] == tag
		}
		flag[i] = tag
		for _, j := range graph[i] {
			if !cal(-tag, j) {
				return false
			}
		}
		return true
	}
	for i, _ := range graph {
		if flag[i] == 0 && !cal(1, i) {
			return false
		}
	}
	return true
}

func isBipartite2(graph [][]int) bool {

	flag := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		if flag[i] == 1 {
			continue
		}
		for j := 0; j < len(graph[i]); j++ {
			if flag[j] != 1 {

			}

		}
	}

	return true

}

func TestIsBip() {
	graph := [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}
	fmt.Print(isBipartite(graph))
}

/*func  numWaterBottles(numBottles int, numExchange int) int {
	count := numBottles
	for numBottles/numExchange > 0{
		count += numBottles/numExchange
		numBottles = numBottles/numExchange + numBottles%numExchange
	}
	return count
}*/
func countSubTrees(n int, edges [][]int, labels string) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		tag := labels[i]
		conut := 1
		if i < len(edges) {
			edg := edges[i]
			for edg != nil {
				sub := edg[1]
				if labels[sub] == tag {
					conut++
				}
				if sub == len(edges) {
					break
				}
				edg = edges[sub]
			}
		}
		res[i] = conut
	}
	return res
}

/*func minPathSum(grid [][]int) int {
	m:=len(grid)
	n:=len(grid[0])

	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < m; i++{
		for j := 1; j < n;j++{
			grid[i][j] += minMP(grid[i-1][j],grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}
func minMP(a,b int)int{
	if a < b{
		return a
	}
	return b
}*/

func splitArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	fmt.Println(len(nums))
	for i := len(nums) - 1; i > 0; i-- {
		if isBoth(nums[0], nums[i]) {
			return splitArray(nums[i+1:]) + 1
		}
	}
	return splitArray(nums[1:]) + 1
}

func isBoth(a, b int) bool {
	if a%2 == 0 && b%2 == 0 {
		return true
	}
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	if a == 1 {
		return false
	}
	return true
}

/*class Solution {
public int shortestSubarray(int[] A, int K) {
int N = A.length;
long[] P = new long[N+1];
for (int i = 0; i < N; ++i)
P[i+1] = P[i] + (long) A[i];

// Want smallest y-x with P[y] - P[x] >= K
int ans = N+1; // N+1 is impossible
Deque<Integer> monoq = new LinkedList(); //opt(y) candidates, as indices of P

for (int y = 0; y < P.length; ++y) {
// Want opt(y) = largest x with P[x] <= P[y] - K;
while (!monoq.isEmpty() && P[y] <= P[monoq.getLast()])
monoq.removeLast();
while (!monoq.isEmpty() && P[y] >= P[monoq.getFirst()] + K)
ans = Math.min(ans, y - monoq.removeFirst());

monoq.addLast(y);
}

return ans < N+1 ? ans : -1;
}
}*/

func search() {

}

func TestnumW() {
	a := []int{782581, 227, 113147}
	fmt.Print(splitArray(a))
}
