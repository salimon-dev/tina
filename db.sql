--
-- PostgreSQL database dump
--

-- Dumped from database version 17.5
-- Dumped by pg_dump version 17.0

-- Started on 2025-07-03 10:16:22 EEST

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 221 (class 1259 OID 16420)
-- Name: access_keys; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.access_keys (
    id uuid NOT NULL,
    user_id uuid,
    name character varying(32),
    value character varying(32),
    created_at bigint,
    updated_at bigint
);


--
-- TOC entry 219 (class 1259 OID 16406)
-- Name: invitations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.invitations (
    id uuid NOT NULL,
    created_by uuid NOT NULL,
    code character varying(16) NOT NULL,
    usage_remaining numeric NOT NULL,
    expires_at bigint,
    status numeric,
    created_at bigint,
    updated_at bigint
);


--
-- TOC entry 224 (class 1259 OID 16435)
-- Name: messages; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.messages (
    id uuid NOT NULL,
    body text,
    user_id uuid,
    thread_id uuid,
    type numeric,
    created_at bigint,
    updated_at bigint
);


--
-- TOC entry 218 (class 1259 OID 16401)
-- Name: permissions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.permissions (
    id uuid NOT NULL,
    user_id uuid,
    permission character varying(32),
    created_at bigint,
    updated_at bigint
);


--
-- TOC entry 223 (class 1259 OID 16430)
-- Name: thread_members; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.thread_members (
    id uuid NOT NULL,
    user_id uuid,
    thread_id uuid,
    created_at bigint,
    updated_at bigint
);


--
-- TOC entry 222 (class 1259 OID 16425)
-- Name: threads; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.threads (
    id uuid NOT NULL,
    name character varying(64),
    created_at bigint,
    updated_at bigint
);


--
-- TOC entry 220 (class 1259 OID 16415)
-- Name: transactions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.transactions (
    id uuid NOT NULL,
    source_id uuid,
    target_id uuid,
    category character varying(32),
    description character varying(256),
    amount bigint,
    fee bigint,
    status smallint,
    created_at bigint,
    updated_at bigint
);


--
-- TOC entry 217 (class 1259 OID 16392)
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    username character varying(32) NOT NULL,
    password character varying(32),
    invitation_id uuid,
    description character varying(256),
    credit bigint,
    score bigint,
    status numeric,
    secret_key character varying(64) NOT NULL,
    is_public boolean,
    hook_url character varying(256),
    registered_at bigint,
    updated_at bigint
);


--
-- TOC entry 3311 (class 2606 OID 16424)
-- Name: access_keys access_keys_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.access_keys
    ADD CONSTRAINT access_keys_pkey PRIMARY KEY (id);


--
-- TOC entry 3305 (class 2606 OID 16412)
-- Name: invitations invitations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.invitations
    ADD CONSTRAINT invitations_pkey PRIMARY KEY (id);


--
-- TOC entry 3317 (class 2606 OID 16441)
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (id);


--
-- TOC entry 3303 (class 2606 OID 16405)
-- Name: permissions permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permissions_pkey PRIMARY KEY (id);


--
-- TOC entry 3315 (class 2606 OID 16434)
-- Name: thread_members thread_members_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.thread_members
    ADD CONSTRAINT thread_members_pkey PRIMARY KEY (id);


--
-- TOC entry 3313 (class 2606 OID 16429)
-- Name: threads threads_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.threads
    ADD CONSTRAINT threads_pkey PRIMARY KEY (id);


--
-- TOC entry 3309 (class 2606 OID 16419)
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- TOC entry 3307 (class 2606 OID 16414)
-- Name: invitations uni_invitations_code; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.invitations
    ADD CONSTRAINT uni_invitations_code UNIQUE (code);


--
-- TOC entry 3299 (class 2606 OID 16400)
-- Name: users uni_users_username; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_username UNIQUE (username);


--
-- TOC entry 3301 (class 2606 OID 16398)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


-- Completed on 2025-07-03 10:16:26 EEST

--
-- PostgreSQL database dump complete
--

