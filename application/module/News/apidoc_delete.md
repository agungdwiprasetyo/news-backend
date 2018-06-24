## Delete News (DELETE in REST API)

### 1.Hapus topik dalam berita
Untuk menghapus satu objek Topik yang berada dalam array ```topics``` dalam objek Berita.
Contoh model object data dari suatu berita:
```json
{
	 "id": "5b2f78b6a150187e295875fe",
	 "title": "Pemilu 2018",
	 "value": "Pemilu serentak 2018 yang dilaksanakan pada hari rabu, 27 Juni 2018 mendatang dikabarkan aman....",
	 "posted_at": "2018-06-24T17:55:50.702+07:00",
	 "status": "konsep",
	 "total_comment": 0,
	 "total_like": 0,
	 "topics": [
	 	{
	 		"id": "5b2f78b6a150187e295875ff",
	 		"string_id": "pemilu",
	 		"name": "Pemilu"
	 	},
	 	{
	 		"id": "5b2f78b6a150187e29587600",
	 		"string_id": "pilkada",
	 		"name": "Pilkada"
	 	}
	 ]
}
```
**Handler Function: ./application/module/News.deleteTopic** 

* Method : DELETE
* URL : http://kumparan-api.agungdwiprasetyo.com/news/topic?news_id={{news_object_id}}&topic_id={{topic_object_id}}
Misal jika ingin menghapus topik "Pilkada" dari contoh model data diatas maka url-nya yaitu: URL : http://kumparan-api.agungdwiprasetyo.com/news/topic?news_id=5b2f78b6a150187e295875fe&topic_id=5b2f78b6a150187e29587600

* Response body example (if success, http status code=200):
```json
{
    "success": true,
    "status": "OK",
    "message": "Success remove topic in news"
}
```

### 2. Arsipkan data berita (soft delete)
Untuk mengarsipkan data berita (soft delete).

* Method : DELETE
* URL : http://kumparan-api.agungdwiprasetyo.com/news/archive/{{news_object_id}}

* Response body example (if success, http status code=200):
```json
{
    "success": true,
    "status": "OK",
    "message": "Success archive news"
}
```

### 3. Hapus data berita (hard delete)
Untuk menghapus data berita dari database (hard delete). Data akan hilang selamanya.

* Method : DELETE
* URL : http://kumparan-api.agungdwiprasetyo.com/news/remove/{{news_object_id}}

* Response body example (if success, http status code=200):
```json
{
    "success": true,
    "status": "OK",
    "message": "Success remove news"
}
```
