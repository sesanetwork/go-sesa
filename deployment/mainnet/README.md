# Prerequisites
Install Docker following the installation guide for Linux OS: [https://docs.docker.com/engine/installation/](https://docs.docker.com/engine/installation/)
* [CentOS](https://docs.docker.com/install/linux/docker-ce/centos) 
* [Ubuntu](https://docs.docker.com/install/linux/docker-ce/ubuntu)

Install docker compose
* [Docker compose](https://docs.docker.com/compose/install/)

Build dependencies
* [Network](https://docs.sesa.network/network/build-dependencies)
* [Create validator wallet](https://docs.sesa.network/network/run-validator-node/mainnet-validator-node#create-a-validator-wallet)

### Buiding image

```
cd <path>/go-sesa
make NET=mainnet sesa-image
```

### Configuration
```
Update password wallet in "/etc/password"
Go to <path>/go-sesa/deployment/mainnet
Update config in docker-compose.yml:
    validator.id <id>
    validator.pubkey <pubkey>
    bootnodes <bootnodes>
```

### Running node
``` 
docker compose up -d
```

### Logging
````
docker logs -f --tail 10 sesa