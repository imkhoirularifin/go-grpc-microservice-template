# Tilt local development configuration
# Docs: https://docs.tilt.dev

load('ext://restart_process', 'docker_build_with_restart')

local_resource(
    'proto-gen',
    'make proto',
    deps=['proto/contracts', 'buf.yaml', 'buf.gen.yaml'],
    labels=['build'],
)

docker_build(
    'go-grpc-template/user',
    '.',
    dockerfile='deploy/docker/user.Dockerfile',
    only=['cmd', 'internal', 'pkg', 'lib', 'gen', 'go.mod', 'go.sum'],
    live_update=[
        sync('./cmd', '/src/cmd'),
        sync('./internal', '/src/internal'),
        sync('./pkg', '/src/pkg'),
        sync('./lib', '/src/lib'),
        sync('./gen', '/src/gen'),
    ],
)

docker_build(
    'go-grpc-template/gateway',
    '.',
    dockerfile='deploy/docker/gateway.Dockerfile',
    only=['cmd', 'internal', 'pkg', 'lib', 'gen', 'go.mod', 'go.sum'],
    live_update=[
        sync('./cmd', '/src/cmd'),
        sync('./internal', '/src/internal'),
        sync('./pkg', '/src/pkg'),
        sync('./lib', '/src/lib'),
        sync('./gen', '/src/gen'),
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
