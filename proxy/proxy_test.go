package proxy_test

import (
	"dataStructuresAndAlgorithms/proxy"
	"testing"
)

func TestUser_Pay(t *testing.T) {
	u := proxy.User{}
	err := u.Pay(20)
	if err != nil {
		t.Error(err)
	}
}
