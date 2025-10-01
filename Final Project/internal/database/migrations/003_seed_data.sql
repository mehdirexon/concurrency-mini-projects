-- +goose Up
INSERT INTO public.users(email, first_name, last_name, password, user_active, is_admin, created_at, updated_at)
VALUES
    ('admin@example.com','Admin','User','$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe',1,1,'2022-03-14 00:00:00','2022-03-14 00:00:00');

SELECT pg_catalog.setval('public.plans_id_seq', 1, false);
SELECT pg_catalog.setval('public.user_id_seq', 2, true);
SELECT pg_catalog.setval('public.user_plans_id_seq', 1, false);

INSERT INTO public.plans(plan_name, plan_amount, created_at, updated_at)
VALUES
    ('Bronze Plan',1000,'2022-05-12 00:00:00','2022-05-12 00:00:00'),
    ('Silver Plan',2000,'2022-05-12 00:00:00','2022-05-12 00:00:00'),
    ('Gold Plan',3000,'2022-05-12 00:00:00','2022-05-12 00:00:00');

-- +goose Down
DELETE FROM public.user_plans;
DELETE FROM public.users WHERE email = 'admin@example.com';
DELETE FROM public.plans WHERE plan_name IN ('Bronze Plan','Silver Plan','Gold Plan');

-- reset sequences back to 1
ALTER SEQUENCE public.plans_id_seq RESTART WITH 1;
ALTER SEQUENCE public.user_id_seq RESTART WITH 1;
ALTER SEQUENCE public.user_plans_id_seq RESTART WITH 1;
