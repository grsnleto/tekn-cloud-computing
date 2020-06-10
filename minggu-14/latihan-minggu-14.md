# Penjelasan Praktikum TCC Minggu Ke 14

## Latihan

1. Pada bagian yang pertama, melakukan clonning repo dari git dockercoins. seperti pada gambar dibawah ini :<br/>
![gambar-01](/minggu-14/gambar-01.jpg)<br/>
2. Kemudian selanjutnya masuk pada direktori dockercoins dan kemudian menjalankan perintah docker-compose up, untuk menjalankan semua container yang ada pada dockercoins. Seperti pada gambar dibawah ini :<br/>
![gambar-02](/minggu-14/gambar-02.jpg)<br/>
3. Kemudian melihat cluster kita. semunya terkoneksi pada server yang sama atau tidak. Seperti pada gambar dibawah ini :<br/> 
![gambar-03](/minggu-14/gambar-03.jpg)<br/>
![gambar-04](/minggu-14/gambar-04.jpg)<br/>
4. Kemudian selanjutnya melakukan ping pada server image docker alpine pada 0.0.0.0 dengan menggunakan perintah seperti pada gambar dibawah ini : <br>
![gambar-05](/minggu-14/gambar-05.jpg)<br/>
5. Selanjutnya, kemudian masuk pada direktori stacks, kemudian membuka file dockercoins.yml dan kemudian melakukan pushing, seperti pada gambar dibawah ini : <br/>
![gambar-06](/minggu-14/gambar-06.jpg)<br/>
6. Kemudian selanjutnya melihat hasil deploy atau pushing pada rng dan worker, seperti pada gambar dibawah ini :<br/>
![gambar-07](/minggu-14/gambar-07.jpg)<br/>
7. Selanjutnya melihat hasil deployment pada port server redis 6379, pada server rng di port 80 dan server hasher di port 80, seperti pada gambar dibawah ini : <br/>
![gambar-08](/minggu-14/gambar-08.jpg)<br/>
8. Selanjutnya melihat logs deploy worker, kemudian melakukan cek service nodeport pada tcp 80 dan port 30001 dan menampilkan svc, seperti pada gambar dibawah ini :<br/>
![gambar-09](/minggu-14/gambar-09.jpg)<br/>
9. Kemudian melakukan export rng menjadi file yml dan sekaligus menampilkannya, seperti pada gambar dibawah ini :<br/>
![gambar-10](/minggu-14/gambar-10.jpg)<br/>
10. Dan selanjutnya melakukan running rng dan melihat describe rng dan selanjutnya melakukan update rng serta menampilkannya, seperti pada gambar dibawah ini :<br/>
![gambar-11](/minggu-14/gambar-11.jpg)<br/>


