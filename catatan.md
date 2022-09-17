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


## HTTP Test
- Unit test untuk handler tanpa harus menjalankan aplikasi webnya
- Dengan menggunakan HTTP Test, developer dapat fokus pada handler function yang ingin ditest
- More: https://golang.org/pkg/net/http/httptest/

**Implementasi**
- `httptest.NewRequest(method, url, body)`, untuk simulasi request
- Dapat menggunakan function di atas untuk simulasi pada unit test
- Dapat menambahkan informasi tambahan lainnya pada request
  - Header
  - Cookie
  - dll.

- `httptest.NewRecorder()`, untuk membuat `ResponseRecorder`
- `ResponseRecorder` adalah struct bantuan untuk merekam HTTP Response dari hasil testing

## Query Parameter
- Mengirim data yang ada di url dari client (browser, postman, dll) ke server 
- Cth: localhost:9090/say-hello?name=rasyidev

## Header
- Informasi tambahan yg biasa dikirim dari client ke server atau sebaliknya.
- Biasanya ada pada HTTP Request dan HTTP response
- Saat menggunakan browser, otomatis header akan ditambahkan oleh browser tersebut seperti informasi browser, dll.

## Request Header
- Untuk menangkap request header yang dikirim oleh client
- `Request.Header`
- Nilainya berupa `map[string]string`
- _case insensitive_

## Form Post
**Mengirim Form Post**
- Untuk mengirim body form menggunakan bantuan `strings.NewReader` lalu dikirimkan menggunakan:
- `request.httptest.NewRequest(http.MethodPost, "/", strings.NewReader)`
- Wajib menambahkan Header `"Content-Type"` denvan value `"application/x-www-form-urlencoded"`

**Menangkap Form Post**
- Parse dari form menggunakan `request.ParseForm()`
- Menangkap value dari form post menggunakan `request.PostForm()` atau `request.Form.Get("key")`

## Response Code
- Dikirim oleh server, default: 200 (ok)
- Semua status code sudah disiapkan Go Lang : https://github.com/golang/go/blob/master/src/net/http/status.go

## Cookie
- Fitur HTTP yang memungkinkan server memberi response cookie (key-value) dan client menyimpan cookie tersebut sebagai tiket untuk request selanjutnya.
- HTTP menggunakan konsep stateless antara client dan server. Server tidak akan menyimpan data apapun untuk mengingat setiap request dari client, agar lebih mudah melakukan scaling.
- Untuk mengingat data dari client tertentu, contohnya pada saat login pada website, server harus ingat bahwa client tersebut sudah login sehingga tidak perlu diminta untuk login lagi.
- Cookie dapat dimanfaatkan untuk mengingat data dari client tertentu
**Membuat Cookie**
```go
cookie := new(http.Cookie)
cookie.Name = "namanya"
cookie.Value = "nilainya"
cookie.Path = "/" // pathnya
``