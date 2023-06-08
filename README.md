# go_study
1.切片：拥有相同类型元素的可变长度的序列,基于数组类型做了一层封装
    使用make或直接[]T的方式进行定义
2.指针是带类型的，不同类型的指针不能进行赋值
3.new是用来初始化值类型指针的
    ```
    var a = new(int) //得到一个int类型指针a
    var b = new([3]int) //得到一个int类型数组，b依然是指针
    ```
4.make是用来初始化slice、map、chan的
5.panic和recover
    panic是代码运行时的错误，会导致程序奔溃，异常退出；可以在代码中调用panic关键字让程序异常退出，类似内核中的bug()函数
    recover将程序从panic中恢复，recover只能在defer函数中使用
6.函数和方法的区别
    函数是谁都可以调用的
    方法是只有某个特定的类型才能调用，函数指定接受者就是方法，又叫方法接收器
        `func (t Type) 函数明(参数) (返回值1， 返回值2)`
		相比函数的定义多指定了(t Type)，表示该方法调用者必须是Type类型实例
		注意：不可以给別的包定义的类型添加方法，可以给自定义的类型添加方法
7.结构体的骚操作
    匿名字段：结构体中某些字段没有指定字段名，只指定类型，如下：
        ```
        type student struct {
            name string
            string
            int
        }```
        如果不指定字段名，不可以出现相同的类型,该写法不推荐，实例化如下：
        ```
        var stu1 student {
            name: "小明"
            string: "想有个家！"
            int: "0"
        }
        ```
        访问结构体中的匿名字段：`fmt.Println(stu1.string)`
    结构体嵌套时：可以直接访问嵌套结构体的一个字段，使用该特性可以模拟“继承”
        ```
        type address struct {
            province string
            city string
        }
        type student struct {
            name string
            age string
            adr address
        }
        ```
        如果访问实例stu1中adr的city字段一般方式：`fmt.Println(stu1.adr.city)`
                                    高级的方式：`fmt.Println(stu1.city)`
8.json序列化和反序列化
    要求结构体字段名首字母大写，小写的话json包中就访问不了结构体变量。
    序列化：json.Marshal()和json.MarshalIndent()
    反序列化：json.Unmarshal()
    成员标签定义：结构体成员在编译期间关联的一些元信息。
9.defer和return：return x，先执行返回值=x然后执行ret指令；defer是在赋值和返回中间执行的，即先执行返回值=x，然后执行defer，最后执行ret指令。
10.接口（interface）是一种抽象的对象，关心动作不关心数据，接口有数个方法组成，定义格式如下：
    ```
    type 接口类型名1 interface {
        方法名1(参数列表1) 返回值列表1
        方法名2(参数列表2) 返回值列表2
    }
    ```
    只要一个类型实现了“方法名1”和“方法名2”，我们就认为这个类型实现了“接口类型1”这个接口，也就是说这个类型是“接口类型1”的一个变量。
    判断方法是否已经实现，方法名相同，参数和返回值只要类型和数量一样即可，不需要参数名和返回值名相同。
11.空接口：没有定义方法的接口，即interface{}类型，任何类型都是空接口的实例，可以接收任何类型
    空接口作为函数参数，可以像函数传递任何参数
    空接口作为map的值，同一个map中存放不同类型的数据
12.接口嵌入：避免在多个地方重复声明相同的方法
13.http.Handler接口
    创建http服务端时会使用http.ListenAndServe()函数，该函数的第二个参数为Handler接口实例，Handler接口中包含方法ServeHTTP（w ResponWriteer， r *Request）。
    Handler接口实例的ServeHTTP方法中可以判断URL.Path，不同的path进行不同的操作，但实际的应用中将不同的操作定义到不同的方法或函数中会很实用，所以net/http包请求多路器ServeMux来简化URL和Handlers的联系，使用方法如下：
        ```
        mux := http.NewServeMux()
        mux.Handle("/list", http.HandlerFunc(db.list))
        mux.Handle("/price", http.HandlerFunc(db.price))
        log.Fatal(http.ListenAndServe("localhost:8000", mux))
        ```
    可使用ServeMuxde的方法Headle将URL.Path对应的操作函数注册到ServeMuxde中的一个map中，map的key值为URL路径名（mux.Handle的第一个参数），一般对应的操作函数不满足http.Handler接口，所以不能直接传给Mux.Handle。
    通过http.HandlerFunc()将操作函数装换为一个满足http.Handler接口参数，所以注册URL对应的操作函数：mux.Handle("/list", http.HandlerFunc(db.list))
    ServeMux还有一个更方便的HandleFunc方法简化了注册的代码：mux.HandleFunc("/list", db.list)
    大多数程序中使用一个web服务器足够了，另外在多个文件添加handle时需要将Mux设置为全局的。为了满足这样的需求，提供了一个全局的ServeMux实例DefaultServerMux和包级别的http.Handle和http.HandleFunc函数。
    现在，为了使用DefaultServeMux作为服务器的主handler，我们不需要将它传给 ListenAndServe函数；nil值就可以工作。
14.类型断言：将空接口转化成指定类型的值，使用x.(T)，其中x是类型为interface{}的变量，T表示x可能的类型（要转换的类型），通常如下使用：
    ```
    switch v := x.(type) {
    case int:
        fmt.Println("x为int类型")
    case string:
        fmt.Println("x为string类型")
    default：
        fmt.Println("x为不知道的类型")
    }
    ```
15.常量生成器iota从0开始，逐项加1，用法如下：
    ```
    type Weekday int
    const (
        Sunday weekday = iota
        Monday
        Tuesday
        Wednesday
        Thursday
        Friday
        Saturday
    )
    ```
    上面表明，Sunday的值为0，Monday的值为1，以此类推。
16.反射：Go语言提供一种机制，能够在运行时更新变量和检查它们的值，调用它们的方法和支持的内在操作而不需要在编译时就知道这些变量的具体类型，这种机制称为反射。
    Go程序在运行期间使用reflect包访问程序的反射信息，reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type。
    应用：各种web框架，配置文件解析库、ORM框架
    reflect.Typeof()可以获取任意值的类型对象（reflect.Type）
        反射中关于类型分为两种：类型（Type）和种类（Kind）,在Go语言中我们可以使用type关键字构造很多自定义类型，而种类（Kind）就是这些自定义类型的底层类型。
        当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）。 
    reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。可使用的reflect.Value类型的方法将reflect.Value转换获取原始值。
        ```
        v := reflect.ValueOf(3)
        x := v.Int()    //reflect.Value类型方法还有Interface()、Uint()、String()、Bool()、Float()等
        fmt.Printf("%T, %d", x, x) //int64, 3
        ```
    通过反射取值：
        ```
        func reflectValue(x interface{}) {
            v := reflect.ValueOf(x)
            k := v.Kind()
            switch k {
            case reflect.Int64:
                // v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
                fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
            case reflect.Float32:
                // v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
                fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
            }
        }
        ```
    通过反射修改变量的值：
        修改变量的值时，必须传递变量地址进行修改，然后使用专有的Elem()方法来获取指针对应的值。
        ```
        func reflectSetValue2(x interface{}) {
            v := reflect.ValueOf(x)
            // 反射中使用 Elem()方法获取指针对应的值
            if v.Elem().Kind() == reflect.Int64 {
                v.Elem().SetInt(200)
            }
        }
        ```
    IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。
    IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。
    结构体反射：任意值通过reflect.TypeOf()获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象（reflect.Type）的NumField()和Field()方法获得结构体成员的详细信息。
