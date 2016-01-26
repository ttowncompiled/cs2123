package main

import "strconv"

type Node struct {
  Item int
  Left *Node
  Right *Node
}

type Tree struct {
  Root *Node
}

func Delete(tree *Tree, node *Node, value int) bool {
  if node == nil {
    return false
  }
  if value == node.Item {
    if node.Right != nil {
      node.Item = node.Right.Item
      if !Delete(tree, node.Right, node.Right.Item) {
        node.Right = nil
      }
    } else if node.Left != nil {
      node.Item = node.Left.Item
      if !Delete(tree, node.Left, node.Left.Item) {
        node.Left = nil
      }
    } else if tree.Root == node {
      tree.Root = nil
    } else {
      return false
    }
    return true
  }
  if value < node.Item && node.Left != nil {
    if !Delete(tree, node.Left, value) {
      if node.Left.Item == value {
        node.Left = nil
        return true
      }
      return false
    }
    return true
  }
  if value > node.Item && node.Right != nil {
    if !Delete(tree, node.Right, value) {
      if node.Right.Item == value {
        node.Right = nil
        return true
      }
      return false
    }
    return true
  }
  return false
}

func Insert(tree *Tree, node *Node, value int) bool {
  if tree.Root == nil {
    tree.Root = &Node{value, nil, nil}
    return true
  }
  if node == nil {
    return false
  }
  if value >= node.Item {
    if node.Right == nil {
      node.Right = &Node{value, nil, nil}
      return true
    }
    return Insert(tree, node.Right, value)
  }
  if value < node.Item {
    if node.Left == nil {
      node.Left = &Node{value, nil, nil}
      return true
    }
    return Insert(tree, node.Left, value)
  }
  return false
}

func ToString(node *Node) string {
  if node == nil {
    return ""
  }
  left := ToString(node.Left)
  if left != "" {
    left += ", "
  }
  right := ToString(node.Right)
  if right != "" {
    right = ", " + right
  }
  return left + strconv.Itoa(node.Item) + right
}

func displayBST(tree *Tree) {
  println(ToString(tree.Root))
}

func main() {
  t := &Tree{}
  Insert(t, t.Root, 4)
  Insert(t, t.Root, 2)
  Insert(t, t.Root, 1)
  Insert(t, t.Root, 3)
  Insert(t, t.Root, 6)
  Insert(t, t.Root, 5)
  Insert(t, t.Root, 7)
  displayBST(t)
  Delete(t, t.Root, 7)
  displayBST(t)
  Delete(t, t.Root, 3)
  displayBST(t)
  Delete(t, t.Root, 1)
  displayBST(t)
}