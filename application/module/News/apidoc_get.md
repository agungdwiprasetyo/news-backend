## Read News (GET in REST API)

### 1. Get list berita

Untuk mendapatkan semua data berita dengan parameter partisi halaman untuk limitasi data (dengan url query ```page``` dan ```limit```). Jika parameter url query ini tidak disertakan, maka nilai default untuk parameter ```page``` = 1 dan parameter ```limit``` = 10.
**Handler Function: ./application/module/News.getAllNews** 

* Method : GET
* URL : http://kumparan-api.agungdwiprasetyo.com/news/list?page=1&limit=10
* Response body example:
```json
{
    "success": true,
    "message": "Success",
    "page": 1,
    "limit": 10,
    "data_per_page": 2,
    "total_page": 1,
    "total_all_data": 2,
    "data": [
        {
            "id": "5b2f409da15018732b3972cf",
            "title": "Kecelakaan mobil gandeng di tol",
            "value": "Telah terjadi kecelakaan dalam tol jakarta",
            "posted_at": "2018-06-24T13:56:29.507+07:00",
            "status": "konsep",
            "total_comment": 0,
            "total_like": 0,
            "topics": [
                {
                    "id": "5b2f284ba150181cbf747e50",
                    "string_id": "news",
                    "name": "News"
                },
                {
                    "id": "5b2f284ba150181cbf747e52",
                    "string_id": "jakarta",
                    "name": "Jakarta"
                }
            ]
        },
        {
            "id": "5b2f4081a15018732b3972ca",
            "title": "CFD Dibuka Lebih Cepat Karena Banyak Warga Telah Kembali ke Jakarta",
            "value": "Hari Bebas Kendaraan Bermotor atau Car Free Day (CFD) dibuka lebih cepat. Sebelumnya, Pemprov DKI sempat mengumumkan bakal meniadakan CFD pada 10, 17, dan 24 Juni 2018. Namun tanggal tersebut direvisi dan kembali dibuka Minggu (24/6) karena banyak penduduk yang telah kembali ke Jakarta setelah mudik Lebaran. Perubahan tanggal tersebut dilakukan jauh hari sebelum Lebaran. Kadishub DKI Jakarta Andri Yansyah mengatakan CFD kembali digelar karena masyarakat sudah kembali dari mudik. Maka itu minggu ini CFD sudah ramai dikunjungi masyarakat. ”Masyarakat sudah pulih. Yang balik mudik juga sudah kembali mereka butuh refreshing kan,” ungkap Andri saat dihubungi kumparan.",
            "posted_at": "2018-06-24T13:56:01.055+07:00",
            "status": "publikasikan",
            "total_comment": 0,
            "total_like": 0,
            "topics": [
                {
                    "id": "5b2f284ba150181cbf747e50",
                    "string_id": "news",
                    "name": "News"
                },
                {
                    "id": "5b2f284ba150181cbf747e51",
                    "string_id": "megapolitan",
                    "name": "Megapolitan"
                },
                {
                    "id": "5b2f284ba150181cbf747e52",
                    "string_id": "jakarta",
                    "name": "Jakarta"
                },
                {
                    "id": "5b2f284ba150181cbf747e53",
                    "string_id": "car-free-day",
                    "name": "Car Free Day"
                }
            ]
        }
    ]
}
```

### 2. Get detail berita

Untuk melihat/membaca isi berita.
**Handler Function: ./application/module/News.getNewsByID** 

* Method : GET
* URL : http://kumparan-api.agungdwiprasetyo.com/news/find/{{news_object_id}}
* Response body example:
```json
{
    "success": true,
    "data": {
        "id": "5b2f4081a15018732b3972ca",
        "title": "CFD Dibuka Lebih Cepat Karena Banyak Warga Telah Kembali ke Jakarta",
        "value": "Hari Bebas Kendaraan Bermotor atau Car Free Day (CFD) dibuka lebih cepat. Sebelumnya, Pemprov DKI sempat mengumumkan bakal meniadakan CFD pada 10, 17, dan 24 Juni 2018. Namun tanggal tersebut direvisi dan kembali dibuka Minggu (24/6) karena banyak penduduk yang telah kembali ke Jakarta setelah mudik Lebaran. Perubahan tanggal tersebut dilakukan jauh hari sebelum Lebaran. Kadishub DKI Jakarta Andri Yansyah mengatakan CFD kembali digelar karena masyarakat sudah kembali dari mudik. Maka itu minggu ini CFD sudah ramai dikunjungi masyarakat. ”Masyarakat sudah pulih. Yang balik mudik juga sudah kembali mereka butuh refreshing kan,” ungkap Andri saat dihubungi kumparan.",
        "posted_at": "2018-06-24T13:56:01.055+07:00",
        "status": "publikasikan",
        "total_comment": 0,
        "total_like": 0,
        "topics": [
            {
                "id": "5b2f284ba150181cbf747e50",
                "string_id": "news",
                "name": "News"
            },
            {
                "id": "5b2f284ba150181cbf747e51",
                "string_id": "megapolitan",
                "name": "Megapolitan"
            },
            {
                "id": "5b2f284ba150181cbf747e52",
                "string_id": "jakarta",
                "name": "Jakarta"
            },
            {
                "id": "5b2f284ba150181cbf747e53",
                "string_id": "car-free-day",
                "name": "Car Free Day"
            }
        ]
    }
}
```

### 3. Filter berita berdasarkan topik
Untuk memfilter berita apa saja yang termasuk dalam topik tertentu (Aktifkan filter berita berdasarkan topiknya).
**Handler Function: ./application/module/News.filterNewsByTopic**

* Method : GET
* URL : http://kumparan-api.agungdwiprasetyo.com/news/filter/topic/{{topic_string_id}}
* Response body example (misal get data dengan parameter topic_string_id="jakarta"):
```json
{
    "success": true,
    "total": 2,
    "data": [
        {
            "id": "5b2f4081a15018732b3972ca",
            "title": "CFD Dibuka Lebih Cepat Karena Banyak Warga Telah Kembali ke Jakarta",
            "value": "Hari Bebas Kendaraan Bermotor atau Car Free Day (CFD) dibuka lebih cepat. Sebelumnya, Pemprov DKI sempat mengumumkan bakal meniadakan CFD pada 10, 17, dan 24 Juni 2018. Namun tanggal tersebut direvisi dan kembali dibuka Minggu (24/6) karena banyak penduduk yang telah kembali ke Jakarta setelah mudik Lebaran. Perubahan tanggal tersebut dilakukan jauh hari sebelum Lebaran. Kadishub DKI Jakarta Andri Yansyah mengatakan CFD kembali digelar karena masyarakat sudah kembali dari mudik. Maka itu minggu ini CFD sudah ramai dikunjungi masyarakat. ”Masyarakat sudah pulih. Yang balik mudik juga sudah kembali mereka butuh refreshing kan,” ungkap Andri saat dihubungi kumparan.",
            "posted_at": "2018-06-24T13:56:01.055+07:00",
            "status": "publikasikan",
            "total_comment": 0,
            "total_like": 0,
            "topics": [
                {
                    "id": "5b2f284ba150181cbf747e50",
                    "string_id": "news",
                    "name": "News"
                },
                {
                    "id": "5b2f284ba150181cbf747e51",
                    "string_id": "megapolitan",
                    "name": "Megapolitan"
                },
                {
                    "id": "5b2f284ba150181cbf747e52",
                    "string_id": "jakarta",
                    "name": "Jakarta"
                },
                {
                    "id": "5b2f284ba150181cbf747e53",
                    "string_id": "car-free-day",
                    "name": "Car Free Day"
                }
            ]
        },
        {
            "id": "5b2f409da15018732b3972cf",
            "title": "Kecelakaan mobil gandeng di tol",
            "value": "Telah terjadi kecelakaan dalam tol jakarta",
            "posted_at": "2018-06-24T13:56:29.507+07:00",
            "status": "konsep",
            "total_comment": 0,
            "total_like": 0,
            "topics": [
                {
                    "id": "5b2f284ba150181cbf747e50",
                    "string_id": "news",
                    "name": "News"
                },
                {
                    "id": "5b2f284ba150181cbf747e52",
                    "string_id": "jakarta",
                    "name": "Jakarta"
                }
            ]
        }
    ]
}
```

### 4. Filter berita berdasarkan statusnya
Untuk mengaktifkan fitur filter berita berdasarkan parameter status ("konsep", "dihapus", "publikasikan"). Dengan parameter enumerasi dari ["status", "dihapus", "publikasikan].
**Handler Function: ./application/module/News.filterNewsByStatus** 

* Method : GET
* URL : http://kumparan-api.agungdwiprasetyo.com/news/filter/status/{{status_enum}}
* Response body example (misal get data dengan parameter status="konsep"):
```json
{
    "success": true,
    "total": 1,
    "data": [
        {
            "id": "5b2f409da15018732b3972cf",
            "title": "Kecelakaan mobil gandeng di tol",
            "value": "Telah terjadi kecelakaan dalam tol jakarta",
            "posted_at": "2018-06-24T13:56:29.507+07:00",
            "status": "konsep",
            "total_comment": 0,
            "total_like": 0,
            "topics": [
                {
                    "id": "5b2f284ba150181cbf747e50",
                    "string_id": "news",
                    "name": "News"
                },
                {
                    "id": "5b2f284ba150181cbf747e52",
                    "string_id": "jakarta",
                    "name": "Jakarta"
                }
            ]
        }
    ]
}
```