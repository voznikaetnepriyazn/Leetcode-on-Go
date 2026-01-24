package leetcode

import "strconv"

func Num1413(nums []int) int {
	startValue := 1
	for {
		total := startValue
		isValid := true
		for _, i := range nums {
			total += i
			if total < 1 {
				isValid = false
				break
			}
		}
		if isValid {
			return startValue
		}
		startValue++
	}
}

func Num119(row int, col int) int {
	if row == 0 || col == 0 || row == col {
		return 1
	}

	return Num119(row-1, col-1) + Num119(row-1, col)
}

func getRow119(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		ans[i] = Num119(rowIndex, i)
	}
	return ans
}

func Num1009(num int) int64 {
	bin := strconv.FormatInt(int64(num), 2)
	run := []rune(bin)
	for i, j := range run {
		switch j {
		case '0':
			run[i] = '1'
		case '1':
			run[i] = '0'
		}
	}
	res, err := strconv.ParseInt(string(run), 2, 64)
	if err != nil {
		return 0
	}
	return res
}

func Num1150(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		count := 0
		if i != target {
			return false

		} else if i == target {
			count++
		}

		if count <= len(nums)/2 {
			return false
		}
	}
	return true
}

func Num830(s string) [][]int {
	var res [][]int
	i := 0

	for j := 0; j < len(s); j++ {
		if s[j] != s[j+1] || j == len(s)-1 {
			if j-i+1 >= 3 {
				res = append(res, []int{i, j})
			}

			i = j + 1
		}
	}
	return res
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func DFS(root *Node, isLonely bool, ans *[]int) {
	if root == nil {
		return
	}

	if isLonely {
		*ans = append(*ans, root.data)
	}

	DFS(root.left, root.right == nil, ans)
	DFS(root.right, root.left == nil, ans)
}

func Num1469(root *Node) []int {
	var ans []int
	DFS(root, false, &ans)
	return ans
}

type Node1 struct {
	data int
	next *Node1
}

func Num1474(head *Node1, m int, n int) *Node1 {
	var currentNode *Node1 = head
	var lastMNode *Node1
	_ = lastMNode

	for currentNode != nil {
		MM, NN := m, n

		for currentNode != nil && MM > 0 {
			lastMNode = currentNode
			currentNode = currentNode.next
			MM--
		}

		for currentNode != nil && NN > 0 {
			currentNode = currentNode.next
			NN--
		}
		lastMNode = currentNode.next
	}
	return head
}

func Num955(arr []string) int {
	var n, m int
	deleteCount := make([]bool, m)
	count := 0

	isSorted := func() bool {
		for i := range n - 1 {
			for j := range m {
				if deleteCount[j] {
					continue
				}
				if arr[i][j] > arr[i+1][j] {
					return false

				} else {
					break
				}
			}
		}
		return true
	}

	for !isSorted() {
		for j := 0; j < m; j++ {
			if deleteCount[j] {
				continue
			}
			for i := range n - 1 {
				if arr[i][j] > arr[i+1][j] {
					deleteCount[j] = true
					break
				}
			}
		}
		for _, c := range deleteCount {
			if c {
				count++
			}
		}
	}
	return count
}

func Num26(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] == nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func Num914(deck []int) bool {
	voc := map[int]int{}
	for i := range deck {
		voc[i]++
	}

	nod := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	g := 0
	for _, j := range voc {
		g = nod(g, j)
	}
	return g > 1
}

func Num844(s, t string) bool {
	ss, tt := len(s)-1, len(t)-1
	skipS, skipT := 0, 0

	for ss >= 0 || tt >= 0 {
		for ss >= 0 {
			if s[ss] == '#' {
				skipS++
				ss--
			} else if skipS > 0 {
				skipS--
				ss--
			} else {
				break
			}
		}

		for tt >= 0 {
			if t[tt] == '#' {
				skipT++
				tt--
			} else if skipT > 0 {
				skipT--
				tt--
			} else {
				break
			}
		}

		if ss >= 0 && tt >= 0 && s[ss] != t[tt] {
			return false
		}
		if (ss >= 0) != (tt >= 0) {
			return false
		}
		ss--
		tt--
	}
	return true
}
