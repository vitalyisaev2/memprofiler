package metrics

import (
	"fmt"

	"github.com/memprofiler/memprofiler/schema"
)

func shortSessionIdentifier(sessionDesc *schema.SessionDescription) string {
	return fmt.Sprintf(
		"%s::%s::%d",
		sessionDesc.InstanceDescription.ServiceName,
		sessionDesc.InstanceDescription.InstanceName,
		sessionDesc.Id,
	)
}
