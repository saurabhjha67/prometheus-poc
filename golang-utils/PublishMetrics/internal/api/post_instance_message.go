package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"com.publish.api/internal/contracts"
	"com.publish.api/internal/pubSub"
	"github.com/prometheus/alertmanager/template"
	"github.com/sirupsen/logrus"
)

func (s *server) PostInstanceMessageHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("post instance message api called")

	jsonBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		writeBadRequestResponse(w, fmt.Sprintf("Error decoding the body: %v", err))
		return
	}

	var request template.Data
	err = json.Unmarshal(jsonBytes, &request)

	if err != nil {
		writeBadRequestResponse(w, fmt.Sprintf("Error decoding the body: %v", err))
		return
	}

	var metricMetadatas []contracts.MetriMetadata

	for _, alert := range request.Alerts {
		metric := contracts.MetriMetadata{
			Instance:  alert.Labels["instance"],
			ProjectId: alert.Labels["project_Id"],
			Zone:      alert.Labels["zone"],
		}
		metricMetadatas = append(metricMetadatas, metric)
	}

	// metricMetadata := contracts.MetriMetadata{
	// 	Instance:  instanceName,
	// 	ProjectId: projectId,
	// 	Zone:      zone,
	// }

	err = pubSub.PublishProtoMessages(w, metricMetadatas)

	if err != nil {
		apiErr, ok := err.(*pubSub.ApiError)
		if ok {
			switch apiErr.Type() {
			case pubSub.DuplicateUser:
				writeConflictResponse(w, "This is an error response")
				return
			case pubSub.UserNotFound:
				writeNotFoundResponse(w, "This is an error response")
				return
			default:
				logrus.Errorf("Unexpected error type %s: %v", apiErr.Type().String(), err)
				writeErrorResponse(w, "This is an error response")
				return
			}
		} else {
			logrus.Errorf("Server Error: %v", err)
			writeErrorResponse(w, "This is an error response")
			return
		}
	}
	for _, metricInfo := range metricMetadatas {
		writeResponse(w, http.StatusOK, metricInfo)
	}
}
