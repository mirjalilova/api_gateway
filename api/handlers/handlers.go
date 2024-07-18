package handlers

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/mirjalilova/api_gateway/genproto/memory"
	"github.com/mirjalilova/api_gateway/kafka"
	"google.golang.org/grpc"
)

type Handlers struct {
	Comment  memory.CommentServiceClient
	Media    memory.MediaServiceClient
	Memory   memory.MemoryServiceClient
	Share    memory.ShareServiceClient
	RDB      *redis.Client
	Producer kafka.KafkaProducer
}

func NewHandler(memoryConn *grpc.ClientConn, rdb *redis.Client) *Handlers {
	pr, err := kafka.NewKafkaProducer([]string{"localhost:9092"})
	if err != nil {
		log.Fatal(err)
	}
	return &Handlers{
		Comment:  memory.NewCommentServiceClient(memoryConn),
		Media:    memory.NewMediaServiceClient(memoryConn),
		Memory:   memory.NewMemoryServiceClient(memoryConn),
		Share:    memory.NewShareServiceClient(memoryConn),
		RDB:      rdb,
		Producer: pr,
	}
}
