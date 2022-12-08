// package main

// //	func main(){
// //		noda := [1,2,3,4,5,6]
// //		fmt.Println(bfs(noda))
// //	}
// //
// // Definition for a Node.
// type Node struct {
// 	Val   int
// 	Left  *Node
// 	Right *Node
// 	Next  *Node
// }

// type NodeWithLevel struct {
// 	node  *Node
// 	level int
// }

// func bfs(root *Node) *Node {
// 	if root == nil { //to handle edge case
// 		return root
// 	}
// 	q := []NodeWithLevel{
// 		{
// 			node:  root,
// 			level: 0,
// 		},
// 	}
// 	visited := []*Node{}
// 	for len(q) > 0 {
// 		vertex := q[0]
// 		node, level := vertex.node, vertex.level
// 		visited = append(visited, node)
// 		q = q[1:] //dequeue first node in queue(fifo)

// 		if node.Left != nil { //have both left and right since it's a perfect binary tree
// 			leftNode := NodeWithLevel{
// 				node:  node.Left,
// 				level: level + 1,
// 			}
// 			q = append(q, leftNode) //append left-child to back of queue(fifo)

// 			rightNode := NodeWithLevel{
// 				node:  node.Right,
// 				level: level + 1,
// 			}
// 			q = append(q, rightNode) //append right-child to back of queue(fifo)
// 		}
// 	}
// 	return root
// }
