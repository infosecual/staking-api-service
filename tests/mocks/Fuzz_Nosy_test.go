package mocks

import (
	context "context"
	"testing"

	indexerdbclient "github.com/babylonlabs-io/staking-api-service/internal/indexer/db/client"
	types "github.com/babylonlabs-io/staking-api-service/internal/shared/types"
	v1dbclient "github.com/babylonlabs-io/staking-api-service/internal/v1/db/client"
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

func Fuzz_Nosy_CoinMarketCapClientInterface_GetLatestBtcPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *CoinMarketCapClientInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetLatestBtcPrice(ctx)
	})
}

// skipping Fuzz_Nosy_DBClient_DeleteUnprocessableMessage__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_DBClient_FindPkMappingsByNativeSegwitAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var nativeSegwitAddresses []string
		fill_err = tp.Fill(&nativeSegwitAddresses)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.FindPkMappingsByNativeSegwitAddress(ctx, nativeSegwitAddresses)
	})
}

func Fuzz_Nosy_DBClient_FindPkMappingsByTaprootAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var taprootAddresses []string
		fill_err = tp.Fill(&taprootAddresses)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.FindPkMappingsByTaprootAddress(ctx, taprootAddresses)
	})
}

func Fuzz_Nosy_DBClient_FindUnprocessableMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.FindUnprocessableMessages(ctx)
	})
}

func Fuzz_Nosy_DBClient_GetLatestPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var symbol string
		fill_err = tp.Fill(&symbol)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetLatestPrice(ctx, symbol)
	})
}

func Fuzz_Nosy_DBClient_InsertPkAddressMappings__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		var taproot string
		fill_err = tp.Fill(&taproot)
		if fill_err != nil {
			return
		}
		var nativeSigwitOdd string
		fill_err = tp.Fill(&nativeSigwitOdd)
		if fill_err != nil {
			return
		}
		var nativeSigwitEven string
		fill_err = tp.Fill(&nativeSigwitEven)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.InsertPkAddressMappings(ctx, stakerPkHex, taproot, nativeSigwitOdd, nativeSigwitEven)
	})
}

func Fuzz_Nosy_DBClient_Ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Ping(ctx)
	})
}

func Fuzz_Nosy_DBClient_SaveUnprocessableMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var messageBody string
		fill_err = tp.Fill(&messageBody)
		if fill_err != nil {
			return
		}
		var receipt string
		fill_err = tp.Fill(&receipt)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SaveUnprocessableMessage(ctx, messageBody, receipt)
	})
}

func Fuzz_Nosy_DBClient_SetLatestPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var symbol string
		fill_err = tp.Fill(&symbol)
		if fill_err != nil {
			return
		}
		var price float64
		fill_err = tp.Fill(&price)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SetLatestPrice(ctx, symbol, price)
	})
}

func Fuzz_Nosy_IndexerDBClient_CheckDelegationExistByStakerPk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *IndexerDBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var address string
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var extraFilter *indexerdbclient.DelegationFilter
		fill_err = tp.Fill(&extraFilter)
		if fill_err != nil {
			return
		}
		if _m == nil || extraFilter == nil {
			return
		}

		_m.CheckDelegationExistByStakerPk(ctx, address, extraFilter)
	})
}

func Fuzz_Nosy_IndexerDBClient_GetBbnStakingParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *IndexerDBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetBbnStakingParams(ctx)
	})
}

func Fuzz_Nosy_IndexerDBClient_GetBtcCheckpointParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *IndexerDBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetBtcCheckpointParams(ctx)
	})
}

func Fuzz_Nosy_IndexerDBClient_GetDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *IndexerDBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetDelegation(ctx, stakingTxHashHex)
	})
}

func Fuzz_Nosy_IndexerDBClient_GetDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *IndexerDBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakerPKHex string
		fill_err = tp.Fill(&stakerPKHex)
		if fill_err != nil {
			return
		}
		var paginationToken string
		fill_err = tp.Fill(&paginationToken)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetDelegations(ctx, stakerPKHex, paginationToken)
	})
}

func Fuzz_Nosy_IndexerDBClient_GetFinalityProviderByPk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *IndexerDBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fpPk string
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetFinalityProviderByPk(ctx, fpPk)
	})
}

func Fuzz_Nosy_IndexerDBClient_GetFinalityProviders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *IndexerDBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetFinalityProviders(ctx)
	})
}

func Fuzz_Nosy_IndexerDBClient_GetLastProcessedBbnHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *IndexerDBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetLastProcessedBbnHeight(ctx)
	})
}

func Fuzz_Nosy_IndexerDBClient_Ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *IndexerDBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Ping(ctx)
	})
}

func Fuzz_Nosy_OrdinalsClient_FetchUTXOInfos__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *OrdinalsClient
		fill_err = tp.Fill(&_m)
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
		if _m == nil {
			return
		}

		_m.FetchUTXOInfos(ctx, utxos)
	})
}

func Fuzz_Nosy_OrdinalsClient_GetBaseURL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *OrdinalsClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetBaseURL()
	})
}

func Fuzz_Nosy_OrdinalsClient_GetDefaultRequestTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *OrdinalsClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetDefaultRequestTimeout()
	})
}

func Fuzz_Nosy_OrdinalsClient_GetHttpClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *OrdinalsClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetHttpClient()
	})
}

func Fuzz_Nosy_V1DBClient_CheckDelegationExistByStakerPk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var address string
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var extraFilter *v1dbclient.DelegationFilter
		fill_err = tp.Fill(&extraFilter)
		if fill_err != nil {
			return
		}
		if _m == nil || extraFilter == nil {
			return
		}

		_m.CheckDelegationExistByStakerPk(ctx, address, extraFilter)
	})
}

// skipping Fuzz_Nosy_V1DBClient_DeleteUnprocessableMessage__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_V1DBClient_FindDelegationByTxHashHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var txHashHex string
		fill_err = tp.Fill(&txHashHex)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.FindDelegationByTxHashHex(ctx, txHashHex)
	})
}

func Fuzz_Nosy_V1DBClient_FindDelegationsByStakerPk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakerPk string
		fill_err = tp.Fill(&stakerPk)
		if fill_err != nil {
			return
		}
		var extraFilter *v1dbclient.DelegationFilter
		fill_err = tp.Fill(&extraFilter)
		if fill_err != nil {
			return
		}
		var paginationToken string
		fill_err = tp.Fill(&paginationToken)
		if fill_err != nil {
			return
		}
		if _m == nil || extraFilter == nil {
			return
		}

		_m.FindDelegationsByStakerPk(ctx, stakerPk, extraFilter, paginationToken)
	})
}

func Fuzz_Nosy_V1DBClient_FindFinalityProviderStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var paginationToken string
		fill_err = tp.Fill(&paginationToken)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.FindFinalityProviderStats(ctx, paginationToken)
	})
}

func Fuzz_Nosy_V1DBClient_FindFinalityProviderStatsByFinalityProviderPkHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var finalityProviderPkHex []string
		fill_err = tp.Fill(&finalityProviderPkHex)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.FindFinalityProviderStatsByFinalityProviderPkHex(ctx, finalityProviderPkHex)
	})
}

func Fuzz_Nosy_V1DBClient_FindPkMappingsByNativeSegwitAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var nativeSegwitAddresses []string
		fill_err = tp.Fill(&nativeSegwitAddresses)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.FindPkMappingsByNativeSegwitAddress(ctx, nativeSegwitAddresses)
	})
}

func Fuzz_Nosy_V1DBClient_FindPkMappingsByTaprootAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var taprootAddresses []string
		fill_err = tp.Fill(&taprootAddresses)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.FindPkMappingsByTaprootAddress(ctx, taprootAddresses)
	})
}

func Fuzz_Nosy_V1DBClient_FindTopStakersByTvl__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var paginationToken string
		fill_err = tp.Fill(&paginationToken)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.FindTopStakersByTvl(ctx, paginationToken)
	})
}

func Fuzz_Nosy_V1DBClient_FindUnprocessableMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.FindUnprocessableMessages(ctx)
	})
}

func Fuzz_Nosy_V1DBClient_GetLatestBtcInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetLatestBtcInfo(ctx)
	})
}

func Fuzz_Nosy_V1DBClient_GetLatestPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var symbol string
		fill_err = tp.Fill(&symbol)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetLatestPrice(ctx, symbol)
	})
}

func Fuzz_Nosy_V1DBClient_GetOrCreateStatsLock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var state string
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetOrCreateStatsLock(ctx, stakingTxHashHex, state)
	})
}

func Fuzz_Nosy_V1DBClient_GetOverallStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetOverallStats(ctx)
	})
}

func Fuzz_Nosy_V1DBClient_GetStakerStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetStakerStats(ctx, stakerPkHex)
	})
}

func Fuzz_Nosy_V1DBClient_IncrementFinalityProviderStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var fpPkHex string
		fill_err = tp.Fill(&fpPkHex)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.IncrementFinalityProviderStats(ctx, stakingTxHashHex, fpPkHex, amount)
	})
}

func Fuzz_Nosy_V1DBClient_IncrementOverallStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.IncrementOverallStats(ctx, stakingTxHashHex, stakerPkHex, amount)
	})
}

func Fuzz_Nosy_V1DBClient_IncrementStakerStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.IncrementStakerStats(ctx, stakingTxHashHex, stakerPkHex, amount)
	})
}

func Fuzz_Nosy_V1DBClient_InsertPkAddressMappings__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		var taproot string
		fill_err = tp.Fill(&taproot)
		if fill_err != nil {
			return
		}
		var nativeSigwitOdd string
		fill_err = tp.Fill(&nativeSigwitOdd)
		if fill_err != nil {
			return
		}
		var nativeSigwitEven string
		fill_err = tp.Fill(&nativeSigwitEven)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.InsertPkAddressMappings(ctx, stakerPkHex, taproot, nativeSigwitOdd, nativeSigwitEven)
	})
}

func Fuzz_Nosy_V1DBClient_Ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Ping(ctx)
	})
}

func Fuzz_Nosy_V1DBClient_SaveActiveStakingDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		var fpPkHex string
		fill_err = tp.Fill(&fpPkHex)
		if fill_err != nil {
			return
		}
		var stakingTxHex string
		fill_err = tp.Fill(&stakingTxHex)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var timelock uint64
		fill_err = tp.Fill(&timelock)
		if fill_err != nil {
			return
		}
		var outputIndex uint64
		fill_err = tp.Fill(&outputIndex)
		if fill_err != nil {
			return
		}
		var startTimestamp int64
		fill_err = tp.Fill(&startTimestamp)
		if fill_err != nil {
			return
		}
		var isOverflow bool
		fill_err = tp.Fill(&isOverflow)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SaveActiveStakingDelegation(ctx, stakingTxHashHex, stakerPkHex, fpPkHex, stakingTxHex, amount, startHeight, timelock, outputIndex, startTimestamp, isOverflow)
	})
}

func Fuzz_Nosy_V1DBClient_SaveTimeLockExpireCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var expireHeight uint64
		fill_err = tp.Fill(&expireHeight)
		if fill_err != nil {
			return
		}
		var txType string
		fill_err = tp.Fill(&txType)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SaveTimeLockExpireCheck(ctx, stakingTxHashHex, expireHeight, txType)
	})
}

func Fuzz_Nosy_V1DBClient_SaveUnbondingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var unbondingTxHashHex string
		fill_err = tp.Fill(&unbondingTxHashHex)
		if fill_err != nil {
			return
		}
		var txHex string
		fill_err = tp.Fill(&txHex)
		if fill_err != nil {
			return
		}
		var signatureHex string
		fill_err = tp.Fill(&signatureHex)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SaveUnbondingTx(ctx, stakingTxHashHex, unbondingTxHashHex, txHex, signatureHex)
	})
}

func Fuzz_Nosy_V1DBClient_SaveUnprocessableMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var messageBody string
		fill_err = tp.Fill(&messageBody)
		if fill_err != nil {
			return
		}
		var receipt string
		fill_err = tp.Fill(&receipt)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SaveUnprocessableMessage(ctx, messageBody, receipt)
	})
}

func Fuzz_Nosy_V1DBClient_ScanDelegationsPaginated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var paginationToken string
		fill_err = tp.Fill(&paginationToken)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.ScanDelegationsPaginated(ctx, paginationToken)
	})
}

func Fuzz_Nosy_V1DBClient_SetLatestPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var symbol string
		fill_err = tp.Fill(&symbol)
		if fill_err != nil {
			return
		}
		var price float64
		fill_err = tp.Fill(&price)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SetLatestPrice(ctx, symbol, price)
	})
}

func Fuzz_Nosy_V1DBClient_SubtractFinalityProviderStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var fpPkHex string
		fill_err = tp.Fill(&fpPkHex)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SubtractFinalityProviderStats(ctx, stakingTxHashHex, fpPkHex, amount)
	})
}

func Fuzz_Nosy_V1DBClient_SubtractOverallStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SubtractOverallStats(ctx, stakingTxHashHex, stakerPkHex, amount)
	})
}

func Fuzz_Nosy_V1DBClient_SubtractStakerStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SubtractStakerStats(ctx, stakingTxHashHex, stakerPkHex, amount)
	})
}

func Fuzz_Nosy_V1DBClient_TransitionToTransitionedState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.TransitionToTransitionedState(ctx, stakingTxHashHex)
	})
}

func Fuzz_Nosy_V1DBClient_TransitionToUnbondedState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var eligiblePreviousState []types.DelegationState
		fill_err = tp.Fill(&eligiblePreviousState)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.TransitionToUnbondedState(ctx, stakingTxHashHex, eligiblePreviousState)
	})
}

func Fuzz_Nosy_V1DBClient_TransitionToUnbondingState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var txHashHex string
		fill_err = tp.Fill(&txHashHex)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var timelock uint64
		fill_err = tp.Fill(&timelock)
		if fill_err != nil {
			return
		}
		var outputIndex uint64
		fill_err = tp.Fill(&outputIndex)
		if fill_err != nil {
			return
		}
		var txHex string
		fill_err = tp.Fill(&txHex)
		if fill_err != nil {
			return
		}
		var startTimestamp int64
		fill_err = tp.Fill(&startTimestamp)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.TransitionToUnbondingState(ctx, txHashHex, startHeight, timelock, outputIndex, txHex, startTimestamp)
	})
}

func Fuzz_Nosy_V1DBClient_TransitionToWithdrawnState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var txHashHex string
		fill_err = tp.Fill(&txHashHex)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.TransitionToWithdrawnState(ctx, txHashHex)
	})
}

func Fuzz_Nosy_V1DBClient_UpsertLatestBtcInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V1DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var confirmedTvl uint64
		fill_err = tp.Fill(&confirmedTvl)
		if fill_err != nil {
			return
		}
		var unconfirmedTvl uint64
		fill_err = tp.Fill(&unconfirmedTvl)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.UpsertLatestBtcInfo(ctx, height, confirmedTvl, unconfirmedTvl)
	})
}

// skipping Fuzz_Nosy_V2DBClient_DeleteUnprocessableMessage__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_V2DBClient_FindPkMappingsByNativeSegwitAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var nativeSegwitAddresses []string
		fill_err = tp.Fill(&nativeSegwitAddresses)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.FindPkMappingsByNativeSegwitAddress(ctx, nativeSegwitAddresses)
	})
}

func Fuzz_Nosy_V2DBClient_FindPkMappingsByTaprootAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var taprootAddresses []string
		fill_err = tp.Fill(&taprootAddresses)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.FindPkMappingsByTaprootAddress(ctx, taprootAddresses)
	})
}

func Fuzz_Nosy_V2DBClient_FindUnprocessableMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.FindUnprocessableMessages(ctx)
	})
}

func Fuzz_Nosy_V2DBClient_GetActiveStakersCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetActiveStakersCount(ctx)
	})
}

func Fuzz_Nosy_V2DBClient_GetFinalityProviderStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetFinalityProviderStats(ctx)
	})
}

func Fuzz_Nosy_V2DBClient_GetLatestPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var symbol string
		fill_err = tp.Fill(&symbol)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetLatestPrice(ctx, symbol)
	})
}

func Fuzz_Nosy_V2DBClient_GetOrCreateStatsLock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var state string
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetOrCreateStatsLock(ctx, stakingTxHashHex, state)
	})
}

func Fuzz_Nosy_V2DBClient_GetOverallStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetOverallStats(ctx)
	})
}

func Fuzz_Nosy_V2DBClient_GetStakerStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakerPKHex string
		fill_err = tp.Fill(&stakerPKHex)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetStakerStats(ctx, stakerPKHex)
	})
}

func Fuzz_Nosy_V2DBClient_HandleActiveStakerStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.HandleActiveStakerStats(ctx, stakingTxHashHex, stakerPkHex, amount)
	})
}

func Fuzz_Nosy_V2DBClient_HandleUnbondingStakerStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var stateHistory []string
		fill_err = tp.Fill(&stateHistory)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.HandleUnbondingStakerStats(ctx, stakingTxHashHex, stakerPkHex, amount, stateHistory)
	})
}

func Fuzz_Nosy_V2DBClient_HandleWithdrawableStakerStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var stateHistory []string
		fill_err = tp.Fill(&stateHistory)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.HandleWithdrawableStakerStats(ctx, stakingTxHashHex, stakerPkHex, amount, stateHistory)
	})
}

func Fuzz_Nosy_V2DBClient_HandleWithdrawnStakerStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var stateHistory []string
		fill_err = tp.Fill(&stateHistory)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.HandleWithdrawnStakerStats(ctx, stakingTxHashHex, stakerPkHex, amount, stateHistory)
	})
}

func Fuzz_Nosy_V2DBClient_IncrementFinalityProviderStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var fpPkHexes []string
		fill_err = tp.Fill(&fpPkHexes)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.IncrementFinalityProviderStats(ctx, stakingTxHashHex, fpPkHexes, amount)
	})
}

func Fuzz_Nosy_V2DBClient_IncrementOverallStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.IncrementOverallStats(ctx, stakingTxHashHex, amount)
	})
}

func Fuzz_Nosy_V2DBClient_InsertPkAddressMappings__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		var taproot string
		fill_err = tp.Fill(&taproot)
		if fill_err != nil {
			return
		}
		var nativeSigwitOdd string
		fill_err = tp.Fill(&nativeSigwitOdd)
		if fill_err != nil {
			return
		}
		var nativeSigwitEven string
		fill_err = tp.Fill(&nativeSigwitEven)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.InsertPkAddressMappings(ctx, stakerPkHex, taproot, nativeSigwitOdd, nativeSigwitEven)
	})
}

func Fuzz_Nosy_V2DBClient_Ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Ping(ctx)
	})
}

func Fuzz_Nosy_V2DBClient_SaveUnprocessableMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var messageBody string
		fill_err = tp.Fill(&messageBody)
		if fill_err != nil {
			return
		}
		var receipt string
		fill_err = tp.Fill(&receipt)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SaveUnprocessableMessage(ctx, messageBody, receipt)
	})
}

func Fuzz_Nosy_V2DBClient_SetLatestPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var symbol string
		fill_err = tp.Fill(&symbol)
		if fill_err != nil {
			return
		}
		var price float64
		fill_err = tp.Fill(&price)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SetLatestPrice(ctx, symbol, price)
	})
}

func Fuzz_Nosy_V2DBClient_SubtractFinalityProviderStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var fpPkHexes []string
		fill_err = tp.Fill(&fpPkHexes)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SubtractFinalityProviderStats(ctx, stakingTxHashHex, fpPkHexes, amount)
	})
}

func Fuzz_Nosy_V2DBClient_SubtractOverallStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *V2DBClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SubtractOverallStats(ctx, stakingTxHashHex, amount)
	})
}
