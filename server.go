package main

import (
"fmt"
"github.com/go-martini/martini"
"github.com/fzzy/radix/redis"
"time"
"os"
)

func main() {
  m := martini.Classic()

  c, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(10)*time.Second)
  errHndlr(err)
  defer c.Close()

  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Run()
}


func errHndlr(err error) {
  if err != nil {
    fmt.Println("error:", err)
    os.Exit(1)
  }
}




