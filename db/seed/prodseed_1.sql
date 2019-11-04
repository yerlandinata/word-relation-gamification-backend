DELETE FROM public.gold_standard WHERE wp_id=684;
DELETE FROM public.gold_standard WHERE wp_id=2234;
DELETE FROM public.gold_standard WHERE wp_id=5429;
UPDATE public.gold_standard SET wrt_id=3 WHERE wp_id=809;
UPDATE public.gold_standard SET wrt_id=1 WHERE wp_id=3968;
INSERT INTO public.gold_standard (wp_id, wrt_id) VALUES (4976, 3);
INSERT INTO public.gold_standard (wp_id, wrt_id) VALUES (5230, 1);
