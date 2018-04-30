Create Table if not exists clients(
    id integer,
    name char,
    email char,
    phone char,
    sms_notification boolean,
    email_notification boolean,
    PRIMARY KEY (ID)
);
Create Table if not exists accounts(
    id integer,
    client_id integer,
    balance integer,
    PRIMARY KEY (ID)
);
Create Table if not exists transactions(
    id integer,
    from_account integer,
    to_account integer,
    amount integer,
    created_at integer,
    add_withdrow boolean,
    PRIMARY KEY (ID)
);
Create Table if not exists transaction_logs(
    id integer,
    transaction_type integer,
    request text,
    PRIMARY KEY (ID)
);
