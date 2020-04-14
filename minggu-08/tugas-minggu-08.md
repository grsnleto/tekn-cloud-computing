# Penjelasan Praktikum TCC Minggu Ke 08

## Tugas

![gambar-15](/minggu-08/gambar-15.jpg)

Dari diagram yang saya rancang diatas, docker compose dapat menerima request dari file yml sehingga feedback yang diberikan dengan membaca file yml dan mengirim ke manage docker client, dimana cli dapat melakukan deklarasi build dan kemudian melakukan running dan kemudian dapat melakukan pull terhadap image container nanti. Apabila proses manage pada docker client melakukan request maka terlebih dahulu melalaui docker swarm untuk melakukan pengelompokan terhadap request yang diberikan dan kemudkan akan dilanjutkan ke dockerd/docker daemond, dimana pada docker dameond image dan container saling mengirim dan menerima feedback untuk selanjutnya memberikan feedback kepada docker cli untuk melakukan manage ulang terhadap kesalahan dll.