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
```

## File Server
- Menggunakan `http.Dir`, saat compile aplikasi, directory harus tetap diupload juga ke dalam server. Hal ini karena static file tidak ikut dicompile
- Menggunakan Go Embed `embed.FS`, saat compile aplikasi, directory tidak perlu diupload juga ke dalam server. Hal ini karena static file sudah ikut dicompile dalam binary aplikasinya.

## Menggunakan Static File tertentu
- `http.ServeFile()`


## HTML Template
- Untuk membuat halaman HTML yang dinamis
- Sebelum menggunakan HTML Template, harus membuat template terlebih dahulu
- Template dapat berupa string atau file
- `{{.}}` untuk membuat bagian html yang dinamis
- Template -> Parse -> Execute

## HTML Template Directory
- Untuk load banyak file html sekaligus
- `template.Glob()`

## Template Data
- Memasukkan multiple value ke dalam html
- Menggunakan `struct` atau `map`
- `{{.NamaField}}`

## Template Action
- if-else, loop, dll.
**If-Else**
- `{{if.Value}}` ... `{{end}}`, Jika value tidak kosong block akan dieksekusi
- `{{if.Value}}` ... `{{else if.Value}}` ... `{{end}}`
- `{{if.Value}}` ... `{{else}}` ... `{{end}}`

**Operator Perbandingan**
- `eq`, equal
- `ne`, not equal
- `lt`, less than
- `le`, less than equal
- `gt`, greater than
- `ge`, greater than equal

**Range**
- Perulangan
- `{{range $index, $element := .Value}}`
- Bisa menggunakan else jika `.Value` kosong

## Template Layout
- Digunakan untuk membuat template html dinamis
- Cocok untuk halaman yang memiliki tampilan yang digunakan oleh banyak halaman seperti header, footer, dll.
- Default: nama file menjadi nama template
- `{{define "nama template"}}`, untuk mengubah nama template

## Template Function
- Mengakses function / method dalam struct. Struct ini dikirim ke html template
- `{{.NamaFunction}}`, memanggil function tanpa parameter
- `{{.NamaFunction "arg1", "arg2"}}`, memanggil function dengan parameter

## Template Caching
- Melakukan parsing template berulang - ulang sebenarnya tidak efisien karena setiap handler dipanggil maka akan selalu melakukan parsing ulang.
- **Idealnya** template hanya melakukan parsing satu kali di awal ketika aplikasi berjalan
- Selanjutnya template akan di-_caching_ (disimpan di dalam memory), sehingga tidak perlu melakukan parsing lagi. Hal ini membuat web kita semakin cepat di-_load_

## Cross Site Scripting (XSS)
- Issue security atau celah keamanan ketika membuat aplikasi berbasis web
- User sengaja memasukkan parameter tertentu pada url yang mengandung JavaScript agar dirender oleh aplikasi web target
- Biasanya digunakan untuk mencuri cookie browser untuk mengambil alih akun kita yang tersimpan di dalam browser


**Dapat Diatasi dengan Template Auto Escape Go-Lang**
- Dapat mengabaikan (escape) secara otomatis apabila terdapat input berupa tag html atau JavaScript
- More : https://github.com/golang/go/blob/master/src/html/template/escape.go
- Secara default belum aman karena sudah di-escape (go 1.19)

## Redirect
- Memindahkan path saat ini ke path tujuan
- Pada HTTP, hanya perlu membuat response code 3xx dan menambah header location
- Di go lang ada function tertentu untuk memudahkan redirect
- `http.Redirect(writer, request, url, code)`

## Upload File: Multipart
- Parsing Multipart menggunakan `Request.ParseMultipartForm(size)` 
- Tanpa parsing, bisa menggunakan `Request.FormFile(name)`

## Download File
- Pada standar HTTP, saat download file terdapat header response `Content-Disposition`
  - `inline` **default**, hanya menampilkan di browser
  - `attachment`, memaksa browser untuk mendownload file

## Middleware, Filter, atau Interceptor
- Fitur yang memungkinkan kita untuk menambahkan block program sebelum dan setelah sebuah handler dieksekusi
- request -> middleware -> handler -> middleware -> response
- Sayangnya di Go Lang tidak ada middleware, kita harus membuat sendiri middlewarenya
- Kita dapat membuat middleware sendiri menggunakan handler

### Middleware Error Handler
- Dapat melakukan recover jika ada error dan mengubah panic menjadi error response dari server
- Dengan menerapkan ini, aplikasi tidak akan berhenti berjalan meskipun ada error. Pesan error akan ada di log