# Penjelasan Praktikum TCC Minggu Ke 13

## Latihan

1. Membuat cluster baru kubernetes.<br/>
![gambar-01](/minggu-13/gambar-01.jpg)<br/>
Dari gambar diatas dimana jika ingin membuat sebuah cluster kubernetes langkah pertama harus menjelankan minikube version. Selanjutnya baru menjalankan cluster kubernetes dengan minikube, perintahnya seperti gambar diatas.
2. Melihat versi cluster yang baru dibuat.<br/>
![gambar-02](/minggu-13/gambar-02.jpg)<br/>
3. Kemudian melihat detail cluster kubernetes yang dibuat dan melihat node dari cluster tersebut.<br/>
![gambar-03](/minggu-13/gambar-03.jpg)<br/>
4. Mendeploy app.<br/>
Yang pertama perlu kita ketahui apakah kubectl terkonfigurasi dengan sluter kita atau tidak.<br/>
![gambar-04](/minggu-13/gambar-04.jpg)<br/>
5. Kemudian setelah itu menampilkan node dari cluster kita kembali.<br/>
![gambar-05](/minggu-13/gambar-05.jpg)<br/>
6. Kemudian melakukan deploy app kita dan melihat list hasil deploy app kita, dengan menjalankan perintah seperti dibawah ini.<br/>
![gambar-06](/minggu-13/gambar-06.jpg)<br/>
7. Kemudian membuka  halaman terminal baru untuk menjalankan server proxi.<br/>
![gambar-07](/minggu-13/gambar-07.jpg)<br/>
Dimana dari gambar diatas proxi kita berjalan pada host 127.0.0.1:8001 atau localhost:8001.
8. Selanjutnya melihat API yang dihosting pada server proxy kita dengan menggunakan versi API yang didapatkan yaitu curl.<br/>
![gambar-09](/minggu-13/gambar-09.jpg)<br/>
9. Kemudian membuat variabel POD_NAME untuk menyimpan nama POD kita.<br/>
![gambar-10](/minggu-13/gambar-10.jpg)<br/>
10. Explore aplikasi.<br/>
Langkah pertama melakukan cek configurasi aplikasi kita yang sudah dideploy dan selanjutnya melihat deskripsi container dan image apa yang berjalan pada pod kita.<br/>
![gambar-11](/minggu-13/gambar-11.jpg)<br/>
11. Kemudian menjalankan kembali server proxi pada terminal baru.<br/>
![gambar-07](/minggu-13/gambar-07.jpg)<br/>
12. Kemudian kita akan mendapatkan nama Pod dengan merequest pod melalui proxy.<br/>
![gambar-12](/minggu-13/gambar-12.jpg)<br/>
13. Kemudian untuk melihat hasil aplikasi kita, perlu kita request ke curl dan mengambil log untuk container pod kita.<br/>
![gambar-13](/minggu-13/gambar-13.jpg)<br/>
14. Setelah pod aktif dan running, maka selanjutnya menjalankan perintah dengan menggunakan pod sebagai parameter pada container untuk membuat environment variabel.<br/>
![gambar-14](/minggu-13/gambar-14.jpg)<br/>
15. Selanjutnya memulai bash kita pada container pod kita dan sekaligus melakukan cek file yang menyimpan source code app NodeJS kita pada container.<br/>
![gambar-15](/minggu-13/gambar-15.jpg)<br/>
16. Selanjutnya melakukan cek curl apakah app kita berjalan didalamnya. Dan kemudian exit untuk keluar dari bash container.<br/>
![gambar-16](/minggu-13/gambar-16.jpg)<br/>
17. Membuat service baru.<br/>
Langkah pertama melakukan cek configurasi aplikasi kita yang sudah dideploy, kemudian mendaftarkan/membuat service kita, kemudian deploy service kita dan melihat hasil deploy service kita.<br/>
![gambar-17](/minggu-13/gambar-17.jpg)<br/>
18. Selanjutnya melihat keseluruhan dari service kita.<br/>
![gambar-18](/minggu-13/gambar-18.jpg)<br/>
19. Selanjutnya membuat environment variabel dan disimpan pada node port dan melakukan cek curl apakah app kita berjalan didalamnya.<br/>
![gambar-19](/minggu-13/gambar-19.jpg)<br/>
20. Kemudian selanjutnya melihat label keseluruhan dari hasil deploy kita.<br/>
![gambar-20](/minggu-13/gambar-20.jpg)<br/>
21. Kemudian menjalankan pod kita dengan menggunakan parameter dari label pod kita, kemudian membuat nama pod pada environment variable POD_NAME serta membuat label baru.<br/>
![gambar-21](/minggu-13/gambar-21.jpg)<br/>
22. Kemudian melihat deskripsi dari pod sudah di label.<br/>
![gambar-22](/minggu-13/gambar-22.jpg)<br/>
![gambar-23](/minggu-13/gambar-23.jpg)<br/>
23. Kemudian melihat daftar pod dengan menggunakan label baru kita.<br/>
![gambar-24](/minggu-13/gambar-24.jpg)<br/>
24. Kemudian menghapus service kita dan kemudian melihat kembali daftar service kita serta konfirmasi dari hasil romoved service kita.<br/>
![gambar-25](/minggu-13/gambar-25.jpg)<br/>
Dari hasil diatas dapat dipastikan app kita tidak dapat berjalan jika berada diluar cluster.
25. Untuk itu melakukan konfirmasi untuk mengetahui apakah app kita masih berjalan pada pod.<br/>
![gambar-26](/minggu-13/gambar-26.jpg)<br/>
26. Scalling a deployment.<br/>
Langkah pertama dengan melihat list dari deploy kita, kemudian melihat replika set yang dibuat auto pada saat deploy, kemudian scale deploy kita untuk membuat 4 replika. Selanjutnya lihat kembali list deploy kita.<br/>
![gambar-27](/minggu-13/gambar-27.jpg)<br/>
27. Selanjutnya cek pada pod kita apakah terdapat perubahan pada nomor list pod kita.<br/>
![gambar-28](/minggu-13/gambar-28.jpg)<br/>
28. Kemudian melihat descripsi dari 4 replika pada hasil deploy scale kita.<br/>
![gambar-29](/minggu-13/gambar-29.jpg)<br/>
29. Selanjutnya melihat deskripsi service yang telah di load, Kemudian membuat kembali environment variabel dan disimpan pada NODE_PORT dengan nilainya yaitu node_port serta memastikan apakah NODE_PORT sedang berjalan atau tidak.<br/>
![gambar-30](/minggu-13/gambar-30.jpg)<br/>
30. Kemudian membuat scale down untuk service kita dengan membuat replika menjadi 2, kemudian melakukan cek perubahan dari hasil replika pada deploy kita dan kemudian cek status replika pada pod kita.<br/>
![gambar-31](/minggu-13/gambar-31.jpg)<br/>
dimana pada gambar status replika kita 2 diantaranya terminating menjadi layanan service.
31. Update App.<br/>
Langkah pertama melihat list deploy dan kemudian melihat list pods kita.<br/>
![gambar-32](/minggu-13/gambar-32.jpg)<br/>
32. Selanjutnya melihat deskripsi pods kita.<br/>
![gambar-33](/minggu-13/gambar-33.jpg)<br/>
33. Kemudian update image kita dengan mendeploy name untuk image version baru kita dan kemudian melihat list pods kita kembali.<br/>
![gambar-34](/minggu-13/gambar-34.jpg)<br/>
34. Kemudian melakukan verifikasi hasil update kita dengan sebelumnya melihat deskripsi service kita.<br/>
![gambar-35](/minggu-13/gambar-35.jpg)<br/>
35. Kemudian membuat kembali environment variabel dan disimpan pada NODE_PORT dengan nilainya yaitu node_port serta memastikan apakah NODE_PORT sedang berjalan atau tidak dan melihat status deploy app kita sduah tersedia atau belum.<br/>
![gambar-36](/minggu-13/gambar-36.jpg)<br/>
36. Selanjutnya melihat deskripsi pods kita kembali.<br/>
![gambar-37](/minggu-13/gambar-37.jpg)<br/>
![gambar-38](/minggu-13/gambar-38.jpg)<br/>
37. Kemudian melakukan deploy image yang sduah diupdate menjadi versi 10, selanjutnya melihat list dari hasil deploy image versi 10 tersebut, dan kemudian melakukan cek list untuk pods kita apakah terdapat perubahan atau tidak.<br/>
![gambar-39](/minggu-13/gambar-39.jpg)<br/>
38. Selanjutnya melihat deskripsi pods kita.<br/>
![gambar-40](/minggu-13/gambar-40.jpg)<br/>
Disitu masih terdapat warning error karena pods sebelumnya sudah di rollout. Sehingga harus undo rollout tersebut. Dan kemudian baru melakukan cek list pods kita kembali.<br/>
![gambar-41](/minggu-13/gambar-41.jpg)<br/>
39. Dan terakhir melihat deskripsi pods kita, untuk melihat perubahan pada pods kita.<br/>
![gambar-42](/minggu-13/gambar-42.jpg)<br/>
Sehingga hasilnya tidak terdapat warning error, dikarenakan rollout pada deploy kita sudah di undo.









