CREATE TABLE IF NOT EXISTS homepage_traduction (
       id SERIAL PRIMARY KEY,
       lang VARCHAR(2),
       title VARCHAR(800),

       section_1_headline VARCHAR(800),
       section_1_card_1_headline VARCHAR(800),
       section_1_card_1_text VARCHAR(800),
       section_1_card_2_headline VARCHAR(800),
       section_1_card_2_text VARCHAR(800),
       section_1_card_3_headline VARCHAR(800),
       section_1_card_3_text VARCHAR(800),

       section_2_headline VARCHAR(800),
       section_2_textblock_1 VARCHAR(800),
       section_2_textblock_2 VARCHAR(800),
       section_2_textblock_3 VARCHAR(800),

       section_3_headline VARCHAR(800),
       section_3_block_headline_backend VARCHAR(800),
       section_3_block_text_backend VARCHAR(800),
       section_3_block_headline_frontend VARCHAR(800),
       section_3_block_text_frontend VARCHAR(800),
       section_3_block_headline_creativedesign VARCHAR(800),
       section_3_block_text_creativedesign VARCHAR(800),
       section_3_block_headline_server_management VARCHAR(800),
       section_3_block_text_server_management VARCHAR(800),

       section_4_headline VARCHAR(800),
       section_4_textblock VARCHAR(800),

       footer_text_1 VARCHAR(800),
       footer_text_2 VARCHAR(800)
);