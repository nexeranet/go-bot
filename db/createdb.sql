create table budget(
    codename varchar(255) primary key,
    daily_limit integer
);

create table category(
    codename varchar(255) primary key,
    name varchar(255),
    is_base_expense boolean
);

--  create table alias(
  --  id integer primary key,
  --  text varchar(255),
  --  category_codename integer,
  --  FOREIGN KEY(category_codename) REFERENCES category(codename)
--  )
create table expense(
    id integer primary key,
    amount integer,
    created datetime,
    category_codename integer,
    raw_text text,
    FOREIGN KEY(category_codename) REFERENCES category(codename)
);

insert into category (codename, name, is_base_expense)
values
    ("products", "продукты", true),
    ("coffee", "кофе", true),
    ("dinner", "обед", true),
    ("cafe", "кафе", true),
    ("transport", "общ. транспорт", false),
    ("taxi", "такси", false),
    ("phone", "телефон", false),
    ("books", "книги", false),
    ("internet", "интернет", true),
    ("subscriptions", "подписки", false),
    ("other", "прочее", true);

insert into budget(codename, daily_limit) values ('base', 500);
