# Minecraft Operator for Kubernetes

## Sample
Note: See [itzg/docker-minecraft-server](https://github.com/itzg/docker-minecraft-server) for a full list of valid inputs to the config environment variables.

Deploy the operator, then deploy this:
```yaml
apiVersion: game.laputacloud.co/v1alpha2
kind: Minecraft
metadata:
  name: sample
spec:
  config:
    EULA: "TRUE"
    TZ: US/New York
    VERSION: LATEST
    WHITELIST: YOUR_USERNAME
  storageSize: 8Gi
  storageClassName: YOUR_STORAGECLASS
  serve: true # this will cause a loadbalancer type service to be created
```
