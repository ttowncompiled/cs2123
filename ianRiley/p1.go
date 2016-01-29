package main

import (
  "fmt"
  "strconv"
)

type Node struct {
  Item int
  Left *Node
  Right *Node
}

func traversal(tree *Node, order int) string {
  if tree == nil {
    return ""
  }
  left := traversal(tree.Left, order)
  right := traversal(tree.Right, order)
  if order < 0 {
    // pre-order traversal
    return strconv.Itoa(tree.Item) + " " + left + right
  }
  if order > 0 {
    // post-order traversal
    return left + right + strconv.Itoa(tree.Item) + " "
  }
  // in order traversal
  return left + strconv.Itoa(tree.Item) + " " + right
}

func preOrderTraversal(tree *Node) string {
  return traversal(tree, -1)
}

func inOrderTraversal(tree *Node) string {
  return traversal(tree, 0)
}

func postOrderTraversal(tree *Node) string {
  return traversal(tree, 1)
}

func insert(tree *Node, element int) *Node {
  if tree == nil {
    return &Node{element, nil, nil}
  }
  if element < tree.Item {
    if tree.Left == nil {
      tree.Left = &Node{element, nil, nil}
      return tree.Left
    }
    return insert(tree.Left, element)
  }
  if element > tree.Item {
    if tree.Right == nil {
      tree.Right = &Node{element, nil, nil}
      return tree.Right
    }
    return insert(tree.Right, element)
  }
  // no duplicates
  return tree
}

func insertList(elementList []int) *Node {
  if len(elementList) == 0 {
    return nil
  }
  tree := insert(nil, elementList[0])
  for idx := 1; idx < len(elementList); idx++ {
    insert(tree, elementList[idx])
  }
  return tree
}

func displayBST(tree *Node) {
  println(inOrderTraversal(tree))
}

func main() {
  l := [10]int{100, 3, 3, 200, 5, 8, 5, 200, 0, -4}
  
  s := l[:]

  insertList(s)

  displayBST(insertList(s))

  fmt.Println()
}

// output: -4 0 3 5 8 100 200