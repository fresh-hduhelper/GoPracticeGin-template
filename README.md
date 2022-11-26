# GoPracticeGin

## 是什么

一系列有关 `Gin` 框架的练习题，目的是帮助参与者巩固所学

## 注意

**一定要先运行对应章节的前置命令，不然没法正常测试 :)**

## 练习

在练习章节，您需要先在命令行输入下面这条命令:

```shell
go run build.go warming-up/
```

之后的命令会在章节一开始给出，不要忘了呦 :P

这里的练习依托 GitHub Classroom

为快速适应环境，让我们先来热热身，[点我](warming-up/README.md)

### 第一章 常见请求参数 : [common-params](common-params/README.md)

```shell
go run build.go common-params/
```

你将学习使用 `Gin` 来处理一些常见的 Http 请求，包括但不限于 GET, POST, PUT 等

### 第二章 中间件 : [middlewire](middlewire/README.md)

```shell
go run build.go middlewire/
```

在这里，你将学习编写自己的中间件来处理特定的请求

### 第三章 鉴权 : [authority](authority/README.md)

```shell
go run build.go authority/
```

你将学习权限控制，限制用户的行为，抵御老王八蛋的攻击(尽管挡不住，你还可以挣扎)

### 第四章 自定义 Server : [custom-server](custom-server/README.md)

```shell
go run build.go custom-server/
```

动手配置一个 Gin 的 Server，而不是用 gin.Default()

### 第五章 配置化 : [configuration](configuration/README.md)

```shell
go run build.go configuration/
```

是时候给你的项目引入适当的配置文件了

### 第六章 项目分层 : [project-layer](project-layer/README.md)

```shell
go run build.go project-layer/
```

一切的尽头是设计! 设计一个合理优雅的项目结构往往强过丰富的功能(大概)

### 第七章 数据库简单应用 : [database-sql](database-sql/README.md)

```shell
go run build.go database-sql/
```

数据，数据，数据!

### 第八章 模板(MVC 设计模式) : [template-mvc](template-mvc/README.md)

```shell
go run build.go template-mvc/
```

前后一把梭(bushi)

### TODO

## 其他

有问题请提 `issue`

欢迎 `pr`
