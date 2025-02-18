package indexerdbclient

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

func Fuzz_Nosy_IndexerDatabase_CheckDelegationExistByStakerPk__(f *testing.F) {
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

		i1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		i1.CheckDelegationExistByStakerPk(c4, stakerPk, extraFilter)
	})
}

func Fuzz_Nosy_IndexerDatabase_GetBbnStakingParams__(f *testing.F) {
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

		db, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		db.GetBbnStakingParams(c4)
	})
}

func Fuzz_Nosy_IndexerDatabase_GetBtcCheckpointParams__(f *testing.F) {
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

		db, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		db.GetBtcCheckpointParams(c4)
	})
}

func Fuzz_Nosy_IndexerDatabase_GetDelegation__(f *testing.F) {
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

		i1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		i1.GetDelegation(c4, stakingTxHashHex)
	})
}

func Fuzz_Nosy_IndexerDatabase_GetDelegations__(f *testing.F) {
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
		if client == nil || cfg == nil {
			return
		}

		i1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		i1.GetDelegations(c4, stakerPKHex, paginationToken)
	})
}

func Fuzz_Nosy_IndexerDatabase_GetFinalityProviderByPk__(f *testing.F) {
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
		var fpPk string
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if client == nil || cfg == nil {
			return
		}

		i1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		i1.GetFinalityProviderByPk(c4, fpPk)
	})
}

func Fuzz_Nosy_IndexerDatabase_GetFinalityProviders__(f *testing.F) {
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

		i1, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		i1.GetFinalityProviders(c4)
	})
}

func Fuzz_Nosy_IndexerDatabase_GetLastProcessedBbnHeight__(f *testing.F) {
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

		db, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		db.GetLastProcessedBbnHeight(c4)
	})
}

// skipping Fuzz_Nosy_IndexerDBClient_CheckDelegationExistByStakerPk__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/indexer/db/client.IndexerDBClient

// skipping Fuzz_Nosy_IndexerDBClient_GetBbnStakingParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/indexer/db/client.IndexerDBClient

// skipping Fuzz_Nosy_IndexerDBClient_GetBtcCheckpointParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/indexer/db/client.IndexerDBClient

// skipping Fuzz_Nosy_IndexerDBClient_GetDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/indexer/db/client.IndexerDBClient

// skipping Fuzz_Nosy_IndexerDBClient_GetDelegations__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/indexer/db/client.IndexerDBClient

// skipping Fuzz_Nosy_IndexerDBClient_GetFinalityProviderByPk__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/indexer/db/client.IndexerDBClient

// skipping Fuzz_Nosy_IndexerDBClient_GetFinalityProviders__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/indexer/db/client.IndexerDBClient

// skipping Fuzz_Nosy_IndexerDBClient_GetLastProcessedBbnHeight__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/indexer/db/client.IndexerDBClient

// skipping Fuzz_Nosy_IndexerDBClient_Ping__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/indexer/db/client.IndexerDBClient
