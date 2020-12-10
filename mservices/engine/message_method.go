package engine
import (
	"context"
	"github.com/boringmary/gomem/gen/gomem/mservices/engine"
	"google.golang.org/protobuf/types/known/emptypb"
)

const oneFilePart = 3 * 1024 * 1024 // 3 MB

func (en Engine) GetMessages(context.Context, *emptypb.Empty) (*engine.MessageResponse, error) {
	return &engine.MessageResponse{}, nil
}
