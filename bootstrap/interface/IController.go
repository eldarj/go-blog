package i

type IController interface {
	InitEndpoints()
	GetEndpoints() map[string]IAction
}
