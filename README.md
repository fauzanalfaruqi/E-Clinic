# Avenger Clinic - Golang Backend Clinic Management System

Avenger Clinic is a backend clinic management system built using Go. This project was created during final project of the Go bootcamp at [Enigmacamp](https://enigmacamp.com/). Implements clean architecture, the program can perform CRUD operations and Authorization using JWT.

## Get Started

- ### Database and Tables creation

  To create required tables for this program to run properly, you can use this sample ddl queries which can be found [here](./DDL.sql). The query looks like this:

  ```sql
  CREATE DATABASE IF NOT EXISTS avenger_clinic_db;
  
  CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
  
  CREATE TYPE user_role AS ENUM ('ADMIN', 'DOCTOR', 'PATIENT');
  
  CREATE TYPE booking_status AS ENUM('WAITING', 'CANCELED', 'DONE');
  
  CREATE TYPE medicine_type AS ENUM('CAIR','TABLET','OLES','TETES','KAPSUL');
  
  CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    username VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    role user_role NOT NULL,
    specialization VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
  );
  
  CREATE TABLE medicines (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR NOT NULL,
    medicine_type medicine_type not null,
    price INT NOT NULL,
    stock INT DEFAULT 0,
    description text,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
  );
  
  CREATE TABLE actions (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    price INT NOT NULL,
    description text,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
  );
  
  CREATE TABLE mst_schedule_time(
    id INT PRIMARY KEY,
    start_at TIME NOT NULL,
    end_at TIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
  );
  
  CREATE TABLE doctor_schedules (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    doctor_id uuid NOT NULL REFERENCES users (id),
    schedule_date DATE NOT NULL,
    start_at INT NOT NULL REFERENCES mst_schedule_time(id),
    end_at INT NOT NULL REFERENCES mst_schedule_time(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
  );
  
  CREATE TABLE bookings (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    patient_id uuid NOT NULL REFERENCES users (id),
    doctor_schedule_id uuid NOT NULL REFERENCES doctor_schedules(id),
    mst_schedule_id int NOT NULL REFERENCES mst_schedule_time(id),
    status booking_status NOT NULL,
    complaint text NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
  );
  
  CREATE TABLE medical_records (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    booking_id uuid NOT NULL REFERENCES bookings (id),
    diagnosis_results text NOT NULL,
    total_medicine int,
    total_action int,
    total_amount int,
    payment_status bool,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
  );
  
  CREATE TABLE medical_record_medicine_details (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    medical_record_id uuid NOT NULL REFERENCES medical_records (id),
    medicine_id uuid NOT NULL REFERENCES medicines (id),
    medicine_price int,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
  );
  
  CREATE TABLE medical_record_action_details (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    medical_record_id uuid NOT NULL REFERENCES medical_records (id),
    action_id uuid NOT NULL REFERENCES actions (id),
    action_price int,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
  );
  ```

- ### Sample Data

  To populate the database with initial/necessary data, you can copy provided dml sample data query which can be found [here](./DMLFeatSched.sql). The query is looks like this:

  ```sql
  -- Sample dml
  INSERT INTO users(id, username, password, role, specialization, created_at, updated_at) 
  VALUES
      ('31b24cdd-c633-4d2d-9044-718378eb3929', 'admin', '$2a$10$bmiD3Nuo3R7CXHTiQcsLFeEhGhNkx6vLfcN50gKgNu6/v.qlWDSZm', 'ADMIN', NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      ('67b65471-eb1f-46ec-a043-959a5cc85778', 'Budi', '$2a$10$bmiD3Nuo3R7CXHTiQcsLFeEhGhNkx6vLfcN50gKgNu6/v.qlWDSZm', 'PATIENT', NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      ('5bc18dd0-58cb-4612-8dc3-5fc2419b7f29', 'Joko', '$2a$10$bmiD3Nuo3R7CXHTiQcsLFeEhGhNkx6vLfcN50gKgNu6/v.qlWDSZm', 'DOCTOR', 'Ortopedi', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
  
  INSERT INTO mst_schedule_time(id, start_at, end_at, created_at, updated_at) 
  VALUES
      (1, '08:00:30', '08:30:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      (2, '08:30:30', '09:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      (3, '09:00:30', '09:30:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      (4, '09:30:30', '10:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      (5, '10:00:30', '10:30:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      (6, '10:30:30', '11:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      (7, '11:00:30', '11:30:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      (8, '11:30:30', '12:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      (9, '13:00:30', '13:30:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      (10, '13:30:30', '14:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      (11, '14:00:30', '14:30:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      (12, '14:30:30', '15:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      (13, '15:00:30', '15:30:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      (14, '15:30:30', '16:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);


  INSERT INTO doctor_schedules(id, doctor_id, schedule_date, start_at, end_at, created_at, updated_at) 
  VALUES
      ('74d93144-6f2e-4bbc-9f89-973c62d3ac54', '5bc18dd0-58cb-4612-8dc3-5fc2419b7f29', '2024-03-14', 1, 8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      ('e076f890-0ed3-4df2-bb57-e9dc6e48fc89', '5bc18dd0-58cb-4612-8dc3-5fc2419b7f29', '2024-03-15', 1, 9, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      ('a98528e4-3fa4-48e5-adaa-14fb3a21d981','5bc18dd0-58cb-4612-8dc3-5fc2419b7f29', '2024-03-16', 1, 9, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

  INSERT INTO bookings(doctor_schedule_id,patient_id, mst_schedule_id, complaint, status, created_at, updated_at)
  VALUES
      ('74d93144-6f2e-4bbc-9f89-973c62d3ac54','67b65471-eb1f-46ec-a043-959a5cc85778', 1,'sakit perut melilit kek mau mateee', 'WAITING', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
      ('74d93144-6f2e-4bbc-9f89-973c62d3ac54','67b65471-eb1f-46ec-a043-959a5cc85778', 2,'sakit kepala dihantam realita', 'CANCELED', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
  ```
  
- ### Create `env` File

  Create `.env` file in the root directory of this project. Copy and modify env settings below into your created `.env` file:

  ```json
  DB_HOST=localhost
  DB_PORT=5432
  DB_USER=postgres
  DB_PASS=your_db_password
  DB_NAME=your_db_name
  MAX_IDLE=1
  MAX_CONN=2
  MAX_LIFE_TIME=1h
  
  PORT=8080
  LOG_MODE=1
  ```

- ### Run the program

  To run the program, simply excecute this command below on your terminal:

  ```shell
  go run .
  ```

## Endpoints

- ### Users

  | Method | Description                                     | Endpoint                     | Role                   |
  | ------ | ----------------------------------------------- | :--------------------------- | ---------------------- |
  | POST   | Insert new user for role *Admin* and *Doctor*   | /api/v1/users                | Admin                  |
  | POST   | Insert new user with role as *Patient*          | /api/v1/users/register       | Public                 |
  | POST   | Login user                                      | /api/v1/users/login          | Public                 |
  | GET    | Get all user records                            | /api/v1/users/getall         | Admin                  |
  | GET    | Get user record based on the given id           | /api/v1/users/{:id}          | Admin, Doctor, Patient |
  | GET    | Get soft deleted users                          | /api/v1/users/trash          | Admin                  |
  | PUT    | Update user record                              | /api/v1/users/{:id}          | Admin, Doctor, Aptient |
  | PUT    | Update password                                 | /api/v1/users/{:id}/password | Admin, Doctor, Patient |
  | PUT    | Restore soft deleted user based on the given id | /api/v1/users/{:id}/restore  | Admin                  |
  | DELETE | Hard delete user based on the given id          | /api/v1/users/{:id}          | Admin                  |
  | DELETE | Soft delete user based on the given id          | /api/v1/users/{:id}/trash    | Admin                  |

- ### Booking

  | Methods | Description                              | Endpoint                     | Role                   |
  | ------- | ---------------------------------------- | ---------------------------- | ---------------------- |
  | POST    | Create new booking record                | /api/v1/booking              | Admin, Patient         |
  | GET     | Get all booking fields                   | /api/v1/booking              | Admin, Doctor          |
  | GET     | Get booking fields based on the given id | /api/v1/booking/{:id}        | Admin, Doctor, Patient |
  | PUT     | Update booking schedule data             | /api/v1/booking/{:id}        | Admin, Patient         |
  | PUT     | Update booking  to mark as done          | /api/v1/booking/done/{:id}   | Admin, Doctor, Patient |
  | PUT     | Cancel booking                           | /api/v1/booking/cancel/{:id} | Admin, Patient         |

- ### Doctor Schedule

  | Method | Description                                      | Endpoint                      | Role                   |
  | ------ | ------------------------------------------------ | ----------------------------- | ---------------------- |
  | POST   | Insert new doctor schedule record                | /api/v1/doctor-schedule       | Admin, Doctor          |
  | GET    | Get all doctor schedule records                  | /api/v1/doctor-schedule       | Admin, Patient         |
  | GET    | Get doctor schedule record based on the given id | /api/v1/doctor-schedule/{:id} | Admin, Doctor, Patient |
  | PUT    | Update doctor schedule record                    | /api/v1/doctor-schedule/{:id} | Admin, Doctor          |
  | DELETE | Soft delete doctor schedule record               | /api/v1/doctor-schedule       | Admin, Doctor          |
  | PUT    | Restore soft deleted doctor schedule record      | /api/v1/doctor-schedule/{:id} | Admin, Doctor          |

- ### Medical Record

  | Method | Description                                     | Endpoint                     | Role          |
  | ------ | ----------------------------------------------- | ---------------------------- | ------------- |
  | POST   | Insert new medical record fields                | /api/v1/medical-record       | Admin, Doctor |
  | GET    | Get all medical record fields                   | /api/v1/medical-record       | Admin         |
  | GET    | Get medical record fields based on the given id | /api/v1/medical-record/{:id} | Admin         |
  | PUT    | Update payment status to done                   | /api/v1/medical-record/{:id} | Admin         |

- ### Medicine

  | Method | Description                               | Endpoint                        | Role  |
  | ------ | ----------------------------------------- | ------------------------------- | ----- |
  | POST   | Inser new medicine record                 | /api/v1/medicines               | Admin |
  | GET    | Get all medicine records                  | /api/v1/medicines               | Admin |
  | GET    | Get medicine record based on the given id | /api/v1/medicines/{:id}         | Admin |
  | PUT    | Update medicine record                    | /api/v1/medicines/{:id}         | Admin |
  | DELETE | Soft delete medicine record               | /api/v1/medicines/{:id}         | Admin |
  | GET    | Get soft delete medicine record           | /api/v1/medicines/trash         | Admin |
  | PUT    | Restore soft deleted medicine record      | /api/v1/medicines/{:id}/restore | Admin |

  - ### Action

  | Method | Description                             | Endpoint                      | Role  |
  | ------ | --------------------------------------- | ----------------------------- | ----- |
  | POST   | Inser new action record                 | /api/v1/actions               | Admin |
  | GET    | Get all action records                  | /api/v1/actions               | Admin |
  | GET    | Get action record based on the given id | /api/v1/actions/{:id}         | Admin |
  | PUT    | Update action record                    | /api/v1/actions/{:id}         | Admin |
  | DELETE | Soft delete action record               | /api/v1/actions/{:id}         | Admin |
  | GET    | Get soft delete action record           | /api/v1/actions/trash         | Admin |
  | PUT    | Restore soft deleted action record      | /api/v1/actions/{:id}/restore | Admin |

## Depencecies

This project uses these packages and all of its dependencies:

- [Gin](https://github.com/gin-gonic/gin): Gin is a HTTP web framework written in Go (Golang).
- [GoDotEnv](https://github.com/joho/godotenv): A Go (golang) port of the Ruby [dotenv](https://github.com/bkeepers/dotenv) project (which loads env vars from a .env file).
- [pq](https://github.com/lib/pq): Pure Go Postgres driver for database/sql.

- [jwt-go](https://github.com/dgrijalva/jwt-go): Golang implementation of JSON Web Tokens (JWT).
- [uuid](https://github.com/google/uuid): Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services.

- [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock): **sqlmock** is a mock library implementing [sql/driver](https://godoc.org/database/sql/driver).
- [validator](https://github.com/go-playground/validator/v10): Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.
- [crypto](https://golang.org/x/crypto): Go supplementary cryptography libraries.
- [zerolog](https://github.com/rs/zerolog): Zero allocation JSON logger.
- [cors](https://github.com/gin-contrib/cors): Official CORS gin's middleware.
- [logger](https://github.com/gin-contrib/logger): Gin middleware/handler to logger url path using rs/zerolog.
- [Testify](https://github.com/stretchr/testify): A toolkit with common assertions and mocks that plays nicely with the standard library.
