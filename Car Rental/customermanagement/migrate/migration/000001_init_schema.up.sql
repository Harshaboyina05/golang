CREATE TABLE IF NOT EXISTS "Customer" (
  "id" serial PRIMARY KEY,
  "first_name" varchar,
  "last_name" varchar,
  "email" varchar,
  "phone_number" varchar,
  "address" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO "Customer" ("first_name", "last_name", "email", "phone_number", "address", "created_at", "updated_at") VALUES
  ('John', 'Doe', 'john.doe@example.com', '123-456-7890', '123 Main St, Anytown, USA', '2024-01-15 10:30:00', '2024-01-15 10:30:00'),
  ('Jane', 'Smith', 'jane.smith@example.com', '987-654-3210', '456 Oak Ave, Othertown, USA', '2024-01-20 11:45:00', '2024-01-20 11:45:00'),
  ('Michael', 'Johnson', 'michael.johnson@example.com', '555-555-5555', '789 Elm St, Anothertown, USA', '2024-02-05 09:15:00', '2024-02-05 09:15:00'),
  ('Emily', 'Williams', 'emily.williams@example.com', '111-222-3333', '321 Maple Ave, Yetanothertown, USA', '2024-02-20 13:30:00', '2024-02-20 13:30:00'),
  ('Daniel', 'Brown', 'daniel.brown@example.com', '444-444-4444', '555 Pine St, Lasttown, USA', '2024-02-25 10:00:00', '2024-02-28 11:45:00'),
  ('Sarah', 'Jones', 'sarah.jones@example.com', '777-888-9999', '888 Cedar Ave, Finaltown, USA', '2024-03-10 12:45:00', '2024-03-15 09:30:00'),
  ('David', 'Miller', 'david.miller@example.com', '333-333-3333', '999 Oak St, Endtown, USA', '2024-03-20 08:00:00', '2024-03-25 10:15:00'),
  ('Jessica', 'Wilson', 'jessica.wilson@example.com', '666-666-6666', '777 Elm Ave, Nearingtown, USA', '2024-04-01 11:30:00', '2024-04-01 11:30:00'),
  ('Alex', 'Martinez', 'alex.martinez@example.com', '222-222-2222', '123 Oak St, Beginningtown, USA', '2024-04-10 09:45:00', '2024-04-15 09:00:00'),
  ('Sophia', 'Taylor', 'sophia.taylor@example.com', '999-999-9999', '456 Pine Ave, Firsttown, USA', '2024-04-20 13:00:00', '2024-04-20 13:00:00');
