CREATE TABLE categories(
/* NAME        TYPE               SPECIFICATION(s) */
  id          INT PRIMARY KEY    NOT NULL,
  uid         INT                NOT NULL,
  category    TEXT               NOT NULL
);

CREATE TABLE dms(
/* NAME        TYPE               SPECIFICATION(s) */
  id          INT PRIMARY KEY    NOT NULL,
  uid         INT                NOT NULL,
  username    TEXT               NOT NULL
);

CREATE TABLE logs(
/* NAME        TYPE               SPECIFICATION(s) */
  id          INT PRIMARY KEY    NOT NULL,
  uid         INT                NOT NULL,
  log         TEXT               NOT NULL
);

CREATE TABLE messages(
/* NAME        TYPE               SPECIFICATION(s) */
  id          INT PRIMARY KEY    NULL,
  uid         INT                NOT NULL,
  sender      TEXT               NOT NULL,
  recipient   TEXT               NOT NULL
);

CREATE TABLE users(
/* NAME        TYPE               SPECIFICATION(s) */
  id          INT PRIMARY KEY    NOT NULL,
  mid         INT                NOT NULL,
  fullname    TEXT               NOT NULL,
  username    TEXT               NOT NULL,
  email       TEXT,
  password    TEXT
);
