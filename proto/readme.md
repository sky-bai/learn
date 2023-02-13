用户 文章两大实体类

.proto 文件夹中含有两个 gRPC 服务，分别为 article 和 user，我们在这两个文件夹中定义各自所需要的 messages 和 services。

一般情况下，我们会将编译生成的 pb.go 文件生成在与 proto 文件相同的目录，这样我们就不需要再创建相同的目录层级结构来存放 pb.go 文件了。

由于同一文件夹下的 pb.go 文件同属于一个 package，所以在定义 proto 文件的时候，

相同文件夹下的 proto 文件也应声明为同一的 package，并且和文件夹同名，这是因为生成的 pb.go 文件的 package 是取自 proto package 的。

同属于一个包内的 proto 文件之间的引用也需要声明 import ，因为每个 proto 文件都是相互独立的

，这点不像 Go（包内所有定义均可见）。我们的项目 user 模块下 service.proto 就需要用到 message.proto 中的 message 定义，代码是这样写的：

user/service.proto:


可以看到，我们在每个 proto 文件中都声明了 package 和 option go_package，这两个声明都是包声明，到底两者有什么关系，这也是我开始比较迷惑的。

我是这样理解的，package 属于 proto 文件自身的范围定义，与生成的 go 代码无关，

它不知道 go 代码的存在（但 go 代码的 package 名往往会取自它）。 如果要指定go文件的包名的话就需要设置  option go_package 则是用来告诉 protoc-gen-go 生成的 go 代码的 package 名。

这个 proto 的 package 的存在是为了避免当导入其他 proto 文件时导致的文件内的命名冲突。

所以，当导入非本包的 message 时，需要加 package 前缀，

如 service.proto 文件中引用的 Article.Articles，点号选择符前为 package，后为 message。同包内的引用不需要加包名前缀。


而 option go_package 的声明就和生成的 go 代码相关了，它定义了生成的 go 文件所属包的完整包名，

所谓完整，是指相对于该项目的完整的包路径，应以项目的 Module Name 为前缀。 需要完整的包路径 。 // 项目名加+ 模块名

如果不声明这一项会怎么样？最开始我是没有加这项声明的，后来发现 依赖这个文件的 其他包的 proto 文件 所生成的 go 代码 中（注意断句，已用斜体和正体标示），

引入本文件所生成的 go 包时，import 的路径并不是基于项目 Module 的完整路径，而是在执行 protoc 命令时相对于 --proto_path 的包路径，

这在 go build 时是找不到要导入的包的。这里听起来可能有点绕，建议大家亲自尝试一下。

// 也就是说指定go_package包名为项目名加模块名 

// 插件 protoc-gen-go 是不支持多包同时编译的