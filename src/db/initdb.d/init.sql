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