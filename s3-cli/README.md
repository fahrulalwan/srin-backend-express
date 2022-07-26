# Try Out S3

Di tugas ini, saya di berikan challenge untuk membuat dokumentasi operasi dari _Amazon Web Services_ (AWS) _Simple Storage Service_ (S3). Disini saya di tantang untuk menunjukkan kode operasi untuk

## Pre-requisites

1. Setup AWS Account
   Untuk dapat mengakses dan menggunakan service AWS, kita perlu [membuat akun terlebih dahulu](https://aws.amazon.com/premiumsupport/knowledge-center/create-and-activate-aws-account/).
2. Install [Python3](https://www.python.org/downloads/)
   AWS CLI membutuhkan environment Python3 sebagai bahasa dasar.
3. Install [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
   AWS CLI diperlukan sebagai penghubung antara komputer lokal kepada service AWS akun kita. Untuk link selengkapnya dapat dijelaskan seperti [berikut](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html), jangan lupa untuk setup akun beserta **download credential terlebih dahulu setelah mendaftar**.
4. Install [Golang](https://go.dev/doc/install)
   Golang diperlukan sebagai penghubung antara _Command Line Interface_ (CLI) dengan AWS _Software Development Kit_ (SDK) yang akan dijelaskan di bawah.

## Pre-configuration

Untuk membuat Simple Storage Service **(S3)** dalam AWS, diperlukan pemahaman pengaturan akses kepada akun terkait. Dari sini kita bisa melakukannya melalui [AWS Console](https://aws.amazon.com/console/).

1. **Pilih Region**
   Pada Halama AWS Console, kita bisa melihat region yang sedang saat ini kita gunakan yang ada di screenshot berikut. Dalam case ini, kita perlu memilih region tempat kita terdekat secara geografis. Kita bisa mengubahnya dengan memilih server Singapore (_ap-southeast-1_).
   ![AWS Console Select Region](/assets/images/aws-console-select-region.png)

2. **_Identity & Access Manager (IAM)_**
   Untuk mengatur akses S3, diperlukan IAM. IAM dibutuhkan untuk mengatur akses user kita kepada service AWS terkait. Kita bisa mengaturnya melalui AWS Console. Kita bisa menambahkan _permissions_ dengan policy **_[AmazonS3FullAccess](https://us-east-1.console.aws.amazon.com/iam/home#/policies/arn%3Aaws%3Aiam%3A%3Aaws%3Apolicy%2FAmazonS3FullAccess)_**. - Create user group (berserta policy S3) & create user baru under user group tersebut (dalam case ini, saya menamakan user-group _s3-dev_ & nama user _srin-dev_) - Download credential untuk created user tersebut (format _.csv_)
3. Konfigurasi **_AWS CLI_**
   Setelah IAM dibuat, kita bisa membuka terminal dan mengetik `aws configure`. kemudian isi sesuai dengan credential berformat **_.csv_**, seperti berikut:
   ` AWS Access Key ID [None]: XXXXXXXXXXXXXXXXXXXX AWS Secret Access Key [None]: XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX Default region name [None]: ap-southeast-1 Default output format [None]: json `
   Setelah itu, kita bisa mengecek konfigurasi kita melalui `aws configure list`

## S3 Operation

Setelah Region & IAM Policy dibuat, kita bisa memanipulasi bucket mulai dari sekarang.

### Bucket Operation

- **Create Bucket**
  ```
  aws s3api create-bucket --bucket <bucket-name> --region <desired-region> [--create-bucket-configuration LocationConstraint=ap-southeast-1]
  ```
  Pada command di atas, kita memerlukan sedikitnya 2 argumen untuk membuat bucket baru. yaitu:
  - `--bucket <bucket-name>`
    Ini adalah nama bucket yang akan kita namakan untuk instance kita nanti
  - `--region <desired-region>`
    Ini adalah region dimana kita akan meletakkan instance S3 kita.
  - `[--create-bucket-configuration LocationConstraint=ap-southeast-1]`
    Bersifat opsional, ini diperlukan ketika kita ingin menetapkan instance bucket kita kepada region yang diinginkan (karena _us-west-1_ adalah default region). Contohnya:
    ` aws s3api create-bucket --bucket srin-bucket --region ap-southeast-1 --create-bucket-configuration LocationConstraint=ap-southeast-1 `
- **List all bucket**
  ```
  aws s3api list-buckets
  ```
  Ini adalah operasi untuk melihat semua bucket yang terdaftar pada akun kita.
- **Delete Bucket**
  Untuk menghapus bucket, diperlukan setidaknya kita **sudah menghapus semua object di dalamnya**. Ketika object dalam bucket tersebut sudah terhapus, maka kita bisa melakukan penghapusan bucket seperti dibawah.
  ` aws s3api delete-bucket --bucket <bucket-name> --region <desired-region> ` - `--bucket <bucket-name>`
  Ini adalah nama bucket yang akan kita namakan untuk instance kita nanti - `--region <desired-region>`
  Ini adalah region dimana kita akan meletakkan region instance S3 kita. Contohnya:
  ` aws s3api delete-bucket --bucket srin-bucket --region ap-southeast-1 `

### Object Operation

Setelah Bucket dibuat, kita bisa memanipulasi data tersebut dengan melakukan beberapa operasi dibawah.

- **Create Object**
  Membuat object dalam AWS CLI sangat mudah, kita hanya perlu menambahkan 3 argumen seperti berikut:
  ` aws s3api put-object --bucket <bucket-name> --key <object-key> --body <path-file-to-upload> ` - `--bucket <bucket-name>`
  Ini adalah nama bucket kita sebagai tempat upload object kita. - `--key <object-key>`
  Ini adalah nama object yang akan kita namakan pada saat object telah terupload di bucket yang telah kita namakan. - `--body <path-file-to-upload>`
  Ini adalah _file-path_ letak folder kita sebagai reference untuk upload object pada bucket kita.

      Kita juga bisa membuat membuat folder terpisah dengan menambahkan `--key ...<folder-name>/<file-name>` didalamnya.
      Kita juga bisa memperbarui object, dengan melakukan command yang sama. Hasilnya adalah object dengan key tersebut tergantikan dengan object baru. Contoh:
      ```
      aws s3api put-object --bucket srin-bucket --key raleway-font.zip --body raleway-v28-latin.zip
      ```

- **Delete Object**
  Untuk menghapus object dalam bucket, kita hanya perlu ketikkan command berikut:
  ` aws s3api delete-object --bucket <bucket-name> --key <object-key> ` - `--bucket <bucket-name>`
  Ini adalah nama bucket kita sebagai tempat upload object kita. - `--key <object-key>`
  Ini adalah nama object yang akan kita namakan pada saat object telah terupload di bucket yang telah kita namakan.

- **List object in specific bucket & sub-bucket**
  Kita bisa melihat list object yang tersedia di bucket seperti berikut
  ` aws s3api list-objects --bucket <bucket-name> ` - `--bucket <bucket-name>`
  Ini adalah nama bucket kita untuk melihat daftar object dalam bucket kita (termasuk sub-bucket).

## S3 Operation Coding

Pada use case specific, kita bisa kustomisasi AWS CLI menyesuaikan dengan kebutuhan pengguna. Seperti contohnya: melakukan custom AWS CLI menggunakan Golang & AWS SDK. **_AWS SDK_** adalah perangkat pengembangan piranti lunak dari AWS untuk memudahkan kita melakukan kustomisasi CLI sehingga kita bisa menggunakan _custom-cli_ untuk kebutuhan kita.

Untuk mengaksesnya, kita bisa melakukan compile terlebih dahulu, seperti berikut:

```
go run srin-aws-cli.go <command>
```

atau

```
go build srin-aws-cli.go
./srin-aws-cli <command>
```

### CLI Commands

pada **srin-aws-cli**, kita bisa melakukan perintahkan untuk melakukan operasi tersebut layaknya AWS CLI, diantaranya:

- **create-bucket**
  Kita bisa memerintahkan srin-aws-cli untuk membuat bucket baru seperti berikut:
  ` srin-aws-cli create-bucket <bucket-name> `
- **list-buckets**
  `srin-aws-cli` bisa dilakukan untuk membuat daftar bucket seperti berikut:
  ` srin-aws-cli list-bucket <bucket-name> `
- **delete-bucket**
  Kita juga bisa memerintahkan `srin-aws-cli` untuk menghapus bucket seperti berikut:
  ` srin-aws-cli delete-bucket <bucket-name> `
- **upload-object**
  `srin-aws-cli` juga bisa digunakan untuk mengunggah object pada bucket seperti berikut:
  ` srin-aws-cli upload-object <bucket-name> --key=<bucket-file-path> --file=<current-file-path> ` - `--key <bucket-file-path>`
  Ini adalah nama object untuk peletakkan direktori & nama file dalam bucket kita. - `--file <current-file-path>`
  Ini adalah path menuju file yang akan kita upload ke bucket AWS S3.

- **list-objects**
  Untuk melihat daftar object dalam bucket, kita bisa perintahkan `srin-aws-cli` seperti berikut:
  ` srin-aws-cli list-object <bucket-name> `
- **delete-object**
  Untuk menghapus object, kita bisa melakukan perintah seperti berikut:
  ` srin-aws-cli delete-object <bucket-name> --key=<bucket-file-path> ` - `--key <bucket-file-path>`
  Ini adalah nama object untuk peletakkan direktori & nama file dalam bucket kita.

# Challenges during operation

Sepanjang percobaan saya terhadap percobaan S3 ini, saya merasa kesulitan untuk melalui _getting started with AWS_ dikarenakan minimnya pemahaman saya untuk menggunakan AWS console.
