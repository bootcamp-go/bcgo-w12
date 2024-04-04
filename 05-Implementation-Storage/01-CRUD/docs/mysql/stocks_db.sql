-- DDL
DROP DATABASE IF EXISTS `stocks_db`;

CREATE DATABASE `stocks_db`;

USE `stocks_db`;

-- Table structure for table `stocks`
CREATE TABLE `stocks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `symbol` varchar(6) NOT NULL,
  `name` varchar(255) NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY `idx_stocks_symbol` (`symbol`),
  PRIMARY KEY (`id`)
);