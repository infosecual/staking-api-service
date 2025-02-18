package v1service

import (
	"context"
	"testing"

	indexerdbmodel "github.com/babylonlabs-io/staking-api-service/internal/indexer/db/model"
	"github.com/babylonlabs-io/staking-api-service/internal/shared/config"
	dbclients "github.com/babylonlabs-io/staking-api-service/internal/shared/db/clients"
	dbmodel "github.com/babylonlabs-io/staking-api-service/internal/shared/db/model"
	"github.com/babylonlabs-io/staking-api-service/internal/shared/http/clients"
	"github.com/babylonlabs-io/staking-api-service/internal/shared/types"
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

func Fuzz_Nosy_V1Service_CheckStakerHasActiveDelegationByPk__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var stakerPk string
		fill_err = tp.Fill(&stakerPk)
		if fill_err != nil {
			return
		}
		var afterTimestamp int64
		fill_err = tp.Fill(&afterTimestamp)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.CheckStakerHasActiveDelegationByPk(c7, stakerPk, afterTimestamp)
	})
}

func Fuzz_Nosy_V1Service_DelegationsByStakerPk__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var stakerPk string
		fill_err = tp.Fill(&stakerPk)
		if fill_err != nil {
			return
		}
		var states []types.DelegationState
		fill_err = tp.Fill(&states)
		if fill_err != nil {
			return
		}
		var pageToken string
		fill_err = tp.Fill(&pageToken)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.DelegationsByStakerPk(c7, stakerPk, states, pageToken)
	})
}

func Fuzz_Nosy_V1Service_FindRegisteredFinalityProvidersNotInUse__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var fpParams []*FpParamsPublic
		fill_err = tp.Fill(&fpParams)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.FindRegisteredFinalityProvidersNotInUse(c7, fpParams)
	})
}

func Fuzz_Nosy_V1Service_FromDelegationDocument__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var d *v1dbmodel.DelegationDocument
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var bbnHeight uint64
		fill_err = tp.Fill(&bbnHeight)
		if fill_err != nil {
			return
		}
		var transitionedFps []*indexerdbmodel.IndexerFinalityProviderDetails
		fill_err = tp.Fill(&transitionedFps)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil || d == nil {
			return
		}

		s, err := New(ctx, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.FromDelegationDocument(d, bbnHeight, transitionedFps)
	})
}

func Fuzz_Nosy_V1Service_GetDelegation__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var txHashHex string
		fill_err = tp.Fill(&txHashHex)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.GetDelegation(c7, txHashHex)
	})
}

func Fuzz_Nosy_V1Service_GetFinalityProvider__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var fpPkHex string
		fill_err = tp.Fill(&fpPkHex)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.GetFinalityProvider(c7, fpPkHex)
	})
}

func Fuzz_Nosy_V1Service_GetFinalityProviders__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var page string
		fill_err = tp.Fill(&page)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.GetFinalityProviders(c7, page)
	})
}

func Fuzz_Nosy_V1Service_GetFinalityProvidersFromGlobalParams__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(ctx, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.GetFinalityProvidersFromGlobalParams()
	})
}

func Fuzz_Nosy_V1Service_GetGlobalParamsPublic__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(ctx, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.GetGlobalParamsPublic()
	})
}

func Fuzz_Nosy_V1Service_GetOverallStats__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.GetOverallStats(c7)
	})
}

func Fuzz_Nosy_V1Service_GetStakerPublicKeysByAddresses__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var addresses []string
		fill_err = tp.Fill(&addresses)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.GetStakerPublicKeysByAddresses(c7, addresses)
	})
}

func Fuzz_Nosy_V1Service_GetStakerStats__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.GetStakerStats(c7, stakerPkHex)
	})
}

func Fuzz_Nosy_V1Service_GetTopStakersByActiveTvl__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var pageToken string
		fill_err = tp.Fill(&pageToken)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.GetTopStakersByActiveTvl(c7, pageToken)
	})
}

func Fuzz_Nosy_V1Service_GetVersionedGlobalParamsByHeight__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(ctx, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.GetVersionedGlobalParamsByHeight(height)
	})
}

func Fuzz_Nosy_V1Service_IsDelegationPresent__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var txHashHex string
		fill_err = tp.Fill(&txHashHex)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.IsDelegationPresent(c7, txHashHex)
	})
}

func Fuzz_Nosy_V1Service_IsEligibleForUnbondingRequest__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.IsEligibleForUnbondingRequest(c7, stakingTxHashHex)
	})
}

func Fuzz_Nosy_V1Service_ProcessAndSaveBtcAddresses__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.ProcessAndSaveBtcAddresses(c7, stakerPkHex)
	})
}

func Fuzz_Nosy_V1Service_ProcessBtcInfoStats__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var btcHeight uint64
		fill_err = tp.Fill(&btcHeight)
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
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.ProcessBtcInfoStats(c7, btcHeight, confirmedTvl, unconfirmedTvl)
	})
}

func Fuzz_Nosy_V1Service_ProcessExpireCheck__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
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
		var txType types.StakingTxType
		fill_err = tp.Fill(&txType)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.ProcessExpireCheck(c7, stakingTxHashHex, startHeight, timelock, txType)
	})
}

func Fuzz_Nosy_V1Service_ProcessStakingStatsCalculation__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
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
		var state types.DelegationState
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.ProcessStakingStatsCalculation(c7, stakingTxHashHex, stakerPkHex, fpPkHex, state, amount)
	})
}

func Fuzz_Nosy_V1Service_SaveActiveStakingDelegation__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var txHashHex string
		fill_err = tp.Fill(&txHashHex)
		if fill_err != nil {
			return
		}
		var stakerPkHex string
		fill_err = tp.Fill(&stakerPkHex)
		if fill_err != nil {
			return
		}
		var finalityProviderPkHex string
		fill_err = tp.Fill(&finalityProviderPkHex)
		if fill_err != nil {
			return
		}
		var value uint64
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var stakingTimestamp int64
		fill_err = tp.Fill(&stakingTimestamp)
		if fill_err != nil {
			return
		}
		var timeLock uint64
		fill_err = tp.Fill(&timeLock)
		if fill_err != nil {
			return
		}
		var stakingOutputIndex uint64
		fill_err = tp.Fill(&stakingOutputIndex)
		if fill_err != nil {
			return
		}
		var stakingTxHex string
		fill_err = tp.Fill(&stakingTxHex)
		if fill_err != nil {
			return
		}
		var isOverflow bool
		fill_err = tp.Fill(&isOverflow)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.SaveActiveStakingDelegation(c7, txHashHex, stakerPkHex, finalityProviderPkHex, value, startHeight, stakingTimestamp, timeLock, stakingOutputIndex, stakingTxHex, isOverflow)
	})
}

func Fuzz_Nosy_V1Service_TransitionToUnbondedState__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var stakingType types.StakingTxType
		fill_err = tp.Fill(&stakingType)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.TransitionToUnbondedState(c7, stakingType, stakingTxHashHex)
	})
}

func Fuzz_Nosy_V1Service_TransitionToUnbondingState__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var unbondingStartHeight uint64
		fill_err = tp.Fill(&unbondingStartHeight)
		if fill_err != nil {
			return
		}
		var unbondingTimelock uint64
		fill_err = tp.Fill(&unbondingTimelock)
		if fill_err != nil {
			return
		}
		var unbondingOutputIndex uint64
		fill_err = tp.Fill(&unbondingOutputIndex)
		if fill_err != nil {
			return
		}
		var unbondingTxHex string
		fill_err = tp.Fill(&unbondingTxHex)
		if fill_err != nil {
			return
		}
		var unbondingStartTimestamp int64
		fill_err = tp.Fill(&unbondingStartTimestamp)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.TransitionToUnbondingState(c7, stakingTxHashHex, unbondingStartHeight, unbondingTimelock, unbondingOutputIndex, unbondingTxHex, unbondingStartTimestamp)
	})
}

func Fuzz_Nosy_V1Service_TransitionToWithdrawnState__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.TransitionToWithdrawnState(c7, stakingTxHashHex)
	})
}

func Fuzz_Nosy_V1Service_UnbondDelegation__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var c7 context.Context
		fill_err = tp.Fill(&c7)
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
		var unbondingTxHex string
		fill_err = tp.Fill(&unbondingTxHex)
		if fill_err != nil {
			return
		}
		var signatureHex string
		fill_err = tp.Fill(&signatureHex)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.UnbondDelegation(c7, stakingTxHashHex, unbondingTxHashHex, unbondingTxHex, signatureHex)
	})
}

func Fuzz_Nosy_V1Service_checkFpStatus__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var fpPk string
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var transitionedFps []*indexerdbmodel.IndexerFinalityProviderDetails
		fill_err = tp.Fill(&transitionedFps)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(ctx, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.checkFpStatus(fpPk, transitionedFps)
	})
}

func Fuzz_Nosy_V1Service_isEligibleForTransition__(f *testing.F) {
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
		var globalParams *types.GlobalParams
		fill_err = tp.Fill(&globalParams)
		if fill_err != nil {
			return
		}
		var finalityProviders []types.FinalityProviderDetails
		fill_err = tp.Fill(&finalityProviders)
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
		var delegation *v1dbmodel.DelegationDocument
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		var bbnHeight uint64
		fill_err = tp.Fill(&bbnHeight)
		if fill_err != nil {
			return
		}
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil || delegation == nil {
			return
		}

		s, err := New(ctx, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.isEligibleForTransition(delegation, bbnHeight)
	})
}

// skipping Fuzz_Nosy_V1ServiceProvider_CheckStakerHasActiveDelegationByPk__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_DelegationsByStakerPk__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_FindRegisteredFinalityProvidersNotInUse__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_GetDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_GetFinalityProvider__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_GetFinalityProviders__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_GetFinalityProvidersFromGlobalParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_GetGlobalParamsPublic__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_GetOverallStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_GetStakerPublicKeysByAddresses__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_GetStakerStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_GetTopStakersByActiveTvl__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_GetVersionedGlobalParamsByHeight__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_IsDelegationPresent__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_IsEligibleForUnbondingRequest__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_ProcessAndSaveBtcAddresses__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_ProcessBtcInfoStats__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_ProcessExpireCheck__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_ProcessStakingStatsCalculation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_SaveActiveStakingDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_TransitionToUnbondedState__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_TransitionToUnbondingState__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_TransitionToWithdrawnState__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

// skipping Fuzz_Nosy_V1ServiceProvider_UnbondDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/v1/service.V1ServiceProvider

func Fuzz_Nosy_buildFallbackFpDetailsPublic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpParams []*FpParamsPublic
		fill_err = tp.Fill(&fpParams)
		if fill_err != nil {
			return
		}

		buildFallbackFpDetailsPublic(fpParams)
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
