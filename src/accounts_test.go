package tradeking

import (
	"testing"
)

func SetupAccounts() {

}

func TearDownAccounts() {

}

func Test_buildEndPoint(t *testing.T) {
    SetupAccounts()
	defer TearDownAccounts()
    t.Errorf("%s != %s; values should be the same", controlEndPoint, testEndPoint)
}
