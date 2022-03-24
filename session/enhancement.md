# Enhanced VMWare Session Management Library

## Summary

The goal of this enhancement is to provide OpenShift, Kubernetes, and VMWare developers with a package that helps maintain VMWare API sessions using common patterns and helps prevent pitfalls of session management (i.e. leaks). 

## Motivation

Session management is a common area of concern in software engineering. Re-inventing or significantly tweaking the wheel is not encouraged due to the, often unnecessary, risk. However, users of the govmomi library have implemented their own extra session management. There are patterns across the various implementations which should be compiled into a single common package to reduce risk of one-off flaws in the code using VMWare sessions. 

### Goals

- Concurrently accessible, in-memory session caching
- Connection parameter and feature builder
- Ensure each VMWare APIs accessible in govmomi is addressed
- One-shot clients with callback that prevent session leaks
- Retain backwards compatibility with govmomi API


### Non-Goals

- New authn/authz methods
- Creating sessions and tokens


## Proposal

### Connection Parameter and Feature Builder

A builder pattern will be used for assembling connection parameters and features. The resulting object can then be passed to API client constructors. 

Connection Parameters are used for the process of building a connection to vSphere. 

Connection Features are anything used *after* the connection process. For example, keepalive.

```go=
type Features struct {
    KeepAlive bool
    KeepAliveTimeout time.Duration
    // TODO: Should caching be a per-client feature? 
    // Cache bool
}

type Parameters struct {
    server     *url.URL
    userinfo   *url.Userinfo
    thumbprint string
    insecure   bool
    timeout    time.Duration // Timeout during connection only
    userAgent  string
    
    features Features
}

func NewParams() *Params {
    return &Params{
        features: Feature {
            KeepAlive: false, // Disabled by default
            // Cache: true,
        }
    }
}

func (p *Params) WithUser(username string, password string) *Params {
    p.userinfo = url.UserPassword(username, password)
    return p
}

func ...
```

This builder comes from the Cluster API Provider for vSphere. It utilizes a [dev friendly builder pattern](https://github.com/kubernetes-sigs/cluster-api-provider-vsphere/blob/master/pkg/session/session.go#L53-L101) for assembling parameters and features before creating an API client.

### Client Construction

A new API Client constructor will be added. The constructor will accept `Parameters` and return a `*govmomi.Client`. This `*govmomi.Client` will be managed by the in-memory cache. Depending on the `Parameters`, a new client reference may be returned *or* an existing client reference may be returned from cache (e.g. a new session is required if a different username is specified). 

*Note* that `govmomi.Client` is composed of `vim25.Client`. All API clients inside the `govmomi` library are also composed of `vim25.Client` (or `soap.Client` which `vim25.Client` is composed of). This means that managing at the `govmomi.Client` level will be sufficient for developers using any other client.

```go=
var clientCache = map[string]*govmomi.Client{}
var clientCacheLock sync.Mutex

func NewCachedClient(ctx context.Context, u *url.URL, insecure bool, key string) (*govmomi.Client, error) {
    // ...
}
```

###  In-memory Cache
The in-memory cache is a `map[string]*govmomi.Client`. The `string` key will based on the server, username, and an optional label to uniquely identify clients within a single process. (e.g. same username but different reconcilation loops within an application).

Cache operations will be managed with a mutex for safe concurrent access.

Sessions validity will be checked any time one is pulled from a cache. If session is no longer valid, a new one will be created in its place. 

## Block Diagrams

*TODO*

### User Story

As an OpenShift, Kubernetes, or VMWare engineer, I want to have VMWare API session caching and management handled by govmomi so that I don't have to copy session management code or reimplement it myself.

### Implementation Details/Notes/Constraints [optional]


### Risks and Mitigations

Implementing session management of course has risks since it will be a relatively low level library, used by multiple applications. Mitigations for this should include high test coverage and peer-review from multiple teams within Red Hat and review from VMWare before merging into govmomi. Mitigations are inherent to this enhancement as well since much of the "risky" behavior is already implemented in govmomi including authentication, authorization, session object creation; existing govmomi behavior will be reused in this enhancement.

## Design Details

### Open Questions

- Automatic logout goroutine wait group? I think this is best way to prevent leaks on app shutdown (or cache termination). 
    - It would be very intrusive to unexpectedly logout. Error handling would be impossible for developers.
    - Perhaps an optional on-shutdown hook would be better?
- Should enabling cache be done through `Features`?
- Can existing `SessionManager` or `govmomi.Client` be enhanced instead of adding yet another layer on the clients?

### Test Plan

#### Failure Modes

#### Support Procedures

## Drawbacks

Existing session management in govmomi is good enough and the duplicated implementations seen across projects actually address the unique needs of their project.

## Alternatives

Develop a public best practices document on how to use govmomi sessions to prevent session leaks and best utilize the session resource (simple caching).
