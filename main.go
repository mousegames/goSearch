package main

import (
  "fmt"
  "strings"
)

func main() {
  links := []string {
    "https://stackoverflow.com/",
    "https://github.com/",
    "https://git-scm.com/",
    "https://chatgpt.com/",
  }

  input := ""
  fmt.Scanln(&input)

  for num, link := range links {
    if strings.Contains(link, input) {
      fmt.Print(link)
      fmt.Println(" enter num: ", num)
    }
  }

  num := 0
  fmt.Scanln(&num)

  fmt.Println("enter command:\ncurl", links[num])
}

