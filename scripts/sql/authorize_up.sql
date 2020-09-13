-- @BeeOverwrite YES
-- @BeeGenerateTime 20200820_230345
CREATE TABLE authorize(
     
        
            `id` int(11) NOT NULL AUTO_INCREMENT,
        
     
        
            `client` varchar(255) NOT NULL,
        
     
        
            `code` varchar(255) NOT NULL,
        
     
        
            `expires_in` int(11) NOT NULL,
        
     
        
            `scope` varchar(255) NOT NULL,
        
     
        
            `redirect_uri` varchar(255) NOT NULL,
        
     
        
            `state` varchar(255) NOT NULL,
        
     
        
            `extra` longtext  NOT NULL,
        
     
        
            `ctime` int(11) NOT NULL,
        
     
     PRIMARY KEY (`id`)
)
