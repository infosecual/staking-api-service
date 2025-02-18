package client

import (
	"testing"

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

// skipping Fuzz_Nosy_HttpClient_GetBaseURL__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/http/client.HttpClient

// skipping Fuzz_Nosy_HttpClient_GetDefaultRequestTimeout__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/http/client.HttpClient

// skipping Fuzz_Nosy_HttpClient_GetHttpClient__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/http/client.HttpClient

// skipping Fuzz_Nosy_SendRequest__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/http/client.HttpClient

func Fuzz_Nosy_isAllowedMethod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, method string) {
		isAllowedMethod(method)
	})
}

// skipping Fuzz_Nosy_sendRequest__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/http/client.HttpClient
