## Create News (POST in REST API)

### 1.Buat berita baru

Untuk membuat data berita baru.

Dalam request body payload, jika tidak terdapat field ```status``` maka nilai defaultnya yaitu "konsep".
Untuk field ```topics``` juga sudah dinamis, jika data ```topic``` yang disertakan dalam payload belum tersedia di database, maka sistem akan menambahkan data ```topic``` tersebut ke database.

**Handler Function: ./application/module/News.addNews** 

* Method : POST
* URL : http://news-api.agungdwiprasetyo.com/news/add
* Header : 
```json
{
	"Content-Type": "application/json"
}
```
* Request body payload example:
```json
{
	"title": "judul-berita [string]",
	"value": "isi-berita [string]",
	"status": "enum 'konsep', 'dihapus', atau 'publikasikan'. Default 'konsep' [string]",
	"topics": [
		{
			"name": "topic-1 [string]"
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
    "message": "Success add new news. Last ID: {{new_news_object_id}}"
}
```
