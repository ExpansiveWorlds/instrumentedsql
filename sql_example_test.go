package instrumentedsql

import (
	"database/sql"
	"context"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/ExpansiveWorlds/instrumentedsql/google"
	"github.com/ExpansiveWorlds/instrumentedsql/opentracing"
)

// WrapDriverGoogle demonstrates how to call wrapDriver and register a new driver.
// This example uses MySQL and google tracing to illustrate this
func ExampleWrapDriver_google() {
	logger := func(ctx context.Context, msg string, keyvals ...interface{}) {
		log.Printf("%s %v", msg, keyvals)
	}

	sql.Register("instrumented-mysql", WrapDriver(mysql.MySQLDriver{}, WithTracer(google.NewTracer()), WithLogger(NewFuncLogger(logger))))
}


// WrapDriverOpentracing demonstrates how to call wrapDriver and register a new driver.
// This example uses MySQL and opentracing to illustrate this
func ExampleWrapDriver_opentracing() {
	logger := func(ctx context.Context, msg string, keyvals ...interface{}) {
		log.Printf("%s %v", msg, keyvals)
	}

	sql.Register("instrumented-mysql", WrapDriver(mysql.MySQLDriver{}, WithTracer(opentracing.NewTracer()), WithLogger(NewFuncLogger(logger))))
}
