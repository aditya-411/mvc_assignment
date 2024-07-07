CREATE TABLE `transactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) DEFAULT NULL,
  `issued_at` date DEFAULT NULL,
  `returned_at` date DEFAULT NULL,
  `title` varchar(200) DEFAULT NULL,
  `request_status` enum('-1','0','1') NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;