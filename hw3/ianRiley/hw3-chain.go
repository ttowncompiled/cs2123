package main

import (
  "bufio"
  "fmt"
  "hash/fnv"
  "os"
  "strings"
)

type Word struct {
  Count int
  Value string
  Next *Word
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

func outputUnigram(unigram []*Word, filename string) {
  f, _ := os.Create(filename)
  writer := bufio.NewWriter(f)
  defer f.Close()
  
  writer.WriteString(toString(unigram))
  writer.Flush()
}

func outputBigram(bigram []*Word, filename string) {
  f, _ := os.Create(filename)
  writer := bufio.NewWriter(f)
  defer f.Close()
  
  writer.WriteString(toString(bigram))
  writer.Flush()
}

func outputCP(unigram, bigram []*Word, filename string) {
  f, _ := os.Create(filename)
  writer := bufio.NewWriter(f)
  defer f.Close()
  
  for i := 0; i < len(bigram); i++ {
    if bigram[i] == nil {
      continue
    }
    words := strings.Split(bigram[i].Value, " ")
    wordCount := probe(unigram, words[0]).Count
    
    writer.WriteString(fmt.Sprintf("P(%s | %s) = %d/%d \n", words[1], words[0], bigram[i].Count, wordCount))
    writer.Flush()
  }
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

func hash(s string, m int) int {
  h := fnv.New32a()
  h.Write([]byte(s))
  return int(h.Sum32()) % m
}

func probe(table []*Word, s string) *Word {
  key := hash(s, len(table))
  word := table[key]
  for word != nil && word.Value != s {
    word = word.Next
  }
  if word == nil {
    fmt.Println(">>> ERROR in probe!")
    os.Exit(1)
  }
  return word
}

func insert(table []*Word, s string) {
  key := hash(s, len(table))
  word := table[key]
  if word == nil {
    table[key] = &Word{1, s, nil}
    return
  }
  for word.Next != nil && word.Value != s {
    word = word.Next
  }
  if word.Value == s {
    word.Count++
    return
  }
  word.Next = &Word{1, s, nil}
}

func insertList(text []string) ([]*Word, []*Word) {
  if text == nil || len(text) == 0 {
    return nil, nil
  }
  
  unigram := make([]*Word, len(text), len(text))
  bigram := make([]*Word, len(text), len(text))
  
  insert(unigram, text[0])
  for idx := 1; idx < len(text); idx++ {
    insert(unigram, text[idx])
    insert(bigram, text[idx-1] + " " + text[idx])
  }
  
  return unigram, bigram
}

func toString(table []*Word) (result string) {
  for i := 0; i < len(table); i++ {
    if table[i] == nil {
      continue
    }
    result += fmt.Sprintf("%s : %d\n", table[i].Value, table[i].Count)
  }
  return
}