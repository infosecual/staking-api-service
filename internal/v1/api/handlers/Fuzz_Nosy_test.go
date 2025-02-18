package v1handlers

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

func Fuzz_Nosy_V1Handler_CheckStakerDelegationExist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V1Handler
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

		h.CheckStakerDelegationExist(request)
	})
}

func Fuzz_Nosy_V1Handler_GetBabylonGlobalParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V1Handler
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

		h.GetBabylonGlobalParams(request)
	})
}

func Fuzz_Nosy_V1Handler_GetDelegationByTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V1Handler
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

		h.GetDelegationByTxHash(request)
	})
}

func Fuzz_Nosy_V1Handler_GetFinalityProviders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V1Handler
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

func Fuzz_Nosy_V1Handler_GetOverallStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V1Handler
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

func Fuzz_Nosy_V1Handler_GetPubKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V1Handler
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

		h.GetPubKeys(request)
	})
}

func Fuzz_Nosy_V1Handler_GetStakerDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V1Handler
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

		h.GetStakerDelegations(request)
	})
}

func Fuzz_Nosy_V1Handler_GetStakersStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V1Handler
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

		h.GetStakersStats(request)
	})
}

func Fuzz_Nosy_V1Handler_GetUnbondingEligibility__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V1Handler
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

		h.GetUnbondingEligibility(request)
	})
}

func Fuzz_Nosy_V1Handler_UnbondDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *V1Handler
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

		h.UnbondDelegation(request)
	})
}

func Fuzz_Nosy_parseTimeframeToAfterTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, timeframe string) {
		parseTimeframeToAfterTimestamp(timeframe)
	})
}

func Fuzz_Nosy_parseUnbondDelegationRequestPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var request *http.Request
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if request == nil {
			return
		}

		parseUnbondDelegationRequestPayload(request)
	})
}
