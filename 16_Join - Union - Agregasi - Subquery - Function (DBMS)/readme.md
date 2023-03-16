# Materi Database Schema dan Data Definition Language

## Pengertian Database

Database adalah sekumpulan data yang terorganisir dalam suatu sistem yang dapat diakses dan dikelola dengan mudah.

## Relationship Database

Relationship Database adalah hubungan antara tabel-tabel dalam suatu database relasional. Terdapat tiga jenis hubungan dalam database relasional, yaitu:

### 1. One-to-One

Hubungan antara dua tabel dalam basis data relasional di mana satu rekaman dalam tabel pertama hanya berkorespondensi dengan satu rekaman dalam tabel kedua, dan sebaliknya.

Contoh: Tabel "customer" dan "customer_address".

### 2. One-to-Many

Hubungan antara dua tabel dalam basis data relasional di mana satu rekaman dalam tabel pertama dapat berkorespondensi dengan banyak rekaman dalam tabel kedua. Namun, setiap rekaman di tabel kedua hanya dapat berkorespondensi dengan satu rekaman di tabel pertama.

Contoh: Hubungan antara tabel "orders" dan "order_details".

### 3. Many-to-Many

Hubungan antara dua tabel dalam basis data relasional di mana banyak rekaman dalam tabel pertama dapat berkorespondensi dengan banyak rekaman dalam tabel kedua, dan sebaliknya.

Contoh: Hubungan antara tabel "students" dan "courses".

## Entity Relationship Diagram

Entity Relationship Diagram (ERD) adalah sebuah diagram yang digunakan untuk merepresentasikan hubungan antara entitas dalam suatu database.

## SQL Statement

SQL (Structured Query Language) adalah bahasa yang digunakan untuk mengakses dan mengelola database. Terdapat beberapa jenis perintah SQL, yaitu:

### 1. Data Definition Language (DDL)

DDL digunakan untuk membuat, mengubah, dan menghapus struktur tabel dalam database. Contoh perintah DDL: CREATE, ALTER, dan DROP.

### 2. Data Manipulation Language (DML)

DML digunakan untuk memanipulasi data dalam tabel, seperti menambah, mengubah, dan menghapus data. Contoh perintah DML: INSERT, SELECT, UPDATE, dan DELETE.

### 3. Data Control Language (DCL)

DCL digunakan untuk mengatur hak akses pengguna dalam database. Contoh perintah DCL: GRANT dan REVOKE.

## Join (Standard SQL ANSI)

Join adalah perintah SQL yang digunakan untuk menggabungkan dua tabel atau lebih berdasarkan kolom yang dimiliki. Terdapat tiga jenis join dalam SQL, yaitu:

### 1. Inner Join

Inner Join digunakan untuk mengambil data yang terdapat pada kedua tabel yang di-join-kan.

### 2. Left Join

Left Join digunakan untuk mengambil semua data dari tabel kiri dan data yang cocok dari tabel kanan.

### 3. Right Join

Right Join digunakan untuk mengambil semua data dari tabel kanan dan data yang cocok dari tabel kiri.

## Pengertian Onion Model

Onion Model adalah sebuah model yang digunakan untuk memisahkan antara data dan aplikasi. Model ini terdiri dari 3 lapisan, yaitu lapisan data, lapisan aplikasi, dan lapisan presentasi. Lapisan data berisi data-data yang akan digunakan oleh aplikasi. Lapisan aplikasi berisi program-program yang akan mengolah data-data yang ada di lapisan data. Lapisan presentasi berisi tampilan-tampilan yang akan ditampilkan kepada pengguna.

## Aggregate Function

Aggregate Function adalah fungsi yang digunakan untuk menghitung nilai-nilai yang berkaitan dengan kumpulan data, seperti nilai rata-rata, nilai maksimum, dan nilai minimum.
