package comms

import (
	"fmt"
	"strings"
	"github.com/mediocregopher/radix"
)

// Interface for functional usage
type RedisItem interface {
	GetID() string
	Key(string) RedisItem

	Get() string
	Integer() int
	Set(string)
	Incr() int

	Map() map[string]string
	KeyVal(string, string) int
	KeyValNX(string, string) int

	Has(string) bool
	Members() []string
	Add(string)
	Rem(string)
	Establish(string) RedisItem
}

func NewRedis(address string) Redis {
	var clusterAddress string
	if(strings.Contains(address, ":")) {
		clusterAddress = address
	} else {
		clusterAddress = address + ":6379"
	}
	redis, err := radix.NewCluster([]string{"tcp://" + clusterAddress})
	if(err != nil) {
		fmt.Printf("Cannot connect to Redis: %s\n", err)
		panic(err)
	}
	return Redis {redis, ""}
}

/*
	Primary Redis holding structure.
	Meets radix.Client interface.
*/
type Redis struct {
	Client* radix.Cluster
	Id string
}

func (r* Redis) Do(act radix.Action) error {
	return r.Client.Do(act)
}
func (r* Redis) Close() error {
	return r.Client.Close()
}

func (r Redis) GetID() string {
	return r.Id
}
func (r Redis) Key(key string) RedisItem {
	if(r.GetID() != "") {
		return Redis {r.Client, r.GetID() + "." + key}
	}
	return Redis {r.Client, key}
}

func (k Redis) Get() string {
	var val string
	err := k.Do(radix.Cmd(&val, "GET", k.Id))
	if(err != nil) {
		fmt.Printf("CANNOT GET KEY: %s\n", k.Id)
		panic(err)
	}
	return val
}
func (k Redis) Integer() int {
	var val int
	err := k.Do(radix.Cmd(&val, "GET", k.Id))
	if(err != nil) {
		fmt.Printf("CANNOT GET KEY: %s\n", k.Id)
		panic(err)
	}
	return val
}
func (k Redis) Set(val string) {
	err := k.Do(radix.Cmd(nil, "SET", k.Id, val))
	if(err != nil) {
		fmt.Printf("CANNOT SET KEY: %s %s\n", k.Id, val)
		panic(err)
	}
}
func (k Redis) Incr() int {
	var val int
	err := k.Do(radix.Cmd(&val, "INCR", k.Id))
	if(err != nil) {
		fmt.Printf("CANNOT INCR KEY: %s\n", k.Id)
		panic(err)
	}
	return val
}

func (m Redis) Map() map[string]string {
	var val map[string]string
	err := m.Do(radix.Cmd(&val, "HGETALL", m.Id))
	if(err != nil) {
		fmt.Printf("CANNOT HGETALL KEY: %s\n", m.Id)
		panic(err)
	}
	return val
}
func (m Redis) KeyVal(k string, v string) int {
	var val int
	err := m.Do(radix.Cmd(&val, "HSET", m.Id, k, v))
	if(err != nil) {
		fmt.Printf("CANNOT HSET KEY/K/V: %s %s %s\n", m.Id, k, v)
		panic(err)
	}
	return val
}
func (m Redis) KeyValNX(k string, v string) int {
	var val int
	err := m.Do(radix.Cmd(&val, "HSETNX", m.Id, k, v))
	if(err != nil) {
		fmt.Printf("CANNOT HSETNX KEY/K/V: %s %s %s\n", m.Id, k, v)
		panic(err)
	}
	return val
}

func (s Redis) Has(value string) bool {
	var val int
	err := s.Do(radix.Cmd(&value, "SISMEMBER", s.Id))
	if(err != nil) {
		fmt.Printf("CANNOT GET SET VALUES: %s\n", s.Id)
		panic(err)
	}
	return (val == 1)
}
func (s Redis) Members() []string {
	var val []string
	err := s.Do(radix.Cmd(&val, "SMEMBERS", s.Id))
	if(err != nil) {
		fmt.Printf("CANNOT GET SET VALUES: %s\n", s.Id)
		panic(err)
	}
	return val
}
func (s Redis) Add(val string) {
	err := s.Do(radix.Cmd(nil, "SADD", s.Id, val))
	if(err != nil) {
		fmt.Printf("CANNOT ADD TO SET: %s %s\n", s.Id, val)
		panic(err)
	}
}
func (s Redis) Rem(val string) {
	err := s.Do(radix.Cmd(nil, "SREM", s.Id, val))
	if(err != nil) {
		fmt.Printf("CANNOT REMOVE FROM SET: %s %s\n", s.Id, val)
		panic(err)
	}
}

// This both adds an id to the set, and then accesses its subkey.
func (s Redis) Establish(val string) RedisItem {
	err := s.Do(radix.Cmd(nil, "SADD", s.Id, val))
	if(err != nil) {
		fmt.Printf("CANNOT ADD TO SET: %s %s\n", s.Id, val)
		panic(err)
	}
	return Redis {s.Client, s.Id + "." + val}
}