import * as opentracing from "opentracing";
import * as jaeger from "jaeger-client";
import {SamplerConfig, TracingOptions} from "jaeger-client";

const initTracer = function (serviceName: string) {
    const cfg : jaeger.TracingConfig = {
        serviceName: serviceName,
        // sampler: {},
        reporter: {},
    };

    const opt : jaeger.TracingOptions = {
        logger : {
            info: function logInfo(msg: string) {
                console.log("INFO : ", msg);
            },
            error: function logInfo(msg: string) {
                console.log("ERROR : ", msg);
            },
        }
    };

    const tracer = jaeger.initTracer(cfg, opt);
    opentracing.initGlobalTracer(tracer);
};

export default initTracer;