package v2service

import (
	"context"
	"testing"

	indexerdbmodel "github.com/babylonlabs-io/staking-api-service/internal/indexer/db/model"
	"github.com/babylonlabs-io/staking-api-service/internal/shared/config"
	dbclients "github.com/babylonlabs-io/staking-api-service/internal/shared/db/clients"
	dbmodel "github.com/babylonlabs-io/staking-api-service/internal/shared/db/model"
	"github.com/babylonlabs-io/staking-api-service/internal/shared/http/clients"
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

func Fuzz_Nosy_V2Service_GetDelegation__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.GetDelegation(c5, stakingTxHashHex)
	})
}

func Fuzz_Nosy_V2Service_GetDelegations__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		var paginationKey string
		fill_err = tp.Fill(&paginationKey)
		if fill_err != nil {
			return
		}
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.GetDelegations(c5, stakerPkHex, paginationKey)
	})
}

func Fuzz_Nosy_V2Service_GetFinalityProvidersWithStats__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
		if fill_err != nil {
			return
		}
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.GetFinalityProvidersWithStats(c5)
	})
}

func Fuzz_Nosy_V2Service_GetLatestPrices__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
		if fill_err != nil {
			return
		}
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.GetLatestPrices(c5)
	})
}

func Fuzz_Nosy_V2Service_GetNetworkInfo__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
		if fill_err != nil {
			return
		}
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.GetNetworkInfo(c5)
	})
}

func Fuzz_Nosy_V2Service_GetOverallStats__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
		if fill_err != nil {
			return
		}
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.GetOverallStats(c5)
	})
}

func Fuzz_Nosy_V2Service_GetParams__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
		if fill_err != nil {
			return
		}
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.GetParams(c5)
	})
}

func Fuzz_Nosy_V2Service_GetStakerPublicKeysByAddresses__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
		if fill_err != nil {
			return
		}
		var addresses []string
		fill_err = tp.Fill(&addresses)
		if fill_err != nil {
			return
		}
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.GetStakerPublicKeysByAddresses(c5, addresses)
	})
}

func Fuzz_Nosy_V2Service_GetStakerStats__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
		if fill_err != nil {
			return
		}
		var stakerPKHex string
		fill_err = tp.Fill(&stakerPKHex)
		if fill_err != nil {
			return
		}
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.GetStakerStats(c5, stakerPKHex)
	})
}

func Fuzz_Nosy_V2Service_MarkV1DelegationAsTransitioned__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.MarkV1DelegationAsTransitioned(c5, stakingTxHashHex)
	})
}

func Fuzz_Nosy_V2Service_ProcessActiveDelegationStats__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
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
		var fpBtcPkHexes []string
		fill_err = tp.Fill(&fpBtcPkHexes)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.ProcessActiveDelegationStats(c5, stakingTxHashHex, stakerPkHex, fpBtcPkHexes, amount)
	})
}

func Fuzz_Nosy_V2Service_ProcessAndSaveBtcAddresses__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.ProcessAndSaveBtcAddresses(c5, stakerPkHex)
	})
}

func Fuzz_Nosy_V2Service_ProcessUnbondingDelegationStats__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
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
		var fpBtcPkHexes []string
		fill_err = tp.Fill(&fpBtcPkHexes)
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
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.ProcessUnbondingDelegationStats(c5, stakingTxHashHex, stakerPkHex, fpBtcPkHexes, amount, stateHistory)
	})
}

func Fuzz_Nosy_V2Service_ProcessWithdrawableDelegationStats__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
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
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.ProcessWithdrawableDelegationStats(c5, stakingTxHashHex, stakerPkHex, amount, stateHistory)
	})
}

func Fuzz_Nosy_V2Service_ProcessWithdrawnDelegationStats__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
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
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.ProcessWithdrawnDelegationStats(c5, stakingTxHashHex, stakerPkHex, amount, stateHistory)
	})
}

func Fuzz_Nosy_V2Service_SaveUnprocessableMessages__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
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
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.SaveUnprocessableMessages(c5, messageBody, receipt)
	})
}

func Fuzz_Nosy_V2Service_doGetLatestBTCPrice__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(ctx, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.doGetLatestBTCPrice()
	})
}

func Fuzz_Nosy_V2Service_getBbnStakingParams__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
		if fill_err != nil {
			return
		}
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.getBbnStakingParams(c5)
	})
}

func Fuzz_Nosy_V2Service_getBtcCheckpointParams__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
		if fill_err != nil {
			return
		}
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.getBtcCheckpointParams(c5)
	})
}

func Fuzz_Nosy_V2Service_getLatestBTCPrice__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var clients *clients.Clients
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}
		var dbClients *dbclients.DbClients
		fill_err = tp.Fill(&dbClients)
		if fill_err != nil {
			return
		}
		var c5 context.Context
		fill_err = tp.Fill(&c5)
		if fill_err != nil {
			return
		}
		if cfg == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, clients, dbClients)
		if err != nil {
			return
		}
		s.getLatestBTCPrice(c5)
	})
}

// skipping Fuzz_Nosy_V2ServiceProvider_GetDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/service.V2ServiceProvider

// skipping Fuzz_Nosy_V2ServiceProvider_GetDelegations__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/service.V2ServiceProvider

// skipping Fuzz_Nosy_V2ServiceProvider_GetFinalityProvidersWithStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/service.V2ServiceProvider

// skipping Fuzz_Nosy_V2ServiceProvider_GetLatestPrices__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/service.V2ServiceProvider

// skipping Fuzz_Nosy_V2ServiceProvider_GetNetworkInfo__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/service.V2ServiceProvider

// skipping Fuzz_Nosy_V2ServiceProvider_GetOverallStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/service.V2ServiceProvider

// skipping Fuzz_Nosy_V2ServiceProvider_GetStakerStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/service.V2ServiceProvider

// skipping Fuzz_Nosy_V2ServiceProvider_MarkV1DelegationAsTransitioned__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/service.V2ServiceProvider

// skipping Fuzz_Nosy_V2ServiceProvider_ProcessActiveDelegationStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/service.V2ServiceProvider

// skipping Fuzz_Nosy_V2ServiceProvider_ProcessAndSaveBtcAddresses__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/service.V2ServiceProvider

// skipping Fuzz_Nosy_V2ServiceProvider_ProcessUnbondingDelegationStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/service.V2ServiceProvider

// skipping Fuzz_Nosy_V2ServiceProvider_ProcessWithdrawableDelegationStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/service.V2ServiceProvider

// skipping Fuzz_Nosy_V2ServiceProvider_ProcessWithdrawnDelegationStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/service.V2ServiceProvider

// skipping Fuzz_Nosy_V2ServiceProvider_SaveUnprocessableMessages__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v2/service.V2ServiceProvider

func Fuzz_Nosy_FromDelegationDocument__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var delegation indexerdbmodel.IndexerDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}

		FromDelegationDocument(delegation)
	})
}

func Fuzz_Nosy_findPublicKeyByNativeSegwitAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var providedNativeSegwitAddress string
		fill_err = tp.Fill(&providedNativeSegwitAddress)
		if fill_err != nil {
			return
		}
		var mappings []*dbmodel.PkAddressMapping
		fill_err = tp.Fill(&mappings)
		if fill_err != nil {
			return
		}

		findPublicKeyByNativeSegwitAddress(providedNativeSegwitAddress, mappings)
	})
}

func Fuzz_Nosy_getUnbondingSignatures__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var covenantSignatures []indexerdbmodel.CovenantSignature
		fill_err = tp.Fill(&covenantSignatures)
		if fill_err != nil {
			return
		}

		getUnbondingSignatures(covenantSignatures)
	})
}
