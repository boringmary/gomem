package card

import (
	"context"
	cardGen "github.com/boringmary/gomem/gen/mservices/card"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (crd CardSrv) GetCards(ctx context.Context, _ *cardGen.Empty) (*cardGen.CardsResponse, error) {

	collection := crd.mongoCli.Database(crd.cfg.MongoDB).Collection(crd.cfg.MongoCardCollection)

	var cards []CardEntity

	cursor, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer cursor.Close(ctx)

	cursor.All(ctx, &cards)

	resp := &cardGen.CardsResponse{}
	for _, card := range cards {
		resp.Cards = append(resp.Cards, &cardGen.Card{
			Question: card.Question,
			Answer: card.Answer,
			Interval: card.Interval,
			Repetitions: card.Repetitions,
			EFactor: card.eFactor,
		})
	}

	return resp, nil
}
