package main

import (
	"github.com/NygmaC/streamming-video/stream-reader/config"
	"github.com/NygmaC/streamming-video/stream-reader/internal/admin"
	"github.com/NygmaC/streamming-video/stream-reader/internal/consumer"
	"github.com/NygmaC/streamming-video/stream-reader/internal/producer"
	"github.com/NygmaC/streamming-video/stream-reader/internal/reader"
	"github.com/NygmaC/streamming-video/stream-reader/pkg/server"
)

func main() {
	config.Init()
	server.Init()
	reader.Init()
	producer.Init()
	admin.Init()
	consumer.Init()

}
