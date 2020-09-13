-- @BeeOverwrite YES
-- @BeeGenerateTime 20200820_195417
CREATE TABLE department(
     
        
            `id` int(11) NOT NULL AUTO_INCREMENT,
        
     
        
            `name` varchar(255) NOT NULL,
        
     
        
            `pid` int(11) NOT NULL,
        
     
        
            `order_num` int(11) NOT NULL,
        
     
        
            `extend_field` longtext  NOT NULL,
        
     
        
            `intro` varchar(255) NOT NULL,
        
     
        
            `created_at` int(11) NOT NULL,
        
     
        
            `updated_at` int(11) NOT NULL,
        
     
     PRIMARY KEY (`id`)
)
