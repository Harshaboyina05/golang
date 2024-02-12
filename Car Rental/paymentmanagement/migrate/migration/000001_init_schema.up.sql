CREATE TABLE IF NOT EXISTS "Payment" (
  "id" serial PRIMARY KEY,
  "reservation_id" int,
  "amount" decimal(10,2),
  "payment_date" timestamp DEFAULT CURRENT_TIMESTAMP,
  "payment_method" varchar,
  "status" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO "Payment" ("id", "reservation_id", "amount", "payment_date", "payment_method", "status", "created_at", "updated_at") VALUES
  (1, 1, 150.00, '2024-01-15 10:30:00', 'Credit Card', 'completed', '2024-01-15 10:30:00', '2024-01-15 10:30:00'),
  (2, 2, 200.50, '2024-01-20 11:45:00', 'PayPal', 'completed', '2024-01-20 11:45:00', '2024-01-20 11:45:00'),
  (3, 3, 100.75, '2024-02-05 09:15:00', 'Credit Card', 'failed', '2024-02-05 09:15:00', '2024-02-05 09:15:00'),
  (4, 4, 175.25, '2024-02-20 13:30:00', 'Credit Card', 'completed', '2024-02-20 13:30:00', '2024-02-20 13:30:00'),
  (5, 5, 120.00, '2024-02-25 10:00:00', 'PayPal', 'completed', '2024-02-25 10:00:00', '2024-02-28 11:45:00'),
  (6, 6, 250.00, '2024-03-10 12:45:00', 'Credit Card', 'completed', '2024-03-10 12:45:00', '2024-03-15 09:30:00'),
  (7, 7, 180.50, '2024-03-20 08:00:00', 'Credit Card', 'completed', '2024-03-20 08:00:00', '2024-03-25 10:15:00'),
  (8, 8, 300.00, '2024-04-01 11:30:00', 'PayPal', 'completed', '2024-04-01 11:30:00', '2024-04-01 11:30:00'),
  (9, 9, 150.75, '2024-04-10 09:45:00', 'Credit Card', 'completed', '2024-04-10 09:45:00', '2024-04-15 09:00:00'),
  (10, 10, 210.25, '2024-04-20 13:00:00', 'PayPal', 'completed', '2024-04-20 13:00:00', '2024-04-20 13:00:00');
