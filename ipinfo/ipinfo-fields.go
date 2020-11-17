// Code generated by gen-fields; DO NOT EDIT.

package ipinfo // import "github.com/ipinfo/go/ipinfo"

import (
	"bytes"
	"net"
)

// GetHostname returns a specific field "hostname" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetHostname(ip net.IP) (string, error) {
	return c.GetHostname(ip)
}

// GetHostname returns a specific field "hostname" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetHostname(ip net.IP) (string, error) {
	s := "hostname"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	if c.Cache == nil {
		return c.requestHostname(s)
	}
	v, err := c.Cache.GetOrRequest(s, func() (interface{}, error) {
		return c.requestHostname(s)
	})
	if err != nil {
		return "", err
	}
	return v.(string), err
}

func (c *Client) requestHostname(s string) (string, error) {
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// GetOrganization returns a specific field "org" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetOrganization(ip net.IP) (string, error) {
	return c.GetOrganization(ip)
}

// GetOrganization returns a specific field "org" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetOrganization(ip net.IP) (string, error) {
	s := "org"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	if c.Cache == nil {
		return c.requestOrganization(s)
	}
	v, err := c.Cache.GetOrRequest(s, func() (interface{}, error) {
		return c.requestOrganization(s)
	})
	if err != nil {
		return "", err
	}
	return v.(string), err
}

func (c *Client) requestOrganization(s string) (string, error) {
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// GetCity returns a specific field "city" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetCity(ip net.IP) (string, error) {
	return c.GetCity(ip)
}

// GetCity returns a specific field "city" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetCity(ip net.IP) (string, error) {
	s := "city"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	if c.Cache == nil {
		return c.requestCity(s)
	}
	v, err := c.Cache.GetOrRequest(s, func() (interface{}, error) {
		return c.requestCity(s)
	})
	if err != nil {
		return "", err
	}
	return v.(string), err
}

func (c *Client) requestCity(s string) (string, error) {
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// GetRegion returns a specific field "region" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetRegion(ip net.IP) (string, error) {
	return c.GetRegion(ip)
}

// GetRegion returns a specific field "region" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetRegion(ip net.IP) (string, error) {
	s := "region"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	if c.Cache == nil {
		return c.requestRegion(s)
	}
	v, err := c.Cache.GetOrRequest(s, func() (interface{}, error) {
		return c.requestRegion(s)
	})
	if err != nil {
		return "", err
	}
	return v.(string), err
}

func (c *Client) requestRegion(s string) (string, error) {
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// GetCountry returns a specific field "country" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetCountry(ip net.IP) (string, error) {
	return c.GetCountry(ip)
}

// GetCountry returns a specific field "country" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetCountry(ip net.IP) (string, error) {
	s := "country"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	if c.Cache == nil {
		return c.requestCountry(s)
	}
	v, err := c.Cache.GetOrRequest(s, func() (interface{}, error) {
		return c.requestCountry(s)
	})
	if err != nil {
		return "", err
	}
	return v.(string), err
}

func (c *Client) requestCountry(s string) (string, error) {
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// GetLocation returns a specific field "loc" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetLocation(ip net.IP) (string, error) {
	return c.GetLocation(ip)
}

// GetLocation returns a specific field "loc" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetLocation(ip net.IP) (string, error) {
	s := "loc"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	if c.Cache == nil {
		return c.requestLocation(s)
	}
	v, err := c.Cache.GetOrRequest(s, func() (interface{}, error) {
		return c.requestLocation(s)
	})
	if err != nil {
		return "", err
	}
	return v.(string), err
}

func (c *Client) requestLocation(s string) (string, error) {
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// GetPhone returns a specific field "phone" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetPhone(ip net.IP) (string, error) {
	return c.GetPhone(ip)
}

// GetPhone returns a specific field "phone" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetPhone(ip net.IP) (string, error) {
	s := "phone"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	if c.Cache == nil {
		return c.requestPhone(s)
	}
	v, err := c.Cache.GetOrRequest(s, func() (interface{}, error) {
		return c.requestPhone(s)
	})
	if err != nil {
		return "", err
	}
	return v.(string), err
}

func (c *Client) requestPhone(s string) (string, error) {
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// GetPostal returns a specific field "postal" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetPostal(ip net.IP) (string, error) {
	return c.GetPostal(ip)
}

// GetPostal returns a specific field "postal" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetPostal(ip net.IP) (string, error) {
	s := "postal"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	if c.Cache == nil {
		return c.requestPostal(s)
	}
	v, err := c.Cache.GetOrRequest(s, func() (interface{}, error) {
		return c.requestPostal(s)
	})
	if err != nil {
		return "", err
	}
	return v.(string), err
}

func (c *Client) requestPostal(s string) (string, error) {
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// GetTimezone returns a specific field "timezone" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetTimezone(ip net.IP) (string, error) {
	return c.GetTimezone(ip)
}

// GetTimezone returns a specific field "timezone" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetTimezone(ip net.IP) (string, error) {
	s := "timezone"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	if c.Cache == nil {
		return c.requestTimezone(s)
	}
	v, err := c.Cache.GetOrRequest(s, func() (interface{}, error) {
		return c.requestTimezone(s)
	})
	if err != nil {
		return "", err
	}
	return v.(string), err
}

func (c *Client) requestTimezone(s string) (string, error) {
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}
