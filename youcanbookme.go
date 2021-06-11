package youcanbookme

import (
	"net/http"
	"fmt"
	"github.com/dghubble/sling"
	"os"
	"github.com/joho/godotenv"
)

// Probably gonna want these
const (

)

const apiBaseUrl = "https://api.youcanbook.me/"

const apiVersion = "v1"

const baseAddress = apiBaseUrl + apiVersion

type Client struct {
	sling *sling.Sling
	baseURL string
	CurrentUser Account
	Accounts *AccountsService
}

func NewClient(client *http.Client) *Client {
	godotenv.Load("auth.env")
	authClient := GetAuth(os.Getenv("email"), os.Getenv("password"))
	base := authClient.New().Client(client).Base(baseAddress + "/")
	currentUser, _, err := GetSelf(base); if err != nil {
		fmt.Println("Error", err)
	}
	return &Client{
		sling: base,
		baseURL: baseAddress,
		CurrentUser: currentUser,
		Accounts: newAccountsService(base.New()),
	}
}

func GetSelf(sling *sling.Sling) (Account, *http.Response, error) {
	apiError := new(APIError)
	var acc Account
	resp, err := sling.New().Get("me").Receive(&acc, apiError)
	acc.Sling = sling.New().Base(baseAddress + "/" + acc.ID + "/")
	acc.Profiles = newProfilesService(acc.Sling.New(), acc)
	return acc, resp, relevantError(err, *apiError)
}