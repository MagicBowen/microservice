/*
Package registration for service registration;
Usage:
	server, err := Registration.NewService().Name().IP().Port().Register()
	defer server.stop()
*/
package registration

type (
	Service struct {
		info
	}
)
