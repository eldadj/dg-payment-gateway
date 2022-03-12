--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.1

-- Started on 2022-03-11 15:41:02 WAT

-- SET statement_timeout = 0;
-- SET lock_timeout = 0;
-- SET idle_in_transaction_session_timeout = 0;
-- SET client_encoding = 'UTF8';
-- SET standard_conforming_strings = on;
-- SELECT pg_catalog.set_config('search_path', '', false);
-- SET check_function_bodies = false;
-- SET xmloption = content;
-- SET client_min_messages = warning;
-- SET row_security = off;

--
-- TOC entry 3387 (class 1262 OID 16384)
-- Name: dgpg; Type: DATABASE; Schema: -; Owner: dgpg
--

-- CREATE DATABASE dgpg WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


-- ALTER DATABASE dgpg OWNER TO dgpg;

-- \connect dgpg
--
-- SET statement_timeout = 0;
-- SET lock_timeout = 0;
-- SET idle_in_transaction_session_timeout = 0;
-- SET client_encoding = 'UTF8';
-- SET standard_conforming_strings = on;
-- SELECT pg_catalog.set_config('search_path', '', false);
-- SET check_function_bodies = false;
-- SET xmloption = content;
-- SET client_min_messages = warning;
-- SET row_security = off;
--
-- SET default_tablespace = '';
--
-- SET default_table_access_method = heap;
--
--
-- TOC entry 214 (class 1259 OID 16413)
-- Name: authorize; Type: TABLE; Schema: public; Owner: dgpg
--

CREATE TABLE public.authorize (
    authorize_id bigint NOT NULL,
    merchant_id bigint,
    credit_card_id bigint,
    currency text,
    amount numeric(18,2),
    date_in date DEFAULT CURRENT_DATE,
    time_in time with time zone DEFAULT CURRENT_TIME,
    authorize_code text,
    status text,
    has_refund boolean
);


ALTER TABLE public.authorize OWNER TO dgpg;

--
-- TOC entry 216 (class 1259 OID 16426)
-- Name: authorize_action; Type: TABLE; Schema: public; Owner: dgpg
--

CREATE TABLE public.authorize_action (
    authorize_action_id bigint NOT NULL,
    date_in date,
    time_in time with time zone,
    amount numeric(18,2),
    action_type text,
    authorize_id bigint
);


ALTER TABLE public.authorize_action OWNER TO dgpg;

--
-- TOC entry 215 (class 1259 OID 16425)
-- Name: authorize_action_authorize_action_id_seq; Type: SEQUENCE; Schema: public; Owner: dgpg
--

CREATE SEQUENCE public.authorize_action_authorize_action_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.authorize_action_authorize_action_id_seq OWNER TO dgpg;

--
-- TOC entry 3375 (class 0 OID 0)
-- Dependencies: 215
-- Name: authorize_action_authorize_action_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dgpg
--

ALTER SEQUENCE public.authorize_action_authorize_action_id_seq OWNED BY public.authorize_action.authorize_action_id;


--
-- TOC entry 218 (class 1259 OID 16438)
-- Name: authorize_action_refund; Type: VIEW; Schema: public; Owner: dgpg
--

CREATE VIEW public.authorize_action_refund AS
 SELECT aa.authorize_action_id,
    aa.date_in,
    aa.time_in,
    aa.amount,
    aa.authorize_id
   FROM public.authorize_action aa
  WHERE (lower(aa.action_type) = 'refund'::text);


ALTER TABLE public.authorize_action_refund OWNER TO dgpg;

--
-- TOC entry 219 (class 1259 OID 16442)
-- Name: authorize_action_void; Type: VIEW; Schema: public; Owner: dgpg
--

CREATE VIEW public.authorize_action_void AS
 SELECT aa.authorize_action_id,
    aa.date_in,
    aa.time_in,
    aa.authorize_id
   FROM public.authorize_action aa
  WHERE (lower(aa.action_type) = 'void'::text);


ALTER TABLE public.authorize_action_void OWNER TO dgpg;

--
-- TOC entry 213 (class 1259 OID 16412)
-- Name: authorize_authorize_id_seq; Type: SEQUENCE; Schema: public; Owner: dgpg
--

CREATE SEQUENCE public.authorize_authorize_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.authorize_authorize_id_seq OWNER TO dgpg;

--
-- TOC entry 3376 (class 0 OID 0)
-- Dependencies: 213
-- Name: authorize_authorize_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dgpg
--

ALTER SEQUENCE public.authorize_authorize_id_seq OWNED BY public.authorize.authorize_id;


--
-- TOC entry 221 (class 1259 OID 16456)
-- Name: capture; Type: TABLE; Schema: public; Owner: dgpg
--

CREATE TABLE public.capture (
    capture_id bigint NOT NULL,
    date_in date DEFAULT CURRENT_DATE,
    amount numeric(18,6),
    authorize_id bigint,
    time_in time with time zone DEFAULT CURRENT_TIME,
    refunded boolean DEFAULT false
);


ALTER TABLE public.capture OWNER TO dgpg;

--
-- TOC entry 220 (class 1259 OID 16455)
-- Name: capture_capture_id_seq; Type: SEQUENCE; Schema: public; Owner: dgpg
--

CREATE SEQUENCE public.capture_capture_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.capture_capture_id_seq OWNER TO dgpg;

--
-- TOC entry 3377 (class 0 OID 0)
-- Dependencies: 220
-- Name: capture_capture_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dgpg
--

ALTER SEQUENCE public.capture_capture_id_seq OWNED BY public.capture.capture_id;


--
-- TOC entry 212 (class 1259 OID 16404)
-- Name: credit_card; Type: TABLE; Schema: public; Owner: dgpg
--

CREATE TABLE public.credit_card (
    credit_card_id bigint NOT NULL,
    owner_name text,
    address text,
    card_no text,
    exp_month integer,
    exp_year integer,
    cvv text,
    currency_code text,
    current_amount numeric(18,6)
);


ALTER TABLE public.credit_card OWNER TO dgpg;

--
-- TOC entry 211 (class 1259 OID 16403)
-- Name: credit_card_credit_card_id_seq; Type: SEQUENCE; Schema: public; Owner: dgpg
--

CREATE SEQUENCE public.credit_card_credit_card_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.credit_card_credit_card_id_seq OWNER TO dgpg;

--
-- TOC entry 3378 (class 0 OID 0)
-- Dependencies: 211
-- Name: credit_card_credit_card_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dgpg
--

ALTER SEQUENCE public.credit_card_credit_card_id_seq OWNED BY public.credit_card.credit_card_id;


--
-- TOC entry 210 (class 1259 OID 16395)
-- Name: merchant; Type: TABLE; Schema: public; Owner: dgpg
--

CREATE TABLE public.merchant (
    merchant_id bigint NOT NULL,
    fullname text,
    user_name text,
    pwd_hash text
);


ALTER TABLE public.merchant OWNER TO dgpg;

--
-- TOC entry 209 (class 1259 OID 16394)
-- Name: merchant_merchant_id_seq; Type: SEQUENCE; Schema: public; Owner: dgpg
--

CREATE SEQUENCE public.merchant_merchant_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.merchant_merchant_id_seq OWNER TO dgpg;

--
-- TOC entry 3379 (class 0 OID 0)
-- Dependencies: 209
-- Name: merchant_merchant_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dgpg
--

ALTER SEQUENCE public.merchant_merchant_id_seq OWNED BY public.merchant.merchant_id;


--
-- TOC entry 223 (class 1259 OID 16466)
-- Name: refund; Type: TABLE; Schema: public; Owner: dgpg
--

CREATE TABLE public.refund (
    refund_id bigint NOT NULL,
    authorize_id bigint,
    date_in date DEFAULT CURRENT_DATE,
    time_in time with time zone DEFAULT CURRENT_TIME
);


ALTER TABLE public.refund OWNER TO dgpg;

--
-- TOC entry 222 (class 1259 OID 16465)
-- Name: refund_refund_id_seq; Type: SEQUENCE; Schema: public; Owner: dgpg
--

CREATE SEQUENCE public.refund_refund_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.refund_refund_id_seq OWNER TO dgpg;

--
-- TOC entry 3380 (class 0 OID 0)
-- Dependencies: 222
-- Name: refund_refund_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dgpg
--

ALTER SEQUENCE public.refund_refund_id_seq OWNED BY public.refund.refund_id;


--
-- TOC entry 217 (class 1259 OID 16434)
-- Name: vw_authorize_action_capture; Type: VIEW; Schema: public; Owner: dgpg
--

CREATE VIEW public.vw_authorize_action_capture AS
 SELECT aa.authorize_action_id,
    aa.date_in,
    aa.time_in,
    aa.amount,
    aa.authorize_id
   FROM public.authorize_action aa
  WHERE (lower(aa.action_type) = 'capture'::text);


ALTER TABLE public.vw_authorize_action_capture OWNER TO dgpg;

--
-- TOC entry 3205 (class 2604 OID 16416)
-- Name: authorize authorize_id; Type: DEFAULT; Schema: public; Owner: dgpg
--

ALTER TABLE ONLY public.authorize ALTER COLUMN authorize_id SET DEFAULT nextval('public.authorize_authorize_id_seq'::regclass);


--
-- TOC entry 3208 (class 2604 OID 16429)
-- Name: authorize_action authorize_action_id; Type: DEFAULT; Schema: public; Owner: dgpg
--

ALTER TABLE ONLY public.authorize_action ALTER COLUMN authorize_action_id SET DEFAULT nextval('public.authorize_action_authorize_action_id_seq'::regclass);


--
-- TOC entry 3209 (class 2604 OID 16459)
-- Name: capture capture_id; Type: DEFAULT; Schema: public; Owner: dgpg
--

ALTER TABLE ONLY public.capture ALTER COLUMN capture_id SET DEFAULT nextval('public.capture_capture_id_seq'::regclass);


--
-- TOC entry 3204 (class 2604 OID 16407)
-- Name: credit_card credit_card_id; Type: DEFAULT; Schema: public; Owner: dgpg
--

ALTER TABLE ONLY public.credit_card ALTER COLUMN credit_card_id SET DEFAULT nextval('public.credit_card_credit_card_id_seq'::regclass);


--
-- TOC entry 3203 (class 2604 OID 16398)
-- Name: merchant merchant_id; Type: DEFAULT; Schema: public; Owner: dgpg
--

ALTER TABLE ONLY public.merchant ALTER COLUMN merchant_id SET DEFAULT nextval('public.merchant_merchant_id_seq'::regclass);


--
-- TOC entry 3213 (class 2604 OID 16469)
-- Name: refund refund_id; Type: DEFAULT; Schema: public; Owner: dgpg
--

ALTER TABLE ONLY public.refund ALTER COLUMN refund_id SET DEFAULT nextval('public.refund_refund_id_seq'::regclass);


--
-- TOC entry 3223 (class 2606 OID 16433)
-- Name: authorize_action authorize_action_pkey; Type: CONSTRAINT; Schema: public; Owner: dgpg
--

ALTER TABLE ONLY public.authorize_action
    ADD CONSTRAINT authorize_action_pkey PRIMARY KEY (authorize_action_id);


--
-- TOC entry 3221 (class 2606 OID 16424)
-- Name: authorize authorize_pkey; Type: CONSTRAINT; Schema: public; Owner: dgpg
--

ALTER TABLE ONLY public.authorize
    ADD CONSTRAINT authorize_pkey PRIMARY KEY (authorize_id);


--
-- TOC entry 3225 (class 2606 OID 16463)
-- Name: capture capture_pk; Type: CONSTRAINT; Schema: public; Owner: dgpg
--

ALTER TABLE ONLY public.capture
    ADD CONSTRAINT capture_pk PRIMARY KEY (capture_id);


--
-- TOC entry 3219 (class 2606 OID 16411)
-- Name: credit_card credit_card_pkey; Type: CONSTRAINT; Schema: public; Owner: dgpg
--

ALTER TABLE ONLY public.credit_card
    ADD CONSTRAINT credit_card_pkey PRIMARY KEY (credit_card_id);


--
-- TOC entry 3217 (class 2606 OID 16402)
-- Name: merchant merchant_pkey; Type: CONSTRAINT; Schema: public; Owner: dgpg
--

ALTER TABLE ONLY public.merchant
    ADD CONSTRAINT merchant_pkey PRIMARY KEY (merchant_id);


--
-- TOC entry 3227 (class 2606 OID 16473)
-- Name: refund refund_pkey; Type: CONSTRAINT; Schema: public; Owner: dgpg
--

ALTER TABLE ONLY public.refund
    ADD CONSTRAINT refund_pkey PRIMARY KEY (refund_id);


-- Completed on 2022-03-11 15:41:02 WAT

--
-- PostgreSQL database dump complete
--

--
-- TOC entry 3371 (class 0 OID 16395)
-- Dependencies: 210
-- Data for Name: merchant; Type: TABLE DATA; Schema: public; Owner: dgpg
--

COPY public.merchant (merchant_id, fullname, user_name, pwd_hash) FROM stdin;
1	Merchant 1	m1	$2a$12$EVduAs9Kht1z2OBwwbXdQe3fQUuvQP5t6TaYJoGV0vWNQB7dki9cu
2	Merchant 2	m2	$2a$12$ppjqU9V7/niX3pMRsO.hmeiDEjAyj8BW.wc2aNxkegUHG5oxbYkuG
3	Merchant 1	m3	$2a$12$y5Kl1kvzNxEbykjOIvc4yuIX8mUoBIpMLl0Dng7zc.JxqS0uLNj7u
4	Merchant 1	m4	$2a$12$3ocxQ77Yoo3Z799IQ/Rv/OZAVwVX6h6xYjWL24iSYkCz4uU/hcK2W
\.

--
-- TOC entry 3381 (class 0 OID 16466)
-- Dependencies: 223
-- Data for Name: refund; Type: TABLE DATA; Schema: public; Owner: dgpg
--