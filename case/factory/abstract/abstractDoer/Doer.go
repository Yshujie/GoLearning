package abstractDoer

import (
	"net/http"
	"net/http/httptest"
)

// NewHTTPClient 和 NewMockHTTPClient 都返回了同一个接口类型 Doer，这使得二者可以互换使用。
// 当你想测试一段调用了 Doer 接口 Do 方法的代码时，这一点特别有用。
// 因为你可以使用一个 Mock 的 HTTP 客户端，从而避免了调用真实外部接口可能带来的失败。

// We define a Doer interface, that has the method signature
// of the 'http.Client' structs `Do` method

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// This gives us a regular HTTP client from the `net/http` package

func NewHttpClient() Doer {
	return &http.Client{}
}

type mockHTTPClient struct{}

func (*mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	// The `NewRecorder` method of the httptest package gives us
	// a new mock request generator
	res := httptest.NewRecorder()

	// calling the `Result` method gives us
	// the default empty *http.Response object
	return res.Result(), nil
}

// This gives us a mock HTTP client, which returns
// an empty response for any request sent to it

func NewNockHTTPClient() Doer {
	return &mockHTTPClient{}
}
