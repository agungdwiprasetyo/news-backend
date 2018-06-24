## Update Topic (PUT in REST API)

### 1.Update data topik

Untuk mengupdate informasi detail dari data topik.

Sudah bersifat dinamis, yaitu cukup update field yg dibutuhkan saja (tidak harus semua field disertakan).

**Handler Function: ./application/module/Topic.updateTopic** 

* Method : PUT
* URL : http://news-api.agungdwiprasetyo.com/topic/update/{{topic_object_id}}
* Header : 
```json
{
	"Content-Type": "application/json"
}
```
* Request body payload example:
```json
{
	"name": "nama-topic-baru [string]"
}
```

* Response body example (if success, http status code=200):
```json
{
    "success": true,
    "status": "OK",
    "message": "Success update topic"
}
```
