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
    PRIMARY KEY (`id`)
);

-- mock data
INSERT INTO `rooms` VALUES
(
    1, "test1", true
),
(
    2, "test2", false
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