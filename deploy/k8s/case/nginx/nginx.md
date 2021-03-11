# k8s 安装nginx

## 创建deployment

```text
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deploy
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx-cluster
  template:
    metadata:
      labels:
        app: nginx-cluster
    spec:
      containers:
        - name: nginx-cluster
          image: nginx
          ports:
            - containerPort: 80
          resources:
            requests:
              cpu: 1
              memory: 500Mi
            limits:
              cpu: 2
              memory: 1024Mi
```


## 创建svc

```text
apiVersion: v1
kind: Service
metadata:
  name: nginx-service
  labels:
    app: nginx-service
spec:
  type: NodePort
  selector:
    app: nginx-cluster
  ports:
    - port: 8000
      targetPort: 80
      nodePort: 32500
```

## 查看服务

- 查看deployment
```text
macdeiMac-Pro:nginx mac$ kubectl get deployment
NAME           READY   UP-TO-DATE   AVAILABLE   AGE
nginx-deploy   1/1     1            1           6m24s

```

- 查看svc
```text
macdeiMac-Pro:nginx mac$ kubectl get svc
NAME            TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
kubernetes      ClusterIP   10.96.0.1        <none>        443/TCP          17m
nginx-service   NodePort    10.102.119.183   <none>        8000:32500/TCP   6m13
```

- 查看网页
```text
macdeiMac-Pro:nginx mac$ curl http://127.0.0.1:32500
```
```html
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>

```