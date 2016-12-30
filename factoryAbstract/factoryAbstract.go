package factoryAbstract

type Producer interface {
	Produce(name *string)
}

type ProducerCreator interface {
	Create() Producer
}
