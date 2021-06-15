##  Note: WIP
# YouCanBookMe Go API Wrapper

This is an API wrapper created for using the [YCBM](https://youcanbook.me/) api in Golang.

## Installation

This package can be installed easily by running:

```bash
go get github.com/nickfiggins/youcanbookme-go/youcanbookme@0.1.0
```

## Usage
First, please set an environment variable for YCBM_EMAIL and YCBM_KEY, which can be either your password or API key for YCBM.

```go
import (
    "fmt"
    ycbm "github.com/nickfiggins/youcanbookme-go/youcanbookme"
)

client := ycbm.NewClient() 
bookings, err := client.currentUser.Bookings.GetBookings(); if err != nil {
    fmt.Println(err)
}
fmt.Println(bookings)
```
## Current Endpoints

Currently, this API only supports GET endpoints for Account, Profile, and Bookings. This is a WIP and more support will be added in the future for other endpoints.
## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.