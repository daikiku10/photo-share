SET NAMES 'utf8mb4';
DROP DATABASE IF EXISTS photo_share;
CREATE DATABASE IF NOT EXISTS photo_share CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `photo_share`.`photos` (
    `id` char(26) not null COMMENT '主キー',
    `title` varchar(255) not null COMMENT '写真タイトル',
    PRIMARY KEY(`id`)
);