package main

import (
  "bufio"
  "crypto/md5"
  "fmt"
  "hash/fnv"
  "os"
  "strings"
)

type Word struct {
  Count int
  Value string
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

func hash1(s string, m int) int {
  h := fnv.New32a()
  h.Write([]byte(s))
  return int(h.Sum32()) % m
}

func hash2(s string, m int) int {
  key := md5.Sum([]byte(s))
  return hash1(string(key[:len(key)]), m)
}

func probe(table []*Word, s string) *Word {
  m := len(table)
  h1 := hash1(s, m)
  h2 := hash2(s, m)
  word := table[h1]
  if word == nil {
    fmt.Println(">>> ERROR in probe! (1)")
    os.Exit(1)
  }
  if word.Value == s {
    return word
  }
  for i := 1; i <= m; i++ {
    word = table[(h1 + i * h2) % m]
    if word == nil {
      fmt.Println(">>> ERROR in probe! (2)")
      os.Exit(1)
    }
    if word.Value == s {
      return word
    }
  }
  fmt.Println(">>> ERROR in probe! (4)")
  os.Exit(1)
  return nil
}

func insert(table []*Word, s string) {
  m := len(table)
  h1 := hash1(s, m)
  h2 := hash2(s, m)
  word := table[h1]
  if word == nil {
    table[h1] = &Word{1, s}
    return
  }
  if word.Value == s {
    word.Count++
    return
  }
  for i := 1; i <= m; i++ {
    word = table[(h1 + i * h2) % m]
    if word == nil {
      table[(h1 + i * h2) % m] = &Word{1, s}
      return
    }
    if word.Value == s {
      word.Count++
      return
    }
  }
  fmt.Println(">>> ERROR in insert!")
  os.Exit(1)
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