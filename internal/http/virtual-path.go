package http

import (
	"fmt"
	"strings"
)

var virtualpath string

func setVirtualPath(host string) {
	virtualpath = host
}

func withVirtualPath(path string) string {
	if !strings.HasPrefix(virtualpath, "/") {
		virtualpath = "/" + virtualpath
	}

	if !strings.HasSuffix(virtualpath, "/") {
		virtualpath = virtualpath + "/"
	}

	if path == "" || path == "/" {
		return virtualpath
	}

	if virtualpath == "/" {
		return path
	}

	if path[0] == '/' {
		return fmt.Sprintf("%s%s", virtualpath, path[1:])
	}

	return fmt.Sprintf("%s%s", virtualpath, path)
}
