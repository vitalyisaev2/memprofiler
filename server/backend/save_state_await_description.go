package backend

import (
	"github.com/memprofiler/memprofiler/schema"
)

var _ saveState = (*saveStateAwaitDescription)(nil)

type saveStateAwaitDescription struct {
	saveStateCommon
}

func (s *saveStateAwaitDescription) addDescription(desc *schema.ServiceDescription) error {

	// run new saver in persistent storage
	dataSaver, err := s.p.getStorage().NewDataSaver(desc)
	if err != nil {
		s.switchState(finished)
		return err
	}

	if err := s.p.setDataSaver(dataSaver); err != nil {
		s.switchState(finished)
		return err
	}

	// set session description
	sd := &schema.SessionDescription{
		ServiceType:     desc.GetServiceType(),
		ServiceInstance: desc.GetServiceInstance(),
		SessionId:       dataSaver.SessionID(),
	}
	if err := s.p.setSessionDescription(sd); err != nil {
		s.switchState(finished)
		return err
	}

	// annotate logger and save it for further usage
	logger := s.p.getLogger().With().Fields(map[string]interface{}{
		"service_type":     desc.GetServiceType(),
		"service_instance": desc.GetServiceInstance(),
		"session_id":       dataSaver.SessionID(),
	}).Logger()
	s.p.setLogger(&logger)
	logger.Info().Msg("Received greeting message from client")

	s.switchState(awaitMeasurement)
	return nil
}
