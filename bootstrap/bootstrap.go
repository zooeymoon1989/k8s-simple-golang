package bootstrap

//func InitTracing() {
//	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
//	agentEndpointURI := "jaegertracing:6831"
//	collectorEndpointURI := "http://localhost:14268/api/traces"
//	exporter, err := jaeger.NewExporter(jaeger.Options{
//		CollectorEndpoint: collectorEndpointURI,
//		AgentEndpoint:     agentEndpointURI,
//		ServiceName:       "product_info",
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	trace.RegisterExporter(exporter)
//}
