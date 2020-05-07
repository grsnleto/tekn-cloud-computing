# Penjelasan Praktikum TCC Minggu Ke 12

## Latihan

1. Membuat container baru yang dijalankan pada node 1. Dengan menggunakan perintah seperti gambar dibawah ini :<br/>
![gambar-01](/minggu-12/gambar-01.jpg)<br/>
![gambar-02](/minggu-12/gambar-02.jpg)<br/>
dimana dari perintah diatas tersebut berguna untuk membuat container baru. Container ini akan tetap berjalan pada background image docker saat sleep. 
2. Kemudian untuk mastikan bahwa container yang baru dibuat tersebut sudah berjalan pada node 1. Dengan menggunakan perintah seperti pada gambar dibawah ini :<br/>
![gambar-03](/minggu-12/gambar-03.jpg)<br/>
3. Kemudian menginisialisasi Swarm baru kemudian join dengan single node dan melakukan verifikasi apabila operasi berhasil. Dengan menggunakan perintah seperti gambar dibawah ini :<br/>
![gambar-04](/minggu-12/gambar-04.jpg)<br/>
![gambar-05](/minggu-12/gambar-05.jpg)<br/>
Dimana dari hasil pembuatan docker swarm baru tersebut akan mendapatkan token yang nantinya digunakan untuk join node yang lain.<br/>
Selanjutnya melakukan verifikasi pada node 1 apabila telah berhasil melakukan confogurasi manager node swarm. Dengan menjalankan perintah seperti gambar dibawah ini :<br/>
![gambar-06](/minggu-12/gambar-06.jpg)<br/>
4. Kemudian melakukan join swarm untuk node 2 dan node 3 dengan menggunakan token yang didapat dari pembauatan swarm baru sebelumnya.<br/>
node 2<br/>
![gambar-07](/minggu-12/gambar-07.jpg)<br/>
node 3<br/>
![gambar-08](/minggu-12/gambar-08.jpg)<br/>
5. Setelah melakukan join node2 dan node3, kemudian kembali ke node1, dan jalankan perintah docker node ls untuk memverifikasi bahwa kedua node tersebut sudah join ke Swarm. Seperti pada gambar dibawah ini :<br/>
![gambar-09](/minggu-12/gambar-09.jpg)<br/>
![gambar-10](/minggu-12/gambar-10.jpg)<br/>
6. Kemudian dari node 1 akan dilakukan deploy sleep pada docker swarm. Deploy Ini merupakan layanan dari docker swarm sendiri.<br/>
![gambar-11](/minggu-12/gambar-11.jpg)<br/>
Kemudian verifikasi bahwa pembuatan service telah diterima oleh Swarm manager. Seperti perintah dan hasilnya pada gambar dibawah ini :<br/>
![gambar-12](/minggu-12/gambar-12.jpg)<br/>
7. Kemudian melakukan replika pada container sampai sleep-app ke 7, ini berguna untuk mendeskripsikan container pada layanan service. Seperti pada perintah dan hasilnya pada gambar dibawah ini :<br/>
![gambar-13](/minggu-12/gambar-13.jpg)<br/>
![gambar-14](/minggu-12/gambar-14.jpg)<br/>
Kemudian melihat replika container yang dibuat secara real-time.
![gambar-15](/minggu-12/gambar-15.jpg)<br/>
![gambar-16](/minggu-12/gambar-16.jpg)<br/>
![gambar-17](/minggu-12/gambar-17.jpg)<br/>
Dimana dari hasil diatas terdapat 7 container yang berjalan.
8. Melakukan replika layanan kembali menjadi hanya 4 kontainer dengan melakukan pembaruan layanan docker container. Seperti pada perintah dan hasilnya dibawah ini :<br/>
![gambar-18](/minggu-12/gambar-18.jpg)<br/>
![gambar-19](/minggu-12/gambar-19.jpg)<br/>
Dimana replika diatas sudah menghasilkan 4 container yang sedang berjalan.<br/>
Kemudian melakukan verifikasi dari hasil replika container diatas, dengan menggunakan perintah seperti pada gambar dibawah ini :<br/>
![gambar-20](/minggu-12/gambar-20.jpg)<br/>
![gambar-21](/minggu-12/gambar-21.jpg)<br/>
Sehingga dari hasilnya diatas, dimana terlihat hanya tinggal 4 container docker yang sendang berjalan pada slepp-app.
9. Kemudian mengecek status node yang sedang berjalan pada node 1. Dengan menggunakan perintah seperti pada gambar dibawah ini :<br/>
![gambar-22](/minggu-12/gambar-22.jpg)<br/>
![gambar-23](/minggu-12/gambar-23.jpg)<br/>
10. Kemudian melihat container yang sedang berjalan pada node 2.<br/>
![gambar-24](/minggu-12/gambar-24.jpg)<br/>
11. Kemudian kembali ke node 1 untuk mengambil node 2 dari layanan service. Dengan menjalankan perintah seperti pada gambar dibawah ini :<br/>
![gambar-25](/minggu-12/gambar-25.jpg)<br/>
Dari hasil diatas, id node 2 akan digunakan untuk melakukan drain untuk menggantikan node kita dengan node 2. Seperti pada gambar dibawah ini :<br/>
![gambar-26](/minggu-12/gambar-26.jpg)<br/>
Setelah melakukan itu, maka cek satatus node, dengan menggunakan perintah seperti pada gambar dibawah ini :<br/>
![gambar-27](/minggu-12/gambar-27.jpg)<br/>
12. Kemudian kembali ke node 2, dan jalankan perintah seperti gambar dibawah ini :<br/>
![gambar-28](/minggu-12/gambar-28.jpg)<br/>
Sehingga hasilnya container pada node 2 sudah tidak berjalan lagi.<br/>
13. Kemudian cek ulang pada node 1 bahwa container sudah dijadwal ulang, dengan melihat bahwa hanya ada 4 coinaer yang sedang berjalan sedangkan node 2 tidak. Seperti pada gambar dibawah ini :<br/>
![gambar-29](/minggu-12/gambar-29.jpg)<br/>
![gambar-30](/minggu-12/gambar-30.jpg)<br/>
14. Kemudian menghapus layanan service pada node 1 dengan menggunakan perintah seperti dibawah ini :<br/>
![gambar-31](/minggu-12/gambar-31.jpg)<br/>
15. Kemudian menjalankan perintah docker ps pada node 1 untuk melihat daftar container yang sedang berjalan :<br/>
![gambar-32](/minggu-12/gambar-32.jpg)<br/>
dan kemudian melakukan kill container. Seperti pada gambar dibawah ini :<br/>
![gambar-33](/minggu-12/gambar-33.jpg)<br/>
16. Kemudian hapus semua swarm container untuk node 1, node 2 dan node 3 seperti pada gambar dibawah ini :<br/>
![gambar-34](/minggu-12/gambar-34.jpg)<br/>

