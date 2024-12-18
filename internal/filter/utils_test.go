package filter

type MockFilter struct {
}

var mockFilterQuery = "mock_query"

func NewMockFilter() *MockFilter {
	return &MockFilter{}
}

func GetMockFilterQuery() string {
	return mockFilterQuery
}

func (f *MockFilter) BuildQuery() (string, error) {
	return mockFilterQuery, nil
}

func (f *MockFilter) Chain(other Filter) (Filter, error) {
	panic("not supported")
}
