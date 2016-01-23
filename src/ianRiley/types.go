package ianRiley

type Node struct {
  Value int
  Left *Node
  Right *Node
}

type Tree struct {
  Root *Node
}