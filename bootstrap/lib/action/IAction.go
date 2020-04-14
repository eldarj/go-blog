package action

import (
	"net/http"
)

type IAction interface {
	Run(w http.ResponseWriter, r *http.Request)
}
