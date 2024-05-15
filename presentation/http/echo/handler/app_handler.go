package handler

type (
	AppHandler interface {
		AuthHandler
		OuranosHandler
	}
)
