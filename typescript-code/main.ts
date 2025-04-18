import initTracer from "./src/jaeger/root"
import {opentracing} from "jaeger-client";

console.log("test");

initTracer("test_tracer");

const tracer: opentracing.Tracer = opentracing.globalTracer();

const span = tracer.startSpan("hello_test");
span.setTag("hello_hello", "test_hello");

span.log({
    event: "string",
    value: "string test",
});

span.finish();