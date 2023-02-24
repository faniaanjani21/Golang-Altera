materi databse scema dan data definition laguage

indtroduction database
relationship data bse
entitiy relationship diagram
sql statment

pengertian
databaseadalah sekumpulan data yang terorganisir

one to one adalah sebuah hubungan anatara database
yang mengacu pada dua table dalam data base relational
hubungan antara dua tabel dalam basis data relasional di mana satu rekaman dalam tabel pertama hanya berkorespondensi dengan satu rekaman dalam tabel kedua, dan sebaliknya. Artinya, setiap rekaman dalam satu tabel dikaitkan dengan satu rekaman tunggal di tabel lainnya, dan tidak ada hubungan duplikat. Hubungan satu-satu sering digunakan untuk memecah tabel yang lebih besar menjadi tabel yang lebih kecil dan lebih mudah dikelola, dan dapat diterapkan dengan menggunakan keterikatan kunci primer-kunci asing.

contoh nya
Contoh hubungan satu-satu dalam basis data bisa diilustrasikan dengan menggunakan tabel "customer" dan "customer_address". Setiap pelanggan (customer) hanya memiliki satu alamat (customer_address) dan setiap alamat hanya dimiliki oleh satu pelanggan.

Tabel "customer" berisi informasi tentang pelanggan seperti id, nama, email, dll. Tabel "customer_address" berisi informasi tentang alamat pelanggan seperti id, alamat, kota, kode pos, dll.

Untuk menetapkan hubungan satu-satu antara kedua tabel tersebut, dapat digunakan kunci asing (foreign key) pada tabel "customer_address" yang merujuk ke kunci primer (primary key) pada tabel "customer". Dengan demikian, setiap rekaman di tabel "customer_address" hanya berkorespondensi dengan satu rekaman di tabel "customer" dan tidak ada rekaman yang terduplikasi.

one to many

- pengertian
  Hubungan satu-ke-banyak (one-to-many) dalam basis data mengacu pada hubungan antara dua tabel dalam basis data relasional di mana satu rekaman dalam tabel pertama dapat berkorespondensi dengan banyak rekaman dalam tabel kedua. Namun, setiap rekaman di tabel kedua hanya dapat berkorespondensi dengan satu rekaman di tabel pertama.

Contoh umum dari hubungan satu-ke-banyak adalah hubungan antara tabel "orders" dan "order_details" dalam sebuah sistem manajemen toko online. Setiap pesanan (order) dapat memiliki banyak detail pesanan (order_details), seperti nama produk, harga, jumlah, dll. Namun, setiap detail pesanan hanya terkait dengan satu pesanan.

Tabel "orders" akan berisi informasi tentang pesanan seperti nomor

contoh
Dalam contoh ini, tabel "orders" memiliki kolom "Order ID" sebagai kunci primer (primary key) dan tabel "order_details" memiliki kolom "Order Detail ID" sebagai kunci primer dan "Order ID" sebagai kunci asing (foreign key) yang merujuk ke kolom "Order ID" di tabel "orders".

Karena hubungan satu-ke-banyak, satu pesanan dalam tabel "orders" dapat berkorespondensi dengan banyak detail pesanan dalam tabel "order_details", tetapi setiap detail pesanan hanya dapat berkorespondensi dengan satu pesanan dalam tabel "orders". Misalnya, pesanan dengan Order ID 1001 memiliki 2 detail pesanan dalam tabel "order_details" dengan Order Detail ID 1 dan 2. Sedangkan detail pesanan dengan Order Detail ID 1 dan 2 hanya berkorespondensi dengan pesanan yang memiliki Order ID 1001 di tabel "orders

many to many
Many-to-many merupakan jenis hubungan antara dua tabel dalam basis data relasional di mana banyak rekaman dalam tabel pertama dapat berkorespondensi dengan banyak rekaman dalam tabel kedua, dan sebaliknya. Artinya, sebuah rekaman dalam tabel pertama dapat memiliki banyak rekaman di tabel kedua yang terkait, dan sebuah rekaman di tabel kedua dapat berkorespondensi dengan banyak rekaman di tabel pertama.

contoh nya
tabel "students" dan "courses" dalam sistem manajemen sekolah. Seorang siswa (student) dapat mendaftar pada banyak kursus (courses), dan sebuah kursus dapat diikuti oleh banyak siswa.

Untuk mengimplementasikan hubungan many-to-many, dibutuhkan sebuah tabel pivot atau tabel relasi yang menyimpan kunci asing (foreign key) dari kedua tabel. Tabel pivot ini berfungsi untuk menghubungkan rekaman dari kedua tabel.

Contoh tabel pivot untuk hubungan many-to-many antara tabel "students" dan "courses":

Tabel "student_courses":
Dalam contoh ini, tabel "students" memiliki kolom "Student ID" sebagai kunci primer dan tabel "courses" memiliki kolom "Course ID" sebagai kunci primer. Tabel pivot "student_courses" memiliki kolom "Student ID" dan "Course ID" sebagai kunci asing yang merujuk ke kolom "Student ID" dan "Course ID" di tabel "students" dan "courses".

Dalam tabel "student_courses", setiap baris mewakili hubungan antara satu siswa dan satu kursus yang diikuti oleh siswa tersebut. Seorang siswa dengan Student ID 001 misalnya, mengikuti dua kursus dengan Course ID 101 dan 102. Sebaliknya, kursus dengan Course ID 101 diikuti oleh dua siswa dengan Student ID 001 dan 002.

cara implementasi
rdbms
ddl = data definition language
dml = data manipulation language
dcl = data control language

statement operation
insert
selec
update
delete

contoh insert
INSERT INTO table_name (column1, column2, column3, ...)

contoh select
SELECT column1, column2, column3, ...

contoh update
UPDATE table_name

contoh delete
DELETE FROM table_name

join standar sql ansi
inner join
left join
right join

contoh inner join
SELECT column_name(s)

contoh left join
SELECT column_name(s)

contoh right join
SELECT column_name(s)

pengertian onion model
onion model adalah sebuah model yang digunakan untuk memisahkan antara data dan aplikasi. Model ini terdiri dari 3 lapisan, yaitu lapisan data, lapisan aplikasi, dan lapisan presentasi. Lapisan data berisi data-data yang akan digunakan oleh aplikasi. Lapisan aplikasi berisi program-program yang akan mengolah data-data yang ada di lapisan data. Lapisan presentasi berisi tampilan tampilan yang akan ditampilkan kepada pengguna.

contoh union
SELECT column_name(s) FROM table1

dungsi agregasi
adalah di mana nilai beberapa baris di kelompokan bersama untuk nilai ringkaasan tunggal
