package main

import (
	"context"

	"go.uber.org/zap"

	"github.com/google/uuid"

	pb "github.com/guycole/elder-mixer/proto"
)

type ServerType struct {
	SugarLog *zap.SugaredLogger
	pb.UnimplementedMixerServer
}

func (st *ServerType) EnqueueCommand(ctx context.Context, in *pb.EnqueueRequest) (*pb.EnqueueResponse, error) {
	st.SugarLog.Debug("enqueue command")

	clientId := in.ClientId
	//message := in.Message
	receiptId := uuid.NewString()

	return &pb.EnqueueResponse{ClientId: clientId, ReceiptId: receiptId}, nil
}
