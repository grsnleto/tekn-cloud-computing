# Penjelasan Praktikum TCC Minggu-03
<br/>

# Latihan
<br/>

1. Pada bagian pertama ini dilakukan penginstalan python dan melakukan registrasi akun heroku, <br/>
tetapi untuk bagian ini sebelumnya saya sudah menginstal pyhton dan mempunyai akun heroku, sehingga untuk<br/>
melakukan proses ini saya lansung melewatinya.

2. Untuk proses berikutnya akan melakukan login pada heroku, disini saya menggunakan command git-bash untuk melakukan penagaksesan heroku. berikut merupakan perintahnya :<br/>
<dd>

![gambar-01](/minggu-03/Latihan//gambar-01.jpg)

Dari perintah diatas apabila sudah mengisi perintah "heroku login", maka secara otomatis akan lansung terhubung ke web heroku yaitu halaman login. Seperti pada gambar dibawah ini :<br/>

![gambar-02](/minggu-03/Latihan/gambar-02.jpg)</dd>

3. Pada bagian ke-3 ini dilakukan proses clonning dari git heroku untuk mengambil repo "python-getting-started" ke bash lokal komputer. untuk perintahnya seperti pada gambar yang tertera dibawah ini :<br/>
<dd>

![gambar-03](/minggu-03/Latihan/gambar-03.jpg)</dd>

4. Selanjutnya masuk pada repo lokal yang baru di push dan kemudian melakukan initialisasi repo lokal yang baru di clonning tersebut. Setelah melakukan initialisasi tersebut maka langkah selanjutnya melakukan remote ke app yag sudah di create sebelumnya pada akun heroku, ini beguna untuk bisa terhubung ke app pada heroku. untuk perintahnya seperti pada gambar dibawah ini :<br/>
<dd>

![gambar-04.jpg](/minggu-03/Latihan/gambar-04.jpg)</dd>

5. Kemudian selanjutnya melakukan konfiguasi pada web scale dengan nilai konfigurasinya 1, apabila 1 maka app yang di create pada heroku dapat diakses, dan kemudian untuk membukanya dengan menggunakan perintah "heroku open". seperti pada gambar dibawah ini :<br/>

<dd>

![gambar-05](/minggu-03/Latihan/gambar-05.jpg)

Sehingga hasil yang di tampilkan pada web browser seperti dibawah ini :<br/>

![gambar-06.jpg](/minggu-03/Latihan/gambar-06.jpg)</dd>

6. Kemudian pada bagian ke-06 ini akan dilakukan proses untuk melihat lalulintas data, yang artinya apabila terdapat pengujung yag mengakases website kita, maka kita bisa melihatnya. sehingga perintah yang digunakan yaitu "heroku logs --tail". Seperti yang tertera pada gambar dibawah ini :<br/>
<dd>

![gambar-07.jpg](/minggu-03/Latihan/gambar-07.jpg)</dd>

7. Selanjutnya akan dilihat dyno yang dipakai oleh kita, untuk perintahnya seperti yang tertera pada gambar dibawah ini :<br/>

<dd>

![gambar-08.jpg](/minggu-03/Latihan/gambar-08.jpg)

Pada bagian pertama diatas menampilkan berapa dyno free yang dipakai oleh kita, dyno yang digunakan disini mempunyai batasan kerena menggunakan dyno yang free.<br/>
dan kemudian perintah dibawahnya untuk melihat web akses kita, apabila bernilai 0 maka web akses kita tidak ada, dan apabila dijalankan pada web browser akan menampilkan error.<br/>
Sehingga yag harus digunakan yaitu web scale yang bernilai 1 untuk dapat mengakses web kita.</dd>

8. Kemudian selanjutnya menjalankan web scale dengan nilainya 0, seperti yang dijelaskan sebelumya, apabila web scale bernilai 0 maka web akses kita tidak ada, dan apabila dijalankan pada web browser akan menampilkan error. Seperti pada gambar dibawah ini :<br/>
<dd>

![gambar-09](/minggu-03/Latihan/gambar-09.jpg)

Untuk hasilnya ketika dijalankan perintah heroku open untuk membuka app kita pada web, Seperti pada gambar dibawah ini :<br/>

![gambar-10.jpg](/minggu-03/Latihan/gambar-10.jpg)</dd>

9. Kemudian selanjutnya menginstal pip dan mendeklarasikan file requirements.txt, untuk perintah dan hasil eksekusinya seperti pada gambar dibawah ini :<br/>
<dd>

![gambar-11.jpg](/minggu-03/Latihan/gambar-11.jpg)</dd>

10. Kemudian pada proses berikutnya membuka file static python dari repo lokal yang sebelumnya sudah di clone. dan kemudian memberikan hak akases melalui local web windows kita, dan kemudian secara otomatis akan menampilkan halaman otomatis windows firewall, disitu lansung di klik tombol "allow access". untuk kedua perintahnya seperti dibawah ini :<br/>

<dd>

![gambar-14.jpg](/minggu-03/Latihan/gambar-14.jpg)

Untuk hasil yang ditampilkan pada web local seperti pada gambar dibawah ini :<br/>

![gambar-16.jpg](/minggu-03/Latihan/gambar-16.jpg)</dd>

11. Selanjutnya untuk melihat running web local pada heroku dan melihat hasil pada web local kita, dengan menggunakan perintah seperti pada gambar dibawah ini :<br/>
<dd>

![gambar-17.jpg](/minggu-03/Latihan/gambar-17.jpg)

Dimana hasil yang ditampilkan pada web local kita, dimana eksekusi file manage.py dari repo local heroku kita, seperti pada gambar dibawah ini :<br/>

![gambar-18.jpg](/minggu-03/Latihan/gambar-18.jpg)
</dd>

12. Kemudian selanjutnya akan dilakukan push ke app kita pada heroku, untuk melihat hasilnya. untuk langkah perintahnya seperti yang tertera pada gambar dibawah ini :<br/>

<dd> 

![gambar-19.jpg](/minggu-03/Latihan/gambar-19.jpg)

Dan apabila proses push heroku berhasil, dan melakukan perintah heroku open, maka otomatis ketika merefres pada halaman app kita, maka hasilnya seperti pada gambar dibawah ini :<br/>

![gambar-20.jpg](/minggu-03/Latihan/gambar-20.jpg)

<dd>Dimana hasil pada web local heroku dan web url app kita, akan menghasilkan hasil yang sama, seperti pada gambar diatas tersebut</dd><br/>

13. Selanjutnya untuk melihat proses runiing dari file shell manage.py, dan kemudian selanjutnya dengan menampilkan hasilnya dengan menjalankan pada bash phyton, seperti hasil dan perintahnya seperti gambar dibawah ini :<br/>

<dd>

![gambar-21.jpg](/minggu-03/Latihan/gambar-21.jpg)