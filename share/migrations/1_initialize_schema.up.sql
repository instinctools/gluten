CREATE SEQUENCE executions_results_id_seq;

CREATE TABLE public.executions_results
(
  id integer NOT NULL DEFAULT nextval('executions_results_id_seq'::regclass),
  created bigint,
  CONSTRAINT executions_results_pkey PRIMARY KEY (id)
);

CREATE TABLE public.metrics
(
  execution_result_id integer,
  key text,
  value text,
  CONSTRAINT metrics_execution_result_id_executions_results_id_foreign FOREIGN KEY (execution_result_id)
      REFERENCES public.executions_results (id) MATCH SIMPLE
      ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE SEQUENCE executions_id_seq;

CREATE TABLE public.executions
(
  id integer NOT NULL DEFAULT nextval('executions_id_seq'::regclass),
  created bigint,
  parameters text,
  result_id integer,
  CONSTRAINT executions_pkey PRIMARY KEY (id),
  CONSTRAINT executions_result_id_executions_results_id_foreign FOREIGN KEY (result_id)
      REFERENCES public.executions_results (id) MATCH SIMPLE
      ON UPDATE CASCADE ON DELETE CASCADE
);

COMMIT;