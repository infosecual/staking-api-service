package ordinals

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/staking-api-service/internal/shared/config"
	"github.com/babylonlabs-io/staking-api-service/internal/shared/types"
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

func Fuzz_Nosy_Ordinals_FetchUTXOInfos__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.OrdinalsConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var utxos []types.UTXOIdentifier
		fill_err = tp.Fill(&utxos)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		c := New(config)
		c.FetchUTXOInfos(ctx, utxos)
	})
}

func Fuzz_Nosy_Ordinals_GetBaseURL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.OrdinalsConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		c := New(config)
		c.GetBaseURL()
	})
}

func Fuzz_Nosy_Ordinals_GetDefaultRequestTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.OrdinalsConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		c := New(config)
		c.GetDefaultRequestTimeout()
	})
}

func Fuzz_Nosy_Ordinals_GetHttpClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.OrdinalsConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		c := New(config)
		c.GetHttpClient()
	})
}

// skipping Fuzz_Nosy_OrdinalsClient_FetchUTXOInfos__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/http/clients/ordinals.OrdinalsClient

// skipping Fuzz_Nosy_OrdinalsClient_GetBaseURL__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/http/clients/ordinals.OrdinalsClient

// skipping Fuzz_Nosy_OrdinalsClient_GetDefaultRequestTimeout__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/http/clients/ordinals.OrdinalsClient

// skipping Fuzz_Nosy_OrdinalsClient_GetHttpClient__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/http/clients/ordinals.OrdinalsClient
