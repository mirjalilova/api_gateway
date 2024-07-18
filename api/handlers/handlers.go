package handlers

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/mirjalilova/api_gateway/genproto/auth"
	"github.com/mirjalilova/api_gateway/genproto/memory"
	"github.com/mirjalilova/api_gateway/genproto/timeline"
	"github.com/mirjalilova/api_gateway/kafka"
	"google.golang.org/grpc"
)

type Handlers struct {
	User        auth.UserServiceClient
	Comment     memory.CommentServiceClient
	Media       memory.MediaServiceClient
	Memory      memory.MemoryServiceClient
	Share       memory.ShareServiceClient
	Millistones timeline.MillistonesServiceClient
	RDB         *redis.Client
	Producer    kafka.KafkaProducer
}

func NewHandler(memoryConn, timelineConn, userConn *grpc.ClientConn, rdb *redis.Client) *Handlers {
	pr, err := kafka.NewKafkaProducer([]string{"localhost:9092"})
	if err != nil {
		log.Fatal(err)
	}
	return &Handlers{
		User:        auth.NewUserServiceClient(userConn),
		Comment:     memory.NewCommentServiceClient(memoryConn),
		Media:       memory.NewMediaServiceClient(memoryConn),
		Memory:      memory.NewMemoryServiceClient(memoryConn),
		Share:       memory.NewShareServiceClient(memoryConn),
		Millistones: timeline.NewMillistonesServiceClient(timelineConn),
		RDB:         rdb,
		Producer:    pr,
	}
}
