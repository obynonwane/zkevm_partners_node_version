# dao_escrow
## Script Requirement

```
Golang needs to be installed on your os, check link for your os: https://go.dev/doc/install
```
### Project setup run below script to clone the repository

```
git clone https://github.com/obynonwane/zkevm_partners_node_version.git
```


### cd into the cloned repository
```
cd zkevm_partners_node_version
```

### download and install the required dependencies

```
go mod tidy
```

### Create your .env file from env.example: copy content of env.example to .env and update accordingly
```
touch .env
```

### Run the below command for node versions on mainnet

```
make mainnet
```

### Run the below command for node versions on testnet

```
make testnet
```

### Customize configuration

See [Configuration Reference](http://cli.vuejs.org/config/).