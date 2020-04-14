package action

import (
	"net/http"
)

type UnderlyingAction func(w http.ResponseWriter, r *http.Request)

type ProtoAction struct {
	UnderlyingRun UnderlyingAction
}

func (action *ProtoAction) Set(run UnderlyingAction) {
	action.UnderlyingRun = run
}

func (action ProtoAction) Run(w http.ResponseWriter, r *http.Request) {
	action.UnderlyingRun(w, r)
}