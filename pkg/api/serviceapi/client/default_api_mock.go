// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package client

import (
	"context"
	"net/http"
	"sync"
)

// Ensure, that DefaultApiMock does implement DefaultApi.
// If this is not the case, regenerate this file with moq.
var _ DefaultApi = &DefaultApiMock{}

// DefaultApiMock is a mock implementation of DefaultApi.
//
//     func TestSomethingThatUsesDefaultApi(t *testing.T) {
//
//         // make and configure a mocked DefaultApi
//         mockedDefaultApi := &DefaultApiMock{
//             CreateKafkaFunc: func(ctx context.Context) ApiCreateKafkaRequest {
// 	               panic("mock out the CreateKafka method")
//             },
//             CreateKafkaExecuteFunc: func(r ApiCreateKafkaRequest) (KafkaRequest, *http.Response, GenericOpenAPIError) {
// 	               panic("mock out the CreateKafkaExecute method")
//             },
//             CreateServiceAccountFunc: func(ctx context.Context) ApiCreateServiceAccountRequest {
// 	               panic("mock out the CreateServiceAccount method")
//             },
//             CreateServiceAccountExecuteFunc: func(r ApiCreateServiceAccountRequest) (ServiceAccount, *http.Response, GenericOpenAPIError) {
// 	               panic("mock out the CreateServiceAccountExecute method")
//             },
//             DeleteKafkaByIdFunc: func(ctx context.Context, id string) ApiDeleteKafkaByIdRequest {
// 	               panic("mock out the DeleteKafkaById method")
//             },
//             DeleteKafkaByIdExecuteFunc: func(r ApiDeleteKafkaByIdRequest) (Error, *http.Response, GenericOpenAPIError) {
// 	               panic("mock out the DeleteKafkaByIdExecute method")
//             },
//             DeleteServiceAccountFunc: func(ctx context.Context, id string) ApiDeleteServiceAccountRequest {
// 	               panic("mock out the DeleteServiceAccount method")
//             },
//             DeleteServiceAccountExecuteFunc: func(r ApiDeleteServiceAccountRequest) (Error, *http.Response, GenericOpenAPIError) {
// 	               panic("mock out the DeleteServiceAccountExecute method")
//             },
//             GetKafkaByIdFunc: func(ctx context.Context, id string) ApiGetKafkaByIdRequest {
// 	               panic("mock out the GetKafkaById method")
//             },
//             GetKafkaByIdExecuteFunc: func(r ApiGetKafkaByIdRequest) (KafkaRequest, *http.Response, GenericOpenAPIError) {
// 	               panic("mock out the GetKafkaByIdExecute method")
//             },
//             ListCloudProviderRegionsFunc: func(ctx context.Context, id string) ApiListCloudProviderRegionsRequest {
// 	               panic("mock out the ListCloudProviderRegions method")
//             },
//             ListCloudProviderRegionsExecuteFunc: func(r ApiListCloudProviderRegionsRequest) (CloudRegionList, *http.Response, GenericOpenAPIError) {
// 	               panic("mock out the ListCloudProviderRegionsExecute method")
//             },
//             ListCloudProvidersFunc: func(ctx context.Context) ApiListCloudProvidersRequest {
// 	               panic("mock out the ListCloudProviders method")
//             },
//             ListCloudProvidersExecuteFunc: func(r ApiListCloudProvidersRequest) (CloudProviderList, *http.Response, GenericOpenAPIError) {
// 	               panic("mock out the ListCloudProvidersExecute method")
//             },
//             ListKafkasFunc: func(ctx context.Context) ApiListKafkasRequest {
// 	               panic("mock out the ListKafkas method")
//             },
//             ListKafkasExecuteFunc: func(r ApiListKafkasRequest) (KafkaRequestList, *http.Response, GenericOpenAPIError) {
// 	               panic("mock out the ListKafkasExecute method")
//             },
//             ListServiceAccountsFunc: func(ctx context.Context) ApiListServiceAccountsRequest {
// 	               panic("mock out the ListServiceAccounts method")
//             },
//             ListServiceAccountsExecuteFunc: func(r ApiListServiceAccountsRequest) (ServiceAccountList, *http.Response, GenericOpenAPIError) {
// 	               panic("mock out the ListServiceAccountsExecute method")
//             },
//             ResetServiceAccountCredsFunc: func(ctx context.Context, id string) ApiResetServiceAccountCredsRequest {
// 	               panic("mock out the ResetServiceAccountCreds method")
//             },
//             ResetServiceAccountCredsExecuteFunc: func(r ApiResetServiceAccountCredsRequest) (ServiceAccount, *http.Response, GenericOpenAPIError) {
// 	               panic("mock out the ResetServiceAccountCredsExecute method")
//             },
//         }
//
//         // use mockedDefaultApi in code that requires DefaultApi
//         // and then make assertions.
//
//     }
type DefaultApiMock struct {
	// CreateKafkaFunc mocks the CreateKafka method.
	CreateKafkaFunc func(ctx context.Context) ApiCreateKafkaRequest

	// CreateKafkaExecuteFunc mocks the CreateKafkaExecute method.
	CreateKafkaExecuteFunc func(r ApiCreateKafkaRequest) (KafkaRequest, *http.Response, GenericOpenAPIError)

	// CreateServiceAccountFunc mocks the CreateServiceAccount method.
	CreateServiceAccountFunc func(ctx context.Context) ApiCreateServiceAccountRequest

	// CreateServiceAccountExecuteFunc mocks the CreateServiceAccountExecute method.
	CreateServiceAccountExecuteFunc func(r ApiCreateServiceAccountRequest) (ServiceAccount, *http.Response, GenericOpenAPIError)

	// DeleteKafkaByIdFunc mocks the DeleteKafkaById method.
	DeleteKafkaByIdFunc func(ctx context.Context, id string) ApiDeleteKafkaByIdRequest

	// DeleteKafkaByIdExecuteFunc mocks the DeleteKafkaByIdExecute method.
	DeleteKafkaByIdExecuteFunc func(r ApiDeleteKafkaByIdRequest) (Error, *http.Response, GenericOpenAPIError)

	// DeleteServiceAccountFunc mocks the DeleteServiceAccount method.
	DeleteServiceAccountFunc func(ctx context.Context, id string) ApiDeleteServiceAccountRequest

	// DeleteServiceAccountExecuteFunc mocks the DeleteServiceAccountExecute method.
	DeleteServiceAccountExecuteFunc func(r ApiDeleteServiceAccountRequest) (Error, *http.Response, GenericOpenAPIError)

	// GetKafkaByIdFunc mocks the GetKafkaById method.
	GetKafkaByIdFunc func(ctx context.Context, id string) ApiGetKafkaByIdRequest

	// GetKafkaByIdExecuteFunc mocks the GetKafkaByIdExecute method.
	GetKafkaByIdExecuteFunc func(r ApiGetKafkaByIdRequest) (KafkaRequest, *http.Response, GenericOpenAPIError)

	// ListCloudProviderRegionsFunc mocks the ListCloudProviderRegions method.
	ListCloudProviderRegionsFunc func(ctx context.Context, id string) ApiListCloudProviderRegionsRequest

	// ListCloudProviderRegionsExecuteFunc mocks the ListCloudProviderRegionsExecute method.
	ListCloudProviderRegionsExecuteFunc func(r ApiListCloudProviderRegionsRequest) (CloudRegionList, *http.Response, GenericOpenAPIError)

	// ListCloudProvidersFunc mocks the ListCloudProviders method.
	ListCloudProvidersFunc func(ctx context.Context) ApiListCloudProvidersRequest

	// ListCloudProvidersExecuteFunc mocks the ListCloudProvidersExecute method.
	ListCloudProvidersExecuteFunc func(r ApiListCloudProvidersRequest) (CloudProviderList, *http.Response, GenericOpenAPIError)

	// ListKafkasFunc mocks the ListKafkas method.
	ListKafkasFunc func(ctx context.Context) ApiListKafkasRequest

	// ListKafkasExecuteFunc mocks the ListKafkasExecute method.
	ListKafkasExecuteFunc func(r ApiListKafkasRequest) (KafkaRequestList, *http.Response, GenericOpenAPIError)

	// ListServiceAccountsFunc mocks the ListServiceAccounts method.
	ListServiceAccountsFunc func(ctx context.Context) ApiListServiceAccountsRequest

	// ListServiceAccountsExecuteFunc mocks the ListServiceAccountsExecute method.
	ListServiceAccountsExecuteFunc func(r ApiListServiceAccountsRequest) (ServiceAccountList, *http.Response, GenericOpenAPIError)

	// ResetServiceAccountCredsFunc mocks the ResetServiceAccountCreds method.
	ResetServiceAccountCredsFunc func(ctx context.Context, id string) ApiResetServiceAccountCredsRequest

	// ResetServiceAccountCredsExecuteFunc mocks the ResetServiceAccountCredsExecute method.
	ResetServiceAccountCredsExecuteFunc func(r ApiResetServiceAccountCredsRequest) (ServiceAccount, *http.Response, GenericOpenAPIError)

	// calls tracks calls to the methods.
	calls struct {
		// CreateKafka holds details about calls to the CreateKafka method.
		CreateKafka []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// CreateKafkaExecute holds details about calls to the CreateKafkaExecute method.
		CreateKafkaExecute []struct {
			// R is the r argument value.
			R ApiCreateKafkaRequest
		}
		// CreateServiceAccount holds details about calls to the CreateServiceAccount method.
		CreateServiceAccount []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// CreateServiceAccountExecute holds details about calls to the CreateServiceAccountExecute method.
		CreateServiceAccountExecute []struct {
			// R is the r argument value.
			R ApiCreateServiceAccountRequest
		}
		// DeleteKafkaById holds details about calls to the DeleteKafkaById method.
		DeleteKafkaById []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// DeleteKafkaByIdExecute holds details about calls to the DeleteKafkaByIdExecute method.
		DeleteKafkaByIdExecute []struct {
			// R is the r argument value.
			R ApiDeleteKafkaByIdRequest
		}
		// DeleteServiceAccount holds details about calls to the DeleteServiceAccount method.
		DeleteServiceAccount []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// DeleteServiceAccountExecute holds details about calls to the DeleteServiceAccountExecute method.
		DeleteServiceAccountExecute []struct {
			// R is the r argument value.
			R ApiDeleteServiceAccountRequest
		}
		// GetKafkaById holds details about calls to the GetKafkaById method.
		GetKafkaById []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// GetKafkaByIdExecute holds details about calls to the GetKafkaByIdExecute method.
		GetKafkaByIdExecute []struct {
			// R is the r argument value.
			R ApiGetKafkaByIdRequest
		}
		// ListCloudProviderRegions holds details about calls to the ListCloudProviderRegions method.
		ListCloudProviderRegions []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// ListCloudProviderRegionsExecute holds details about calls to the ListCloudProviderRegionsExecute method.
		ListCloudProviderRegionsExecute []struct {
			// R is the r argument value.
			R ApiListCloudProviderRegionsRequest
		}
		// ListCloudProviders holds details about calls to the ListCloudProviders method.
		ListCloudProviders []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// ListCloudProvidersExecute holds details about calls to the ListCloudProvidersExecute method.
		ListCloudProvidersExecute []struct {
			// R is the r argument value.
			R ApiListCloudProvidersRequest
		}
		// ListKafkas holds details about calls to the ListKafkas method.
		ListKafkas []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// ListKafkasExecute holds details about calls to the ListKafkasExecute method.
		ListKafkasExecute []struct {
			// R is the r argument value.
			R ApiListKafkasRequest
		}
		// ListServiceAccounts holds details about calls to the ListServiceAccounts method.
		ListServiceAccounts []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// ListServiceAccountsExecute holds details about calls to the ListServiceAccountsExecute method.
		ListServiceAccountsExecute []struct {
			// R is the r argument value.
			R ApiListServiceAccountsRequest
		}
		// ResetServiceAccountCreds holds details about calls to the ResetServiceAccountCreds method.
		ResetServiceAccountCreds []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// ResetServiceAccountCredsExecute holds details about calls to the ResetServiceAccountCredsExecute method.
		ResetServiceAccountCredsExecute []struct {
			// R is the r argument value.
			R ApiResetServiceAccountCredsRequest
		}
	}
	lockCreateKafka                     sync.RWMutex
	lockCreateKafkaExecute              sync.RWMutex
	lockCreateServiceAccount            sync.RWMutex
	lockCreateServiceAccountExecute     sync.RWMutex
	lockDeleteKafkaById                 sync.RWMutex
	lockDeleteKafkaByIdExecute          sync.RWMutex
	lockDeleteServiceAccount            sync.RWMutex
	lockDeleteServiceAccountExecute     sync.RWMutex
	lockGetKafkaById                    sync.RWMutex
	lockGetKafkaByIdExecute             sync.RWMutex
	lockListCloudProviderRegions        sync.RWMutex
	lockListCloudProviderRegionsExecute sync.RWMutex
	lockListCloudProviders              sync.RWMutex
	lockListCloudProvidersExecute       sync.RWMutex
	lockListKafkas                      sync.RWMutex
	lockListKafkasExecute               sync.RWMutex
	lockListServiceAccounts             sync.RWMutex
	lockListServiceAccountsExecute      sync.RWMutex
	lockResetServiceAccountCreds        sync.RWMutex
	lockResetServiceAccountCredsExecute sync.RWMutex
}

// CreateKafka calls CreateKafkaFunc.
func (mock *DefaultApiMock) CreateKafka(ctx context.Context) ApiCreateKafkaRequest {
	if mock.CreateKafkaFunc == nil {
		panic("DefaultApiMock.CreateKafkaFunc: method is nil but DefaultApi.CreateKafka was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockCreateKafka.Lock()
	mock.calls.CreateKafka = append(mock.calls.CreateKafka, callInfo)
	mock.lockCreateKafka.Unlock()
	return mock.CreateKafkaFunc(ctx)
}

// CreateKafkaCalls gets all the calls that were made to CreateKafka.
// Check the length with:
//     len(mockedDefaultApi.CreateKafkaCalls())
func (mock *DefaultApiMock) CreateKafkaCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockCreateKafka.RLock()
	calls = mock.calls.CreateKafka
	mock.lockCreateKafka.RUnlock()
	return calls
}

// CreateKafkaExecute calls CreateKafkaExecuteFunc.
func (mock *DefaultApiMock) CreateKafkaExecute(r ApiCreateKafkaRequest) (KafkaRequest, *http.Response, GenericOpenAPIError) {
	if mock.CreateKafkaExecuteFunc == nil {
		panic("DefaultApiMock.CreateKafkaExecuteFunc: method is nil but DefaultApi.CreateKafkaExecute was just called")
	}
	callInfo := struct {
		R ApiCreateKafkaRequest
	}{
		R: r,
	}
	mock.lockCreateKafkaExecute.Lock()
	mock.calls.CreateKafkaExecute = append(mock.calls.CreateKafkaExecute, callInfo)
	mock.lockCreateKafkaExecute.Unlock()
	return mock.CreateKafkaExecuteFunc(r)
}

// CreateKafkaExecuteCalls gets all the calls that were made to CreateKafkaExecute.
// Check the length with:
//     len(mockedDefaultApi.CreateKafkaExecuteCalls())
func (mock *DefaultApiMock) CreateKafkaExecuteCalls() []struct {
	R ApiCreateKafkaRequest
} {
	var calls []struct {
		R ApiCreateKafkaRequest
	}
	mock.lockCreateKafkaExecute.RLock()
	calls = mock.calls.CreateKafkaExecute
	mock.lockCreateKafkaExecute.RUnlock()
	return calls
}

// CreateServiceAccount calls CreateServiceAccountFunc.
func (mock *DefaultApiMock) CreateServiceAccount(ctx context.Context) ApiCreateServiceAccountRequest {
	if mock.CreateServiceAccountFunc == nil {
		panic("DefaultApiMock.CreateServiceAccountFunc: method is nil but DefaultApi.CreateServiceAccount was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockCreateServiceAccount.Lock()
	mock.calls.CreateServiceAccount = append(mock.calls.CreateServiceAccount, callInfo)
	mock.lockCreateServiceAccount.Unlock()
	return mock.CreateServiceAccountFunc(ctx)
}

// CreateServiceAccountCalls gets all the calls that were made to CreateServiceAccount.
// Check the length with:
//     len(mockedDefaultApi.CreateServiceAccountCalls())
func (mock *DefaultApiMock) CreateServiceAccountCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockCreateServiceAccount.RLock()
	calls = mock.calls.CreateServiceAccount
	mock.lockCreateServiceAccount.RUnlock()
	return calls
}

// CreateServiceAccountExecute calls CreateServiceAccountExecuteFunc.
func (mock *DefaultApiMock) CreateServiceAccountExecute(r ApiCreateServiceAccountRequest) (ServiceAccount, *http.Response, GenericOpenAPIError) {
	if mock.CreateServiceAccountExecuteFunc == nil {
		panic("DefaultApiMock.CreateServiceAccountExecuteFunc: method is nil but DefaultApi.CreateServiceAccountExecute was just called")
	}
	callInfo := struct {
		R ApiCreateServiceAccountRequest
	}{
		R: r,
	}
	mock.lockCreateServiceAccountExecute.Lock()
	mock.calls.CreateServiceAccountExecute = append(mock.calls.CreateServiceAccountExecute, callInfo)
	mock.lockCreateServiceAccountExecute.Unlock()
	return mock.CreateServiceAccountExecuteFunc(r)
}

// CreateServiceAccountExecuteCalls gets all the calls that were made to CreateServiceAccountExecute.
// Check the length with:
//     len(mockedDefaultApi.CreateServiceAccountExecuteCalls())
func (mock *DefaultApiMock) CreateServiceAccountExecuteCalls() []struct {
	R ApiCreateServiceAccountRequest
} {
	var calls []struct {
		R ApiCreateServiceAccountRequest
	}
	mock.lockCreateServiceAccountExecute.RLock()
	calls = mock.calls.CreateServiceAccountExecute
	mock.lockCreateServiceAccountExecute.RUnlock()
	return calls
}

// DeleteKafkaById calls DeleteKafkaByIdFunc.
func (mock *DefaultApiMock) DeleteKafkaById(ctx context.Context, id string) ApiDeleteKafkaByIdRequest {
	if mock.DeleteKafkaByIdFunc == nil {
		panic("DefaultApiMock.DeleteKafkaByIdFunc: method is nil but DefaultApi.DeleteKafkaById was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteKafkaById.Lock()
	mock.calls.DeleteKafkaById = append(mock.calls.DeleteKafkaById, callInfo)
	mock.lockDeleteKafkaById.Unlock()
	return mock.DeleteKafkaByIdFunc(ctx, id)
}

// DeleteKafkaByIdCalls gets all the calls that were made to DeleteKafkaById.
// Check the length with:
//     len(mockedDefaultApi.DeleteKafkaByIdCalls())
func (mock *DefaultApiMock) DeleteKafkaByIdCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockDeleteKafkaById.RLock()
	calls = mock.calls.DeleteKafkaById
	mock.lockDeleteKafkaById.RUnlock()
	return calls
}

// DeleteKafkaByIdExecute calls DeleteKafkaByIdExecuteFunc.
func (mock *DefaultApiMock) DeleteKafkaByIdExecute(r ApiDeleteKafkaByIdRequest) (Error, *http.Response, GenericOpenAPIError) {
	if mock.DeleteKafkaByIdExecuteFunc == nil {
		panic("DefaultApiMock.DeleteKafkaByIdExecuteFunc: method is nil but DefaultApi.DeleteKafkaByIdExecute was just called")
	}
	callInfo := struct {
		R ApiDeleteKafkaByIdRequest
	}{
		R: r,
	}
	mock.lockDeleteKafkaByIdExecute.Lock()
	mock.calls.DeleteKafkaByIdExecute = append(mock.calls.DeleteKafkaByIdExecute, callInfo)
	mock.lockDeleteKafkaByIdExecute.Unlock()
	return mock.DeleteKafkaByIdExecuteFunc(r)
}

// DeleteKafkaByIdExecuteCalls gets all the calls that were made to DeleteKafkaByIdExecute.
// Check the length with:
//     len(mockedDefaultApi.DeleteKafkaByIdExecuteCalls())
func (mock *DefaultApiMock) DeleteKafkaByIdExecuteCalls() []struct {
	R ApiDeleteKafkaByIdRequest
} {
	var calls []struct {
		R ApiDeleteKafkaByIdRequest
	}
	mock.lockDeleteKafkaByIdExecute.RLock()
	calls = mock.calls.DeleteKafkaByIdExecute
	mock.lockDeleteKafkaByIdExecute.RUnlock()
	return calls
}

// DeleteServiceAccount calls DeleteServiceAccountFunc.
func (mock *DefaultApiMock) DeleteServiceAccount(ctx context.Context, id string) ApiDeleteServiceAccountRequest {
	if mock.DeleteServiceAccountFunc == nil {
		panic("DefaultApiMock.DeleteServiceAccountFunc: method is nil but DefaultApi.DeleteServiceAccount was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteServiceAccount.Lock()
	mock.calls.DeleteServiceAccount = append(mock.calls.DeleteServiceAccount, callInfo)
	mock.lockDeleteServiceAccount.Unlock()
	return mock.DeleteServiceAccountFunc(ctx, id)
}

// DeleteServiceAccountCalls gets all the calls that were made to DeleteServiceAccount.
// Check the length with:
//     len(mockedDefaultApi.DeleteServiceAccountCalls())
func (mock *DefaultApiMock) DeleteServiceAccountCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockDeleteServiceAccount.RLock()
	calls = mock.calls.DeleteServiceAccount
	mock.lockDeleteServiceAccount.RUnlock()
	return calls
}

// DeleteServiceAccountExecute calls DeleteServiceAccountExecuteFunc.
func (mock *DefaultApiMock) DeleteServiceAccountExecute(r ApiDeleteServiceAccountRequest) (Error, *http.Response, GenericOpenAPIError) {
	if mock.DeleteServiceAccountExecuteFunc == nil {
		panic("DefaultApiMock.DeleteServiceAccountExecuteFunc: method is nil but DefaultApi.DeleteServiceAccountExecute was just called")
	}
	callInfo := struct {
		R ApiDeleteServiceAccountRequest
	}{
		R: r,
	}
	mock.lockDeleteServiceAccountExecute.Lock()
	mock.calls.DeleteServiceAccountExecute = append(mock.calls.DeleteServiceAccountExecute, callInfo)
	mock.lockDeleteServiceAccountExecute.Unlock()
	return mock.DeleteServiceAccountExecuteFunc(r)
}

// DeleteServiceAccountExecuteCalls gets all the calls that were made to DeleteServiceAccountExecute.
// Check the length with:
//     len(mockedDefaultApi.DeleteServiceAccountExecuteCalls())
func (mock *DefaultApiMock) DeleteServiceAccountExecuteCalls() []struct {
	R ApiDeleteServiceAccountRequest
} {
	var calls []struct {
		R ApiDeleteServiceAccountRequest
	}
	mock.lockDeleteServiceAccountExecute.RLock()
	calls = mock.calls.DeleteServiceAccountExecute
	mock.lockDeleteServiceAccountExecute.RUnlock()
	return calls
}

// GetKafkaById calls GetKafkaByIdFunc.
func (mock *DefaultApiMock) GetKafkaById(ctx context.Context, id string) ApiGetKafkaByIdRequest {
	if mock.GetKafkaByIdFunc == nil {
		panic("DefaultApiMock.GetKafkaByIdFunc: method is nil but DefaultApi.GetKafkaById was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetKafkaById.Lock()
	mock.calls.GetKafkaById = append(mock.calls.GetKafkaById, callInfo)
	mock.lockGetKafkaById.Unlock()
	return mock.GetKafkaByIdFunc(ctx, id)
}

// GetKafkaByIdCalls gets all the calls that were made to GetKafkaById.
// Check the length with:
//     len(mockedDefaultApi.GetKafkaByIdCalls())
func (mock *DefaultApiMock) GetKafkaByIdCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockGetKafkaById.RLock()
	calls = mock.calls.GetKafkaById
	mock.lockGetKafkaById.RUnlock()
	return calls
}

// GetKafkaByIdExecute calls GetKafkaByIdExecuteFunc.
func (mock *DefaultApiMock) GetKafkaByIdExecute(r ApiGetKafkaByIdRequest) (KafkaRequest, *http.Response, GenericOpenAPIError) {
	if mock.GetKafkaByIdExecuteFunc == nil {
		panic("DefaultApiMock.GetKafkaByIdExecuteFunc: method is nil but DefaultApi.GetKafkaByIdExecute was just called")
	}
	callInfo := struct {
		R ApiGetKafkaByIdRequest
	}{
		R: r,
	}
	mock.lockGetKafkaByIdExecute.Lock()
	mock.calls.GetKafkaByIdExecute = append(mock.calls.GetKafkaByIdExecute, callInfo)
	mock.lockGetKafkaByIdExecute.Unlock()
	return mock.GetKafkaByIdExecuteFunc(r)
}

// GetKafkaByIdExecuteCalls gets all the calls that were made to GetKafkaByIdExecute.
// Check the length with:
//     len(mockedDefaultApi.GetKafkaByIdExecuteCalls())
func (mock *DefaultApiMock) GetKafkaByIdExecuteCalls() []struct {
	R ApiGetKafkaByIdRequest
} {
	var calls []struct {
		R ApiGetKafkaByIdRequest
	}
	mock.lockGetKafkaByIdExecute.RLock()
	calls = mock.calls.GetKafkaByIdExecute
	mock.lockGetKafkaByIdExecute.RUnlock()
	return calls
}

// ListCloudProviderRegions calls ListCloudProviderRegionsFunc.
func (mock *DefaultApiMock) ListCloudProviderRegions(ctx context.Context, id string) ApiListCloudProviderRegionsRequest {
	if mock.ListCloudProviderRegionsFunc == nil {
		panic("DefaultApiMock.ListCloudProviderRegionsFunc: method is nil but DefaultApi.ListCloudProviderRegions was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockListCloudProviderRegions.Lock()
	mock.calls.ListCloudProviderRegions = append(mock.calls.ListCloudProviderRegions, callInfo)
	mock.lockListCloudProviderRegions.Unlock()
	return mock.ListCloudProviderRegionsFunc(ctx, id)
}

// ListCloudProviderRegionsCalls gets all the calls that were made to ListCloudProviderRegions.
// Check the length with:
//     len(mockedDefaultApi.ListCloudProviderRegionsCalls())
func (mock *DefaultApiMock) ListCloudProviderRegionsCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockListCloudProviderRegions.RLock()
	calls = mock.calls.ListCloudProviderRegions
	mock.lockListCloudProviderRegions.RUnlock()
	return calls
}

// ListCloudProviderRegionsExecute calls ListCloudProviderRegionsExecuteFunc.
func (mock *DefaultApiMock) ListCloudProviderRegionsExecute(r ApiListCloudProviderRegionsRequest) (CloudRegionList, *http.Response, GenericOpenAPIError) {
	if mock.ListCloudProviderRegionsExecuteFunc == nil {
		panic("DefaultApiMock.ListCloudProviderRegionsExecuteFunc: method is nil but DefaultApi.ListCloudProviderRegionsExecute was just called")
	}
	callInfo := struct {
		R ApiListCloudProviderRegionsRequest
	}{
		R: r,
	}
	mock.lockListCloudProviderRegionsExecute.Lock()
	mock.calls.ListCloudProviderRegionsExecute = append(mock.calls.ListCloudProviderRegionsExecute, callInfo)
	mock.lockListCloudProviderRegionsExecute.Unlock()
	return mock.ListCloudProviderRegionsExecuteFunc(r)
}

// ListCloudProviderRegionsExecuteCalls gets all the calls that were made to ListCloudProviderRegionsExecute.
// Check the length with:
//     len(mockedDefaultApi.ListCloudProviderRegionsExecuteCalls())
func (mock *DefaultApiMock) ListCloudProviderRegionsExecuteCalls() []struct {
	R ApiListCloudProviderRegionsRequest
} {
	var calls []struct {
		R ApiListCloudProviderRegionsRequest
	}
	mock.lockListCloudProviderRegionsExecute.RLock()
	calls = mock.calls.ListCloudProviderRegionsExecute
	mock.lockListCloudProviderRegionsExecute.RUnlock()
	return calls
}

// ListCloudProviders calls ListCloudProvidersFunc.
func (mock *DefaultApiMock) ListCloudProviders(ctx context.Context) ApiListCloudProvidersRequest {
	if mock.ListCloudProvidersFunc == nil {
		panic("DefaultApiMock.ListCloudProvidersFunc: method is nil but DefaultApi.ListCloudProviders was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockListCloudProviders.Lock()
	mock.calls.ListCloudProviders = append(mock.calls.ListCloudProviders, callInfo)
	mock.lockListCloudProviders.Unlock()
	return mock.ListCloudProvidersFunc(ctx)
}

// ListCloudProvidersCalls gets all the calls that were made to ListCloudProviders.
// Check the length with:
//     len(mockedDefaultApi.ListCloudProvidersCalls())
func (mock *DefaultApiMock) ListCloudProvidersCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockListCloudProviders.RLock()
	calls = mock.calls.ListCloudProviders
	mock.lockListCloudProviders.RUnlock()
	return calls
}

// ListCloudProvidersExecute calls ListCloudProvidersExecuteFunc.
func (mock *DefaultApiMock) ListCloudProvidersExecute(r ApiListCloudProvidersRequest) (CloudProviderList, *http.Response, GenericOpenAPIError) {
	if mock.ListCloudProvidersExecuteFunc == nil {
		panic("DefaultApiMock.ListCloudProvidersExecuteFunc: method is nil but DefaultApi.ListCloudProvidersExecute was just called")
	}
	callInfo := struct {
		R ApiListCloudProvidersRequest
	}{
		R: r,
	}
	mock.lockListCloudProvidersExecute.Lock()
	mock.calls.ListCloudProvidersExecute = append(mock.calls.ListCloudProvidersExecute, callInfo)
	mock.lockListCloudProvidersExecute.Unlock()
	return mock.ListCloudProvidersExecuteFunc(r)
}

// ListCloudProvidersExecuteCalls gets all the calls that were made to ListCloudProvidersExecute.
// Check the length with:
//     len(mockedDefaultApi.ListCloudProvidersExecuteCalls())
func (mock *DefaultApiMock) ListCloudProvidersExecuteCalls() []struct {
	R ApiListCloudProvidersRequest
} {
	var calls []struct {
		R ApiListCloudProvidersRequest
	}
	mock.lockListCloudProvidersExecute.RLock()
	calls = mock.calls.ListCloudProvidersExecute
	mock.lockListCloudProvidersExecute.RUnlock()
	return calls
}

// ListKafkas calls ListKafkasFunc.
func (mock *DefaultApiMock) ListKafkas(ctx context.Context) ApiListKafkasRequest {
	if mock.ListKafkasFunc == nil {
		panic("DefaultApiMock.ListKafkasFunc: method is nil but DefaultApi.ListKafkas was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockListKafkas.Lock()
	mock.calls.ListKafkas = append(mock.calls.ListKafkas, callInfo)
	mock.lockListKafkas.Unlock()
	return mock.ListKafkasFunc(ctx)
}

// ListKafkasCalls gets all the calls that were made to ListKafkas.
// Check the length with:
//     len(mockedDefaultApi.ListKafkasCalls())
func (mock *DefaultApiMock) ListKafkasCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockListKafkas.RLock()
	calls = mock.calls.ListKafkas
	mock.lockListKafkas.RUnlock()
	return calls
}

// ListKafkasExecute calls ListKafkasExecuteFunc.
func (mock *DefaultApiMock) ListKafkasExecute(r ApiListKafkasRequest) (KafkaRequestList, *http.Response, GenericOpenAPIError) {
	if mock.ListKafkasExecuteFunc == nil {
		panic("DefaultApiMock.ListKafkasExecuteFunc: method is nil but DefaultApi.ListKafkasExecute was just called")
	}
	callInfo := struct {
		R ApiListKafkasRequest
	}{
		R: r,
	}
	mock.lockListKafkasExecute.Lock()
	mock.calls.ListKafkasExecute = append(mock.calls.ListKafkasExecute, callInfo)
	mock.lockListKafkasExecute.Unlock()
	return mock.ListKafkasExecuteFunc(r)
}

// ListKafkasExecuteCalls gets all the calls that were made to ListKafkasExecute.
// Check the length with:
//     len(mockedDefaultApi.ListKafkasExecuteCalls())
func (mock *DefaultApiMock) ListKafkasExecuteCalls() []struct {
	R ApiListKafkasRequest
} {
	var calls []struct {
		R ApiListKafkasRequest
	}
	mock.lockListKafkasExecute.RLock()
	calls = mock.calls.ListKafkasExecute
	mock.lockListKafkasExecute.RUnlock()
	return calls
}

// ListServiceAccounts calls ListServiceAccountsFunc.
func (mock *DefaultApiMock) ListServiceAccounts(ctx context.Context) ApiListServiceAccountsRequest {
	if mock.ListServiceAccountsFunc == nil {
		panic("DefaultApiMock.ListServiceAccountsFunc: method is nil but DefaultApi.ListServiceAccounts was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockListServiceAccounts.Lock()
	mock.calls.ListServiceAccounts = append(mock.calls.ListServiceAccounts, callInfo)
	mock.lockListServiceAccounts.Unlock()
	return mock.ListServiceAccountsFunc(ctx)
}

// ListServiceAccountsCalls gets all the calls that were made to ListServiceAccounts.
// Check the length with:
//     len(mockedDefaultApi.ListServiceAccountsCalls())
func (mock *DefaultApiMock) ListServiceAccountsCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockListServiceAccounts.RLock()
	calls = mock.calls.ListServiceAccounts
	mock.lockListServiceAccounts.RUnlock()
	return calls
}

// ListServiceAccountsExecute calls ListServiceAccountsExecuteFunc.
func (mock *DefaultApiMock) ListServiceAccountsExecute(r ApiListServiceAccountsRequest) (ServiceAccountList, *http.Response, GenericOpenAPIError) {
	if mock.ListServiceAccountsExecuteFunc == nil {
		panic("DefaultApiMock.ListServiceAccountsExecuteFunc: method is nil but DefaultApi.ListServiceAccountsExecute was just called")
	}
	callInfo := struct {
		R ApiListServiceAccountsRequest
	}{
		R: r,
	}
	mock.lockListServiceAccountsExecute.Lock()
	mock.calls.ListServiceAccountsExecute = append(mock.calls.ListServiceAccountsExecute, callInfo)
	mock.lockListServiceAccountsExecute.Unlock()
	return mock.ListServiceAccountsExecuteFunc(r)
}

// ListServiceAccountsExecuteCalls gets all the calls that were made to ListServiceAccountsExecute.
// Check the length with:
//     len(mockedDefaultApi.ListServiceAccountsExecuteCalls())
func (mock *DefaultApiMock) ListServiceAccountsExecuteCalls() []struct {
	R ApiListServiceAccountsRequest
} {
	var calls []struct {
		R ApiListServiceAccountsRequest
	}
	mock.lockListServiceAccountsExecute.RLock()
	calls = mock.calls.ListServiceAccountsExecute
	mock.lockListServiceAccountsExecute.RUnlock()
	return calls
}

// ResetServiceAccountCreds calls ResetServiceAccountCredsFunc.
func (mock *DefaultApiMock) ResetServiceAccountCreds(ctx context.Context, id string) ApiResetServiceAccountCredsRequest {
	if mock.ResetServiceAccountCredsFunc == nil {
		panic("DefaultApiMock.ResetServiceAccountCredsFunc: method is nil but DefaultApi.ResetServiceAccountCreds was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockResetServiceAccountCreds.Lock()
	mock.calls.ResetServiceAccountCreds = append(mock.calls.ResetServiceAccountCreds, callInfo)
	mock.lockResetServiceAccountCreds.Unlock()
	return mock.ResetServiceAccountCredsFunc(ctx, id)
}

// ResetServiceAccountCredsCalls gets all the calls that were made to ResetServiceAccountCreds.
// Check the length with:
//     len(mockedDefaultApi.ResetServiceAccountCredsCalls())
func (mock *DefaultApiMock) ResetServiceAccountCredsCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockResetServiceAccountCreds.RLock()
	calls = mock.calls.ResetServiceAccountCreds
	mock.lockResetServiceAccountCreds.RUnlock()
	return calls
}

// ResetServiceAccountCredsExecute calls ResetServiceAccountCredsExecuteFunc.
func (mock *DefaultApiMock) ResetServiceAccountCredsExecute(r ApiResetServiceAccountCredsRequest) (ServiceAccount, *http.Response, GenericOpenAPIError) {
	if mock.ResetServiceAccountCredsExecuteFunc == nil {
		panic("DefaultApiMock.ResetServiceAccountCredsExecuteFunc: method is nil but DefaultApi.ResetServiceAccountCredsExecute was just called")
	}
	callInfo := struct {
		R ApiResetServiceAccountCredsRequest
	}{
		R: r,
	}
	mock.lockResetServiceAccountCredsExecute.Lock()
	mock.calls.ResetServiceAccountCredsExecute = append(mock.calls.ResetServiceAccountCredsExecute, callInfo)
	mock.lockResetServiceAccountCredsExecute.Unlock()
	return mock.ResetServiceAccountCredsExecuteFunc(r)
}

// ResetServiceAccountCredsExecuteCalls gets all the calls that were made to ResetServiceAccountCredsExecute.
// Check the length with:
//     len(mockedDefaultApi.ResetServiceAccountCredsExecuteCalls())
func (mock *DefaultApiMock) ResetServiceAccountCredsExecuteCalls() []struct {
	R ApiResetServiceAccountCredsRequest
} {
	var calls []struct {
		R ApiResetServiceAccountCredsRequest
	}
	mock.lockResetServiceAccountCredsExecute.RLock()
	calls = mock.calls.ResetServiceAccountCredsExecute
	mock.lockResetServiceAccountCredsExecute.RUnlock()
	return calls
}