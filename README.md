# go-gen

Scaffolding for quickly generating frameworks for some languages.

## 项目目的

- 对一直给类语言的web框架，除一些重型框架外(django、beego等)，这些开发者只有自己的经验去构建工程目录，在开源项目上很少看见关于此类的框架脚手架，缺少一个类似npm的脚手架。
- 一个良好的工程项目应该从第一步构建清晰完善的工程目录结构开始。go-gen依赖的模板仓库都是结合了多个参考文章与部分开发者开发经验得来的

## 使用

- Windows：您可以直接到发布的版本中去下载可执行文件，并把它加入到您电脑的环境变量中快速用作与开发
- 对于其它系统使用者，您可以clone最新代码去构建打包

## 目前进度

- 目前支持python语言的fastapi命令构建项目
- gen 【command】 [flags]

-----

### example：

```shell
gen fastapi --name=demo-norm
```

### help：

```shell
gen -h
gen 【command】 -h
```

- 关于指令操作您完全可以去键入 -h 或者 --help查询帮助与示例

## 代码模板仓库

| Name         | description                                                 | Address                                                      |
| ------------ | ----------------------------------------------------------- | ------------------------------------------------------------ |
| fastapi-norm | 这是一个标准的fastapi框架的工程结构                         | [gen-fastapi-norm](https://github.com/NC-Cj/gen-fastapi-norm) |
| fastapi-easy | 这是一个简单的fastapi框架的工程结构，给开发者更多的操作空间 | [gen-fastapi-easy](https://github.com/NC-Cj/gen-fastapi-easy) |

## 关于数据库orm

1. 不论您来自任何语言，我非常推荐您的数据表都由 [Prisma ORM ](https://www.prisma.io/)去管理以及迁移更新
2. 关于它的学习成本对于开发者来说，他是低廉的，它带来的收益是客观的，您可以点击连接进入它的官网阅读