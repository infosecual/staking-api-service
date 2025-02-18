package handler

import (
	"net/http"
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
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

func Fuzz_Nosy_Handler_HealthCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Handler
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

		h.HealthCheck(request)
	})
}

func Fuzz_Nosy_Handler_VerifyUTXOs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Handler
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

		h.VerifyUTXOs(request)
	})
}

func Fuzz_Nosy_ParseBooleanQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *http.Request
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var queryName string
		fill_err = tp.Fill(&queryName)
		if fill_err != nil {
			return
		}
		var isOptional bool
		fill_err = tp.Fill(&isOptional)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		ParseBooleanQuery(r, queryName, isOptional)
	})
}

func Fuzz_Nosy_ParseBtcAddressQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *http.Request
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var queryName string
		fill_err = tp.Fill(&queryName)
		if fill_err != nil {
			return
		}
		var netParam *chaincfg.Params
		fill_err = tp.Fill(&netParam)
		if fill_err != nil {
			return
		}
		if r == nil || netParam == nil {
			return
		}

		ParseBtcAddressQuery(r, queryName, netParam)
	})
}

func Fuzz_Nosy_ParseBtcAddressesQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *http.Request
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var queryName string
		fill_err = tp.Fill(&queryName)
		if fill_err != nil {
			return
		}
		var netParam *chaincfg.Params
		fill_err = tp.Fill(&netParam)
		if fill_err != nil {
			return
		}
		var limit int
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if r == nil || netParam == nil {
			return
		}

		ParseBtcAddressesQuery(r, queryName, netParam, limit)
	})
}

func Fuzz_Nosy_ParseFPSearchQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *http.Request
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var queryName string
		fill_err = tp.Fill(&queryName)
		if fill_err != nil {
			return
		}
		var isOptional bool
		fill_err = tp.Fill(&isOptional)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		ParseFPSearchQuery(r, queryName, isOptional)
	})
}

func Fuzz_Nosy_ParseFPStateQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *http.Request
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var isOptional bool
		fill_err = tp.Fill(&isOptional)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		ParseFPStateQuery(r, isOptional)
	})
}

func Fuzz_Nosy_ParsePaginationQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *http.Request
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		ParsePaginationQuery(r)
	})
}

func Fuzz_Nosy_ParsePublicKeyQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *http.Request
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var queryName string
		fill_err = tp.Fill(&queryName)
		if fill_err != nil {
			return
		}
		var isOptional bool
		fill_err = tp.Fill(&isOptional)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		ParsePublicKeyQuery(r, queryName, isOptional)
	})
}

func Fuzz_Nosy_ParseStateFilterQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *http.Request
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var queryName string
		fill_err = tp.Fill(&queryName)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		ParseStateFilterQuery(r, queryName)
	})
}

func Fuzz_Nosy_ParseTxHashQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *http.Request
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var queryName string
		fill_err = tp.Fill(&queryName)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		ParseTxHashQuery(r, queryName)
	})
}

func Fuzz_Nosy_parseRequestPayload__(f *testing.F) {
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
		var maxUTXOs uint32
		fill_err = tp.Fill(&maxUTXOs)
		if fill_err != nil {
			return
		}
		var netParam *chaincfg.Params
		fill_err = tp.Fill(&netParam)
		if fill_err != nil {
			return
		}
		if request == nil || netParam == nil {
			return
		}

		parseRequestPayload(request, maxUTXOs, netParam)
	})
}
