package action

import (
	"net/http"
)

// UnderlyingViewAction Returns path to file to be served, e.g.
//  "./www/home.html"
type UnderlyingViewAction func() string

type ViewAction struct {
	UnderlyingRun UnderlyingViewAction
}

func (action *ViewAction) Set(run UnderlyingViewAction) {
	action.UnderlyingRun = run
}

func (action ViewAction) Run(w http.ResponseWriter, r *http.Request) {
	run := action.UnderlyingRun()
	http.ServeFile(w, r, run)
}

// Alternative, improvement

type ViewAction2 struct {
	UnderlyingRun UnderlyingViewAction
}

func (action *ViewAction2) Set(run UnderlyingViewAction) {
	action.UnderlyingRun = run
}

func (action ViewAction2) Run(w http.ResponseWriter, r *http.Request) {
	run := action.UnderlyingRun()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_,_ = w.Write([]byte(run))
}