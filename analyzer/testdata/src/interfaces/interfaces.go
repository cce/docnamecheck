package interfaces

// handler shows interface doc drift.
type handler interface {
	// ServeHTTP handles requests on interface methods too.
	serveHHTP() // want `doc comment starts with 'ServeHTTP' but symbol is 'serveHHTP' \(possible typo or old name\)`
}
