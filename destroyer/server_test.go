package destroyer

import (
	"context"
	"github.com/adigunhammedolalekan/microservices-test/destroyer/proto/pb"
	"github.com/adigunhammedolalekan/microservices-test/types"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

const (
	bufSize = 1024 * 1024
)

var (
	mockTargets = []*pb.Target{
		{Id: uuid.New().String(), Message: "some message to send", CreatedOn: "2020-06-25T16:31:18.993Z", UpdatedOn: "2020-06-25T16:31:18.993Z"},
		{Id: uuid.New().String(), Message: "some other message to send", CreatedOn: "2020-06-25T16:31:18.993Z", UpdatedOn: "2020-06-25T16:31:18.993Z"},
	}
)

type mockStore struct{}
type mockProducer struct{}
type mockMessageID struct{}

func newMockProducer() *mockProducer {
	return &mockProducer{}
}

func newMockStore() *mockStore {
	return &mockStore{}
}

func newMockMessage() pulsar.MessageID {
	return &mockMessageID{}
}

func (m *mockProducer) Send(ctx context.Context, opt *pulsar.ProducerMessage) (pulsar.MessageID, error) {
	return newMockMessage(), nil
}

func (ms *mockStore) ListTargets() ([]*types.Target, error) {
	return []*types.Target{
		{Id: "1", Message: "hello"}, {Id: "2", Message: "world"},
	}, nil
}

func (mm *mockMessageID) Serialize() []byte {
	return []byte(uuid.New().String())
}

func TestService_AcquireTargets(t *testing.T) {
	producer := newMockProducer()
	store := newMockStore()
	svc := NewService(store, producer)
	srv, err := New(svc)
	if err != nil {
		t.Fatal(err)
	}
	lis := bufconn.Listen(bufSize)
	go func() {
		if err := srv.Run(lis); err != nil {
			t.Fatal(err)
		}
	}()
	client, err := grpc.Dial("net", grpc.WithContextDialer(createDialer(lis)), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	destroyerClient := pb.NewDestroyerServiceClient(client)
	r, err := destroyerClient.AcquireTargets(context.Background(), &pb.EventRequest{
		Id:   uuid.New().String(),
		Name: "targets.acquired",
		Data: mockTargets,
	})
	if err != nil {
		t.Fatal(err)
	}
	if r.MessageId == "" {
		t.Fatal("message id is expected to not be nil")
	}
	t.Log(r.MessageId)
}

func TestService_ListTargets(t *testing.T) {
	s := newMockStore()
	svc := NewService(s, nil)
	srv, err := New(svc)
	if err != nil {
		t.Fatal(err)
	}
	lis := bufconn.Listen(bufSize)
	go func() {
		if err := srv.Run(lis); err != nil {
			t.Fatal(err)
		}
	}()
	client, err := grpc.Dial("net", grpc.WithContextDialer(createDialer(lis)), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	destroyerClient := pb.NewDestroyerServiceClient(client)
	r, err := destroyerClient.ListTargets(context.Background(), &pb.ListTargetsRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Data) != 2 {
		t.Fatalf("expected 2 targets. got %d instead", len(r.Data))
	}
	t.Log(r.Data)
}

func createDialer(lis *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(i context.Context, s string) (conn net.Conn, e error) {
		return lis.Dial()
	}
}
