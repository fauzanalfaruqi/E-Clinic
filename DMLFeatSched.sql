INSERT INTO users(id username, password, role, specialization) VALUES('67b65471-eb1f-46ec-a043-959a5cc85778', 'urip', 'aaadasadadasaa', 'PATIENT', ''), ('5bc18dd0-58cb-4612-8dc3-5fc2419b7f29', 'martono', 'aaadasadadasaa', 'DOCTOR', 'ortopedi');


INSERT INTO doctor_schedules(id, doctor_id, day_of_week, start_at, end_at) VALUES('74d93144-6f2e-4bbc-9f89-973c62d3ac54', '5bc18dd0-58cb-4612-8dc3-5fc2419b7f29', 1, '09:00:00', '15:00:00'), ('5bc18dd0-58cb-4612-8dc3-5fc2419b7f29', 2, '09:00:00', '15:00:00'), ('a98528e4-3fa4-48e5-adaa-14fb3a21d981','5bc18dd0-58cb-4612-8dc3-5fc2419b7f29', 4, '09:00:00', '15:00:00')


INSERT INTO bookings(patient_id, schedule_id, complaint, start_at, end_at, status) VALUES ('67b65471-eb1f-46ec-a043-959a5cc85778', 'a98528e4-3fa4-48e5-adaa-14fb3a21d981', 'sakit perut melilit kek mau mateee', '10:00:00', '11:00:00', 'WAITING');
