CREATE TABLE sdn (
                       uid int NOT NULL,
                       firstname text NOT NULL,
                       lastname text NOT NULL,
                       publish timestamp,
                       CONSTRAINT "pk_user_uid" PRIMARY KEY (uid)
);