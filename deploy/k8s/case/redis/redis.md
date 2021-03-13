# k8s 安装 redis

## 创建pv
```text
apiVersion: v1
kind: PersistentVolume
metadata:
  name: redis-pv
  namespace: redis
  labels:
    pv: redis-pv
spec:
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: /Volumes/backup/k8s/redis
  persistentVolumeReclaimPolicy: Retain
  volumeMode: Filesystem
```


## 创建pvc

## 创建deployment

## 创建svc