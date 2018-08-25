package web

import (
	"net/http"
	"sort"

	"github.com/julienschmidt/httprouter"
	"github.com/vitalyisaev2/memprofiler/schema"
	"github.com/vitalyisaev2/memprofiler/server/storage"
)

func (s *server) computeSessionMetrics(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// parse session id
	sessionID, err := storage.SessionIDFromString(params.ByName("session"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ask storage for session data
	sessionDescription := &storage.SessionDescription{
		ServiceDescription: &schema.ServiceDescription{
			Type:     params.ByName("type"),
			Instance: params.ByName("instance"),
		},
		SessionID: sessionID,
	}

	// get session metrics
	result, err := s.computer.GetSessionMetrics(r.Context(), sessionDescription)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// sort results by InUseBytes rate, since it the most relevant indicator for memory leak
	s.logger.Debug("sorting slice")
	sort.Slice(result.Locations, func(i, j int) bool {
		// descending order
		return result.Locations[i].Rates.InUseBytes > result.Locations[j].Rates.InUseBytes
	})

	// dump to JSON
	m := newJSONMarshaler()
	w.Header().Set("Content-Type", "application/json")
	if err = m.Marshal(w, result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
