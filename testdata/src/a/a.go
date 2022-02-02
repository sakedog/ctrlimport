package a

import (
	"net/http/cgi"
	"net/http/httptest" // want "import net/http/httptest is not allowed"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts" // want "import github.com/google/go-cmp/cmp/cmpopts is not allowed"
)

func main() {
	_ = cmpopts.IgnoreUnexported(struct{}{})
	cmp.Diff("", "")

	_ = cgi.Handler{}
	_ = httptest.Server{}
}
