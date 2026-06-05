package tests

// import (
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/go-redis/redismock/v9"
// 	"github.com/stretchr/testify/assert"

// 	"project/internal/domain"
// 	"project/internal/infrastructure"
// 	"project/tests/mocks"
// )

// const cacheTTL = 2 * time.Minute

// func TestGetByLink_CacheHit(t *testing.T) {
// 	client, rMock := redismock.NewClientMock()
// 	rMock.ExpectGet("abc1234").SetVal("https://example.com")

// 	repo := new(mocks.MockLinksRepo)
// 	cached := infrastructure.NewCachedRepoWith(repo, client, cacheTTL)

// 	link, err := cached.GetByLink("abc1234")

// 	assert.NoError(t, err)
// 	assert.Equal(t, "https://example.com", link.FullUrl)
// 	assert.Equal(t, "abc1234", link.ShortCode)
// 	repo.AssertNotCalled(t, "GetByLink")
// 	assert.NoError(t, rMock.ExpectationsWereMet())
// }

// func TestGetByLink_CacheMiss_FetchesDB(t *testing.T) {
// 	client, rMock := redismock.NewClientMock()
// 	rMock.ExpectGet("abc1234").RedisNil()
// 	rMock.ExpectSet("abc1234", "https://example.com", cacheTTL).SetVal("OK")

// 	repo := new(mocks.MockLinksRepo)
// 	repo.On("GetByLink", "abc1234").Return(&domain.Link{ShortCode: "abc1234", FullUrl: "https://example.com"}, nil)

// 	cached := infrastructure.NewCachedRepoWith(repo, client, cacheTTL)

// 	link, err := cached.GetByLink("abc1234")

// 	assert.NoError(t, err)
// 	assert.Equal(t, "https://example.com", link.FullUrl)
// 	repo.AssertExpectations(t)
// 	assert.NoError(t, rMock.ExpectationsWereMet())
// }

// func TestGetByLink_CacheMiss_DBError(t *testing.T) {
// 	client, rMock := redismock.NewClientMock()
// 	rMock.ExpectGet("abc1234").RedisNil()

// 	repo := new(mocks.MockLinksRepo)
// 	repo.On("GetByLink", "abc1234").Return(nil, errors.New("db error"))

// 	cached := infrastructure.NewCachedRepoWith(repo, client, cacheTTL)

// 	link, err := cached.GetByLink("abc1234")

// 	assert.Error(t, err)
// 	assert.Nil(t, link)
// 	repo.AssertExpectations(t)
// }

// func TestDeleteItemByShortLink_InvalidatesCache(t *testing.T) {
// 	client, rMock := redismock.NewClientMock()
// 	rMock.ExpectDel("abc1234").SetVal(1)

// 	repo := new(mocks.MockLinksRepo)
// 	repo.On("DeleteItemByShortLink", "abc1234").Return(true, nil)

// 	cached := infrastructure.NewCachedRepoWith(repo, client, cacheTTL)

// 	ok, err := cached.DeleteItemByShortLink("abc1234")

// 	assert.NoError(t, err)
// 	assert.True(t, ok)
// 	repo.AssertExpectations(t)
// 	assert.NoError(t, rMock.ExpectationsWereMet())
// }

// func TestDeleteItemByShortLink_DBError(t *testing.T) {
// 	client, rMock := redismock.NewClientMock()
// 	rMock.ExpectDel("abc1234").SetVal(1)

// 	repo := new(mocks.MockLinksRepo)
// 	repo.On("DeleteItemByShortLink", "abc1234").Return(false, errors.New("db error"))

// 	cached := infrastructure.NewCachedRepoWith(repo, client, cacheTTL)

// 	ok, err := cached.DeleteItemByShortLink("abc1234")

// 	assert.Error(t, err)
// 	assert.False(t, ok)
// 	repo.AssertExpectations(t)
// }

// func TestCheckExsist_DelegatesToDB(t *testing.T) {
// 	client, _ := redismock.NewClientMock()
// 	repo := new(mocks.MockLinksRepo)
// 	repo.On("CheckExsist", "abc1234").Return(true, nil)

// 	cached := infrastructure.NewCachedRepoWith(repo, client, cacheTTL)

// 	exists, err := cached.CheckExsist("abc1234")

// 	assert.NoError(t, err)
// 	assert.True(t, exists)
// 	repo.AssertExpectations(t)
// }

// func TestTryAddItem_DelegatesToDB(t *testing.T) {
// 	client, _ := redismock.NewClientMock()
// 	repo := new(mocks.MockLinksRepo)
// 	newLink := &domain.Link{ShortCode: "abc1234", FullUrl: "https://example.com"}
// 	repo.On("TryAddItem", newLink).Return(true, nil)

// 	cached := infrastructure.NewCachedRepoWith(repo, client, cacheTTL)

// 	ok, err := cached.TryAddItem(newLink)

// 	assert.NoError(t, err)
// 	assert.True(t, ok)
// 	repo.AssertExpectations(t)
// }
