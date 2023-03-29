psql -d webapp < sql/webapp.sql
cd data/location/bin
sh run.sh
cd ../../profile/bin
sh run.sh