package controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type RabbitMQGateway struct {
	host      string
	namespace string
	username  string
	password  string
}

type RabbitOutput struct {
	Messages      int32 `json:"messages"`
	MessagesReady int32 `json:"messages_ready"`
}

func NewRabbitMQGateway(host, namespace, user, pass string) *RabbitMQGateway {
	return &RabbitMQGateway{host, namespace, user, pass}
}

func (r *RabbitMQGateway) RequestQuery(queue string) (*RabbitOutput, error) {
	log := logrus.WithField("Type", "[RABBITMQ_GATEWAY]")
	log.Info("Requesting RabbitMQ")
	defer log.Info("Finished")
	url := fmt.Sprintf("http://%s.%s:15672/api/queues/%s/%s", r.host, r.namespace, "%2f", queue)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Error("Failed to create request")
		return nil, err
	}
	auth := fmt.Sprintf("%s:%s", r.username, r.password)
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", authHeader)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Error("Failed to request RabbitMQ")
		return nil, err
	}
	defer resp.Body.Close()

	log.Infof("Status code: %d", resp.StatusCode)
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("unauthorized")
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("not found queue")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	responseBty, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Error("Failed to read response")
		return nil, err
	}
	output := &RabbitOutput{}
	if err := json.Unmarshal(responseBty, output); err != nil {
		return nil, err
	}
	return output, nil
}
