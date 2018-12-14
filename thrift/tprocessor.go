package thrift

/**
 * A processor is a generic object which operates upon an input stream and
 * writes to some output stream.
 *
 */
type TProcessor interface {
	Process(in, out TProtocol) (bool, TException)
}

type TProcessorFunction interface {
	Process(seqId int32, in, out TProtocol) (bool, TException)
}

type TProcessorFactory interface {
	GetProcessor(trans TTransport) TProcessor
}

type tProcessorFactory struct {
	processor TProcessor
}

func NewTProcessorFactory(p TProcessor) TProcessorFactory {
	return &tProcessorFactory{processor: p}
}

func (p *tProcessorFactory) GetProcessor(trans TTransport) TProcessor {
	return p.processor
}

/**
 * The default processor factory just returns a singleton
 * instance.
 */
type TProcessorFunctionFactory interface {
	GetProcessorFunction(trans TTransport) TProcessorFunction
}

type tProcessorFunctionFactory struct {
	processor TProcessorFunction
}

func NewTProcessorFunctionFactory(p TProcessorFunction) TProcessorFunctionFactory {
	return &tProcessorFunctionFactory{processor: p}
}

func (p *tProcessorFunctionFactory) GetProcessorFunction(trans TTransport) TProcessorFunction {
	return p.processor
}
