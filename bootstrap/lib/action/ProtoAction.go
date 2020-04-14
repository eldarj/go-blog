package action

import (
	"net/http"
)

type _underlyingAction func()

type ProtoAction struct {
	run _underlyingAction
}

func (action *ProtoAction) Set(run _underlyingAction) {
	action.run = run
}

func (action ProtoAction) Run(w http.ResponseWriter, r *http.Request) {
	action.run()
}