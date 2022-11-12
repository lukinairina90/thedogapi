package main

import (
	"context"
	"fmt"
	"golangNinga/work_with_JSON_API_for_coincap/dog/thedogapi"
	"log"
)

// [GET] https://api.thedogapi.com/v1/votes - LIST
// [GET] https://api.thedogapi.com/v1/votes/{id} - GET
// [POST] https://api.thedogapi.com/v1/votes - без ID и с body - CREATE
// [PUT] https://api.thedogapi.com/v1/votes/{id} - c ID и с body - UPDATE
// [DELETE] https://api.thedogapi.com/v1/votes/{id} - c ID - DELETE

const APIKey = "live_cVcXvz3CNxPuHBdAGv3DOvubhk8lhY1hWf8nSURPsqQ8x9oZj9kHwYbH0Hfty2IW"
const BaseURL = "https://api.thedogapi.com/v1"

func main() {
	dogAPIClient := thedogapi.NewClient(BaseURL, APIKey)

	ctx := context.Background()

	list, err := dogAPIClient.List(ctx, &thedogapi.ListParams{
		Limit:     20,
		Page:      1,
		DescOrder: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	if err := dogAPIClient.Vote(ctx, list[0].ID, true); err != nil {
		log.Fatal(err)
	}

	fmt.Println(list)
}
