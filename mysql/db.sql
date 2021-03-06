DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(225) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `relationships`;
CREATE TABLE `relationships` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `first_email_id` INT DEFAULT NULL,
  `second_email_id` INT DEFAULT NULL,
  `status` INT NOT NULL,
  `action_user_id` INT DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;