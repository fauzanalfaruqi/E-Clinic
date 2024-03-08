CREATE DATABASE IF NOT EXISTS avenger_clinic;

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

CREATE TABLE doctor_schedules (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  doctor_id uuid NOT NULL REFERENCES users (id),
  day_of_week INT NOT NULL,
  start_at TIME NOT NULL,
  end_at TIME NOT NULL,
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

CREATE TABLE mst_schedule(
  id INT PRIMARY KEY,
  start_at TIME NOT NULL,
  end_at TIME NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
CREATE TABLE bookings (
  id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
  patient_id uuid NOT NULL REFERENCES users (id),
  doctor_schedule_id uuid NOT NULL REFERENCES doctor_schedules(id),
  booking_date DATE NOT NULL,
  mst_schedule_id int NOT NULL REFERENCES mst_schedule(id),
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

