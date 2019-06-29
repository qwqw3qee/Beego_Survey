START TRANSACTION;
DROP database IF EXISTS `beego_survey`;
create database beego_survey DEFAULT CHARACTER
SET utf8;
use beego_survey;
CREATE TABLE IF NOT EXISTS `users` (
  `id` INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
  `user_name` varchar (32) NOT NULL UNIQUE,
  `nick_name` varchar (32) NOT NULL,
  `user_type` INTEGER NOT NULL DEFAULT 1 CHECK (user_type in (0, 1, 2)),
  `password` varchar (256) NOT NULL DEFAULT "",
  `img_name` varchar (256) DEFAULT "",
  `register_time` datetime NOT NULL,
  `last_login_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
);
CREATE TABLE IF NOT EXISTS `papers` (
  `paper_id` INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT UNIQUE,
  `author_id` INTEGER NOT NULL,
  `type` INTEGER NOT NULL DEFAULT 0 CHECK (type in (0, 1)),
  `title` varchar (256) DEFAULT "",
  `subtitle` varchar (256) DEFAULT "",
  `description` varchar (5000) DEFAULT "",
  `add_date` datetime NOT NULL,
  `start_date` datetime NOT NULL,
  `end_date` datetime NOT NULL,
  `time_limit` INTEGER NOT NULL DEFAULT 0,
  `remark` varchar (5000) DEFAULT "",
  `img_name` varchar (256) DEFAULT "",
  FOREIGN KEY (`author_id`) REFERENCES `users` (`id`) ON
DELETE CASCADE ON
UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS `questions` (
  `question_id` INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT UNIQUE,
  `paper_id` INTEGER NOT NULL,
  `no` INTEGER DEFAULT 0 CHECK (no >= 0),
  `type` INTEGER NOT NULL DEFAULT 0 CHECK (type in (0, 1, 2)),
  `description` varchar (5000) DEFAULT "",
  `options_num` INTEGER NOT NULL DEFAULT 0,
  `options_info` varchar (5000) DEFAULT "",
  `remark` varchar (5000) DEFAULT "",
  FOREIGN KEY (`paper_id`) REFERENCES `papers` (`paper_id`) ON
DELETE CASCADE ON
UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS `records` (
  `record_id` INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT UNIQUE,
  `paper_id` INTEGER NOT NULL,
  `answer_user_id` INTEGER NOT NULL,
  `answer_date` datetime NOT NULL,
  `use_time` INTEGER NOT NULL DEFAULT 0,
  `scores` INTEGER DEFAULT 0,
  FOREIGN KEY (`answer_user_id`) REFERENCES `users` (`id`) ON
DELETE CASCADE ON
UPDATE CASCADE,
  FOREIGN KEY(`paper_id`) REFERENCES `papers` (`paper_id`) ON
DELETE CASCADE ON
UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS `answers` (
  `id` INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT UNIQUE,
  `record_id` INTEGER NOT NULL,
  `paper_id` INTEGER NOT NULL,
  `question_id` INTEGER NOT NULL,
  `answer_no` INTEGER DEFAULT 0,
  `answer_text` varchar (5000) DEFAULT "",
  `scores` INTEGER DEFAULT 0,
  FOREIGN KEY (`paper_id`) REFERENCES `papers` (`paper_id`) ON
DELETE CASCADE ON
UPDATE CASCADE,
  FOREIGN KEY(`question_id`) REFERENCES `questions` (`question_id`) ON
DELETE CASCADE ON
UPDATE CASCADE,
  FOREIGN KEY(`record_id`) REFERENCES `records` (`record_id`) ON
DELETE CASCADE ON
UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS `standards` (
  `id` INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT UNIQUE,
  `paper_id` INTEGER NOT NULL,
  `question_id` INTEGER NOT NULL,
  `standard_no` INTEGER DEFAULT 0,
  `standard_text` varchar (5000) DEFAULT "",
  `scores` INTEGER NOT NULL DEFAULT 0,
  FOREIGN KEY (`question_id`) REFERENCES `questions` (`question_id`) ON
DELETE CASCADE ON
UPDATE CASCADE,
  FOREIGN KEY(`paper_id`) REFERENCES `papers` (`paper_id`) ON
DELETE CASCADE ON
UPDATE CASCADE
);
INSERT INTO
  `beego_survey`.`users` (
    `user_name`,
    `nick_name`,
    `user_type`,
    `password`,
    `register_time`,
    `last_login_time`
  )
VALUES
  (
    'anonymous',
    '匿名',
    '2',
    '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918',
    now(),
    now()
  );
INSERT INTO
  `beego_survey`.`users` (
    `user_name`,
    `nick_name`,
    `user_type`,
    `password`,
    `register_time`,
    `last_login_time`
  )
VALUES
  (
    'admin',
    '管理员',
    '0',
    '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918',
    now(),
    now()
  );
COMMIT;