# (21) Clean Architecture 

## 1. Introduction
Clean architecture merupakan suatu pendekatan untuk membangun sebuah aplikasi yang mudah untuk di maintain dan di test.
Clean architecture sendiri merupakan sebuah pendekatan yang diusulkan oleh Robert C. Martin (Uncle Bob) pada tahun 2012. Clean architecture sendiri merupakan sebuah pendekatan yang berbeda dengan pendekatan lainnya seperti MVC, MVP, dll.

## 2. Clean Architecture
Clean architecture sendiri terdiri dari 3 layer, yaitu:
- Entities
- Use Cases
- Interfaces

### 2.1 Entities
Entities merupakan sebuah layer yang berisi semua model yang digunakan dalam aplikasi. Entities ini merupakan sebuah layer yang tidak memiliki dependensi terhadap layer lainnya.

### 2.2 Use Cases
Use Cases merupakan sebuah layer yang berisi semua use case yang digunakan dalam aplikasi. Use Cases ini merupakan sebuah layer yang tidak memiliki dependensi terhadap layer lainnya.

### 2.3 Interfaces
Interfaces merupakan sebuah layer yang berisi semua interface yang digunakan dalam aplikasi. Interfaces ini merupakan sebuah layer yang memiliki dependensi terhadap layer lainnya.

