# 简单的介质仓库

##### 因为需要一个仓库存放打包好的介质去提供给相关人员下载，所以开发了这样一个简单的介质仓库程序
### 使用方法
1、clone项目或下载相关zip包到本地

2、进入Release目录中，修改conf目录下的app.conf文件，填写保存路径和MySQL数据库信息 

3、根据操作系统、位数，选择运行packageRepo-xxx(.exe)等可执行文件即可（Linux等操作系统可能需要手动赋予可执行权限）

4、打开浏览器访问http://ip:port就可使用

### 自行编译
##### 方法一
```
go get -u github.com/godfather1103/packageRepo
```
上述命令运行结束后，在GoPath下bin中将生成对应的可执行文件，然后将相关文件拷贝到与conf，views等文件夹同级的目录下，运行即可。

##### 方法二
下载相关代码到对应目录下，即：$GoPath/src/github.com/godfather1103/packageRepo，打开终端并将路径移动到$GoPath/src/github.com/godfather1103/packageRepo，执行如下命令即可。
```
go build -i
```
###### 该项目需要[beego](https://github.com/astaxie/beego)等依赖库，需要自行下载到对应文件夹中

