-- Decrease the length of the "user_id" column
ALTER TABLE user_settings ALTER COLUMN user_id TYPE CHAR(20);

-- Drop all the created tables
DROP TABLE bills;
DROP TABLE recurring_bills;
DROP TABLE bank_details;
DROP TABLE contacts;
DROP TABLE accounts;
DROP TABLE banks;
DROP TABLE categories;