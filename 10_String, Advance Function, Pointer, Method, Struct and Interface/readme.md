# (10) String, Advance Function, Pointer, Method, Struct and Interface

## String
Pada golang terdapat library strings, library ini mempermudah kita ketika melakukan manipulasi string. Berikut adalah beberapa fungsi yang ada pada library strings.
1. Len
2. Contains
3. Compare
4. Replace

Contoh penggunaan:
```go
str1 := "Hello World"
str2 := strings.Replace(str1, "Hello", "Hi", 1) // Output: Hi World
```

## Advance Function
- Variadic Function, merupakan fungsi yang dapat menerima banyak parameter. 
Contoh:
```go
func sum(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}
```
- Anonymous function, merupakan fungsi yang tidak memiliki nama.
Contoh:
```go
func(){
	fmt.Println("Hello World")
}
```
- Closure, merupakan fungsi yang dapat mengakses variabel diluar fungsi.
Contoh:
```go
func newCounter() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
```
- Defer function, merupakan fungsi yang akan dieksekusi setelah fungsi yang memanggilnya selesai dieksekusi.
Contoh:
```go
func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
```

## Pointer
Pointer merupakan variabel yang menyimpan alamat memori. Untuk mendapatkan alamat memori dari suatu variabel, kita dapat menggunakan operator &. Untuk mengakses nilai dari suatu variabel yang berada pada alamat memori tertentu, kita dapat menggunakan operator *.
Contoh:
```go
func main() {
    var numberA int = 4
    var numberB *int = &numberA

    fmt.Println("numberA (value)  :", numberA)
    fmt.Println("numberA (address):", &numberA)
    fmt.Println("numberB (value)  :", *numberB)
    fmt.Println("numberB (address):", numberB)
}
```
## Struct
Struct merupakan tipe data yang dapat menyimpan data dengan berbagai tipe data. Struct dapat dianggap sebagai sebuah class pada bahasa pemrograman lainnya.
Contoh:
```go
type Person struct {
    Name string
    Age  int
}
```
## Method
Method merupakan fungsi yang berada di dalam struct. Method dapat mengakses data yang ada di dalam struct.
Contoh:
```go
func (person Person) sayHello(name string) {
    fmt.Println("Hello", name, "My name is", person.Name)
}
```
## Interface
Interface merupakan tipe data yang berisi definisi dari method-method yang harus ada pada struct yang mengimplementasikan interface tersebut.
Contoh:
```go
type HasName interface {
    GetName() string
}
```