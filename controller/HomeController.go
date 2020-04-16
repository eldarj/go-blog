package controller

import (
	a "go-blog/bootstrap/lib/action"
	v "go-blog/bootstrap/lib/view"
	"go-blog/model"
	"log"
	"net/http"
)

type HomeController struct {}

func (controller *HomeController) GetEndpoints() map[string]a.IAction {
	return map[string]a.IAction{
		"/":         indexAction,
		"/home":     homeAction,
		"/about-us": aboutUsAction,
		"/form": formDataAction,
		"/blog": blogDataInHTMLAction,
		"/blog-2": blogModelInHTMLAction,
	}
}

var indexAction = a.ProtoAction{UnderlyingRun: func(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello from my home action!")
}}

var homeAction = a.ProtoAction{UnderlyingRun: func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./www/home.html")
}}

var aboutUsAction = a.ViewAction{UnderlyingRun: func() string {
	return v.View("about-us")
}}


var formDataAction = a.ViewAction2{UnderlyingRun: func() string {
	return v.ViewData("form", map[string]string {
		"data": "Hello friend.",
	})
}}

var blogDataInHTMLAction = a.ViewAction2{UnderlyingRun: func() string {
	return v.ViewData("first-blog-page", map[string]string {
		"mainTitle": "The Go Blog",
		"title": "Go maps in action",
		"content": "One of the most useful data structures in computer science is the hash table. Many hash table" +
			"implementations exist with varying properties, but in general they offer fast lookups, adds, and deletes." +
			"Go provides a built-in map type that implements a hash table.",
		"content2" : "A Go map type looks like this:",
		"code": "map[KeyType]ValueType",
		"content3": "where KeyType may be any type that is comparable (more on this later)," +
			"and ValueType may be any type at all, including another map! This variable m is a map " +
			"of string keys to int values:",
		"subtitle": "Declaration and initialization",
		"subcontent": "Map types are reference types, like pointers or slices, and so the value of m above is nil;" +
			"it doesn't point to an initialized map. A nil map behaves like an empty map when reading, but attempts to" +
			"write to a nil map will cause a runtime panic; don't do that. To initialize a map, use the built in make" +
			"function:",
	})
}}

var blogModelInHTMLAction = a.ViewAction2{UnderlyingRun: func() string {
	model := model.BlogPageModel{ MainTitle: "The Go Blog - Part 2",
		Title: "Compare Structs",
		Content: "You can compare struct values with the comparison operators == and !=." +
			"Two values are equal if their corresponding" +
			"fields are equal.",
	}
	return v.ViewModel("second-blog-page", model)
}}
