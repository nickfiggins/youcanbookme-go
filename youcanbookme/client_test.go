package youcanbookme

import (
	"net/http"
	"testing"
)

func TestClient(t *testing.T) { 
	ycbm := NewClient(http.DefaultClient)
	account, _, err := ycbm.Accounts.Get("17f60ab5-958a-4c27-89e7-750801c03df8")
	
	if err != nil {
		t.Fatal(err)
	}

	if account.ID != "17f60ab5-958a-4c27-89e7-750801c03df8" {
		t.Fatalf("Account ID was incorrect, value %s", account.ID)
	}
}