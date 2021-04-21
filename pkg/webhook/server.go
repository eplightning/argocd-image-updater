package webhook

import (
	"fmt"
	"net/http"

	"github.com/argoproj-labs/argocd-image-updater/pkg/log"
)

type webhookServer struct {
	trigger *UpdateTrigger
}

func StartWebhookServer(port int, trigger *UpdateTrigger) chan error {
	server := &webhookServer{trigger: trigger}

	errCh := make(chan error)
	go func() {
		sm := http.NewServeMux()
		sm.HandleFunc("/webhook/generic", server.receiveGenericWebhook)
		errCh <- http.ListenAndServe(fmt.Sprintf(":%d", port), sm)
	}()
	return errCh
}

func (ws *webhookServer) receiveGenericWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ws.trigger.Trigger()
		log.Infof("webhook triggered")
	}

	fmt.Fprintf(w, "OK\n")
}
