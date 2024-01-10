CREATE TABLE `quotes` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `service` longtext,
  `deadline` longtext,
  `price` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;