package main

func isSameTree(p *Tree, q *Tree) bool {
	var result bool
	result = false

	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val == q.Val {
			result = isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}
	}

	return result
}
