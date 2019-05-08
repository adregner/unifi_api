package unifi_api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/bah2830/unifi_api/pkg/client"
	runtimeClient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Client struct {
	*client.Unifi
	user       string
	pass       string
	basePath   string
	httpClient *http.Client
}

type loginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

func NewClient(host, username, password string) (*Client, error) {
	config := client.DefaultTransportConfig()
	config.Host = host
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 60 * time.Second,
	}

	tranport := runtimeClient.NewWithClient(
		config.Host,
		config.BasePath,
		config.Schemes,
		httpClient,
	)

	c := &Client{
		Unifi:      client.New(tranport, strfmt.Default),
		user:       username,
		pass:       password,
		basePath:   fmt.Sprintf("%s://%s%s", config.Schemes[0], config.Host, config.BasePath),
		httpClient: httpClient,
	}

	if err := c.authenticate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) authenticate() error {
	login := loginParams{
		Username: c.user,
		Password: c.pass,
		Remember: true,
	}

	loginData, err := json.Marshal(login)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.basePath+"login", bytes.NewBuffer(loginData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	spew.Dump(string(respBody))

	return nil
}
