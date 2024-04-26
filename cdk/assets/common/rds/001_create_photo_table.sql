SET NAMES 'utf8mb4';
DROP DATABASE IF EXISTS photo_share;
CREATE DATABASE IF NOT EXISTS photo_share CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `photo_share`.`photos` (
    `id` char(26) not null COMMENT '主キー',
    `created_at` datetime not null default current_timestamp COMMENT '作成日時',
    `title` varchar(255) not null COMMENT '写真タイトル',
    `description` varchar(255) not null COMMENT '説明',
    `imageUrl` varchar(255) not null COMMENT '写真URL',
    `userId` char(26) not null COMMENT '投稿者ID',
    `categoryId` char(26) not null COMMENT 'カテゴリID',
    CONSTRAINT `Photo_pkey` PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='写真'
    ;