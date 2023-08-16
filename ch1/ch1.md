#### 1.1 hello, world

Go是一门**编译型语言**，Go语言的工具链将源代码及其依赖转换成计算机的机器指令（译注：静态编译）。Go语言提供的工具都通过一个单独的命令go调用，go命令有一系列子命令。最简单的一个子命令就是run。这个命令编译一个或多个以.go结尾的源文件，链接库文件，并运行最终生成的可执行文件。

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```

```shell
# 一次性实验
go run helloworld.go

# 生成一个名为helloworld的可执行的二进制文件，支持复用运行
go build helloworld.go

# 执行可执行程序
./helloworld
```

执行 go get gopl.io/ch1/helloworld 命令，就会从网上获取代码，并放到对应目录中（需要先安装Git或Hg之类的版本管理工具，并将对应的命令添加到PATH环境变量中。**需要先设置好GOPATH环境变量**，下载的代码会放在$GOPATH/src/gopl.io/ch1/helloworld目录）。

一个**包**由位于**单个目录下的一个或多个.go源代码文件组成**，**目录名字定义包的作用**。**每个源文件都以一条<u>package</u>声明语句开始，这个例子里就是package main，表示<u>该文件属于哪个包</u>**，紧跟着一系列导入（import）的包，之后是存储在这个文件里的程序语句。

**<u>main包</u>**比较特殊。它**定义了一个独立可执行的程序**，而不是一个库。在main里的**main *函数*** 也很特殊，它是**整个程序执行时的入口**（译注：C系语言差不多都这样）。main函数所做的事情就是程序做的。

必须恰当导入需要的包，缺少了必要的包或者导入了不需要的包，程序都无法编译通过。这项严格要求**避免了程序开发过程中引入未使用的包**（译注：**Go语言编译过程没有警告信息**，争议特性之一）。

**import声明必须跟在文件的package声明之后**。随后，则是组成程序的函数、变量、常量、类型的声明语句（分别由关键字func、var、const、type定义）。这些内容的声明顺序并不重要（译注：最好还是定一下规范）。

**Go语言不需要在语句或者声明的末尾添加分号**，除非一行上有多条语句。

函数的左括号{必须和func函数声明在同一行上，且位于末尾，不能独占一行，而在表达式x + y中，可在+后换行，不能在+前换行（译注：以+结尾的话不会被插入分号分隔符，但是以x结尾的话则会被分号分隔符，从而导致编译错误）。

**Go语言在代码格式上采取了很强硬的态度**。**gofmt工具把代码格式化为标准格式**（译注：**这个格式化工具没有任何可以调整代码格式的参数**，Go语言就是这么任性），并且go工具中的fmt子命令会对指定包，否则默认为当前目录中所有.go源文件应用gofmt命令。

**很多文本编辑器都可以配置为保存文件时自动执行gofmt，这样你的源代码总会被恰当地格式化**。



#### 1.2 命令行参数

**os包以跨平台的方式，提供了一些与操作系统交互的函数和变量**。程序的命令行参数可从os包的Args变量获取；os包外部使用os.Args访问该变量。

**os.Args**变量是一个**字符串（string）的*切片*（slice）**（译注：slice和Python语言中的切片类似，是一个简版的动态数组。和大多数编程语言类似，**区间索引**时，Go言里也采用**左闭右开**形式，即，区间**包括第一个索引元素**，**不包括最后一个**，因为这样可以简化逻辑。

os.Args的第一个元素：**os.Args[0]**，是**命令本身的名字**。s[m:n]形式的切片表达式，产生从第m个元素到第n-1个元素的切片。**如果省略切片表达式的m或n，会默认传入0或len(s)**。

```go
// echo1 输出其命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
  
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
```

对string类型，+运算符连接字符串（译注：和C++或者js是一样的）。

符号:=是***短变量声明***（short variable declaration）的一部分，这是定义一个或多个变量并根据它们的初始值为这些变量赋予适当类型的语句。

**自增语句i++给i加1；这和i += 1以及i = i + 1都是等价的**。对应的还有i--给i减1。它们是语句，而不像C系的其它语言那样是表达式。所以j = i++非法，而且++和--都只能放在变量名后面，因此--i也非法（只有后++，没有前++）。

```go
for initialization; condition; post {
    // 零个或多个语句
}
```

**大括号强制要求，左大括号必须和*post*语句在同一行**。*initialization*语句是可选的，在循环开始前执行。***initalization***如果存在，必须是一条***简单语句***（simple statement），即，**短变量声明、自增语句、赋值语句或函数调用**。**condition是一个布尔表达式**（boolean expression），其值在每次循环迭代开始时计算。

for循环的这三个部分每个都可以省略，如果省略initialization和post，分号也可以省略：

```go
// 传统的while循环
for condition {
    // ...
}
```

如果连condition也省略了，像下面这样：

```go
// 传统的无限循环
for {
    // ...
}
```

for循环的另一种形式，在某种数据类型的区间（range）上遍历，如字符串或切片。

```go
// echo2输出其命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
```

每次循环迭代，**range**产生一对值；**索引**以及**在该索引处的元素值**。**Go语言不允许使用无用的局部变量（local variables），因为这会导致编译错误**。

Go语言中这种情况的解决方法是用**空标识符**（blank identifier），即_（也就是下划线）。**空标识符可用于在任何语法需要变量名但程序逻辑不需要的时候**（如：在循环里）丢弃不需要的循环索引，并保留元素值。

**每次循环迭代字符串s的内容都会更新**。+=连接原字符串、空格和下个参数，产生新字符串，并把它赋值给s。**s原来的内容已经不再使用，将在适当时机对它进行垃圾回收**。

如果连接涉及的数据量很大，这种方式代价高昂。一种简单且高效的解决方案是使用strings包的Join函数：

```go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
```

##### 练习 1.1

```go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
```

##### 练习 1.2

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args {
		fmt.Println(idx, ":", arg)
	}
}
```



#### 1.3 查找重复的行

```go
// Dup1 输出标准输入中出现次数大于1的行，前面是次数
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() { 
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        counts[input.Text()]++
    }
  // 注意: 忽略input.Err()中可能出现得错误
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
```

```shell
# 注意在运行的时候通过ctrl+d发送EOF
```

**map**存储了键/值（key/value）的集合，对集合元素，提供常数时间的存、取或测试操作。键可以是任意类型，只要其值能用==运算符比较，最常见的例子是字符串；值则可以是任意类型。

**map的迭代顺序并不确定，从实践来看，该顺序随机，每次运行都会变化**。这种设计是有意为之的，因为能防止程序依赖特定遍历顺序，而这是无法保证的。（译注：具体可以参见这里http://stackoverflow.com/questions/11853396/google-go-lang-assignment-order）

继续来看**bufio**包，它使处理输入和输出方便又高效。**Scanner类型是该包最有用的特性之一，它读取输入并将其拆成行或单词**；通常是处理行形式的输入最简单的方法。

程序使用短变量声明创建bufio.Scanner类型的变量input。

```go
input := bufio.NewScanner(os.Stdin)
```

该变量从程序的标准输入中读取内容。每次调用input.Scan()，即读入下一行，并移除行末的换行符；读取的内容可以调用input.Text()得到。Scan函数在读到一行时返回true，不再有输入时返回false。

Print“转换字符”，基本同C语言，新增的转移字符如下：

```shell
%q          带双引号的字符串"abc"或带单引号的字符'c'
%v          变量的自然形式（natural format）
%T          变量的类型
%%          字面上的百分号标志（无操作数）
```

默认情况下，Printf不会换行。按照惯例，**以字母f结尾的格式化函数**，如log.Printf和fmt.Errorf，都采用fmt.Printf的格式化准则。而**以ln结尾的格式化函数，则遵循Println的方式**，**以跟%v差不多的方式格式化参数，并在最后添加一个换行符**。（译注：后缀f指format，ln指line。）

```go
// Dup2 打印输入中多次出现的行的个数和文本，它从stdin或指定的文件列表读取
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// 注意: 忽略input.Err()中可能出现得错误
}
```

**os.Open函数返回两个值**。第一个值是**被打开的文件**(*os.File），其后被Scanner读取。os.Open返回的**第二个值是内置error类型的值**。如果err等于内置值nil（译注：相当于其它语言里的NULL），那么文件被成功打开。读取文件，直到文件结束，然后调用Close关闭该文件，并释放占用的所有资源。相反的话，如果err的值不是nil，说明打开文件时出错了。这种情况下，错误值描述了所遇到的问题。我们的错误处理非常简单，**只是使用Fprintf与表示任意类型默认格式值的%v向标准错误流打印一条信息**，然后dup继续处理下一个文件；continue语句直接跳到for循环的下个迭代开始执行。

注意countLines函数在其声明前被调用。**函数和包级别的变量（package-level entities）可以任意顺序声明，并不影响其被调用**。（译注：最好还是遵循一定的规范）

**map是一个由make函数创建的数据结构的<u>引用</u>**。**map作为参数传递给某函数时，该函数接收这个引用的一份拷贝（copy，或译为副本），被调用函数对map底层数据结构的任何修改，调用者函数都可以通过持有的map引用看到**。在我们的例子中，countLines函数向counts插入的值，也会被main函数看到。（译注：类似于C++里的引用传递，实际上指针是另一个指针了，但内部存的值指向同一块内存）

引入了ReadFile函数（来自于io/ioutil包），其读取指定文件的全部内容，strings.Split函数把字符串分割成子串的切片。（Split的作用与前文提到的strings.Join相反。）

dup的前两个版本以"流”模式读取输入，并根据需要拆分成多个行。理论上，这些程序可以处理任意数量的输入数据。还有另一个方法，就是一口气把全部输入数据读到内存中，一次分割为多行，然后处理它们。

```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
```

**ReadFile函数返回一个字节切片（byte slice），必须把它转换为string，才能用strings.Split分割**。

**实现上，bufio.Scanner、ioutil.ReadFile和ioutil.WriteFile都使用*os.File的Read和Write方法，但是，大多数程序员很少需要直接调用那些低级（lower-level）函数**。**高级（higher-level）函数，像bufio和io/ioutil包中所提供的那些，用起来要容易点**。

##### 练习 1.4

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

type LineInfo struct {
	count int
	files []string
}

func main() {
	lineMap := make(map[string]*LineInfo)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, lineMap)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, lineMap)
			f.Close()
		}
	}
	for line, n := range lineMap {
		if n.count > 1 {
			fmt.Printf("%d\t%s\n", n.count, line)
		}
	}
}

func countLines(f *os.File, lineMap map[string]*LineInfo) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if _, exist := lineMap[line]; exist {
			lineMap[line].count++
			lineMap[line].files = append(lineMap[line].files, f.Name())
		} else {
			lineMap[line] = &LineInfo{1, []string{f.Name()}}
		}
	}
	// 注意: 忽略input.Err()中可能出现得错误
}
```



#### 1.4 GIF动画

```go
// lissajous 产生随机丽萨茹图形的GIF动画
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // 画板中的第一种颜色
	blackIndex = 1 // 画板中的第二种颜色
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 完整的x振荡器变化次数
		res     = 0.001 // 角度分辨率
		size    = 100   // 图像画布大小
		nframes = 64    // 动画的帧数
		delay   = 8     // 以10ms为单位的延迟，总计80ms延迟
	)

	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 初相位差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // 创建画板
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay) // 指定动画的延迟
		anim.Image = append(anim.Image, img)   // 追加帧
	}
	gif.EncodeAll(out, &anim) // 注意：忽略编码错误
}
```

这个程序里的常量声明给出了一系列的常量值，常量是指在程序编译后运行时始终都不会变化的值，比如圈数、帧数、延迟值。**常量声明和变量声明一般都会出现在包级别**，所以**这些常量在整个包中都是可以共享的**，或者你也可以**把常量声明定义在函数体内部，那么这种常量就只能在函数体内用**。目前常量声明的值必须是一个数字值、字符串或者一个固定的boolean值。

##### 练习 1.5

```go
// lissajous 产生随机丽萨茹图形的GIF动画
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	blackIndex = 0 // 画板中的第一种颜色
	greenIndex = 1 // 画板中的第二种颜色
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 完整的x振荡器变化次数
		res     = 0.001 // 角度分辨率
		size    = 100   // 图像画布大小
		nframes = 64    // 动画的帧数
		delay   = 8     // 以10ms为单位的延迟，总计80ms延迟
	)

	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 初相位差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // 创建画板
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay) // 指定动画的延迟
		anim.Image = append(anim.Image, img)   // 追加帧
	}
	gif.EncodeAll(out, &anim) // 注意：忽略编码错误
}
```

##### 练习 1.6

```go
// lissajous 产生随机丽萨茹图形的GIF动画
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
	color.RGBA{0xff, 0x00, 0xff, 0xff},
	color.RGBA{0xAA, 0x55, 0x00, 0xff},
	color.RGBA{0x00, 0xAA, 0x55, 0xff},
	color.RGBA{0xAA, 0x00, 0x55, 0xff},
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 完整的x振荡器变化次数
		res     = 0.001 // 角度分辨率
		size    = 100   // 图像画布大小
		nframes = 64    // 动画的帧数
		delay   = 8     // 以10ms为单位的延迟，总计80ms延迟
	)

	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 初相位差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // 创建画板
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(rand.Int()%len(palette)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay) // 指定动画的延迟
		anim.Image = append(anim.Image, img)   // 追加帧
	}
	gif.EncodeAll(out, &anim) // 注意：忽略编码错误
}
```



#### 1.5 获取URL

为了最简单地展示基于HTTP获取信息的方式，下面给出一个示例程序fetch，这个程序将获取对应的url，并将其源文本打印出来；这个例子的灵感来源于curl工具（译注：unix下的一个用来发http请求的工具，具体可以man curl）。

```go
// Fetch 输出从URL获取的内容
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
```

net/http和io/ioutil包，**http.Get**函数是创建HTTP请求的函数，如果获取过程没有出错，那么会在resp这个结构体中得到访问的请求结果。**resp的Body字段包括一个可读的服务器响应流**。ioutil.ReadAll函数从response中读取到全部内容；将其结果保存在变量b中。

##### 练习 1.7

```go
// Fetch 输出从URL获取的内容
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}
```

##### 练习 1.8

```go
// Fetch 输出从URL获取的内容
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
```

##### 练习 1.9

```go
// Fetch 输出从URL获取的内容
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%v\n", resp.Status)
	}
}
```



#### 1.6 并发获取多个URL

Go语言最有意思并且最新奇的特性就是对并发编程的支持。

```go
// Fetchall 并发获取url并报告它们的时间和大小
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
        go fetch(url, ch) // 启动一个goroutine
    }
    for range os.Args[1:] {
        fmt.Println(<-ch) // 从通道ch接收
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err) // 发送到通道ch
        return
    }
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close() // 不要泄露资源
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
```

**goroutine**是一种函数的**并发执行方式**，而**channel**是用来**在goroutine之间进行参数传递**。main函数本身也运行在一个goroutine中，而**go function**则表示**创建一个新的goroutine**，并在这个新的goroutine中执行这个函数。

main函数中用make函数创建了一个传递string类型参数的channel，对每一个命令行参数，我们都用go这个关键字来创建一个goroutine，并且让函数在这个goroutine异步执行http.Get方法。这个程序里的io.Copy会把响应的Body内容拷贝到**ioutil.Discard**输出流中（译注：可以把这个变量看作一个**垃圾桶**，可以向里面**写一些不需要的数据**），因为我们需要这个方法返回的字节数，但是又不想要其内容。

**当一个goroutine尝试在一个channel上做send或者receive操作时，这个goroutine会阻塞在调用处，直到另一个goroutine从这个channel里接收或者写入值，这样两个goroutine才会继续执行channel操作之后的逻辑**。在这个例子中，每一个fetch函数在执行时都会往channel里发送一个值（ch <- expression），主函数负责接收这些值（<-ch）。



#### 1.7 一个Web服务器

Go语言的内置库使得写一个类似fetch的web服务器变得异常地简单。比如用户访问的是 http://localhost:8000/hello ，那么响应是URL.Path = "/hello"。

```go
// Server1 是一个迷你回声服务器
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // 回声请求调用处理程序
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

// 处理程序回显请求URL r的路径部分
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
```

main函数将一个处理函数和以/开头的URL链接在一起，**代表所有的URL使用这个函数处理**，然后启动服务器监听进人8000端口处的请求。一个**请求由一个http.Request类型的结构体表示**，它包含很多关联的域，其中一个是所请求的URL。当一个请求到达时，它被转交给处理函数，并从请求的URL 中提取路径部分（/hello），使用fmt.Printf格式化，然后作为响应发送回去。

**<u>在后台启动服务器</u>，在Mac OS X或者Linux上，<u>在命令行后添加一个&符号</u>**。

```shell
./server1 &
```

```go
// server2 是一个迷你的回声和计数器服务器
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

// 处理程序回显请求的URL的路径部分
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter 回显目前为止调用的次数
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
```

这个服务器有两个处理函数，通过请求的URL来决定哪一个被调用：请求/count调用counter，其他的调用handler。**以/结尾的处理模式匹配所有含有这个前缀的URL**。**在后台，对于每个传入的请求，服务器在不同的goroutine中运行该处理函数，这样它可以同时处理多个请求**。然而，如果两个并发的请求试图同时更新计数值count，它可能会不一致地增加，程序会产生一个严重的竞态bug。为避免该问题，必须确保最多只有一个goroutine在同一时间访问变量，这正是mu.Lock()和mu.Unlock()语句的作用。

```go
// server2 是一个迷你的回声和计数器服务器
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

// 处理程序回显请求的URL的路径部分
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	// 解析http请求中传入的参数
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter 回显目前为止调用的次数
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
```

这些程序中，我们看到了作为**输出流**的三种非常不同的类型。fetch程序复制HTTP响应到文件 **os.Stdout**，像lissajous一样； fetchall程序通过将响应复制到**ioutil.Discard**中进行**丢弃**（在统计其长度时）； Web服务器使用**fmt.Fprintf** 通过写人**http.Responsewriter**来**让浏览器显示**。
尽管**三种类型细节不同**，**但都满足一个通用的接口**（**<u>interface</u>**），该**接口允许它们按需使用任何一种输出流**。该接口（称为io.writer）。

##### 练习 1.12

```go
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // 画板中的第一种颜色
	blackIndex = 1 // 画板中的第二种颜色
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))

	return
}

func lissajous(out io.Writer, r *http.Request) {
	cycles := 5 // 完整的x振荡器变化次数

	const (
		res     = 0.001 // 角度分辨率
		size    = 100   // 图像画布大小
		nframes = 64    // 动画的帧数
		delay   = 8     // 以10ms为单位的延迟，总计80ms延迟
	)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		if k == "cycles" && v[0] != "" {
			var err error
			cycles, err = strconv.Atoi(v[0])

			if err != nil {
				log.Print(err)
				fmt.Fprintf(out, "cycles parameter wrong")

				return
			}
		}
	}

	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 初相位差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // 创建画板
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay) // 指定动画的延迟
		anim.Image = append(anim.Image, img)   // 追加帧
	}
	gif.EncodeAll(out, &anim) // 注意：忽略编码错误
}
```



#### 1.8 其他内容

**switch语句不需要操作数**，它就像一个case语句列表，每条case语句都是一个布尔表达式：

```go
func Signum(x int) int {
    switch {
    case x > 0:
        return +1
    default:
        return 0
    case x < 0:
        return -1
    }
}
```

这种形式称为**无标签（tagless）选择**，它**等价于switch true**。

与for和if语句类似， **switch可以包含一个可选的简单语句**：一个短变量声明，一个递增或赋值语句，或者一个函数调用，用来在判断条件前设置一个值。

**break可以打断<u>for</u>、<u>switch</u> 或<u>select</u>的最内层调用**。

**type**声明给已有类型命名。因为**结构体**类型通常很长，所以它们基本上都**独立命名**。一个熟悉的例子是定义一个2D图形系统的Point类型：

```go
type Point struct {
    X, Y int
}
var p Point
```

**方法和接口**：一个**关联了命名类型的函数**称为**<u>方法</u>**。Go里面的方法可以关联到几乎所有的命名类型。**<u>接口</u>**可以**用相同的方式处理不同的具体类型的<u>抽象类型</u>**，**它基于这些类型所包含的方法**，**而不是类型的描述或实现**。

**包**： Go自带一个可扩展并且实用的标准库，Go社区创建和共享了更多的库。

**注释**：**在声明任何函数前，写一段注释来说明它的行为是一个好的风格**。这个约定很重要，**因为它们可以被go doc 和godoc工具定位和作为文档显示**。

```shell
go doc io.writer
```

```shell
package io // import "io"

type Writer interface {
	Write(p []byte) (n int, err error)
}
    Writer is the interface that wraps the basic Write method.

    Write writes len(p) bytes from p to the underlying data stream. It returns
    the number of bytes written from p (0 <= n <= len(p)) and any error
    encountered that caused the write to stop early. Write must return a non-nil
    error if it returns n < len(p). Write must not modify the slice data, even
    temporarily.

    Implementations must not retain p.

var Discard Writer = discard{}
func MultiWriter(writers ...Writer) Writer
```

```shell
go doc http.HandlerFunc
```

```shell
package http // import "net/http"

type HandlerFunc func(ResponseWriter, *Request)
    The HandlerFunc type is an adapter to allow the use of ordinary functions as
    HTTP handlers. If f is a function with the appropriate signature,
    HandlerFunc(f) is a Handler that calls f.

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
```

对于跨越多行的注释，可以使用类似其他语言中的/*...*/注释。这样可以避免在文件的开始有一大块说明文本时每一行都有//。**在注释内部，//和/*没有特殊的含义，所以注释不能嵌套**。
