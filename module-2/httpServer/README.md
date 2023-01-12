## 打包镜像:https://docs.docker.com/engine/reference/commandline/build/
```shell
docker build --force-rm -t hufeifei/cloud-native-study:1.0.0.1 .
```

#### --force-rm：用来删除构建过程中生成得中间镜像

## 将镜像推送至 docker 官方镜像仓库:https://docs.docker.com/docker-hub/repos/
```shell
docker push hufeifei/cloud-native-study:1.0.0.1
```

#### 推送地址：https://hub.docker.com/layers/hufeifei/cloud-native-study/1.0.0.1/images/sha256-a4e1990de093fc7bdcb62aabce22f895bfdcb006fdbaaf1abe9a88dc775ef994?context=repo

#### Q: denied: requested access to the resource is denied
A: 需要登录dockerhud:
```shell
docker login -u username -p password
```

## 通过 docker 命令本地启动 httpserver
```shell
docker run -p 80:80 -d hufeifei/cloud-native-study:1.0.0.1
```

## 通过 nsenter 进入容器查看 IP 配置
```shell
1. 获取CONTAINER ID
docker ps
2. 获取PID
docker inspect -f be99655e4dd8
3. 进入该容器的网络命令空间
nsenter -t 13364 -n
4. 查看IP配置
ip addr
```


