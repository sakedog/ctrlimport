package b_test

import (
	"net/http/cgi"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestHoge(t *testing.T) {
	_ = cmpopts.IgnoreUnexported(struct{}{})
	cmp.Diff("", "")

	_ = cgi.Handler{}
	_ = httptest.Server{}
}
