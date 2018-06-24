## Read Topic (GET in REST API)

### 1. Get list topic

Untuk mendapatkan list data topik dengan parameter partisi halaman untuk limitasi data (dengan url query ```page``` dan ```limit```). Jika parameter url query ini tidak disertakan, maka nilai default untuk parameter ```page``` = 1 dan parameter ```limit``` = 10. Hasil response data berupa model objek data topik lengkap dengan list berita dari masing-masing topik.

**Handler Function: ./application/module/Topic.getListTopic** 

* Method : GET
* URL : http://news-api.agungdwiprasetyo.com/topic/list?page=1&limit=10
* Response body example:
```json
{
    "success": true,
    "message": "Success",
    "page": 1,
    "limit": 10,
    "data_per_page": 8,
    "total_page": 1,
    "total_all_data": 8,
    "data": [
        {
            "id": "5b2f284ba150181cbf747e50",
            "string_id": "news",
            "name": "News",
            "news": [
                {
                    "id": "5b2f409da15018732b3972cf",
                    "title": "Kecelakaan mobil gandeng di tol",
                    "value": "Telah terjadi kecelakaan dalam tol jakarta",
                    "posted_at": "2018-06-24T13:56:29.507+07:00",
                    "status": "konsep",
                    "total_comment": 0,
                    "total_like": 0
                }
            ]
        },
        {
            "id": "5b2f284ba150181cbf747e51",
            "string_id": "megapolitan",
            "name": "Megapolitan",
            "news": [
                {
                    "id": "5b2f4081a15018732b3972ca",
                    "title": "CFD Dibuka Lebih Cepat Karena Banyak Warga Telah Kembali ke Jakarta",
                    "value": "Hari Bebas Kendaraan Bermotor atau Car Free Day (CFD) dibuka lebih cepat. Sebelumnya, Pemprov DKI sempat mengumumkan bakal meniadakan CFD pada 10, 17, dan 24 Juni 2018. Namun tanggal tersebut direvisi dan kembali dibuka Minggu (24/6) karena banyak penduduk yang telah kembali ke Jakarta setelah mudik Lebaran. Perubahan tanggal tersebut dilakukan jauh hari sebelum Lebaran. Kadishub DKI Jakarta Andri Yansyah mengatakan CFD kembali digelar karena masyarakat sudah kembali dari mudik. Maka itu minggu ini CFD sudah ramai dikunjungi masyarakat. ”Masyarakat sudah pulih. Yang balik mudik juga sudah kembali mereka butuh refreshing kan,” ungkap Andri saat dihubungi kumparan.",
                    "posted_at": "2018-06-24T13:56:01.055+07:00",
                    "status": "publikasikan",
                    "total_comment": 0,
                    "total_like": 0
                }
            ]
        },
        {
            "id": "5b2f284ba150181cbf747e52",
            "string_id": "jakarta",
            "name": "Jakarta",
            "news": [
                {
                    "id": "5b2f4081a15018732b3972ca",
                    "title": "CFD Dibuka Lebih Cepat Karena Banyak Warga Telah Kembali ke Jakarta",
                    "value": "Hari Bebas Kendaraan Bermotor atau Car Free Day (CFD) dibuka lebih cepat. Sebelumnya, Pemprov DKI sempat mengumumkan bakal meniadakan CFD pada 10, 17, dan 24 Juni 2018. Namun tanggal tersebut direvisi dan kembali dibuka Minggu (24/6) karena banyak penduduk yang telah kembali ke Jakarta setelah mudik Lebaran. Perubahan tanggal tersebut dilakukan jauh hari sebelum Lebaran. Kadishub DKI Jakarta Andri Yansyah mengatakan CFD kembali digelar karena masyarakat sudah kembali dari mudik. Maka itu minggu ini CFD sudah ramai dikunjungi masyarakat. ”Masyarakat sudah pulih. Yang balik mudik juga sudah kembali mereka butuh refreshing kan,” ungkap Andri saat dihubungi kumparan.",
                    "posted_at": "2018-06-24T13:56:01.055+07:00",
                    "status": "publikasikan",
                    "total_comment": 0,
                    "total_like": 0
                },
                {
                    "id": "5b2f409da15018732b3972cf",
                    "title": "Kecelakaan mobil gandeng di tol",
                    "value": "Telah terjadi kecelakaan dalam tol jakarta",
                    "posted_at": "2018-06-24T13:56:29.507+07:00",
                    "status": "konsep",
                    "total_comment": 0,
                    "total_like": 0
                }
            ]
        },
        {
            "id": "5b2f284ba150181cbf747e53",
            "string_id": "car-free-day",
            "name": "Car Free Day",
            "news": [
                {
                    "id": "5b2f4081a15018732b3972ca",
                    "title": "CFD Dibuka Lebih Cepat Karena Banyak Warga Telah Kembali ke Jakarta",
                    "value": "Hari Bebas Kendaraan Bermotor atau Car Free Day (CFD) dibuka lebih cepat. Sebelumnya, Pemprov DKI sempat mengumumkan bakal meniadakan CFD pada 10, 17, dan 24 Juni 2018. Namun tanggal tersebut direvisi dan kembali dibuka Minggu (24/6) karena banyak penduduk yang telah kembali ke Jakarta setelah mudik Lebaran. Perubahan tanggal tersebut dilakukan jauh hari sebelum Lebaran. Kadishub DKI Jakarta Andri Yansyah mengatakan CFD kembali digelar karena masyarakat sudah kembali dari mudik. Maka itu minggu ini CFD sudah ramai dikunjungi masyarakat. ”Masyarakat sudah pulih. Yang balik mudik juga sudah kembali mereka butuh refreshing kan,” ungkap Andri saat dihubungi kumparan.",
                    "posted_at": "2018-06-24T13:56:01.055+07:00",
                    "status": "publikasikan",
                    "total_comment": 0,
                    "total_like": 0
                }
            ]
        },
        {
            "id": "5b2f33a3a150181c75aa5c18",
            "string_id": "kecelakaan",
            "name": "Kecelakaan"
        },
        {
            "id": "5b2f4ab8a1501832ac4dac1b",
            "string_id": "cuaca",
            "name": "CUACA"
        },
        {
            "id": "5b2f78b6a150187e295875ff",
            "string_id": "pemilu",
            "name": "Pemilu",
            "news": [
                {
                    "id": "5b2f78b6a150187e295875fe",
                    "title": "Pemilu 2018",
                    "value": "Pemilu serentak 2018 yang dilaksanakan pada hari rabu, 27 Juni 2018 mendatang dikabarkan aman",
                    "posted_at": "2018-06-24T17:55:50.702+07:00",
                    "status": "konsep",
                    "total_comment": 0,
                    "total_like": 0
                }
            ]
        },
        {
            "id": "5b2f78b6a150187e29587600",
            "string_id": "pilkada",
            "name": "Pilkada",
            "news": [
                {
                    "id": "5b2f78b6a150187e295875fe",
                    "title": "Pemilu 2018",
                    "value": "Pemilu serentak 2018 yang dilaksanakan pada hari rabu, 27 Juni 2018 mendatang dikabarkan aman",
                    "posted_at": "2018-06-24T17:55:50.702+07:00",
                    "status": "konsep",
                    "total_comment": 0,
                    "total_like": 0
                }
            ]
        }
    ]
}
```

### 2. Get All Topic

Untuk mendapatkan seluruh data topic yang ada dalam database, nantinya untuk memilih topik-topik yang tersedia untuk data berita.

**Handler Function: ./application/module/Topic.getAllTopic** 

* Method : GET
* URL : http://news-api.agungdwiprasetyo.com/topic/all
* Response body example:
```json
{
    "success": true,
    "total": 8,
    "data": [
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
        },
        {
            "id": "5b2f33a3a150181c75aa5c18",
            "string_id": "kecelakaan",
            "name": "Kecelakaan"
        },
        {
            "id": "5b2f4ab8a1501832ac4dac1b",
            "string_id": "cuaca",
            "name": "CUACA"
        },
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

### 3. Get Topic by ID

Untuk mendapatkan suatu objek data topik berdasarkan stringID-nya.

**Handler Function: ./application/module/Topic.getTopicByStringID**

* Method : GET
* URL : http://news-api.agungdwiprasetyo.com/topic/find/{{topic_string_id}}
* Response body example (misal get data dengan parameter topic_string_id="pemilu"):
```json
{
    "success": true,
    "data": {
        "id": "5b2f78b6a150187e295875ff",
        "string_id": "pemilu",
        "name": "Pemilu",
        "news": [
            {
                "id": "5b2f78b6a150187e295875fe",
                "title": "Pemilu 2018",
                "value": "Pemilu serentak 2018 yang dilaksanakan pada hari rabu, 27 Juni 2018 mendatang dikabarkan aman",
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
        ]
    }
}
```