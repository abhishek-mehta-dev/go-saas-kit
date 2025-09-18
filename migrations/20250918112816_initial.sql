-- Create "migrations" table
CREATE TABLE "migrations" (
  "id" bigserial NOT NULL,
  "name" text NOT NULL,
  "applied_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_migrations_name" UNIQUE ("name")
);
-- Create "plans" table
CREATE TABLE "plans" (
  "id" bigserial NOT NULL,
  "name" smallint NOT NULL,
  "period" smallint NOT NULL,
  "price" numeric NOT NULL,
  "description" character varying(255) NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_plans_deleted_at" to table: "plans"
CREATE INDEX "idx_plans_deleted_at" ON "plans" ("deleted_at");
-- Create "users" table
CREATE TABLE "users" (
  "id" bigserial NOT NULL,
  "first_name" character varying(100) NOT NULL,
  "last_name" character varying(100) NULL,
  "user_name" character varying(50) NOT NULL,
  "email" character varying(100) NOT NULL,
  "password" text NOT NULL,
  "user_type" character varying(20) NOT NULL DEFAULT 'member',
  "stripe_customer_id" character varying(100) NULL,
  "stripe_payment_method_id" character varying(100) NULL,
  "card_last4" character varying(4) NULL,
  "card_brand" character varying(50) NULL,
  "refresh_token" character varying(255) NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "users" ("email");
-- Create index "idx_users_user_name" to table: "users"
CREATE UNIQUE INDEX "idx_users_user_name" ON "users" ("user_name");
-- Create "subscriptions" table
CREATE TABLE "subscriptions" (
  "id" bigserial NOT NULL,
  "user_id" bigint NOT NULL,
  "status" smallint NOT NULL DEFAULT 0,
  "next_payment_date" timestamptz NULL,
  "is_free_trial" boolean NULL DEFAULT false,
  "is_renew_cancelled" boolean NULL DEFAULT false,
  "plan_id" bigint NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_plans_subscriptions" FOREIGN KEY ("plan_id") REFERENCES "plans" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_users_subscriptions" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_subscriptions_deleted_at" to table: "subscriptions"
CREATE INDEX "idx_subscriptions_deleted_at" ON "subscriptions" ("deleted_at");
-- Create index "idx_subscriptions_plan_id" to table: "subscriptions"
CREATE INDEX "idx_subscriptions_plan_id" ON "subscriptions" ("plan_id");
-- Create index "idx_subscriptions_user_id" to table: "subscriptions"
CREATE INDEX "idx_subscriptions_user_id" ON "subscriptions" ("user_id");
-- Create "payments" table
CREATE TABLE "payments" (
  "id" bigserial NOT NULL,
  "user_id" bigint NOT NULL,
  "subscription_id" bigint NULL,
  "amount" numeric NOT NULL,
  "stripe_payment_status" character varying(50) NULL,
  "stripe_payment_intent" character varying(100) NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_payments_subscription" FOREIGN KEY ("subscription_id") REFERENCES "subscriptions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_users_payments" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_payments_deleted_at" to table: "payments"
CREATE INDEX "idx_payments_deleted_at" ON "payments" ("deleted_at");
-- Create index "idx_payments_subscription_id" to table: "payments"
CREATE INDEX "idx_payments_subscription_id" ON "payments" ("subscription_id");
-- Create index "idx_payments_user_id" to table: "payments"
CREATE INDEX "idx_payments_user_id" ON "payments" ("user_id");
