package youcanbookme

import (
	"net/http"
	"os"
	"testing"
)

func TestClient(t *testing.T) { 
	ycbm := NewTestClient(http.DefaultClient)
	testAccId := os.Getenv("TEST_ACCOUNT_ID")
	account, _, err := ycbm.Accounts.Get(testAccId)
	
	if err != nil {
		t.Fatal(err)
	}

	if account.ID != testAccId {
		t.Fatalf("Account ID was incorrect, value %s", account.ID)
	}
}