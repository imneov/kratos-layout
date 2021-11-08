package model

import "github.com/google/wire"

// ProviderSet is model providers.
var ProviderSet = wire.NewSet(NewGreeterUsecase)
