-- Active: 1709309448667@@127.0.0.1@3306
create schema if not exists puuclocks;

use puuclocks;

DROP TABLE IF EXISTS `account`;

CREATE TABLE `account` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `nickname` VARCHAR(255) NOT NULL,
    `password_hash` VARCHAR(255) NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `clock`;
CREATE TABLE `clock` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `image` LONGTEXT NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
DROP TABLE IF EXISTS `card`;
CREATE TABLE `card` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `time` float NOT NULL,
    `clock_id` bigint(20) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`clock_id`) REFERENCES clock(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

