package pkg_redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisConnectivityRepository struct {
	ctx    context.Context
	client *redis.Client
}

func NewRedisConnectivityRepository(ctx context.Context, client *redis.Client) *RedisConnectivityRepository {
	return &RedisConnectivityRepository{
		ctx:    ctx,
		client: client,
	}
}

func (r *RedisConnectivityRepository) IsConnected(clientId string) (bool, error) {
	clientId = generateKey(clientId)

	_, err := r.client.Get(r.ctx, clientId).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (r *RedisConnectivityRepository) GetConnectionId(clientId string) (string, error) {
	clientId = generateKey(clientId)

	connectionId, err := r.client.Get(r.ctx, clientId).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	} else {
		return connectionId, nil
	}
}

func (r *RedisConnectivityRepository) SaveConnection(clientId string, connectionId string) error {
	clientId = generateKey(clientId)

	if err := r.client.Set(r.ctx, clientId, connectionId, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (r *RedisConnectivityRepository) RemoveConnection(clientId string) error {
	clientId = generateKey(clientId)

	if err := r.client.Del(r.ctx, clientId).Err(); err != nil {
		return err
	}
	return nil
}

func generateKey(clientId string) string {
	return "things:connectivity:" + clientId
}
