# Golang 学习项目
本项目是一个 go 语言学习项目. 包含了一些作者在学习 go 语言的过程中写的练习草稿.  
其中最重要的是 meal-coupon项目.  
## meal-coupon
### 需求
加班餐券管理系统
1. 已有整个部门的人员名单和各team的人员名单，和每个team的leader和部门admin的人员信息
2. 可以通过员工号登陆
3. team member登陆后能看到自己team的当月的餐券申请数量，
4. team leader登陆后不仅可以看到数量，每月5号之前还可以编辑餐券数量
5. admin登陆后可以看到各个team的申请数量
6. 可以没有界面和数据库(不会写前端的悲哀)
### 设计
- 因为没有数据库, 因此数据持久化使用文件的方式
- 系统启动时读取文件加载到内存
- 偷懒一下,除了餐券数量可以修改外, 其他数据均不做修改
- 对外接口使用 url 参数的方式提供 
- 登陆信息存入 cookie, 这个是个大坑
### 使用到的知识点
- http, time, cookie, 
- map, list
- 字符串数字转换
- 异常抛出与捕获
- 读写文件
### 系统部署
1. 程序运行当前文件目录中增加一个 data 目录
2. 进入 data 目录中增加两个文件 team.dat, staff.dat
- team.dat 中存放小组信息, 必须包含标题栏. 标题包括如下字段：组号, 组名, 组长, 餐券数
- staff.dat 中存放员工信息, 必须包含标题栏. 标题包括如下字段：编号, 姓名, 密码, 所在组号, 角色
3. go-learn 目录下执行 go build 命令, 即可生成 go-learn 的可以执行文件, 直接运行即可
4. 对于 meal-coupon 项目, 外侧的其他几个日常学习包并没有使用到, 可以删除

### 用户使用
1. 登录 http://localhost:8080/login?username=10101&password=1234
2. 数据操作
- 查看餐券数量 http://localhost:8080/query
- 修改餐券数量 http://localhost:8080/modify?couponNum=2