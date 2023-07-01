//go:build !wireinject
// +build !wireinject

package DI

import "github.com/google/wire"

func InitializedService() *SimpleService {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)

	// nil karena nanti otomatis wire yang handle
	return nil
}
