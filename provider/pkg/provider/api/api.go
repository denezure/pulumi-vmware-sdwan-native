package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	vcoUrl    string
	apiToken  string
	portalUrl string
	client    http.Client
}

func NewClient(vcoUrl string, apiToken string, timeout int) Client {
	return Client{
		vcoUrl:    vcoUrl,
		apiToken:  apiToken,
		portalUrl: fmt.Sprintf("https://%s/portal/", vcoUrl),
		client: http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

type rpcRequest struct {
	JsonRpcVer string      `json:"jsonrpc"`
	ID         int         `json:"id"`
	Method     string      `json:"method"`
	Params     interface{} `json:"params"`
}

type rpcResponse struct {
	Result json.RawMessage `json:"result"`
	Error  json.RawMessage `json:"error"`
}

type rpcError struct {
	Code    int64           `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func (c *Client) DoPortal(method string, params interface{}) (json.RawMessage, error) {
	requestBody, err := json.Marshal(rpcRequest{
		JsonRpcVer: "2.0",
		ID:         1,
		Method:     method,
		Params:     params,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.portalUrl, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Token "+c.apiToken)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var rpcResult rpcResponse
	err = json.Unmarshal(resBody, &rpcResult)
	if err != nil {
		return nil, err
	}

	if rpcResult.Result == nil {
		var errInfo rpcError
		json.Unmarshal(rpcResult.Error, &errInfo)
		return nil, fmt.Errorf("API error: [%d] %s [[ %v ]]", errInfo.Code, errInfo.Message, errInfo.Data)
	}

	return rpcResult.Result, nil
}
