package service

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Registration(account string) error {
	return nil
}
