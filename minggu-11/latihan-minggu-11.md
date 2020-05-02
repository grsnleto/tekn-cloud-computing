# Penjelasan Praktikum TCC Minggu Ke 11

## Latihan
1. Melakukan clonning repo dari git dan kemudian mengakases derektori serta memeriksa branch dari repo tersebut.<br/>
![gambar-01](/minggu-11/gambar-01.jpg)<br/>
2. Selanjutnya memeriksa branch step0 untuk melihat file didalamnya. Dan kemudian melihat isi dari file python linkextractor.<br/>
![gambar-02](/minggu-11/gambar-02.jpg)<br/>
3. Menajalankan file linkextractor.py.<br/>
![gambar-03](/minggu-11/gambar-03.jpg)<br/>
4. Memeriksa branch step1 untuk melihat file didalamnya.<br/>
![gambar-04](/minggu-11/gambar-04.jpg)<br/>
5. Kemudian melihat isi dari file Dockerfile.<br/>
![gambar-05](/minggu-11/gambar-05.jpg)<br/>
6. Membanun docker image kita. Dimana dengan menjalankan perintah seperti dibawah ini, dan sekaligus manghasilkan outputnya.<br/>
![gambar-06](/minggu-11/gambar-06.jpg)<br/>
![gambar-07](/minggu-11/gambar-07.jpg)<br/>
7. Kemudian setelah membuat docker image bernama linkextractor: step1 selanjutnya kita mengecek dengan melihat daftar/ist dari image docker yang sudah ada. Kemudian melakukan ekstrak docker image kita dengan mendapat URL.<br/>
![gambar-08](/minggu-11/gambar-08.jpg)<br/>
Dan ini merupakan output apabila kita mengakses alamat URL tersebut.
![gambar-09](/minggu-11/gambar-09.jpg)<br/>
8. Selanjutnya melakukan percobaan pada halaman web untuk melihat lebih banyak tautannya.<br/>
![gambar-10](/minggu-11/gambar-10.jpg)<br/>
9. Kemudian memeriksa branch step2 dan daftar file yang ada di dalamnya.<br/>
![gambar-11](/minggu-11/gambar-11.jpg)<br/>
10. Sehingga secara otomatis file linkextractor.py akan diupdate, berikut merupakan hasil update dari file tersebut.<br/>
![gambar-12](/minggu-11/gambar-12.jpg)<br/>
11. Kemudian membuat image baru.<br/>
![gambar-13](/minggu-11/gambar-13.jpg)<br/>
12. Sehingga image docker yang baru dibuat dengan nama linkextractor:step02, kita cek pada pada list image docker yang sudah ada.<br/>
![gambar-14](/minggu-11/gambar-14.jpg)<br/>
13. Kemudian menjalankan image docker tersebut dan menghasilkan keluaran seperti dibawah ini :<br/>
![gambar-15](/minggu-11/gambar-15.jpg)<br/>
14. Kemudian menjalankan image docker step1 yang sebelumnya dan menghasilkan keluaran yang masih sama seperti dibawah ini :<br/>
![gambar-16](/minggu-11/gambar-16.jpg)<br/>
15. Selanjutnya memeriksa branch step3 dan isi file didalamnya.<br/>
![gambar-17](/minggu-11/gambar-17.jpg)<br/>
16. Kemudian cek file Dockerfile untuk melihat perubahannya.<br/>
![gambar-18](/minggu-11/gambar-18.jpg)<br/>
17. Selanjutnya melihat isi dari file main.py yang baru ditambahkan tersebut.<br/>
![gambar-19](/minggu-11/gambar-19.jpg)<br/>
18. Kemudian update image docker step3 ini dengan beberapa langkah perubahan, seperti dibawah ini.<br/>
![gambar-20](/minggu-11/gambar-20.jpg)<br/>
![gambar-21](/minggu-11/gambar-21.jpg)<br/>
19. Selanjutnya menjalankan container dalam mode (-d flag) sehingga terminal dapat tersedia untuk perintah yang lain saat container masih berjalan. Perhatikan juga bahwa disitu terdapat port 5000 dari container dengan host 5000 (menggunakan perintah -p 5000: 5000) agar dapat diakses dari host. Dan juga memberikan nama (--name = linkextractor) ke container untuk lebih mudah melihat log atau menghapus container. Serta melihat list image container yagg baru dibuat tersebut.<br/>
![gambar-22](/minggu-11/gambar-22.jpg)<br/>
20. Membuat permintaan HTTP dalam bentuk /api /url untuk mengakses server ini dan mengambil respons berisi link yang diekstrak.<br/>
![gambar-23](/minggu-11/gambar-23.jpg)<br/>
21. Karena container berjalan dalam mode terpisah, jadi tidak dapat melihat proses yang terjadi di dalam, tetapi dapat melihat log menggunakan linkextractor yang di tetapkan untuk container. Serta menghapus image container ini.<br/>
![gambar-24](/minggu-11/gambar-24.jpg)<br/>
22. Selanjutnya memeriksa branch step4 dan isi file didalamnya.<br/>
![gambar-25](/minggu-11/gambar-25.jpg)<br/>
23. Kemudian melihat isi dari file docker-compose.yml dan www / index.php.<br/>
![gambar-26](/minggu-11/gambar-26.jpg)<br/>
![gambar-27](/minggu-11/gambar-27.jpg)<br/>
24. Membuat mode terpisah untuk container.<br/>
![gambar-28](/minggu-11/gambar-28.jpg)<br/>
![gambar-29](/minggu-11/gambar-29.jpg)<br/>
25. Memeriksa daftar container yang sedang berjalan serta memastikan bahwa kedua container tersebut benar berjalan. Dan kemudian mengkases layanan API.<br/>
![gambar-30](/minggu-11/gambar-30.jpg)<br/>
Sehingga ketika megkasesnya pada URL, maka hasilnya seperti pada gambar dibawah ini :<br/>
![gambar-31](/minggu-11/gambar-31.jpg)<br/>
26. Selanjutnya memodifikasi file www/index.php mengganti semua kemunculan Link Extractor dengan Super Link Extractor. Kemudian mengembalikan perubahan, dan menonaktifkan container ini.<br/>
![gambar-32](/minggu-11/gambar-32.jpg)<br/>
27. Kemudian memeriksa branch step5 dan melihat isi file didalamnya.<br/>
![gambar-33](/minggu-11/gambar-33.jpg)<br/>
28. Memeriksa file Dockerfile yang baru di dalam direktori www.<br/>
![gambar-34](/minggu-11/gambar-34.jpg)<br/>
29. Selanjutnya,  melihat isi file api/main.py dengan menggunakan server redis.<br/>
![gambar-35](/minggu-11/gambar-35.jpg)<br/>
30. Melihat hasil perubahan pada file docker-compose.yml.<br/>
![gambar-36](/minggu-11/gambar-36.jpg)<br/>
31. Selanjutnya melakukan eksekusi container ini, untuk bisa di buka pada browser.<br/>
![gambar-37](/minggu-11/gambar-37.jpg)<br/>
![gambar-38](/minggu-11/gambar-38.jpg)<br/>
Sehingga hasilnya ketika ditampilkan pada browser, seperti pada gambar dibawah ini :<br/>
![gambar-31](/minggu-11/gambar-31.jpg)<br/>
32. Selanjutnya memeriksa apakah layanan redis dipakai atau tidak.<br/>
![gambar-39](/minggu-11/gambar-39.jpg)<br/>
33. Memeriksa ketika folder www tidak tersedia pada container yang sedang berjalan. kemudian melakukan verifikasi bahwa perubahan yang dibuat secara lokal tidak berada dalam layanan yang berjalan dengan memuat ulang halaman web dan kemudian mengembalikan perubahan. Dan menonaktifkan cointainer ini. <br/>
![gambar-40](/minggu-11/gambar-40.jpg)<br/>
34. Kemudian memeriksa branch step6 dan melihat isi file didalamnya.<br/>
![gambar-41](/minggu-11/gambar-41.jpg)<br/>
35. Selanjutnya melihat isi file linkextractor.rb, ini merupakan file ruby untuk mengelola dependency. Dan kemudian juga melihat isi file Dockerfile.<br/>
![gambar-42](/minggu-11/gambar-42.jpg)<br/>
![gambar-43](/minggu-11/gambar-43.jpg)<br/>
36. Kemudian mengecek perubahan pada file docker-compose.yml.<br/>
![gambar-44](/minggu-11/gambar-44.jpg)<br/>
37. Kemudian selanjutnya mengeksekusi container ini.<br/>
![gambar-45](/minggu-11/gambar-45.jpg)<br/>
![gambar-46](/minggu-11/gambar-46.jpg)<br/>
Dan kemudian selanjutnya harus dapat mengakses API (menggunakan nomor port yang diperbarui):<br/>
![gambar-47](/minggu-11/gambar-47.jpg)<br/>
Sehingga hasil yang ditampilkan pada halaman browser seperti dibawah ini :<br/>
![gambar-31](/minggu-11/gambar-31.jpg)<br/>
38. Kemudian shut down container ini tetapi lognya akan tetap ada walaupun containernya hilang.<br/>
![gambar-48](/minggu-11/gambar-48.jpg)<br/>



