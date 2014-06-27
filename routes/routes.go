package routes

import (
  "github.com/go-martini/martini"

  "github.com/owais/martini-structure-example/views"
  "github.com/owais/martini-structure-example/apps/games/routes"
)

func Include(m *martini.ClassicMartini) {
  m.Get("/",  views.Home);
  m.Group("/games", routes.Patterns);
}
