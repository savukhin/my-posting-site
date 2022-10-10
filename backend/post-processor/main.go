package main

import (
	"context"
	"my-posting-site/backend/post-processor/controllers"
)

func main() {
	controllers.Consume(context.Background())
}
