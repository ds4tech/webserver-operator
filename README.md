# Webserver-operator

### Web-server -operator maintains Webserver application.

## Introduction <a name="intro"></a>

Remember to run build Operator image every time, a change is made on helm templates.
For istance when required to change container port:

```          ports:
            - name: http
              containerPort: 9090
```
upgrade image version in helm-charts/webserver/Chart.yaml
and run 

`make docker-build docker-push`
