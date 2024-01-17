package homepage

import (
	"cms/internal/database/postgres"
	"cms/internal/privates/structure"
)

func All() ([]HomepageModel, error) {
	var sql = `
	SELECT
	    title,
		section_1_headline,
		section_1_card_1_headline,
		section_1_card_1_text,
		section_1_card_2_headline,
		section_1_card_2_text,
		section_1_card_3_headline,
		section_1_card_3_text,
		section_2_headline,
		section_2_textblock_1,
		section_2_textblock_2,
		section_2_textblock_3,
		section_3_headline,
		section_3_block_headline_backend,
		section_3_block_text_backend,
		section_3_block_headline_frontend,
		section_3_block_text_frontend,
		section_3_block_headline_creativedesign,
		section_3_block_text_creativedesign,
		section_3_block_headline_server_management,
		section_3_block_text_server_management,
		section_4_headline,
		section_4_textblock,
		footer_text_1,
		footer_text_2
	FROM homepage_traduction
`
	conn, err := postgres.Connection()
	if err != nil {
		return nil, err
	}

	query, err := conn.Query(sql)
	if err != nil {
		return nil, err
	}

	var trads []HomepageModel
	var buffer HomepageModel

	for query.Next() {
		println(query.Columns())
		err := query.Scan(
			&buffer.Title,
			&buffer.Section_1_headline,
			&buffer.Section_1_card_1_headline,
			&buffer.Section_1_card_1_text,
			&buffer.Section_1_card_2_headline,
			&buffer.Section_1_card_2_text,
			&buffer.Section_1_card_3_headline,
			&buffer.Section_1_card_3_text,
			&buffer.Section_2_headline,
			&buffer.Section_2_textblock_1,
			&buffer.Section_2_textblock_2,
			&buffer.Section_2_textblock_3,
			&buffer.Section_3_headline,
			&buffer.Section_3_block_headline_backend,
			&buffer.Section_3_block_text_backend,
			&buffer.Section_3_block_headline_frontend,
			&buffer.Section_3_block_text_frontend,
			&buffer.Section_3_block_headline_creativedesign,
			&buffer.Section_3_block_text_creativedesign,
			&buffer.Section_3_block_headline_server_management,
			&buffer.Section_3_block_text_server_management,
			&buffer.Section_4_headline,
			&buffer.Section_4_textblock,
			&buffer.Footer_text_1,
			&buffer.Footer_text_2,
		)
		if err != nil {
			return nil, err
		}
		trads = append(trads, buffer)
	}
	conn.Close()

	return trads, nil
}

func New(lang string) (*HomepageModel, error) {
	var sql = `
	SELECT
	    title,
		section_1_headline,
		section_1_card_1_headline,
		section_1_card_1_text,
		section_1_card_2_headline,
		section_1_card_2_text,
		section_1_card_3_headline,
		section_1_card_3_text,
		section_2_headline,
		section_2_textblock_1,
		section_2_textblock_2,
		section_2_textblock_3,
		section_3_headline,
		section_3_block_headline_backend,
		section_3_block_text_backend,
		section_3_block_headline_frontend,
		section_3_block_text_frontend,
		section_3_block_headline_creativedesign,
		section_3_block_text_creativedesign,
		section_3_block_headline_server_management,
		section_3_block_text_server_management,
		section_4_headline,
		section_4_textblock,
		footer_text_1,
		footer_text_2
	FROM homepage_traduction
	WHERE lang=$1;
`
	conn, err := postgres.Connection()
	if err != nil {
		return nil, err
	}

	var buffer HomepageModel
	err = conn.QueryRow(sql, lang).Scan(
		&buffer.Title,
		&buffer.Section_1_headline,
		&buffer.Section_1_card_1_headline,
		&buffer.Section_1_card_1_text,
		&buffer.Section_1_card_2_headline,
		&buffer.Section_1_card_2_text,
		&buffer.Section_1_card_3_headline,
		&buffer.Section_1_card_3_text,
		&buffer.Section_2_headline,
		&buffer.Section_2_textblock_1,
		&buffer.Section_2_textblock_2,
		&buffer.Section_2_textblock_3,
		&buffer.Section_3_headline,
		&buffer.Section_3_block_headline_backend,
		&buffer.Section_3_block_text_backend,
		&buffer.Section_3_block_headline_frontend,
		&buffer.Section_3_block_text_frontend,
		&buffer.Section_3_block_headline_creativedesign,
		&buffer.Section_3_block_text_creativedesign,
		&buffer.Section_3_block_headline_server_management,
		&buffer.Section_3_block_text_server_management,
		&buffer.Section_4_headline,
		&buffer.Section_4_textblock,
		&buffer.Footer_text_1,
		&buffer.Footer_text_2,
	)
	if err != nil {
		return nil, err
	}
	conn.Close()

	return &buffer, nil
}

func (m HomepageModel) Create(lang string) error {
	var sql = `
		INSERT INTO 
		    homepage_traduction(
			title,
			section_1_headline,
			section_1_card_1_headline,
			section_1_card_1_text,
			section_1_card_2_headline,
			section_1_card_2_text,
			section_1_card_3_headline,
			section_1_card_3_text,
			section_2_headline,
			section_2_textblock_1,
			section_2_textblock_2,
			section_2_textblock_3,
			section_3_headline,
			section_3_block_headline_backend,
			section_3_block_text_backend,
			section_3_block_headline_frontend,
			section_3_block_text_frontend,
			section_3_block_headline_creativedesign,
			section_3_block_text_creativedesign,
			section_3_block_headline_server_management,
			section_3_block_text_server_management,
			section_4_headline,
			section_4_textblock,
			footer_text_1,
			footer_text_2,
			lang
		) VALUES (
			 $1,
			 $2,
			 $3,
			 $4,
			 $5,
			 $6,
			 $7,
			 $8,
			 $9,
			 $10,
			 $11,
			 $12,
			 $13,
			 $14,
			 $15,
			 $16,
			 $17,
			 $18,
			 $19,
			 $20,
			 $21,
			 $22,
			 $23,
			 $24,
			 $25,
			 $26
		);
	`

	conn, err := postgres.Connection()
	if err != nil {
		return err
	}

	var querys []any = structure.StructToArray(m)
	querys = append(querys, lang)

	_, err = conn.Exec(sql, querys...)
	if err != nil {
		return err
	}

	return nil
}

func (m HomepageModel) Update(lang string) error {
	var sql = `
		UPDATE
		    homepage_traduction
		SET
		title = $1,
		section_1_headline = $2,
		section_1_card_1_headline = $3,
		section_1_card_1_text = $4,
		section_1_card_2_headline = $5,
		section_1_card_2_text = $6,
		section_1_card_3_headline = $7,
		section_1_card_3_text = $8,
		section_2_headline = $9,
		section_2_textblock_1 = $10,
		section_2_textblock_2 = $11,
		section_2_textblock_3 = $12,
		section_3_headline = $13,
		section_3_block_headline_backend = $14,
		section_3_block_text_backend = $15,
		section_3_block_headline_frontend = $16,
		section_3_block_text_frontend = $17,
		section_3_block_headline_creativedesign = $18,
		section_3_block_text_creativedesign = $19,
		section_3_block_headline_server_management = $20,
		section_3_block_text_server_management = $21,
		section_4_headline = $22,
		section_4_textblock = $23,
		footer_text_1 = $24,
		footer_text_2 = $25
		WHERE lang = $26
	`

	conn, err := postgres.Connection()
	if err != nil {
		return err
	}

	var querys []any = structure.StructToArray(m)
	querys = append(querys, lang)
	for i := range querys {
		println(i)
	}

	_, err = conn.Exec(sql, querys...)
	if err != nil {
		return err
	}

	return nil
}

func Delete(lang string) error {
	var sql = `DELETE FROM homepage_traduction WHERE lang=$1`

	conn, err := postgres.Connection()
	if err != nil {
		return err
	}

	println(lang)

	_, err = conn.Exec(sql, lang)
	if err != nil {
		return err
	}

	return nil
}
