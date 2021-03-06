package backend

import (
	"io"
	"net"

	"github.com/rs/zerolog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/memprofiler/memprofiler/schema"
	"github.com/memprofiler/memprofiler/server/config"
	"github.com/memprofiler/memprofiler/server/locator"
)

var _ Service = (*server)(nil)

type server struct {
	grpcServer      *grpc.Server
	listener        net.Listener
	protocolFactory protocolFactory
	logger          *zerolog.Logger
	errChan         chan<- error
	cfg             *config.BackendConfig
}

func (s *server) Start() {
	s.errChan <- s.grpcServer.Serve(s.listener)
}

func (s *server) Stop() {
	s.grpcServer.GracefulStop()
}

func (s *server) SaveReport(stream schema.MemprofilerBackend_SaveReportServer) error {
	s.logger.Debug().Msg("Started request handling")

	// create object that will be responsible for handling incoming messages
	protocol := s.protocolFactory.save()
	defer func() {
		if err := protocol.close(); err != nil {
			s.logger.Err(err).Msg("Failed to close save protocol")
		}
	}()

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&schema.SaveReportResponse{})
		}
		if err != nil {
			return err
		}
		// FIXME: validate request, all required fields must exist
		switch request.Payload.(type) {
		case *schema.SaveReportRequest_InstanceDescription:
			err = protocol.addDescription(request.GetInstanceDescription())
		case *schema.SaveReportRequest_Measurement:
			err = protocol.addMeasurement(request.GetMeasurement())
		}
		if err != nil {
			s.logger.Err(err).Msg("Save error")
			return err
		}
	}
}

// NewServer builds new GRPC server
func NewServer(
	cfg *config.BackendConfig,
	locator *locator.Locator,
	errChan chan<- error,
) (Service, error) {

	listener, err := net.Listen("tcp", cfg.ListenEndpoint)
	if err != nil {
		return nil, err
	}

	s := &server{
		protocolFactory: &defaultProtocolFactory{locator: locator},
		listener:        listener,
		cfg:             cfg,
		errChan:         errChan,
		logger:          locator.Logger,
	}

	s.grpcServer = grpc.NewServer()
	schema.RegisterMemprofilerBackendServer(s.grpcServer, s)
	reflection.Register(s.grpcServer)

	return s, nil
}
