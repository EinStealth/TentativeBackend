-- CREATE DATABASE
DROP DATABASE IF EXISTS `tentative-db`;
CREATE DATABASE `tentative-db`;

-- map指定
USE `tentative-db`;

DROP TABLE IF EXISTS `rooms`;
CREATE TABLE `rooms` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `secret_words` varchar(10),
    `is_start` boolean,
    `deamon` int,
    PRIMARY KEY (`id`)
);

-- mock data
INSERT INTO `rooms` VALUES
(
    1, "test1", 0, 1
),
(
    2, "test2", 0, 1
);

DROP TABLE IF EXISTS `players`;
CREATE TABLE `players` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `secret_words` varchar(10),
    `name` varchar(10),
    `icon` int,
    `status` int,
    PRIMARY KEY (`id`)
);

-- mock data
INSERT INTO `players` VALUES
(
    1, "test1", "player01", 1, 1
),
(
    2, "test1", "player02", 2, 1
);

DROP TABLE IF EXISTS `locations`;
CREATE TABLE `locations` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `secret_words` varchar(10),
    `relative_time` varchar(10),
    `latitude` float,
    `longitude` float,
    `status` int,
    PRIMARY KEY (`id`)
);

-- mock data
INSERT INTO `locations` VALUES
(
    1, "test1", "12:00:00", 10.0, 10.0, 1
),
(
    2, "test1", "12:00:00", 10.1, 10.1, 2
);
