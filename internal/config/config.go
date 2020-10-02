package config

const (
	// ServerPort .
	ServerPort = 8080

	// EnvURI .
	EnvURI = "DB_URI"

	// EnvOldDBURI .
	EnvOldDBURI = "OLD_DB_URI"

	// EnvSigner .
	EnvSigner = "SIGNER"

	// ErrNoObjectFoundForID .
	ErrNoObjectFoundForID = "no object found for id"

	// ErrInvalidContentType .
	ErrInvalidContentType = "content-type is not application/json"

	// ErrMissingAuthorization .
	ErrMissingAuthorization = "missing authorization details"

	// ErrInvalidAuthorization .
	ErrInvalidAuthorization = "invalid authorization details"

	// ErrUsernameTaken .
	ErrUsernameTaken = "username taken"

	// ErrInvalidToken .
	ErrInvalidToken = "invalid token"

	// ErrUnexpectedSigningMethod .
	ErrUnexpectedSigningMethod = "unexpected signing method: %v"

	// HeaderContentType .
	HeaderContentType = "Content-Type"

	// HeaderApplicationJSON .
	HeaderApplicationJSON = "application/json"

	// HeaderAuthorization .
	HeaderAuthorization = "Authorization"
)
