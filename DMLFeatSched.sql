INSERT INTO users(id username, password, role, specialization) VALUES('67b65471-eb1f-46ec-a043-959a5cc85778', 'urip', 'aaadasadadasaa', 'PATIENT', ''), ('5bc18dd0-58cb-4612-8dc3-5fc2419b7f29', 'martono', 'aaadasadadasaa', 'DOCTOR', 'ortopedi');


INSERT INTO mst_schedule_time(id, start_at, end_at) VALUES (1, '08:00:30', '08:30:00'), (2, '08:30:30', '09:00:00'), (3, '09:00:30', '09:30:00'), (4, '09:30:30', '10:00:00'), (5, '10:00:30', '10:30:00'), (6, '10:30:30', '11:00:00'), (7, '11:00:30', '11:30:00'), (8, '11:30:30', '12:00:00'), (9, '13:00:30', '13:30:00'), (10, '13:30:30', '14:00:00'), (11, '14:00:30', '14:30:00'), (12, '14:30:30', '15:00:00'), (13, '15:00:30', '15:30:00'), (14, '15:30:30', '16:00:00');


INSERT INTO doctor_schedules(id, doctor_id, schedule_date, start_at, end_at) VALUES('74d93144-6f2e-4bbc-9f89-973c62d3ac54', '5bc18dd0-58cb-4612-8dc3-5fc2419b7f29', '2024-03-14', 1, 8), ('e076f890-0ed3-4df2-bb57-e9dc6e48fc89', '5bc18dd0-58cb-4612-8dc3-5fc2419b7f29', '2024-03-15', 1, 9), ('a98528e4-3fa4-48e5-adaa-14fb3a21d981','5bc18dd0-58cb-4612-8dc3-5fc2419b7f29', '2024-03-16', 1, 9);



INSERT INTO bookings(doctor_schedule_id,patient_id, mst_schedule_id, complaint, status) VALUES ('74d93144-6f2e-4bbc-9f89-973c62d3ac54','67b65471-eb1f-46ec-a043-959a5cc85778', 1,'sakit perut melilit kek mau mateee', 'WAITING'), ('74d93144-6f2e-4bbc-9f89-973c62d3ac54','67b65471-eb1f-46ec-a043-959a5cc85778', 2,'sakit kepala dihantam realita', 'CANCELED');


