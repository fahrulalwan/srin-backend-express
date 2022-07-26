# API Server

Di tugas kedua, saya di berikan challenge untuk membuat Backend server, memiliki fitur CRUD, dengan memilih diantara **Express.js**, **Spring Framework** atau **Golang**. Saya memilih untuk melakukan development menggunakan **Express.js**.

## Pre-requisites

1. Setup [Node.js](https://nodejs.org/en/)
   Untuk membuat aplikasi backend Express memerlukan dependensi Node.js latest **LTS** version.
   Untuk dapat mengakses dan menggunakan service AWS, kita perlu [membuat akun terlebih dahulu](https://aws.amazon.com/premiumsupport/knowledge-center/create-and-activate-aws-account/).
2. Setup [Docker](https://docs.docker.com/)
   Pada OS Windows, bisa merujuk pada [link ini](https://docs.docker.com/desktop/install/windows-install/) untuk download **_Docker Desktop_**. Namun untuk Linux-based, bisa merujuk pada [link ini](https://docs.docker.com/desktop/install/linux-install/).

## Docker

```
docker login
docker build -t srin-api-express .
docker run -d -p 3000:3000 srin-api-express
docker push fahrulalwan/srin-api-express
```

# Challenges during development

Sepanjang pengembangan backend menggunakan Express.js, saya menemukan beberapa tantangan, diantaranya:

1. Saya merasa kesulitan untuk mengintegrasikan API pada Swagger, sehingga dokumentasi API Swagger tidak terlihat. Kemungkinan besar ada kesalahan dari saya untuk mengintegrasikan pada _swagger-jsdoc_ library.
2. Saya kesulitan untuk mengoneksikan API ke MongoDB, karena kekurangan saya terhadap pengintegrasian backend menggunakan MongoDB.
