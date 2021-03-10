# 使用k8s安装mysql


## 1.简单安装

### 1.1.新建mysql-rc.yaml

```shell
vi mysql-rc.yaml
```


```text
apiVersion: v1
kind: ReplicationController
metadata:
  name: mysql-rc
  labels:
    name: mysql-rc
spec:
  replicas: 1
  selector:
    name: mysql-pod
  template:
    metadata:
      labels:
        name: mysql-pod
    spec:
      containers:
        - name: mysql
          image: mysql
          imagePullPolicy: IfNotPresent
          ports:
            - containerPort: 3306
          env:
            - name: MYSQL_ROOT_PASSWORD
              value: "123456"
```

### 1.2.创建mysql-svc.yaml
```text
vi mysql-svc.yaml
```

```text
apiVersion: v1
kind: Service
metadata:
  name: mysql-svc
  labels:
    name: mysql-svc
spec:
  type: NodePort
  ports:
    - port: 3306
      protocol: TCP
      targetPort: 3306
      name: http
      nodePort: 30306
  selector:
    name: mysql-pod
```

### 1.3.k8s 执行文件，下载mysql镜像和运行mysqlr容器
```text
  [root@k8s-master ~]# kubectl create -f mysql-rc.yaml
  replicationcontroller "mysql-rc" created
  [root@k8s-master ~]# kubectl create -f mysql-svc.yaml
  service "mysql-svc" created
```


## 2.数据持久化挂载安装


### 2.1、创建PV
```text
apiVersion: v1
kind: PersistentVolume
metadata:
  name: model-db-pv
spec:
  accessModes:
    - ReadWriteOnce
  capacity:
    storage: 5Gi
  claimRef:
    apiVersion: v1
    kind: PersistentVolumeClaim
    name: model-db-pv-claim
    namespace: default
  hostPath:
    path: /Volumes/DATA/k8s/case/mysql/persistence/data
  persistentVolumeReclaimPolicy: Retain
  volumeMode: Filesystem
```



### 2.2、创建PVC
```text
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: model-db-pv-claim
  namespace: default
  labels:
    app: model-mysql
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi
```

### 2.3、确认pv及pvc的状态
```text
# kubectl get pv
# kubectl get pvc
```

### 2.4. 创建mysql configMap
```text
apiVersion: v1
kind: ConfigMap
metadata:
  name: model-db-config
  namespace: default
  labels:
    app: model-db
data:
  my.cnf: |-
    [client]
    default-character-set=utf8mb4
    [mysql]
    default-character-set=utf8mb4
    [mysqld]
    character-set-server = utf8mb4
    collation-server = utf8mb4_bin
    init_connect='SET NAMES utf8mb4'
    skip-character-set-client-handshake = true
    max_connections=2000
    secure_file_priv=/var/lib/mysql
    bind-address=0.0.0.0
    symbolic-links=0
    sql_mode=''
```



### 2.5、部署(Deployment)文件
```text
apiVersion: apps/v1
kind: Deployment
metadata:
  name: model-db
  namespace: default
spec:
  replicas: 1
  selector:
    matchLabels:
      app: model-mysql
  template:
    metadata:
      labels:
        app: model-mysql
    spec:
      containers:
        - args:
            - --datadir
            - /var/lib/mysql/datadir
          env:
            - name: MYSQL_ROOT_PASSWORD
              value: root
            - name: MYSQL_USER
              value: user
            - name: MYSQL_PASSWORD
              value: user
          image: mysql:8.0
          name: model-db-container
          ports:
            - containerPort: 3306
              name: dbapi
          volumeMounts:
            - mountPath: /var/lib/mysql
              name: model-db-storage
            - name: config
              mountPath: /etc/mysql/conf.d/my.cnf
              subPath: my.cnf
      volumes:
        - name: model-db-storage
          persistentVolumeClaim:
            claimName: model-db-pv-claim
        - name: config
          configMap:
            name: model-db-config
        - name: localtime
          hostPath:
            type: File
            path: /etc/localtime
```



### 2.6、创建svc（service）
```text
apiVersion: v1
kind: Service
metadata:
  labels:
    app: model-mysql
  name: model-db-svc
  namespace: default
spec:
  type: NodePort
  ports:
    - name: http
      port: 3306
      nodePort: 30336
      protocol: TCP
      targetPort: 3306
  selector:
    app: model-mysql
```
