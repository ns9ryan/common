package config

import (
	"context"
	"crypto/tls"
	"errors"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"
	"github.com/zeromicro/go-zero/core/logx"
)

// RedisConf Redis 配置
type RedisConf struct {
	Host         string `json:",env=REDIS_HOST"`                               // Redis 地址，多个地址使用逗号分隔
	Db           int    `json:",default=0,env=REDIS_DB"`                       // Redis 数据库编号
	Mode         string `json:",optional,default=single,env=REDIS_MODE"`       // Redis 模式：single 或 cluster
	Username     string `json:",optional,env=REDIS_USERNAME"`                  // Redis 用户名
	Pass         string `json:",optional,env=REDIS_PASSWORD"`                  // Redis 密码
	Tls          bool   `json:",optional,env=REDIS_TLS"`                       // 是否启用 TLS
	Master       string `json:",optional,env=REDIS_MASTER"`                    // Redis Master 名称
	PoolSize     int    `json:",optional,default=48,env=REDIS_POOL_SIZE"`      // 连接池大小
	MaxIdleConns int    `json:",optional,default=12,env=REDIS_MAX_IDLE_CONNS"` // 最大空闲连接数
}

const (
	RedisModeSingle  = "single"  // 单机模式
	RedisModeCluster = "cluster" // 集群模式
)

// Validate 校验 Redis 配置
func (r RedisConf) Validate() error {
	if len(r.Host) == 0 {
		return errors.New("host cannot be empty")
	}
	return nil
}

// EffectiveMode 获取实际使用的 Redis 模式
func (r RedisConf) EffectiveMode() string {
	// 配置多个地址时自动使用集群模式
	if strings.Contains(r.Host, ",") {
		return RedisModeCluster
	}

	// 显式配置 cluster 时使用集群模式
	if strings.EqualFold(strings.TrimSpace(r.Mode), RedisModeCluster) {
		return RedisModeCluster
	}

	return RedisModeSingle
}

// IsClusterMode 判断是否为 Redis 集群模式
func (r RedisConf) IsClusterMode() bool {
	return r.EffectiveMode() == RedisModeCluster
}

// parsedAddrs 解析 Redis 地址并去除空值和重复地址
func (r RedisConf) parsedAddrs() []string {
	parts := strings.Split(r.Host, ",")           // 拆分 Redis 地址
	addrs := make([]string, 0, len(parts))        // 保存有效地址
	seen := make(map[string]struct{}, len(parts)) // 记录已存在的地址

	for _, part := range parts {
		addr := strings.TrimSpace(part)
		if addr == "" {
			continue
		}

		// 跳过重复地址
		if _, ok := seen[addr]; ok {
			continue
		}

		seen[addr] = struct{}{}
		addrs = append(addrs, addr)
	}

	// 没有解析到有效地址时返回原始地址
	if len(addrs) == 0 {
		return []string{strings.TrimSpace(r.Host)}
	}

	return addrs
}

// NewUniversalRedis 创建 Redis UniversalClient
func (r RedisConf) NewUniversalRedis() (redis.UniversalClient, error) {
	// 校验 Redis 配置
	err := r.Validate()
	if err != nil {
		return nil, err
	}

	// 创建 Redis 客户端配置
	opt := &redis.UniversalOptions{
		Addrs:         r.parsedAddrs(),   // Redis 节点地址
		IsClusterMode: r.IsClusterMode(), // 是否使用集群模式
		DB:            r.Db,              // Redis 数据库编号
		Password:      r.Pass,            // Redis 密码
		Username:      r.Username,        // Redis 用户名
		PoolSize:      r.PoolSize,        // 连接池大小
		MaxIdleConns:  r.MaxIdleConns,    // 最大空闲连接数

		// Redis 维护通知配置
		MaintNotificationsConfig: &maintnotifications.Config{
			Mode: maintnotifications.ModeDisabled, // 关闭 Redis 维护通知
		},
	}

	// 设置 Master 名称
	if r.Master != "" {
		opt.MasterName = r.Master
	}

	// 启用 TLS
	if r.Tls {
		opt.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
	}

	// 创建 Redis 客户端
	rds := redis.NewUniversalClient(opt)

	// 检查 Redis 连接
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = rds.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return rds, nil
}

// MustNewUniversalRedis 创建 Redis UniversalClient，失败时直接终止程序
func (r RedisConf) MustNewUniversalRedis() redis.UniversalClient {
	rds, err := r.NewUniversalRedis()
	logx.Must(err)

	return rds
}
