package revel

import (
	"runtime/debug"
)

// PanicFilter wraps the action invocation in a protective defer block
// that converts panics into 500 error pages.
func PanicFilter(c *Controller, fc []Filter) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			handleInvocationPanic(c, panicErr)
		}
	}()
	fc[0](c, fc[1:])
}

// handleInvocationPanic processes a panic in an action invocation.
// It logs the panic and displays an error page, showing the stack trace
// in development mode, while being more secure in production.
func handleInvocationPanic(c *Controller, panicErr interface{}) {
	panickedError := NewErrorFromPanic(panicErr)

	if panickedError == nil && DevMode {
		// In development mode, show full stack trace and error details.
		ERROR.Print(panicErr, "\n", string(debug.Stack()))
		c.Response.Out.WriteHeader(500)
		_, _ = c.Response.Out.Write(debug.Stack())
		return
	}

	// In production, we log the error and show a generic error page.
	ERROR.Print(panicErr, "\n", panickedError.Stack)
	c.Result = c.RenderError(panickedError)
}
