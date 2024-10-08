// Code generated by protoc-gen-twirplura v1.0.0, DO NOT EDIT.
// source: protos/color-service/service.proto

package color

import context "context"
import json "encoding/json"
import fmt "fmt"

import "github.com/kyawmyintthein/lura-twirp"
import "github.com/luraproject/lura/config"
import "github.com/luraproject/lura/logging"
import twirp "github.com/twitchtv/twirp"
import proto "google.golang.org/protobuf/proto"

// Version compatibility assertion.
// If the constant is not defined in the package, that likely means
// the package needs to be updated to work with this generated code.
// See https://twitchtv.github.io/twirp/docs/version_matrix.html
const _ = twirp.TwirpPackageMinVersion_8_1_0

// ========================
// ColorService Lura Client
// ========================

type colorServiceLuraClient struct {
	id      string
	service ColorService
	l       logging.Logger
}

// ====================
// ColorService Methods
// ====================

const (
	_ColorServiceMethod_GetRandomColor = "GetRandomColor"
	_ColorServiceMethod_GetRGBColor    = "GetRGBColor"
)

// ===============================================================================================
// NewColorServiceLuraClient creates a Protobuf client that implements the ColorService interface.
// ===============================================================================================

func NewColorServiceLuraClient(config *config.ServiceConfig, id string, client HTTPClient, l logging.Logger, opts ...twirp.ClientOption) (luratwirp.LuraTwirpStub, error) {
	baseURL, err := getBaseURLByColorServiceClientID(config, id)
	if err != nil {
		return nil, err
	}
	protobufClient := NewColorServiceProtobufClient(baseURL, client, opts...)
	return &colorServiceLuraClient{
		id:      id,
		service: protobufClient,
		l:       l,
	}, nil
}

// =================================
// ColorService getBaseURLByClientID
// =================================

func getBaseURLByColorServiceClientID(config *config.ServiceConfig, id string) (string, error) {
	for _, endpoint := range config.Endpoints {
		val, ok := endpoint.ExtraConfig[luratwirp.TwirpServiceIdentifierConst].(string)
		if ok {
			if val != id {
				continue
			}
			for _, backend := range endpoint.Backend {
				_, ok := backend.ExtraConfig[luratwirp.TwirpServiceIdentifierConst].(string)
				if ok {
					if val != id {
						continue
					}
					if len(backend.Host) <= 0 {
						return "", twirp.InternalError("invalid host configuration")
					}
				}
				return backend.Host[0], nil
			}
		}
	}
	return "", twirp.InternalError(fmt.Sprintf("invalid %s", luratwirp.TwirpServiceIdentifierConst))
}

// ==============================================================
// Invoke invoke RPC function regarding given service and method.
// ==============================================================

func (c *colorServiceLuraClient) Invoke(ctx context.Context, service string, method string, in proto.Message) (proto.Message, error) {
	switch method {
	case _ColorServiceMethod_GetRandomColor:
		req, ok := in.(*GetRandomColorRequest)
		if !ok {
			return nil, twirp.InternalError("invalid protobuf message")
		}
		resp, err := c.service.GetRandomColor(ctx, req)
		if err != nil {
			c.l.Error(err, "failed to invoke : ", _ColorServiceMethod_GetRandomColor)
			return resp, err
		}
		return resp, err
	case _ColorServiceMethod_GetRGBColor:
		req, ok := in.(*GetRGBColorRequest)
		if !ok {
			return nil, twirp.InternalError("invalid protobuf message")
		}
		resp, err := c.service.GetRGBColor(ctx, req)
		if err != nil {
			c.l.Error(err, "failed to invoke : ", _ColorServiceMethod_GetRGBColor)
			return resp, err
		}
		return resp, err
	}
	return nil, twirp.InternalError(fmt.Sprintf("invalid %s", luratwirp.TwirpServiceIdentifierConst))
}

// ===================================================================
// Identifier return client identifier to lura-twirp backend registery
// ===================================================================

func (c *colorServiceLuraClient) Identifier() string {
	return c.id
}

// ====================================
// Encode convert JSON to proto.Message
// ====================================

func (c *colorServiceLuraClient) Encode(ctx context.Context, method string, data []byte) (proto.Message, error) {
	switch method {
	case _ColorServiceMethod_GetRandomColor:
		out := new(GetRandomColorRequest)
		err := json.Unmarshal(data, out)
		if err != nil {
			c.l.Error(err, "failed to unmarhsal : ", _ColorServiceMethod_GetRandomColor)
			return out, err
		}
		return out, err
	case _ColorServiceMethod_GetRGBColor:
		out := new(GetRGBColorRequest)
		err := json.Unmarshal(data, out)
		if err != nil {
			c.l.Error(err, "failed to unmarhsal : ", _ColorServiceMethod_GetRGBColor)
			return out, err
		}
		return out, err
	}
	return nil, twirp.InternalError(fmt.Sprintf("invalid method %s", method))
}

// ====================================
// Decode convert proto.Message to JSON
// ====================================

func (c *colorServiceLuraClient) Decode(ctx context.Context, method string, msg proto.Message) ([]byte, error) {
	switch method {
	case _ColorServiceMethod_GetRandomColor:
		out, err := proto.Marshal(msg)
		if err != nil {
			c.l.Error(err, "failed to marshal : ", _ColorServiceMethod_GetRandomColor)
			return out, err
		}
		return out, err
	case _ColorServiceMethod_GetRGBColor:
		out, err := proto.Marshal(msg)
		if err != nil {
			c.l.Error(err, "failed to marshal : ", _ColorServiceMethod_GetRGBColor)
			return out, err
		}
		return out, err
	}
	return nil, twirp.InternalError(fmt.Sprintf("invalid method %s", method))
}
