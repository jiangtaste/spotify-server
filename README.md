# spotify_server

for spotify token exchange & refresh

## usage

### RUN

1. rename ```config/config.yaml.dist``` to ```config/config.yaml```
2. edit config file.
3. ```go run main.go```

### Docker

1. build the docker
2. rename ```config/config.yaml.dist``` to ```config/config.yaml```
3. edit config file
4. ```docker run -p 8080:8080 -v YOUR_CONFIG_FILE /optconfig YOUR_DOCKER_REGISTER --name spotify_server```

