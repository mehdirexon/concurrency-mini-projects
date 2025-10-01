-- +goose Up
ALTER TABLE ONLY public.plans
    ADD CONSTRAINT plans_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.user_plans
    ADD CONSTRAINT user_plans_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.user_plans
    ADD CONSTRAINT user_plans_plan_id_fkey FOREIGN KEY (plan_id) REFERENCES public.plans(id) ON UPDATE RESTRICT ON DELETE CASCADE;

ALTER TABLE ONLY public.user_plans
    ADD CONSTRAINT user_plans_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE CASCADE;

-- +goose Down
ALTER TABLE ONLY public.user_plans DROP CONSTRAINT IF EXISTS user_plans_plan_id_fkey;
ALTER TABLE ONLY public.user_plans DROP CONSTRAINT IF EXISTS user_plans_user_id_fkey;
ALTER TABLE ONLY public.user_plans DROP CONSTRAINT IF EXISTS user_plans_pkey;
ALTER TABLE ONLY public.users DROP CONSTRAINT IF EXISTS users_pkey;
ALTER TABLE ONLY public.plans DROP CONSTRAINT IF EXISTS plans_pkey;
