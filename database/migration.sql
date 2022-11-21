
CREATE TABLE `users`
(
    id   bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `users` (`name`, `status`)
VALUES ('Solomon', 0),
       ('Menelik', 0);



CREATE TABLE `todos`
(
    id   bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`),
    content text NOT NULL,
    status int NOT NULL
);

INSERT INTO `users` (`name`)
VALUES ('Solomon'),
       ('Menelik');

INSERT INTO `todos` (`name`, `content`, `status`)
VALUES ('Solomon', 'Shopping', 0),
       ('Solomon', 'Laundry', 0);