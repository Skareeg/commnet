package comms

import (
	nats "github.com/nats-io/nats.go"
)

func NewComm(addressRedis string, addressNATS string) Comm {
	return Comm {NewRedis(addressRedis), NewNATS(addressNATS)}
}

/*
	Holds the Redis/NATS cluster and provides a basic way to interact with it.
*/
type Comm struct {
	Memory RedisItem
	Web* nats.EncodedConn
}

func (c Comm) GetID() string {
	return c.Memory.GetID()
}
func (c Comm) Key(key string) Comm {
	return Comm {c.Memory.Key(key), c.Web}
}


func (c Comm) Get() string {
	return c.Memory.Get()
}
func (c Comm) Integer() int {
	return c.Memory.Integer()
}
func (c Comm) Set(val string) {
	c.Memory.Set(val)
}
func (c Comm) Incr() int {
	return c.Memory.Incr()
}

func (c Comm) Map() map[string]string {
	return c.Memory.Map()
}
func (c Comm) KeyVal(key string, val string) int {
	return c.Memory.KeyVal(key, val)
}
func (c Comm) KeyValNX(key string, val string) int {
	return c.Memory.KeyVal(key, val)
}

func (c Comm) Has(val string) bool {
	return c.Memory.Has(val)
}
func (c Comm) Members() []string {
	return c.Memory.Members()
}
func (c Comm) Add(val string) {
	c.Memory.Add(val)
}
func (c Comm) Rem(val string) {
	c.Memory.Rem(val)
}
func (c Comm) Establish(val string) RedisItem {
	return c.Memory.Establish(val)
}

func (c Comm) SendString(val string) {
	c.Web.Publish(c.Memory.GetID(), val)
}
func (c Comm) SendInt(val int) {
	c.Web.Publish(c.Memory.GetID(), val)
}
func (c Comm) SendFloat(val float64) {
	c.Web.Publish(c.Memory.GetID(), val)
}
func (c Comm) Recv(handler nats.Handler) {
	c.Web.Subscribe(c.Memory.GetID(), handler)
}