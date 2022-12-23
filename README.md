# Webserver-operator

### Web-server -operator maintains Webserver application.

## Introduction <a name="intro"></a>

Remember to build the Operator image on helm templates when a change is made.
For instance, when required to change container port:

```
   spec:
      containers:
        - name: {{ .Chart.Name }}
          ports:
            - name: http
              containerPort: 9090
```
upgrade image version in helm-charts/webserver/Chart.yaml
```
apiVersion: v2
appVersion: 0.0.4
description: A Helm chart for Kubernetes
name: webserver
type: application
version: 0.0.2
```
then upgrade docker image version in a Makefile
`VERSION ?= 0.0.2`

and run 

`make docker-build docker-push`
