# Builder Pattern With Director

Dengan menggunakan director, kita bisa menggunakan kembali kode untuk _product_, _builder_, dan _director_ untuk pembuat object product selanjutnya. 

Namun kekurangannya adalah kita tidak bisa membuat object product dengan nilai atribut sesuai input ketika di client code. Karena sudah di set di _concrete builder_