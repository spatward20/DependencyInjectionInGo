package postalservice

type PostalService interface {
	Send(origin, destination string) (string, error)
}
