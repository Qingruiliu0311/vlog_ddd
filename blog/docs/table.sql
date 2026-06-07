CREATE TABLE `blogs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(100) DEFAULT NULL,
  `content` text,
  `summary` varchar(255) DEFAULT NULL,
  `tag` longtext,
  `catelog` varchar(100) DEFAULT NULL,
  `stage` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_blogs_deleted_at` (`deleted_at`),
  KEY `idx_blogs_stage` (`stage`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;