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

// StocksReserverMock implements order.StocksReserver
type StocksReserverMock struct {
	t minimock.Tester

	funcReserveStocks          func(ctx context.Context, items []models.ItemOrderInfo) (err error)
	inspectFuncReserveStocks   func(ctx context.Context, items []models.ItemOrderInfo)
	afterReserveStocksCounter  uint64
	beforeReserveStocksCounter uint64
	ReserveStocksMock          mStocksReserverMockReserveStocks
}

// NewStocksReserverMock returns a mock for order.StocksReserver
func NewStocksReserverMock(t minimock.Tester) *StocksReserverMock {
	m := &StocksReserverMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ReserveStocksMock = mStocksReserverMockReserveStocks{mock: m}
	m.ReserveStocksMock.callArgs = []*StocksReserverMockReserveStocksParams{}

	return m
}

type mStocksReserverMockReserveStocks struct {
	mock               *StocksReserverMock
	defaultExpectation *StocksReserverMockReserveStocksExpectation
	expectations       []*StocksReserverMockReserveStocksExpectation

	callArgs []*StocksReserverMockReserveStocksParams
	mutex    sync.RWMutex
}

// StocksReserverMockReserveStocksExpectation specifies expectation struct of the StocksReserver.ReserveStocks
type StocksReserverMockReserveStocksExpectation struct {
	mock    *StocksReserverMock
	params  *StocksReserverMockReserveStocksParams
	results *StocksReserverMockReserveStocksResults
	Counter uint64
}

// StocksReserverMockReserveStocksParams contains parameters of the StocksReserver.ReserveStocks
type StocksReserverMockReserveStocksParams struct {
	ctx   context.Context
	items []models.ItemOrderInfo
}

// StocksReserverMockReserveStocksResults contains results of the StocksReserver.ReserveStocks
type StocksReserverMockReserveStocksResults struct {
	err error
}

// Expect sets up expected params for StocksReserver.ReserveStocks
func (mmReserveStocks *mStocksReserverMockReserveStocks) Expect(ctx context.Context, items []models.ItemOrderInfo) *mStocksReserverMockReserveStocks {
	if mmReserveStocks.mock.funcReserveStocks != nil {
		mmReserveStocks.mock.t.Fatalf("StocksReserverMock.ReserveStocks mock is already set by Set")
	}

	if mmReserveStocks.defaultExpectation == nil {
		mmReserveStocks.defaultExpectation = &StocksReserverMockReserveStocksExpectation{}
	}

	mmReserveStocks.defaultExpectation.params = &StocksReserverMockReserveStocksParams{ctx, items}
	for _, e := range mmReserveStocks.expectations {
		if minimock.Equal(e.params, mmReserveStocks.defaultExpectation.params) {
			mmReserveStocks.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmReserveStocks.defaultExpectation.params)
		}
	}

	return mmReserveStocks
}

// Inspect accepts an inspector function that has same arguments as the StocksReserver.ReserveStocks
func (mmReserveStocks *mStocksReserverMockReserveStocks) Inspect(f func(ctx context.Context, items []models.ItemOrderInfo)) *mStocksReserverMockReserveStocks {
	if mmReserveStocks.mock.inspectFuncReserveStocks != nil {
		mmReserveStocks.mock.t.Fatalf("Inspect function is already set for StocksReserverMock.ReserveStocks")
	}

	mmReserveStocks.mock.inspectFuncReserveStocks = f

	return mmReserveStocks
}

// Return sets up results that will be returned by StocksReserver.ReserveStocks
func (mmReserveStocks *mStocksReserverMockReserveStocks) Return(err error) *StocksReserverMock {
	if mmReserveStocks.mock.funcReserveStocks != nil {
		mmReserveStocks.mock.t.Fatalf("StocksReserverMock.ReserveStocks mock is already set by Set")
	}

	if mmReserveStocks.defaultExpectation == nil {
		mmReserveStocks.defaultExpectation = &StocksReserverMockReserveStocksExpectation{mock: mmReserveStocks.mock}
	}
	mmReserveStocks.defaultExpectation.results = &StocksReserverMockReserveStocksResults{err}
	return mmReserveStocks.mock
}

// Set uses given function f to mock the StocksReserver.ReserveStocks method
func (mmReserveStocks *mStocksReserverMockReserveStocks) Set(f func(ctx context.Context, items []models.ItemOrderInfo) (err error)) *StocksReserverMock {
	if mmReserveStocks.defaultExpectation != nil {
		mmReserveStocks.mock.t.Fatalf("Default expectation is already set for the StocksReserver.ReserveStocks method")
	}

	if len(mmReserveStocks.expectations) > 0 {
		mmReserveStocks.mock.t.Fatalf("Some expectations are already set for the StocksReserver.ReserveStocks method")
	}

	mmReserveStocks.mock.funcReserveStocks = f
	return mmReserveStocks.mock
}

// When sets expectation for the StocksReserver.ReserveStocks which will trigger the result defined by the following
// Then helper
func (mmReserveStocks *mStocksReserverMockReserveStocks) When(ctx context.Context, items []models.ItemOrderInfo) *StocksReserverMockReserveStocksExpectation {
	if mmReserveStocks.mock.funcReserveStocks != nil {
		mmReserveStocks.mock.t.Fatalf("StocksReserverMock.ReserveStocks mock is already set by Set")
	}

	expectation := &StocksReserverMockReserveStocksExpectation{
		mock:   mmReserveStocks.mock,
		params: &StocksReserverMockReserveStocksParams{ctx, items},
	}
	mmReserveStocks.expectations = append(mmReserveStocks.expectations, expectation)
	return expectation
}

// Then sets up StocksReserver.ReserveStocks return parameters for the expectation previously defined by the When method
func (e *StocksReserverMockReserveStocksExpectation) Then(err error) *StocksReserverMock {
	e.results = &StocksReserverMockReserveStocksResults{err}
	return e.mock
}

// ReserveStocks implements order.StocksReserver
func (mmReserveStocks *StocksReserverMock) ReserveStocks(ctx context.Context, items []models.ItemOrderInfo) (err error) {
	mm_atomic.AddUint64(&mmReserveStocks.beforeReserveStocksCounter, 1)
	defer mm_atomic.AddUint64(&mmReserveStocks.afterReserveStocksCounter, 1)

	if mmReserveStocks.inspectFuncReserveStocks != nil {
		mmReserveStocks.inspectFuncReserveStocks(ctx, items)
	}

	mm_params := &StocksReserverMockReserveStocksParams{ctx, items}

	// Record call args
	mmReserveStocks.ReserveStocksMock.mutex.Lock()
	mmReserveStocks.ReserveStocksMock.callArgs = append(mmReserveStocks.ReserveStocksMock.callArgs, mm_params)
	mmReserveStocks.ReserveStocksMock.mutex.Unlock()

	for _, e := range mmReserveStocks.ReserveStocksMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmReserveStocks.ReserveStocksMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmReserveStocks.ReserveStocksMock.defaultExpectation.Counter, 1)
		mm_want := mmReserveStocks.ReserveStocksMock.defaultExpectation.params
		mm_got := StocksReserverMockReserveStocksParams{ctx, items}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmReserveStocks.t.Errorf("StocksReserverMock.ReserveStocks got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmReserveStocks.ReserveStocksMock.defaultExpectation.results
		if mm_results == nil {
			mmReserveStocks.t.Fatal("No results are set for the StocksReserverMock.ReserveStocks")
		}
		return (*mm_results).err
	}
	if mmReserveStocks.funcReserveStocks != nil {
		return mmReserveStocks.funcReserveStocks(ctx, items)
	}
	mmReserveStocks.t.Fatalf("Unexpected call to StocksReserverMock.ReserveStocks. %v %v", ctx, items)
	return
}

// ReserveStocksAfterCounter returns a count of finished StocksReserverMock.ReserveStocks invocations
func (mmReserveStocks *StocksReserverMock) ReserveStocksAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmReserveStocks.afterReserveStocksCounter)
}

// ReserveStocksBeforeCounter returns a count of StocksReserverMock.ReserveStocks invocations
func (mmReserveStocks *StocksReserverMock) ReserveStocksBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmReserveStocks.beforeReserveStocksCounter)
}

// Calls returns a list of arguments used in each call to StocksReserverMock.ReserveStocks.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmReserveStocks *mStocksReserverMockReserveStocks) Calls() []*StocksReserverMockReserveStocksParams {
	mmReserveStocks.mutex.RLock()

	argCopy := make([]*StocksReserverMockReserveStocksParams, len(mmReserveStocks.callArgs))
	copy(argCopy, mmReserveStocks.callArgs)

	mmReserveStocks.mutex.RUnlock()

	return argCopy
}

// MinimockReserveStocksDone returns true if the count of the ReserveStocks invocations corresponds
// the number of defined expectations
func (m *StocksReserverMock) MinimockReserveStocksDone() bool {
	for _, e := range m.ReserveStocksMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ReserveStocksMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterReserveStocksCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcReserveStocks != nil && mm_atomic.LoadUint64(&m.afterReserveStocksCounter) < 1 {
		return false
	}
	return true
}

// MinimockReserveStocksInspect logs each unmet expectation
func (m *StocksReserverMock) MinimockReserveStocksInspect() {
	for _, e := range m.ReserveStocksMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StocksReserverMock.ReserveStocks with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ReserveStocksMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterReserveStocksCounter) < 1 {
		if m.ReserveStocksMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StocksReserverMock.ReserveStocks")
		} else {
			m.t.Errorf("Expected call to StocksReserverMock.ReserveStocks with params: %#v", *m.ReserveStocksMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcReserveStocks != nil && mm_atomic.LoadUint64(&m.afterReserveStocksCounter) < 1 {
		m.t.Error("Expected call to StocksReserverMock.ReserveStocks")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *StocksReserverMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockReserveStocksInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *StocksReserverMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *StocksReserverMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockReserveStocksDone()
}
