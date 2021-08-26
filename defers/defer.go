package defers

import "database/sql"

// Defer was made good in 1.16 (or 1.15?). Before that, it was pretty slow. So now, use it.
// A common error made with defers is to not check the error inside. For example :

func deferFuncClose(db *sql.DB) error {
	rows, err := db.Query("SELECT things FROM stuff")
	if err != nil {
		return err
	}
	// Here we want to defer, but rows.Close implements io.Closer, and that specifies it can return an error. So :
	defer rows.Close()
	// This defer is highlighted on Goland as "Unhandled error". Because it is!

	// A better way to do it :
	defer func() {
		// Notice the ":=" here, we're creating a new err variable (shadowing previous ones)
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}()
	// And if you don't want to panic, you can assign to an error declared outside of the closure function.

	defer func() {
		// Now we use "=", we assign to err.
		// We assign this way to allow the code inbetween the return and this defer
		// to still use 'err' if necessary.
		tmperr := rows.Close()
		if tmperr != nil {
			err = tmperr
		}
	}()

	// Careful to not overwrite err here! =/

	// This returns whatever is deferred!
	return err
}
