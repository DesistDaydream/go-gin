package session

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

// RedisSessionData 表示一个用户的 SessionData 应该具有的属性
type RedisSessionData struct {
	ID   string
	Data map[string]interface{}
	// 读写锁，锁的是上面的 Data
	rwLock sync.RWMutex
	// Redis 中保存的 SessionData 的过期时间
	expired int
	// Redis 连接
	client *redis.Client
}

// NewRedisSessionData 实例化 RedisSessionData
func NewRedisSessionData(id string, client *redis.Client) Data {
	return &RedisSessionData{
		ID:     id,
		Data:   make(map[string]interface{}, 8),
		client: client,
	}
}

// GetID is
func (r *RedisSessionData) GetID() string {
	return r.ID
}

// Get Data 支持的操作,根据给定的 key 获取值
func (r *RedisSessionData) Get(key string) (values interface{}, err error) {
	// 获取读锁
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()

	value, ok := r.Data[key]
	if !ok {
		err = fmt.Errorf("invalid key")
		return
	}

	return value, nil
}

// Set Data 支持的操作,根据给定的 k/v 设定这些值
func (r *RedisSessionData) Set(key string, value interface{}) {
	// 获取读锁
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	r.Data[key] = value
}

// Del Data 支持的操作,根据给定的 key，删除对应的 k/v 对
func (r *RedisSessionData) Del(key string) {
	// 获取读锁
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	delete(r.Data, key)
}

// Save 保存 SessionData
func (r *RedisSessionData) Save() {
	var (
		value []byte
		err   error
	)
	// 将最新的 SessionData 保存到 Redis 中
	if value, err = json.Marshal(r.Data); err != nil {
		fmt.Printf("marshal SessionData 失败:%v\n", err)
		return
	}
	// 将数据保存到 Redis
	r.client.Set(r.ID, value, time.Second*time.Duration(r.expired))
}

// RedisManager 存储 SessionData 的 Redis 后端管理器
type RedisManager struct {
	Session map[string]Data
	rwLock  sync.RWMutex
	// Redis 连接池
	client *redis.Client
}

// NewRedisManager 实例化存储 SessionData 的 Redis 后端
func NewRedisManager() Manager {
	return &RedisManager{
		Session: make(map[string]Data, 1024),
	}
}

func (r *RedisManager) loadFromRedis(sessionID string) (err error) {
	var value string
	// 连接 Redis
	if value, err = r.client.Get(sessionID).Result(); err != nil {
		// redis 中没有该 SessionID 对应的 SessionData
		return
	}

	// 根据 SessionID 找到对应的数据
	// 把数据取出来反序列化到
	if err = json.Unmarshal([]byte(value), &r.Session); err != nil {
		// 从 Redis 中取出来的数据反序列化失败
		return
	}
	return
}

// Init 初始化
func (r *RedisManager) Init(addr string, options ...string) {
	var (
		password string
		dbString string
		db       int
		err      error
	)

	switch len(options) {
	case 1:
		password = options[0]
	case 2:
		password = options[0]
		dbString = options[1]
	}

	if db, err = strconv.Atoi(dbString); err != nil {
		db = 0
	}

	r.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if _, err = r.client.Ping().Result(); err != nil {
		panic(err)
	}
}

// GetSessionData 获取 SessionID 对应的 SessionData
func (r *RedisManager) GetSessionData(sessionID string) (d Data, err error) {
	// 如果 SessionData 为空，去 Redis 里根据 SessionID 加载 SessionData
	if r.Session == nil {
		if err = r.loadFromRedis(sessionID); err != nil {
			return nil, err
		}
	}
	// 然后根据 r.Session[sessionID] 拿到对应的 SessionData
	// r.Session[sessionID] 就是从存储 SessionData 的大 map 中根据 key 找到 SessionData
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()
	d, ok := r.Session[sessionID]
	if !ok {
		err = fmt.Errorf("无效的 SessionID")
	}
	return
}

// CreateSession 创建一条 Session 记录
func (r *RedisManager) CreateSession() (d Data) {
	uuidObj := uuid.New()
	d = NewRedisSessionData(uuidObj.String(), r.client)
	r.Session[d.GetID()] = d
	return
}
