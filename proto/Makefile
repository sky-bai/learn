op1:
	protoc --go_out=plugins=grpc,paths=source_relative:. ./user/*.proto



# proto 文件中使用 import 关键字导入的路径一定是要基于 --proto_path 参数所指定的路径的。该参数如果不指定，默认为 pwd ，也可以指定多个以包含所有所需文件。

# --go_out 参数是用来指定 protoc-gen-go 插件的工作方式 和 go 代码目录架构的生成位置  // 以什么方式去生成go代码 以及生成的文件目录

# --go_out 参数的写法是，参数之间用逗号隔开，最后加上冒号来指定代码目录架构的生成位置，例如：--go_out=plugins=grpc,paths=import:. 。paths 参数有两个选项，import 和 source_relative 。

# 默认为 import ，代表按照生成的 go 代码的包的全路径去创建目录层级，source_relative 代表按照 proto 源文件的目录层级去创建 go 代码的目录层级，如果目录已存在则不用创建。

# --go_out=plugins=grpc,paths=import:.  // 生成的go代码的包的全路径去创建目录层级   go_package 选项指定的包名

# import 代表按照生成的 go 代码的包的全路径去创建目录层级，source_relative 代表按照 proto 源文件的目录层级去创建 go 代码的目录层级，如果目录已存在则不用创建。



# option go_package 声明 是为了让生成的其他 go 包（依赖方）可以正确 import 到本包（被依赖方）  // todo 查看依赖房import的代码
# --go_out=paths=source_relative:. 参数 是为了让加了 option go_package 声明的 proto 文件可以将 go 代码编译到与其同目录。

# 为了统一性，我会将所有 proto 文件中的 import 路径写为相对于项目根目录的路径，然后 protoc 的执行总是在项目根目录下进行：

op2:
	protoc --proto_path=. --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./user/*.proto