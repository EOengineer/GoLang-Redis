package main

import (
"github.com/go-martini/martini"
"github.com/fzzy/radix/redis"
"time"
)


func errHndlr(err error) {
  if err != nil {
  }
}


func main() {
  m := martini.Classic()

  c, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(10)*time.Second)
  errHndlr(err)
  defer c.Close()

  m.Group("/listings", func(r martini.Router) {
    r.Get("", GetListingsIndex)
    r.Get("/:id", GetListing)
    r.Post("/new", NewListing)
    r.Put("/update/:id", UpdateListing)
    r.Delete("/delete/:id", DeleteListing)
})

  m.Run()
}

func GetListingsIndex(params martini.Params) {
  // return all listings
}

func GetListing(params martini.Params) {
  // return single listing
}

func NewListing(params martini.Params) {
  // create new listing
}

func UpdateListing(params martini.Params) {
  // update a listing
}

func DeleteListing(params martini.Params) {
  // delete a listing
}






