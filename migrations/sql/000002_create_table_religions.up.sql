-- eventori.model_schedules definition

CREATE TABLE IF NOT EXISTS `model_schedules` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `model_id` int(11) NOT NULL,
  `schedule_date` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `model_schedules_FK` (`model_id`),
  CONSTRAINT `model_schedules_FK` FOREIGN KEY (`model_id`) REFERENCES `model_catalogues` (`id`)
)