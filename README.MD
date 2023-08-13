clone后注意修改`config/config.go`中的数据库配置，程序第一次运行后会生成`config.json`配置文件，可以在`main`函数中进行修改

如果要使用`sqlite`数据库，请修改`model/db`中`InitDB`函数的数据库初始化，但是提交的时候不需要将修改提交上来

编写代码的时候注意`orm`中的字段响应到前端的`json`标签，避免响应敏感字段

项目结构：

~~~shell
├─config // 配置文件和程序
├─controller // 接口的handler
├─middleware // 中间件:认证中间件...
├─model // orm model和db连接
├─public // 公开资源
├─test // 单元测试文件
├─utils // 工具包:jwt...
├─router.go // 接口注册
├─.gitignore // gitignore文件
└─main.go // 工具包

~~~

参考下面的规范、但是不完全照搬

1. 主分支指的是在整个软件生命周期内一直存在的分支，不允许删除commit的历史记录，更不允许删除分支。例如，master分支记录的就是整个软件的发布记录（换而言之，可以拉取其中的任一commit，无须修改即可对外发布），只能从release
branches或hotfixes两个分支中合并，每次commit必须打标签。

2. 辅助性分支则一般用于新功能开发、bug修复、新版QA，完成相应的功能之后可以删除这些分支，但是为了使整个开发过程有据可查，一般都予以保留。例如，已经对外发布的V0.1版需要紧急修复一个自测、内测均未发现的bug，则需要直接从master分支的V0.1开出一个紧急修复分支hotfixes-XXX，完成后同时并入master、develop两个主分支，hotfixes-XXX分支的存留取决于项目管理员。


远程仓库分支：
- main 或 master 分支： 这是生产环境的主要分支，存放着稳定、可部署的代码。
- develop 分支： 这是集成开发的主要分支，用于开发和测试新功能、bug 修复等。其他分支的更改最终都会合并到这个分支。
- feature 分支： 用于开发新功能的分支，从 develop 分支分离出来。每个功能都在独立的 feature 分支上进行开发，开发完成后合并回 develop 分支。
- bugfix 分支： 用于修复 bug 的分支，从 develop 分支分离出来。修复完成后合并回 develop 分支。

本地分支：
- main 或 master 分支： 本地仓库上的主分支，通常跟踪远程的 main 或 master 分支。
- develop 分支： 本地仓库上的集成开发分支，用于与远程 develop 分支同步，进行功能开发和测试。
- 功能分支： 每个人可以在本地创建独立的功能分支，用于开发单个功能。这些分支从本地 develop 分支分离出来，开发完成后再合并回 develop 分支。
- 个人分支： 有些团队成员可能会创建个人分支，用于测试、实验、备份等。这些分支不会直接影响主要分支。