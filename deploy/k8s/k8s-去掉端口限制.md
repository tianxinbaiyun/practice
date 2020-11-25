@[TOC]()
# k8s-去掉端口限制.md


## linux 
```
vi /etc/kubernetes/manifests/kube-apiserver.yaml
```


添加到如下位置就行了
```
- command:

   - kube-apiserver

   - --service-node-port-range=1-65535 
```


直接删除kube-apiserver  pod 就行了 会自动重启

```
kubectl delete pod kube-apiserver -n kube-system
```

## mac

### 1.登陆到Docker VM

```
docker run \
--rm \
-it \
--privileged \
--pid=host \
walkerlee/nsenter -t 1 -m -u -i -n bash
```
### 2.编辑kube-apiserver.yaml
```
vi /etc/kubernetes/manifests/kube-apiserver.yaml
```

### 3.添加service-node-port-range范围
```
spec:
  containers:
  - command:
    - kube-apiserver
    - --advertise-address=192.168.65.3
    ...
    - --service-node-port-range=1-65535
    ...
```

### 4. 重启api-server服务
```
kubectl get po -n kube-system

kubectl delete po -n kube-system kube-apiserver-docker-desktop
```

打开docker桌面app，重启

### 5.修改k8s端口尝试


