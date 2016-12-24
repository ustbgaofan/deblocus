package glog

import (
	"container/list"
	"sync"
	"sync/atomic"

	"github.com/imkira/go-observer"
)

type LogQueue struct {
	sync.RWMutex
	capacity int
	seq      uint32
	list     *list.List
	property observer.Property
}

type LogEntry struct {
	Id  uint32
	Msg string
}

var LogQueueInstance = LogQueue{
	capacity: 1024,
	list:     list.New(),
	property: observer.NewProperty(1),
}

func (l *LogQueue) appendBuffer(b *buffer) {
	var id = atomic.AddUint32(&l.seq, 1)
	var entry = &LogEntry{
		Id:  id,
		Msg: b.String(),
	}
	l.property.Update(entry)
	l.addNewEntry(entry)
}

func (l *LogQueue) trimQueue() {
	var ls = l.list
	for ls.Len() > l.capacity {
		ls.Remove(ls.Front())
	}
}

func (l *LogQueue) addNewEntry(entry *LogEntry) {
	l.Lock()
	defer l.Unlock()
	l.list.PushBack(entry)
	l.trimQueue()
}

func (l *LogQueue) NewMonitor() observer.Stream {
	return l.property.Observe()
}

func (l *LogQueue) Retrieve(callback func(string) bool) {
	l.RLock()
	defer l.RUnlock()
	var ls = l.list
	var entry *LogEntry
	for i, e := ls.Len(), ls.Front(); i > 0; i, e = i-1, e.Next() {
		if e != nil {
			entry = e.Value.(*LogEntry)
			if !callback(entry.Msg) {
				return
			}
		}
	}
}
