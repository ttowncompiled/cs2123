package main

import "strconv"

type Tree struct {
  Length int
  Nodes []int
}

func Insert(tree *Tree, value int) bool {
  if tree == nil || tree.Nodes == nil {
    tree.Nodes = make([]int, 1023, 1023)
  }
  tree.Nodes[tree.Length] = value
  tree.Length++
  return true
}

func Delete(tree *Tree, value int) bool {
  if tree == nil || tree.Nodes == nil {
    return false
  }
  for i := 0; i < tree.Length; i++ {
    if tree.Nodes[i] == value {
      tree.Nodes[i] = tree.Nodes[tree.Length-1]
      tree.Length--
      return true
    }
  }
  return false
}

func ToString(tree *Tree, node int) string {
  if tree == nil || tree.Nodes == nil {
    return ""
  }
  if node >= tree.Length {
    return ""
  }
  left := ToString(tree, 2 * node + 1)
  if left != "" {
    left += ", "
  }
  right := ToString(tree, 2 * node + 2)
  if right != "" {
    right = ", " + right
  }
  return left + strconv.Itoa(tree.Nodes[node]) + right
}

func main() {
  t := &Tree{}
  Insert(t, 1)
  Insert(t, 2)
  Insert(t, 3)
  Insert(t, 4)
  Insert(t, 5)
  Insert(t, 6)
  Insert(t, 7)
  println(ToString(t, 0))
  Delete(t, 7)
  println(ToString(t, 0))
  Delete(t, 3)
  println(ToString(t, 0))
  Delete(t, 1)
  println(ToString(t, 0))
}