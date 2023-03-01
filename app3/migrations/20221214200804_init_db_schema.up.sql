CREATE TABLE budgets(
    id bigserial not null primary key,
    name varchar(255) not null,
    description varchar(255) not null,
    amount decimal(10, 2) not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

CREATE TABLE categories(
    id bigserial not null primary key,
    name varchar(255) not null,
    description varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

CREATE TABLE transactions(
    id bigserial not null primary key,
    title varchar(255) not null,
    budget_id bigint not null references budgets(id),
    category_id bigint not null references categories(id),
    amount decimal(10, 2) not null,
    type varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

CREATE TABLE user_budgets(
    id bigserial not null primary key,
    user_id bigint not null references "users"(id),
    budget_id bigint not null references budgets(id),
    ownership boolean not null,
    readonly boolean not null,
    created_at timestamp not null,
    updated_at timestamp not null
);