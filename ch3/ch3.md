计算机**底层全是位**，而实际操作则是基于大小固定的单元中的数值，称为字（word），这些值**可解释为整数、浮点数、位集（bitset）或内存地址等**，进面构成更大的聚合体，以表示数据包、像素、文件、诗集，以及其他种种。

Go的数据类型分四大类：**基础类型**（basic type）、**聚合类型**（aggregate type）、**引用类型** （reference type）和**接口类型**（interface type）。**引用**是一大分类，其中包含多种不同类型，如指针（pointer，见2.3.2节），slice（见4.2节），map（见4.3节），函数（function，见第5章），以及通道（channel，见第8章）。它们的共同点是全都**间接指向程序变量或状态**，于是操作所引用数据的**效果**就会**遍及该数据的全部引用**。



#### 3.1 整数

有符号整数分四种大小：8位、16位、32位、64位，用int8、int16、int32、int64表示，对应的无符号整数是uint8、uint16、uint32、uint64。**int**是目前使用最广泛的数值类型，都是32位成64位，但**不能认为它们一定就是32位，或一定就是64位**；**即使在同样的硬件平台上，不同的编译器可能选用不同的大小**。

**rune**类型是**int32**类型的同义词，常常用于指明一个值是**Unicode码点**（code point）。 这两个名称可互换使用。同样，**byte**类型是**uint8**类型的同义词，强调一个值是原始数据，而非量值。

还有一种无符号整数**uintptr**，其大小并不明确，但足以完整存放指针。**uintptr类型仅仅用于底层编程**，例如在Go程序与C程序库或操作系统的接口界面。

有符号整数以补码表示，保留最高位作为符号位。

**取模运算符%仅能用于整数**。取模运算符%的行为因编程语言而异。就Go而言，取模余数的正负号总是与被除数一致，于是-5％3和-5％-3都得-2。

全部基本类型的值（布尔值、数值、字符申）都可以比较，这意味着两个相同类型的值可用＝=和!＝运算符比较。

```go
&^ // 位清空(AND NOT)
```

表达式z ＝ x &^ y中，若y的某位是1，则z的对应位等于0；否则，它就等于x的对应位。

Printf用谓词％b以二进制形式输出数值，副词08在这个输出结果前被零，补够8位。

```go
package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5

	fmt.Printf("%08b\n", x)
}
```

```shell
00100010
```

**左移以0填补右边空位**，无符号整数右移同样以0填补左边空位，但**有符号数的右移操作是按符号位的值填补空位**。因此，请注意，**如果将整数以位模式处理，须使用无符号整型**。

**无符号整数**往往只用于**位运算符**和特定算术运算符，如**实现位集**时，**解析二进制格式的文件**，或**散列和加密**。一般而言，**无符号整数极少用于表示非负值**。

浮点型转成整型，会舍弃小数部分，趋零截尾（正值向下取整，负值向上取整）。

源码中的整数都能写成常见的十进制数；也能写成八进制数，以0开头，如0666；还能写成十六进制数，以0x或0X开头，如0xdeadbeef。十六进制的数字（或字母）大小写皆可。当前，八进制数似乎仅有一种用途——表示POSIX文件系统的权限——而十六进制数广泛用于强调其位模式，而非数值大小。

```go
package main

import "fmt"

func main() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
}
```

```shell
438 666 0666
3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
```

注意fmt的两个技巧。通常Printf的格式化字符串含有多个%谓词，这要求提供相同数目的操作数，而**％后的副词［1］告知Printf重复使用第一个操作数**。其次，%o、%x或%x之前的**副词\#告知Printf输出相应的前缀0、0x或0x**。



#### 3.2 浮点数

算术特性**遵从IEEE 754标准**。常量math.MaxFloat32是float32的最大值，大约为3.4e38，面math.MaxFloat64则大约为1.8e308。

绝大多数情况下，应**优先选用float64**，因为除非格外小心，否则float32的运算会迅速累积误差。非常小或非常大的数字最好使用科学记数法表示，此方法在数量级指数前写字母e或E。

浮点值能方便地通过Printf的谓词%g输出，该谓词会自动保持足够的精度。

math包还有函数用于创建和判断IEEE 754标准定义的特殊值：正**无穷大**和负无穷大，它表示超出最大许可值的数及除以零的商；以及**NaN**（Not a Number），它表示数学上无意义的运算结果（如0／0或Sqrt(-1)）。

在数字运算中，我们倾向于将NaN当作信号值（sentinel value），但直接判断具体的计算结果是否为NaN可能导致潜在错误，因为**与NaN的比较总不成立**。

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var z float64

	fmt.Println(z, 1/z, -1/z, z/z)

	nan := math.NaN()
	// 比较不成立
	fmt.Println(nan == nan, nan < nan, nan > nan)
}
```

```shell
0 +Inf -Inf NaN
false false false
```

##### 练习 3.1

```go
// surface 函数根据一个三维曲面函数计算并生成SVG
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // 以像素表示的画布大小
	cells         = 100                 // 网格单元的个数
	xyrange       = 30.0                // 坐标轴的范围(-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x或y轴上每个单位长度的像素
	zscale        = height * 0.4        // z轴上每个单位长度的像素
	angle         = math.Pi / 6         // x、y轴的角度
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok := corner(i+1, j)
			if !ok {
				continue
			}

			bx, by, ok := corner(i, j)
			if !ok {
				continue
			}

			cx, cy, ok := corner(i, j+1)
			if !ok {
				continue
			}

			dx, dy, ok := corner(i+1, j+1)
			if !ok {
				continue
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// 求出网格单元(i,j)的顶点坐标(x,y)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 计算曲面高度z
	z := f(x, y)

	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, false
	}

	// 将(x,y,z)等角投射到二维SVG绘图平面上，坐标是(sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // 到(0,0)的距离
	return math.Sin(r) / r
}
```

##### 练习 3.4

```go
// surface 函数根据一个三维曲面函数计算并生成SVG
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // 以像素表示的画布大小
	cells         = 100                 // 网格单元的个数
	xyrange       = 30.0                // 坐标轴的范围(-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x或y轴上每个单位长度的像素
	zscale        = height * 0.4        // z轴上每个单位长度的像素
	angle         = math.Pi / 6         // x、y轴的角度
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler) // 回声请求调用处理程序
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok := corner(i+1, j)
			if !ok {
				continue
			}

			bx, by, ok := corner(i, j)
			if !ok {
				continue
			}

			cx, cy, ok := corner(i, j+1)
			if !ok {
				continue
			}

			dx, dy, ok := corner(i+1, j+1)
			if !ok {
				continue
			}

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// 求出网格单元(i,j)的顶点坐标(x,y)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 计算曲面高度z
	z := f(x, y)

	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, false
	}

	// 将(x,y,z)等角投射到二维SVG绘图平面上，坐标是(sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // 到(0,0)的距离
	return math.Sin(r) / r
}
```



#### 3.3 复数

Go具备两种大小的复数complex64和complex128，二者分别由float32和float64构成。 内置的complex函数根据给定的实部和虚部创建复数，面内置的real函数和imag函数则分别提取复数的实部和虚部。

```go
var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i
fmt.Println(x*y)                 // "(-5+10i)"
fmt.Println(real(x*y))           // "-5"
fmt.Println(imag(x*y))           // "10"
```

根据常量运算规则，复数常量可以和其他常量相加（整型或浮点型，实数和虚数皆可）。可以用==或!=判断复数是否等值。若两个复数的实部和虚部都相等，则它们相等。



#### 3.4 布尔值

短路行为：如果运算符左边的操作数已经能直接确定总体结果，则右边的操作数不会计算在内。

&&较||优先级更高（助记窍门：&&表示逻辑乘法，||表示逻辑加法）。



#### 3.5 字符串

字符串是**不可变的字节序列**，它可以包含任意数据，包括0值字节，但主要是人类可读的文本。试图访问许可范围以外的字节会触发宕机异常。**加号（＋）运算符连接两个字符串而生成一个新字符串**。

> 这一点同Python，区别与C/C++。

**字符串可以通过比较运算符做比较**，如＝=和＜；比较运算按字节进行，结果服从本身的字典排序。

可以将新值赋予字符串变量，但是字符串值无法改变；字符串值本身所包含的字节序列永不可变。**因为字符串不可改变，所以字符串内部的数据不允许修改**。

不可变意味着**两个字符串能安全地共用同一段底层内存**，使得复制任何长度字符串的开销都低廉。类似地，字符串s及其子串（如s[7:]）可以安全地共用数据，因此**子串生成操作的开销低廉**。这两种情况下都没有分配新内存。



##### 3.5.1 字符串字面量

Go的源文件总是按UTF-8编码，并且习惯上Go的字符串会按UTF-8解读，所以在源码中我们可以将Unicode码点写入字符串字面量。

**原生的字符串字面量**的书写形式是\`...\`，使用**反引号**而不是双引号。原生的字符串字面量内，**转义序列不起作用**；**实质内容与字面写法严格一致，包括反斜杠和换行符**，因此，在程序源码中，原生的字符申字面量可以展开多行。

**正则表达式**往往含有大量反斜杠，可以方便地写成原生的字符串字面量。原生的字面量也适用于**HTML模板**、**JSON字面量**、**命令行提示信息**，以及需要多行文本表达的场景。



##### 3.5.2 Unicode

**Unicode囊括了世界上所有文书体系的全部字符**，还有重音符和其他变音符，控制码（如制表符和回车符），以及许多特有文字，对它们各自赋予一个叫Unicode 码点的标准数字。在Go的术语中，这些字符记号称为**文字符号rune**。天然适合保存单个文字符号的数据类型就是int32，为Go所采用；正因如此，**rune类型作为int32类型的别名**。我们可以将文字符号的序列表示成int32值序列，这种表示方式称作UTF-32或UCS-4，每个Unicode码点的编码长度相同，都是32位。



##### 3.5.3 UTF-8

**UTF-8以字节为单位对Unicode码点作变长编码**。UTF-8是现行的一种Unicode标准，由Go的两位创建者**Ken Thompson**和**Rob Pike**发明。每个文字符号用1～4个字节表示。ASCII字符的编码仅占1个字节，而其他常用的文书字符的编码只是2或3个字节。**一个文字符号编码的首字节的高位指明了后面还有多少字节**。若最高位为0，则标示着它是7位的ASCII码，其文字符号的编码仅占1字节，这样就与传统的ASCHI码一致。若最高几位是110，则文字符号的编码占用2个字节，第二个字节以10开始。更长的编码以此类推。

```shell
0xxxxxxx                             runes 0-127    (ASCII)
110xxxxx 10xxxxxx                    128-2047       少于128个未使用的值
1110xxxx 10xxxxxx 10xxxxxx           2048-65535     少于2048个未使用的值
11110xxx 10xxxxxx 10xxxxxx 10xxxxxx  65536-0x10ffff 其他未使用的值
```

UTF-8编码紧凑，兼容ASCII，并且自同步；最多追溯3字节，就能定位一个字符的起始位置。**UTF-8还是前缀编码**，因此它能从左向右解码而不产生歧义，也无须超前预读。**文字符号的字典字节顺序与Unicode码点顺序一致**（Unicode 设计如此），因此按UTF-8编码排序自然就是对文字符号排序。

**Go的源文件总是以UTF-8编码，同时，需要用Go程序操作的文本字符串也优先采用UTF-8编码**。unicode包具备针对单个文字符号的函数（例如区分字母和数字，转换大小写），而unicode/utf8包则提供了按UTF-8编码和解码文字符号的函数。

按UTF-8编码的文本的逻辑同样也适用原生字节序列，但其他编码则无法如此。

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"

	for i := 0; i < len(s); {
		// 每次DecodeRuneIn5tring的调用都返回r（文字符号本身）和一个值（表示r按UTF—8编码所占用的字节数）。
		// 这个值用来更新下标i，定位字符串内的下一个文字符号。对于非ASCII文字符号，下标增量大于1。
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	// Go的range循环也适用于字符串，按UTF-8隐式解码。
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%[2]d\n", i, r)
	}
}
```

```shell
13
9
0       H
1       e
2       l
3       l
4       o
5       ,
6        
7       世
10      界
0       'H'     72
1       'e'     101
2       'l'     108
3       'l'     108
4       'o'     111
5       ','     44
6       ' '     32
7       '世'    19990
10      '界'    30028
```

每次UTF-8解码器读入一个不合理的字节，无论是**显式调用utf8.DecodeRuneInString**， 还是在**range循环内隐式读取**，都会产生一个专门的Unicode字符“＼uFFFD＇替换它，其输出通常是个**黑色六角形或类似钻石的形状，里面有个白色问号**。如果程序碰到这个文字符号值，通常意味着，生成字符串数据的系统上游部分在处理**文本编码方面存在瑕疵**。

```go
package main

import "fmt"

func main() {
	s := "你好，世界"
	// 按16进制打印UTF-8编码
	fmt.Printf("% x\n", s)

	// 当[]rune转换作用于UTF—8编码的字符串时，返回该字符串的Unicode码点序列
	r := []rune(s)
	fmt.Printf("% x\n", r)

	// 如果把rune slice转换成一个字符申，它会输出各个文字符号的UTF—8编码拼接结果
	fmt.Println(string(r))

	// 若将一个整数值转换成字符串，其值按rune类型解读，并且string()转换产生代表该文字符号值的UTF-8码
	fmt.Println(string(0x4eac)) // "京"
}
```

```shell
# e 4 b d a 0(UFT-8)
# 1110 0100 1011 1101 1010 0000
# 0100 1111 0110 0000
# 4 f 6 0(Unicode)
e4 bd a0 e5 a5 bd ef bc 8c e4 b8 96 e7 95 8c
[ 4f60  597d  ff0c  4e16  754c]
你好，世界
京
```

第一个Printt里的谓词%x（注意，%和x之间有空格）以十六进制数形式输出，并在每两个数位间插入空格。



##### 3.5.4 字符串和字节slice

4个标准包对字符串操作特别重要：bytes、strings、strconv和unicode。

**<u>strings包</u>**提供了许多函数，用于**搜索、替换、比较、修整、切分与连接字符串**。

**<u>bytes包</u>**也有类似的函数，用于**操作字节slice**（[]byte类型，其某些属性和字符申相同）。由于字符串不可变，因此按增量方式构建字符串会导致多次内存分配和复制。这种情况下，使用bytes.Buffer类型会更高效。

**<u>strconv包</u>**具备的函数，主要用于**转换布尔值、整数、浮点数为与之对应的字符串形式**，或者**把字符串转换为布尔值、整数、浮点数**，另外还有为字符串添加／去除引号的函数。

**<u>unicode包</u>**备有判别文字符号值特性的函数，如IsDigit、IsLetter、IsUpper和IsLower。 **每个函数以单个文字符号值作为参数**，并返回布尔值。转换函数（如ToUpper 和ToLower）将其转换成指定的大小写。strings包也有类似的函数，函数名也是ToUpper和ToLower，它们**对原字符串的每个字符做指定变换**，生成并返回一个新字符串。

**basename函数**模仿UNIX shell中的同名实用程序。只要s的前缀看起来像是文件系统路径（各部分由斜杠分隔），该版本的basename(s)就将其移除，貌似文件类型的后缀也被移除。

若字符串包含一个字节数组，创建后它就无法改变。相反地，字节slice的元素允许随意修改。字符串可以和字节slice相互转换：

```go
s := "abc"
b := []byte(s)
s2 := string(b)
```

[]byte(s)转换操作会分配新的字节数组，拷贝填入s含有的字节，并生成一个slice引用，指向整个数组。具备优化功能的编译器在某些情况下可能会避免分配内存和复制内容，但一般而言，复制有必要确保s的字节维持不变（即使b的字节在转换后发生改变）。反之，用string(b)将字节slice转换成字符串也会产生一份副本，保证s2也不可变。

bytes包为高效处理字节slice提供了Buffer类型。Buffer起初为空，其大小随着各种类型数据的写人而增长。

```go
package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer

	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')

	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
}
```

```shell
[1, 2, 3]
```

若要在bytes.Buffer变量后面添加任意文字符号的UTF-8编码，最好使用bytes.Buffer的writeRune方法，而追加ASCII字符，如“［＇和＇］＇，则使用writeByte亦可。

**bytes.Buffer**类型用途极广，假若I/O函数需要一个字节接收器（**io.writer**）或字节发生器（**io.Reader**），我们将看到能如何**用其来代替文件**，其中接收器的作用就如上例中的Fprintf一样。

##### 练习 3.10

```go
package comma // import "github.com/AntiBargu/gopl/ch3/comma"

import (
	"bytes"
)

func Comma(s string) string {
	var buf bytes.Buffer

	n := len(s)

	if n == 0 {
		return buf.String()
	}

	i := n % 3

	if i == 0 {
		i = 3
	}

	buf.WriteString(s[:i])
	for ; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}
```

##### 练习 3.11

```go
package comma1 // import "github.com/AntiBargu/gopl/ch3/comma1"

import (
	"bytes"
	"strings"
)

func Comma(s string) string {
	var buf bytes.Buffer

	if len(s) == 0 {
		return buf.String()
	}

	// 处理正负号
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	// 分成整数部和小数部，分别处理
	var intStr, floatStr string

	dotIdx := strings.Index(s, ".")
	if -1 == dotIdx {
		intStr = s
	} else {
		intStr, floatStr = s[:dotIdx], s[dotIdx:]
	}

	i := len(intStr) % 3
	if i == 0 {
		i = 3
	}

	buf.WriteString(intStr[:i])
	for ; i < len(intStr); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(intStr[i : i+3])
	}
	buf.WriteString(floatStr)

	return buf.String()
}
```

##### 练习 3.12

```go
package same // import "github.com/AntiBargu/gopl/ch3/same"

func Same(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m1, m2 := make(map[rune]int), make(map[rune]int)

	for _, c := range s1 {
		m1[c]++
	}

	for _, c := range s2 {
		m2[c]++
	}

	if len(m1) != len(m2) {
		return false
	}

	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}

	return true
}
```



##### 3.5.5 字符串和数字的相互转换

要将整数转换成字符串，一种选择是使用fmt.Sprintf，另一种做法是用函数strconv.Itoa()。

strconv包内的Atoi函数或Parseint函数用于解释表示整数的字符串，而ParseUint用于无符号整数。

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(strconv.Itoa(x), y)

	a, _ := strconv.Atoi("123")             // x是一个整数
	b, _ := strconv.ParseInt("123", 10, 64) // 十进制，最长为64位

	fmt.Println(a, b)
}
```

```shell
123 123
123 123
```

有时候，单行输入由字符串和数字依次混合构成，需要用fmt.Scanf解释，可惜fmt.Scanf也许不够灵活，处理不完整或不规则输入时尤甚。



#### 3.6 常量

常量是一种表达式，其可以保证在**编译阶段就计算出表达式的值**，并不需要等到运行时，从而使编译器得以知晓其值。所有常量本质上都属于基本类型：布尔型、字符串或数字。

若同时声明一组常量，除了第一项之外，其他项在等号右侧的表达式都可以省略，这意味着会复用前面一项的表达式及其类型。

```go
const (
    a = 1
    b
    c = 2
    d
)

fmt.Println(a, b, c, d) // "1 1 2 2"
```



##### 3.6.1 常量生成器iota

常量声明中，iota从0开始取值，逐项加1。这种类型通常称为枚举型（enumeration，或缩写成enum）。

> 类似于C/C++中的enum。

```go
package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Monday, Friday, Tuesday)
}
```

```shell
1 5 2
```

针对相应的位执行判定、设置或清除操作，就会用到这些常量：

```go
package main

import "fmt"

type Flags uint

const (
	FlagUp           Flags = 1 << iota // 向上
	FlagBroadcast                      // 支持广播访问
	FlagLoopback                       // 是环回接口
	FlagPointToPoint                   // 属于点对点链路
	FlagMulticast                      // 支持多播访问
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}
```

```shell
10001 true
10000 false
10010 false
10010 true
```

1024的幂：

```go
const (
    _ = 1 << (10 * iota)
    KiB // 1024
    MiB // 1048576
    GiB // 1073741824
    TiB // 1099511627776             (exceeds 1 << 32)
    PiB // 1125899906842624
    EiB // 1152921504606846976
    ZiB // 1180591620717411303424    (exceeds 1 << 64)
    YiB // 1208925819614629174706176
)
```

##### 练习 3.13

```go
const (
	KB = 1000
  MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)
```

 

##### 3.6.2 无类型常量

许多**常量并不从属某一具体类型**。编译器将这些从属类型待定的常量表示成某些值，这些值比基本类型的数字精度更高，且算术精度高于原生的机器精度。可以认为它们的精度至少达到256位。从属类型待定的常量共有6种，分别是无类型布尔、无类型整数、无类型文字符号、无类型浮点数、无类型复数、无类型字符串。

无类型常量不仅能**暂时维持更高的精度**，与类型已确定的常量相比，它们还能写进更多表达式而无需转换类型。比如，上例中ZiB和YiB的值过大，用哪种整型都无法存储，但它们都是合法常量并且可以用在下面的表达式中：

```go
package main

import "fmt"

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func main() {
	fmt.Println(YiB / ZiB)
}
```

```shell
1024
```

只有常量才可以是无类型的。若将无类型常量声明为变量，或在类型明确的变量赋值的右方出现无类型常量，则常量会被隐式转换成该变量的类型。
