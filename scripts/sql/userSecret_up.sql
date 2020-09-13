-- @BeeOverwrite YES
-- @BeeGenerateTime 20200831_101837
CREATE TABLE user_secret(
     
        
            `id` int(11) NOT NULL AUTO_INCREMENT,
        
     
        
            `uid` int(11) NOT NULL,
        
     
        
            `secret` varchar(255) NOT NULL,
        
     
        
            `is_bind` int(11) NOT NULL,
        
     
        
            `ctime` int(11) NOT NULL,
        
     
        
            `utime` int(11) NOT NULL,
        
     
     PRIMARY KEY (`id`)
)
