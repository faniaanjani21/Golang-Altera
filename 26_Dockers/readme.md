#mengecek versi
docker version

#mengecek images
docker images

#mengecek container semua
docker container ls --all

#mengecek container yang jalan
docker container ls

#membuat container
docker container create --name golangserver1 golang

#mendownload images
docker pull golang

#menjalankan container
docker container start golangserver1

#menstop container
docker container stop golangserver1

#menghapus docker container
docker container rm golangserver1

#menghapus imager
docker image rm golang

#membuka port pada docker
docker container create --name golang1 -p 8000:27017 golang
