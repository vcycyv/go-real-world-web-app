# go-real-world-web-app

## Gin + Gorm2 搭建企业级应用
做了一个视频，介绍和演示了这个项目：https://www.bilibili.com/video/BV15P4y147Ce  
（对应的UI项目在这里：https://github.com/vcycyv/ui-real-world-web-app 基于react/redux）
此示例包含：  
1. 企业级应用的分层架构
2. gin + 最新版gorm的示例 （gorm2相对于jinzhu包的第一版gorm）
3. 四个middleware  
  1.1 cors  
  1.2 authentication  
  1.3 log  
  1.4 error message  
4. ldap 登录功能  
5. 支持持久存储的postgre sql docker脚本
6. 基于sqlite的自动化测试 
7. 初始化/客户化脚本

## 一些关于技术的选择  
### 1. 分层
  * DDD:  
    - controller  
    - service  
    - infrastructure
  * The clean architecture  https://zhuanlan.zhihu.com/p/64343082
### 2. 数据库主键
* 自增优点：
  - 存储空间小
  - 查询效率高
* uuid优点：
  - 分布式友好
  - 避免暴露业务规模
* snowflake
### 3. 返回值struct vs. pointer  
https://www.ardanlabs.com/bookshop/2014/12/using-pointers-in-go.html
### 4. 尽量避免在运行时创建实例
### 5. ldap
apache directory studio的使用：
https://www.bilibili.com/video/BV1kh411h7yB?from=search&seid=4698533963557832110
### 6. 运行在docker中的postgre sql
### 7. 命名规范  
包的名称用单数  
https://rakyll.org/style-packages/  
https://github.com/golang-standards/project-layout/issues/7
### 8. 自动化测试



## 初始化/客户化脚本：
for filename in `find . -type f -name 'book*'`; do mv -v "$filename" "${filename//book/[replacement]}"; done
find ./ -type f \( ! -iname "replace.sh" \) -exec sed -i -e 's/book/[replacement]/g' {} \;
find ./ -type f \( ! -iname "replace.sh" \) -exec sed -i -e 's/Book/[Replacement]/g' {} \;
find ./ -type f \( ! -iname "replace.sh" \) -exec sed -i -e 's/"bookshop"/"[project]"/g' {} \;
find ./ -type f \( ! -iname "replace.sh" \) -exec sed -i -e 's/POSTGRES_DB: bookshop/POSTGRES_DB: [project]/g' {} \;

