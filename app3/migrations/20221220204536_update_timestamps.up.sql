ALTER TABLE
    budgets
ALTER COLUMN
    created_at
SET
    DEFAULT now(),
ALTER COLUMN
    updated_at
SET
    DEFAULT now();

---------------------------------------------
ALTER TABLE
    categories
ALTER COLUMN
    created_at
SET
    DEFAULT now(),
ALTER COLUMN
    updated_at
SET
    DEFAULT now();

---------------------------------------------
ALTER TABLE
    transactions
ALTER COLUMN
    created_at
SET
    DEFAULT now(),
ALTER COLUMN
    updated_at
SET
    DEFAULT now();

---------------------------------------------
ALTER TABLE
    user_budgets
ALTER COLUMN
    created_at
SET
    DEFAULT now(),
ALTER COLUMN
    updated_at
SET
    DEFAULT now();