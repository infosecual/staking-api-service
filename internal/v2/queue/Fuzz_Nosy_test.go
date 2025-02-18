package queue

import (
	"testing"

	"github.com/babylonlabs-io/staking-api-service/internal/shared/services"
	"github.com/babylonlabs-io/staking-api-service/internal/shared/types"
	"github.com/babylonlabs-io/staking-queue-client/config"
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

func Fuzz_Nosy_Queues_IsConnectionHealthy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.QueueConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var service *services.Services
		fill_err = tp.Fill(&service)
		if fill_err != nil {
			return
		}
		if cfg == nil || service == nil {
			return
		}

		q, err := New(cfg, service)
		if err != nil {
			return
		}
		q.IsConnectionHealthy()
	})
}

func Fuzz_Nosy_Queues_StartReceivingMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.QueueConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var service *services.Services
		fill_err = tp.Fill(&service)
		if fill_err != nil {
			return
		}
		if cfg == nil || service == nil {
			return
		}

		q, err := New(cfg, service)
		if err != nil {
			return
		}
		q.StartReceivingMessages()
	})
}

func Fuzz_Nosy_Queues_StopReceivingMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.QueueConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var service *services.Services
		fill_err = tp.Fill(&service)
		if fill_err != nil {
			return
		}
		if cfg == nil || service == nil {
			return
		}

		q, err := New(cfg, service)
		if err != nil {
			return
		}
		q.StopReceivingMessages()
	})
}

func Fuzz_Nosy_recordErrorLog__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err *types.Error
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}
		if err == nil {
			return
		}

		recordErrorLog(err)
	})
}
