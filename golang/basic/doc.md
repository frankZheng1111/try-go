## try-go

- 尝试一下go

#### 关于import

1. import 导入的参数是路径，而非包名。

2. 尽管习惯将包名和目录名保证一致，但这不是强制规定；

3. 在代码中引用包成员时，使用包名而非目录名；

4. 同一目录下，所有源文件必须使用相同的包名称（因为导入时使用绝对路径，所以在搜索路径下，包必须有唯一路径，但无须是唯一名字）；

5. go的package是以绝对路径GOPATH来寻址的，不要用相对路径来import

#### 关于 := 定义变量

- ":=": 对于使用:=定义的变量，如果新变量与那个同名已定义变量 (比如说同名全局变量)不在一个作用域中时，那么golang会新定义这个变量p，遮盖住全局变量p，这就是导致这个问题的真凶。

- 使用简短方式 := 声明的变量的作用域只存在于 if 结构中（在 if 结构的大括号之间，如果使用 if-else 结构则在 else 代码块中变量也会存在）。如果变量在 if 结构之前就已经存在，那么在 if 结构中，该变量原来的值会被隐藏。最简单的解决方案就是不要在初始化语句中声明变量

#### 关于package

- golang的package和其他语言的组织方式完全不同，刚开始接触时，很不适应。

- golang的package的特点：

  1. go的package不局限于一个文件，可以由多个文件组成。组成一个package的多个文件，编译后实际上和一个文件类似，组成包的不同文件相互之间可以直接引用变量和函数，不论是否导出；因此，组成包的多个文件中不能有相同的全局变量和函数（这里有一个例外就是包的初始化函数：init函数，下面还有讨论）

  2. go不要求package的名称和所在目录名相同，但是你最好保持相同，否则容易引起歧义。因为引入包的时候，go会使用子目录名作为包的路径，而你在代码中真正使用时，却要使用你package的名称。

  3. 每个子目录中只能存在一个package，否则编译时会报错。


#### 关于package的初始化函数init的说明：

- 每个package中应该是每个init都会被调用，且顺序固定

  1. 对同一个go文件的init()调用顺序是从上到下的
 
  2. 对同一个package中不同文件是按文件名字符串比较“从小到大”顺序调用各文件中的init()函数,对于
 
  3. 对不同的package，如果不相互依赖的话，按照main包中"先import的先调用"的顺序调用其包中的init()
 
  4. 如果package存在依赖，则先调用最早被依赖的package中的init(), 即最底层的init

