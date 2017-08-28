CREATE TABLE public.execution(
  id SERIAL PRIMARY KEY NOT NULL,
  created BIGINT NOT NULL,
  parameters VARCHAR(255) NOT NULL
);

CREATE TABLE public.execution_result(
  id SERIAL PRIMARY KEY NOT NULL,
  created BIGINT NOT NULL,
  execution_id INT NOT NULL REFERENCES execution(id)
  ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE public.metric(
  key VARCHAR(100) NOT NULL,
  value VARCHAR(100) NOT NULL,
  execution_result_id INT NOT NULL REFERENCES execution_result(id)
  ON UPDATE CASCADE ON DELETE CASCADE
);