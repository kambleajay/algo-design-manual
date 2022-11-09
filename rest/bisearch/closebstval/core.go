package closebstval

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func absDiff(x, y float64) float64 {
	return abs(x - y)
}

func closestValue(root *TreeNode, target float64) int {
	x := root
	closest := root.Val
	for x != nil {
		xf := float64(x.Val)
		if absDiff(xf, target) < absDiff(float64(closest), target) {
			closest = x.Val
		}
		if target < xf {
			x = x.Left
		} else if target > xf {
			x = x.Right
		} else {
			return x.Val
		}
	}
	return closest
}
