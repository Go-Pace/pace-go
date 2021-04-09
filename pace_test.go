package pace_test

import (
	"testing"

	"github.com/gopace/pace-go"
)

func TestInit(t *testing.T) {
	gopace := &pace.GoPaceRequest{}
	gopace.Init(&pace.AuthRequest{
		PrivateKey: "hello",
	})
	if gopace.Token == "" {
		t.Error("No token present and auth failed")
	}
	t.Logf("Go Pace Auth Token %s", gopace.Token)
}

func BenchmarkInit(b *testing.B) {
	gopace := &pace.GoPaceRequest{}
	for i := 0; i < b.N; i++ {
		gopace.Init(&pace.AuthRequest{
			PrivateKey: "hello",
		})
	}
}
