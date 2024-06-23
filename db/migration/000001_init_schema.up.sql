create table "accounts" (
                            "id" bigserial PRIMARY KEY,
                            "owner" varchar NOT NULL,
                            "balance" bigint NOT NULL,
                            "currency" varchar NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

create table "entries" (
                           "id" bigserial PRIMARY KEY,
                           "account_id" bigint NOT NULL,
                           "amount" bigint NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now())
);

create table "transfers" (
                             "id" bigserial PRIMARY KEY,
                             "from_account_id" bigint,
                             "to_account_id" bigint,
                             "amount" bigint NOT NULL,
                             "created_at" timestamptz NOT NULL DEFAULT (now())
);

alter table "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
alter table "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");
alter table "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

create index on "accounts" ("owner");
create index on "entries" ("account_id");
create index on "transfers" ("from_account_id");
create index on "transfers" ("to_account_id");
create index on "transfers" ("from_account_id", "to_account_id");
