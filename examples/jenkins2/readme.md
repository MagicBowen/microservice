## Jenkins2

### setup by docker

```sh
docker run --name myjenkins -p 10000:8080 -p 50000:50000 -v ~/codes/jenkins jenkins/jenkins
```

When the jenkins server startup, it will show a key for admin user login, just like `4a13c7db6c8646859ea78610513fcbe8`

tutorial:
- https://hub.docker.com/_/jenkins
- https://www.cloudbees.com/blog/get-started-jenkins-20-docker

### 核心概念

pipeline是一级概念。通过Groovy脚本或者DSL描述pipeline。
pipeline as code， Jenkinfile

- 节点（node）：标签、主目录、executor数量、环境变量等
    - 主节点（master）
    - 从节点（slave）
- 代理器（agent）：在节点上配置连接方式（SSH和证书）用于和主节点通讯安装agent。
- 执行器（executor）：基于主目录的工作空间

标签的作用：识别节点；分组；标识节点特征；

### pipeline

pipeline :
- stage
    - step
  post
post
  - mail
  - im
  - html

parallel
node
lock
script
parameter
input
conditions

trigger -> stage -> post

### security

- RBAC : Role-Based Access Control
- Vault: 在有限生命周期内存储凭证的方法

允许用户注册。后台管理员为用户分组，基于组和用户进行权限勾选。
Jenkins可以存储各种用户账号密码和凭证文件。然后设置按照一定的域去使用各种凭证。凭证范围：系统、全局、用户

配置角色，为角色分配权限。为不同的用户和用户组分配角色。在pipeline的代码中可以使用凭证。

### other Tools

- Travis CI
- Strider
- GitLab CI
- cloudBee

将对应的语言技术栈的常规配置和依赖打包成docker，提供运行环境的托管。简化了对应的配置和环境管理工作。
按照使用资源进行收费。

### TODO

- [ ] try Artifactory
- [ ] try Vault
- [ ] try Jenkins2
- [ ] integrate jenkins with dynamic resources management

