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

// ReserveCancellerMock implements order.ReserveCanceller
type ReserveCancellerMock struct {
	t minimock.Tester

	funcReserveCancel          func(ctx context.Context, items []models.ItemOrderInfo) (err error)
	inspectFuncReserveCancel   func(ctx context.Context, items []models.ItemOrderInfo)
	afterReserveCancelCounter  uint64
	beforeReserveCancelCounter uint64
	ReserveCancelMock          mReserveCancellerMockReserveCancel
}

// NewReserveCancellerMock returns a mock for order.ReserveCanceller
func NewReserveCancellerMock(t minimock.Tester) *ReserveCancellerMock {
	m := &ReserveCancellerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ReserveCancelMock = mReserveCancellerMockReserveCancel{mock: m}
	m.ReserveCancelMock.callArgs = []*ReserveCancellerMockReserveCancelParams{}

	return m
}

type mReserveCancellerMockReserveCancel struct {
	mock               *ReserveCancellerMock
	defaultExpectation *ReserveCancellerMockReserveCancelExpectation
	expectations       []*ReserveCancellerMockReserveCancelExpectation

	callArgs []*ReserveCancellerMockReserveCancelParams
	mutex    sync.RWMutex
}

// ReserveCancellerMockReserveCancelExpectation specifies expectation struct of the ReserveCanceller.ReserveCancel
type ReserveCancellerMockReserveCancelExpectation struct {
	mock    *ReserveCancellerMock
	params  *ReserveCancellerMockReserveCancelParams
	results *ReserveCancellerMockReserveCancelResults
	Counter uint64
}

// ReserveCancellerMockReserveCancelParams contains parameters of the ReserveCanceller.ReserveCancel
type ReserveCancellerMockReserveCancelParams struct {
	ctx   context.Context
	items []models.ItemOrderInfo
}

// ReserveCancellerMockReserveCancelResults contains results of the ReserveCanceller.ReserveCancel
type ReserveCancellerMockReserveCancelResults struct {
	err error
}

// Expect sets up expected params for ReserveCanceller.ReserveCancel
func (mmReserveCancel *mReserveCancellerMockReserveCancel) Expect(ctx context.Context, items []models.ItemOrderInfo) *mReserveCancellerMockReserveCancel {
	if mmReserveCancel.mock.funcReserveCancel != nil {
		mmReserveCancel.mock.t.Fatalf("ReserveCancellerMock.ReserveCancel mock is already set by Set")
	}

	if mmReserveCancel.defaultExpectation == nil {
		mmReserveCancel.defaultExpectation = &ReserveCancellerMockReserveCancelExpectation{}
	}

	mmReserveCancel.defaultExpectation.params = &ReserveCancellerMockReserveCancelParams{ctx, items}
	for _, e := range mmReserveCancel.expectations {
		if minimock.Equal(e.params, mmReserveCancel.defaultExpectation.params) {
			mmReserveCancel.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmReserveCancel.defaultExpectation.params)
		}
	}

	return mmReserveCancel
}

// Inspect accepts an inspector function that has same arguments as the ReserveCanceller.ReserveCancel
func (mmReserveCancel *mReserveCancellerMockReserveCancel) Inspect(f func(ctx context.Context, items []models.ItemOrderInfo)) *mReserveCancellerMockReserveCancel {
	if mmReserveCancel.mock.inspectFuncReserveCancel != nil {
		mmReserveCancel.mock.t.Fatalf("Inspect function is already set for ReserveCancellerMock.ReserveCancel")
	}

	mmReserveCancel.mock.inspectFuncReserveCancel = f

	return mmReserveCancel
}

// Return sets up results that will be returned by ReserveCanceller.ReserveCancel
func (mmReserveCancel *mReserveCancellerMockReserveCancel) Return(err error) *ReserveCancellerMock {
	if mmReserveCancel.mock.funcReserveCancel != nil {
		mmReserveCancel.mock.t.Fatalf("ReserveCancellerMock.ReserveCancel mock is already set by Set")
	}

	if mmReserveCancel.defaultExpectation == nil {
		mmReserveCancel.defaultExpectation = &ReserveCancellerMockReserveCancelExpectation{mock: mmReserveCancel.mock}
	}
	mmReserveCancel.defaultExpectation.results = &ReserveCancellerMockReserveCancelResults{err}
	return mmReserveCancel.mock
}

// Set uses given function f to mock the ReserveCanceller.ReserveCancel method
func (mmReserveCancel *mReserveCancellerMockReserveCancel) Set(f func(ctx context.Context, items []models.ItemOrderInfo) (err error)) *ReserveCancellerMock {
	if mmReserveCancel.defaultExpectation != nil {
		mmReserveCancel.mock.t.Fatalf("Default expectation is already set for the ReserveCanceller.ReserveCancel method")
	}

	if len(mmReserveCancel.expectations) > 0 {
		mmReserveCancel.mock.t.Fatalf("Some expectations are already set for the ReserveCanceller.ReserveCancel method")
	}

	mmReserveCancel.mock.funcReserveCancel = f
	return mmReserveCancel.mock
}

// When sets expectation for the ReserveCanceller.ReserveCancel which will trigger the result defined by the following
// Then helper
func (mmReserveCancel *mReserveCancellerMockReserveCancel) When(ctx context.Context, items []models.ItemOrderInfo) *ReserveCancellerMockReserveCancelExpectation {
	if mmReserveCancel.mock.funcReserveCancel != nil {
		mmReserveCancel.mock.t.Fatalf("ReserveCancellerMock.ReserveCancel mock is already set by Set")
	}

	expectation := &ReserveCancellerMockReserveCancelExpectation{
		mock:   mmReserveCancel.mock,
		params: &ReserveCancellerMockReserveCancelParams{ctx, items},
	}
	mmReserveCancel.expectations = append(mmReserveCancel.expectations, expectation)
	return expectation
}

// Then sets up ReserveCanceller.ReserveCancel return parameters for the expectation previously defined by the When method
func (e *ReserveCancellerMockReserveCancelExpectation) Then(err error) *ReserveCancellerMock {
	e.results = &ReserveCancellerMockReserveCancelResults{err}
	return e.mock
}

// ReserveCancel implements order.ReserveCanceller
func (mmReserveCancel *ReserveCancellerMock) ReserveCancel(ctx context.Context, items []models.ItemOrderInfo) (err error) {
	mm_atomic.AddUint64(&mmReserveCancel.beforeReserveCancelCounter, 1)
	defer mm_atomic.AddUint64(&mmReserveCancel.afterReserveCancelCounter, 1)

	if mmReserveCancel.inspectFuncReserveCancel != nil {
		mmReserveCancel.inspectFuncReserveCancel(ctx, items)
	}

	mm_params := &ReserveCancellerMockReserveCancelParams{ctx, items}

	// Record call args
	mmReserveCancel.ReserveCancelMock.mutex.Lock()
	mmReserveCancel.ReserveCancelMock.callArgs = append(mmReserveCancel.ReserveCancelMock.callArgs, mm_params)
	mmReserveCancel.ReserveCancelMock.mutex.Unlock()

	for _, e := range mmReserveCancel.ReserveCancelMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmReserveCancel.ReserveCancelMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmReserveCancel.ReserveCancelMock.defaultExpectation.Counter, 1)
		mm_want := mmReserveCancel.ReserveCancelMock.defaultExpectation.params
		mm_got := ReserveCancellerMockReserveCancelParams{ctx, items}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmReserveCancel.t.Errorf("ReserveCancellerMock.ReserveCancel got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmReserveCancel.ReserveCancelMock.defaultExpectation.results
		if mm_results == nil {
			mmReserveCancel.t.Fatal("No results are set for the ReserveCancellerMock.ReserveCancel")
		}
		return (*mm_results).err
	}
	if mmReserveCancel.funcReserveCancel != nil {
		return mmReserveCancel.funcReserveCancel(ctx, items)
	}
	mmReserveCancel.t.Fatalf("Unexpected call to ReserveCancellerMock.ReserveCancel. %v %v", ctx, items)
	return
}

// ReserveCancelAfterCounter returns a count of finished ReserveCancellerMock.ReserveCancel invocations
func (mmReserveCancel *ReserveCancellerMock) ReserveCancelAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmReserveCancel.afterReserveCancelCounter)
}

// ReserveCancelBeforeCounter returns a count of ReserveCancellerMock.ReserveCancel invocations
func (mmReserveCancel *ReserveCancellerMock) ReserveCancelBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmReserveCancel.beforeReserveCancelCounter)
}

// Calls returns a list of arguments used in each call to ReserveCancellerMock.ReserveCancel.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmReserveCancel *mReserveCancellerMockReserveCancel) Calls() []*ReserveCancellerMockReserveCancelParams {
	mmReserveCancel.mutex.RLock()

	argCopy := make([]*ReserveCancellerMockReserveCancelParams, len(mmReserveCancel.callArgs))
	copy(argCopy, mmReserveCancel.callArgs)

	mmReserveCancel.mutex.RUnlock()

	return argCopy
}

// MinimockReserveCancelDone returns true if the count of the ReserveCancel invocations corresponds
// the number of defined expectations
func (m *ReserveCancellerMock) MinimockReserveCancelDone() bool {
	for _, e := range m.ReserveCancelMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ReserveCancelMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterReserveCancelCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcReserveCancel != nil && mm_atomic.LoadUint64(&m.afterReserveCancelCounter) < 1 {
		return false
	}
	return true
}

// MinimockReserveCancelInspect logs each unmet expectation
func (m *ReserveCancellerMock) MinimockReserveCancelInspect() {
	for _, e := range m.ReserveCancelMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ReserveCancellerMock.ReserveCancel with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ReserveCancelMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterReserveCancelCounter) < 1 {
		if m.ReserveCancelMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ReserveCancellerMock.ReserveCancel")
		} else {
			m.t.Errorf("Expected call to ReserveCancellerMock.ReserveCancel with params: %#v", *m.ReserveCancelMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcReserveCancel != nil && mm_atomic.LoadUint64(&m.afterReserveCancelCounter) < 1 {
		m.t.Error("Expected call to ReserveCancellerMock.ReserveCancel")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ReserveCancellerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockReserveCancelInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ReserveCancellerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ReserveCancellerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockReserveCancelDone()
}
