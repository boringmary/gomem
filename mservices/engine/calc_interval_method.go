package engine
import (
	"context"
	"github.com/boringmary/gomem/gen/mservices/engine"
)

func (en Engine) CalcInterval(cx context.Context, r *engine.CalcRequest) (*engine.CalcResponse, error) {
		var interval int32
		var easyFactor float64
		repetitions := r.Repetitions
		if r.Quality >= 3 {
			switch repetitions {
			case 0:
				interval = 1
				break
			case 1:
				interval = 6
				break
			default:
				interval = int32((float64(r.PreviousInterval) * r.PreviousEasyFactor))
			}

			repetitions++
			easyFactor = r.PreviousEasyFactor +
				(0.1 - (5.0 - float64(r.Quality)) * (0.08 + (5.0 - float64(r.Quality)) * 0.02))
		} else {
			repetitions = 0
			interval = 1
			easyFactor = r.PreviousEasyFactor
		}

		if easyFactor < 1.3 {
			easyFactor = 1.3
		}
		return &engine.CalcResponse{
			Interval: interval,
			Repetitions: repetitions,
			EasyFactor: easyFactor,
		}, nil
}
