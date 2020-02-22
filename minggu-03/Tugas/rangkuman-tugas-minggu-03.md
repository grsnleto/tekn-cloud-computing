# Penjelasan Praktikum TCC Minggu-03
<br/>

# Tugas
<br/>

1. Untuk proses berikutnya akan melakukan login pada heroku, disini saya menggunakan command git-bash untuk melakukan penagaksesan heroku. berikut merupakan perintahnya :<br/>
<dd>

![gambar-22](/minggu-03/Tugas/gambar-22.jpg)

Dari perintah diatas apabila sudah mengisi perintah "heroku login", maka secara otomatis akan lansung terhubung ke web heroku yaitu halaman login. Seperti pada gambar dibawah ini :<br/>

![gambar-23](/minggu-03/Tugas/gambar-23.jpg)</dd>

2. Selanjutnya mengecek versi dari php, composer dan git dari komputer local yang sebelumnya sudah diinstal, untuk perintah dan hasil eksekusinya seperti pada gambar dibawah ini :<br/>
<dd>

![gambar-24](/minggu-03/Tugas/gambar-24.jpg)</dd>

3. Pada bagian ke-3 ini dilakukan proses clonning dari git heroku untuk mengambil repo "php-getting-started" ke bash lokal komputer dan masuk ke repo lokal yang baru di clonning, dan kemudian selanjutnya membuat app bauru kita pada heroku. untuk perintahnya seperti pada gambar yang tertera dibawah ini :<br/>
<dd>

![gambar-05.jpg](/minggu-03/Tugas/gambar-25.jpg)</dd>

4. Kemudian selanjutnya melakukan push ke branch master heroku dengan menggunakan perintah "git push heroku master". Untuk perintah dan hasilnya seperti pada gambar dibawah ini :<br/>
<dd>

![gambar-26](/minggu-03/Tugas/gambar-26.jpg)</dd>

5. Kemudian selanjutnya melakukan konfiguasi pada web scale dengan nilai konfigurasinya 1, apabila 1 maka app yang di create pada heroku dapat diakses, dan kemudian untuk membukanya dengan menggunakan perintah "heroku open". seperti pada gambar dibawah ini :<br/>
<dd>

![gambar-28.jpg](/minggu-03/Tugas/gambar-28.jpg)</dd>

<dd>Sehingga hasil yang di tampilkan pada web browser seperti dibawah ini :<br/>

![gambar-27](/minggu-03/Tugas/gambar-27.jpg)</dd>

6. Kemudian selanjutnya melakukan update composer, dari bash commmand git, untuk perintah dan hasilnya seperti yang terlihat pada gambar dibawah ini :<br/>
<dd>

![gambar-29](/minggu-03/Tugas/gambar-29.jpg)

Sehingga hasil yang di tampilkan pada repo lokal yang di conning sebelumnya, dimana pada direktori vendor akan terdapat file autoload.php dari hasil update composer tersebut :<br/>

![gambar-30](/minggu-03/Tugas/gambar-30.jpg)</dd>

7. Selanjutnya dari hasil update composer tersebut masih memerlukan direktori arlik11es dan didalamnya akan di install atau ditambahkan direktori cowsayphp, dan kemudian baru dilakukan update composernya lagi. Untuk perintah dan hasilnya seperti pada gambar dibawah ini :<br/>
<dd>

![gambar-31](/minggu-03/Tugas/gambar-31.jpg)

Sehingga hasil yang di tampilkan pada repo lokal yang di conning sebelumnya, dimana pada direktori vendor akan terdapat direktori arlik11es/cowsayphp dari hasil update composer tersebut :<br/>

![gambar-32](/minggu-03/Tugas/gambar-32.jpg)</dd>

8. Kemudian selanjutnya melakukan perubahan pada file index.php yang terdapat pada direktori web pada repo lokal hasil clonning sebelumnya. untuk prosesnya seperti pada gambar dibawah ini : <br/>
<dd>

![gambar-33](/minggu-03/Tugas/gambar-33.jpg)</dd>

9. Dan selanjutnya dilakukan push pada app heroku kita, dengan langkah perintah seperti yang tertera pada gambar dibawah ini :<br/>
<dd>

![gambar-34](/minggu-03/Tugas/gambar-34.jpg)

Dan kemudian untuk melihat hasil dari push tersebut, dengan menggunakan perintah "git open cowsay", seperti pada gambar dibawah ini :<br/>

![gambar-35](/minggu-03/Tugas/gambar-35.jpg)

Sehingga hasil yang ditapilkan pada halaman url app heroku kita, seperti pada gambar dibawah ini :<br/>

![gambar-36](/minggu-03/Tugas/gambar-36.jpg)</dd>

10. Selanjutnya akan dilakukan pengecekan dengan menjalankan pada bash lokal heroku, apabila berhasil maka proses running app kita akan dilihat disitu, seperti pada gambar dibawah ini :<br/>
<dd>

![gambar-37](/minggu-03/Tugas/gambar-37.jpg)</dd>