```go
type Server struct {
	// Addr: host:port, 如果为空: ":http" (port -> 80)
	Addr string
	// Handler 路由处理器
	Handler Handler

	// DisableGeneralOptionsHandler
	// true: 将OPTIONS请求传递给实际的请求处理程序进行处理
	// false: 返回200, 响应头包含一个Content-Length字段, 其值为0
	DisableGeneralOptionsHandler bool

	// TLSConfig optionally provides a TLS configuration for use
	// by ServeTLS and ListenAndServeTLS. Note that this value is
	// cloned by ServeTLS and ListenAndServeTLS, so it's not
	// possible to modify the configuration with methods like
	// tls.Config.SetSessionTicketKeys. To use
	// SetSessionTicketKeys, use Server.Serve with a TLS Listener
	// instead.
	TLSConfig *tls.Config

	// ReadTimeout 读取整个请求的最长持续时间(包括body), 针对整个请求, 而不是对请求体的个别部分
	// <=0 表示不会超时, 并不允许处理程序在每个请求的请求体上做出针对请求体可接受截止时间或上传速率的单独决定
	// 限制header+body, 不能对每个请求的请求体上做出自定义的决定, 比如设定每个请求的请求体的可接受截止时间(比如某个特定请求的请求体必须在一定时间内完成传输)
	ReadTimeout time.Duration

	// ReadHeaderTimeout 表示读取请求头允许的最长时间, 如果为0, 使用ReadTimeout, 都没有不限制
	// 在读取了请求头之后, 连接的读取截止时间将被重置
	ReadHeaderTimeout time.Duration

	// WriteTimeout 用于配置服务器在响应的写入过程中的超时时限, 会在每个新请求到来时重新启动计时
	// 计时器会在每个新请求的请求头被读取时重新启动, 目的是为了确保服务器能够及时地处理响应, 避免响应写入操作因为某些原因而长时间阻塞
	// <= 0不做限制
	WriteTimeout time.Duration

	// IdleTimeout 用于控制在开启长连接(keep-alives)时等待下一个请求的最大时间
	// 如果为0, 使用ReadTimeout的值作为超时时限, 如果都为0, 表示没有设置超时限制.
	// 这个选项的目的是为了控制空闲连接的超时, 以便及时释放资源
	IdleTimeout time.Duration

	// MaxHeaderBytes 用于控制服务器在解析请求头时所读取的最大字节数, 包括请求行, 不会影响请求体的大小
	// 如果没有单独设置, 使用默认值DefaultMaxHeaderBytes. 目的是为了防止恶意或异常的请求头过大
	// If zero, DefaultMaxHeaderBytes is used.
	MaxHeaderBytes int

	// TLSNextProto optionally specifies a function to take over
	// ownership of the provided TLS connection when an ALPN
	// protocol upgrade has occurred. The map key is the protocol
	// name negotiated. The Handler argument should be used to
	// handle HTTP requests and will initialize the Request's TLS
	// and RemoteAddr if not already set. The connection is
	// automatically closed when the function returns.
	// If TLSNextProto is not nil, HTTP/2 support is not enabled
	// automatically.
	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)

	// ConnState specifies an optional callback function that is
	// called when a client connection changes state. See the
	// ConnState type and associated constants for details.
	ConnState func(net.Conn, ConnState)

	// ErrorLog specifies an optional logger for errors accepting
	// connections, unexpected behavior from handlers, and
	// underlying FileSystem errors.
	// If nil, logging is done via the log package's standard logger.
	ErrorLog *log.Logger

	// BaseContext optionally specifies a function that returns
	// the base context for incoming requests on this server.
	// The provided Listener is the specific Listener that's
	// about to start accepting requests.
	// If BaseContext is nil, the default is context.Background().
	// If non-nil, it must return a non-nil context.
	BaseContext func(net.Listener) context.Context

	// ConnContext optionally specifies a function that modifies
	// the context used for a new connection c. The provided ctx
	// is derived from the base context and has a ServerContextKey
	// value.
	ConnContext func(ctx context.Context, c net.Conn) context.Context
	// true: server关闭
	inShutdown atomic.Bool

	disableKeepAlives atomic.Bool
	nextProtoOnce     sync.Once // guards setupHTTP2_* init
	nextProtoErr      error     // result of http2.ConfigureServer if used

	mu         sync.Mutex
	listeners  map[*net.Listener]struct{}
	activeConn map[*conn]struct{}
	onShutdown []func()

	listenerGroup sync.WaitGroup
}
// ListenAndServe 监听tcp, 然后启动服务处理请求
func (srv *Server) ListenAndServe() error {
	// 利用srv.InShutdown字段判断server是否已经关闭
	if srv.shuttingDown() {
		return ErrServerClosed
	}
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve(ln)
}
// Serve 接受传入的连接, 并未每个连接创建一个新的goroutine来处理请求
func (srv *Server) Serve(l net.Listener) error {
	// 设置http2
	// 追踪所有的连接, 如果服务已经关闭, 不再接受新地连接, 无法运行服务(一个server可能监听不同的端口, 有多个监听器)
	for {
		l.Accept()
		// 失败重试机制
		c := srv.newConn()
		c.setState()
		go c.serve()
	}	
}

// 追踪所有的连接, 如果服务已经关闭, 则无法接受新的连接
// set: 
// srv := &http.Server{}
// go srv.Serve(listener1)
// go srv.Serve(listener2)
func (s *Server) trackListener(ln *net.Listener, add bool) bool {
	if s.listeners == nil {
		s.listeners = make(map[*net.Listener]struct{})
	}
	if add {
		if s.shuttingDown() {
			return false
		}
	}
	
	return true
}

func (c *conn) serve(ctx context.Context) {
	// tls处理
	// 创建一个reader, bufReader, bufWriter
	for {
		w,err := c.readRequest()
		// 标记状态, 处理错误等
		// 调用ServeHTTP处理请求
		serverHandler{c.server}.ServeHTTP(w, w.req)
	}
}
```

```go
// Handler 响应http请求的接口. 将resp的header和body写入ResponseWriter, 此后将不再有效地使用ResponseWriter或者并发地读取Request.Body(确保顺序性, 如果写了后再读可能会出错, 导致请求和响应的数据交错)
// Handler 在 ServeHTTP 调用完成后或同时使用 ResponseWriter 或从 Request.Body 读取是无效的
// 谨慎处理程序读取Request.Body, 除了读取请求体之外, 处理程序不应该修改提供的Request
// 发生panic会recover
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

```

```go
// ServeMux 维护path到pathHandler的映射关系
type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	es    []muxEntry // slice of entries sorted from longest to shortest.
	hosts bool       // whether any patterns contain hostnames
}
// path + handler
type muxEntry struct {
	h       Handler
	pattern string
}
```