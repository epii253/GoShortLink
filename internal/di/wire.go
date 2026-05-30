//go:build wireinject

package di

import (
	"project/internal/application/repositories"
	"project/internal/application/services/links"
	linkhandlers "project/internal/controllers/link_handlers"
	"project/internal/infrastructure"

	"github.com/google/wire"
)

var RepoSet = wire.NewSet(
	infrastructure.NewLinksInMemoryRepo,
	wire.Bind(new(repositories.ILinksRepo), new(*infrastructure.LinksInMemoryRepo)),
)

var ServiceSet = wire.NewSet(
	links.NewLinkService,
	wire.Bind(new(links.ILinkService), new(*links.LinkService)),
)

var HandlerSet = wire.NewSet(
	linkhandlers.NewLinkHandler,
)

func InitializeHandler() (*linkhandlers.LinkHandler, error) {
	wire.Build(
		RepoSet,
		ServiceSet,
		HandlerSet
	)
	
	return nil, nil
}
