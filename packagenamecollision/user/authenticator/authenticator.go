package authenticator

type state struct {
	isAuthenticated bool
}

var authenticationState *state

func init() {
	authenticationState = &state{
		isAuthenticated: false,
	}
}

func IsAuthenticated() bool {
	authenticationState.isAuthenticated = true
	return authenticationState.isAuthenticated
}
