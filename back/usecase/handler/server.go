package handler

import "photo-share/back/usecase"

type Server struct {
	resourceProvider usecase.ResourceProviderInterface
}

func NewServer(rp usecase.ResourceProviderInterface) *Server {
	return &Server{
		resourceProvider: rp,
	}
}
