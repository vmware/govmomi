// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"container/list"
	"sync"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

var (
	maxPageSize           = 1000
	maxCollectors   int32 = 32 // the VC default limit
	defaultPageSize       = 10

	errInvalidArgMaxCount = &types.InvalidArgument{InvalidProperty: "maxCount"}
)

type HistoryCollector struct {
	parent *history
	root   types.ManagedObjectReference
	size   int

	fill func(*Context)
	page *list.List
	pos  *list.Element
}

type history struct {
	sync.Mutex

	page       *list.List
	collectors map[types.ManagedObjectReference]mo.Reference
}

func newHistory() *history {
	return &history{
		page:       list.New(),
		collectors: make(map[types.ManagedObjectReference]mo.Reference),
	}
}

func (c *history) add(ctx *Context, collector mo.Reference) types.ManagedObjectReference {
	ref := ctx.Session.Put(collector).Reference()

	c.Lock()
	c.collectors[ref] = collector
	c.Unlock()

	return ref
}

func (c *history) remove(ctx *Context, ref types.ManagedObjectReference) {
	ctx.Session.Remove(ctx, ref)

	c.Lock()
	delete(c.collectors, ref)
	c.Unlock()
}

func newHistoryCollector(ctx *Context, h *history, size int) *HistoryCollector {
	return &HistoryCollector{
		parent: h,
		root:   ctx.Map.content().RootFolder,
		size:   size,
		page:   list.New(),
	}
}

func (c *HistoryCollector) SetCollectorPageSize(ctx *Context, req *types.SetCollectorPageSize) soap.HasFault {
	body := new(methods.SetCollectorPageSizeBody)
	size, err := validatePageSize(req.MaxCount)
	if err != nil {
		body.Fault_ = err
		return body
	}

	c.size = size
	c.page = list.New()
	c.fill(ctx)

	body.Res = new(types.SetCollectorPageSizeResponse)
	return body
}

func (c *HistoryCollector) ResetCollector(ctx *Context, req *types.ResetCollector) soap.HasFault {
	c.pos = c.page.Back()

	return &methods.ResetCollectorBody{
		Res: new(types.ResetCollectorResponse),
	}
}

func (c *HistoryCollector) RewindCollector(ctx *Context, req *types.RewindCollector) soap.HasFault {
	c.pos = c.page.Front()

	return &methods.RewindCollectorBody{
		Res: new(types.RewindCollectorResponse),
	}
}

func (c *HistoryCollector) DestroyCollector(ctx *Context, req *types.DestroyCollector) soap.HasFault {
	ctx.Session.Remove(ctx, req.This)

	c.parent.remove(ctx, req.This)

	return &methods.DestroyCollectorBody{
		Res: new(types.DestroyCollectorResponse),
	}
}

func (c *HistoryCollector) read(max int32, next func() *list.Element, take func(*list.Element)) {
	for i := 0; i < int(max); i++ {
		e := next()
		if e == nil {
			break
		}

		take(e)
		c.pos = e
	}
}

func (c *HistoryCollector) next(max int32, take func(*list.Element)) {
	next := func() *list.Element {
		if c.pos != nil {
			return c.pos.Next()
		}
		return c.page.Front()
	}

	c.read(max, next, take)
}

func (c *HistoryCollector) prev(max int32, take func(*list.Element)) {
	next := func() *list.Element {
		if c.pos != nil {
			return c.pos.Prev()
		}
		return c.page.Back()
	}

	c.read(max, next, take)
}

func pushHistory(l *list.List, item types.AnyType) {
	if l.Len() > maxPageSize {
		l.Remove(l.Front()) // Prune history
	}
	l.PushBack(item)
}

func validatePageSize(count int32) (int, *soap.Fault) {
	size := int(count)

	if size == 0 {
		size = defaultPageSize
	} else if size < 0 || size > maxPageSize {
		return -1, Fault("", errInvalidArgMaxCount)
	}

	return size, nil
}
