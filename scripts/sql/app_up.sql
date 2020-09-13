-- @BeeOverwrite YES
-- @BeeGenerateTime 20200831_174645
CREATE TABLE app(
     
        
            `aid` int(11) NOT NULL AUTO_INCREMENT,
        
     
        
            `client_id` varchar(255) NOT NULL,
        
     
        
            `name` varchar(255) NOT NULL,
        
     
        
            `secret` varchar(255) NOT NULL,
        
     
        
            `redirect_uri` varchar(255) NOT NULL,
        
     
        
            `url` varchar(255) NOT NULL,
        
     
        
            `extra` longtext  NOT NULL,
        
     
        
            `call_no` int(11) NOT NULL,
        
     
        
            `state` int(11) NOT NULL,
        
     
        
            `ctime` int(11) NOT NULL,
        
     
        
            `utime` int(11) NOT NULL,
        
     
        
            `dtime` int(11) NOT NULL,
        
     
     PRIMARY KEY (`aid`)
)
