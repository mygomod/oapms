-- @BeeOverwrite NO
-- @BeeGenerateTime 20200831_224648
CREATE TABLE menu(


            `id` int(11) NOT NULL AUTO_INCREMENT,



            `pid` int(11) NOT NULL,



            `app_id` int(11) NOT NULL,



            `name` varchar(255) NOT NULL,



            `path` varchar(255) NOT NULL,



            `pms_code` varchar(255) NOT NULL,


            `menu_type` int(11) NOT NULL,



            `icon` varchar(255) NOT NULL,



            `order_num` int(11) NOT NULL,



            `ctime` int(11) NOT NULL,



            `utime` int(11) NOT NULL,



            `state` int(11) NOT NULL,


     PRIMARY KEY (`id`)
)
