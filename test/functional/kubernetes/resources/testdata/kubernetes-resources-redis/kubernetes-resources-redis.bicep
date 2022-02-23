import kubernetes from kubernetes

resource redisService 'kubernetes.core/Service@v1' existing = {
  metadata: {
    name: 'redis-svc'
  }
}

resource redisSecret 'kubernetes.core/Secret@v1' existing = {
  metadata: {
    name: 'redis-pw'
  }
}

resource app 'radius.dev/Application@v1alpha3' = {
  name: 'kubernetes-resources-redis'

  resource webapp 'Container' = {
    name: 'todoapp'
    properties: {
      container: {
        image: 'radius.azurecr.io/magpie:latest'
        env: {
        }
      }
      connections: {
        redis: {
          kind: 'redislabs.com/Redis'
          source: redis.id
        }
      }
    }
  }

  resource redis 'redislabs.com.RedisCache' = {
    name: 'redis'
    properties: {
      host: '${redisService.metadata.name}.${redisService.metadata.namespace}.svc.cluster.local'
      port: redisService.spec.ports[0].port
      secrets: {
        connectionString: '${redisService.metadata.name}.${redisService.metadata.namespace}.svc.cluster.local:${redisService.spec.ports[0].port},password=${base64ToString(redisSecret.data['redis-password'])}'
        password: base64ToString(redisSecret.data['redis-password'])
      }
    }
  }
}