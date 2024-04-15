package mock

func NewMock() Mock {
	return Mock{
		Calls:         make(map[string]Call),
		ExpectedCalls: make(map[string]Values),
	}
}

type Values []any

type Call struct {
	// ok is true if the method was called
	Ok bool
	// args are the arguments passed to the method
	Args Values
}

type Mock struct {
	// Calls
	Calls map[string]Call

	// Expected Calls
	// - key: method name
	// - value: args to return (output)
	ExpectedCalls map[string]Values
}

func (m *Mock) On(method string, args ...any) {
	m.ExpectedCalls[method] = args
}

func (m *Mock) Called(method string, args Values) (output Values) {
	// registry
	m.Calls[method] = Call{Ok: true, Args: args}

	// expected call
	output = m.ExpectedCalls[method]
	return
}