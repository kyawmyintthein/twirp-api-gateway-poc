package number_impl

import (
	"context"

	numbersvc "github.com/kyawmyintthein/twirp-api-gateway-poc/rpc/number"
)

type numberHandler struct{}

func NewNumberHandler() numbersvc.NumberService {
	return &numberHandler{}
}

func (h *numberHandler) Add(ctx context.Context, getRandomColorRequest *numbersvc.AddNumberRequest) (*numbersvc.AddResultResponse, error) {
	val := getRandomColorRequest.A + getRandomColorRequest.B
	if val > 255 {
		val = 255
	}
	return &numbersvc.AddResultResponse{Result: val}, nil
}
