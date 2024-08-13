package main

import (
	"fmt"
	"strings"
)

// makes building a complex object easier by providing an abstraction over
// object construction
//  Instead of directly constructing an object with a large number of constructor
//arguments, the pattern splits the construction process into multiple smaller steps

type Url struct {
	protocol string
	username string
	password string
	hostname string
	port     string
	path     string
}

func (u *Url) Validate() error {
	if u.protocol == "" || u.hostname == "" {
		return fmt.Errorf("must specify at least a protocol and a hostname")
	}
	return nil
}

func (u *Url) String() string {
	var url strings.Builder
	url.WriteString(fmt.Sprintf("%s://", u.protocol))
	if u.username != "" && u.password != "" {
		url.WriteString(fmt.Sprintf("%s:%s@", u.username, u.password))
	}
	url.WriteString(u.hostname)
	if u.port != "" {
		url.WriteString(fmt.Sprintf(":%s", u.port))
	}
	if u.path != "" {
		url.WriteString(u.path)
	}
	return url.String()
}

type UrlBuilder struct {
	protocol string
	username string
	password string
	hostname string
	port     string
	path     string
}

func (b *UrlBuilder) SetProtocol(protocol string) *UrlBuilder {
	b.protocol = protocol
	return b
}

func (b *UrlBuilder) SetAuthentication(Uame, pass string) *UrlBuilder {
	b.username = Uame
	b.password = pass
	return b
}

func (b *UrlBuilder) SetHostName(hostname string) *UrlBuilder {
	b.hostname = hostname
	return b
}

func (b *UrlBuilder) SetPort(port string) *UrlBuilder {
	b.port = port
	return b
}

func (b *UrlBuilder) SetPath(path string) *UrlBuilder {
	b.path = path
	return b
}

func (b *UrlBuilder) Build() *Url {
	return &Url{
		protocol: b.protocol,
		username: b.username,
		password: b.password,
		hostname: b.hostname,
		port:     b.port,
		path:     b.path,
	}
}

// usecase:B ;using interface
type IUrlBuilder interface {
	SetProtocol(protocol string) *UrlBuilder
	SetAuthentication(Uame, pass string) *UrlBuilder
	SetHostName(hostname string) *UrlBuilder
	SetPort(port string) *UrlBuilder
	SetPath(path string) *UrlBuilder
}

func NewIUrlBuilder() IUrlBuilder {
	return &UrlBuilder{}
}

func main() {
	//usecase:A ;using method struct
	urlbuilder := &UrlBuilder{}
	urlbuilder.SetProtocol("http").SetAuthentication("username", "password").SetHostName("host").SetPort("port").SetPath("path")
	url := urlbuilder.Build()
	fmt.Println(url.String())

	//usecase:B ;using interface
	ub := NewIUrlBuilder()
	l := ub.SetAuthentication("", "").SetHostName("").Build()
	fmt.Println(l.String())
	l.Validate()
}
