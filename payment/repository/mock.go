// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package repository

import (
	"context"
	"sync"
)

// Ensure, that QuerierMock does implement Querier.
// If this is not the case, regenerate this file with moq.
var _ Querier = &QuerierMock{}

// QuerierMock is a mock implementation of Querier.
//
// 	func TestSomethingThatUsesQuerier(t *testing.T) {
//
// 		// make and configure a mocked Querier
// 		mockedQuerier := &QuerierMock{
// 			CreatePaymentFunc: func(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
// 				panic("mock out the CreatePayment method")
// 			},
// 			DiscardPaymentFunc: func(ctx context.Context, id int64) (int64, error) {
// 				panic("mock out the DiscardPayment method")
// 			},
// 			GetPaymentStatusByIDFunc: func(ctx context.Context, id int64) (ValidStatus, error) {
// 				panic("mock out the GetPaymentStatusByID method")
// 			},
// 			ListUserPaymentsByEmailFunc: func(ctx context.Context, arg ListUserPaymentsByEmailParams) ([]Payment, error) {
// 				panic("mock out the ListUserPaymentsByEmail method")
// 			},
// 			ListUserPaymentsByIDFunc: func(ctx context.Context, arg ListUserPaymentsByIDParams) ([]Payment, error) {
// 				panic("mock out the ListUserPaymentsByID method")
// 			},
// 			UpdatePaymentStatusFunc: func(ctx context.Context, arg UpdatePaymentStatusParams) (int64, error) {
// 				panic("mock out the UpdatePaymentStatus method")
// 			},
// 		}
//
// 		// use mockedQuerier in code that requires Querier
// 		// and then make assertions.
//
// 	}
type QuerierMock struct {
	// CreatePaymentFunc mocks the CreatePayment method.
	CreatePaymentFunc func(ctx context.Context, arg CreatePaymentParams) (Payment, error)

	// DiscardPaymentFunc mocks the DiscardPayment method.
	DiscardPaymentFunc func(ctx context.Context, id int64) (int64, error)

	// GetPaymentStatusByIDFunc mocks the GetPaymentStatusByID method.
	GetPaymentStatusByIDFunc func(ctx context.Context, id int64) (ValidStatus, error)

	// ListUserPaymentsByEmailFunc mocks the ListUserPaymentsByEmail method.
	ListUserPaymentsByEmailFunc func(ctx context.Context, arg ListUserPaymentsByEmailParams) ([]Payment, error)

	// ListUserPaymentsByIDFunc mocks the ListUserPaymentsByID method.
	ListUserPaymentsByIDFunc func(ctx context.Context, arg ListUserPaymentsByIDParams) ([]Payment, error)

	// UpdatePaymentStatusFunc mocks the UpdatePaymentStatus method.
	UpdatePaymentStatusFunc func(ctx context.Context, arg UpdatePaymentStatusParams) (int64, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreatePayment holds details about calls to the CreatePayment method.
		CreatePayment []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg CreatePaymentParams
		}
		// DiscardPayment holds details about calls to the DiscardPayment method.
		DiscardPayment []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
		}
		// GetPaymentStatusByID holds details about calls to the GetPaymentStatusByID method.
		GetPaymentStatusByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
		}
		// ListUserPaymentsByEmail holds details about calls to the ListUserPaymentsByEmail method.
		ListUserPaymentsByEmail []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg ListUserPaymentsByEmailParams
		}
		// ListUserPaymentsByID holds details about calls to the ListUserPaymentsByID method.
		ListUserPaymentsByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg ListUserPaymentsByIDParams
		}
		// UpdatePaymentStatus holds details about calls to the UpdatePaymentStatus method.
		UpdatePaymentStatus []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg UpdatePaymentStatusParams
		}
	}
	lockCreatePayment           sync.RWMutex
	lockDiscardPayment          sync.RWMutex
	lockGetPaymentStatusByID    sync.RWMutex
	lockListUserPaymentsByEmail sync.RWMutex
	lockListUserPaymentsByID    sync.RWMutex
	lockUpdatePaymentStatus     sync.RWMutex
}

// CreatePayment calls CreatePaymentFunc.
func (mock *QuerierMock) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	if mock.CreatePaymentFunc == nil {
		panic("QuerierMock.CreatePaymentFunc: method is nil but Querier.CreatePayment was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg CreatePaymentParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockCreatePayment.Lock()
	mock.calls.CreatePayment = append(mock.calls.CreatePayment, callInfo)
	mock.lockCreatePayment.Unlock()
	return mock.CreatePaymentFunc(ctx, arg)
}

// CreatePaymentCalls gets all the calls that were made to CreatePayment.
// Check the length with:
//     len(mockedQuerier.CreatePaymentCalls())
func (mock *QuerierMock) CreatePaymentCalls() []struct {
	Ctx context.Context
	Arg CreatePaymentParams
} {
	var calls []struct {
		Ctx context.Context
		Arg CreatePaymentParams
	}
	mock.lockCreatePayment.RLock()
	calls = mock.calls.CreatePayment
	mock.lockCreatePayment.RUnlock()
	return calls
}

// DiscardPayment calls DiscardPaymentFunc.
func (mock *QuerierMock) DiscardPayment(ctx context.Context, id int64) (int64, error) {
	if mock.DiscardPaymentFunc == nil {
		panic("QuerierMock.DiscardPaymentFunc: method is nil but Querier.DiscardPayment was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int64
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDiscardPayment.Lock()
	mock.calls.DiscardPayment = append(mock.calls.DiscardPayment, callInfo)
	mock.lockDiscardPayment.Unlock()
	return mock.DiscardPaymentFunc(ctx, id)
}

// DiscardPaymentCalls gets all the calls that were made to DiscardPayment.
// Check the length with:
//     len(mockedQuerier.DiscardPaymentCalls())
func (mock *QuerierMock) DiscardPaymentCalls() []struct {
	Ctx context.Context
	ID  int64
} {
	var calls []struct {
		Ctx context.Context
		ID  int64
	}
	mock.lockDiscardPayment.RLock()
	calls = mock.calls.DiscardPayment
	mock.lockDiscardPayment.RUnlock()
	return calls
}

// GetPaymentStatusByID calls GetPaymentStatusByIDFunc.
func (mock *QuerierMock) GetPaymentStatusByID(ctx context.Context, id int64) (ValidStatus, error) {
	if mock.GetPaymentStatusByIDFunc == nil {
		panic("QuerierMock.GetPaymentStatusByIDFunc: method is nil but Querier.GetPaymentStatusByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int64
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetPaymentStatusByID.Lock()
	mock.calls.GetPaymentStatusByID = append(mock.calls.GetPaymentStatusByID, callInfo)
	mock.lockGetPaymentStatusByID.Unlock()
	return mock.GetPaymentStatusByIDFunc(ctx, id)
}

// GetPaymentStatusByIDCalls gets all the calls that were made to GetPaymentStatusByID.
// Check the length with:
//     len(mockedQuerier.GetPaymentStatusByIDCalls())
func (mock *QuerierMock) GetPaymentStatusByIDCalls() []struct {
	Ctx context.Context
	ID  int64
} {
	var calls []struct {
		Ctx context.Context
		ID  int64
	}
	mock.lockGetPaymentStatusByID.RLock()
	calls = mock.calls.GetPaymentStatusByID
	mock.lockGetPaymentStatusByID.RUnlock()
	return calls
}

// ListUserPaymentsByEmail calls ListUserPaymentsByEmailFunc.
func (mock *QuerierMock) ListUserPaymentsByEmail(ctx context.Context, arg ListUserPaymentsByEmailParams) ([]Payment, error) {
	if mock.ListUserPaymentsByEmailFunc == nil {
		panic("QuerierMock.ListUserPaymentsByEmailFunc: method is nil but Querier.ListUserPaymentsByEmail was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg ListUserPaymentsByEmailParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockListUserPaymentsByEmail.Lock()
	mock.calls.ListUserPaymentsByEmail = append(mock.calls.ListUserPaymentsByEmail, callInfo)
	mock.lockListUserPaymentsByEmail.Unlock()
	return mock.ListUserPaymentsByEmailFunc(ctx, arg)
}

// ListUserPaymentsByEmailCalls gets all the calls that were made to ListUserPaymentsByEmail.
// Check the length with:
//     len(mockedQuerier.ListUserPaymentsByEmailCalls())
func (mock *QuerierMock) ListUserPaymentsByEmailCalls() []struct {
	Ctx context.Context
	Arg ListUserPaymentsByEmailParams
} {
	var calls []struct {
		Ctx context.Context
		Arg ListUserPaymentsByEmailParams
	}
	mock.lockListUserPaymentsByEmail.RLock()
	calls = mock.calls.ListUserPaymentsByEmail
	mock.lockListUserPaymentsByEmail.RUnlock()
	return calls
}

// ListUserPaymentsByID calls ListUserPaymentsByIDFunc.
func (mock *QuerierMock) ListUserPaymentsByID(ctx context.Context, arg ListUserPaymentsByIDParams) ([]Payment, error) {
	if mock.ListUserPaymentsByIDFunc == nil {
		panic("QuerierMock.ListUserPaymentsByIDFunc: method is nil but Querier.ListUserPaymentsByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg ListUserPaymentsByIDParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockListUserPaymentsByID.Lock()
	mock.calls.ListUserPaymentsByID = append(mock.calls.ListUserPaymentsByID, callInfo)
	mock.lockListUserPaymentsByID.Unlock()
	return mock.ListUserPaymentsByIDFunc(ctx, arg)
}

// ListUserPaymentsByIDCalls gets all the calls that were made to ListUserPaymentsByID.
// Check the length with:
//     len(mockedQuerier.ListUserPaymentsByIDCalls())
func (mock *QuerierMock) ListUserPaymentsByIDCalls() []struct {
	Ctx context.Context
	Arg ListUserPaymentsByIDParams
} {
	var calls []struct {
		Ctx context.Context
		Arg ListUserPaymentsByIDParams
	}
	mock.lockListUserPaymentsByID.RLock()
	calls = mock.calls.ListUserPaymentsByID
	mock.lockListUserPaymentsByID.RUnlock()
	return calls
}

// UpdatePaymentStatus calls UpdatePaymentStatusFunc.
func (mock *QuerierMock) UpdatePaymentStatus(ctx context.Context, arg UpdatePaymentStatusParams) (int64, error) {
	if mock.UpdatePaymentStatusFunc == nil {
		panic("QuerierMock.UpdatePaymentStatusFunc: method is nil but Querier.UpdatePaymentStatus was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg UpdatePaymentStatusParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockUpdatePaymentStatus.Lock()
	mock.calls.UpdatePaymentStatus = append(mock.calls.UpdatePaymentStatus, callInfo)
	mock.lockUpdatePaymentStatus.Unlock()
	return mock.UpdatePaymentStatusFunc(ctx, arg)
}

// UpdatePaymentStatusCalls gets all the calls that were made to UpdatePaymentStatus.
// Check the length with:
//     len(mockedQuerier.UpdatePaymentStatusCalls())
func (mock *QuerierMock) UpdatePaymentStatusCalls() []struct {
	Ctx context.Context
	Arg UpdatePaymentStatusParams
} {
	var calls []struct {
		Ctx context.Context
		Arg UpdatePaymentStatusParams
	}
	mock.lockUpdatePaymentStatus.RLock()
	calls = mock.calls.UpdatePaymentStatus
	mock.lockUpdatePaymentStatus.RUnlock()
	return calls
}
