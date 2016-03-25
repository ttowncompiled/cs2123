package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type Node struct {
  Count int
  HeightLeft int
  HeightRight int
  Item string
  Left *Node
  Parent *Node
  Right *Node
}

func getNode(tree *Node, element string) *Node {
  if tree == nil {
    return nil
  }
  if element < tree.Item {
    return getNode(tree.Left, element)
  }
  if element > tree.Item {
    return getNode(tree.Right, element)
  }
  return tree
}

func getHeight(node *Node) int {
  if node == nil {
    return 0
  }
  if node.HeightLeft > node.HeightRight {
    return node.HeightLeft + 1
  }
  return node.HeightRight + 1
}

func singleRotationToTheRight(grandparent *Node, parent *Node, child *Node) *Node {
  grandparent.Left = parent.Right
  if grandparent.Left != nil {
    grandparent.Left.Parent = grandparent
  }
  grandparent.HeightLeft = getHeight(grandparent.Left)
  parent.Right = grandparent
  parent.HeightRight = getHeight(parent.Right)
  if grandparent.Parent != nil && grandparent.Parent.Left == grandparent {
    grandparent.Parent.Left = parent
  }
  if grandparent.Parent != nil && grandparent.Parent.Right == grandparent {
    grandparent.Parent.Right = parent
  }
  parent.Parent = grandparent.Parent
  grandparent.Parent = parent
  if parent.Parent == nil {
    return parent
  }
  return rebalance(parent.Parent, parent, child)
}

func singleRotationToTheLeft(grandparent *Node, parent *Node, child *Node) *Node {
  grandparent.Right = parent.Left
  if grandparent.Right != nil {
    grandparent.Right.Parent = grandparent
  }
  grandparent.HeightRight = getHeight(grandparent.Right)
  parent.Left = grandparent
  parent.HeightLeft = getHeight(parent.Left)
  if grandparent.Parent != nil && grandparent.Parent.Left == grandparent {
    grandparent.Parent.Left = parent
  }
  if grandparent.Parent != nil && grandparent.Parent.Right == grandparent {
    grandparent.Parent.Right = parent
  }
  parent.Parent = grandparent.Parent
  grandparent.Parent = parent
  if parent.Parent == nil {
    return parent
  }
  return rebalance(parent.Parent, parent, child)
}

func doubleRotationToTheRight(grandparent *Node, parent *Node, child *Node) *Node {
  parent.Right = child.Left
  if parent.Right != nil {
    parent.Right.Parent = parent
  }
  parent.HeightRight = getHeight(parent.Right)
  child.Left = parent
  child.HeightLeft = getHeight(child.Left)
  grandparent.Left = child.Right
  if grandparent.Left != nil {
    grandparent.Left.Parent = grandparent
  }
  grandparent.HeightLeft = getHeight(grandparent.Left)
  child.Right = grandparent
  child.HeightRight = getHeight(child.Right)
  if grandparent.Parent != nil && grandparent.Parent.Left == grandparent {
    grandparent.Parent.Left = child
  }
  if grandparent.Parent != nil && grandparent.Parent.Right == grandparent {
    grandparent.Parent.Right = child
  }
  child.Parent = grandparent.Parent
  parent.Parent = child
  grandparent.Parent = child
  if child.Parent == nil {
    return child
  }
  return rebalance(child.Parent, child, parent)
}

func doubleRotationToTheLeft(grandparent *Node, parent *Node, child *Node) *Node {
  parent.Left = child.Right
  if parent.Left != nil {
    parent.Left.Parent = parent
  }
  parent.HeightLeft = getHeight(parent.Left)
  child.Right = parent
  child.HeightRight = getHeight(child.Right)
  grandparent.Right = child.Left
  if grandparent.Right != nil {
    grandparent.Right.Parent = grandparent
  }
  grandparent.HeightRight = getHeight(grandparent.Right)
  child.Left = grandparent
  child.HeightLeft = getHeight(child.Left)
  if grandparent.Parent != nil && grandparent.Parent.Left == grandparent {
    grandparent.Parent.Left = child
  }
  if grandparent.Parent != nil && grandparent.Parent.Right == grandparent {
    grandparent.Parent.Right = child
  }
  child.Parent = grandparent.Parent
  parent.Parent = child
  grandparent.Parent = child
  if child.Parent == nil {
    return child
  }
  return rebalance(child.Parent, child, parent)
}

func rebalance(node *Node, child *Node, grandchild *Node) *Node {
  node.HeightLeft = getHeight(node.Left)
  node.HeightRight = getHeight(node.Right)
  balance := node.HeightRight - node.HeightLeft
  if balance < -1 || balance > 1 {
    if child == node.Left && grandchild == child.Left {
      return singleRotationToTheRight(node, child, grandchild)
    } 
    if child == node.Left && grandchild == child.Right {
      return doubleRotationToTheRight(node, child, grandchild)
    }
    if child == node.Right && grandchild == child.Left {
      return doubleRotationToTheLeft(node, child, grandchild)
    }
    if child == node.Right && grandchild == child.Right {
      return singleRotationToTheLeft(node, child, grandchild)
    }
    fmt.Println("ERROR", balance, node, node.Left, node.Right)
  }
  if node.Parent == nil {
    return node
  }
  return rebalance(node.Parent, node, child)
}

func avlInsert(root *Node, tree *Node, element string) *Node {
  if tree == nil {
    return &Node{1, 0, 0, element, nil, nil, nil}
  }
  if element < tree.Item {
    if tree.Left == nil {
      tree.Left = &Node{1, 0, 0, element, nil, tree, nil}
      return rebalance(tree, tree.Left, nil)
    }
    return avlInsert(root, tree.Left, element)
  }
  if element > tree.Item {
    if tree.Right == nil {
      tree.Right = &Node{1, 0, 0, element, nil, tree, nil}
      return rebalance(tree, tree.Right, nil)
    }
    return avlInsert(root, tree.Right, element)
  }
  tree.Count++
  return root
}

func insertList(text []string) (*Node, *Node) {
  if text == nil || len(text) == 0 {
    return nil, nil
  }
  
  unigram := avlInsert(nil, nil, text[0])
  unigram = avlInsert(unigram, unigram, text[1])
  bigram := avlInsert(nil, nil, text[0] + " " + text[1])
  
  for idx := 2; idx < len(text); idx++ {
    unigram = avlInsert(unigram, unigram, text[idx])
    bigram = avlInsert(bigram, bigram, text[idx-1] + " " + text[idx])
  }
  
  return unigram, bigram
}

func inOrderTraversal(tree *Node) string {
  if tree == nil {
    return ""
  }
  left := inOrderTraversal(tree.Left)
  right := inOrderTraversal(tree.Right)
  return left + fmt.Sprintf("%s : %d\n", tree.Item, tree.Count) + right
}

func processText(filename string) []string {
  f, _ := os.Open(filename)
  reader := bufio.NewReader(f)
  scanner := bufio.NewScanner(reader)
  defer f.Close()
  
  var text []string
  for scanner.Scan() {
    if text == nil {
      text = strings.Split(scanner.Text(), " ")
    } else {
      text = append(text, strings.Split(scanner.Text(), " ")...)
    }
  }
  
  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "error:", err)
    os.Exit(1)
  }
  
  return text
}

func outputUnigram(unigram *Node, filename string) {
  f, _ := os.Create(filename)
  writer := bufio.NewWriter(f)
  defer f.Close()
  
  writer.WriteString(inOrderTraversal(unigram))
  writer.Flush()
}

func outputBigram(bigram *Node, filename string) {
  f, _ := os.Create(filename)
  writer := bufio.NewWriter(f)
  defer f.Close()
  
  writer.WriteString(inOrderTraversal(bigram))
  writer.Flush()
}

func cpTraversal(unigram *Node, bigram *Node, writer *bufio.Writer) {
  if bigram == nil {
    return
  }
  
  words := strings.Split(bigram.Item, " ")
  wordCount := getNode(unigram, words[0]).Count
  
  writer.WriteString(fmt.Sprintf("P(%s | %s) = %d/%d \n", words[1], words[0], bigram.Count, wordCount))
  writer.Flush()
  
  cpTraversal(unigram, bigram.Left, writer)
  cpTraversal(unigram, bigram.Right, writer)
}

func outputCP(unigram *Node, bigram *Node, filename string) {
  f, _ := os.Create(filename)
  writer := bufio.NewWriter(f)
  defer f.Close()
  
  cpTraversal(unigram, bigram, writer)
}

func main() {
  textInputFile := os.Args[1]
  outputFilePrefix := os.Args[2]
  
  text := processText(textInputFile)
  unigram, bigram := insertList(text)
  
  fmt.Println("file size:", len(text))
  
  outputUnigram(unigram, outputFilePrefix + ".uni")
  outputBigram(bigram, outputFilePrefix + ".bi")
  outputCP(unigram, bigram, outputFilePrefix + ".cp")
}