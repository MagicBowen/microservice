## MYSQL

### docker

```sh
docker pull mysql:5.6

docker run --name mysql -e MYSQL_ROOT_PASSWORD=omit_pass -d mysql:5.6
```

```sh
docker run --net=examples_default --name bcdb-mysql -e MYSQL_ROOT_PASSWORD=omit_pass -d mysql:5.6
```

command:
```sh
docker run -it --net=examples_default --rm mysql:5.6 mysql -hbcdb-mysql -uroot -p
```

docker compose
```yml
  db:
    image: mysql
    command: --default-authentication-plugin=mysql_native_password
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: example

  adminer:
    image: adminer
    restart: always
    ports:
      - 8080:8080
```

run mysql in local host network:
```sh
docker run -p 3306:3306 --name bcdb-mysql -e MYSQL_ROOT_PASSWORD=omit_pass -d mysql:5.6
```

### useful commands

创建数据库：
create database 库名
删除数据库：
drop database 库名
转到数据库
use database 库名

创建表：
create table 表名(列1 类型, 列2 类型...)

使用旧表创建新表：
create table 表名 like 旧表名
create table 表名 as select 列1, 列2... from 旧表 definition only

删除表：
drop table 表名
显示所有表：
show tables
显示表详情：
desc 表名
show create table 表名
清空表：
truncate table 表名
delete from 表名
修改表名：
alter table 原名 rename to 新名
添加列：
alter table 表名 add column 列名 类型
删除列：
alter table 表名 drop column 列名
修改列名：
alter table 表名 change 原名 新名 类型
修改列类型：
alter table 表名 modify 列名 类型
更新数据：
update 表名 set 字段=值 where 条件
删除数据：
delete from 表名 where 条件
插入数据：
insert into 表名(字段1, 字段2...) values(值1, 值2...)
排序：
select * from 表名 order by 字段 [desc|asc]
计数行：
select count(*) from 表名

