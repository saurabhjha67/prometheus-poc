package pubSub

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"cloud.google.com/go/pubsub"
	"com.publish.api/internal/contracts"
	"github.com/sirupsen/logrus"
)

func PublishProtoMessages(w io.Writer, metricMetadatas []contracts.MetriMetadata) error {

	projectID := "v2agent-9423a"
	topicID := "VM_CPU_UTIL"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	// Get the topic encoding type.
	t := client.Topic(topicID)

	for _, metricInfo := range metricMetadatas {
		logrus.Info("PublishProtoMessages called with message : ", metricInfo.Instance)
		logrus.Info("PublishProtoMessages called with message : ", metricInfo.ProjectId)
		logrus.Info("PublishProtoMessages called with message : ", metricInfo.Zone)
		message, err := json.Marshal(metricInfo)

		if err != nil {
			return fmt.Errorf("json.Marshal: %v", err)
		}

		result := t.Publish(ctx, &pubsub.Message{
			Data: message,
		})
		serverId, err := result.Get(ctx)
		if err != nil {
			return fmt.Errorf("result.Get: %v", err)
		}
		fmt.Fprintf(w, "Published proto message with %#v with serverId : %s\n", serverId, string(message))
	}
	return nil
}
