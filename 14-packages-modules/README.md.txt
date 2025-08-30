# Packages and Modules Example

Contoh penggunaan **package** dan **module** di Go.

## Struktur Folder
├── go.mod
├── main.go
└── greet/
└── greet.go
## Penjelasan
- **go.mod** → inisialisasi module Go (`go mod init github.com/kevinandrea/hellopkg`)
- **main.go** → file utama yang memanggil fungsi dari package `greet`
- **greet/greet.go** → berisi fungsi `Hello` dan `GetMessage`

## Cara Menjalankan
1. Masuk ke folder project:
   ```bash
   cd 14-packages-modules
