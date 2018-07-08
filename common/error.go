package common

// Errors
const (
	None        = 100
	InvalidType = 200

	InvalidAuthentication = 300
	InvalidKey            = 301
	InvalidCard           = 302
	InvalidProof          = 303
	InvalidSignature      = 304
	InvalidPassphrase     = 305
	InvalidProject        = 306
	InvalidFile           = 310
	InvalidJSON           = 320
	InvalidEncoding       = 321

	NetworkError    = 400
	NetworkError404 = 404

	AlreadyExist = 500

	UnknownError    = 600
	UnknownElement  = 601
	UnknownProvider = 602

	EthOpError      = 800
	EthNotEnoughGas = 801

	NotYetImplemented = 999
)

// FBKError struct
type FBKError struct {
	err  string
	code int
}

// NewFBKError default Constructor
func NewFBKError(text string, code int) *FBKError {
	return &FBKError{text, code}
}

func (e *FBKError) Error() string {
	return e.err
}

// Type type
func (e *FBKError) Type() int {
	return e.code
}
