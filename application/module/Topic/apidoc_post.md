## Create Topic (POST in REST API)

### 1.Input data topik baru

Untuk membuat data topik baru. Sudah ada validasi untuk mencegah duplikasi data.

**Handler Function: ./application/module/Topic.addNewTopic** 

* Method : POST
* URL : http://news-api.agungdwiprasetyo.com/topic/add
* Header : 
```json
{
	"Content-Type": "application/json"
}
```
* Request body payload example:
```json
{
	"name": "topic-name [string]"
}
```
* Response body example (if success, http status code=200):
```json
{
    "success": true,
    "status": "OK",
    "message": "Success add new topic. Last ID: {{new_topic_object_id}}"
}
```
