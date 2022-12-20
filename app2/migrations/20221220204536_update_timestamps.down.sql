ALTER TABLE
    budgets
ALTER COLUMN
    created_at DROP DEFAULT,
ALTER COLUMN
    updated_at DROP DEFAULT;

---------------------------------------------
ALTER TABLE
    categories
ALTER COLUMN
    created_at DROP DEFAULT,
ALTER COLUMN
    updated_at DROP DEFAULT;

---------------------------------------------
ALTER TABLE
    transactions
ALTER COLUMN
    created_at DROP DEFAULT,
ALTER COLUMN
    updated_at DROP DEFAULT;

---------------------------------------------
ALTER TABLE
    user_budgets
ALTER COLUMN
    created_at DROP DEFAULT,
ALTER COLUMN
    updated_at DROP DEFAULT;