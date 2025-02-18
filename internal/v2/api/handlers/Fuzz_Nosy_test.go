package v2handlers

import (
	"net/http"
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

func Fuzz_Nosy_V2Handler_GetDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V2Handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var request *http.Request
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if h == nil || request == nil {
			return
		}

		h.GetDelegation(request)
	})
}

func Fuzz_Nosy_V2Handler_GetDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V2Handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var request *http.Request
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if h == nil || request == nil {
			return
		}

		h.GetDelegations(request)
	})
}

func Fuzz_Nosy_V2Handler_GetFinalityProviders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V2Handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var request *http.Request
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if h == nil || request == nil {
			return
		}

		h.GetFinalityProviders(request)
	})
}

func Fuzz_Nosy_V2Handler_GetNetworkInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V2Handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var request *http.Request
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if h == nil || request == nil {
			return
		}

		h.GetNetworkInfo(request)
	})
}

func Fuzz_Nosy_V2Handler_GetOverallStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V2Handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var request *http.Request
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if h == nil || request == nil {
			return
		}

		h.GetOverallStats(request)
	})
}

func Fuzz_Nosy_V2Handler_GetPrices__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V2Handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var request *http.Request
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if h == nil || request == nil {
			return
		}

		h.GetPrices(request)
	})
}

func Fuzz_Nosy_V2Handler_GetStakerStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V2Handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var request *http.Request
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if h == nil || request == nil {
			return
		}

		h.GetStakerStats(request)
	})
}
