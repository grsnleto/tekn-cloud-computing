# Latihan Minggu Ke-06

<h3>A. Implementasi Program Go Untuk Membaca Data Pada Mysql</h3>

1. Melakukan instalasi driver mysql untuk bisa terkoneksi ke mysql server, seperti pada gambar dibawah ini: <br/>
<dd>

![gambar-02](/minggu-06/Latihan/gambar-02.jpg)
</dd>

2. Kemudian sebelumnya mengecek ulang database dan tabel, serta isian data yang sebelumnya sudah di insert pada tabel.
Seperti pada gambar dibawah ini :
<dd>

![gambar-01](/minggu-06/Latihan/gambar-01.jpg)
</dd>

3. Selanjutnya membuat program go untuk melakukan koneksi ke database dengan melakukan remote ke mysql server dan kemudian mengekusinya dan melakukan baca data dari database dan tabel, Seperti pada script program dibawah ini : <br/>
<dd>

| Koneksi  | Membaca Data Mysql |
|---|---|
|[Lihat Script](/minggu-06/Latihan/main.go)   | [Lihat Script](/minggu-06/Latihan/baca-Mysql.go)  |   

Sehingga hasil yang ditampilkan ketika melakukan eksekusi terhadapat program koneksi dan baca data mysql, akan menghasilkan output seperti pada ambar dibawah ini :<br/>
1. Output untuk program koneksi :<br/>
![gambar-04](/minggu-06/Latihan/gambar-04.jpg)
2. Output untuk Program Membaca Data MySql :<br/>
![gambar-03](/minggu-06/Latihan/gambar-03.jpg)
</dd>

<h3>B. Implementasi Program Go Untuk Membaca Data Pada MongoDB</h3>

1. Sebelumnya perlu dijelaskan bahwa, disini saya mengunakan mongodb atlas, sifat penggunaannya dengan mongodb compas tidak jauh berbeda, hanya saja pada mongodb atlas harus membuat clusters(ini bermaksud untuk mendapatkan URL untuk nantinya dikonfigurasikan dengan drivers yang sudah tersedia). Kemudian baru mendeklarasikan database dan juga collectionsnya.

2. MongoDB atlas juga merupakan jenis database NoSQL yang disediakan pihak layanan cloud, yang artinya dia memiliki versi cloud sehingga sedikit berbeda dengan mongodb compas. Dan juga keuntungannya lebih mempermudah dalam proses kontrol database.

3. Pada bagian pertama, dimana sebelumnya saya sudah membuat clusters terlebih dahulu dengan nama <b>tcc-175610055</b>, sehingga halaman yang ditampilkan seperti gambar dibawah ini :
<dd>

![gambar-05](/minggu-06/Latihan/gambar-05.jpg)
</dd>

4. Pada halaman diatas tersebut klik tombol <b>CONNECT</b> untuk melakukan koneksi dengan aplikasi saya melalui driver. Kemudian selanjutnya untuk tahap berikutnya pilih bagian yang ditandai merah seperti gambar dibawah ini :<br/>
<dd>

![gambar-06](/minggu-06/Latihan/gambar-06.jpg)

Dan kemudian setelah diklik secara otomatis akan mendapatkan code URL, seperti pada gambar dibawah ini :<br/>

![gambar-07](/minggu-06/Latihan/gambar-07.jpg)
</dd>

5. Kemudian selanjutnya melakukan download dirver yang dibutuhkan yang nantinya digunakan untuk konfiurasi dengan clusters yang sebelumnya sudah dibuat, seperti pada gambar dibawah ini :<br/>
<dd>

![gambar-08](/minggu-06/Latihan/gambar-08.jpg)
</dd>

6. Kemudian melihat kembali database dan collection yang dibuat serta insert data pada collection, seperti pada gambar dibawah ini :<br/>
<dd>

![gambar-09](/minggu-06/Latihan/gambar-09.jpg)
dimana saya membuat database dengan nama <b>tccMhs</b> didalamnya dibuat colection dengan nama <b>mahasiswa</b>. 
</dd>

7. Selanjutnya membuat program go untuk melakukan koneksi ke database dengan melakukan remote ke mongodb server dan kemudian mengekusinya dan melakukan baca data dari database dan collection, Seperti pada script program dibawah ini : <br/>
<dd>

| Koneksi  | Membaca Data MongoDB |
|---|---|
|[Lihat Script](/minggu-06/Latihan/konek-mongo.go)   | [Lihat Script](/minggu-06/Latihan/baca-Mongo.go)  | 

Sehingga hasil yang ditampilkan ketika melakukan eksekusi terhadapat program koneksi dan baca data mongodb, akan menghasilkan output seperti pada ambar dibawah ini :<br/>
1. Output untuk program koneksi :<br/>
![gambar-10](/minggu-06/Latihan/gambar-10.jpg)
2. Output untuk Program Membaca Data MySql :<br/>
![gambar-11](/minggu-06/Latihan/gambar-11.jpg)
</dd>
</dd>
