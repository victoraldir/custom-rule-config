package main

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInvokingEvent(t *testing.T) {
	file, err := os.Open("../events/invoking.json")
	assert.Nil(t, err)
	invokingEventStr, err := io.ReadAll(file)
	assert.Nil(t, err)

	// Marshal the invoking event
	var invokingEvent InvokingEvent
	err = json.Unmarshal(invokingEventStr, &invokingEvent)
	assert.Nil(t, err)

	// Check the configuration item
	assert.NotNil(t, invokingEvent.ConfigurationItem)
	assert.NotNil(t, invokingEvent.ConfigurationItem.Tags)

}
