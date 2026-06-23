# Tilt local development for polyrepo layout
# Docs: https://docs.tilt.dev

local_resource(
    'proto-gen',
    'make proto',
    deps=['repos/proto-contracts', 'go.work'],
    labels=['build'],
)

docker_build(
    'go-grpc-template/user',
    '.',
    dockerfile='repos/user-service/Dockerfile',
    only=[
        'go.work',
        'repos/go-platform',
        'repos/proto-contracts',
        'repos/user-service',
    ],
)

docker_build(
    'go-grpc-template/gateway',
    '.',
    dockerfile='repos/gateway-service/Dockerfile',
    only=[
        'go.work',
        'repos/go-platform',
        'repos/proto-contracts',
        'repos/gateway-service',
    ],
)

k8s_yaml(kustomize('deploy/kubernetes/overlays/dev'))

k8s_resource(
    'user',
    port_forwards=['50051:50051', '9091:9091'],
    resource_deps=['proto-gen'],
    labels=['services'],
)

k8s_resource(
    'gateway',
    port_forwards=['8080:8080', '9090:9090'],
    resource_deps=['user'],
    labels=['services'],
)
