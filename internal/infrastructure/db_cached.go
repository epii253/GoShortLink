package infrastructure

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"project/internal/application/repositories"
	"project/internal/domain"
	settings "project/internal/settings"
)

type CachedRepo struct {
	mainDb     repositories.ILinksRepo
	redisCache *redis.Client
	ttl        time.Duration
}

func NewCachedRepo(mainDb *LinksDbRepo, redisCache *redis.Client, cnf *settings.Config) *CachedRepo {
	return newCachedRepo(mainDb, redisCache, cnf.CacheTTL)
}

// NewCachedRepoWith accepts ILinksRepo so tests can inject mocks.
func NewCachedRepoWith(mainDb repositories.ILinksRepo, redisCache *redis.Client, ttl time.Duration) *CachedRepo {
	return newCachedRepo(mainDb, redisCache, ttl)
}

func newCachedRepo(mainDb repositories.ILinksRepo, redisCache *redis.Client, ttl time.Duration) *CachedRepo {
	return &CachedRepo{mainDb: mainDb, redisCache: redisCache, ttl: ttl}
}

func (repo *CachedRepo) GetByLink(shortLink string) (*domain.Link, error) {
	ctx, cansel := context.WithTimeout(context.Background(), time.Second*2)
	defer cansel()

	fullUrl, redisErr := repo.redisCache.Get(ctx, shortLink).Result()

	if redisErr == nil { // cahce hit
		return &domain.Link{
				FullUrl:   fullUrl,
				ShortCode: shortLink,
			},
			nil
	}

	link, err := repo.mainDb.GetByLink(shortLink)
	if err != nil {
		return nil, err
	}

	repo.redisCache.Set(ctx, shortLink, link.FullUrl, repo.ttl)

	return link, nil
}

func (repo *CachedRepo) CheckExsist(shrortLink string) (bool, error) {
	return repo.mainDb.CheckExsist(shrortLink)
}

func (repo *CachedRepo) TryAddItem(newItem *domain.Link) (bool, error) {
	return repo.mainDb.TryAddItem(newItem)
}

func (repo *CachedRepo) DeleteItemByShortLink(shrortLink string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	repo.redisCache.Del(ctx, shrortLink)

	return repo.mainDb.DeleteItemByShortLink(shrortLink)
}
