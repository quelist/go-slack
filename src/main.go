package main

import (
  "os"
  "fmt"
  "slacksender"
)

func main() {

  if len(os.Args) > 1 {
    msg := os.Args[1]
    webhook_url := "https://hooks.slack.com/services/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"
    slacksender.SendTOSlack(msg, webhook_url)
    } else {
      fmt.Println("Please enter message in CL")
    }
 
}