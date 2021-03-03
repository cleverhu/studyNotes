# 第一章:go和操作系统
# 1.golang的优点
	1.现代化语言，读写代码简单，是被有经验的程序员所创造。
	2.Go编译器打印实际的错误和警告帮助解决实际问题，简单来说Go编译器帮助你，而不是使生活更糟糕，打印无意义的输出。
	3.Go的代码是可移植的，特别是在unix系统之间。
	4.Go支持过程，并发和分布式编程。
	5.Go有垃圾回收机制，不需要考虑内存的申请和释放。
	6.Go没有预处理器，但可以进行高速编译。 因此，Go也可以用作脚本语言。
	7.标准的Go库提供了许多简化开发人员工作的软件包。 另外，在标准go库中找到的功能是由开发Go的人员预先测试和调试的，这意味着在大多数情况下，它们没有错误。
	8.Go默认情况下使用静态链接，这意味着生成的二进制文件可以轻松转移到具有相同操作系统的其他计算机上。 因此，一旦成功编译了Go程序并生成了可执行文件，您就不必担心库，依赖项和不同的库版本。
	9.Go支持Unicode，这意味着您不需要任何其他代码即可打印多种人类语言的字符。
	10.Go使概念保持正交，因为一些正交特征比许多重叠特征更有效。
# 2.golang的缺点
	1.go没有直接支持面向对象的编程，这对习惯于以面向对象的方式编写代码的程序员来说可能是个问题。 不过，您可以在Go中使用组合来模仿继承。
	2.对于某些人来说，Go永远不会取代C。
	3.对于系统编程，C仍然比任何其他编程语言都要快，这主要是因为UNIX是用C编写的。

# 3.godoc单元
	安装手册：https://blog.csdn.net/qq_43267773/article/details/109162333
	linux:man ps
	golang:godoc fmt/go doc fmt.Printf	...
	godoc -http=:8001 创建一个网站，端口是8001，可以查看所有go文档http://localhost:8001/pkg/

# 4.文件命名
	source_file.go或者aSourceFile.go，无论选择哪种，风格保持一致

# 5.执行
	go build生成二进制可执行文件
	go run执行比go build简单，go run执行完毕会把二进制文件给删除

# 6.两条规则
	1.不能导入不用到的包，如果要只需要用到init，包前加_
	2.{}问题，{要跟在函数名()后面

# 7.下载go packages
	go get -v github.com/mactsouk/go/simpleGitHub

	go clean -i -v -x github.com/mactsouk/go/simpleGitHub
	等效于
	rm -rf ~/go/src/github.com/mactsouk/go/simpleGitHub

# 8.UNIX stdin, stdout, and stderr
	默认情况下，所有UNIX系统都支持三种特殊和标准文件名：/dev/stdin，/dev/stdout和/dev/stderr，
    也可以分别使用文件描述符0、1和2访问它们。 这三个文件描述符也分别称为标准输入，标准输出和标准错误。 
    另外，文件描述符0可以在macOS机器上以/dev/fd/0访问，在Debian Linux机器上可以以/dev/fd/0和/dev/pts/0访问。
	Go使用os.Stdin访问标准输入，使用os.Stdout访问标准输出，使用os.Stderr访问标准错误。 
    尽管您仍然可以使用/dev/stdin，/dev/stdout和/dev/ stderr或相关的文件描述符值来访问相同的设备，
    但坚持使用os.Stdin，os会更好，更安全且更可移植。 Stdout和os.Stderr由Go提供。

# 9.打印输出
	fmt.Println(v1, v2) 等效 fmt.Print(v1, " ", v2, "\n")
	fmt.Printf()格式化打印
	fmt.Sprintln(), fmt.Sprint(),fmt.Sprintf()这些函数是基于给定的格式创建字符串
	fmt.Fprintln(), fmt.Fprint(), and fmt.Fprintf()被用来写文件通过使用io.Writer

# 10.标准输出
	io.WriteString(agr1,arg2),io.WriteString(os.Stdout, "\n") 函数工作方式和fmt.Print()等效
	第一个参数是需要写入的文件，第二个是字符串变量
	io.Writer(agr1,agr2)第二个参数需要一个字节数组([]byte)或者一个切片

# 11.得到用户输入
	:= 短的赋值语句
	m:=123 等效于 var m int = 123

	func main() {
	 
	    var f *os.File 
	    f = os.Stdin 
	    defer f.Close() 
	 
	    scanner := bufio.NewScanner(f) 
	    for scanner.Scan() { 
	        fmt.Println(">", scanner.Text()) 
	    } 
	} 
	Ctrl + D终止读取

# 12.错误输出
	 io.WriteString(os.Stderr, myString) 和标准输出差不多

# 13.log函数
	log.Panic()输出包括其他底层信息，这些信息有望帮助您解决Go代码中发生的困难情况
	与log.Fatal（）函数类似，使用log.Panic（）函数会将条目添加到适当的日志文件中，并立即终止Go程序。

# 14.go里面错误处理
	错误和错误处理是两个非常重要的Go主题。 Go非常喜欢错误消息，以至于它有一个单独的错误数据类型，称为error。 
	这也意味着，如果您发现Go所提供的功能不足，则可以轻松创建自己的错误消息。
	请注意，出现错误情况是一回事，但是决定如何对错误情况做出反应是完全不同的事情。 
	简而言之，并非所有错误条件都相同，这意味着某些错误条件可能要求您立即停止执行程序，
    而其他错误情况则可能需要打印警告消息，以供用户在继续执行程序时查看。 
	该程序。 开发人员应使用常识并决定如何处理程序可能获得的每个错误值。
	在许多情况下，开发自己的Go应用程序时最终可能不得不处理新的错误情况。 错误数据类型在这里可以帮助您定义自己的错误。

	创建属于自己的error变量=>errors标准库里的New() errors.New()需要一个string类型作为参数，
    如果一个函数应该返回error变量，但是没有error去报告，应该返回nil
	fmt.Errorf(format,args)底层也是errors.New(),格式化一个字符串然后，调用errors.New()
	err.Error()将一个error变量转化为string类型

# 15.Dockerfile
	FROM golang:alpine 
 
	RUN mkdir /files 
	COPY hw.go /files 
	WORKDIR /files 
	 
	RUN go build -o /files/hw hw.go 
	ENTRYPOINT ["/files/hw"] 


	docker build -t go_hw:v1 .

	docker run go_hw:v1

	docker login
	docker tag go_hw:v1 "mactsouk/go_hw:v1"
	docker push "mactsouk/go_hw:v1"


# 第二章:理解go的本质

# 1.go编译器
	go tool compile test.go
	test.go->test.o
	go tool compile -pack test.go
	test.go->test.a

	file test.a
	结果:
	current ar archive

	file test.o
	current ar archive

	archive file 存档文件是包含一个或多个文件的二进制文件，主要是用于将多个文件分组为一个文件。 这些格式之一是ar，这是被Go使用。

	ar t test.a
	结果
	__.PKGDEF
	_go_.o

	go tool compile -S test.go更详细的信息

# 2.垃圾回收
	垃圾回收是释放未使用的内存空间的过程。 换句话说，垃圾收集器查看哪些对象不在范围内并且无法再被引用，并释放它们消耗的内存空间。 
	此过程在Go程序运行时（而不是在程序执行之前或之后）以并行方式发生。

	三色算法
	三色标级和清除算法

	三色标记和清除算法背后的主要原理是，它根据堆的颜色将堆的对象划分为三个不同的集合，这些颜色由算法分配。 
	现在是时候讨论每种颜色集的含义了。 保证黑色集合的对象没有指向白色集合的任何对象的指针。
	但是，白色集的对象可以具有指向黑色集的对象的指针，因为这对垃圾收集器的操作没有影响。 
	灰色集合的对象可能具有被指白色集合的某些对象指向的指针。 
	最后，白色集合的对象是垃圾收集的候选对象。

	请注意，没有对象可以直接从黑色集转到白色集，这允许算法进行操作并能够清除白色集上的对象。 
	此外，黑色组的任何对象都不能直接指向白色组的对象。
	因此，当垃圾回收开始时，所有对象都是白色的，垃圾回收器访问所有根对象，并将它们着色为灰色。 
	根是可以由应用程序直接访问的对象，其中包括全局变量和堆栈上的其他内容。 这些对象主要取决于特定程序的Go代码。
	此后，垃圾收集器选择一个灰色对象，使其变为黑色，然后开始查看该对象是否具有指向白色集合中其他对象的指针。 
	这意味着，当扫描灰色对象以寻找指向其他对象的指针时，该对象将变为黑色。 
	如果该扫描发现该特定对象具有一个或多个指向白色对象的指针，则会将该白色对象放入灰色集中。 
	只要灰色集中存在对象，此过程就会持续进行。 
	之后，白色集中的对象将无法访问，并且它们的存储空间可以被重用。 
	因此，在这一点上，白色集合的元素被认为是垃圾收集的。

	请注意，如果在垃圾回收周期的某个时刻灰色集合的某个对象变得不可访问，
	则不会在该垃圾回收周期中将其收集，而是在下一个垃圾回收中将其收集！ 尽管这不是最佳情况，但还不错。
















