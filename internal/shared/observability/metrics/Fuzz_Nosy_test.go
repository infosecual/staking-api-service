package metrics

import (
	"testing"

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

func Fuzz_Nosy_Outcome_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var O Outcome
		fill_err = tp.Fill(&O)
		if fill_err != nil {
			return
		}

		O.String()
	})
}

func Fuzz_Nosy_Init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, metricsPort int) {
		Init(metricsPort)
	})
}

func Fuzz_Nosy_RecordDbError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, method string) {
		RecordDbError(method)
	})
}

func Fuzz_Nosy_RecordHttpResponseWriteFailure__(f *testing.F) {
	f.Fuzz(func(t *testing.T, statusCode int) {
		RecordHttpResponseWriteFailure(statusCode)
	})
}

func Fuzz_Nosy_RecordQueueOperationFailure__(f *testing.F) {
	f.Fuzz(func(t *testing.T, operation string, queuename string) {
		RecordQueueOperationFailure(operation, queuename)
	})
}

func Fuzz_Nosy_RecordServiceCrash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, service string) {
		RecordServiceCrash(service)
	})
}

func Fuzz_Nosy_RecordUnprocessableEntity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, entity string) {
		RecordUnprocessableEntity(entity)
	})
}

func Fuzz_Nosy_StartClientRequestDurationTimer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, baseUrl string, method string, path string) {
		StartClientRequestDurationTimer(baseUrl, method, path)
	})
}

func Fuzz_Nosy_StartEventProcessingDurationTimer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, queuename string, attempts int32) {
		StartEventProcessingDurationTimer(queuename, attempts)
	})
}

func Fuzz_Nosy_StartHttpRequestDurationTimer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, endpoint string) {
		StartHttpRequestDurationTimer(endpoint)
	})
}

func Fuzz_Nosy_initMetricsRouter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, metricsPort int) {
		initMetricsRouter(metricsPort)
	})
}
