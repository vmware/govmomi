/*
Copyright (c) 2017-2018 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package simulator

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type SessionManager struct {
	mo.SessionManager

	ServiceHostName string

	*sessionMap
}

func NewSessionManager(ref types.ManagedObjectReference) object.Reference {
	s := &SessionManager{
		sessionMap: &sessionMap{sessions: make(map[string]Session)},
	}
	s.Self = ref
	return s
}

func (s *SessionManager) Login(ctx *Context, login *types.Login) soap.HasFault {
	body := &methods.LoginBody{}

	if login.Locale == "" {
		login.Locale = session.Locale
	}

	if login.UserName == "" || login.Password == "" || ctx.Session != nil {
		body.Fault_ = invalidLogin
	} else {
		session := Session{
			UserSession: types.UserSession{
				Key:            uuid.New().String(),
				UserName:       login.UserName,
				FullName:       login.UserName,
				LoginTime:      time.Now(),
				LastActiveTime: time.Now(),
				Locale:         login.Locale,
				MessageLocale:  login.Locale,
			},
			Registry: NewRegistry(),
		}

		ctx.SetSession(session, true)

		body.Res = &types.LoginResponse{
			Returnval: session.UserSession,
		}
	}

	return body
}

func (s *SessionManager) Logout(ctx *Context, _ *types.Logout) soap.HasFault {
	s.remove(ctx.Session.Key)
	return &methods.LogoutBody{Res: new(types.LogoutResponse)}
}

func (s *SessionManager) TerminateSession(ctx *Context, req *types.TerminateSession) soap.HasFault {
	body := new(methods.TerminateSessionBody)

	for _, id := range req.SessionId {
		if id == ctx.Session.Key {
			body.Fault_ = Fault("", new(types.InvalidArgument))
			return body
		}
		s.remove(id)
	}

	body.Res = new(types.TerminateSessionResponse)
	return body
}

func (s *SessionManager) AcquireCloneTicket(ctx *Context, _ *types.AcquireCloneTicket) soap.HasFault {
	session := *ctx.Session
	session.Key = uuid.New().String()
	s.save(session)

	return &methods.AcquireCloneTicketBody{
		Res: &types.AcquireCloneTicketResponse{
			Returnval: session.Key,
		},
	}
}

func (s *SessionManager) CloneSession(ctx *Context, ticket *types.CloneSession) soap.HasFault {
	body := new(methods.CloneSessionBody)

	s.mu.Lock()
	session, exists := s.sessions[ticket.CloneTicket]
	s.mu.Unlock()

	if exists {
		s.remove(ticket.CloneTicket) // A clone ticket can only be used once
		session.Key = uuid.New().String()
		ctx.SetSession(session, true)

		body.Res = &types.CloneSessionResponse{
			Returnval: session.UserSession,
		}
	} else {
		body.Fault_ = invalidLogin
	}

	return body
}

func (s *SessionManager) AcquireGenericServiceTicket(ticket *types.AcquireGenericServiceTicket) soap.HasFault {
	return &methods.AcquireGenericServiceTicketBody{
		Res: &types.AcquireGenericServiceTicketResponse{
			Returnval: types.SessionManagerGenericServiceTicket{
				Id:       uuid.New().String(),
				HostName: s.ServiceHostName,
			},
		},
	}
}

// sessionMap allows Session.Get to clone SessionManager, copying the mo.SessionManager field, but sharing the sessionMap fields.
type sessionMap struct {
	mu       sync.Mutex
	sessions map[string]Session
}

func (s *sessionMap) remove(key string) {
	s.mu.Lock()
	delete(s.sessions, key)
	s.mu.Unlock()
}

func (s *sessionMap) load(key string) (Session, bool) {
	s.mu.Lock()
	session, ok := s.sessions[key]
	s.mu.Unlock()
	return session, ok
}

func (s *sessionMap) save(session Session) {
	s.mu.Lock()
	s.sessions[session.Key] = session
	s.mu.Unlock()
}

// internalContext is the session for use by the in-memory client (Service.RoundTrip)
var internalContext = &Context{
	Context: context.Background(),
	Session: &Session{
		UserSession: types.UserSession{
			Key: uuid.New().String(),
		},
		Registry: NewRegistry(),
	},
}

var invalidLogin = Fault("Login failure", new(types.InvalidLogin))

// Context provides per-request Session management.
type Context struct {
	req *http.Request
	res http.ResponseWriter
	m   *SessionManager

	context.Context
	Session *Session
}

// mapSession maps an HTTP cookie to a Session.
func (c *Context) mapSession() {
	if cookie, err := c.req.Cookie(soap.SessionCookieName); err == nil {
		if val, ok := c.m.load(cookie.Value); ok {
			c.SetSession(val, false)
		}
	}
}

// SetSession should be called after successful authentication.
func (c *Context) SetSession(session Session, login bool) {
	session.UserAgent = c.req.UserAgent()
	session.IpAddress = strings.Split(c.req.RemoteAddr, ":")[0]
	session.LastActiveTime = time.Now()

	c.m.save(session)

	c.Session = &session

	if login {
		http.SetCookie(c.res, &http.Cookie{
			Name:  soap.SessionCookieName,
			Value: session.Key,
		})
	}
}

// Session combines a UserSession and a Registry for per-session managed objects.
type Session struct {
	types.UserSession
	*Registry
}

// Put wraps Registry.Put, setting the moref value to include the session key.
func (s *Session) Put(item mo.Reference) mo.Reference {
	ref := item.Reference()
	if ref.Value == "" {
		ref.Value = fmt.Sprintf("session[%s]%s", s.Key, uuid.New())
	}
	s.Registry.setReference(item, ref)
	return s.Registry.Put(item)
}

// Get wraps Registry.Get, session-izing singleton objects such as SessionManager and the root PropertyCollector.
func (s *Session) Get(ref types.ManagedObjectReference) mo.Reference {
	obj := s.Registry.Get(ref)
	if obj != nil {
		return obj
	}

	switch ref.Type {
	case "SessionManager":
		// Clone SessionManager so the PropertyCollector can properly report CurrentSession
		m := *Map.SessionManager()
		m.CurrentSession = &s.UserSession

		// TODO: we could maintain SessionList as part of the SessionManager singleton
		m.mu.Lock()
		for _, session := range m.sessions {
			m.SessionList = append(m.SessionList, session.UserSession)
		}
		m.mu.Unlock()

		return &m
	case "PropertyCollector":
		if ref == Map.content().PropertyCollector {
			return s.Put(NewPropertyCollector(ref))
		}
	}

	return Map.Get(ref)
}
