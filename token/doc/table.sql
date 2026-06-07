CREATE TABLE `tokens` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ref_user_id` longtext,
  `access_token` varchar(191) DEFAULT NULL,
  `issued_at` datetime(3) DEFAULT NULL,
  `access_token_expired_at` datetime(3) DEFAULT NULL,
  `refresh_token` varchar(191) DEFAULT NULL,
  `refresh_token_expired_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_tokens_access_token` (`access_token`),
  UNIQUE KEY `uni_tokens_refresh_token` (`refresh_token`),
  KEY `idx_tokens_access_token` (`access_token`),
  KEY `idx_tokens_refresh_token` (`refresh_token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;