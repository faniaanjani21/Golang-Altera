CREATE DATABASE db_kehadiran;

USE db_kehadiran;

CREATE TABLE staff (
  id_staff INT AUTO_INCREMENT PRIMARY KEY,
  nama_staff VARCHAR(255),
  jabatan VARCHAR(255),
  email VARCHAR(255),
  password VARCHAR(255),
  xppoint INT DEFAULT 0
);

CREATE TABLE attendance_record (
  id_attendance INT AUTO_INCREMENT PRIMARY KEY,
  id_staff INT,
  tgl_masuk DATE,
  waktu_masuk TIME,
  status VARCHAR(255),
  FOREIGN KEY (id_staff) REFERENCES staff(id_staff)
);

CREATE TABLE attendance_notification (
  id_notification INT AUTO_INCREMENT PRIMARY KEY,
  id_staff INT,
  tgl_notifikasi DATE,
  waktu_notifikasi TIME,
  pesan VARCHAR(255),
  FOREIGN KEY (id_staff) REFERENCES staff(id_staff)
);
