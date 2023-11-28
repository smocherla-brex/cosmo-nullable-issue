package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ravilushqa/otelgqlgen"
	"github.com/rs/cors"
	"github.com/wundergraph/cosmo-federation-demos/demos/go/pkg/injector"
	"github.com/wundergraph/cosmo-federation-demos/demos/go/pkg/otel"
	"github.com/wundergraph/cosmo-federation-demos/demos/go/pkg/subgraphs"
	"github.com/wundergraph/cosmo-federation-demos/demos/go/pkg/subgraphs/family"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

const (
	defaultPort = "4002"
	serviceName = "family"
)

func main() {
	otel.InitTracing(context.Background(), otel.Options{ServiceName: serviceName})
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := subgraphs.NewDemoServer(family.NewSchema())

	srv.Use(&debug.Tracer{})
	srv.Use(otelgqlgen.Middleware(otelgqlgen.WithCreateSpanFromFields(func(ctx *graphql.FieldContext) bool {
		return true
	})))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	})

	http.Handle("/", c.Handler(playground.Handler("GraphQL playground", "/graphql")))
	http.Handle("/graphql", injector.HTTP(c.Handler(otelhttp.NewHandler(srv, "", otelhttp.WithSpanNameFormatter(func(_operation string, r *http.Request) string {
		return fmt.Sprintf("%s %s", r.Method, r.URL.Path)
	})))))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
