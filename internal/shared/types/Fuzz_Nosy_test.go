package types

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

func Fuzz_Nosy_Error_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var statusCode int
		fill_err = tp.Fill(&statusCode)
		if fill_err != nil {
			return
		}
		var errorCode ErrorCode
		fill_err = tp.Fill(&errorCode)
		if fill_err != nil {
			return
		}
		var msg string
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		e := NewErrorWithMsg(statusCode, errorCode, msg)
		e.Error()
	})
}

func Fuzz_Nosy_DelegationState_ToString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		s1, err := FromStringToDelegationState(s)
		if err != nil {
			return
		}
		s1.ToString()
	})
}

func Fuzz_Nosy_ErrorCode_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e ErrorCode
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.String()
	})
}

func Fuzz_Nosy_StakingTxType_ToString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		s1, err := StakingTxTypeFromString(s)
		if err != nil {
			return
		}
		s1.ToString()
	})
}

func Fuzz_Nosy_NewFinalityProviders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		NewFinalityProviders(path)
	})
}

func Fuzz_Nosy_NewGlobalParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		NewGlobalParams(path)
	})
}

func Fuzz_Nosy_parseCovenantPubKeyFromHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pkStr string) {
		parseCovenantPubKeyFromHex(pkStr)
	})
}
