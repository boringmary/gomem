package card

import (
	"context"
	cardGen "github.com/boringmary/gomem/gen/mservices/card"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
)

func (crd CardSrv) Create(ctx context.Context, msg *cardGen.Card) (*empty.Empty, error) {
	collection := crd.mongoCli.Database(crd.cfg.MongoDB).Collection(crd.cfg.MongoCardCollection)

	cardEntity := CardEntity{
		Question: msg.Question,
		Answer: msg.Answer,
		Interval: msg.Interval,
		Repetitions: msg.Repetitions,
		eFactor: msg.EFactor,
	}
	_, err := collection.InsertOne(context.TODO(), cardEntity)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp := &empty.Empty{}

	return resp, nil
}
