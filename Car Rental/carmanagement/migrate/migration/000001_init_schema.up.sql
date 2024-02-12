CREATE TABLE IF NOT EXISTS "Car" (
  "id" serial PRIMARY KEY,
  "make" varchar,
  "model" varchar,
  "year" int,
  "license_plate" varchar UNIQUE,
  "color" varchar,
  "mileage" int,
  "status" varchar CHECK (status IN ('available', 'in use', 'maintenance')),
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

-- Sample data insertion
INSERT INTO "Car" ("make", "model", "year", "license_plate", "color", "mileage", "status", "created_at", "updated_at") VALUES
  ('Toyota', 'Camry', 2020, 'ABC123', 'Red', 15000, 'available', '2023-01-10 08:30:00', '2023-01-10 08:30:00'),
  ('Honda', 'Civic', 2019, 'XYZ456', 'Blue', 20000, 'in use', '2022-08-05 12:45:00', '2023-02-20 10:15:00'),
  ('Ford', 'Fusion', 2021, 'DEF789', 'Black', 10000, 'maintenance', '2022-12-20 16:20:00', '2023-01-05 09:00:00'),
  ('Chevrolet', 'Malibu', 2018, 'GHI101', 'Silver', 18000, 'available', '2023-03-15 11:00:00', '2023-03-15 11:00:00'),
  ('Nissan', 'Altima', 2017, 'JKL202', 'White', 22000, 'available', '2022-11-25 09:30:00', '2023-01-30 14:45:00'),
  ('Hyundai', 'Sonata', 2019, 'MNO303', 'Gray', 19000, 'in use', '2023-02-01 10:20:00', '2023-03-10 08:30:00'),
  ('Toyota', 'Corolla', 2022, 'PQR404', 'Blue', 12000, 'maintenance', '2022-12-02 14:00:00', '2023-02-28 13:00:00'),
  ('Honda', 'Accord', 2016, 'STU505', 'Black', 25000, 'available', '2022-10-10 07:45:00', '2023-01-15 09:30:00'),
  ('Ford', 'Focus', 2019, 'VWX606', 'Red', 20000, 'available', '2023-01-20 13:10:00', '2023-02-10 11:20:00'),
  ('Chevrolet', 'Impala', 2015, 'YZA707', 'Green', 28000, 'in use', '2022-09-15 15:30:00', '2023-02-05 10:00:00');
