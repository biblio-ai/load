package main

func createBiblioDB(env *Env) {
  /*
  statement_entity, err := env.db.Prepare("DROP TABLE stg_slv_primo")
  if err != nil {
    panic(err)
  }
  statement_entity.Exec()
  */
  statement, err := env.db.Prepare("CREATE TABLE IF NOT EXISTS stg_uni_melb_library (header_identifier TEXT, date_latest TEXT, metadata_identifier TEXT, metadata_identifier_handle_id TEXT, metadata_identifier_cms_id TEXT, metadata_identifier_accession_id TEXT, metadata_identifier_file_id TEXT, url TEXT, operation_id TEXT, status TEXT)")
  if err != nil {
    panic(err)
  }
  statement.Exec()

}
