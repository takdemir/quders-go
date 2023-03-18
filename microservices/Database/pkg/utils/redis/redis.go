package redis

import (
	"context"
	"database/pkg/utils/common"
	"errors"
	"github.com/redis/go-redis/v9"
	"os"
	"strings"
)

var rdb *redis.Client
var ctx context.Context

type QudersRedis struct {
}

func (br *QudersRedis) Connect(redisEnvName string) {
	connectionStrings := ConnectionStringParser(os.Getenv(redisEnvName))
	if len(connectionStrings) < 4 {
		errorMessage := "redis connection string error"
		common.FailOnError(errors.New(errorMessage), errorMessage)
		panic(errorMessage)
	}
	_, isHostExist := connectionStrings["host"]
	if !isHostExist {
		errorMessage := "redis connection string host error"
		common.FailOnError(errors.New(errorMessage), errorMessage)
		panic(errorMessage)
	}

	_, isUsernameExist := connectionStrings["username"]
	if !isUsernameExist {
		errorMessage := "redis connection string username error"
		common.FailOnError(errors.New(errorMessage), errorMessage)
		panic(errorMessage)
	}

	_, isPasswordExist := connectionStrings["password"]
	if !isPasswordExist {
		errorMessage := "redis connection string password error"
		common.FailOnError(errors.New(errorMessage), errorMessage)
		panic(errorMessage)
	}

	_, isSchemeExist := connectionStrings["scheme"]
	if !isSchemeExist {
		errorMessage := "redis connection string scheme error"
		common.FailOnError(errors.New(errorMessage), errorMessage)
		panic(errorMessage)
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     connectionStrings["host"],
		Password: connectionStrings["password"],
		DB:       0, // use default DB
	})
	ctx = context.Background()
}

func (br *QudersRedis) Get(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (br *QudersRedis) Set(key string, value interface{}) error {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func ConnectionStringParser(connectionString string) map[string]string {
	explodeFromAt := strings.Split(connectionString, "@")
	if len(explodeFromAt) < 2 {
		errorMessage := "no @ in connection string"
		common.FailOnError(errors.New(errorMessage), errorMessage)
		panic(errorMessage)
	}
	explodeFromComma := strings.Split(explodeFromAt[0], ":")
	if len(explodeFromAt) < 2 {
		errorMessage := "no : in connection string"
		common.FailOnError(errors.New(errorMessage), errorMessage)
		panic(errorMessage)
	}
	explodeFromQuestionMark := strings.Split(explodeFromAt[1], "?")
	if len(explodeFromAt) < 2 {
		errorMessage := "no ? in connection string"
		common.FailOnError(errors.New(errorMessage), errorMessage)
		panic(errorMessage)
	}

	return map[string]string{
		"host":     explodeFromQuestionMark[0],
		"username": explodeFromComma[0],
		"password": explodeFromComma[1],
		"scheme":   explodeFromQuestionMark[1],
	}

}
