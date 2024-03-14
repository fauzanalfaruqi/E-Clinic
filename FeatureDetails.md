# Doctor Schedule
## Get All
    - Get All doctor schedule disediakan untuk admin
    - Bisa mencari dengan parameter startDate(sd) & endDate(ed)

## Get My Schedule
    - Disediakan hanya untuk role dokter
    - Bisa mencari dengan query parameter :
        - startDate(sd) & endDate(ed)
        - day of week(dow) secara multiple([0,1,2])
        - booking(schedule) berdasarkan status('WAITING','DONE', 'CANCELED') secara multiple

## Get By ID
    - Disediakan untuk admin, pasien & dokter
    - Bisa mencari dengan query parameter :
        - booking(schedule) berdasarkan status('WAITING','DONE', 'CANCELED') secara multiple


## Insert
    - Disediakan untuk role dokter & admin
    - Insert beberapa data sekaligus
    - Validasi startAt harus < endAt
    - Validasi format tanggal di setiap data yang akan diinsert

## Update
    - Disediakan untuk role dokter & admin
    - Validasi startAt harus < endAt
    - Validasi format tanggal yang akan diupdate
    - Validasi tanggal yang diinsert apakah sudah ada di db

## Delete
    - Disediakan untuk role dokter & admin

## Restore
    - Disediakan untuk role dokter & admin


### Notes :
    - Seluruh Query Param sudah divalidasi agar tidak diinject sql
    - day of week(dow) : minggu=0 ... sabtu=6

### TODO Doctor Schedule :
    - validasi exist date saat insert 

