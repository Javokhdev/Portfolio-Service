package main

import (
	"log"
	"net"

	pb "github.com/Javokhdev/Portfolio-Service/genprotos"
	"github.com/Javokhdev/Portfolio-Service/service"
	postgres "github.com/Javokhdev/Portfolio-Service/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	liss, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterSkillServiceServer(s, service.NewSkillService(db))
	pb.RegisterProjectServiceServer(s, service.NewProjectService(db))
	pb.RegisterExperienceServiceServer(s, service.NewExperienceService(db))
	pb.RegisterEducationServiceServer(s, service.NewEducationService(db))
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
