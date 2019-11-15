--
-- PostgreSQL database dump
--

-- Dumped from database version 10.10
-- Dumped by pg_dump version 10.10

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- Name: adminpack; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION adminpack; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: annotation; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.annotation (
    wp_id integer NOT NULL,
    wrt_id integer,
    player_id bigint NOT NULL,
    player_time_ms integer NOT NULL,
    created_at timestamp with time zone,
    is_valid boolean NOT NULL
);


ALTER TABLE public.annotation OWNER TO postgres;

--
-- Name: gold_standard; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.gold_standard (
    wp_id integer NOT NULL,
    wrt_id integer
);


ALTER TABLE public.gold_standard OWNER TO postgres;

--
-- Name: player; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.player (
    id bigint NOT NULL,
    birth_date integer,
    display_name character varying(64) NOT NULL,
    campaign_source character varying(64) NOT NULL,
    score integer NOT NULL,
    annotation_count integer NOT NULL,
    elapsed integer NOT NULL,
    game_level integer NOT NULL,
    full_name character varying(64),
    onboarding_time_ms integer
);


ALTER TABLE public.player OWNER TO postgres;

--
-- Name: word_pair; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.word_pair (
    id integer NOT NULL,
    word_1 character varying(64) NOT NULL,
    word_2 character varying(64) NOT NULL,
    word_1_freq integer,
    word_2_freq integer,
    active_status integer NOT NULL,
    pmi double precision NOT NULL
);


ALTER TABLE public.word_pair OWNER TO postgres;

--
-- Name: word_pair_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.word_pair_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.word_pair_id_seq OWNER TO postgres;

--
-- Name: word_pair_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.word_pair_id_seq OWNED BY public.word_pair.id;


--
-- Name: word_relation_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.word_relation_type (
    id integer NOT NULL,
    short_desc character varying(16) NOT NULL,
    human_desc character varying(32) NOT NULL
);


ALTER TABLE public.word_relation_type OWNER TO postgres;

--
-- Name: word_relation_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.word_relation_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.word_relation_type_id_seq OWNER TO postgres;

--
-- Name: word_relation_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.word_relation_type_id_seq OWNED BY public.word_relation_type.id;


--
-- Name: word_pair id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.word_pair ALTER COLUMN id SET DEFAULT nextval('public.word_pair_id_seq'::regclass);


--
-- Name: word_relation_type id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.word_relation_type ALTER COLUMN id SET DEFAULT nextval('public.word_relation_type_id_seq'::regclass);


--
-- Data for Name: annotation; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.annotation (wp_id, wrt_id, player_id, player_time_ms, created_at, is_valid) FROM stdin;
\.


--
-- Data for Name: gold_standard; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.gold_standard (wp_id, wrt_id) FROM stdin;
\.


--
-- Data for Name: player; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.player (id, birth_date, display_name, campaign_source, score, annotation_count, elapsed, game_level, full_name, onboarding_time_ms) FROM stdin;
\.


--
-- Data for Name: word_pair; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.word_pair (id, word_1, word_2, word_1_freq, word_2_freq, active_status, pmi) FROM stdin;
54	muson	perubahan	255	14117	1	0
80	sarung	kain	420	3207	1	7.0474132128236366
107	neutrino	partikel	123	2300	1	8.3003877051837289
301	stasiun_televisi	stasiun	3044	29256	1	1.9624629877528454
472	pepatah	ayat	221	5493	1	2.7408176425545272
582	topik	acara	1863	28224	1	1.9585806960575312
614	model	televisi	14468	22447	1	1.7389442523803615
744	intuisi	kompleks	131	6719	1	3.0801249920992513
809	sabuk	bahan	1261	10642	1	2.5639417881521327
869	substansi	hakikat	530	528	1	6.8979699118153777
901	organ	alat	1727	11951	1	3.0338800982989951
933	mayoritas	gereja	7212	31500	1	0.3462137835808623
1223	ranjau_darat	alat	56	11951	1	5.1449834783532804
1233	bunga	grup	11926	24103	1	0
1246	serambi_jantung	ruang	2	12648	1	8.5942539944604217
1291	deisme	ajaran	79	5300	1	0
1299	gabardin	kain	5	3207	1	11.478144519207627
1473	rotan_dahanan	palem	13	340	1	10.882658705448259
1554	friksi	gaya	73	10891	1	6.9258212546020594
1618	jangka	alat	1942	11951	1	0
1727	pengumpat	orang	2	118113	1	0
1758	belacan	bumbu	17	1747	1	9.6390781203576612
1796	wilayah	daerah	83313	53990	1	1.8452888418700153
1886	sepak_bola	olahraga	42096	10151	1	3.5366076847031391
2022	cerita	adaptasi	14553	2950	1	3.8495588795223119
2027	ikan	vertebrata	10300	684	1	4.4369485582373605
2121	basil	bakteri	329	2781	1	4.7399638483209667
2139	suara	gelombang	18526	4499	1	3.7638097001744337
2199	ingatan	fungsi	759	8325	1	0
2295	aorta	arteri	78	302	1	12.162010810102537
2324	mikroskop_elektron	mikroskop	49	307	1	11.772895774794572
2351	bandongan	pola	44	4209	1	5.2459769133275191
2370	mega	jaringan	700	17954	1	0
2484	bitumen	cairan	27	2472	1	0
2539	kolibri	burung	52	9739	1	9.0034424702881228
2785	aluminium	unsur	922	6065	1	3.7757910989065233
2797	kurator	pengurus	407	2234	1	0
2840	torium	aktinida	91	74	1	12.546007972439996
2951	ancaman	usaha	3478	10869	1	0
2998	bahasa	himpunan	71150	1237	1	0
3016	jasa	aktivitas	5541	7128	1	1.6505183077221952
3073	anjing	mamalia	3608	1224	1	3.4480884592818963
3132	janturan	wacana	14	762	1	9.6928677626090867
3154	tawes	ikan	25	10300	1	7.1017249456269127
3448	konjungtiva	lapisan_tipis	17	100	1	12.227538637003601
3450	embargo	pelarangan	142	525	1	0
3820	karbon_dioksida	pemadam	495	51	1	0
3885	imlek	religi	637	557	1	5.4535860649177632
3901	pasung	rangka	38	5349	1	5.0693910882578272
4056	rotan_dahanan	rotan	13	581	1	11.116553235121531
4105	kapitalisme	manusia	736	29736	1	2.3947291800610899
4164	asidosis	proses	36	29983	1	3.8591759902998111
4235	laboratorium	tempat	3313	4426	1	0.63650326701993931
4469	garam	sumber	2673	14585	1	1.1359836389400459
4606	badik	senjata	104	8532	1	6.5893578126764369
4635	gula	bahan	2531	10642	1	3.7711028926969097
4645	lutung	kera	165	571	1	9.2499769112180594
4657	laras_bahasa	bentuk_bahasa	15	119	1	13.207763068242278
4717	helikopter	alat	2009	11951	1	1.2028237065188971
4805	beras_kencur	minuman	12	3481	1	0
4856	zina	perbuatan	128	2112	1	0
4904	pulung	kekuatan	172	14634	1	1.567393071712782
4976	kapal	donasi	18712	322	1	1.1384286285028988
5032	tanjung	daratan	6911	4774	1	0.99951079680143173
5035	belut	sumber	251	14585	1	2.05121428095556
5061	konsumer	organisme	39	1943	1	6.578702574274172
5098	kerabu	subang	16	1027	1	8.9122446499648298
5115	sasando	alat	43	11951	1	6.7695997132408081
5135	zodiak	sabuk	308	1261	1	5.4451468593657264
5196	sanggraloka	tempat	19	4426	1	0
5230	nasi	makanan	1927	18596	1	6.5600709404576767
5235	huruf	grafem	6377	31	1	7.7058839096563352
41	agronomi	ilmu_terapan	103	68	1	10.128455311241435
118	tanda_petik	tanda_baca	15	180	1	12.755190551467637
5286	sufisme	kekuatan	90	14634	1	2.6523460734090074
478	novel	narasi	11653	803	1	3.4389732197862073
493	lenong	kesenian	189	4465	1	4.0774329485147023
1328	revitalisasi	proses	193	29983	1	2.613289866296427
1654	egoisme	motivasi	67	842	1	8.031733674538966
1903	tanin	senyawa	115	4512	1	7.5203186154268806
2003	siaran_pers	tulisan	88	5893	1	0
2100	alu	alat	194	11951	1	4.4182840705362292
2386	pertumbuhan_ekonomi	proses	847	29983	1	2.2090887749501551
2519	perang_sipil	kekuatan	421	14634	1	1.360342958601545
2580	usia_sekolah	tahap	122	6643	1	3.2270271616029707
3035	liberalisme	paham	350	1446	1	7.8295918673524909
3118	jantung	organ	2422	1727	1	5.5355378066957686
3155	paus	kepala	11105	20548	1	1.0284337051726375
3403	prefiks	afiks	134	42	1	11.627985274606607
3968	besi_tempa	bahan	42	10642	1	6.3312233941337981
4361	modal	setoran	2110	41	1	10.302142454026349
5468	turunan	objek	1594	12800	1	0
1870	amilopektin	polisakarida	24	125	1	13.154240252484
2344	modal_tetap	modal	10	2110	1	10.543481788760822
2851	sekretaris_jenderal	jabatan_struktural	1362	56	1	6.5448319684205281
3203	lisimeter	alat_ukur	1	106	1	16.802696719089258
3479	manusia	setumpuk	29736	62	1	0
4563	kapal_tunda	kapal	48	18712	1	6.6467641103352539
4649	superkonduktivitas	fenomena	50	2678	1	7.817439922613846
5269	jagal	perjalanan	202	14168	1	0
\.


--
-- Data for Name: word_relation_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.word_relation_type (id, short_desc, human_desc) FROM stdin;
2	synonymy	%a sama artinya dengan %b
1	hyponymy	%a adalah sejenis %b
3	unrelated	
4	notsure	
\.


--
-- Name: word_pair_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.word_pair_id_seq', 1, false);


--
-- Name: word_relation_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.word_relation_type_id_seq', 1, false);


--
-- Name: annotation annotation_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.annotation
    ADD CONSTRAINT annotation_pkey PRIMARY KEY (wp_id, player_id);


--
-- Name: gold_standard gold_standard_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gold_standard
    ADD CONSTRAINT gold_standard_pkey PRIMARY KEY (wp_id);


--
-- Name: player player_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player
    ADD CONSTRAINT player_pkey PRIMARY KEY (id);


--
-- Name: word_pair word_pair_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.word_pair
    ADD CONSTRAINT word_pair_pkey PRIMARY KEY (id);


--
-- Name: word_relation_type word_relation_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.word_relation_type
    ADD CONSTRAINT word_relation_type_pkey PRIMARY KEY (id);


--
-- Name: annotation_player_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX annotation_player_id ON public.annotation USING btree (player_id);


--
-- Name: annotation annotation_player_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.annotation
    ADD CONSTRAINT annotation_player_id_fkey FOREIGN KEY (player_id) REFERENCES public.player(id) ON UPDATE CASCADE;


--
-- Name: annotation annotation_wp_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.annotation
    ADD CONSTRAINT annotation_wp_id_fkey FOREIGN KEY (wp_id) REFERENCES public.word_pair(id);


--
-- Name: annotation annotation_wrt_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.annotation
    ADD CONSTRAINT annotation_wrt_id_fkey FOREIGN KEY (wrt_id) REFERENCES public.word_relation_type(id);


--
-- Name: gold_standard gold_standard_wp_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gold_standard
    ADD CONSTRAINT gold_standard_wp_id_fkey FOREIGN KEY (wp_id) REFERENCES public.word_pair(id);


--
-- Name: gold_standard gold_standard_wrt_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gold_standard
    ADD CONSTRAINT gold_standard_wrt_id_fkey FOREIGN KEY (wrt_id) REFERENCES public.word_relation_type(id);


--
-- PostgreSQL database dump complete
--

