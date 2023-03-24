package contract

const IDKey = "skyscraper:id"

type IDService interface {
	NewID() string
}
