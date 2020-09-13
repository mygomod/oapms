-- @BeeOverwrite YES
-- @BeeGenerateTime 20200820_230345
CREATE TABLE access(
     
        
            `id` int(11) NOT NULL AUTO_INCREMENT,
        
     
        
            `client` varchar(255) NOT NULL,
        
     
        
            `authorize` varchar(255) NOT NULL,
        
     
        
            `previous` varchar(255) NOT NULL,
        
     
        
            `access_token` varchar(255) NOT NULL,
        
     
        
            `refresh_token` varchar(255) NOT NULL,
        
     
        
            `expires_in` int(11) NOT NULL,
        
     
        
            `scope` varchar(255) NOT NULL,
        
     
        
            `redirect_uri` varchar(255) NOT NULL,
        
     
        
            `extra` longtext  NOT NULL,
        
     
        
            `ctime` int(11) NOT NULL,
        
     
     PRIMARY KEY (`id`)
)
