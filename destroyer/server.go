package destroyer

import (
	"context"
	"encoding/json"
	"github.com/adigunhammedolalekan/microservices-test/destroyer/proto/pb"
	"github.com/adigunhammedolalekan/microservices-test/types"
	"github.com/apache/pulsar-client-go/pulsar"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	conn *grpc.Server
}

func New(svc *Service) (*Server, error) {
	s := grpc.NewServer()
	pb.RegisterDestroyerServiceServer(s, svc)
	return &Server{conn: s}, nil
}

func (srv *Server) Run(lis net.Listener) error {
	return srv.conn.Serve(lis)
}

type Producer interface {
	Send(ctx context.Context, opt *pulsar.ProducerMessage) (pulsar.MessageID, error)
}

type Service struct {
	producer Producer
	store Store
}


func NewService(store Store, producer Producer) *Service {
	return &Service{producer: producer, store: store}
}

func (svc *Service) AcquireTargets(ctx context.Context, req *pb.EventRequest) (*pb.EventResponse, error) {
	targets := make([]*types.Target, 0, len(req.Data))
	for _, tgt := range req.Data {
		target, err := types.NewTarget(tgt)
		if err != nil {
			return nil, err
		}
		targets = append(targets, target)
	}
	ev := types.NewEvent(req.Name, targets)
	payload, err := json.Marshal(ev)
	if err != nil {
		return nil, err
	}
	res, err := svc.producer.Send(ctx, &pulsar.ProducerMessage{
		Payload: payload,
	})
	if err != nil {
		return nil, err
	}
	return &pb.EventResponse{MessageId: string(res.Serialize())}, nil
}

func (svc *Service) ListTargets(ctx context.Context, req *pb.ListTargetsRequest) (*pb.TargetResponse, error) {
	data, err := svc.store.ListTargets()
	if err != nil {
		return nil, err
	}
	r := types.Convert(data)
	return &pb.TargetResponse{
		Data: r,
	}, nil
}
