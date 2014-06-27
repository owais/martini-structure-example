package routes

import (
  "github.com/go-martini/martini"

  "github.com/owais/martini-structure-example/apps/games/views"
)


func Patterns(r martini.Router) {
  r.Get("/", views.Home)
}
