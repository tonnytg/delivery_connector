package event_consumer

type EventDelivery struct{}

func NewEventDelivery() *EventDelivery {
	return &EventDelivery{}
}

func (ed *EventDelivery) Start() {
	// Implementação do método Start
}
