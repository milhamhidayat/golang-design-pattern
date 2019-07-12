Singleton

Singleton merupakan pattern dimana sebuah kelas hanya memiliki sebuah instance, 
dan menyediakan akses global untuk instance tersebut

Inline-style: 
![alt text](https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "Logo Title Text 1")

Reference-style: 
![alt text][logo]

[logo]: https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "Logo Title Text 2"

Contoh situasi menggunakan Singleton Pattern :
- Ketika ingin menggunakan koneksi yang sama ke database untuk membuat query

Tujuannya / masalah yang diselesaikan adalah :
- Memastikan sebuah class hanya memiliki satu instance. Cth : database, file

    Contoh :
    membuat sebuah object, namun memutuskan untuk membuat objek baru lagi. Maka akan
    mengembalikan object yang telah dibuat

- Menyediakan akses global pada instance tersebut

Aplikasi singleton pattern :
- Ketika class di program hanya diperbolehkan memiliki satu instance saja. Cth : database object
- Ketika butuh kontrol yang ketat variabel global yang menampung instance tersebut

Keuntungan :
- Memastikan sebuah class memiliki satu single instance
- Mendapatkan global akses untuk instance tersebut
- Object singleton hanya dibuat sekali saja

Kekurangan :
- Melanggar Single Responsibility Principle. Karena fungsi membuat object dan fungsionalitas class menjadi satu, 
responsiblity terlalu besar
- Melanggar Open / Closed Principle. Karena jika ingin ubah behaviour singleton, harus ubah class tersebut.
- Melanggar Dependency Inversion Principle, karena consumber bergantung langsung pada object singleton. Sehingga tingkat coupling-nya tinggi
- Harus dicek apakah multi-threaded environment tidak membuat object singleton beberapa kali

NOTE :
- Singleton pattern bisa disebut anti pattern karena menggunakan global variable
- Dapat susah didebug karena berbagai proses dapat menggunakannya secara langsung, tidak boleh digunakan terlalu banyak