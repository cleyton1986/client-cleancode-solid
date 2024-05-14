package monitoring

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PrometheusHandler retorna um gin.HandlerFunc para expor as métricas do Prometheus
func PrometheusHandler() gin.HandlerFunc {
	handler := promhttp.Handler()
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

// InitMetrics inicializa as métricas do Prometheus.
func InitMetrics() {
	opsProcessed := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed operations.",
	})
	prometheus.MustRegister(opsProcessed)
}
