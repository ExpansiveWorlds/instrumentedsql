package tracedSQL

import (
	"database/sql"
	"context"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/ExpansiveWorlds/traced-sql/google"
	"github.com/ExpansiveWorlds/traced-sql/opentracing"
)

// WrapDriverGoogle demonstrates how to call wrapDriver and register a new driver.
// This example uses MySQL and google tracing to illustrate this
func ExampleWrapDriver_Google() {
	logger := func(ctx context.Context, msg string, keyvals ...interface{}) {
		log.Printf("%s %v", msg, keyvals)
	}

	sql.Register("traced-mysql", WrapDriver(mysql.MySQLDriver{}, google.NewTracer(), NewFuncLogger(logger)))
}


// WrapDriverOpentracing demonstrates how to call wrapDriver and register a new driver.
// This example uses MySQL and opentracing to illustrate this
func ExampleWrapDriver_Opentracing() {
	logger := func(ctx context.Context, msg string, keyvals ...interface{}) {
		log.Printf("%s %v", msg, keyvals)
	}

	sql.Register("traced-mysql", WrapDriver(mysql.MySQLDriver{}, opentracing.NewTracer(), NewFuncLogger(logger)))
}
