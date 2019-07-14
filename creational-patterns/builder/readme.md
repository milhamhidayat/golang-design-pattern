Builder

Builder pattern merupakan pattern dimana membuat objek secara bertahap, selangkah demi selangkah.

![alt text](https://refactoring.guru/images/patterns/content/builder/builder.png)


**Masalah:**

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
- Beberapa tahapan pembangunan object memiliki implementasi yang berbeda tergantung pada produk yang akan dibuat.

    Contoh:

    - Rumah : butuh 4 jendela
    - Istana : butuh 30 jendela

**Elemen penyusun builder pattern

- Product : jenis produk yang akan dibuat oleh builder. Menentukan elemen penyusunnya (blueprint)
- Builder : builder interface menentukan langkah / tahapan yang dilakukan untuk membuat sebuah product
- Concrete builder : menyediakan implementasi langkah / tahapan yang dilakukan untuk membuat sebuah product. Implementasi dari builder interface
- Director : menentukan urutan tahapan / langkah dalam membuat product. Sehingga dapat dipanggil kembali jika ingin membuat sebuah product

![alt text](https://refactoring.guru/images/patterns/diagrams/builder/structure.png)
