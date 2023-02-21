package handler

import (
	"github.com/adough/warehouse_api/internal/parser"
	"github.com/adough/warehouse_api/internal/repository"
)

type service struct {
	rep    repository.Repository
	parser parser.Parser
}

func NewServie(rep repository.Repository, parser parser.Parser) *service {
	return &service{
		rep:    rep,
		parser: parser,
	}
}
