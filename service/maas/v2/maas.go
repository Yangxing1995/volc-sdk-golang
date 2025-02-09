package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/maas"
	"github.com/volcengine/volc-sdk-golang/service/maas/models/api/v2"
	"github.com/volcengine/volc-sdk-golang/service/maas/sse"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// MaaS ... use base client
type MaaS struct {
	*base.Client
}

// NewInstance ...
func NewInstance(host, region string) *MaaS {
	instance := &MaaS{}
	instance.Client = base.NewClient(&base.ServiceInfo{
		Timeout: maas.ServiceTimeout,
		Scheme:  "https",
		Host:    host,
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{
			Region:  region,
			Service: maas.ServiceName,
		},
	}, map[string]*base.ApiInfo{
		maas.APIChat: {
			Method: http.MethodPost,
			Path:   "/api/v2/endpoint/%s/chat",
		},
		maas.APIStreamChat: {
			Method: http.MethodPost,
			Path:   "/api/v2/endpoint/%s/chat",
		},
		maas.APICert: {
			Method: http.MethodPost,
			Path:   "/api/v2/endpoint/%s/cert",
		},
		maas.APIClassification: {
			Method: http.MethodPost,
			Path:   "/api/v2/endpoint/%s/classification",
		},
		maas.APITokenization: {
			Method: http.MethodPost,
			Path:   "/api/v2/endpoint/%s/tokenization",
		},
		maas.APIEmbeddings: {
			Method: http.MethodPost,
			Path:   "/api/v2/endpoint/%s/embeddings",
		},
	})

	return instance
}

// POST method
// Chat ...
func (cli *MaaS) Chat(endpointId string, req *api.ChatReq) (*api.ChatResp, int, error) {
	return cli.ChatWithCtx(context.Background(), endpointId, req)
}

// POST method
// ChatWithCtx ...
func (cli *MaaS) ChatWithCtx(ctx context.Context, endpointId string, req *api.ChatReq) (*api.ChatResp, int, error) {
	req.Stream = false

	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()), "")
	}

	return cli.ChatImpl(ctx, endpointId, bts)
}

// POST method
// SecretChat is like `Chat`, except its messages are encrypted
// to ensure that messages are not intercepted by receivers other than the model.
func (cli *MaaS) SecretChat(endpointId string, req *api.ChatReq) (*api.ChatResp, int, error) {
	return cli.SecretChatWithCtx(context.Background(), endpointId, req)
}

// POST method
// SecretChatWithCtx is like `ChatWithCtx`, except its messages are encrypted
// to ensure that messages are not intercepted by receivers other than the model.
func (cli *MaaS) SecretChatWithCtx(ctx context.Context, endpointId string, req *api.ChatReq) (*api.ChatResp, int, error) {
	key, nonce, req, err := cli.encryptChatRequest(ctx, endpointId, req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to encrypt chat request: %v", err), "")
	}

	output, status, err := cli.ChatWithCtx(ctx, endpointId, req)
	if err != nil {
		return nil, status, err
	}

	output, err = cli.decryptChatResponse(key, nonce, output)
	if err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to decrypt chat response: %v", err), "")
	}
	return output, status, nil
}

// POST method
// SecretStreamChat is like `StreamChat`, except its messages are encrypted
// to ensure that messages are not intercepted by receivers other than the model.
func (cli *MaaS) SecretStreamChat(endpointId string, req *api.ChatReq) (ch <-chan *api.ChatResp, err error) {
	return cli.SecretStreamChatWithCtx(context.Background(), endpointId, req)
}

// POST method
// SecretStreamChatWithCtx is like `StreamChatWithCtx`, except its messages are encrypted
// to ensure that messages are not intercepted by receivers other than the model.
func (cli *MaaS) SecretStreamChatWithCtx(ctx context.Context, endpointId string, req *api.ChatReq) (ch <-chan *api.ChatResp, err error) {
	key, nonce, req, err := cli.encryptChatRequest(ctx, endpointId, req)
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to encrypt chat request: %v", err), "")
	}

	resps, err := cli.StreamChatWithCtx(ctx, endpointId, req)
	if err != nil {
		return nil, err
	}

	outputs := make(chan *api.ChatResp, maas.RespBufferSize)
	go func() {
		defer func() {
			_ = recover()
			close(outputs)
		}()

		for resp := range resps {
			output, err := cli.decryptChatResponse(key, nonce, resp)
			if err != nil {
				resp.Error = api.NewClientSDKRequestError(fmt.Sprintf("failed to decrypt chat response: %v", err), "")
				outputs <- resp
				continue
			}
			outputs <- output
		}
	}()
	return outputs, nil
}

// POST method
// StreamChat make stream chat request
//  1. if any error returned, a channel=`nil` is returned;
//  2. if no error returned, the channel are closed after all responses processed.
func (cli *MaaS) StreamChat(endpointId string, req *api.ChatReq) (ch <-chan *api.ChatResp, err error) {
	return cli.StreamChatWithCtx(context.Background(), endpointId, req)
}

// POST method
// StreamChat make stream chat request
//  1. if any error returned, a channel=`nil` is returned;
//  2. if no error returned, the channel are closed after all responses processed.
func (cli *MaaS) StreamChatWithCtx(ctx context.Context, endpointId string, req *api.ChatReq) (ch <-chan *api.ChatResp, err error) {
	req.Stream = true

	bts, err := json.Marshal(req)
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()), "")
	}

	return cli.StreamChatImpl(ctx, endpointId, bts)
}

func (cli *MaaS) ChatImpl(ctx context.Context, endpointId string, body []byte) (*api.ChatResp, int, error) {
	respBody, status, err, logId := cli.request(ctx, maas.APIChat, nil, endpointId, body)
	if err != nil {
		errVal := &api.ChatResp{}
		if er := json.Unmarshal(respBody, errVal); er != nil {
			errVal.Error = api.NewClientSDKRequestError(err.Error(), logId)
		}
		errVal.Error.ReqId = errVal.ReqId
		return errVal, status, errVal.Error
	}

	output := new(api.ChatResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()), logId)
	}
	return output, status, nil
}

func (cli *MaaS) StreamChatImpl(ctx context.Context, endpointId string, body []byte) (<-chan *api.ChatResp, error) {
	apiInfo := cli.ApiInfoList[maas.APIStreamChat]
	if apiInfo == nil {
		return nil, api.NewClientSDKRequestError("the related api does not exist", "")
	}

	// build request
	req, err := maas.MakeRequest(apiInfo, endpointId, cli.ServiceInfo, nil, "application/json")
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to make request: %v", err), "")
	}
	req.Body = io.NopCloser(bytes.NewReader(body))
	timeout := maas.GetTimeout(cli.ServiceInfo.Timeout, apiInfo.Timeout)

	req = cli.ServiceInfo.Credentials.Sign(req)

	ctx, cancel := context.WithTimeout(ctx, timeout)
	req = req.WithContext(ctx)

	// do request
	resp, err := cli.Client.Client.Do(req)
	logId := resp.Header.Get("x-tt-logid")
	if err != nil {
		cancel()
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("request error: %v", err), logId)
	}

	if resp.StatusCode != 200 { // fast fail
		res := &api.ChatResp{}
		if er := json.NewDecoder(resp.Body).Decode(res); er != nil {
			res.Error = api.NewClientSDKRequestError(fmt.Sprintf("failed to call service: http status_code=%d", resp.StatusCode), logId)
		}
		cancel()
		_ = resp.Body.Close()
		return nil, res.Error
	}

	// parse response
	ch := make(chan *api.ChatResp, maas.RespBufferSize)
	go func() {
		defer func() {
			_ = recover()
			_ = resp.Body.Close()
			cancel()
			close(ch)
		}()

		stream := sse.NewEventStreamFromReader(resp.Body, maas.MaxBufferSize)
		for {
			event, err := stream.Next()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				if errors.Is(err, context.DeadlineExceeded) {
					ch <- &api.ChatResp{
						Error: api.NewClientSDKRequestError(fmt.Sprintf("call service timeout: timeout=%s", timeout.String()), logId),
					}
				} else {
					ch <- &api.ChatResp{
						Error: api.NewClientSDKRequestError(err.Error(), logId),
					}
				}
				return
			}
			if event != nil {
				if bytes.Equal(event.Data, []byte(maas.Terminator)) {
					return
				}

				item := &api.ChatResp{}
				if err = json.Unmarshal(event.Data, item); err != nil {
					ch <- &api.ChatResp{
						Error: api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response(data=%s): %v", string(event.Data), err), logId),
					}
					return
				}
				item.ReqId = logId
				if item.Error != nil {
					item.Error.ReqId = logId
				}
				ch <- item
			}
		}
	}()

	return ch, nil
}

func (cli *MaaS) initCertByReq(ctx context.Context, endpointId string, req *api.ChatReq) (*maas.KeyAgreementClient, error) {
	certReq := &api.CertReq{}
	body, err := json.Marshal(certReq)
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()), "")
	}
	respBody, _, err, logId := cli.request(ctx, maas.APICert, nil, endpointId, body)
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to get CA from proxy: %s", err.Error()), logId)
	}
	output := new(api.CertResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()), logId)
	}

	// todo: check chain
	return maas.NewP256KeyAgreementClient(output.Cert)
}

func (cli *MaaS) encryptChatRequest(ctx context.Context, endpointId string, req *api.ChatReq) ([]byte, []byte, *api.ChatReq, error) {
	ka, err := cli.initCertByReq(ctx, endpointId, req)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to init cert: %w", err)
	}
	key, nonce, token, err := ka.GenerateECIESKeyPair()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to generate session token: %w", err)
	}

	req.CryptoToken = token

	for i := range req.Messages {
		content, ok := req.Messages[i].Content.(string)
		if ok && content != "" {
			secret, err := maas.AesGcmEncryptBase64String(key, nonce, content)
			if err != nil {
				return nil, nil, nil, fmt.Errorf("failed to encrypt message: %w", err)
			}
			req.Messages[i].Content = secret
		}
	}

	return key, nonce, req, nil
}

func (cli *MaaS) decryptChatResponse(key, nonce []byte, resp *api.ChatResp) (*api.ChatResp, error) {
	if len(resp.Choices) == 0 {
		return resp, nil
	}
	if secret, ok := resp.Choices[0].Message.Content.(string); ok && secret != "" {
		plain, err := maas.AesGcmDecryptBase64String(key, nonce, secret)
		if err != nil {
			return nil, err
		}
		resp.Choices[0].Message.Content = plain
	}
	return resp, nil
}

// POST method
// Tokenization
func (cli *MaaS) Tokenization(endpointId string, req *api.TokenizeReq) (*api.TokenizeResp, int, error) {
	return cli.TokenizationWithCtx(context.Background(), endpointId, req)
}

func (cli *MaaS) TokenizationWithCtx(ctx context.Context, endpointId string, req *api.TokenizeReq) (*api.TokenizeResp, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()), "")
	}
	return cli.tokenizationImpl(ctx, endpointId, bts)
}

func (cli *MaaS) tokenizationImpl(ctx context.Context, endpointId string, body []byte) (*api.TokenizeResp, int, error) {
	respBody, status, err, logId := cli.request(ctx, maas.APITokenization, nil, endpointId, body)
	if err != nil {
		errVal := &api.TokenizeResp{}
		if er := json.Unmarshal(respBody, errVal); er != nil {
			errVal.Error = api.NewClientSDKRequestError(err.Error(), logId)
		}
		errVal.Error.ReqId = errVal.ReqId
		return nil, status, errVal.Error
	}

	output := new(api.TokenizeResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()), logId)
	}
	return output, status, nil
}

// POST method
// Classification
func (cli *MaaS) Classification(endpointId string, req *api.ClassificationReq) (*api.ClassificationResp, int, error) {
	return cli.ClassificationWithCtx(context.Background(), endpointId, req)
}

func (cli *MaaS) ClassificationWithCtx(ctx context.Context, endpointId string, req *api.ClassificationReq) (*api.ClassificationResp, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()), "")
	}
	return cli.classificationImpl(ctx, endpointId, bts)
}

func (cli *MaaS) classificationImpl(ctx context.Context, endpointId string, body []byte) (*api.ClassificationResp, int, error) {
	respBody, status, err, logId := cli.request(ctx, maas.APIClassification, nil, endpointId, body)
	if err != nil {
		errVal := &api.ClassificationResp{}
		if er := json.Unmarshal(respBody, errVal); er != nil {
			errVal.Error = api.NewClientSDKRequestError(err.Error(), logId)
		}
		errVal.Error.ReqId = errVal.ReqId
		return nil, status, errVal.Error
	}

	output := new(api.ClassificationResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()), logId)
	}
	return output, status, nil
}

// POST method
// Embeddings
func (cli *MaaS) Embeddings(endpointId string, req *api.EmbeddingsReq) (*api.EmbeddingsResp, int, error) {
	return cli.EmbeddingsWithCtx(context.Background(), endpointId, req)
}

func (cli *MaaS) EmbeddingsWithCtx(ctx context.Context, endpointId string, req *api.EmbeddingsReq) (*api.EmbeddingsResp, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()), "")
	}
	return cli.embeddingsImpl(ctx, endpointId, bts)
}

func (cli *MaaS) embeddingsImpl(ctx context.Context, endpointId string, body []byte) (*api.EmbeddingsResp, int, error) {
	respBody, status, err, logId := cli.request(ctx, maas.APIEmbeddings, nil, endpointId, body)
	if err != nil {
		errVal := &api.EmbeddingsResp{}
		if er := json.Unmarshal(respBody, errVal); er != nil {
			errVal.Error = api.NewClientSDKRequestError(err.Error(), logId)
		}
		errVal.Error.ReqId = errVal.ReqId
		return nil, status, errVal.Error
	}

	output := new(api.EmbeddingsResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()), logId)
	}
	return output, status, nil
}

func (cli *MaaS) doRequest(inputContext context.Context, api string, req *http.Request, timeout time.Duration) ([]byte, int, error, bool, string) {
	req = cli.ServiceInfo.Credentials.Sign(req)

	ctx := inputContext
	if ctx == nil {
		ctx = context.Background()
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	req = req.WithContext(ctx)

	resp, err := cli.Client.Client.Do(req)
	if err != nil {
		// should retry when client sends request error.
		return []byte(""), 500, err, true, ""
	}
	defer resp.Body.Close()

	logId := resp.Header.Get("x-tt-logid")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), resp.StatusCode, err, false, logId
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		needRetry := false
		// should retry when server returns 5xx error.
		if resp.StatusCode >= http.StatusInternalServerError {
			needRetry = true
		}
		return body, resp.StatusCode, fmt.Errorf("api %s http code %d body %s", api, resp.StatusCode, string(body)), needRetry, logId
	}

	return body, resp.StatusCode, nil, false, logId
}

func (cli *MaaS) request(ctx context.Context, apiKey string, query url.Values, endpointId string, body []byte) ([]byte, int, error, string) {
	apiInfo := cli.ApiInfoList[apiKey]
	if apiInfo == nil {
		return []byte(""), 500, errors.New("The related api does not exist"), ""
	}

	// build request
	req, err := maas.MakeRequest(apiInfo, endpointId, cli.ServiceInfo, query, "application/json")
	if err != nil {
		return nil, 500, api.NewClientSDKRequestError(fmt.Sprintf("failed to make request: %v", err), ""), ""
	}
	requestBody := bytes.NewReader(body)
	timeout := maas.GetTimeout(cli.ServiceInfo.Timeout, apiInfo.Timeout)
	retrySettings := maas.GetRetrySetting(&cli.ServiceInfo.Retry, &apiInfo.Retry)

	var resp []byte
	var code int
	var logId string

	err = backoff.Retry(func() error {
		_, err = requestBody.Seek(0, io.SeekStart)
		if err != nil {
			// if seek failed, stop retry.
			return backoff.Permanent(err)
		}
		req.Body = ioutil.NopCloser(requestBody)
		var needRetry bool
		resp, code, err, needRetry, logId = cli.doRequest(ctx, apiKey, req, timeout)
		if needRetry {
			return err
		} else {
			return backoff.Permanent(err)
		}
	}, backoff.WithMaxRetries(backoff.NewConstantBackOff(*retrySettings.RetryInterval), *retrySettings.RetryTimes))
	return resp, code, err, logId
}
