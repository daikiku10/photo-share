SET NAMES 'utf8mb4';
DROP DATABASE IF EXISTS appTemplateDataBase;
CREATE DATABASE IF NOT EXISTS appTemplateDataBase CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE appTemplateDataBase;
DROP TABLE IF EXISTS appTemplateTable;
CREATE TABLE IF NOT EXISTS appTemplateTable(
        ID char(26) NOT NULL,
        Name varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
        PRIMARY KEY (ID)
    );