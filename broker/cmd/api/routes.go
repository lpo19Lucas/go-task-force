package main

import (
	"net/http"
	"time"
	"fmt"
	// "log"
	"encoding/json"

	// "context"
	// "jwtauth"
	// "httplog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Article struct {
	Name string `json:"name"`
	Id int      `json:"id"`
}

// type Articles struct {
// 	articles list(Article)
// }

func (app *Config) routes() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
    // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
    AllowedOrigins:   []string{"https://*", "http://*"},
    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: true,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	
	// r.Use(middleware.Heartbeat, "/heartbeat")

  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hi"))
	})

	// !!!!!!!!! Ajust routes bellow, decople crud function from this file, call chi router on main go

	// RESTy routes for "articles" resource
  r.Route("/articles", func(r chi.Router) {
    r.Get("/", listArticles)                           // GET /articles
    // r.With(paginate).Get("/{month}-{day}-{year}", listArticlesByDate) // GET /articles/01-16-2017

    // r.Post("/", createArticle)                                        // POST /articles
    // r.Get("/search", searchArticles)                                  // GET /articles/search

    // // Regexp url parameters:
    // r.Get("/{articleSlug:[a-z-]+}", getArticleBySlug)                // GET /articles/home-is-toronto

    // // Subrouters:
    // r.Route("/{articleID}", func(r chi.Router) {
    //   r.Use(ArticleCtx)
    //   r.Get("/", getArticle)                                          // GET /articles/123
    //   r.Put("/", updateArticle)                                       // PUT /articles/123
    //   r.Delete("/", deleteArticle)                                    // DELETE /articles/123
    // })
  })

	return r
}

// func ArticleCtx(next http.Handler) http.Handler {
//   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//     articleID := chi.URLParam(r, "articleID")
//     article, err := dbGetArticle(articleID)
//     if err != nil {
//       http.Error(w, http.StatusText(404), 404)
//       return
//     }
//     ctx := context.WithValue(r.Context(), "article", article)
//     next.ServeHTTP(w, r.WithContext(ctx))
//   })
// }

// func dbGetArticle(aID string) (*Article, error) {
// 	article := &Article{
// 		Id: 4,
// 		Name: "As tracas do rei careca",
// 	}
	
// 	// err := nil
// 	return article, nil
// }

func listArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")	
  // ctx := r.Context()
  // article, ok := ctx.Value("article").(*Article)
	article := &Article{
		Id: 4,
		Name: "Vida Loka",
	}
	
	a, err := json.Marshal(article)
	if err != nil {
			fmt.Printf("Error: %s", err)
			return;
	}

	ok := true

  if !ok {
    http.Error(w, http.StatusText(422), 422)
    return
  }

  w.Write(a)
}