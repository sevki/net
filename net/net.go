package net

import (
	"context"
	"time"

	"github.com/domainr/whois"
	"sevki.org/net/transport/v1"
)

func New() v1.Net {
	return &netMonkey{whois.NewClient(time.Second * 15)}
}

type netMonkey struct{ whois *whois.Client }

func (m *netMonkey) Whois(ctx context.Context, name string) (*v1.Domain, error) {
	req, err := whois.NewRequest(name)
	if err != nil {
		return nil, err
	}
	resp, err := m.whois.Fetch(req)
	if err != nil {
		return nil, err
	}

	return &v1.Domain{
		Name:  name,
		Whois: resp.String(),
	}, nil
}
