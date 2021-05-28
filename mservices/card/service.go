package card

import (
	"context"
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	cardGen "github.com/boringmary/gomem/gen/mservices/card"
)

type CardSrv struct {
	mongoCli          *mongo.Client
	cfg				  config
}

type config struct {
	MongoUrl        			string `env:"MONGO_URI"`
	MongoDB        				string `env:"MONGO_DB"`
	MongoCardCollection         string `env:"MONGO_CARD_COLLECTION"`
}

type Dependences struct {
	Registrator *grpc.Server
	GrpcConn  *grpc.ClientConn
}

func NewCard(dep Dependences) error {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return errors.WithStack(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoUrl))
	if err != nil {
		return errors.WithStack(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return errors.WithStack(err)
	}
	card := &CardSrv{
		mongoCli:       client,
		cfg: 			cfg,
	}

	cardGen.RegisterCardSrvServer(dep.Registrator, card)

	return nil
}
