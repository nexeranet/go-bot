create table users (
  id  integer primary key,
  tg_id integer unique
);

create table category(
  codename varchar(255) primary key,
  name varchar(255),
  tg_id integer,
  FOREIGN KEY(tg_id) REFERENCES users(tg_id)
);

create table alias(
  id integer primary key,
  text varchar(255),
  category_codename integer,
  tg_id integer,
  FOREIGN KEY(category_codename) REFERENCES category(codename),
  FOREIGN KEY(tg_id) REFERENCES users(tg_id)
);

create table expense(
  id integer primary key,
  amount integer,
  created datetime,
  category_codename integer,
  raw_text text,
  tg_id integer,
  FOREIGN KEY(category_codename) REFERENCES category(codename),
  FOREIGN KEY(tg_id) REFERENCES users(tg_id)
);

--  insert into category (codename, name, is_base_expense)
--  values
    --  ("products", "продукты", true),
    --  ("coffee", "кофе", true),
    --  ("dinner", "обед", true),
    --  ("cafe", "кафе", true),
    --  ("transport", "общ. транспорт", false),
    --  ("taxi", "такси", false),
    --  ("phone", "телефон", false),
    --  ("books", "книги", false),
    --  ("internet", "интернет", true),
    --  ("subscriptions", "подписки", false),
    --  ("other", "прочее", true);
--
--  insert into alias(category_codename, text)
--  values
    --  ("products", "продукты"),
    --  ("coffee", "кофе"),
    --  ("dinner", "обед"),
    --  ("cafe", "кафе"),
    --  ("transport", "общ. транспорт"),
    --  ("taxi", "такси"),
    --  ("phone", "телефон"),
    --  ("books", "книги"),
    --  ("internet", "интернет"),
    --  ("subscriptions", "подписки"),
    --  ("other", "прочее");
