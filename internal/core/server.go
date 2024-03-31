package core

import (
	"context"

	"github.com/huy2272/FormulaOneLite/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	Db     *mongo.Client
	config *config.Configuration // Configuration
	err    error
}

// NewServer will create a new instance of the application
func NewServer(config *config.Configuration) *Server {
	server := &Server{}
	server.config = config

	server.Db, server.err = mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Uri))

	if server.err != nil {
		panic(server.err)
	}

	return server
}
