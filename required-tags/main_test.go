package main

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	t.Run("Should parse invocation event", func(t *testing.T) {
		// Open the test event file
		file, err := os.Open("../events/invoking.json")
		assert.Nil(t, err)
		defer file.Close()

		// Read the file content
		invokingEventStr, err := io.ReadAll(file)
		assert.Nil(t, err)

		// Unmarshal the invoking event
		var invokingEvent InvokingEvent
		err = json.Unmarshal(invokingEventStr, &invokingEvent)
		assert.Nil(t, err)

		// Check the configuration item
		assert.NotNil(t, invokingEvent.ConfigurationItem)
		assert.NotNil(t, invokingEvent.ConfigurationItem.Tags)
	})

	t.Run("Should handle config event", func(t *testing.T) {
		// Open the test event file
		file, err := os.Open("../events/config-valid-tags.json")
		assert.Nil(t, err)
		defer file.Close()

		// Read the file content
		configEventStr, err := io.ReadAll(file)
		assert.Nil(t, err)

		// Unmarshal the config event
		var configEvent events.ConfigEvent
		err = json.Unmarshal(configEventStr, &configEvent)
		assert.Nil(t, err)

		// Check the configuration item
		assert.NotNil(t, configEvent.AccountID)
		assert.NotNil(t, configEvent.ConfigRuleArn)

		// Call the handler
		response, err := handler(configEvent)

		// Check the response
		assert.Nil(t, err)
		assert.Equal(t, "CONFORMANT", response)
	})

	t.Run("Should handle noncompliant resource and remediate", func(t *testing.T) {
		// Open the test event file
		file, err := os.Open("../events/config-invalid-tags.json")
		assert.Nil(t, err)
		defer file.Close()

		// Read the file content
		configEventStr, err := io.ReadAll(file)
		assert.Nil(t, err)

		// Unmarshal the config event
		var configEvent events.ConfigEvent
		err = json.Unmarshal(configEventStr, &configEvent)
		assert.Nil(t, err)

		// Call the handler
		response, err := handler(configEvent)

		// Check the response
		assert.Nil(t, err)
		assert.Equal(t, "CONFORMANT", response)
	})

	t.Run("Should handle no existing resource", func(t *testing.T) {
		// Open the test event file
		file, err := os.Open("../events/config-no-existing-object-id.json")
		assert.Nil(t, err)
		defer file.Close()

		// Read the file content
		configEventStr, err := io.ReadAll(file)
		assert.Nil(t, err)

		// Unmarshal the config event
		var configEvent events.ConfigEvent
		err = json.Unmarshal(configEventStr, &configEvent)
		assert.Nil(t, err)

		// Call the handler
		response, err := handler(configEvent)

		// Check the response
		assert.Nil(t, err)
		assert.Equal(t, "CONFORMANT", response)
	})
}
