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

6. Kemudian selanjutnya melakukan running untuk project/direktori yang dibuat tadi dengan menggunakan docker compose, yang perintahnya yaitu <b>docker-compose up</b>, seperti pada gambar dibawah ini :<br/>
![gambar-06](/minggu-08/gambar-06.jpg)
![gambar-07](/minggu-08/gambar-07.jpg)
<dd>

```markdown 
Ketika perintah docker-compose up dijalankan maka otomatis sever akan runnig, dimana setiap proses running dilakukan maka akan otomatis merujuk pada image docker dan kemudian merequest ke server redis.
```
Sehingga hasilnya, apabila di open pada browser dengan mengakses server host yakni <b>192.168.99.100:5000</b>. Host ini fungsinya sama seperti penggunaan pada docker dekstop yakni <b>localhost:5000</b>, karena disini saya menggunakan docker tollbox, maka host yang digunakan seperti diatas tersebut.<br/>
![gambar-08](/minggu-08/gambar-08.jpg) 
</dd>

7. Kemudian apabila melakukan refresh pada browser maka secara otomatis akan di <b>count</b> dimana setiap kali request dilakukan maka nilainya akan bertambah, artinya akan menghitung jumlah riquest time dari client. Seperti pada gambar dibawah ini :<br/>
![gambar-09](/minggu-08/gambar-09.jpg) 

8. Selanjutnya mengeksekusi perintah untuk melihat daftar image yang ada, dengan meggunakan perintah <b>docker image ls</b>, seperti pada gambar dibawah ini :<br/>
![gambar-10](/minggu-08/gambar-10.jpg) 
<dd>

```markdown 
Dimana dari gambar diatas repo/image yang digunakan sekarang yaitu python dan redis.
```
</dd>

9. Kemudian selanjutnya melakukan perubahan pada file .yml yang ada, dengan menambahkan script didalamnya seperti pada gambar dibawah ini :<br/>
![gambar-11](/minggu-08/gambar-11.jpg) 
<dd>

```markdown 
Dimana perubahan dengan menambahkan dua variabel yaitu volumes dan environtment. 
```
</dd>

10. Kemudian selanjutnya melakukan running ulang dengan perintah docker-compose up, dan disitu terlihat perubahan, seperti pada gambar dibawah ini :<br/>
![gambar-12](/minggu-08/gambar-12.jpg) 
<dd>

```markdown 
Dimana hasil update pada file yml akan ditampilkan ketika proses runnig ulang ke host seperti pada gambar diatas tersebut. 
```
</dd>

11. Kemudian melakukan update pada file app.py dengan merubah pesan seperti pada gambar dibawah ini :<br/>
![gambar-13](/minggu-08/gambar-13.jpg) 

12. Dan selajutnya melakukan eksperimen pada beberapa perintah docker yang ada. Dimana beberapa perintah ini akan melihat container yang sedang running pada backgroung, melakukan stop running container, melihat running pada server web dan volume container, seperti pada gambar dibawah ini :<br/>
![gambar-14](/minggu-08/gambar-14.jpg) 