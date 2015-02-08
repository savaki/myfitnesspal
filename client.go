package myfitnesspal

import "net/http"

const (
	Codebase     = "https://www.myfitnesspal.com/"
	LoginUrl     = Codebase + "account/login"
	FoodDiaryUrl = Codebase + "food/diary/"
)

type Client struct {
	client   *http.Client
	username string
}

func New(username, password string) (*Client, error) {
	client, err := login(username, password)
	if err != nil {
		return nil, err
	}

	return &Client{
		client:   client,
		username: username,
	}, nil
}
