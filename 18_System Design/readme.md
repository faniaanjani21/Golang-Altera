# Diagram Design

## Karakteristik Sistem Terdistribusi

Sistem Terdistribusi memiliki karakteristik yaitu:

- Terdiri dari beberapa komputer yang bekerja sama untuk menyelesaikan tugas
- Terdapat komunikasi dan koordinasi antar komponen
- Terdapat kemampuan untuk melakukan scaling secara horizontal atau vertikal

## Horizontal Scaling vs Vertical Scaling

- Horizontal Scaling: Menambah jumlah instance pada suatu aplikasi dengan tujuan membagi beban kerja
- Vertical Scaling: Menambah kapasitas pada suatu instance, misalnya dengan menambah RAM atau CPU

## Job/Work Queue

Job Queue atau Work Queue adalah mekanisme untuk menyelesaikan tugas secara terstruktur. Tugas-tugas tersebut ditempatkan pada antrian, dan instance yang tersedia dapat mengambil tugas dari antrian tersebut untuk dieksekusi.

## Load Balancing

Load Balancing adalah mekanisme untuk membagi beban kerja pada beberapa instance agar tidak terjadi overload pada suatu instance. Load Balancer akan memantau beban kerja pada setiap instance dan memutuskan instance mana yang akan menerima permintaan selanjutnya.

## Monolitik vs Mikroservice

- Monolitik: Aplikasi yang dibangun sebagai satu kesatuan yang besar, biasanya sulit untuk melakukan scaling secara horizontal
- Mikroservice: Aplikasi yang dibangun sebagai kumpulan service yang terpisah, memungkinkan scaling secara horizontal

## SQL vs NoSQL

- SQL: Database yang memanfaatkan Structured Query Language sebagai bahasa query, cocok untuk data terstruktur
- NoSQL: Database yang tidak memanfaatkan Structured Query Language, cocok untuk data yang tidak terstruktur

## Caching

Caching adalah mekanisme untuk menyimpan data sementara pada suatu instance, sehingga permintaan selanjutnya tidak perlu memuat data dari sumber asli. Hal ini dapat meningkatkan performa dan mengurangi beban kerja pada sumber asli.

## Database Indexing

Indexing adalah mekanisme untuk mempercepat pencarian data pada suatu database dengan membuat indeks pada kolom-kolom yang sering digunakan.

## Database Replication

Database Replication adalah mekanisme untuk membuat salinan data dari suatu database pada instance lain, sehingga terdapat cadangan data jika terjadi kegagalan pada suatu instance.

## Tools Diagram

Beberapa tools yang dapat digunakan untuk membuat diagram:

- SmartDraw
- Lucidchart
- Draw.io
- Visio
- Whimsical

## Software Design

Beberapa jenis design yang umum digunakan adalah:

- Flowchart
- Use Case
- Entity Relationship Diagram (ERD)
- High-Level System Design (HLS)

## Pengertian dan Konsep Scalability

Scalability adalah kemampuan suatu sistem untuk meningkatkan kinerjanya dengan menambah sumber daya, baik secara horizontal atau vertikal. Scalability juga berhubungan dengan kemampuan sistem untuk merespon perubahan beban kerja atau data dengan cepat dan efektif. Pemilihan arsitektur hardware/software yang tepat juga sangat penting dalam scalability.

## Kelayakan, Ketersediaan, dan Efisiensi

Kelayakan (Reliability), Ketersediaan (Availability), dan Efisiensi (Efficiency) adalah tiga hal yang sering dibahas dalam lingkungan TI dan layanan digital. Kelayakan mengacu pada daya tahan sistem terhadap gangguan atau masalah teknis, ketersediaan mengacu pada responsivitas sistem dalam menyediakan layanan, dan efisiensi mengacu pada penggunaan sumber daya yang ada untuk menghasilkan solusi dan output yang diinginkan.
