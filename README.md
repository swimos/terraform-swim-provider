# Terraform Provider Swim

### Install

Run the formatter
```shell
go fmt ./...
```

Build the provider.
```shell
go build -o terraform-provider-swim
```

Build and install the provider locally.
```shell
make install
```

### Run

Run the Swim server.
```shell
(cd swim-server && ./gradlew run)
```

Initialize the workspace and apply the terraform plan.
```shell
(cd examples && terraform init && terraform apply)
```

### Clean up
Destroy the terraform stack.
```shell
(cd examples && terraform destroy)
```