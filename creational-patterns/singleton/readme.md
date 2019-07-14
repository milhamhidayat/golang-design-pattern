Singleton

Singleton merupakan pattern dimana sebuah kelas hanya memiliki sebuah instance, 
dan menyediakan akses global untuk instance tersebut

![alt text](https://refactoring.guru/images/patterns/content/singleton/singleton.png)

![alt text](https://refactoring.guru/images/patterns/content/singleton/singleton-comic-1-en.png)

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

Reference :
- Refactoring Guru - Singleton : https://refactoring.guru/design-patterns/singleton
- Go: Design Patterns for Real World Projects : https://www.amazon.com/Go-Design-Patterns-Real-World-Projects-ebook/dp/B072R4FJ45
- Go: Design Patterns for Real World Projects source code : https://github.com/PacktPublishing/Go-Design-Patterns-for-Real-World-Projects
- 6 reasons why you should avoid singleton : https://www.davidtanzer.net/david's%20blog/2016/03/14/6-reasons-why-you-should-avoid-singletons.html
- Singleton Pattern Pitfalls : https://www.vojtechruzicka.com/singleton-pattern-pitfalls

