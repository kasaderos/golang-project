package mock

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"route256/loms/internal/models"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// OrderProviderMock implements order.OrderProvider
type OrderProviderMock struct {
	t minimock.Tester

	funcGetOrderByID          func(ctx context.Context, orderID models.OrderID) (op1 *models.Order, err error)
	inspectFuncGetOrderByID   func(ctx context.Context, orderID models.OrderID)
	afterGetOrderByIDCounter  uint64
	beforeGetOrderByIDCounter uint64
	GetOrderByIDMock          mOrderProviderMockGetOrderByID
}

// NewOrderProviderMock returns a mock for order.OrderProvider
func NewOrderProviderMock(t minimock.Tester) *OrderProviderMock {
	m := &OrderProviderMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetOrderByIDMock = mOrderProviderMockGetOrderByID{mock: m}
	m.GetOrderByIDMock.callArgs = []*OrderProviderMockGetOrderByIDParams{}

	return m
}

type mOrderProviderMockGetOrderByID struct {
	mock               *OrderProviderMock
	defaultExpectation *OrderProviderMockGetOrderByIDExpectation
	expectations       []*OrderProviderMockGetOrderByIDExpectation

	callArgs []*OrderProviderMockGetOrderByIDParams
	mutex    sync.RWMutex
}

// OrderProviderMockGetOrderByIDExpectation specifies expectation struct of the OrderProvider.GetOrderByID
type OrderProviderMockGetOrderByIDExpectation struct {
	mock    *OrderProviderMock
	params  *OrderProviderMockGetOrderByIDParams
	results *OrderProviderMockGetOrderByIDResults
	Counter uint64
}

// OrderProviderMockGetOrderByIDParams contains parameters of the OrderProvider.GetOrderByID
type OrderProviderMockGetOrderByIDParams struct {
	ctx     context.Context
	orderID models.OrderID
}

// OrderProviderMockGetOrderByIDResults contains results of the OrderProvider.GetOrderByID
type OrderProviderMockGetOrderByIDResults struct {
	op1 *models.Order
	err error
}

// Expect sets up expected params for OrderProvider.GetOrderByID
func (mmGetOrderByID *mOrderProviderMockGetOrderByID) Expect(ctx context.Context, orderID models.OrderID) *mOrderProviderMockGetOrderByID {
	if mmGetOrderByID.mock.funcGetOrderByID != nil {
		mmGetOrderByID.mock.t.Fatalf("OrderProviderMock.GetOrderByID mock is already set by Set")
	}

	if mmGetOrderByID.defaultExpectation == nil {
		mmGetOrderByID.defaultExpectation = &OrderProviderMockGetOrderByIDExpectation{}
	}

	mmGetOrderByID.defaultExpectation.params = &OrderProviderMockGetOrderByIDParams{ctx, orderID}
	for _, e := range mmGetOrderByID.expectations {
		if minimock.Equal(e.params, mmGetOrderByID.defaultExpectation.params) {
			mmGetOrderByID.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetOrderByID.defaultExpectation.params)
		}
	}

	return mmGetOrderByID
}

// Inspect accepts an inspector function that has same arguments as the OrderProvider.GetOrderByID
func (mmGetOrderByID *mOrderProviderMockGetOrderByID) Inspect(f func(ctx context.Context, orderID models.OrderID)) *mOrderProviderMockGetOrderByID {
	if mmGetOrderByID.mock.inspectFuncGetOrderByID != nil {
		mmGetOrderByID.mock.t.Fatalf("Inspect function is already set for OrderProviderMock.GetOrderByID")
	}

	mmGetOrderByID.mock.inspectFuncGetOrderByID = f

	return mmGetOrderByID
}

// Return sets up results that will be returned by OrderProvider.GetOrderByID
func (mmGetOrderByID *mOrderProviderMockGetOrderByID) Return(op1 *models.Order, err error) *OrderProviderMock {
	if mmGetOrderByID.mock.funcGetOrderByID != nil {
		mmGetOrderByID.mock.t.Fatalf("OrderProviderMock.GetOrderByID mock is already set by Set")
	}

	if mmGetOrderByID.defaultExpectation == nil {
		mmGetOrderByID.defaultExpectation = &OrderProviderMockGetOrderByIDExpectation{mock: mmGetOrderByID.mock}
	}
	mmGetOrderByID.defaultExpectation.results = &OrderProviderMockGetOrderByIDResults{op1, err}
	return mmGetOrderByID.mock
}

// Set uses given function f to mock the OrderProvider.GetOrderByID method
func (mmGetOrderByID *mOrderProviderMockGetOrderByID) Set(f func(ctx context.Context, orderID models.OrderID) (op1 *models.Order, err error)) *OrderProviderMock {
	if mmGetOrderByID.defaultExpectation != nil {
		mmGetOrderByID.mock.t.Fatalf("Default expectation is already set for the OrderProvider.GetOrderByID method")
	}

	if len(mmGetOrderByID.expectations) > 0 {
		mmGetOrderByID.mock.t.Fatalf("Some expectations are already set for the OrderProvider.GetOrderByID method")
	}

	mmGetOrderByID.mock.funcGetOrderByID = f
	return mmGetOrderByID.mock
}

// When sets expectation for the OrderProvider.GetOrderByID which will trigger the result defined by the following
// Then helper
func (mmGetOrderByID *mOrderProviderMockGetOrderByID) When(ctx context.Context, orderID models.OrderID) *OrderProviderMockGetOrderByIDExpectation {
	if mmGetOrderByID.mock.funcGetOrderByID != nil {
		mmGetOrderByID.mock.t.Fatalf("OrderProviderMock.GetOrderByID mock is already set by Set")
	}

	expectation := &OrderProviderMockGetOrderByIDExpectation{
		mock:   mmGetOrderByID.mock,
		params: &OrderProviderMockGetOrderByIDParams{ctx, orderID},
	}
	mmGetOrderByID.expectations = append(mmGetOrderByID.expectations, expectation)
	return expectation
}

// Then sets up OrderProvider.GetOrderByID return parameters for the expectation previously defined by the When method
func (e *OrderProviderMockGetOrderByIDExpectation) Then(op1 *models.Order, err error) *OrderProviderMock {
	e.results = &OrderProviderMockGetOrderByIDResults{op1, err}
	return e.mock
}

// GetOrderByID implements order.OrderProvider
func (mmGetOrderByID *OrderProviderMock) GetOrderByID(ctx context.Context, orderID models.OrderID) (op1 *models.Order, err error) {
	mm_atomic.AddUint64(&mmGetOrderByID.beforeGetOrderByIDCounter, 1)
	defer mm_atomic.AddUint64(&mmGetOrderByID.afterGetOrderByIDCounter, 1)

	if mmGetOrderByID.inspectFuncGetOrderByID != nil {
		mmGetOrderByID.inspectFuncGetOrderByID(ctx, orderID)
	}

	mm_params := &OrderProviderMockGetOrderByIDParams{ctx, orderID}

	// Record call args
	mmGetOrderByID.GetOrderByIDMock.mutex.Lock()
	mmGetOrderByID.GetOrderByIDMock.callArgs = append(mmGetOrderByID.GetOrderByIDMock.callArgs, mm_params)
	mmGetOrderByID.GetOrderByIDMock.mutex.Unlock()

	for _, e := range mmGetOrderByID.GetOrderByIDMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.op1, e.results.err
		}
	}

	if mmGetOrderByID.GetOrderByIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetOrderByID.GetOrderByIDMock.defaultExpectation.Counter, 1)
		mm_want := mmGetOrderByID.GetOrderByIDMock.defaultExpectation.params
		mm_got := OrderProviderMockGetOrderByIDParams{ctx, orderID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetOrderByID.t.Errorf("OrderProviderMock.GetOrderByID got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetOrderByID.GetOrderByIDMock.defaultExpectation.results
		if mm_results == nil {
			mmGetOrderByID.t.Fatal("No results are set for the OrderProviderMock.GetOrderByID")
		}
		return (*mm_results).op1, (*mm_results).err
	}
	if mmGetOrderByID.funcGetOrderByID != nil {
		return mmGetOrderByID.funcGetOrderByID(ctx, orderID)
	}
	mmGetOrderByID.t.Fatalf("Unexpected call to OrderProviderMock.GetOrderByID. %v %v", ctx, orderID)
	return
}

// GetOrderByIDAfterCounter returns a count of finished OrderProviderMock.GetOrderByID invocations
func (mmGetOrderByID *OrderProviderMock) GetOrderByIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetOrderByID.afterGetOrderByIDCounter)
}

// GetOrderByIDBeforeCounter returns a count of OrderProviderMock.GetOrderByID invocations
func (mmGetOrderByID *OrderProviderMock) GetOrderByIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetOrderByID.beforeGetOrderByIDCounter)
}

// Calls returns a list of arguments used in each call to OrderProviderMock.GetOrderByID.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetOrderByID *mOrderProviderMockGetOrderByID) Calls() []*OrderProviderMockGetOrderByIDParams {
	mmGetOrderByID.mutex.RLock()

	argCopy := make([]*OrderProviderMockGetOrderByIDParams, len(mmGetOrderByID.callArgs))
	copy(argCopy, mmGetOrderByID.callArgs)

	mmGetOrderByID.mutex.RUnlock()

	return argCopy
}

// MinimockGetOrderByIDDone returns true if the count of the GetOrderByID invocations corresponds
// the number of defined expectations
func (m *OrderProviderMock) MinimockGetOrderByIDDone() bool {
	for _, e := range m.GetOrderByIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetOrderByIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetOrderByIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetOrderByID != nil && mm_atomic.LoadUint64(&m.afterGetOrderByIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetOrderByIDInspect logs each unmet expectation
func (m *OrderProviderMock) MinimockGetOrderByIDInspect() {
	for _, e := range m.GetOrderByIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to OrderProviderMock.GetOrderByID with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetOrderByIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetOrderByIDCounter) < 1 {
		if m.GetOrderByIDMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to OrderProviderMock.GetOrderByID")
		} else {
			m.t.Errorf("Expected call to OrderProviderMock.GetOrderByID with params: %#v", *m.GetOrderByIDMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetOrderByID != nil && mm_atomic.LoadUint64(&m.afterGetOrderByIDCounter) < 1 {
		m.t.Error("Expected call to OrderProviderMock.GetOrderByID")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *OrderProviderMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetOrderByIDInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *OrderProviderMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *OrderProviderMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetOrderByIDDone()
}
