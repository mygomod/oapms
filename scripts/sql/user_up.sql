-- @BeeOverwrite NO
-- @BeeGenerateTime 20200820_230345
CREATE TABLE user(


            `uid` int(11) NOT NULL AUTO_INCREMENT,



            `nickname` varchar(255) NOT NULL,



            `email` varchar(255) NOT NULL,



            `avatar` varchar(255) NOT NULL,



            `password` varchar(500) NOT NULL,



            `state` int(11) NOT NULL,



            `gender` int(11) NOT NULL,



            `birthday` int(11) NOT NULL,



            `ctime` int(11) NOT NULL,



            `utime` int(11) NOT NULL,



            `last_login_ip` varchar(255) NOT NULL,



            `last_login_time` int(11) NOT NULL,


     PRIMARY KEY (`uid`)
)
