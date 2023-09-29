package client

import (
	"bytes"
	"context"
	"errors"
	v1alpha2 "github.com/tektoncd/results/proto/v1alpha2/results_go_proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"k8s.io/client-go/transport"
	"net/http"
	"net/url"
)

const (
	basePath = "/apis/results.tekton.dev/v1alpha2/parents"
)

type restClient struct {
	httpClient *http.Client
	url        *url.URL
}

// NewRESTClient creates a new REST client.
func NewRESTClient(o *Options) (Client, error) {
	u, err := url.Parse(o.Host)
	if err != nil {
		return nil, err
	}

	u.Path = basePath

	rt, err := transport.New(&transport.Config{
		BearerToken: o.Token,
		TLS:         *o.TLSConfig,
		Impersonate: *o.ImpersonationConfig,
	})
	if err != nil {
		return nil, err
	}

	rc := &restClient{
		httpClient: &http.Client{
			Transport: rt,
			Timeout:   o.Timeout,
		},
		url: u,
	}

	return rc, nil
}

// TODO: Get these methods from a generated client

// GetResult makes request to get result
func (c *restClient) GetResult(ctx context.Context, in *v1alpha2.GetResultRequest, _ ...grpc.CallOption) (*v1alpha2.Result, error) {
	out := &v1alpha2.Result{}
	return out, c.send(ctx, http.MethodGet, []string{in.Name}, in, out)
}

// ListResults makes request and get result list
func (c *restClient) ListResults(ctx context.Context, in *v1alpha2.ListResultsRequest, _ ...grpc.CallOption) (*v1alpha2.ListResultsResponse, error) {
	out := &v1alpha2.ListResultsResponse{}
	return out, c.send(ctx, http.MethodGet, []string{in.Parent, "results"}, in, out)
}

// DeleteResult makes request to delete result
func (c *restClient) DeleteResult(ctx context.Context, in *v1alpha2.DeleteResultRequest, _ ...grpc.CallOption) (*emptypb.Empty, error) {
	out := &emptypb.Empty{}
	return &emptypb.Empty{}, c.send(ctx, http.MethodDelete, []string{in.Name}, in, out)
}

func (c *restClient) CreateResult(_ context.Context, _ *v1alpha2.CreateResultRequest, _ ...grpc.CallOption) (*v1alpha2.Result, error) {
	//TODO implement me
	panic("not implemented")
}

func (c *restClient) UpdateResult(_ context.Context, _ *v1alpha2.UpdateResultRequest, _ ...grpc.CallOption) (*v1alpha2.Result, error) {
	//TODO: implement
	panic("not implemented")
}

// GetRecord makes request to get record
func (c *restClient) GetRecord(ctx context.Context, in *v1alpha2.GetRecordRequest, _ ...grpc.CallOption) (*v1alpha2.Record, error) {
	out := &v1alpha2.Record{}
	return out, c.send(ctx, http.MethodGet, []string{in.Name}, in, out)
}

// ListRecords makes request to get record list
func (c *restClient) ListRecords(ctx context.Context, in *v1alpha2.ListRecordsRequest, _ ...grpc.CallOption) (*v1alpha2.ListRecordsResponse, error) {
	out := &v1alpha2.ListRecordsResponse{}
	return out, c.send(ctx, http.MethodGet, []string{in.Parent, "records"}, in, out)
}

// DeleteRecord makes request to delete record
func (c *restClient) DeleteRecord(ctx context.Context, in *v1alpha2.DeleteRecordRequest, _ ...grpc.CallOption) (*emptypb.Empty, error) {
	out := &emptypb.Empty{}
	return &emptypb.Empty{}, c.send(ctx, http.MethodDelete, []string{in.Name}, in, out)
}

func (c *restClient) CreateRecord(_ context.Context, _ *v1alpha2.CreateRecordRequest, _ ...grpc.CallOption) (*v1alpha2.Record, error) {
	//TODO: implement
	panic("not implemented")
}

func (c *restClient) UpdateRecord(_ context.Context, _ *v1alpha2.UpdateRecordRequest, _ ...grpc.CallOption) (*v1alpha2.Record, error) {
	//TODO: implement
	panic("not implemented")
}

func (c *restClient) GetLog(ctx context.Context, in *v1alpha2.GetLogRequest, _ ...grpc.CallOption) (v1alpha2.Logs_GetLogClient, error) {
	out := &v1alpha2.Log{}
	return nil, c.send(ctx, http.MethodGet, []string{in.Name}, in, out)
}

func (c *restClient) ListLogs(ctx context.Context, in *v1alpha2.ListRecordsRequest, _ ...grpc.CallOption) (*v1alpha2.ListRecordsResponse, error) {
	out := &v1alpha2.ListRecordsResponse{}
	return out, c.send(ctx, http.MethodGet, []string{in.Parent, "records"}, in, out)
}

func (c *restClient) DeleteLog(ctx context.Context, in *v1alpha2.DeleteLogRequest, _ ...grpc.CallOption) (*emptypb.Empty, error) {
	out := &emptypb.Empty{}
	return &emptypb.Empty{}, c.send(ctx, http.MethodDelete, []string{in.Name}, in, out)
}

func (c *restClient) UpdateLog(_ context.Context, _ ...grpc.CallOption) (v1alpha2.Logs_UpdateLogClient, error) {
	panic("not implemented")
}

func (c *restClient) send(ctx context.Context, method string, values []string, in, out proto.Message) error {
	u := c.url.JoinPath(values...)
	q := u.Query()
	in.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.JSONName() == "parent" || !fd.HasJSONName() || fd.Kind() == protoreflect.BytesKind {
			return true
		}
		q.Set(fd.JSONName(), v.String())
		return true
	})
	u.RawQuery = q.Encode()

	var body io.Reader
	if method == http.MethodPost || method == http.MethodPut {
		b, err := protojson.Marshal(in)
		if err != nil {
			return err
		}
		body = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New(http.StatusText(res.StatusCode))
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return protojson.Unmarshal(b, out)
}
