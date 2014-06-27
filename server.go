package main

import (
  "github.com/go-martini/martini"

  "github.com/owais/martini-structure-example/routes"
)

func main() {
  m := martini.Classic()
  routes.Include(m)
  m.Run();
}
