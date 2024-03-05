CREATE DATABASE IF NOT EXISTS avenger_clinic;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE role AS ENUM ('ADMIN', 'DOCTOR', 'PATIENT');

CREATE TABLE users (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  username varchar NOT NULL,
  email varchar NOT NULL,
  password varchar NOT NULL,
  role role NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp,
  deleted_at timestamp
);

CREATE TABLE doctors (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  user_id uuid NOT NULL,
  name varchar NOT NULL,
  specialization varchar NOT NULL,
  consultation_fee int NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp,
  deleted_at timestamp
);

CREATE TABLE doctor_schedules (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  doctor_id uuid NOT NULL,
  start_date timestamp NOT NULL,
  end_date timestamp NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp,
  deleted_at timestamp
);

CREATE TABLE patients (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  user_id uuid NOT NULL,
  name varchar NOT NULL,
  address text,
  phone_number varchar NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp,
  deleted_at timestamp
);

CREATE TABLE medicines (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  name varchar NOT NULL,
  medicine_type varchar NOT NULL,
  price int NOT NULL,
  stock int DEFAULT 0,
  description text,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp,
  deleted_at timestamp
);

CREATE TABLE bookings (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  patient_id uuid NOT NULL,
  doctor_id uuid NOT NULL,
  schedule timestamp NOT NULL,
  complaint text NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp,
  deleted_at timestamp
);

CREATE TABLE medical_records (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  booking_id uuid NOT NULL,
  diagnosis_results text NOT NULL,
  consultation_fee int,
  total_amount int,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp,
  deleted_at timestamp
);

CREATE TABLE medical_record_details (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  medical_record_id uuid NOT NULL,
  medicine_id uuid NOT NULL,
  medicine_price int,
  quantity int NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp,
  deleted_at timestamp
);

CREATE TABLE bills (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  medical_record_id uuid NOT NULL,
  bill_date timestamp NOT NULL,
  total_amount int,
  payment_status bool,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp,
  deleted_at timestamp
);

ALTER TABLE doctors ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE doctor_schedules ADD FOREIGN KEY (doctor_id) REFERENCES doctors (id);

ALTER TABLE patients ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE bookings ADD FOREIGN KEY (doctor_id) REFERENCES doctors (id);

ALTER TABLE medical_records ADD FOREIGN KEY (booking_id) REFERENCES bookings (id);

ALTER TABLE bills ADD FOREIGN KEY (medical_record_id) REFERENCES medical_records (id);

ALTER TABLE bookings ADD FOREIGN KEY (patient_id) REFERENCES patients (id);

ALTER TABLE medical_record_details ADD FOREIGN KEY (medical_record_id) REFERENCES medical_records (id);

ALTER TABLE medical_record_details ADD FOREIGN KEY (medicine_id) REFERENCES medicines (id);