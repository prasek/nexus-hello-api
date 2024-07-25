# nexus-hello-api

## Getting started with Temporal OSS server

```
git clone https://github.com/prasek/nexus-hello-api.git
git clone https://github.com/prasek/nexus-hello-provider.git
git clone https://github.com/prasek/nexus-hello-consumer.git
```

### Get `temporal` CLI v0.14.0-nexus.0

```
curl -sSf https://temporal.download/cli.sh | sh -s -- --version v0.14.0-nexus.0 --dir .

./bin/temporal --version
```

### Spin up environment

#### Start temporal server

```
./bin/temporal server start-dev --dynamic-config-value system.enableNexus=true --http-port 7243
```

Open Temporal UI on http://localhost:8233/

### Initialize environment

In a separate terminal window

#### Create caller and target namespaces

```
./bin/temporal operator namespace create --namespace my-target-namespace
./bin/temporal operator namespace create --namespace my-caller-namespace
```

#### Create Nexus endpoint

```
./bin/temporal operator nexus endpoint delete --name myendpoint
./bin/temporal operator nexus endpoint create \
  --name myendpoint \
  --target-namespace my-target-namespace \
  --target-task-queue my-handler-task-queue \
  --description-file ./nexus-hello-api/description.md
```

### Make Nexus calls across namespace boundaries

In separate terminal windows:

### Nexus handler worker

```
cd nexus-hello-provider/app

go run ./worker \
    -target-host localhost:7233 \
    -namespace my-target-namespace
```

### Nexus consumer workflow worker

```
cd nexus-hello-consumer/app
go run ./worker \
    -target-host localhost:7233 \
    -namespace my-caller-namespace
```

### Start caller workflow

```
cd nexus-hello-consumer/app
go run ./starter \
    -target-host localhost:7233 \
    -namespace my-caller-namespace
```

### Output

which should result in:
```
2024/07/23 19:57:40 Workflow result: Nexus Echo 👋
2024/07/23 19:57:40 Started workflow WorkflowID nexus_hello_caller_workflow_20240723195740 RunID c9789128-2fcd-4083-829d-95e43279f6d7
2024/07/23 19:57:40 Workflow result: ¡Hola! Nexus 👋
```
