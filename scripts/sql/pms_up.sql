-- @BeeOverwrite YES
-- @BeeGenerateTime 20200820_230345
CREATE TABLE pms(
     
        
            `id` int(11) NOT NULL AUTO_INCREMENT,
        
     
        
            `app_id` int(11) NOT NULL,
        
     
        
            `pid` int(11) NOT NULL,
        
     
        
            `name` varchar(255) NOT NULL,
        
     
        
            `pms_code` varchar(255) NOT NULL,
        
     
        
            `pms_rule` varchar(255) NOT NULL,
        
     
        
            `pms_type` int(11) NOT NULL,
        
     
        
            `order_num` int(11) NOT NULL,
        
     
        
            `intro` varchar(255) NOT NULL,
        
     
        
            `ctime` int(11) NOT NULL,
        
     
        
            `utime` int(11) NOT NULL,
        
     
     PRIMARY KEY (`id`)
)
