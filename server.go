package main

import (
"github.com/go-martini/martini"
"github.com/fzzy/radix/redis"
"time"
"net/http"
)


func main() {
  m := martini.Classic()

  c, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(10)*time.Second)
  errHndlr(err)
  defer c.Close()

  r := c.Cmd("select", 8)
  errHndlr(r.Err)

  m.Group("/listings", func(r martini.Router) {
    r.Get("", GetListingsIndex)
    r.Get("/:id", GetListing)
    r.Post("/new", NewListing)
    r.Put("/update/:id", UpdateListing)
    r.Delete("/delete/:id", DeleteListing)
  })

  m.Run()
}


type Listing struct {
        Id      int
        Title   string
        Description   string
        Price string
}

func errHndlr(err error) {
  if err != nil {
  }
}

func GetListingsIndex(params martini.Params, w http.ResponseWriter, r *http.Request, c *redis.Client) (int, string) {
  // return all listings
  return 200, "all good"
}

func GetListing(params martini.Params, w http.ResponseWriter, r *http.Request, c *redis.Client) {
  // return single listing
}

func NewListing(params martini.Params, w http.ResponseWriter, r *http.Request, c *redis.Client) {
  // create new listing
}

func UpdateListing(params martini.Params, w http.ResponseWriter, r *http.Request, c *redis.Client) {
  // update a listing
}

func DeleteListing(params martini.Params, w http.ResponseWriter, r *http.Request, c *redis.Client) {
  // delete a listing
}






