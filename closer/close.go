package closer

import (
	"database/sql"
	"io"
	"os"
)

// Some objects ask you to call Close() on them - databases, file descriptors, etc.
// The key point is that any structure with a Close() method implements io.Closer.
// It can be a good way to centralize error management / closing stuff instead of handling this in defer func().

// Here is an example of a "bad" Close management, because pretty "heavy" on the eyes.
func badClose() error {
	var someRows = new(sql.Rows)
	var someDB = new(sql.DB)
	var someFile = new(os.File)

	// Do some operations on all of those...
	// let's say: open the DB, query something, record in a CSV

	// Time to close : this code is full of repetition
	if err := someRows.Close(); err != nil {
		return err
	}
	if err := someDB.Close(); err != nil {
		return err
	}
	if err := someFile.Close(); err != nil {
		return err
	}

	return nil
}

// A better way to do this could be :
func niceClose() []error {
	var collectClosers []io.Closer
	var someRows = new(sql.Rows)
	var someDB = new(sql.DB)
	var someFile = new(os.File)
	// Since all those have Close(), they *are* "io.Closer" interfaces.
	collectClosers = append(collectClosers, someRows, someDB, someFile)

	// Do some operations...

	// Finish the function call : close everything and return errors
	var allErrors []error
	for _, thing := range collectClosers {
		if err := thing.Close(); err != nil {
			allErrors = append(allErrors, err)
		}
	}

	return allErrors
}

// Variadic functions allow us to use either 1 or more closers at the same time,
// which makes this function actually useable in every case.
func CloseAll(obj ...io.Closer) []error {
	var errs []error
	for _, o := range obj {
		if err := o.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

// Arbitrary type to make a closeAll unit test
type Closeable struct{}

// Implement io.Closer
func (c *Closeable) Close() error {
	return nil
}
