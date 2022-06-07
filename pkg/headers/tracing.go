package headers

const (
	SpanID        = "X-B3-SpanId"
	TraceID       = "X-B3-TraceId"
	ParentSpanID  = "X-B3-ParentSpanId"
	Sampled       = "X-B3-Sampled"
	Flags         = "X-B3-Flags"
	SpanContext   = "X-Ot-Span-Context"
	CloudTrace    = "X-Cloud-Trace-Context"
	JaegerBaggage = "Jaeger-Baggage"
	UberTraceID   = "Uber-Trace-Id"
	AwsTraceID    = "X-Amzn-Trace-Id"
	AwsRequestID  = "x-amzn-RequestId"
	RequestID     = "X-Request-Id"
	TransactionID = "X-Transaction-Id"
)
