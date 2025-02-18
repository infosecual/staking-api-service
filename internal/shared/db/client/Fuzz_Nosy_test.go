package dbclient

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

// skipping Fuzz_Nosy_Database_DeleteUnprocessableMessage__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Database_FindPkMappingsByNativeSegwitAddress__(f *testing.F) {
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
		var nativeSegwitAddresses []string
		fill_err = tp.Fill(&nativeSegwitAddresses)
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
		db.FindPkMappingsByNativeSegwitAddress(c4, nativeSegwitAddresses)
	})
}

func Fuzz_Nosy_Database_FindPkMappingsByTaprootAddress__(f *testing.F) {
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
		var taprootAddresses []string
		fill_err = tp.Fill(&taprootAddresses)
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
		db.FindPkMappingsByTaprootAddress(c4, taprootAddresses)
	})
}

func Fuzz_Nosy_Database_FindUnprocessableMessages__(f *testing.F) {
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
		db.FindUnprocessableMessages(c4)
	})
}

func Fuzz_Nosy_Database_GetLatestPrice__(f *testing.F) {
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
		var symbol string
		fill_err = tp.Fill(&symbol)
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
		db.GetLatestPrice(c4, symbol)
	})
}

func Fuzz_Nosy_Database_InsertPkAddressMappings__(f *testing.F) {
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
		var pkHex string
		fill_err = tp.Fill(&pkHex)
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
		if client == nil || cfg == nil {
			return
		}

		db, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		db.InsertPkAddressMappings(c4, pkHex, taproot, nativeSigwitOdd, nativeSigwitEven)
	})
}

func Fuzz_Nosy_Database_Ping__(f *testing.F) {
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
		db.Ping(c4)
	})
}

func Fuzz_Nosy_Database_SaveUnprocessableMessage__(f *testing.F) {
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
		if client == nil || cfg == nil {
			return
		}

		db, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		db.SaveUnprocessableMessage(c4, messageBody, receipt)
	})
}

func Fuzz_Nosy_Database_SetLatestPrice__(f *testing.F) {
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
		if client == nil || cfg == nil {
			return
		}

		db, err := New(c1, client, cfg)
		if err != nil {
			return
		}
		db.SetLatestPrice(c4, symbol, price)
	})
}

// skipping Fuzz_Nosy_DBClient_DeleteUnprocessableMessage__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/db/client.DBClient

// skipping Fuzz_Nosy_DBClient_FindPkMappingsByNativeSegwitAddress__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/db/client.DBClient

// skipping Fuzz_Nosy_DBClient_FindPkMappingsByTaprootAddress__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/db/client.DBClient

// skipping Fuzz_Nosy_DBClient_FindUnprocessableMessages__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/db/client.DBClient

// skipping Fuzz_Nosy_DBClient_GetLatestPrice__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/db/client.DBClient

// skipping Fuzz_Nosy_DBClient_InsertPkAddressMappings__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/db/client.DBClient

// skipping Fuzz_Nosy_DBClient_Ping__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/db/client.DBClient

// skipping Fuzz_Nosy_DBClient_SaveUnprocessableMessage__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/db/client.DBClient

// skipping Fuzz_Nosy_DBClient_SetLatestPrice__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-api-service/internal/shared/db/client.DBClient
