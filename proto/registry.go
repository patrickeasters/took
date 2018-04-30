package proto

type RefreshOption int

// Constants for how to get the token
const (
	UseDefault RefreshOption = iota
	UseRefresh
	UseReAuth
)

type OutputOption int

// Output options
const (
	OutputToken OutputOption = iota
	OutputHeader
)

// Protocol defines a protocol
type Protocol interface {
	// GetDataInstance returns a new data block into which the data
	// part of the configuration will be unmarshaled
	GetDataInstance() interface{}
	// GetConfigInstance returns a new configuration instance into which the configuration will be unmarshaled
	GetConfigInstance() interface{}
	// GetToken returns the token with the given configuration and data blocks
	GetToken(RefreshOption, OutputOption) (string, error)
}

var protocols = make(map[string]func() Protocol)

// Register registers a protocol
func Register(name string, factory func() Protocol) {
	protocols[name] = factory
}

// Get retrieves a protocol by name
func Get(name string) Protocol {
	p, ok := protocols[name]
	if ok {
		return p()
	}
	return nil
}