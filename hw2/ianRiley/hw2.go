package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type Node struct {
  Count int
  Item string
  Left *Node
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

func insert(tree *Node, element string) *Node {
  if tree == nil {
    return &Node{1, element, nil, nil}
  }
  if element < tree.Item {
    if tree.Left == nil {
      tree.Left = &Node{1, element, nil, nil}
      return tree.Left
    }
    return insert(tree.Left, element)
  }
  if element > tree.Item {
    if tree.Right == nil {
      tree.Right = &Node{1, element, nil, nil}
      return tree.Right
    }
    return insert(tree.Right, element)
  }
  tree.Count++
  return tree
}

func insertList(text []string) (*Node, *Node) {
  if text == nil || len(text) == 0 {
    return nil, nil
  }
  
  unigram := insert(nil, text[0])
  insert(unigram, text[1])
  bigram := insert(nil, text[0] + " " + text[1])
  
  for idx := 2; idx < len(text); idx++ {
    insert(unigram, text[idx])
    insert(bigram, text[idx-1] + " " + text[idx])
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