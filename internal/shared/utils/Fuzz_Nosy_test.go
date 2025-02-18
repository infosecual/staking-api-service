package utils

import (
	"testing"

	"github.com/babylonlabs-io/staking-api-service/internal/shared/types"
	"github.com/btcsuite/btcd/wire"
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

// skipping Fuzz_Nosy_Contains__ because parameters include func, chan, or unsupported interface: []T

func Fuzz_Nosy_GetCovenantPksFromStrings__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pkStrings []string
		fill_err = tp.Fill(&pkStrings)
		if fill_err != nil {
			return
		}

		GetCovenantPksFromStrings(pkStrings)
	})
}

func Fuzz_Nosy_GetSchnorrPkFromHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pkHex string) {
		GetSchnorrPkFromHex(pkHex)
	})
}

func Fuzz_Nosy_IsBase64Encoded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		IsBase64Encoded(s)
	})
}

func Fuzz_Nosy_IsValidSignatureFormat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, sigHex string) {
		IsValidSignatureFormat(sigHex)
	})
}

func Fuzz_Nosy_IsValidTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, txHash string) {
		IsValidTxHash(txHash)
	})
}

func Fuzz_Nosy_IsValidTxHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, txHex string) {
		IsValidTxHex(txHex)
	})
}

func Fuzz_Nosy_ParseTimestampToIsoFormat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochtime int64) {
		ParseTimestampToIsoFormat(epochtime)
	})
}

func Fuzz_Nosy_QualifiedStatesToUnbonded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var unbondTxType types.StakingTxType
		fill_err = tp.Fill(&unbondTxType)
		if fill_err != nil {
			return
		}

		QualifiedStatesToUnbonded(unbondTxType)
	})
}

func Fuzz_Nosy_outputsAreEqual__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *wire.TxOut
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *wire.TxOut
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if a == nil || b == nil {
			return
		}

		outputsAreEqual(a, b)
	})
}
