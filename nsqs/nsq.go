package nsqs

import (
	"github.com/daqiancode/jsoniter"
	"github.com/nsqio/go-nsq"
)

type Nsq struct {
	conf     *nsq.Config
	producer *nsq.Producer
	json     jsoniter.API
}

func NewNsq(addr string, config *nsq.Config) *Nsq {
	r := &Nsq{
		conf: config,
		json: jsoniter.Config{Decapitalize: true}.Froze(),
	}
	producer, err := nsq.NewProducer(addr, r.conf)
	if err != nil {
		panic(err)
	}
	r.producer = producer
	return r
}

func (s *Nsq) Send(topic string, message interface{}) error {

	// defer producer.Stop()
	bs, err := s.json.Marshal(message)
	if err != nil {
		return err
	}
	return s.producer.Publish(topic, bs)
}

func (s *Nsq) SetJsoniter(json jsoniter.API) {
	s.json = json
}

func (s *Nsq) Stop() {
	s.producer.Stop()
}
func (s *Nsq) GetProducer() *nsq.Producer {
	return s.producer
}
