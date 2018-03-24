// Package multicloser is for easily closing groups of io.Closer types in a safe
// and ergonomic way that reduces code repetition and reduces the risk of
// forgetting to handle the errors returned by Close methods.
package multicloser

import "io"

// MultiCloser provides a simple API around io.Closer types for ergonomically
// allowing lots of things that return io.Closers to be closed together. It's
// designed for resources that are always only constructed on startup and need
// to be collectively shut down gracefully.
//
// It embeds the io.Closer interface and declares a single additional method,
// Add, which is designed to wrap functions that construct types that satisfy the
// io.Closer interface and also may return an error. It will add the closer to
// its internal list and close it when its Close method is called.
//
// For example, say you have a bunch of APIs that look like this:
//
//     NewDatabaseThing(dsn string, opts Options) (DatabaseThinger, error)
//
// Because this returns a paid of some type that implements io.Closer and an
// error, it can be passed directly to the Add method:
//
//     m.Add(NewDatabaseThing(dsn, opts))
//
// Also note that to remain generic, the Add method returns an io.Closer, the
// same as the input. So to use your own type, you need to do a simply cast. No
// need to worry about checking the conversion though since you know exactly
// what type was passed in:
//
//     db := m.Add(NewDatabaseThing(dsn, opts)).(DatabaseThinger)
//
// If your constructors don't return errors, simply fill it with nil:
//
//     db := m.Add(NewSimpleThing(opts), nil).(SimpleThinger)
//
// If your constructors return more than (io.Closer, error) then... well,
// consider simplifying them a bit!
//
type MultiCloser interface {
	io.Closer
	Add(io.Closer, error) io.Closer
}
