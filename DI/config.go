package DI

type SimpleRepository struct {
}

type SimpleService struct {
	SimpleRepository
}

// membuat provider

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

func NewSimpleService(repository *SimpleRepository) *SimpleService {
	return &SimpleService{
		SimpleRepository: *repository,
	}
}
