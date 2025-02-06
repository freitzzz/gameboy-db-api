package http

import "github.com/labstack/echo/v4"

const (
	defaultServerHost = "::1"
	defaultServerPort = "8080"
)

type httpServer struct {
	host    string
	port    string
	cert    *string
	certKey *string

	echo *echo.Echo
}

type httpServerBuilder struct {
	// conn
	host    *string
	port    *string
	cert    *string
	certKey *string

	// di
	serviceContainer serviceContainer

	// routing
	virtualPath *string
}

func Builder() *httpServerBuilder {
	return &httpServerBuilder{}
}

func (b *httpServerBuilder) WithHostPort(host, port *string) *httpServerBuilder {
	b.host = host
	b.port = port
	return b
}

func (b *httpServerBuilder) WithSelfSignedCertificate(cert, certKey *string) *httpServerBuilder {
	b.cert = cert
	b.certKey = certKey
	return b
}

func (b *httpServerBuilder) WithServiceContainer(sc serviceContainer) *httpServerBuilder {
	b.serviceContainer = sc
	return b
}

func (b *httpServerBuilder) WithVirtualPath(virtualPath *string) *httpServerBuilder {
	b.virtualPath = virtualPath
	return b
}

func (b *httpServerBuilder) Build() *httpServer {
	e := echo.New()

	var ternary = func(x *string, y string) string {
		if x != nil {
			return *x
		} else {
			return y
		}
	}

	if b.virtualPath != nil {
		setVirtualPath(virtualpath)
	}

	registerHandlers(e)
	registerMiddlewares(e, b.serviceContainer)

	return &httpServer{
		host:    ternary(b.host, defaultServerHost),
		port:    ternary(b.port, defaultServerPort),
		cert:    b.cert,
		certKey: b.certKey,
		echo:    e,
	}
}
