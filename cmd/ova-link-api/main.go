package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ozonva/ova-link-api/internal/api"

	linkAPI "github.com/ozonva/ova-link-api/pkg/ova-link-api"

	"google.golang.org/grpc"

	"github.com/ozonva/ova-link-api/internal/config"
	"github.com/ozonva/ova-link-api/internal/link"
	"github.com/ozonva/ova-link-api/internal/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "config" {
		err := config.UpdateConfig("./configs/config.json", config.InfiniteUpdater)
		if err != nil {
			log.Fatal(err)
		}
	}

	if len(args) > 0 && args[0] == "server" {
		const (
			grpcPort = ":82"
		)
		listen, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		linkAPI.RegisterLinkAPIServer(s, api.NewLinkAPI())

		if err := s.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		return
	}

	projectName := "ova-link-api"
	fmt.Printf("It's my project: %q\n", projectName)

	fmt.Println(utils.SliceChunk([]int{1, 2, 3, 4, 5, 6, 7}, 2))
	fmt.Println(utils.SliceFilterByList([]int{-5, -4, -3, -3, -2, -2, 0, 0, 1, 2, 3, 4, 5, 6, 6, 7, 7, 7}))
	result, err := utils.MapInvert(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	userLink := link.New(1, 1, "github.com/ozonva/ova-link-api")
	userLink.SetDescription("Ozon Go School. Project.")
	tags := make(map[link.Tag]struct{})
	userLink.SetTags(tags)
	userLink.AddTag("tag1")
	userLink.AddTag("tag2")
	userLink.AddTag("tag3")
	userLink.RemoveTag("tag2")
	fmt.Println(userLink)
}
