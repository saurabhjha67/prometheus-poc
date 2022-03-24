package pubSub

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/pubsub"
	"github.com/sirupsen/logrus"
)

func PublishProtoMessages(w io.Writer, message string) error {
	logrus.Info("PublishProtoMessages called with message : ", message)
	projectID := "v2agent-9423a"
	topicID := "VM_CPU_UTIL"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	// Get the topic encoding type.
	t := client.Topic(topicID)

	messageBytes := []byte(message)

	result := t.Publish(ctx, &pubsub.Message{
		Data: messageBytes,
	})
	serverId, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("result.Get: %v", err)
	}
	fmt.Fprintf(w, "Published proto message with %#v with serverId : %s\n", serverId, string(messageBytes))
	return nil
}
