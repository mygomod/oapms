-- @BeeOverwrite YES
-- @BeeGenerateTime 20200902_214035
CREATE TABLE menu_pms(
     
        
            `id` int(11) NOT NULL AUTO_INCREMENT,
        
     
        
            `pms_code` varchar(255) NOT NULL,
        
     
        
            `key` varchar(255) NOT NULL,
        
     
        
            `app_id` int(11) NOT NULL,
        
     
        
            `ctime` int(11) NOT NULL,
        
     
        
            `utime` int(11) NOT NULL,
        
     
     PRIMARY KEY (`id`)
)
