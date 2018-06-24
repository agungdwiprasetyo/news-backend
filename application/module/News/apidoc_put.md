## Update News (PUT in REST API)

### 1.Update informasi berita
Untuk mengupdate informasi detail dari sebuah berita.
Sudah bersifat dinamis, yaitu cukup update field yg dibutuhkan saja (tidak harus semua field disertakan).
Dalam fungsi ini, sudah termasuk update status berita ("konsep", "dihapus", atau "publikasikan").
Dan juga dalam fungsi ini sudah termasuk update informasi topik untuk sebuah berita (dalam hal ini hanya menambahkan topik, untuk menghapus topik ada dalam fungsi ```deleteTopic```). 
Daftar ```topic``` dalam satu object berita bersifat ```unique```, jadi tidak ada duplikasi data.
Sama seperti fungsi tambah berita baru, jika data ```topic``` yang disertakan dalam payload belum tersedia di database, maka sistem akan menambahkan data ```topic``` tersebut ke database.
**Handler Function: ./application/module/News.updateNews** 

* Method : PUT
* URL : http://kumparan-api.agungdwiprasetyo.com/news/update/{news_object_id}
* Header : 
```json
{
	"Content-Type": "application/json"
}
```
* Request body payload example:
```json
{
	"title": "judul-berita-baru [string]",
	"value": "isi-berita-baru [string]",
	"status": "enum 'konsep', 'dihapus', atau 'publikasikan' [string]"
	"topics": [
		{
			"name": "topic-1 [string]",
		},
		{
			"name": "topic-2 [string]"
		},
		...
	]
}
```
Jika hanya ingin mengupdate status berita, maka payload:
```json
{
	"status": "enum 'konsep', 'dihapus', dan 'publikasikan' [string]"
}
```
Jika hanya ingin mengupdate info topik (menambahkan data), maka payload:
```json
{
	"topics": [
		{
			"name": "topic-1 [string]",
		},
		{
			"name": "topic-2 [string]"
		},
		...
	]
}
```

* Response body example (if success, http status code=200):
```json
{
    "success": true,
    "status": "OK",
    "message": "Success update news"
}
```
