# Menjalankan Docker pada Terminal

## Mengecek Versi Docker

Untuk mengecek versi Docker yang terinstall pada komputer, gunakan perintah berikut pada terminal:

```
docker version
```

## Mengecek Images

Untuk mengecek images yang sudah terdownload pada komputer, gunakan perintah berikut pada terminal:

```
docker images
```

## Mengecek Container

Untuk mengecek container yang ada pada komputer, gunakan perintah berikut pada terminal:

```
docker container ls --all
```

Perintah ini akan menampilkan semua container yang pernah dibuat pada komputer.

## Membuat Container

Untuk membuat container baru, gunakan perintah berikut pada terminal:

```
docker container create --name [nama_container] [nama_image]
```

Contoh:

```
docker container create --name golangserver1 golang
```

## Mendownload Images

Untuk mendownload images baru dari Docker Hub, gunakan perintah berikut pada terminal:

```
docker pull [nama_image]
```

Contoh:

```
docker pull golang
```

## Menjalankan Container

Untuk menjalankan container yang sudah dibuat, gunakan perintah berikut pada terminal:

```
docker container start [nama_container]
```

Contoh:

```
docker container start golangserver1
```

## Menghentikan Container

Untuk menghentikan container yang sedang berjalan, gunakan perintah berikut pada terminal:

```
docker container stop [nama_container]
```

Contoh:

```
docker container stop golangserver1
```

## Menghapus Container

Untuk menghapus container yang sudah tidak dibutuhkan, gunakan perintah berikut pada terminal:

```
docker container rm [nama_container]
```

Contoh:

```
docker container rm golangserver1
```

## Menghapus Images

Untuk menghapus images yang sudah tidak dibutuhkan, gunakan perintah berikut pada terminal:

```
docker image rm [nama_image]
```

Contoh:

```
docker image rm golang
```

## Membuka Port Pada Container

Untuk membuka port pada container, gunakan perintah berikut pada terminal:

```
docker container create --name [nama_container] -p [port_host]:[port_container] [nama_image]
```

Contoh:

```
docker container create --name golang1 -p 8000:27017 golang
```

Perintah ini akan membuka port 8000 pada host dan menghubungkannya ke port 27017 pada container.
