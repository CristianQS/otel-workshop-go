package common

import (
	"time"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

func ConsoleExporter(durationTime time.Duration) trace.TracerProvider {
	consoleExporter, _ := stdouttrace.New(stdouttrace.WithPrettyPrint())
	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(consoleExporter, trace.WithBatchTimeout(durationTime)))
	return *tracerProvider
}
