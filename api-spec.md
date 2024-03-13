# Avenger Clinic API Specs

## Samples

1. ### User Endpoints

- #### Register new patient

    Method

        POST

    Endpoint

        localhost:8080/api/v1/register

    Sample body request

        {
            "name":"",
            "username":"",
            "password":"",
            "role":"",
            "address":"",
            "phone_number":""
        }

    Sample response

        {
            "code":"",
            "message",
            "data":
                {
                    "id":"",
                    "name":"",
                    "username":"",
                    "role":"",
                    "address":"",
                    "phone_number":""
                }
        }

- #### Register new doctor

    Method

        POST

    Endpoint

        localhost:8080/api/v1/users/doctor

    Sample body request

        {
            "name":"",
            "username":"",
            "password":"",
            "role":"",
            "spesialization":"",
            "consultation_fee":""
        }

    Sample response

        {
            "code":"",
            "message",
            "data":
                {
                    "id":"",
                    "name":"",
                    "username":"",
                    "role":"",
                    "spesialization":"",
                    "consultation_fee":""
                    
                }
        }

- #### Register new admin

    Method

        POST

        Endpoint

            localhost:8080/api/v1/users/doctor

        Sample body request

            {
                "username":"",
                "password":"",
                "role":""
            }

        Sample response

            {
                "code":"",
                "message",
                "data":
                    {
                        "name":"",
                        "username":"",
                        "role":""
                    }
            }

- #### Login

    Method

        POST

    Endpoint

        localhost:8080/api/v1/users/doctor

    Sample body request

        {
            "username":"",
            "password":"",
            "role":""
        }

    Sample response

        {
            "code":"",
            "message",
            "data":
                {
                    "token":""
                }
        }

2. ### Patient Endpoints

- Create booking


3. ### Doctor Endpoints

- Create medical records

    Method

        POST

    Endpoint

        localhost:8080/api/v1/doctor

    Sample body request

        {
            "booking_id":"",
            "diagnosis_result":"",
            "items": [
                {
                    "medicine_id":"",
                    "quantity":""
                },
                {
                    "medicine_id":"",
                    "quantity":""
                }
            ]
        }

    Sample response

        {
            "id":"",
            "booking_id":"",
            "diagnosis_result":"",
            "items": [
                {
                    "medicine_id":"",
                    "quantity":""
                },
                {
                    "medicine_id":"",
                    "quantity":""
                }
            ]
        }

- Get medical records

    Method

        GET

    Endpoint

        localhost:8080/api/v1/doctor

    Sample response

        [
            {
                "id":"",
                "booking_id":"",
                "diagnosis_result":"",
                "items": [
                    {
                        "medicine_id":"",
                        "quantity":""
                    },
                    {
                        "medicine_id":"",
                        "quantity":""
                    }
                ]
            },
            {
                "id":"",
                "booking_id":"",
                "diagnosis_result":"",
                "items": [
                    {
                        "medicine_id":"",
                        "quantity":""
                    },
                    {
                        "medicine_id":"",
                        "quantity":""
                    }
                ]
            }
        ]

- Get medical record by id

    Method

        POST

    Endpoint

        localhost:8080/api/v1/doctor/:id

    Sample response

        {
            "id":"",
            "booking_id":"",
            "diagnosis_result":"",
            "items": [
                {
                    "medicine_id":"",
                    "quantity":""
                },
                {
                    "medicine_id":"",
                    "quantity":""
                }
            ]
        }

4. #### Medicine Enpoints

- Create medicine

    Method

        POST

    Endpoint

        localhost:8080/api/v1/medicines

    Sample body request

        {
            "name":"",
            "type":"",
            "price":"",
            "stock":"",
            "description":""
        }

    Sample response

        {
            "name":"",
            "type":"",
            "price":"",
            "stock":"",
            "description":""
        }

- Get medicines

    Method

        GET

    Endpoint

        localhost:8080/api/v1/medicines

    Sample response

        [
            {
                "name":"",
                "type":"",
                "price":"",
                "stock":"",
                "description":"",
                "created_at":"",
                "updated_at":""
            },
            {
                "name":"",
                "type":"",
                "price":"",
                "stock":"",
                "description":"",
                "created_at":"",
                "updated_at":""
            }
        ]

- Get medicine by id

    Method

        GET

    Endpoint

        localhost:8080/api/v1/medicines

    Sample response

        {
            "name":"",
            "type":"",
            "price":"",
            "stock":"",
            "description":"",
            "created_at":"",
            "updated_at":""
        }

- Update medicine by id

    Method

        PUT

    Endpoint

        localhost:8080/api/v1/medicines

    Sample body request

        {
            "name":"",
            "type":"",
            "price":"",
            "stock":"",
            "description":""
        }

    Sample response

        {
            "name":"",
            "type":"",
            "price":"",
            "stock":"",
            "description":""
        }

- Delete medicine by id

    Method

        DELETE

    Endpoint

        localhost:8080/api/v1/medicines/:id

    Sample response
    
        {
            "data": "OK",
            "message": "Customer deleted successfully"
        }

5. #### Bill Features Endpoints

- Create transaction

    Method

        POST

    Endpoint

        localhost:8080/api/v1/bills

    Sample body request

        {
            "total_amount":"",
            "payment_status":""
        }

    Sample response

        {
            "id":"",
            "total_amoutn":"",
            "diagnosis_result":""
        }

- Get transactions

    Method

        GET

    Endpoint

        localhost:8080/api/v1/bills

    Sample response

        [
            {
                "medical_record_id":"",
                "total_amount":"",
                "payment_status":"",
                "created_at":"",
                "updated_at":""
            },
            {
                "medical_record_id":"",
                "total_amount":"",
                "payment_status":"",
                "created_at":"",
                "updated_at":""
            }
        ]

- Get transaction by id

    Method

        GET

    Endpoint

        localhost:8080/api/v1/bills/:id

    Sample response

    {
        "medical_record_id":"",
        "total_amount":"",
        "payment_status":"",
        "created_at":"",
        "updated_at":""
    },
    {
        "medical_record_id":"",
        "total_amount":"",
        "payment_status":"",
        "created_at":"",
        "updated_at":""
    }