## 打包镜像:https://docs.docker.com/engine/reference/commandline/build/
```shell
docker build --force-rm -t hufeifei/cloud-native-study:1.0.0.1 .
```

#### --force-rm：用来删除构建过程中生成得中间镜像

## 推送到dockerhub
```shell
docker push hufeifei/cloud-native-study:1.0.0.1
```

#### Q: denied: requested access to the resource is denied
A: 需要登录dockerhud:
```shell
docker login
```