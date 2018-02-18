package common

const None = 100
const InvalidType = 200

const InvalidAuthentication = 300
const InvalidKey = 301
const InvalidCard = 302
const InvalidProof = 303
const InvalidSignature = 304
const InvalidPassphrase = 305
const InvalidFile = 310
const InvalidJson = 320
const InvalidEncoding = 321

const NetworkError = 400
const NetworkError404 = 404

const AlreadyExist = 500

const UnknownError = 600
const UnknownElement = 601
const UnknownProvider = 602

const EthOpError = 800
const EthNotEnoughGas = 801

const NotYetImplemented = 999

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
