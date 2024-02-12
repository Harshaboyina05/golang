CREATE TABLE IF NOT EXISTS "Reservation" (
  "id" serial PRIMARY KEY,
  "car_id" int,
  "customer_id" int,
  "start_date" timestamp DEFAULT CURRENT_TIMESTAMP,
  "end_date" timestamp DEFAULT CURRENT_TIMESTAMP,
  "status" varchar,
  "payment_status" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO "Reservation" ("car_id", "customer_id", "start_date", "end_date", "status", "payment_status", "created_at", "updated_at") VALUES
  (1, 1, '2024-02-01 09:00:00', '2024-02-05 17:00:00', 'confirmed', 'paid', '2024-01-15 10:30:00', '2024-01-15 10:30:00'),
  (2, 2, '2024-02-10 08:30:00', '2024-02-15 16:00:00', 'pending', 'pending', '2024-01-20 11:45:00', '2024-01-20 11:45:00'),
  (3, 3, '2024-03-05 10:00:00', '2024-03-10 14:00:00', 'cancelled', 'paid', '2024-02-05 09:15:00', '2024-02-15 09:00:00'),
  (4, 4, '2024-03-20 11:30:00', '2024-03-25 18:00:00', 'pending', 'pending', '2024-02-20 13:30:00', '2024-02-20 13:30:00'),
  (5, 5, '2024-04-02 08:00:00', '2024-04-06 16:30:00', 'confirmed', 'paid', '2024-02-25 10:00:00', '2024-02-28 11:45:00'),
  (6, 6, '2024-04-15 10:30:00', '2024-04-20 15:00:00', 'confirmed', 'paid', '2024-03-10 12:45:00', '2024-03-15 09:30:00'),
  (7, 7, '2024-05-01 09:00:00', '2024-05-05 17:30:00', 'confirmed', 'paid', '2024-03-20 08:00:00', '2024-03-25 10:15:00'),
  (8, 8, '2024-05-10 08:30:00', '2024-05-15 16:00:00', 'pending', 'pending', '2024-04-01 11:30:00', '2024-04-01 11:30:00'),
  (9, 9, '2024-06-05 10:00:00', '2024-06-10 14:00:00', 'confirmed', 'paid', '2024-04-10 09:45:00', '2024-04-15 09:00:00'),
  (10, 10, '2024-06-20 11:30:00', '2024-06-25 18:00:00', 'pending', 'pending', '2024-04-20 13:00:00', '2024-04-20 13:00:00');
