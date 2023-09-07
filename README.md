# Partners Node version: Mainnet or Testnet
## Script Requirement

 Install golang on your os: [See installation](https://go.dev/doc/install)

### Run below script to clone the repository to your working directory

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
