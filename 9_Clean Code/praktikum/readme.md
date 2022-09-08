# Praktikum Clean Code

## Problem 1


Kode awal:
```go
package main

type user struct {
    id        int
    username  int
    password  int
}

type userservice struct {
	t []user
}

func (u userservice) getallusers() []user {
    return u.t
}

func (u userservice) getuserbyid(id int) user {
    for _, v := range u.t {
        if v.id == id {
            return v
        }
    }
    return user{}
}
```

Kode yang sudah diperbaiki

```go
package main

type User struct {
    id        int
    username  int
    password  int
}

type userService struct {
	users []User
}

func (u userService) getAllUsers() []User {
    return u.users
}

func (u userService) getUserById(id int) User {
    for _, userServices := range u.users {
        if userServices.id == id {
            return userServices
        }
    }
    return User{}
}
```

## Problem 2

Kode sebelum di-rewrite:
```go
package main

type kendaraan struct {
	totalroda int
	kecepatanperjam int
}

type mobil struct {
    kendaraan
}

func (m *mobil) berjalan(){
	m.tambahkecepatan(10)
}

func (m *mobil) tambahkecepatan(kecepatanbaru int){
    m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
}
func main() {
	mobilcepat := mobil{}
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	
	mobillamban := mobil{}
	mobillamban.berjalan()
}
```

Kode setelah di-rewrite:
```go
package main

type Kendaraan struct {
	kecepatan int
}

type Mobil struct {
    Kendaraan
}

func (m *Mobil) berjalan(){
	m.kecepatan += 10
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	
	mobillamban := Mobil{}
	mobillamban.berjalan()
}
```