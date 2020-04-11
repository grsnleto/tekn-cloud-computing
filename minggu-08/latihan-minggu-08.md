# Penjelasan Praktikum TCC Minggu Ke 08

## Latihan

1. Terlebih dulu membuat direktori baru untuk project dan kemudian masuk pada direktori tersebut, seperti pada gambar dibawah ini :<br/>
![gambar-01](/minggu-08/gambar-01.jpg)
<dd>

```markdown 
Perintah yang digunakan yakni :
1. mkdir composetest
2. cd composetest
```
</dd>

2. Selanjutnya membuat file baru pada project dengan nama app.py file python ini akan menghubungkan ke jaringan redis yang nantinya dapat diakses. Dengan menambahkan script seperti pada gambar dibawah ini :<br/>
![gambar-02](/minggu-08/gambar-02.jpg)

3. Kemudian membuat file txt dengan nama requirements dan mengisihkan script pada file tersebut, seperti dibawah ini :<br/>
![gambar-03](/minggu-08/gambar-03.jpg)

4. Kemudian membuat file baru dengan nama DockerFile dan kemudian mengisihkan script seperti dibawah ini :<br/>
![gambar-04](/minggu-08/gambar-04.jpg)
<dd>

```markdown 
Pada script diatas tersebut akan berguna untuk mengeksekusi beberapa perintah yang berada pada beberapa file sebelumnya.
```
</dd>

5. Dan kemudian untuk selanjutnya membuat file yml, dengan nama docker-compose. File ini akan berguna sebagai wadah untuk merujuk pada port 5000 yang nantinya dapat diakses. Seperti pada gambar dibawah ini :<br/>
![gambar-05](/minggu-08/gambar-05.jpg)
<dd>

```markdown 
Script diatas merupakan script uang akan menjadi layanan dengan menggunakan port default untuk server web Flask, 5000.
```
</dd>
