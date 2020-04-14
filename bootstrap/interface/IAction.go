package i

import (
	"net/http"
)

type IAction struct {
	Filepath string
}

func (action IAction) ServeFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, action.Filepath)
}

func (action IAction) Do(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, action.Filepath)
}
