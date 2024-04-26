SET NAMES 'utf8mb4';
DROP DATABASE IF EXISTS photo_share;
CREATE DATABASE IF NOT EXISTS photo_share CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreatePhotos
CREATE TABLE `photo_share`.`photos` (
    `id` char(26) not null COMMENT '主キー',
    `created_at` datetime not null default current_timestamp COMMENT '作成日時',
    `title` varchar(255) not null COMMENT '写真タイトル',
    `description` varchar(255) not null COMMENT '説明',
    `imageUrl` varchar(255) not null COMMENT '写真URL',
    `authorId` char(26) not null COMMENT '投稿者ID',
    `categoryId` char(26) not null COMMENT 'カテゴリID',
    CONSTRAINT `Photo_pkey` PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='写真'
    ;

-- CreateCategory
CREATE TABLE `photo_share`.`category` (
    `id` char(26) not null COMMENT '主キー',
    `created_at` datetime not null default current_timestamp COMMENT '作成日時',
    `name` varchar(255) not null COMMENT 'カテゴリ名',
    `label` varchar(255) not null COMMENT 'ラベル',
    `description` varchar(255) not null COMMENT '説明',
    `imageUrl` varchar(255) not null COMMENT '写真URL',
    CONSTRAINT `Category_pkey` PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='カテゴリ'
    ;

-- CreateLike
CREATE TABLE `photo_share`.`like` (
    `id` char(26) not null COMMENT '主キー',
    `created_at` datetime not null default current_timestamp COMMENT '作成日時',
    `userId` char(26) not null COMMENT 'ユーザーID',
    `photoId` char(26) not null COMMENT '写真ID',
    CONSTRAINT `Like_pkey` PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='カテゴリ'
    ;

-- CreateComment
CREATE TABLE `photo_share`.`comment` (
    `id` char(26) not null COMMENT '主キー',
    `created_at` datetime not null default current_timestamp COMMENT '作成日時',
    `authorId` char(26) not null COMMENT '投稿者ID',
    `photoId` char(26) not null COMMENT '写真ID',
    `comment` varchar(255) not null COMMENT 'コメント',
    CONSTRAINT `Like_pkey` PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='カテゴリ'
    ;
