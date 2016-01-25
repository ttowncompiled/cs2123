package main

type Node struct {
  Item int
  Left *Node
  Right *Node
}

type Tree struct {
  Root *Node
}

type BinarySearch interface {
  Delete(value int)
  Insert(value int)
  ToString() string
}

func (t *Tree) Delete(value int) {
  
}

func (t *Tree) Insert(value int) {
  
}

func (t *Tree) ToString() (result string) {
  return
}

func displayBST(tree BinarySearch) {
  
}

func main() {
  
}