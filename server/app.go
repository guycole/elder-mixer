package main

import (
	"fmt"
	"net"
	"time"

	"go.uber.org/zap"

	"google.golang.org/grpc"

	pb "github.com/guycole/elder-mixer/proto"
)

type AppType struct {
	FeatureFlags uint32             // control run time features
	GrpcPort     int                // gRPC port
	SugarLog     *zap.SugaredLogger // logging

	Quantum time.Duration
	RunFlag bool // true while schedule
}

func (at *AppType) initialize() {
	if isDevelopmentModeLogging(at.FeatureFlags) {
		at.SugarLog = zapSetup(true)
		at.SugarLog.Debug("debug level log entry")
	}

	at.RunFlag = true
}

// Run pacifier
func (at *AppType) run() {
	at.SugarLog.Info("run run run")

	//at.SugarLog.Fatal(http.ListenAndServe(":"+address, at.Router))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", at.GrpcPort))
	if err != nil {
		at.SugarLog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMixerServer(grpcServer, &ServerType{})
	at.SugarLog.Infof("server listening at %v", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		at.SugarLog.Fatalf("failed to serve: %v", err)
	}
}
