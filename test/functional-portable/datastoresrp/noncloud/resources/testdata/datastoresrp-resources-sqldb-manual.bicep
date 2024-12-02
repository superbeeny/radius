extension radius

@description('Specifies the location for resources.')
param location string = 'global'

@description('Specifies the image for the container resource.')
param magpieImage string

@description('Specifies the port for the container resource.')
param magpiePort int = 3000

@description('Specifies the environment for resources.')
param environment string = 'test'

@description('Specifies the image for the sql container resource.')
param sqlImage string = 'mcr.microsoft.com/mssql/server:2019-latest'

@description('Specifies the port for the container resource.')
param sqlPort int = 1433

@description('Specifies the SQL username.')
param username string = 'sa'

@description('Specifies the SQL password.')
@secure()
param password string = newGuid()

resource app 'Applications.Core/applications@2023-10-01-preview' = {
  name: 'dsrp-resources-sql'
  location: location
  properties: {
    environment: environment
  }
}

resource webapp 'Applications.Core/containers@2023-10-01-preview' = {
  name: 'sql-app-ctnr'
  location: location
  properties: {
    application: app.id
    connections: {
      sql: {
        source: db.id
      }
    }
    container: {
      image: magpieImage
      env: {
        CONNECTION_SQL_CONNECTIONSTRING: {
          value: db.listSecrets().connectionString
        }
      }
      readinessProbe: {
        kind: 'httpGet'
        containerPort: magpiePort
        path: '/healthz'
      }
    }
  }
}

resource db 'Applications.Datastores/sqlDatabases@2023-10-01-preview' = {
  name: 'sql-db'
  location: location
  properties: {
    application: app.id
    environment: environment
    server: 'sql-ctnr'
    database: 'master'
    resourceProvisioning: 'manual'
    port: sqlPort
    username: username
    secrets: {
      password: password
    }
  }
}

resource sqlContainer 'Applications.Core/containers@2023-10-01-preview' = {
  name: 'sql-ctnr'
  location: location
  properties: {
    application: app.id
    container: {
      image: sqlImage
      env: {
        ACCEPT_EULA: {
          value: 'Y'
        }
        MSSQL_PID: {
          value: 'Developer'
        }
        MSSQL_SA_PASSWORD: {
          value: password
        }
      }
      ports: {
        sql: {
          containerPort: sqlPort
        }
      }
    }
    runtimes: {
      kubernetes: {
        pod: {
          resources: {
            requests: {
              cpu: '1200m'
              memory: '2048Mi'
            }
            limits: {
              cpu: '1500m'
              memory: '4096Mi'
            }
          }
        }
      }
    }
  }
}
