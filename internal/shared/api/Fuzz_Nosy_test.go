package api

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/staking-api-service/internal/shared/config"
	"github.com/babylonlabs-io/staking-api-service/internal/shared/services"
	"github.com/go-chi/chi"
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

func Fuzz_Nosy_Server_SetupRoutes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var services *services.Services
		fill_err = tp.Fill(&services)
		if fill_err != nil {
			return
		}
		var r *chi.Mux
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if cfg == nil || services == nil || r == nil {
			return
		}

		a, err := New(ctx, cfg, services)
		if err != nil {
			return
		}
		a.SetupRoutes(r)
	})
}

func Fuzz_Nosy_Server_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var services *services.Services
		fill_err = tp.Fill(&services)
		if fill_err != nil {
			return
		}
		if cfg == nil || services == nil {
			return
		}

		a, err := New(ctx, cfg, services)
		if err != nil {
			return
		}
		a.Start()
	})
}

// skipping Fuzz_Nosy_registerHandler__ because parameters include func, chan, or unsupported interface: func(*net/http.Request) (*github.com/babylonlabs-io/staking-api-service/internal/shared/api/handlers/handler.Result, *github.com/babylonlabs-io/staking-api-service/internal/shared/types.Error)

// skipping Fuzz_Nosy_writeResponse__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter
