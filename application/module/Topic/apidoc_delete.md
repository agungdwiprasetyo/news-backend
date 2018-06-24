## Delete Topic (DELETE in REST API)

### 1. Hapus data topik (hard delete)

Untuk menghapus data topik dari database (hard delete). Data akan hilang selamanya.

**Handler Function: ./application/module/Topic.removeTopic** 

* Method : DELETE
* URL : http://news-api.agungdwiprasetyo.com/topic/remove/{{topic_object_id}}

* Response body example (if success, http status code=200):
```json
{
    "success": true,
    "status": "OK",
    "message": "Success remove topic"
}
```
