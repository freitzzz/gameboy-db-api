package http

import "github.com/labstack/echo/v4"

type httpServer struct {
	host    string
	port    string
	cert    *string
	certKey *string

	echo *echo.Echo
}

type httpServerBuilder struct {
	// conn
	host    string
	port    string
	cert    *string
	certKey *string

	// di
	serviceContainer *serviceContainer

	// routing
	virtualPath *string
}

func New() *httpServerBuilder {
	return &httpServerBuilder{}
}

func (b *httpServerBuilder) WithHostPort(host string, port string) *httpServerBuilder {
	b.host = host
	b.port = port
	return b
}

func (b *httpServerBuilder) WithSelfSignedCertificate(cert, certKey string) *httpServerBuilder {
	b.cert = &cert
	b.certKey = &certKey
	return b
}

func (b *httpServerBuilder) WithServiceContainer(sc serviceContainer) *httpServerBuilder {
	b.serviceContainer = &sc
	return b
}

func (b *httpServerBuilder) WithVirtualPath(virtualPath string) *httpServerBuilder {
	b.virtualPath = &virtualPath
	return b
}

func (b *httpServerBuilder) Build() *httpServer {
	e := echo.New()

	if b.virtualPath != nil {
		setVirtualPath(virtualpath)
	}

	registerHandlers(e)
	registerMiddlewares(e, *b.serviceContainer)

	return &httpServer{
		host:    b.host,
		port:    b.port,
		cert:    b.cert,
		certKey: b.certKey,
		echo:    e,
	}
}
