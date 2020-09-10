package main

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"github.com/vishal1132/cafebucks/eventbus"
)

// eventHandler is the handler for events
func (s *server) eventHandler(msg kafka.Message, produce bool) {
	switch string(msg.Key) {
	case string(eventbus.OrderAccept):
		var event eventbus.EventC
		err := json.Unmarshal(msg.Value, &event)
		if err != nil {
			s.Logger.
				Error().
				Err(err).
				Msg("error unmarshaling order accept event")
			return
		}
		// sleep for some random time and then create an order finished event
		event.Event = eventbus.OrderProcessed
		event.Order.Status = eventbus.OrderProcessed
		b, err := json.Marshal(event)
		if produce {
			err = s.EventBus.Publish(context.Background(), eventbus.OrderProcessed, b)
		}
		if err != nil {
			s.Logger.Error().Err(err).Msg("error pushing event to kafka")
		}
	}
}
