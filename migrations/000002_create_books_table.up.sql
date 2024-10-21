CREATE TABLE IF NOT EXISTS `books` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `title` VARCHAR(255) NOT NULL,
    `author` VARCHAR(255) NOT NULL,
    `isbn` VARCHAR(255) NOT NULL,
    `price` DECIMAL(10, 2) NOT NULL,

    INDEX `idx_author_title` (`author`, `title`),
    INDEX `idx_title` (`title`)
);
