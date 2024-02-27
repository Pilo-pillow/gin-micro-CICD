package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{ID: 1, Title: "Gin", Content: "Gin is a fast and modular framework for Go."},
	{ID: 2, Title: "Gorilla", Content: "Gorilla is a fast and modular framework for Go."},
	{ID: 3, Title: "Beego", Content: "Beego is a fast and modular framework for Go."},
}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}
