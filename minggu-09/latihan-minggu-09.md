# Penjelasan Praktikum TCC Minggu Ke 09

## Latihan

1. Pada bagian yang pertama ini, akan melakukan clonning repo dari github dengan mengambil direktori linux_tweet_app. Perintah yang digunakan seperti pada gambar dibawah ini :<br/>
![gambar-01](/minggu-09/gambar-01.jpg)<br/>
2. Apabila proses clonning berhasil dilakukan, maka langkah selanjutnya akses direktori yang baru di clone dan kemudian jalankan perintah untuk menajalankan hostname melalui linux container yang berapa didalam wadah alphine. Seperti pada gambar dibawah ini :<br/>
![gambar-02](/minggu-09/gambar-02.jpg)<br/>
dari hasil output di atas ini menunjukkan bahwa image alpine:latest terbaru tidak dapat ditemukan secara lokal. Ketika ini terjadi, Docker secara otomatis menariknya dari Docker Hub. sehingga ketika image di pull, nama host kontainer ditampilkan (b4e4e25ce4f5 seperti pada gambar diatas tersebut).<br/>
3. Kemudian melihat list container yang tersedia, dengan menggunakan perintah seperti dibawah ini :<br/>
![gambar-03](/minggu-09/gambar-03.jpg)<br/>
dimana seperti yang diberi tanda bahwa cantainer dengan hostname kita sedang running.<br/>
4. Kemudian selanjutnya melakukan running container ubuntu dengan menggunakan perintah seperti dibawah ini :<br/>
![gambar-04](/minggu-09/gambar-04.jpg)<br/>
5. Selanjutnya menjalankan perintah seperti dibawah ini didalam container ubuntu :<br/>
![gambar-05](/minggu-09/gambar-05.jpg)<br/>
dimana ls / akan mencantumkan isi direktori root, ps aux akan menunjukkan proses yang berjalan, cat / etc / issue akan menunjukkan proses Linux yang sedang berjalan, dalam hal ini Ubuntu 18.04.3 LTS. <br/>
dan kemudian menjalankan perintah exit untuk keluar dari bash container seperti pada gambar dibawah ini :<br/>
![gambar-06](/minggu-09/gambar-06.jpg)<br/>
6. Kemudian selanjutnya menjalankan perintah untuk melihat host yang sedang berjalan, seperti pada gambar dibawah ini :<br/>
![gambar-07](/minggu-09/gambar-07.jpg)<br/>
7. Kemudian mejalankan perintah untuk running container MYSQL, seperti pada gambar dibawah ini :<br/>
![gambar-08](/minggu-09/gambar-08.jpg)<br/>
Sehingga image container mysql akan berjalan pada background.
8. Dan kemudian melakukan cek list untuk container yang sedang berjalan, Seperti pada gambar dibawah ini :<br/>
![gambar-09](/minggu-09/gambar-09.jpg)<br/>
dimanan pada gambar diatas dapat dilihat bahwa container image mysql sedang berjalan.<br/>
9. Kemudian melakukan cek terhadap container kita yang sedang berjalan dalam wadah, dengan menggunakan perintah seperti gambar dibawah ini :<br/>
![gambar-10](/minggu-09/gambar-10.jpg)<br/>
10. Selanjutnya melihat container yang sedang berjalan pada wadah dengan perintah seperti dibawah ini :<br/>
![gambar-11](/minggu-09/gambar-11.jpg)<br/>
sehingga dapat dilihat bahwa container mysql sedang barjalan dalam wadah.<br/>
11. Kemudian melihat versi container mysql yang sedang berjalan, dengan menggunakan perintah seperti dibawah ini :<br/>
![gambar-12](/minggu-09/gambar-12.jpg)<br/>
12. Selanjutnya melakukan koneksi container docker dengan container mysql yang sedang berjalan dan kemudian melakukan cek versi dengan menggunakan perintah shell baru dalam wadah container ini, dan kemudian menjalankan perintah exit untuk keluar dari wadah container. dengan menggunakan perintah seperti dibawah ini :<br/>
![gambar-13](/minggu-09/gambar-13.jpg)<br/>
13. Membuat situs web sederhana dengan menggunakan image. Pada bagian pertama melakukan check isi dari file Dockerfile, dengan menggunakan perintah seperti dibawah ini :<br/>
![gambar-14](/minggu-09/gambar-14.jpg)<br/>
14. Selanjutnya mengunakan Docker ID kita untuk nantinya bisa membuat image docker baru dan selanjutnya menjalankan perintan docker image build, seperti perintah yang digunakan pada gambar dibawah ini :<br/>
![gambar-15](/minggu-09/gambar-15.jpg)
![gambar-16](/minggu-09/gambar-16.jpg)<br/>
15. Dan kemudian menggunakan perintah seperti dibawah ini untuk menjalankan image container docker kita, seperti pada gambar dibawah ini :<br/>
![gambar-17](/minggu-09/gambar-17.jpg)<br/>
apabila proses tersebut berhasil, maka selanjutnya image docker kita dapat diakses di browser, dengan mengklik link seperti gambar dibawah ini : <br/>
![gambar-18](/minggu-09/gambar-18.jpg)<br/>
Sehingga tampilan web sederhana kita dari image docker yang dibuat, seperti pada gambar dibawah ini :
![gambar-19](/minggu-09/gambar-19.jpg)<br/>
16. Apabila sudah melakukan akses terhadap situs web sederhana kita, maka selanjutnya mencoba untuk menonaktifkan dan remove sementara web kita, dengan menggunakan perintah seperti dibawah ini :<br/>
![gambar-20](/minggu-09/gambar-20.jpg)<br/>
Sehingga apabila kita mengakses ulang web sederhana kita akan menampilkan halaman, seperti pada gambar dibawah ini :
![gambar-21](/minggu-09/gambar-21.jpg)<br/>
ini dikarenakan image docker yang dibuat menjadi web sederhana tadi sudah dinonaktifkan dan di remove sementara.<br/>
17. Kemudian melakukan modifikasi web kita dengan terlebih dahulu menjalankan web sederhana kita kembali, dikarenakan sebelumnya sudah dilakukan nonaktif dan di remove sementara, perintah yang digunakan seperti pada gambar dibawah ini :<br/>
![gambar-22](/minggu-09/gambar-22.jpg)<br/>
18. Kemudian selanjutnya copy file index.html pada container kita, untuk perintahnya seperti pada gambar dibawah ini :<br/>
![gambar-23](/minggu-09/gambar-23.jpg)<br/>
Kemudian lakukan refresh pada halaman web sederhana kita, dan akan menampilkan seperti gambar dibawah ini :<br/>
![gambar-24](/minggu-09/gambar-24.jpg)<br/>
19. Kemudian memberentikan container update kita dan kemudian menjalankan container versi sebelumnya yaitu 1.0. untuk perintahnya seperti pada gambar dibawah ini :<br/>
![gambar-25](/minggu-09/gambar-25.jpg)<br/>
sehingga apabila kita melakukan refresh ulang pada website kita, akan menampilkan halaman pada versi sebelumnya, seperti gambar dibawah ini :<br/>
![gambar-26](/minggu-09/gambar-26.jpg)<br/>
apabila sudah, maka remove dan nonaktifkan sementara kembali web kita. untuk perintahnya seperti dibawah ini :<br/>
![gambar-27](/minggu-09/gambar-27.jpg)<br/>
20. Kemudian melakukan update image. Dengan membuat image baru dengan versinya yaitu 2.0. Perintah yang digunakan seperti pada gambar dibawah ini :
![gambar-28](/minggu-09/gambar-28.jpg)<br/>
21. Selanjutnya melihat list image yang berjalan, dengan menggunakan perintah seperti dibawah ini :<br/>
![gambar-29](/minggu-09/gambar-29.jpg)<br/>
dimana dari gambar diatas sudah host untuk versi terbaru yang baru saja diupdate yaitu 2.0.<br/>
22. Selanjutnya melakukan test untuk versi terbaru image kita, perintah yang digunakan seperti gambar dibawah ini :<br/>
![gambar-30](/minggu-09/gambar-30.jpg)<br/>
dan kemudian tampilan untuk web versi terbaru kita yaitu 2.0, seperti pada gambar dibawah ini :<br/>
![gambar-31](/minggu-09/gambar-31.jpg)<br/>
23. Kemudian apabila kita ingin menjalakan web versi sebelumnya yaitu 1.0, dengan menggunakan perintah seperti dibawah ini :<br/>
![gambar-32](/minggu-09/gambar-32.jpg)<br/>
dan tampilan pada browser ketika di open, seperti pada gambar dibawah ini :<br/>
![gambar-33](/minggu-09/gambar-33.jpg)<br/>
24. Kemudian selanjutnya melakukan push image kita ke docker hub. untuk pertama melihat list image pada docker host kita, untuk perintahnya seperti pada gambar dibawah ini :<br/>
![gambar-34](/minggu-09/gambar-34.jpg)<br/>
dimana terdapat 2 image kita yaitu versio 1.0 dan 2.0, kedua ini akan nantinya di push ke docker hub.
25. Untuk melakukan push ke docker hub, sebelumnya harus melakukan login dengan akun docker kita, untuk perintahnya seperti pada gambar dibawah ini :<br/>
![gambar-35](/minggu-09/gambar-35.jpg)<br/>
login diatas menggunakan ID kita dan Password kita. Apabila sudah melakukan login maka selanjutnya melakukan push image docker kita untuk versi 1.0 dan versi 2.0, dengan menggunakan perintah seperti dibawah ini :<br/>
![gambar-36](/minggu-09/gambar-36.jpg)<br/>
Apabila proses push image docker kita berhasi, maka selanjutnya akses url https://hub.docker.com/r/175610055/. Ini bermaksud untuk melihat dan memastikan hasil image docker yang push tadi, seperti pada gambar dibawah ini :<br/>
![gambar-37](/minggu-09/gambar-37.jpg)<br/>
dimana dari gambar diatas tersebut 2 image docker berhasi di push dan ditampilkan pada container docker saya.
