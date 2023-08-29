#### 2.1 名称

Go 有25个像if和switch这样的关键字，只能用在语法允许的地方，它们不能作为名称：

```go
break      default       func     interface   select
case       defer         go       map         struct
chan       else          goto     package     switch
const      fallthrough   if       range       type
continue   for           import   return      var
```

 另外，还有三十几个内置的预声明的常量、类型和函数：

```go
内建常量: true false iota nil

内建类型:	int int8 int16 int32 int64
         uint uint8 uint16 uint32 uint64 uintptr
         float32 float64 complex128 complex64
         bool byte rune string error

内建函数: make len cap new append copy close delete
        complex real imag
        panic recover
```

这些名称不是预留的，可以在声明中使用它们。我们将在很多地方看到对其中的名称进行重声明，但是要知道这有冲突的风险。

如果一个实体在函数中声明，它只在函数局部有效。如果声明在函数外，它将对包里面的所有源文件可见。实体**第一个字母的大小写决定其可见性是否跨包**。如果**名称以大写字母的开头，它是导出的，意味着它对包外是可见和可访问的**，可以被自己包之外的其他程序所引用，像fmt包中的Printf。**包名本身总是由小写字母组成**。

名称本身没有长度限制，但是习惯以及Go的编程风格倾向于使用短名称，特别是作用域较小的局部变量，你更喜欢看到一个变量叫i而不是theLoopIndex。**通常，名称的作用域越大，就使用越长且更有意义的名称**。

**Go程序员使用“驼峰式”的风格**——更喜欢使用大写字母而不是下划线。



#### 2.2 声明

有4个主要的声明：**变量**(var)、**常量**(const)、**类型**(type)和**函数**(func)。

每一个文件以package声明开头，表明文件属于哪个包。**package声明**后面是**import 声明**，然后是**包级**别的**类型**、**变量**、**常量**、**函数**的声明，不区分顺序。

> 良好的编码习惯，应规约好一种次序。

```go
// boiling 输出水的沸点
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// 输出：
	// boiling point = 212°F or 100°C
}
```

常量**boilingF**是一个包级别的声明(**main包**)。**包级别的实体名字不仅对于包含其声明的<u>源文件可见</u>，而且对于<u>同一个包里面的所有源文件</u>都可见**。

如果函数不返回任何内容，返回值列表可以省略。函数的执行从第一个语句开始，直到遇到一个返回语句，或者执行到无返回结果的函数的结尾。

```go
// ftoc 输出两个华氏温度—摄氏温度的转换
package main

import "fmt"

func main() {
	// 局部常量
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
```



#### 2.3 变量

```go
var name type = expression
```

**类型和表达式部分可以省略一个，但是不能都省略**。如果类型省略，它的类型将由初始化表达式决定；如果表达式省略，其**初始值对应于类型的零值**—对于数字是0，对于布尔值是false，对于字符串是""，对于接口和引用类型(slice、指针、map、通道、函数)是nil。对于一个像数组或结构体这样的复合类型，零值是其所有元素或成员的零值。

**零值机制保障所有的变量是良好定义的，Go里面不存在未初始化变量**。

> 区别与C/C++

**包级别的初始化在main开始之前进行(运行前)**，**局部变量初始化和声明一样在函数执行期间进行(运行时)**。



##### 2.3.1 短变量声明

一种称作**短变量声明**的可选形式可以用来**声明**和**初始化**局部变量。

```go
name := expression
```

name 的类型由 expression的类型决定。

因其短小、灵活，故而在局部变量的声明和初始化中主要使用短声明。**var声明通常是为那些跟初始化表达式类型不一致的局部变量保留的**，或者用于**后面才对变量赋值**以及变量初始值不重要的情况。

:=表示声明(加初始化)，而=表示赋值。

短变量声明最少声明一个新变量，否则，代码编译将无法通过。

```go
f, err := os.Open(infile)
// ...
f, err := os.Create(outfile) // 编译错误：没有新的变量
```



##### 2.3.2 指针

**不是所有的值都有地址，但是所有的变量都有**。使用指针，可以在无须知道变量名字的情况下，间接读取或更新变量的值。

**函数返回局部变量的地址是非常安全的**。

> 区别与C/C++，在C/C++中是不安全的

通过调用f产生的局部变量v即使在调用返回后依然存在，指针p依然引用它：

```go
package main

import "fmt"

func f() *int {
	v := 1
	return &v
}

func main() {
	fmt.Println(f() == f())
}
```

```shell
false
```

**指针别名允许我们不用变量的名字来访问变量**。

flag是一个用于做命令行解析的包。

```go
// echo4 输出其命令行参数package main
package main

import (
	"flag"
	"fmt"
	"strings"
)

// flag参数：标志、默认值、help信息，返回值是指针
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	// 需要调用Parse()进行参数解析
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
```



##### 2.3.3 new 函数

表达式**new(T)创建一个未命名的T类型变量**，初始化为T类型的零值，并**返回其地址**(地址类型为*T)。

使用new创建的变量和取其地址的普通局部变量没有什么不同，只是不需要引入(和声明)一个虚拟的名字，通过new(T)就可以直接在表达式中使用。因此new只是语法上的便利，不是一个基础概念。

**new**是一个**预声明的函数**，**不是一个关键字**，所以它可以重定义为另外的其他类型。

```go
// 在delta函数内，内置的new函数是不可用的
func delta(old, new int) int { return new - old }
```

> 这里区别与C++，C++中new是一个关键字。另外，不要用到上面的技巧。



##### 2.3.4 变量的生命周期

**包级别变量的生命周期是整个程序的执行时间**。相反，**局部变量**有一个动态的生命周期：每次执行声明语句时创建一个新的实体，**变量一直生存到它变得不可访问，这时它占用的存储空间被回收**。函数的**参数**和**返回值**也是**局部变量**，它们在其闭包函数被调用的时候创建。

那么**垃圾回收**器如何知道一个变量是否应该被回收？说来话长，基本思路是每一个包级别的变量，以及每一个当前执行函数的局部**变量**，可以**作为追溯该变量的路径的源头**，**通过指针和其他方式的引用可以找到变量**。如果变量的**路径不存在**，那么变量**变得不可访问**，因此它不会影响任何其他的计算过程。

因为变量的生命周期是通过它是否可达来确定的，所以**局部变量可在包含它的循环的一次选代之外继续存活**。**即使包含它的循环已经返回，它的存在还可能延续**。

> 这一点区别与C/C++，C/C++局部变量的生命周期是在其特定的作用域，作用域之外若试图通过指针访问它的话，会出现难以预料的问题。

编译器可以选择使用堆或栈上的空间来分配，令人惊奇的是，这个选择不是基于使用var或new关键字来声明变量。

```go
var global *int

func f() {
    var x int
    x = 1
    global = &x
}

func g() {
    y := new(int)
    *y = 1
}
```

x一定使用**堆空间**，因为**它在f函数返回以后还可以从global变量访问**，尽管它被声明为一个局部变量。这种情况我们说**x从f中逃逸**。相反，当g函数返回时，变量\*y变得**不可访问**，可回收。因为\*y**没有从g中逃逸**，所以编译器可以安全地在**栈**上分配\*y，**即便使用new函数创建它**。任何情况下，逃逸的概念使你不需要额外费心来写正确的代码，但要记住它在性能优化的时候是有好处的，因为**每一次变量逃逸都需要一次额外的内存分配过程**。

> 这一点跟C/C++有很大的区别，行文之间似乎也在试图跟C++对照。

垃圾回收对于写出正确的程序有巨大的帮助，但是免不了**考虑内存的负担**。不需要显式分配和释放内存，但是变量的生命周期是写出高效程序所必需清楚的。例如，在长生命周期对象中**保持短生命周期对象不必要的指针**，特别是在全局变量中，**会阻止垃圾回收器回收短生命周期的对象空间**。



#### 2.4 赋值

##### 2.4.1 多重赋值

从风格上考虑，如果表达式比较复杂，则避免使用多重赋值形式；一系列独立的语句更易读。

表达式(例如一个有多个返回值的函数调用)产生多个值。当在一个赋值语句中使用这样的调用时，左边的变量个数需要和函数的返回值一样多。

**map查询**、**类型断言**或者**通道接收**动作出现在两个结果的赋值语句中，都会产生一个额外的布尔型结果。



##### 2.4.2 可赋值性

程序中很多地方的赋值是隐式的：一个函数调用**隐式地将参数的值赋给对应参数的变量**；一个return语句隐式地将**return操作数赋值给结果变量**。

不管隐式还是显式赋值，如果左边的(变量)和右边的(值)类型相同，它就是合法的。通俗地说，**赋值只有在值对于变量类型是可赋值的时候才合法**。

**<u>可赋值性</u>**规则很简单：类型必须精准匹配， **nil**可以被赋给任何**接口**变量或**引用**类型。



#### 2.5 类型声明

**type**声明**定义一个新的命名类型**，它和某个已有类型使用同样的底层类型。命名类型提供了一种方式来区分底层类型的不同或者不兼容使用，这样它们就不会在无意中混用。

**类型的声明**通常**出现在包级别**，这里命名的类型在整个包中可见，如果**名字是导出的**(开头使用大写字母)，**其他的包也可以访问它**。

```go
// 包tempconv进行摄氏温度和华氏温度的转换计算
package tempconv

import "fmt"

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
    AbsoluteZeroC Celsius = -273.15 // 绝对零度
    FreezingC     Celsius = 0       // 结冰点温度
    BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
```

这个包定义了两个类型—Celsius(摄氏温度)和Fahrenheit (华氏温度)，它们分别对应两种温度计量单位。**即使使用相同的底层类型float64**，它们也不是相同的类型，所以**它们不能使用算术表达式进行比较和合并**。区分这些类型可以防止无意间合并不同计量单位的温度值；从float64 转换为Celsius(t) 或Fahrenheit(t)需要显式类型转换。**Celsius(t)和Fahrenheit(t)是类型转换，而不是函数调用**。它们不会改变值和表达方式，但改变了显式意义。另一方面，函数CToF和FToc用来在两种温度计量单位之间转换，返回不同的数值。

对于每个类型T，都有一个对应的类型转换操作T(x)将值x转换为类型T。

如果两个类型**具有相同的底层类型**或**二者都是指向相同底层类型变量的未命名指针类型**，则二者是**可以相互转换**的。类型转换不改变类型值的表达方式，仅改变类型。

数字类型间的转换，字符串和一些slice类型间的转换是允许的，我们将在下一章详细讨论。这些转换会改变值的表达方式。例如，从**浮点型转化为整型会丢失小数部分**，从**字符串转换成字节([]byte) slice会分配一份字符串数据副本**。任何情况下，运行时的转换不会失败。

命名类型的值**可以与其相同类型的值或者底层类型相同的未命名类型的值相比较**。但是**不同命名类型的值不能直接比较**。

注意最后一种情况。无论名字如何，类型转换Celsius(f)没有改变参数的值，只改变其类型。

```go
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
```

Celsius参数c出现在函数名字前面，名字叫**String的方法关联到Celsius类型**，返回c变量的数字值，后面跟着摄氏温度的符号℃。

很多类型都声明这样一个String方法，在变量通过fmt包作为字符串输出时，它可以控制类型值的显示方式。

```go
package main

import "fmt"

type Celsius float64

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func main() {
	var c Celsius = 100.0

	fmt.Printf("%s\n", c)
}
```



#### 2.6 包和文件

在Go语言中**包的作用**和其他语言中的库或模块作用类似，用于支持**模块化**、**封装**、**编译隔离**和**重用**。一个包的源代码保存在一个或多个以.go结尾的文件中。

**每一个包给它的声明提供独立的命名空间**。例如，在image包中，Decode标识符和unicode/utf16包中的标识符一样，但是关联了不同的函数。为了从包外部引用一个函数，我们必须明确修饰标识符来指明所指的是image.Decode 或 utf16.Decode。

在Go里，通过一条简单的规则来管理标识符是否对外可见：**导出的标识符以大写字母开头**。

**package 声明前面紧挨着的文档注释对整个包进行描述**。习惯上，**应该在开头用一句话对包进行总结性的描述**。每一个包里只有一个文件应该包含该包的文档注释。**扩展的文档注释**通常放在一个文件中，按惯例名字叫作**doc.go**。

##### 练习 2.1

```go
// tempconv 包负责摄氏温度与华氏温度的转换
package tempconv // import "github.com/AntiBargu/gopl/ch2/tempconv"

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
```

```go
package tempconv

// CToF把摄氏温度转换为华氏温度
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK把摄氏温度转换为开尔文温度
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FTOC把华氏温度转换为摄氏温度
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FTOK把华氏温度转换为开尔文温度
func FToK(f Fahrenheit) Kelvin { return Kelvin((f-32)*5/9 + 273.15) }

// KToC把开尔文温度转换为摄氏温度
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToF把开尔文温度转换为华氏温度
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.15)*9/5 + 32) }

```



##### 2.6.1 导入

在Go程序里，每一个**包通过称为导入路径**（import path）的唯一字符串来**标识**。

每个包还有一个**包名**，它以短名字的形式（且不必是唯一的）出现在包的声明中。按约定，包名匹配**导入路径的最后一段**，这样可以方便地预测gopl.io/ch2/tempconv的包名是tempconv。

**导入声明可以给导入的包绑定一个短名字**，用来在整个文件中引用包的内容。上面的import 可以使用修饰标识符来引用gopl.io/ch2/tempconv包里的变量名，如tempconv.CToF。默认这个短名字是包名，在本例中是tempconv，但是**导入声明可以设定一个可选的名字来避免冲突**。

**如果导入一个没有被引用的包，就会触发一个错误**。



##### 2.6.2 包初始化

包的初始化**从初始化包级别的变量开始**，这些变量按照声明顺序初始化，在依赖已解析完毕的情况下，根据**依赖的顺序进行**。

```go
var a = b + c // a 第三个初始化, 为 3
var b = f()   // b 第二个初始化, 为 2, 通过调用 f (依赖c)
var c = 1     // c 第一个初始化, 为 1

func f() int { return c + 1 }
```

包的初始化按照在程序中导入的顺序来进行，**依赖顺序优先**，**每次初始化一个包**。因此，如果包p导入了包q，可以确保q在p之前已完全初始化。**初始化过程是自下向上的，main包最后初始化**。在这种方式下，在程序的main函数开始执行前，所有的包已初始化完毕。

```go
package popcount

// pc[i] 是i的种群统计
var pc [256]byte

// 计算0-255（1个字节）所有值的置位数
func init() {
    // for i, _ := range pc {
    // 简要写法
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

// PopCount 返回x的种群统计（置位的个数）
func PopCount(x uint64) int {
    return int(pc[byte(x>>(0*8))] +
        pc[byte(x>>(1*8))] +
        pc[byte(x>>(2*8))] +
        pc[byte(x>>(3*8))] +
        pc[byte(x>>(4*8))] +
        pc[byte(x>>(5*8))] +
        pc[byte(x>>(6*8))] +
        pc[byte(x>>(7*8))])
}
```

##### 练习 2.3

```go
package popcount1 // import "github.com/AntiBargu/gopl/ch2/popcount1"

// pc[i] 是i的种群统计
var pc [256]byte

// 计算0-255（1个字节）所有值的置位数
func init() {
	// for i, _ := range pc {
	// 简要写法
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount 返回x的种群统计（置位的个数）
func PopCount(x uint64) int {
	var rslt int

	for i := 0; i < 8; i++ {
		rslt += int(pc[byte(x>>(i*8))])
	}

	return rslt
}
```

##### 练习 2.4

```go
package popcount2 // import "github.com/AntiBargu/gopl/ch2/popcount2"

// PopCount 返回x的种群统计（置位的个数）
func PopCount(x uint64) int {
	var rslt int

	for i := 0; i < 64; i++ {
		if (x & (1 << i)) != 0 {
			rslt++
		}
	}

	return rslt
}
```

##### 练习 2.5

```go
package popcount3 // import "github.com/AntiBargu/gopl/ch2/popcount3"

// PopCount 返回x的种群统计（置位的个数）
func PopCount(x uint64) int {
	var rslt int

	for x != 0 {
		rslt++
		x = x & (x - 1)
	}

	return rslt
}
```



#### 2.7 作用域

**声明**将**名字和程序实体关联起来**，如一个**函数**或一个**变量**。**声明的作用域**是指**用到**声明时所声明**名字的源代码段**。
不要将作用域和生命周期混淆。**声明的作用域是声明在程序文本中出现的区域，它是一个编译时属性**。变量的**生命周期是变量在程序执行期间能被程序的其他部分所引用的起止时间，它是一个运行时属性**。

**语法块（block）是由大括号围起来的一个语句序列**，比如一个循环体或函数体。在语法块内部声明的变量对块外部不可见。块把声明包围起来，并且决定了它的可见性。我们可以把块的概念推广到其他没有显式包含在大括号中的声明代码，将其统称为词法块。

**当编译器遇到一个名字的引用时**，将**从最内层的封闭词法块到全局块寻找其声明**。如果没有找到，它会报“undeclared name”错误；**如果在内层和外层块都存在这个声明，内层的将先被找到**。这种情况下，**内层声明将覆盖外部声明**。

> 这一点同C/C++

