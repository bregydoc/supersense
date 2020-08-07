package sources

import (
	"time"

	"github.com/minskylab/supersense"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// Dummy is a minimal source implementation,
// it's util when you need to test supersense
type Dummy struct {
	id         string
	sourceName string
	clock      *time.Ticker
	message    string
	events     *chan supersense.Event
}

// NewDummy creates and init a new dummy source
func NewDummy(period time.Duration, message string) (*Dummy, error) {
	eventsChan := make(chan supersense.Event, 1)
	source := &Dummy{
		id:         uuid.NewV4().String(),
		sourceName: "dummy",
		clock:      time.NewTicker(period),
		events:     &eventsChan,
		message:    message,
	}
	return source, nil
}

// Run starts the recurrent message issuer
func (s *Dummy) Run() error {
	if s.events == nil {
		return errors.New("invalid Source, it not have an events channel")
	}

	go func() {
		for {
			event := <-s.clock.C
			*s.events <- supersense.Event{
				ID:        uuid.NewV4().String(),
				Message:   s.message,
				EmmitedAt: event,
				Person: supersense.Person{
					Name:        "John Doe",
					Photo:       "https://pic.jpeg",
					SourceOwner: s.sourceName,
				},
			}
		}
	}()

	return nil
}

// Events implements the supersense.Source interface
func (s *Dummy) Events() *chan supersense.Event {
	return s.events
}