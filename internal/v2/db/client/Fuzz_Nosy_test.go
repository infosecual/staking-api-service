package v2dbclient

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/staking-api-service/internal/shared/config"
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

func Fuzz_Nosy_V2Database_GetActiveStakersCount__(f *testing.F) {
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
		v1.GetActiveStakersCount(c4)
	})
}

func Fuzz_Nosy_V2Database_GetFinalityProviderStats__(f *testing.F) {
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
		v1.GetFinalityProviderStats(c4)
	})
}

func Fuzz_Nosy_V2Database_GetOrCreateStatsLock__(f *testing.F) {
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

func Fuzz_Nosy_V2Database_GetOverallStats__(f *testing.F) {
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

func Fuzz_Nosy_V2Database_GetStakerStats__(f *testing.F) {
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

func Fuzz_Nosy_V2Database_HandleActiveStakerStats__(f *testing.F) {
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
		v1.HandleActiveStakerStats(c4, stakingTxHashHex, stakerPkHex, amount)
	})
}

func Fuzz_Nosy_V2Database_HandleUnbondingStakerStats__(f *testing.F) {
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
		var stateHistory []string
		fill_err = tp.Fill(&stateHistory)
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
		v1.HandleUnbondingStakerStats(c4, stakingTxHashHex, stakerPkHex, amount, stateHistory)
	})
}

func Fuzz_Nosy_V2Database_HandleWithdrawableStakerStats__(f *testing.F) {
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
		var stateHistory []string
		fill_err = tp.Fill(&stateHistory)
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
		v1.HandleWithdrawableStakerStats(c4, stakingTxHashHex, stakerPkHex, amount, stateHistory)
	})
}

func Fuzz_Nosy_V2Database_HandleWithdrawnStakerStats__(f *testing.F) {
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
		var stateHistory []string
		fill_err = tp.Fill(&stateHistory)
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
		v1.HandleWithdrawnStakerStats(c4, stakingTxHashHex, stakerPkHex, amount, stateHistory)
	})
}

func Fuzz_Nosy_V2Database_IncrementFinalityProviderStats__(f *testing.F) {
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
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.IncrementFinalityProviderStats(c4, stakingTxHashHex, fpPkHexes, amount)
	})
}

func Fuzz_Nosy_V2Database_IncrementOverallStats__(f *testing.F) {
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
		v1.IncrementOverallStats(c4, stakingTxHashHex, amount)
	})
}

func Fuzz_Nosy_V2Database_SubtractFinalityProviderStats__(f *testing.F) {
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
		if client == nil || cfg == nil {
			return
		}

		v1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		v1.SubtractFinalityProviderStats(c4, stakingTxHashHex, fpPkHexes, amount)
	})
}

func Fuzz_Nosy_V2Database_SubtractOverallStats__(f *testing.F) {
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
		v1.SubtractOverallStats(c4, stakingTxHashHex, amount)
	})
}

func Fuzz_Nosy_V2Database_generateOverallStatsId__(f *testing.F) {
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

// skipping Fuzz_Nosy_V2Database_updateFinalityProviderStats__ because parameters include func, chan, or unsupported interface: []go.mongodb.org/mongo-driver/mongo.WriteModel

// skipping Fuzz_Nosy_V2Database_updateStakerStats__ because parameters include func, chan, or unsupported interface: go.mongodb.org/mongo-driver/bson/primitive.M

func Fuzz_Nosy_V2Database_updateStatsLockByFieldName__(f *testing.F) {
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

// skipping Fuzz_Nosy_V2DBClient_GetActiveStakersCount__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/db/client.V2DBClient

// skipping Fuzz_Nosy_V2DBClient_GetFinalityProviderStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/db/client.V2DBClient

// skipping Fuzz_Nosy_V2DBClient_GetOrCreateStatsLock__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/db/client.V2DBClient

// skipping Fuzz_Nosy_V2DBClient_GetOverallStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/db/client.V2DBClient

// skipping Fuzz_Nosy_V2DBClient_GetStakerStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/db/client.V2DBClient

// skipping Fuzz_Nosy_V2DBClient_HandleActiveStakerStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/db/client.V2DBClient

// skipping Fuzz_Nosy_V2DBClient_HandleUnbondingStakerStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/db/client.V2DBClient

// skipping Fuzz_Nosy_V2DBClient_HandleWithdrawableStakerStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/db/client.V2DBClient

// skipping Fuzz_Nosy_V2DBClient_HandleWithdrawnStakerStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/db/client.V2DBClient

// skipping Fuzz_Nosy_V2DBClient_IncrementFinalityProviderStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/db/client.V2DBClient

// skipping Fuzz_Nosy_V2DBClient_IncrementOverallStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/db/client.V2DBClient

// skipping Fuzz_Nosy_V2DBClient_SubtractFinalityProviderStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/db/client.V2DBClient

// skipping Fuzz_Nosy_V2DBClient_SubtractOverallStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/db/client.V2DBClient

func Fuzz_Nosy_constructStatsLockId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHashHex string, state string) {
		constructStatsLockId(stakingTxHashHex, state)
	})
}
