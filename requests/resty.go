package requests

import "github.com/go-resty/resty/v2"

type Resty struct {
	client *resty.Client
}

func (s *Resty) Get() *resty.Client {
	return s.client
}
func (s *Resty) WithJSON() *Resty {
	s.client.SetHeader("Accept", "application/json").SetHeader("Content-Type", "application/json")
	return s
}

func (s *Resty) WithRetry() *Resty {

	return s
}

func (s *Resty) WithTLS() *Resty {

	return s
}
