package dao

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
)

type Cache struct {
	conn redis.Conn
}

func NewCache(network, host string, port int) *Cache {
	conn, err := redis.Dial(network, fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		logs.Error(err)
		return nil
	}
	return &Cache{
		conn: conn,
	}
}

func (c *Cache) Set(key string, value []byte, expiredSecond int) error {
	var args []interface{}
	if expiredSecond == 0 {
		args = []interface{}{key, value}
	} else {
		args = []interface{}{key, value, "ex", expiredSecond}
	}
	if _, err := c.conn.Do("set", args...); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (c *Cache) Get(key string) ([]byte, error) {
	reply, err := c.conn.Do("get", key)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return reply.([]byte), nil
}

func (c *Cache) Del(key string) error {
	if _, err := c.conn.Do("del", key); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
