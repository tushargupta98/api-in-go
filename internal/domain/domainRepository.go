package domain

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/tushargupta98/api-in-go/cache"
	"github.com/tushargupta98/api-in-go/internal"
	"github.com/tushargupta98/api-in-go/logger"
)

const (
	cacheKeyAll  = "domain"
	cacheKeyById = "domain_%d"
)

type DomainRepository interface {
	List() ([]Domain, error)
	Create(domain Domain) (int, error)
	Get(id int) (*Domain, error)
	Update(id int, domain Domain) error
	Delete(id int) error
}

type domainRepository struct {
	*internal.BaseDBRepository
	*internal.BaseCacheRepository
}

var (
	selectAllDomainsQuery = "SELECT * FROM domain"
	selectDomainByIdQuery = "SELECT * FROM domain WHERE id = $1"
	insertDomainQuery     = "INSERT INTO domain (label) VALUES ($1) RETURNING id"
	updateDomainQuery     = "UPDATE domain SET label = $1 WHERE id = $2"
	deleteDomainQuery     = "DELETE FROM domain WHERE id = $1"
)

func NewDomainRepository(db *sqlx.DB, cache cache.RedisClient) DomainRepository {
	return &domainRepository{
		BaseDBRepository:    internal.NewBaseDBRepository(db),
		BaseCacheRepository: internal.NewBaseCacheRepository(cache),
	}
}

func (r *domainRepository) List() ([]Domain, error) {
	var domains []Domain
	if err := r.BaseCacheRepository.Get(cacheKeyAll, &domains); err != nil {
		if err := r.BaseDBRepository.List(selectAllDomainsQuery, nil, &domains); err != nil {
			return nil, err
		}
		if err := r.BaseCacheRepository.Set(cacheKeyAll, domains, 60); err != nil {
			return nil, err
		}
	}
	return domains, nil
}

func (r *domainRepository) Create(domain Domain) (int, error) {
	var id int
	id, err := r.BaseDBRepository.Create(insertDomainQuery, []interface{}{domain.Label})
	if err != nil {
		return 0, err
	}

	if err := r.BaseCacheRepository.Delete(cacheKeyAll); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *domainRepository) Get(id int) (*Domain, error) {
	var domain Domain
	cacheKey := fmt.Sprintf(cacheKeyById, id)
	if err := r.BaseCacheRepository.Get(cacheKey, &domain); err != nil {
		if err := r.BaseDBRepository.GetByID(selectDomainByIdQuery, []interface{}{id}, &domain); err != nil {
			logger.Logger.WithField("id", id).WithError(err).Error("Error fetching domain from database")
			return nil, err
		}
		if err := r.BaseCacheRepository.Set(cacheKey, domain, 60); err != nil {
			logger.Logger.Error("Error caching domain")
			return nil, err
		}
		logger.Logger.Info("Successfully fetched domain from database and cached")
	} else {
		logger.Logger.Info("Fetched domain from cache")
	}
	return &domain, nil
}

func (r *domainRepository) Update(id int, domain Domain) error {
	args := []interface{}{domain.Label, id}
	if err := r.BaseDBRepository.Update(updateDomainQuery, args); err != nil {
		return err
	}
	if err := r.BaseCacheRepository.Delete(cacheKeyAll); err != nil {
		return err
	}
	cacheKey := fmt.Sprintf(cacheKeyById, id)
	if err := r.BaseCacheRepository.Delete(cacheKey); err != nil {
		return err
	}
	return nil
}

func (r *domainRepository) Delete(id int) error {
	args := []interface{}{id}
	if err := r.BaseDBRepository.Delete(deleteDomainQuery, args); err != nil {
		return err
	}
	if err := r.BaseCacheRepository.Delete(cacheKeyAll); err != nil {
		return err
	}
	cacheKey := fmt.Sprintf(cacheKeyById, id)
	if err := r.BaseCacheRepository.Delete(cacheKey); err != nil {
		return err
	}
	return nil
}
