## Go-Lang Web
- Populer digunakan untuk aplikasi web backend
- Disediakan package untuk membuat web beserta testingnya

## Package `net/http`
- Built-in package untuk membuat aplikasi berbasis web
- Direkomendasikan untuk menggunakan framework yang baik dalam developer experience.

## Server `net/http`
- Struct yang ada di package `net/http`

## Interface `Handler`
- Digunakan untuk menerima HTTP Request yang masuk ke server
- Direpresentasikan dalam interface, di dalamnya terdapat function `ServeHTTP()`

## `HandlerFunc`
- Implementasi dari `Handler`
- Membuat function handler HTTP
- Hanya support 1 endpoint saja

## `ServeMux`
- Implementasi dari `Handler`
- Mendukung banyak endpoint
- Di bahasa pemrograman lain istilahnya adalah **router**

## URL Pattern pada `ServeMux`
- `/` di akhir pattern = re"/.+"
- Cth: `/images/` dapat menerima endpoint berikut
  - `/images/house`
  - `/images/appartment`
  - `/images/lampung/pringsewu`
  - dll

## Request
- Struct yang merepresentasikan HTTP Request yang dikirim oleh Client (Web browser, Postman, dll)
- Semua informasi request dapat diambil:
  - Cth: URL, http method, http header, http body, dll.