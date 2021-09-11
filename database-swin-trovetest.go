package main

func createBiblioDB(env *Env) {

  statement_trove, err := env.db.Prepare("DROP TABLE stg_swin_trovetest")
  if err != nil {
    panic(err)
  }
  statement_trove.Exec()

  statement_swin, err := env.db.Prepare("CREATE TABLE IF NOT EXISTS stg_swin_trovetest (header_identifier TEXT, date_latest TEXT, metadata_identifier TEXT, metadata_identifier_file_id TEXT, url TEXT)")
  if err != nil {
    panic(err)
  }
  statement_swin.Exec()
}
