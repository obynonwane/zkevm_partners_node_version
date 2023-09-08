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
![Image Alt Text](https://res.cloudinary.com/dxec82vds/image/upload/v1694132865/Screenshot_2023-09-08_at_01.23.52_x5umpx.png)
### Run the below command for node versions on testnet

```
make testnet
```

![Image Alt Text](https://res.cloudinary.com/dxec82vds/image/upload/v1694132523/Screenshot_2023-09-08_at_01.21.27_upojd3.png)

## Contributors
- Obinna Johnson [@darealobinna](https://twitter.com/darealobinna)
