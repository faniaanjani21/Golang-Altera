CREATE DATABASE db_kehadiran;

\c db_kehadiran;

CREATE TABLE staff (
  id_staff SERIAL PRIMARY KEY,
  nama_staff VARCHAR(255),
  jabatan VARCHAR(255),
  email VARCHAR(255),
  password VARCHAR(255)
);

CREATE TABLE attendance_record (
  id_attendance SERIAL PRIMARY KEY,
  id_staff INT,
  tgl_masuk DATE,
  waktu_masuk TIME,
  status VARCHAR(255),
  FOREIGN KEY (id_staff) REFERENCES staff(id_staff)
);

CREATE TABLE attendance_notification (
  id_notification SERIAL PRIMARY KEY,
  id_staff INT,
  tgl_notifikasi DATE,
  waktu_notifikasi TIME,
  pesan VARCHAR(255),
  FOREIGN KEY (id_staff) REFERENCES staff(id_staff)
);
