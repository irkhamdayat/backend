package uploadfile

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Repository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func New() *Repository {
	return new(Repository)
}

func (r *Repository) WithPostgresDB(DB *gorm.DB) *Repository {
	r.db = DB
	return r
}

func (r *Repository) WithRedisClient(rdb *redis.Client) *Repository {
	r.rdb = rdb
	return r
}
