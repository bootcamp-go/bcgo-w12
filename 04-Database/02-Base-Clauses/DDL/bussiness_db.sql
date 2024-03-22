-- DDL (Data Definition Language)
DROP DATABASE IF EXISTS `bussiness_db`;
CREATE DATABASE `bussiness_db`;
USE `bussiness_db`;

CREATE TABLE `departments` (
	`id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    `location` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `employees` (
	`id` INT NOT NULL AUTO_INCREMENT,
    `first_name` VARCHAR(100) NOT NULL,
    `last_name` VARCHAR(100) NOT NULL,
    `position` VARCHAR (150) NOT NULL,
    `department_id` INT NOT NULL,
    PRIMARY KEY(`id`),
    KEY `idx_employee_department_id` (`department_id`),
    CONSTRAINT `fk_employees_department_id` FOREIGN KEY (`department_id`) REFERENCES `departments` (`id`) 
);