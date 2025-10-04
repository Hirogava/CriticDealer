CREATE TYPE weather AS ENUM (
  'Clear',
  'Clouds',
  'Rain',
  'Snow',
  'Drizzle',
  'Thunderstorm',
  'Mist',
  'Fog',
  'Haze',
  'Ice',
  'Freezing_Rain'
);

CREATE TYPE day_type AS ENUM (
  'Weekday',
  'Weekend'
);

CREATE TABLE accident (
  id serial PRIMARY KEY,
  movement_id integer,
  dtp_time integer,
  month integer,
  day_type day_type,
  weather_id integer
);

CREATE TABLE global_accident_statistic (
    dtp_count integer,
    dtp_koef decimal(10, 2),
    region varchar(255)
);

CREATE TABLE weather (
  id serial PRIMARY KEY,
  weather_type weather
);

ALTER TABLE accident ADD FOREIGN KEY (weather_id) REFERENCES weather (id);
