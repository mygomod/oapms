-- @BeeOverwrite NO
-- @BeeGenerateTime 20200824_145130
CREATE TABLE role(
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `name` varchar(255) NOT NULL,
   `app_id` int(11) NOT NULL,
   `intro` varchar(255) NOT NULL,
   `menu_ids` json NOT NULL,
   `menu_ids_ele` longtext  NOT NULL,
    PRIMARY KEY (`id`)
)
