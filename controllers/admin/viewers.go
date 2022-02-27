package admin

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/owncast/owncast/controllers"
	"github.com/owncast/owncast/metrics"
	log "github.com/sirupsen/logrus"
)

// GetViewersOverTime will return the number of viewers at points in time.
func GetViewersOverTime(w http.ResponseWriter, r *http.Request) {
	windowStartAtStr := r.URL.Query().Get("windowStart")
	windowStartAtUnix, err := strconv.Atoi(windowStartAtStr)
	if err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}

	windowStartAt := time.Unix(int64(windowStartAtUnix), 0)
	windowEnd := time.Now()

	viewersOverTime := metrics.GetViewersOverTime(windowStartAt, windowEnd)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(viewersOverTime)
	if err != nil {
		log.Errorln(err)
	}
}
