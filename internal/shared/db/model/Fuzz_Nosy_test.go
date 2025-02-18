package dbmodel

import (
	"context"
	"testing"

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

func Fuzz_Nosy_DecodePaginationToken__(f *testing.F) {
	f.Fuzz(func(t *testing.T, token string) {
		DecodePaginationToken(token)
	})
}

// skipping Fuzz_Nosy_GetPaginationToken__ because parameters include func, chan, or unsupported interface: PaginationType

func Fuzz_Nosy_createCollection__(f *testing.F) {
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
		var database *mongo.Database
		fill_err = tp.Fill(&database)
		if fill_err != nil {
			return
		}
		var collectionName string
		fill_err = tp.Fill(&collectionName)
		if fill_err != nil {
			return
		}
		if database == nil {
			return
		}

		createCollection(ctx, database, collectionName)
	})
}

func Fuzz_Nosy_createIndex__(f *testing.F) {
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
		var database *mongo.Database
		fill_err = tp.Fill(&database)
		if fill_err != nil {
			return
		}
		var collectionName string
		fill_err = tp.Fill(&collectionName)
		if fill_err != nil {
			return
		}
		var idx index
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		if database == nil {
			return
		}

		createIndex(ctx, database, collectionName, idx)
	})
}
