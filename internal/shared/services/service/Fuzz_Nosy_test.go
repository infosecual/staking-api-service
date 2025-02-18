package service

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/staking-api-service/internal/shared/config"
	dbclients "github.com/babylonlabs-io/staking-api-service/internal/shared/db/clients"
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

func Fuzz_Nosy_Service_DoHealthCheck__(f *testing.F) {
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
		s.DoHealthCheck(c7)
	})
}

func Fuzz_Nosy_Service_SaveUnprocessableMessages__(f *testing.F) {
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
		if cfg == nil || globalParams == nil || clients == nil || dbClients == nil {
			return
		}

		s, err := New(c1, cfg, globalParams, finalityProviders, clients, dbClients)
		if err != nil {
			return
		}
		s.SaveUnprocessableMessages(c7, messageBody, receipt)
	})
}

func Fuzz_Nosy_Service_VerifyUTXOs__(f *testing.F) {
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
		var utxos []types.UTXOIdentifier
		fill_err = tp.Fill(&utxos)
		if fill_err != nil {
			return
		}
		var address string
		fill_err = tp.Fill(&address)
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
		s.VerifyUTXOs(c7, utxos, address)
	})
}

func Fuzz_Nosy_Service_verifyViaOrdinalService__(f *testing.F) {
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
		var utxos []types.UTXOIdentifier
		fill_err = tp.Fill(&utxos)
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
		s.verifyViaOrdinalService(c7, utxos)
	})
}

// skipping Fuzz_Nosy_SharedServiceProvider_DoHealthCheck__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/services/service.SharedServiceProvider

// skipping Fuzz_Nosy_SharedServiceProvider_SaveUnprocessableMessages__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/services/service.SharedServiceProvider

// skipping Fuzz_Nosy_SharedServiceProvider_VerifyUTXOs__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/services/service.SharedServiceProvider
