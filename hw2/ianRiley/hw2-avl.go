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

func singleRotationToTheRight(grandparent *Node, parent *Node, child *Node) {
  
}

func singleRotationToTheLeft(grandparent *Node, parent *Node, child *Node) {
  
}

func doubleRotationToTheRight(grandparent *Node, parent *Node, child *Node) {
  
}

func doubleRotationToTheLeft(grandparent *Node, parent *Node, child *Node) {
  
}

func rebalance(node *Node, child *Node, grandchild *Node) {
  node.HeightLeft = getHeight(node.Left)
  node.HeightRight = getHeight(node.Right)
  balance := node.HeightRight - node.HeightLeft
  if balance < -1 || balance > 1 {
    if child == nil {
      fmt.Println(balance, node, child, grandchild)
    }
    if child == node.Left && grandchild == child.Left {
      singleRotationToTheRight(node, child, grandchild)
      return
    } 
    if child == node.Left && grandchild == child.Right {
      doubleRotationToTheRight(node, child, grandchild)
      return
    }
    if child == node.Right && grandchild == child.Left {
      doubleRotationToTheLeft(node, child, grandchild)
      return
    }
    if child == node.Right && grandchild == child.Right {
      singleRotationToTheLeft(node, child, grandchild)
      return
    }
    fmt.Println("ERROR", balance, node, child, grandchild)
  }
  if node.Parent == nil {
    return
  }
  rebalance(node.Parent, node, child)
}

func avlInsert(tree *Node, element string) *Node {
  if tree == nil {
    return &Node{1, 0, 0, element, nil, nil, nil}
  }
  if element < tree.Item {
    if tree.Left == nil {
      tree.Left = &Node{1, 0, 0, element, nil, tree, nil}
      rebalance(tree.Left, nil, nil)
      return tree.Left
    }
    return avlInsert(tree.Left, element)
  }
  if element > tree.Item {
    if tree.Right == nil {
      tree.Right = &Node{1, 0, 0, element, nil, tree, nil}
      rebalance(tree.Right, nil, nil)
      return tree.Right
    }
    return avlInsert(tree.Right, element)
  }
  tree.Count++
  return tree
}

func insertList(text []string) (*Node, *Node) {
  if text == nil || len(text) == 0 {
    return nil, nil
  }
  
  unigram := avlInsert(nil, text[0])
  avlInsert(unigram, text[1])
  bigram := avlInsert(nil, text[0] + " " + text[1])
  
  for idx := 2; idx < len(text); idx++ {
    avlInsert(unigram, text[idx])
    avlInsert(bigram, text[idx-1] + " " + text[idx])
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