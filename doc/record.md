# 外部认证

## 基本过程
[Envoy外部授权](https://cloudnative.to/envoy/intro/arch_overview/security/ext_authz_filter.html)过滤器调用授权服务以检查传入请求是否被授权,支持以gRpc和Http两种形式调用外部服务。

gRpc的接口定义如下，参考[完整DEMO](https://github.com/salrashid123/envoy_external_authz):

- [service.auth.v3.CheckRequest](https://cloudnative.to/envoy/api-v3/service/auth/v3/external_auth.proto.html#envoy-v3-api-msg-service-auth-v3-checkrequest):传递给授权服务的信息
- [service.auth.v3.CheckResponse](https://cloudnative.to/envoy/api-v3/service/auth/v3/external_auth.proto.html#service-auth-v3-checkresponse):要求授权服务返回的结果
- grpc的集群配置必须要有`http2_protocol_options`字段

http则按照Json格式交互，解析请求头，以返回的状态码作为决策结果：
- http.Request: 默认带 Authorization字段，如果是其他字段，需要在[AuthorizationRequest](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/http/ext_authz/v3/ext_authz.proto#envoy-v3-api-msg-extensions-filters-http-ext-authz-v3-httpservice)中配置。
- http.Status: 返回200状态码（默认）表示通过，其他状态表示不通过。如果需要携带其他字段，在AuthorizationResponse配置。
- http_service配置项中要求必须填写`uri`，但是该uri实际无效，所以不要用这个来定位请求URL，见[issue](https://github.com/envoyproxy/envoy/issues/5357)


`failure_mode_allow` 字段兜底：设置当外部认证服务不可访问的时候的通过策略，`false`表示拒绝，`true`表示通过。

## Tips

1. Docker Compose: 确定服务端口、通信网络
2. http_service配置项中要求必须填写uri，但是该uri实际无效，所以不要用这个来定位请求URL，见[issue](https://github.com/envoyproxy/envoy/issues/5357)
3. envoy.yaml极容易写错，写完验证一下配置

```yaml
http_service:
server_uri:
  uri: nothing_but_need
  cluster: ext_authz
  timeout: 0.250s
```

### 参考资料
- [GITHUB:Enovy Example](https://github.com/envoyproxy/envoy/tree/main/examples/ext_authz):Demo级示范，包括了Grpc和HttpService两种,[官方说明文档](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/ext_authz)
- [GITHUB:envoy_external_authz](https://github.com/salrashid123/envoy_external_authz):详细解释了gRpc方式的相关接口和流程
- [GITHUB:Istio Example](https://github.com/istio/istio/tree/master/samples/extauthz): Demo级示范，参考接口对接使用规范
- [GITHUB:authService](https://github.com/istio-ecosystem/authservice):产品级实现
- [GITHUB:envoy_ratelimit_example](https://github.com/Cluas/envoy_ratelimit_example):Demo+级示范，考虑了认证和限流，还有Redis集成，但是有许多小问题，没Work
