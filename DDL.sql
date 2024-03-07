CREATE DATABASE IF NOT EXISTS avenger_clinic;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE user_role AS ENUM ('ADMIN', 'DOCTOR', 'PATIENT');

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
  start_date TIMESTAMP NOT NULL,
  end_date TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE medicines (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  name VARCHAR NOT NULL,
  medicine_type VARCHAR NOT NULL,
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

CREATE TABLE bookings (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  patient_id uuid NOT NULL REFERENCES users (id),
  doctor_id uuid NOT NULL REFERENCES users (id),
  schedule TIMESTAMP NOT NULL,
  complaINT text NOT NULL,
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