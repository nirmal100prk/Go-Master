package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	root := &Node{data: 10}
	root.left = &Node{data: 20}
	root.right = &Node{data: 30}
	root.left.left = &Node{data: 5}
	root.left.right = &Node{data: 25}
	//fmt.Println(FindMaxElementByIteration(root))
	//ReverseLevelOrder(root)
	//PostOrderTraversal(root)
	DeleteTree(&root)
	fmt.Println(root)
	PostOrderTraversal(root)

}

// Tree traversal
// Post-order-traversal
func PostOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.left)
	PostOrderTraversal(root.right)
	fmt.Println(root.data)
}

// Give an algorithm for finding maximum element in a binary tree
func FindMaxElement(n *Node) int {
	var max int
	if n != nil {
		rootVal := n.data
		left := FindMaxElement(n.left)
		right := FindMaxElement(n.right)

		if left > right {
			max = left
		} else {
			max = right
		}
		if rootVal > max {
			return rootVal
		}

	}
	return max
}

// Give an algorithm for finding maximum element in a binary tree without recursion
// Find max iterative
func FindMaxElementByIteration(n *Node) int {
	if n == nil {
		return -1 << 31
	}
	maxVal := n.data
	// the nodes are added by level, each node is appended into the queue
	// queue for level order traversal(Breadth-First Search)
	queue := []*Node{n}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.data > maxVal {
			maxVal = node.data
		}

		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}

	}
	return maxVal
}

// Give an algorithm for inserting an element in a binary tree
// insert using level order traversal
func Insert(n *Node, val int) *Node {
	if n == nil {
		return &Node{data: val}
	}
	queue := []*Node{n}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.left == nil {
			node.left = &Node{data: val}
			break
		} else {
			queue = append(queue, node.left)
		}

		if node.right == nil {
			node.right = &Node{data: val}
			break
		} else {
			queue = append(queue, node.right)
		}
	}
	return n
}

func sizeOfBinaryTree(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + sizeOfBinaryTree(root.left) + sizeOfBinaryTree(root.right)
}

// Give	an	algorithm for printing the level order data in reverse order
func ReverseLevelOrder(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}
	stack := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		stack = append(stack, node.data)
		if node.left != nil {
			queue = append(queue, node.left)
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Println(stack[i])
	}

}

// Give an algorithm for deleting the tree
// post order traversal = left >> right >> root
func DeleteTree(root **Node) {
	if *root == nil {
		return
	}
	DeleteTree(&((*root).left))
	DeleteTree(&((*root).right))
	*root = nil
}
