package v1dbclient

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/staking-api-service/internal/shared/config"
	"github.com/babylonlabs-io/staking-api-service/internal/shared/types"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"go.mongodb.org/mongo-driver/mongo"
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

func Fuzz_Nosy_V1Database_CheckDelegationExistByStakerPk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var stakerPk string
		fill_err = tp.Fill(&stakerPk)
		if fill_err != nil {
			return
		}
		var extraFilter *DelegationFilter
		fill_err = tp.Fill(&extraFilter)
		if fill_err != nil {
			return
		}
		if client == nil || cfg == nil || extraFilter == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.CheckDelegationExistByStakerPk(c4, stakerPk, extraFilter)
	})
}

func Fuzz_Nosy_V1Database_FindDelegationByTxHashHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.FindDelegationByTxHashHex(c4, stakingTxHashHex)
	})
}

func Fuzz_Nosy_V1Database_FindDelegationsByStakerPk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var stakerPk string
		fill_err = tp.Fill(&stakerPk)
		if fill_err != nil {
			return
		}
		var extraFilter *DelegationFilter
		fill_err = tp.Fill(&extraFilter)
		if fill_err != nil {
			return
		}
		var paginationToken string
		fill_err = tp.Fill(&paginationToken)
		if fill_err != nil {
			return
		}
		if client == nil || cfg == nil || extraFilter == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.FindDelegationsByStakerPk(c4, stakerPk, extraFilter, paginationToken)
	})
}

func Fuzz_Nosy_V1Database_FindFinalityProviderStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var paginationToken string
		fill_err = tp.Fill(&paginationToken)
		if fill_err != nil {
			return
		}
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.FindFinalityProviderStats(c4, paginationToken)
	})
}

func Fuzz_Nosy_V1Database_FindFinalityProviderStatsByFinalityProviderPkHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var finalityProviderPkHex []string
		fill_err = tp.Fill(&finalityProviderPkHex)
		if fill_err != nil {
			return
		}
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.FindFinalityProviderStatsByFinalityProviderPkHex(c4, finalityProviderPkHex)
	})
}

func Fuzz_Nosy_V1Database_FindTopStakersByTvl__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var paginationToken string
		fill_err = tp.Fill(&paginationToken)
		if fill_err != nil {
			return
		}
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.FindTopStakersByTvl(c4, paginationToken)
	})
}

func Fuzz_Nosy_V1Database_GetLatestBtcInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.GetLatestBtcInfo(c4)
	})
}

func Fuzz_Nosy_V1Database_GetOrCreateStatsLock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var txType string
		fill_err = tp.Fill(&txType)
		if fill_err != nil {
			return
		}
		if client == nil || cfg == nil {
			return
		}

		db, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		db.GetOrCreateStatsLock(c4, stakingTxHashHex, txType)
	})
}

func Fuzz_Nosy_V1Database_GetOverallStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.GetOverallStats(c4)
	})
}

func Fuzz_Nosy_V1Database_GetStakerStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.GetStakerStats(c4, stakerPkHex)
	})
}

func Fuzz_Nosy_V1Database_IncrementFinalityProviderStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
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
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.IncrementFinalityProviderStats(c4, stakingTxHashHex, fpPkHex, amount)
	})
}

func Fuzz_Nosy_V1Database_IncrementOverallStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
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
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.IncrementOverallStats(c4, stakingTxHashHex, stakerPkHex, amount)
	})
}

func Fuzz_Nosy_V1Database_IncrementStakerStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
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
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.IncrementStakerStats(c4, stakingTxHashHex, stakerPkHex, amount)
	})
}

func Fuzz_Nosy_V1Database_SaveActiveStakingDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
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
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.SaveActiveStakingDelegation(c4, stakingTxHashHex, stakerPkHex, fpPkHex, stakingTxHex, amount, startHeight, timelock, outputIndex, startTimestamp, isOverflow)
	})
}

func Fuzz_Nosy_V1Database_SaveTimeLockExpireCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
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
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.SaveTimeLockExpireCheck(c4, stakingTxHashHex, expireHeight, txType)
	})
}

func Fuzz_Nosy_V1Database_SaveUnbondingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var txHashHex string
		fill_err = tp.Fill(&txHashHex)
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
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.SaveUnbondingTx(c4, stakingTxHashHex, txHashHex, txHex, signatureHex)
	})
}

func Fuzz_Nosy_V1Database_ScanDelegationsPaginated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var paginationToken string
		fill_err = tp.Fill(&paginationToken)
		if fill_err != nil {
			return
		}
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.ScanDelegationsPaginated(c4, paginationToken)
	})
}

func Fuzz_Nosy_V1Database_SubtractFinalityProviderStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
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
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.SubtractFinalityProviderStats(c4, stakingTxHashHex, fpPkHex, amount)
	})
}

func Fuzz_Nosy_V1Database_SubtractOverallStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
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
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.SubtractOverallStats(c4, stakingTxHashHex, stakerPkHex, amount)
	})
}

func Fuzz_Nosy_V1Database_SubtractStakerStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
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
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.SubtractStakerStats(c4, stakingTxHashHex, stakerPkHex, amount)
	})
}

func Fuzz_Nosy_V1Database_TransitionToTransitionedState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.TransitionToTransitionedState(c4, stakingTxHashHex)
	})
}

func Fuzz_Nosy_V1Database_TransitionToUnbondedState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
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
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.TransitionToUnbondedState(c4, stakingTxHashHex, eligiblePreviousState)
	})
}

func Fuzz_Nosy_V1Database_TransitionToUnbondingState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
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
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.TransitionToUnbondingState(c4, txHashHex, startHeight, timelock, outputIndex, txHex, startTimestamp)
	})
}

func Fuzz_Nosy_V1Database_TransitionToWithdrawnState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var txHashHex string
		fill_err = tp.Fill(&txHashHex)
		if fill_err != nil {
			return
		}
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.TransitionToWithdrawnState(c4, txHashHex)
	})
}

func Fuzz_Nosy_V1Database_UpsertLatestBtcInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
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
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.UpsertLatestBtcInfo(c4, height, confirmedTvl, unconfirmedTvl)
	})
}

func Fuzz_Nosy_V1Database_generateOverallStatsId__(f *testing.F) {
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
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(ctx, client, cfg)
		if err != nil {
			return
		}
		v1.generateOverallStatsId()
	})
}

// skipping Fuzz_Nosy_V1Database_transitionState__ because parameters include func, chan, or unsupported interface: map[string]interface{}

// skipping Fuzz_Nosy_V1Database_updateFinalityProviderStats__ because parameters include func, chan, or unsupported interface: go.mongodb.org/mongo-driver/bson/primitive.M

// skipping Fuzz_Nosy_V1Database_updateStakerStats__ because parameters include func, chan, or unsupported interface: go.mongodb.org/mongo-driver/bson/primitive.M

func Fuzz_Nosy_V1Database_updateStatsLockByFieldName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var client *mongo.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
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
		var fieldName string
		fill_err = tp.Fill(&fieldName)
		if fill_err != nil {
			return
		}
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.updateStatsLockByFieldName(c4, stakingTxHashHex, state, fieldName)
	})
}

// skipping Fuzz_Nosy_V1DBClient_CheckDelegationExistByStakerPk__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_FindDelegationByTxHashHex__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_FindDelegationsByStakerPk__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_FindFinalityProviderStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_FindFinalityProviderStatsByFinalityProviderPkHex__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_FindTopStakersByTvl__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_GetLatestBtcInfo__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_GetOrCreateStatsLock__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_GetOverallStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_GetStakerStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_IncrementFinalityProviderStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_IncrementOverallStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_IncrementStakerStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_SaveActiveStakingDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_SaveTimeLockExpireCheck__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_SaveUnbondingTx__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_ScanDelegationsPaginated__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_SubtractFinalityProviderStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_SubtractOverallStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_SubtractStakerStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_TransitionToTransitionedState__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_TransitionToUnbondedState__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_TransitionToUnbondingState__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_TransitionToWithdrawnState__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

// skipping Fuzz_Nosy_V1DBClient_UpsertLatestBtcInfo__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/db/client.V1DBClient

func Fuzz_Nosy_constructStatsLockId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHashHex string, state string) {
		constructStatsLockId(stakingTxHashHex, state)
	})
}
