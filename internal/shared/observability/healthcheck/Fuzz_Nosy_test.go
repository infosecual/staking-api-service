package healthcheck

import (
	"testing"

	"github.com/babylonlabs-io/staking-api-service/internal/v2/queue"
	"github.com/rs/zerolog"
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

func Fuzz_Nosy_SetLogger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var customLogger zerolog.Logger
		fill_err = tp.Fill(&customLogger)
		if fill_err != nil {
			return
		}

		SetLogger(customLogger)
	})
}

func Fuzz_Nosy_queueHealthCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var queues *queue.Queues
		fill_err = tp.Fill(&queues)
		if fill_err != nil {
			return
		}
		if queues == nil {
			return
		}

		queueHealthCheck(queues)
	})
}
