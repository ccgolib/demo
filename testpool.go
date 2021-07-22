package main

import (
	"golang.org/x/exp/rand"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type factory func(host string) conn
type conn interface {
	Close() error
}
type pool struct {
	m    map[string]chan conn
	mu   sync.RWMutex
	fact factory
}

func main() {
	var pts uint64
	p := &pool{
		m:  make(map[string]chan conn),
		mu: sync.RWMutex{},
		fact: func(target string) conn {
			c, _ := net.Dial("", "8080")
			return c
		},
	}
	//打印线程，打印get，put效率
	be := time.Now()
	go func() {
		for true {
			//此处先休眠一秒是为了避免第一次时差计算为0导致的除法错误
			time.Sleep(1 * time.Second)
			cost := time.Since(be) / time.Second
			println(atomic.LoadUint64(&pts)/uint64(cost), "pt/s")
		}
	}()
	time.Sleep(1 * time.Second)
	//打印线程完，此处等待一秒是为对应打印线程第一次休眠，尽量减少误差

	//集群规模
	hosts := []string{"192.168.0.1", "192.168.0.2", "192.168.0.3", "192.168.0.4"}
	//并发线程数量
	threadnum := 1
	for i := 0; i < threadnum; i++ {
		go func() {
			for true {
				target := hosts[rand.Int()%len(hosts)]
				conn := p.Get3(target)
				//------------------使用连接开始
				//time.Sleep(1*time.Nanosecond)
				//------------------使用连接完毕
				p.Put3(target, conn)
				atomic.AddUint64(&pts, 1)
			}
		}()
	}
	time.Sleep(100 * time.Second)
}

func (p *pool) Get(host string) (c conn) {
	if _, ok := p.m[host]; !ok {
		p.m[host] = make(chan conn, 100)
	}
	select {
	case c = <-p.m[host]:
		{
		}
	default:
		c = p.New(host)
	}
	return
}
func (p *pool) Put(host string, c conn) {
	select {
	case p.m[host] <- c:
		{
		}
	default:
		c.Close()
	}
}
func (p *pool) New(host string) conn {
	return p.fact(host)
}

func (p *pool) Get3(host string) (c conn) {
	p.mu.RLock()
	if _, ok := p.m[host]; !ok {
		p.mu.RUnlock()
		p.mu.Lock()
		p.m[host] = make(chan conn, 100)
		p.mu.Unlock()
	} else {
		p.mu.RUnlock()
	}
	p.mu.RLock()
	defer p.mu.RUnlock()
	select {
	case c = <-p.m[host]:
		{
		}
	default:
		c = p.New(host)
	}
	return
}
func (p *pool) Put3(host string, c conn) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	select {
	case p.m[host] <- c:
		{
		}
	default:
		c.Close()
	}
}
