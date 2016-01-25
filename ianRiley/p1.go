package main

import "strconv"

type Tree struct {
  Length int
  Nodes []int
}

type Binary interface {
  Delete(value int)
  Insert(value int)
  ToString() string
}

func (t *Tree) Insert(value int) {
  if t.Nodes == nil {
    t.Nodes = make([]int, 1023, 1023)
  }
  t.Nodes[t.Length] = value
  t.Length++
}

func (t *Tree) Delete(value int) {
  if t.Nodes == nil {
    return
  }
  for i := 0; i < t.Length; i++ {
    if t.Nodes[i] == value {
      t.Nodes[i] = t.Nodes[t.Length-1]
      t.Length--
      return
    }
  }
}

func (t * Tree) ToString() (result string) {
  if t.Nodes == nil {
    return
  }
  return
}

func main() {
  
}