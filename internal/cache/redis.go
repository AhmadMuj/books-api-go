package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AhmadMuj/books-api-go/internal/config"
	"github.com/AhmadMuj/books-api-go/internal/models"
	"github.com/redis/go-redis/v9"
)

const (
	bookKeyPrefix     = "book:"
	bookListKeyPrefix = "books:page:"
	defaultExpiration = 24 * time.Hour
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(cfg *config.Config) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RedisCache{
		client: client,
	}, nil
}

func (c *RedisCache) GetBook(ctx context.Context, id uint) (*models.Book, error) {
	key := fmt.Sprintf("%s%d", bookKeyPrefix, id)
	data, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var book models.Book
	if err := json.Unmarshal(data, &book); err != nil {
		return nil, err
	}

	return &book, nil
}

func (c *RedisCache) SetBook(ctx context.Context, book *models.Book) error {
	data, err := json.Marshal(book)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s%d", bookKeyPrefix, book.ID)
	return c.client.Set(ctx, key, data, defaultExpiration).Err()
}

func (c *RedisCache) DeleteBook(ctx context.Context, id uint) error {
	key := fmt.Sprintf("%s%d", bookKeyPrefix, id)
	return c.client.Del(ctx, key).Err()
}

func (c *RedisCache) GetBooksList(ctx context.Context, page, pageSize int) ([]models.Book, int64, error) {
	key := fmt.Sprintf("%s%d:%d", bookListKeyPrefix, page, pageSize)

	// Get cached data
	data, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	var result struct {
		Books []models.Book `json:"books"`
		Total int64         `json:"total"`
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, 0, err
	}

	return result.Books, result.Total, nil
}

func (c *RedisCache) SetBooksList(ctx context.Context, books []models.Book, total int64, page, pageSize int) error {
	data, err := json.Marshal(struct {
		Books []models.Book `json:"books"`
		Total int64         `json:"total"`
	}{
		Books: books,
		Total: total,
	})
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s%d:%d", bookListKeyPrefix, page, pageSize)
	return c.client.Set(ctx, key, data, defaultExpiration).Err()
}

func (c *RedisCache) InvalidateBooksList(ctx context.Context) error {
	pattern := bookListKeyPrefix + "*"
	keys, err := c.client.Keys(ctx, pattern).Result()
	if err != nil {
		return err
	}

	if len(keys) > 0 {
		return c.client.Del(ctx, keys...).Err()
	}

	return nil
}

func (c *RedisCache) Clear(ctx context.Context) error {
	return c.client.FlushDB(ctx).Err()
}

func (c *RedisCache) Close() error {
	return c.client.Close()
}
