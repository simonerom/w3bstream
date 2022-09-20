package types

import (
	"context"

	"github.com/iotexproject/Bumblebee/conf/log"
	"github.com/iotexproject/Bumblebee/conf/mqtt"
	"github.com/iotexproject/Bumblebee/conf/postgres"
	"github.com/iotexproject/Bumblebee/kit/sqlx"
	"github.com/iotexproject/Bumblebee/x/contextx"
	"github.com/iotexproject/Bumblebee/x/misc/must"
)

type Context uint8

type (
	CtxDBExecutor   struct{} // CtxDBExecutor sqlx.DBExecutor
	CtxPgEndpoint   struct{} // CtxPgEndpoint postgres.Endpoint
	CtxLogger       struct{} // CtxLogger log.Logger
	CtxMqttBroker   struct{} // CtxMqttBroker mqtt.Broker
	CtxUploadConfig struct{} // CtxUploadConfig UploadConfig
)

func WithDBExecutor(ctx context.Context, v sqlx.DBExecutor) context.Context {
	return contextx.WithValue(ctx, CtxDBExecutor{}, v)
}

func WithDBExecutorContext(v sqlx.DBExecutor) contextx.WithContext {
	return func(ctx context.Context) context.Context {
		return contextx.WithValue(ctx, CtxDBExecutor{}, v)
	}
}

func DBExecutorFromContext(ctx context.Context) (sqlx.DBExecutor, bool) {
	v, ok := ctx.Value(CtxDBExecutor{}).(sqlx.DBExecutor)
	return v, ok
}

func MustDBExecutorFromContext(ctx context.Context) sqlx.DBExecutor {
	v, ok := DBExecutorFromContext(ctx)
	must.BeTrue(ok)
	return v
}

func WithPgEndpoint(ctx context.Context, v postgres.Endpoint) context.Context {
	return contextx.WithValue(ctx, CtxPgEndpoint{}, v)
}

func WithPgEndpointContext(v *postgres.Endpoint) contextx.WithContext {
	return func(ctx context.Context) context.Context {
		return contextx.WithValue(ctx, CtxPgEndpoint{}, v)
	}
}

func PgEndpointFromContext(ctx context.Context) (*postgres.Endpoint, bool) {
	v, ok := ctx.Value(CtxPgEndpoint{}).(*postgres.Endpoint)
	return v, ok
}

func MustPgEndpointFromContext(ctx context.Context) *postgres.Endpoint {
	v, ok := PgEndpointFromContext(ctx)
	must.BeTrue(ok)
	return v
}

func WithLogger(ctx context.Context, v log.Logger) context.Context {
	return contextx.WithValue(ctx, CtxLogger{}, v)
}

func WithLoggerContext(v log.Logger) contextx.WithContext {
	return func(ctx context.Context) context.Context {
		return contextx.WithValue(ctx, CtxLogger{}, v)
	}
}

func LoggerFromContext(ctx context.Context) (log.Logger, bool) {
	v, ok := ctx.Value(CtxLogger{}).(log.Logger)
	return v, ok
}

func MustLoggerFromContext(ctx context.Context) log.Logger {
	v, ok := LoggerFromContext(ctx)
	must.BeTrue(ok)
	return v
}

func WithMqttBroker(ctx context.Context, v *mqtt.Broker) context.Context {
	return contextx.WithValue(ctx, CtxMqttBroker{}, v)
}

func WithMqttBrokerContext(v *mqtt.Broker) contextx.WithContext {
	return func(ctx context.Context) context.Context {
		return contextx.WithValue(ctx, CtxMqttBroker{}, v)
	}
}

func MqttBrokerFromContext(ctx context.Context) (*mqtt.Broker, bool) {
	v, ok := ctx.Value(CtxMqttBroker{}).(*mqtt.Broker)
	return v, ok
}

func MustMqttBrokerFromContext(ctx context.Context) *mqtt.Broker {
	v, ok := MqttBrokerFromContext(ctx)
	must.BeTrue(ok)
	return v
}

func WithUploadConfig(ctx context.Context, v *UploadConfig) context.Context {
	return contextx.WithValue(ctx, CtxUploadConfig{}, v)
}

func WithUploadConfigContext(v *UploadConfig) contextx.WithContext {
	return func(ctx context.Context) context.Context {
		return contextx.WithValue(ctx, CtxUploadConfig{}, v)
	}
}

func UploadConfigFromContext(ctx context.Context) (*UploadConfig, bool) {
	v, ok := ctx.Value(CtxUploadConfig{}).(*UploadConfig)
	return v, ok
}

func MustUploadConfigFromContext(ctx context.Context) *UploadConfig {
	v, ok := UploadConfigFromContext(ctx)
	must.BeTrue(ok)
	return v
}