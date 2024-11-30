#ambil golang(aplikasi induk) sebagai base
FROM golang:1:21

#daftarkan direktory
ADD . /stock-inventory
WORKDIR /stock-inventory

#copy modules dan addons
COPY go.mod go.sum ./
RUN go mod download

#copy semua kode aplikasi
COPY . .

#build aplikasi
RUN go build -o main.go

#ambil port
EXPOSE 8080

#jalankan aplikasi
CMD [ "/main" ]