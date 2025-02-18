package v2queuehandler

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/staking-api-service/internal/shared/services"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_V2QueueHandler_ActiveStakingHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var services *services.Services
		fill_err = tp.Fill(&services)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var messageBody string
		fill_err = tp.Fill(&messageBody)
		if fill_err != nil {
			return
		}
		if services == nil {
			return
		}

		h := NewV2QueueHandler(services)
		h.ActiveStakingHandler(ctx, messageBody)
	})
}

func Fuzz_Nosy_V2QueueHandler_HandleUnprocessedMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var services *services.Services
		fill_err = tp.Fill(&services)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var messageBody string
		fill_err = tp.Fill(&messageBody)
		if fill_err != nil {
			return
		}
		var receipt string
		fill_err = tp.Fill(&receipt)
		if fill_err != nil {
			return
		}
		if services == nil {
			return
		}

		qh := NewV2QueueHandler(services)
		qh.HandleUnprocessedMessage(ctx, messageBody, receipt)
	})
}

func Fuzz_Nosy_V2QueueHandler_UnbondingStakingHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var services *services.Services
		fill_err = tp.Fill(&services)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var messageBody string
		fill_err = tp.Fill(&messageBody)
		if fill_err != nil {
			return
		}
		if services == nil {
			return
		}

		h := NewV2QueueHandler(services)
		h.UnbondingStakingHandler(ctx, messageBody)
	})
}

func Fuzz_Nosy_V2QueueHandler_WithdrawableStakingHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var services *services.Services
		fill_err = tp.Fill(&services)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var messageBody string
		fill_err = tp.Fill(&messageBody)
		if fill_err != nil {
			return
		}
		if services == nil {
			return
		}

		h := NewV2QueueHandler(services)
		h.WithdrawableStakingHandler(ctx, messageBody)
	})
}

func Fuzz_Nosy_V2QueueHandler_WithdrawnStakingHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var services *services.Services
		fill_err = tp.Fill(&services)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var messageBody string
		fill_err = tp.Fill(&messageBody)
		if fill_err != nil {
			return
		}
		if services == nil {
			return
		}

		h := NewV2QueueHandler(services)
		h.WithdrawnStakingHandler(ctx, messageBody)
	})
}
