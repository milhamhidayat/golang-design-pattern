Builder

Builder pattern merupakan pattern dimana 
- memisahkan kode untuk membuat object ke _builders_ sehingga terpisah dari object tersebut. Akan membuat kode untuk membuat product dapat memiliki implementasi yang berbeda. 
- membuat objek secara bertahap, selangkah demi selangkah.

![alt text](https://refactoring.guru/images/patterns/content/builder/builder.png)


**Masalah:**

- Ketika

- Parameter terlalu banyak 

Ketika ingin membuat object rumah, maka harus menyediakan jendela, pintu, kamar. Opsi tambahan adalah garasi, kolam renang, taman, dll. Jika menggunakan constructor biasa, maka parameter constructor akan menjadi sangat panjang. 


```java
class House {
    public House(int windows, int doors, int rooms, int phone ,bool hasGarage, bool HasSwimPool, boolHasGarden){}
}

House h1 = new House(8, 6, 4, 021723343, 1, 1, 1, 1, 1)
```

![alt text](https://refactoring.guru/images/refactoring/content/smells/long-parameter-list-01.png)


Ketika sebuah rumah tidak memiliki nomor telepon rumah, maka nilainya tidak ada dan harus set nilai phone menjadi null

```java
House h1 = new House(8, 6, 4, null, 1, 1, 1, 1, 1)
```

Hal ini membuat kode menjadi lebih susah dibaca karena tidak tahu nilai null digunakan untuk set variabel apa. 

![alt text](https://refactoring.guru/images/patterns/diagrams/builder/problem2.png)


Kita bisa membuat overloading constructor untuk mengatasi masalah tersebut.

```java
class House {
    public House(int windows, int doors, int rooms, int phone ,bool hasGarage, bool HasSwimPool, boolHasGarden){}
    public House(int windows, int doors, int rooms, bool hasGarage, bool HasSwimPool, boolHasGarden){}
}
```

Namun akan mengakibatkan jumlah constructor yang banyak karena bergantung pada parameter yang diberikan.

**Solusi:**

- Solusi dari builder pattern adalah memindahkan fungsi untuk pembuatan object ke object lain yang disebut builder.

![alt text](https://refactoring.guru/images/patterns/diagrams/builder/solution1.png)


- Pembuatan object dilakukan dengan beberapa tahapan (buildWalls, buildDoor)
- Beberapa tahapan pembangunan object memiliki implementasi yang berbeda tergantung pada object product yang akan dibuat.

    Contoh:

    - Rumah : butuh 4 jendela
    - Istana : butuh 30 jendela

**Elemen penyusun builder pattern

- Product : jenis product  yang akan dibuat oleh builder. Menentukan elemen penyusunnya (blueprint)
- Builder : builder interface menentukan langkah / tahapan yang dilakukan untuk membuat sebuah product
- Concrete builder : object product menyediakan implementasi langkah / tahapan yang dilakukan untuk membuat sebuah product. Implementasi dari builder interface
- Director : menentukan urutan tahapan / langkah dalam membuat object product. Sehingga dapat dipanggil kembali jika ingin membuat sebuah object product

Note : director dipakai ketika urutan tahapan / langkah sudah fix dan dapat diimplementasikan di setiap pembuatan object product. Sehingga 2 object product dapat memiliki nilai atribut yang sama. Tidak bisa memberikan nilai input saat ada di client code.

![alt text](https://refactoring.guru/images/patterns/diagrams/builder/structure.png)

**Implementasi Builder Pattern**
- Buat _product_ (cetak biru) yang akan dibuat
- Buat _interface builder_, digunakan untuk menentukan tahapan dalam membuat product
- Buat _director_, untuk menentukan urutan. _director_ akan memanggil implementasi _interface builder_
- Buat _concrete builder_, implementasi dari _interface builder_. Kode implementasi untuk membuat product

**Penggunaan Builder Pattern**

- Gunakan builder pattern untuk menghapus override constructor yang memiliki jumlah parameter berbeda-beda untuk membuat sebuah object
- Gunakan builder pattern ketika tahapan membuat product sama, namun berbeda di implementasinya 

**Keuntungan**
- Dapat membuat object secara bertahap
- Dapat menggunakan kembali code untuk membuat tahapan dalam membuat product (product, builder, director)
- Menggunakan konsep _Single Responsiblity Principle_. Dimana bagian untuk membuat product dipisahkan dari business logic product.

**Kekurangan**
- _Code Complexity_ akan meningkat karena harus membuat class baru setiap kali ingin membuat concrete builder baru.

**NOTE**
- kita bisa membuat builder pattern tanpa menggunakan director
- kita bisa membuat builder pattern dengan menggunakan functional options
- Kita bisa menggunakan struct sebagai builder pattern, karena kita bisa langsung memberikan value pada attribut. Sehingga kita tidak perlu menggunkan nilai null jika kita tidak ingin memberikan nilai pada sebuah atribut, karena sudah dihandle golang secara otomatis.

```go
type User struct {
    name string
    phone int
    address string
}

func main(){
    user := User{
        name: "sadf",
        phone: 3425809,
    },
}
```

Reference :
- Refactoring guru - Builder pattern : https://refactoring.guru/design-patterns/builder
- Jon Callhoun - Using functional options instead of method chaining in Go :  https://www.calhoun.io/using-functional-options-instead-of-method-chaining-in-go/
