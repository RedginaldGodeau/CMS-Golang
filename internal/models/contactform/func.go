package contactform

import (
	"cms/internal/database/postgres"
	"cms/internal/privates/structure"
)

func (c ContactFormModel) Create() error {
	var sql = `
		INSERT INTO 
		    contact_form(
			username,
		    email,
		    message,
		    clientIp,
		    sendAt
		) VALUES (
			 $1,
			 $2,
			 $3,
			 $4,
		     $5
		);
	`

	conn, err := postgres.Connection()
	if err != nil {
		return err
	}

	var querys []any = structure.StructToArray(c)
	_, err = conn.Exec(sql, querys...)
	if err != nil {
		return err
	}

	return nil
}
