import { opentracing } from "jaeger-client";
import initTracer from "./src/jaeger/root";
import * as express from "express";

console.log("test");

initTracer("test_tracer");

const app = express();

app.get("/api/resource", (req, res) => {
    const tracer: opentracing.Tracer = opentracing.globalTracer();

    const headersCarieer = req.headers;
    const spaneContext = tracer.extract(
        opentracing.FORMAT_HTTP_HEADERS,
        headersCarieer
    )

    const span = tracer.startSpan("hello_test", {childOf: spaneContext});
    span.setTag("hello_hello", "test_hello");
    span.setTag(opentracing.Tags.HTTP_METHOD, "GET");
    span.setTag(opentracing.Tags.HTTP_URL, "/api/resource");

    span.log({
        event: "string",
        value: "string test",
    });

    span.finish();
    res.send("Success")
});

app.listen(8080, () => {
    console.log("server started at", 8080);
});