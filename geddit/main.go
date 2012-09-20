package main

import (
  "fmt"
  "github.com/ginty/reddit"
  "log"
)

func main() {
  feeds := []string{"golang", "atheism"}

  for _, feed := range feeds {
    items, err := reddit.Get(feed)
    if err != nil {
      log.Fatal(err)
    }
    for _, item := range items {
      fmt.Println(item)
    }
  }
}
