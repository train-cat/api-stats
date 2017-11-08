package route

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/train-cat/api-stats/helper"
	"github.com/train-cat/api-stats/model"
)

// Issue record an issue to the database
func Issue(w http.ResponseWriter, r *http.Request) {
	i, err := model.GetIssueFromHTTPRequest(r)

	if helper.HTTPError(w, err) {
		return
	}

	e, err := i.ToEntity()

	if helper.HTTPError(w, err) {
		return
	}

	err = e.Persist()

	if helper.HTTPError(w, err) {
		return
	}

	log.Infof("ACK station:%d code:%s", i.StationID, i.Code)

	w.WriteHeader(http.StatusNoContent)
}
