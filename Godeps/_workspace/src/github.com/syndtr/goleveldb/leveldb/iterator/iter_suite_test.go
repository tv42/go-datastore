package iterator_test

import (
	"testing"

	"github.com/jbenet/go-datastore/Godeps/_workspace/src/github.com/syndtr/goleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
